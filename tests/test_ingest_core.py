from __future__ import annotations

import tempfile
import unittest
from datetime import UTC, datetime, timedelta, timezone
from pathlib import Path

from app.codex_dashboard.aggregation import (
    build_buckets,
    build_project_stacks,
    event_metric_tokens,
    is_over_redline,
    project_weekly_burn,
)
from app.codex_dashboard.config import DashboardConfig
from app.codex_dashboard.models import SessionContextMarker, TokenEvent
from app.codex_dashboard.scanner import ingest_once
from app.codex_dashboard.storage import (
    connect,
    count_events,
    initialize_db,
    load_cursor,
    load_session_context_markers,
)


TOKEN_EVENT_LINE = (
    b'{"timestamp":"2026-04-03T00:09:11.080Z","type":"event_msg","payload":{"type":"token_count",'
    b'"info":{"total_token_usage":{"total_tokens":7242153},"last_token_usage":{"input_tokens":100970,'
    b'"cached_input_tokens":97920,"output_tokens":398,"reasoning_output_tokens":30,"total_tokens":101368}},'
    b'"rate_limits":{"secondary":{"used_percent":42.0,"window_minutes":10080,"resets_at":1775638824}}}}'
)


class IngestCoreTests(unittest.TestCase):
    def setUp(self) -> None:
        self.temp_dir = tempfile.TemporaryDirectory()
        self.root = Path(self.temp_dir.name)
        self.codex_root = self.root / ".codex"
        self.session_dir = self.codex_root / "sessions" / "2026" / "04" / "03"
        self.session_dir.mkdir(parents=True, exist_ok=True)
        self.db_path = self.root / "dashboard.db"
        self.connection = connect(self.db_path)
        initialize_db(self.connection)
        self.config = DashboardConfig(
            codex_root=str(self.codex_root),
            db_path=str(self.db_path),
            weekly_budget_tokens=8_000_000,
        )

    def tearDown(self) -> None:
        self.connection.close()
        self.temp_dir.cleanup()

    def test_ingest_skips_incomplete_trailing_line_until_next_scan(self) -> None:
        session_file = self.session_dir / "rollout-test.jsonl"
        session_file.write_bytes(TOKEN_EVENT_LINE + b"\n" + TOKEN_EVENT_LINE[:50])

        first_run = ingest_once(self.connection, self.config)
        self.assertEqual(first_run.events_ingested, 1)
        self.assertEqual(count_events(self.connection), 1)

        last_offset, _ = load_cursor(self.connection, str(session_file))
        self.assertLess(last_offset, session_file.stat().st_size)

        with session_file.open("ab") as handle:
            handle.write(TOKEN_EVENT_LINE[50:] + b"\n")

        second_run = ingest_once(self.connection, self.config)
        self.assertEqual(second_run.events_ingested, 1)
        self.assertEqual(count_events(self.connection), 2)

    def test_ingest_dedupes_same_line_offset(self) -> None:
        session_file = self.session_dir / "rollout-dedupe.jsonl"
        session_file.write_bytes(TOKEN_EVENT_LINE + b"\n")

        ingest_once(self.connection, self.config)
        ingest_once(self.connection, self.config)

        self.assertEqual(count_events(self.connection), 1)

    def test_aggregation_builds_expected_buckets(self) -> None:
        base = datetime(2026, 4, 3, 0, 0, tzinfo=UTC)
        events = [
            TokenEvent(
                session_path="a",
                line_offset=0,
                event_timestamp=base,
                total_tokens=100,
                input_tokens=0,
                cached_input_tokens=0,
                output_tokens=0,
                reasoning_output_tokens=0,
                cumulative_total_tokens=100,
                weekly_used_percent=None,
                weekly_window_minutes=None,
                weekly_resets_at=None,
                raw_json="{}",
            ),
            TokenEvent(
                session_path="a",
                line_offset=1,
                event_timestamp=base.replace(minute=4),
                total_tokens=200,
                input_tokens=0,
                cached_input_tokens=0,
                output_tokens=0,
                reasoning_output_tokens=0,
                cumulative_total_tokens=300,
                weekly_used_percent=None,
                weekly_window_minutes=None,
                weekly_resets_at=None,
                raw_json="{}",
            ),
        ]
        buckets = build_buckets(events, "5m", bucket_count=2, now=base.replace(minute=5))
        self.assertEqual(len(buckets), 2)
        self.assertEqual(buckets[0].total_tokens, 300)
        self.assertEqual(buckets[1].total_tokens, 0)

    def test_event_metric_tokens_supports_normalized_mode(self) -> None:
        event = TokenEvent(
            session_path="a",
            line_offset=0,
            event_timestamp=datetime(2026, 4, 3, 0, 0, tzinfo=UTC),
            total_tokens=999,
            input_tokens=1_000,
            cached_input_tokens=900,
            output_tokens=10,
            reasoning_output_tokens=5,
            cumulative_total_tokens=999,
            weekly_used_percent=None,
            weekly_window_minutes=None,
            weekly_resets_at=None,
            raw_json="{}",
        )

        self.assertEqual(event_metric_tokens(event, "total"), 999)
        self.assertEqual(event_metric_tokens(event, "norm"), 280)

    def test_aggregation_builds_normalized_buckets(self) -> None:
        base = datetime(2026, 4, 3, 0, 0, tzinfo=UTC)
        events = [
            TokenEvent(
                session_path="a",
                line_offset=0,
                event_timestamp=base,
                total_tokens=500,
                input_tokens=200,
                cached_input_tokens=100,
                output_tokens=10,
                reasoning_output_tokens=0,
                cumulative_total_tokens=500,
                weekly_used_percent=None,
                weekly_window_minutes=None,
                weekly_resets_at=None,
                raw_json="{}",
            ),
            TokenEvent(
                session_path="a",
                line_offset=1,
                event_timestamp=base.replace(minute=4),
                total_tokens=600,
                input_tokens=300,
                cached_input_tokens=100,
                output_tokens=0,
                reasoning_output_tokens=5,
                cumulative_total_tokens=1_100,
                weekly_used_percent=None,
                weekly_window_minutes=None,
                weekly_resets_at=None,
                raw_json="{}",
            ),
        ]

        buckets = build_buckets(events, "5m", bucket_count=2, now=base.replace(minute=5), metric_mode="norm")

        self.assertEqual(len(buckets), 2)
        self.assertEqual(buckets[0].total_tokens, 410)
        self.assertEqual(buckets[1].total_tokens, 0)

    def test_daily_buckets_align_to_local_midnight(self) -> None:
        est = timezone(timedelta(hours=-5), name="EST")
        now = datetime(2026, 4, 3, 1, 30, tzinfo=est)
        event_local = datetime(2026, 4, 2, 23, 45, tzinfo=est)
        events = [
            TokenEvent(
                session_path="a",
                line_offset=0,
                event_timestamp=event_local.astimezone(UTC),
                total_tokens=250,
                input_tokens=0,
                cached_input_tokens=0,
                output_tokens=0,
                reasoning_output_tokens=0,
                cumulative_total_tokens=250,
                weekly_used_percent=None,
                weekly_window_minutes=None,
                weekly_resets_at=None,
                raw_json="{}",
            )
        ]

        buckets = build_buckets(events, "1d", bucket_count=2, now=now, display_tz=est)

        self.assertEqual(buckets[0].start_at, datetime(2026, 4, 2, 0, 0, tzinfo=est))
        self.assertEqual(buckets[1].start_at, datetime(2026, 4, 3, 0, 0, tzinfo=est))
        self.assertEqual(buckets[0].total_tokens, 250)
        self.assertEqual(buckets[1].total_tokens, 0)

    def test_redline_projection_uses_interval_rate(self) -> None:
        projected = project_weekly_burn(10_000, 3600)
        self.assertGreater(projected, 1_000_000)
        self.assertTrue(is_over_redline(10_000, 3600, 1_500_000))
        self.assertFalse(is_over_redline(500, 3600, 8_000_000))

    def test_ingest_ignores_token_event_with_null_info(self) -> None:
        session_file = self.session_dir / "rollout-null-info.jsonl"
        session_file.write_text(
            '{"timestamp":"2026-04-03T00:09:11.080Z","type":"event_msg","payload":{"type":"token_count","info":null}}\n',
            encoding="utf-8",
        )

        ingest_result = ingest_once(self.connection, self.config)

        self.assertEqual(ingest_result.events_ingested, 0)
        self.assertEqual(count_events(self.connection), 0)

    def test_ingest_backfills_session_context_markers_from_session_meta(self) -> None:
        repo_root = self.root / "RepoAlpha"
        (repo_root / ".git").mkdir(parents=True, exist_ok=True)
        session_file = self.session_dir / "rollout-meta.jsonl"
        session_file.write_text(
            (
                '{"timestamp":"2026-04-03T00:00:00.000Z","type":"session_meta",'
                f'"payload":{{"cwd":"{repo_root.as_posix()}"}}}}\n'
            )
            + TOKEN_EVENT_LINE.decode("utf-8")
            + "\n",
            encoding="utf-8",
        )

        ingest_once(self.connection, self.config)

        markers = load_session_context_markers(self.connection, [str(session_file)])
        self.assertIn(str(session_file), markers)
        self.assertEqual(markers[str(session_file)][0].project_label, "RepoAlpha")
        self.assertEqual(markers[str(session_file)][0].project_source, "repo")

    def test_build_project_stacks_uses_latest_context_marker_before_event(self) -> None:
        base = datetime(2026, 4, 3, 0, 0, tzinfo=UTC)
        events = [
            TokenEvent(
                session_path="a",
                line_offset=5,
                event_timestamp=base,
                total_tokens=100,
                input_tokens=0,
                cached_input_tokens=0,
                output_tokens=0,
                reasoning_output_tokens=0,
                cumulative_total_tokens=100,
                weekly_used_percent=None,
                weekly_window_minutes=None,
                weekly_resets_at=None,
                raw_json="{}",
            ),
            TokenEvent(
                session_path="a",
                line_offset=15,
                event_timestamp=base,
                total_tokens=200,
                input_tokens=0,
                cached_input_tokens=0,
                output_tokens=0,
                reasoning_output_tokens=0,
                cumulative_total_tokens=300,
                weekly_used_percent=None,
                weekly_window_minutes=None,
                weekly_resets_at=None,
                raw_json="{}",
            ),
        ]
        markers = {
            "a": [
                SessionContextMarker("a", 0, "C:/RepoA", "C:/RepoA", "RepoA", "repo"),
                SessionContextMarker("a", 10, "C:/RepoB", "C:/RepoB", "RepoB", "repo"),
            ]
        }

        buckets, legend, repo_totals = build_project_stacks(
            events,
            markers,
            "1h",
            bucket_count=1,
            now=base,
            top_n=5,
        )

        self.assertEqual(len(buckets), 1)
        self.assertEqual(dict(legend), {"C:/RepoB": "RepoB", "C:/RepoA": "RepoA"})
        self.assertEqual(repo_totals[0]["C:/RepoA"], 100)
        self.assertEqual(repo_totals[0]["C:/RepoB"], 200)

    def test_build_project_stacks_supports_normalized_mode(self) -> None:
        base = datetime(2026, 4, 3, 0, 0, tzinfo=UTC)
        events = [
            TokenEvent(
                session_path="a",
                line_offset=1,
                event_timestamp=base,
                total_tokens=1_000,
                input_tokens=400,
                cached_input_tokens=300,
                output_tokens=5,
                reasoning_output_tokens=0,
                cumulative_total_tokens=1_000,
                weekly_used_percent=None,
                weekly_window_minutes=None,
                weekly_resets_at=None,
                raw_json="{}",
            ),
            TokenEvent(
                session_path="b",
                line_offset=1,
                event_timestamp=base,
                total_tokens=500,
                input_tokens=90,
                cached_input_tokens=10,
                output_tokens=20,
                reasoning_output_tokens=0,
                cumulative_total_tokens=500,
                weekly_used_percent=None,
                weekly_window_minutes=None,
                weekly_resets_at=None,
                raw_json="{}",
            ),
        ]
        markers = {
            "a": [SessionContextMarker("a", 0, "C:/RepoA", "C:/RepoA", "RepoA", "repo")],
            "b": [SessionContextMarker("b", 0, "C:/RepoB", "C:/RepoB", "RepoB", "repo")],
        }

        buckets, legend, repo_totals = build_project_stacks(
            events,
            markers,
            "1h",
            bucket_count=1,
            now=base,
            top_n=5,
            metric_mode="norm",
        )

        self.assertEqual(len(buckets), 1)
        self.assertEqual(buckets[0].total_tokens, 361)
        self.assertEqual(dict(legend), {"C:/RepoB": "RepoB", "C:/RepoA": "RepoA"})
        self.assertEqual(repo_totals[0]["C:/RepoA"], 160)
        self.assertEqual(repo_totals[0]["C:/RepoB"], 201)

    def test_build_project_stacks_caps_repo_legend_to_top_five_plus_other(self) -> None:
        base = datetime(2026, 4, 3, 0, 0, tzinfo=UTC)
        events: list[TokenEvent] = []
        markers: dict[str, list[SessionContextMarker]] = {}
        totals = [700, 600, 500, 400, 300, 200, 100]
        for index, total in enumerate(totals, start=1):
            session_path = f"session-{index}"
            events.append(
                TokenEvent(
                    session_path=session_path,
                    line_offset=index,
                    event_timestamp=base,
                    total_tokens=total,
                    input_tokens=0,
                    cached_input_tokens=0,
                    output_tokens=0,
                    reasoning_output_tokens=0,
                    cumulative_total_tokens=total,
                    weekly_used_percent=None,
                    weekly_window_minutes=None,
                    weekly_resets_at=None,
                    raw_json="{}",
                )
            )
            markers[session_path] = [
                SessionContextMarker(
                    session_path,
                    0,
                    f"C:/Repo{index}",
                    f"C:/Repo{index}",
                    f"Repo{index}",
                    "repo",
                )
            ]

        _buckets, legend, repo_totals = build_project_stacks(
            events,
            markers,
            "1h",
            bucket_count=1,
            now=base,
            top_n=5,
        )

        self.assertEqual(len(legend), 6)
        self.assertEqual(legend[-1], ("__other__", "Other"))
        self.assertEqual(repo_totals[0]["__other__"], 300)


if __name__ == "__main__":
    unittest.main()
