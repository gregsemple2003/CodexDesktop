from __future__ import annotations

import argparse
from datetime import UTC, datetime, timedelta
from pathlib import Path

from .aggregation import INTERVAL_SECONDS, build_buckets, is_over_redline, project_weekly_burn
from .config import DashboardConfig, load_config
from .storage import connect, count_events, initialize_db, load_events_since
from .scanner import ingest_once
from .ui import DashboardApp


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="CodexDashboard ingest utility")
    parser.add_argument("--config-path", type=Path, default=None)
    parser.add_argument("--scan-once", action="store_true")
    parser.add_argument("--print-summary", action="store_true")
    parser.add_argument("--no-ui", action="store_true")
    parser.add_argument("--smoke-artifact-dir", type=Path, default=None)
    parser.add_argument("--db-path", type=Path, default=None)
    parser.add_argument("--codex-root", type=Path, default=None)
    parser.add_argument(
        "--interval",
        choices=sorted(INTERVAL_SECONDS.keys()),
        default="1h",
    )
    return parser


def summary_text(connection, config: DashboardConfig, interval_key: str) -> str:
    initialize_db(connection)
    interval_seconds = INTERVAL_SECONDS[interval_key]
    now = datetime.now(UTC)
    events = load_events_since(connection, now - timedelta(days=7))
    buckets = build_buckets(events, interval_key, bucket_count=12, now=now)
    current_bucket_tokens = buckets[-1].total_tokens if buckets else 0
    projected = project_weekly_burn(current_bucket_tokens, interval_seconds)
    redline = is_over_redline(
        current_bucket_tokens,
        interval_seconds,
        config.weekly_budget_tokens,
    )
    total_7d = sum(event.total_tokens for event in events)
    lines = [
        f"events={count_events(connection)}",
        f"last_7d_tokens={total_7d}",
        f"interval={interval_key}",
        f"current_bucket_tokens={current_bucket_tokens}",
        f"projected_weekly_burn={projected}",
        f"weekly_budget_tokens={config.weekly_budget_tokens}",
        f"over_redline={redline}",
    ]
    return "\n".join(lines)


def main() -> int:
    parser = build_parser()
    args = parser.parse_args()
    config = load_config(args.config_path)
    if args.db_path is not None:
        config.db_path = str(args.db_path)
    if args.codex_root is not None:
        config.codex_root = str(args.codex_root)
    connection = connect(Path(config.db_path))
    try:
        if args.scan_once:
            ingest_result = ingest_once(connection, config)
            if args.print_summary:
                print(
                    f"files_scanned={ingest_result.files_scanned} "
                    f"files_updated={ingest_result.files_updated} "
                    f"events_ingested={ingest_result.events_ingested}"
                )
                print(summary_text(connection, config, args.interval))
            return 0
        if args.no_ui:
            parser.print_help()
            return 0
        app = DashboardApp(args.config_path, args.smoke_artifact_dir)
        app.run()
        return 0
    finally:
        connection.close()


if __name__ == "__main__":
    raise SystemExit(main())
