# Task 0003

## Title

Add a Total/Norm metric toggle for CodexDashboard chart analysis.

## Summary

This task adds a top-level metric-mode toggle so the overlay can switch between raw total tokens and a cost-weighted normalized proxy for chart analysis.

The goal is to make burst inspection more useful without replacing the existing total-token budget model or the current dashboard operator workflow.

## Goals

- Add a `Total` / `Norm` toggle in the overlay header.
- Drive velocity and repo charts from the selected metric mode.
- Keep hover values aligned with the selected metric mode.
- Keep the right-click investigation flow honest by preserving raw bucket context for investigation prompts.
- Avoid misleading raw budget-line comparisons when the chart is in normalized mode.

## Non-Goals

- Replacing the existing weekly budget model with a new billing model.
- Claiming the normalized metric is an exact OpenAI billing number.
- Reworking persistence or scanner ingestion.
- Expanding the overlay into a new multi-pane analytics product.

## Implementation Home

Keep task-owned artifacts under `Tracking/Task-0003/`.

Implement the feature in the existing desktop app under `app/codex_dashboard/`.

## Acceptance Criteria

- The overlay header includes a visible `Total` / `Norm` toggle.
- In `Total` mode, the chart behaves as it does today with raw token totals.
- In `Norm` mode, the chart and hover values use a normalized weighted metric derived from uncached input, cached input, output, and reasoning output.
- Repo mode uses the selected metric mode consistently for stacked bar totals.
- Right-click investigation still analyzes the real raw bucket range rather than a misleading normalized “total tokens” value.
- The unit test suite passes.
- The repo-root desktop overlay regression lane is rerun honestly after the change.
