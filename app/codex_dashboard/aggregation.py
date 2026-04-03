from __future__ import annotations

from datetime import UTC, datetime, timedelta

from .models import AggregationBucket, TokenEvent


INTERVAL_SECONDS = {
    "1m": 60,
    "5m": 300,
    "15m": 900,
    "1h": 3600,
    "1d": 86400,
}


def build_buckets(
    events: list[TokenEvent],
    interval_key: str,
    bucket_count: int,
    now: datetime | None = None,
) -> list[AggregationBucket]:
    if interval_key not in INTERVAL_SECONDS:
        raise ValueError(f"Unsupported interval: {interval_key}")
    now_utc = (now or datetime.now(UTC)).astimezone(UTC)
    interval_seconds = INTERVAL_SECONDS[interval_key]
    aligned_end_epoch = int(now_utc.timestamp()) // interval_seconds * interval_seconds
    buckets: list[AggregationBucket] = []
    totals: dict[int, int] = {}
    for event in events:
        bucket_epoch = (
            int(event.event_timestamp.astimezone(UTC).timestamp()) // interval_seconds
        ) * interval_seconds
        totals[bucket_epoch] = totals.get(bucket_epoch, 0) + event.total_tokens
    first_epoch = aligned_end_epoch - ((bucket_count - 1) * interval_seconds)
    for bucket_epoch in range(first_epoch, aligned_end_epoch + 1, interval_seconds):
        start_at = datetime.fromtimestamp(bucket_epoch, UTC)
        end_at = start_at + timedelta(seconds=interval_seconds)
        buckets.append(
            AggregationBucket(
                start_at=start_at,
                end_at=end_at,
                total_tokens=totals.get(bucket_epoch, 0),
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
