from __future__ import annotations

from datetime import UTC, datetime, timedelta, tzinfo

from .models import AggregationBucket, TokenEvent


INTERVAL_SECONDS = {
    "1m": 60,
    "5m": 300,
    "15m": 900,
    "1h": 3600,
    "1d": 86400,
}


def _resolve_display_timezone(
    now: datetime | None,
    display_tz: tzinfo | None,
) -> tzinfo:
    if display_tz is not None:
        return display_tz
    if now is not None and now.tzinfo is not None:
        return now.tzinfo
    return datetime.now().astimezone().tzinfo or UTC


def _align_bucket_start(dt_local: datetime, interval_key: str) -> datetime:
    if interval_key == "1d":
        return dt_local.replace(hour=0, minute=0, second=0, microsecond=0)
    if interval_key == "1h":
        return dt_local.replace(minute=0, second=0, microsecond=0)

    minute_step = INTERVAL_SECONDS[interval_key] // 60
    aligned_minute = dt_local.minute - (dt_local.minute % minute_step)
    return dt_local.replace(minute=aligned_minute, second=0, microsecond=0)


def build_buckets(
    events: list[TokenEvent],
    interval_key: str,
    bucket_count: int,
    now: datetime | None = None,
    display_tz: tzinfo | None = None,
) -> list[AggregationBucket]:
    if interval_key not in INTERVAL_SECONDS:
        raise ValueError(f"Unsupported interval: {interval_key}")
    timezone = _resolve_display_timezone(now, display_tz)
    now_local = (now or datetime.now(timezone)).astimezone(timezone)
    interval_seconds = INTERVAL_SECONDS[interval_key]
    aligned_end = _align_bucket_start(now_local, interval_key)
    buckets: list[AggregationBucket] = []
    totals: dict[datetime, int] = {}
    for event in events:
        bucket_start = _align_bucket_start(event.event_timestamp.astimezone(timezone), interval_key)
        totals[bucket_start] = totals.get(bucket_start, 0) + event.total_tokens

    for offset in range(bucket_count - 1, -1, -1):
        start_at = aligned_end - timedelta(seconds=interval_seconds * offset)
        end_at = start_at + timedelta(seconds=interval_seconds)
        buckets.append(
            AggregationBucket(
                start_at=start_at,
                end_at=end_at,
                total_tokens=totals.get(start_at, 0),
            )
        )
    return buckets


def project_weekly_burn(tokens_in_interval: int, interval_seconds: int) -> int:
    if interval_seconds <= 0:
        raise ValueError("interval_seconds must be positive")
    return int(tokens_in_interval * (7 * 24 * 60 * 60 / interval_seconds))


def is_over_redline(
    tokens_in_interval: int,
    interval_seconds: int,
    weekly_budget_tokens: int,
) -> bool:
    return project_weekly_burn(tokens_in_interval, interval_seconds) > weekly_budget_tokens
