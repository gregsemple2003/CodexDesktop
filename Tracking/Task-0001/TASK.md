# Task 0001

## Title

Codex token-velocity dashboard with a hotkey-first overlay.

## Summary

Build a Windows-first background app that watches `C:\Users\gregs\.codex` for new and modified session-history `.jsonl` files, extracts token-usage telemetry from the live append stream, and renders a compact bar-chart view of token velocity over multiple time buckets.

The primary user value is fast situational awareness:

- how many tokens were consumed recently
- whether current usage velocity is accelerating
- whether the current pace is likely to exhaust the weekly budget before reset
- whether that status can be checked with a single global hotkey instead of managing a normal foreground app window

## Goals

- detect new and modified `.jsonl` session-history files under `C:\Users\gregs\.codex`
- parse `token_count` events from those files instead of inferring usage from file counts or file sizes
- aggregate token usage into `1m`, `5m`, `15m`, `1h`, and `1d` intervals
- render those aggregates as a simple bar chart that can be read quickly
- compute and display a clear `redline` state when the current rate projects a weekly-budget overrun
- run quietly in the background with a global hotkey that summons and dismisses the dashboard overlay
- survive restarts without double-counting previously ingested events
- tolerate partially written trailing lines and files that are still being appended to by Codex

## Non-Goals

- building a full Codex session browser, transcript viewer, or orchestration console
- editing or rewriting files under `C:\Users\gregs\.codex`
- matching provider-side billing or quota calculations exactly when the local telemetry cannot prove them
- cross-platform support in the first version
- multi-user, networked, or cloud-synced telemetry storage
- a complex charting surface beyond the single token-velocity dashboard
- relying on a traditional maximize/minimize workflow as the primary UX

## Constraints And Baseline

- the real telemetry source already exists in local `.jsonl` session files under `C:\Users\gregs\.codex\sessions\YYYY\MM\DD\`
- sampled files contain JSONL records with:
  - top-level `timestamp`
  - `type: "event_msg"`
  - `payload.type: "token_count"`
  - `payload.info.last_token_usage.total_tokens`
  - `payload.info.total_token_usage.total_tokens`
  - `payload.rate_limits.secondary.window_minutes`
  - `payload.rate_limits.secondary.used_percent`
  - `payload.rate_limits.secondary.resets_at`
- `session_index.jsonl` exists, but the primary budget and velocity signal appears in the session-history files themselves
- files are append-only in practice, but the reader must assume:
  - partial trailing lines can appear while Codex is still writing
  - more than one active session file can be updated during the same wall-clock interval
  - the same files will be revisited across app restarts
- the weekly `redline` model must not assume that `used_percent` can be losslessly converted back into exact token counts
- the first version should prefer an always-available hotkey overlay over a conventional desktop window-management flow

## Implementation Home

Keep task-owned artifacts under `Tracking/Task-0001/`.

Target the implementation under repo-root `app/`.

If the recommended Tauri direction is used, keep:

- the Rust shell and background watcher in `app/src-tauri/`
- the overlay UI in `app/src/`

## Acceptance Criteria

- the app discovers new or modified session-history `.jsonl` files under `C:\Users\gregs\.codex` without requiring manual refresh
- the ingest pipeline extracts token usage from `token_count` events and records enough cursor state to avoid double-counting after restart
- the ingest pipeline ignores or safely retries incomplete trailing lines instead of crashing or corrupting aggregates
- the dashboard can switch between `1m`, `5m`, `15m`, `1h`, and `1d` bar views
- each interval view updates from the ingested event stream and reflects combined usage across all active session files
- the UI exposes a clearly labeled weekly-budget status with:
  - current recent velocity
  - projected weekly burn at that velocity
  - a visible `redline` state when the projection exceeds the configured weekly budget
- the first version supports a user-configured weekly token budget in absolute tokens
- when `rate_limits.secondary` metadata is available, the app stores and surfaces it as advisory context instead of pretending it is an exact local budget denominator
- a global hotkey can show and hide the overlay while the app remains backgrounded
- the overlay can be used without forcing the user to maximize or restore a conventional main window
- the app can start with Windows and keep its background footprint small enough that it is reasonable to leave running continuously
