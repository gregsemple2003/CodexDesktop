# Task 0002

## Title

CodexDashboard Stitch-aligned overlay redesign.

## Summary

This task reworks the existing desktop overlay in `app/` into a Stitch-inspired control-room UI without changing the product's core role as a hotkey-first private cockpit.

The Stitch mockup is the primary composition reference. `Design/GENERAL-DESIGN.md` remains the product-intent anchor where the two differ.

## Goals

- Redesign the current overlay shell around the Stitch composition: compact header, clear dismiss path, interval toggle row, metric strip, single main token-velocity chart, status area, and footer controls.
- Preserve the existing hotkey-overlay operator-console behavior: quick summon/dismiss, immediate readability, and no maximize-first workflow.
- Keep the visual language dark, dense, and operational while still honoring the repo-root general design for the weekly burn/redline presentation, budget editing, startup toggle, and advisory context.
- Keep the work inside the existing desktop app under `app/` rather than creating a new product surface.

## Non-Goals

- Reworking token ingest, persistence, or aggregation logic.
- Turning the app into a transcript browser, session explorer, or multi-pane console.
- Rewriting the desktop app architecture from scratch.
- Adding decorative motion, light-theme variants, or consumer-product chrome.
- Changing the repo-root product intent in `Design/GENERAL-DESIGN.md`.

## Constraints And Baseline

- `Task-0001` already delivered a working hotkey-first dashboard baseline with background ingest and overlay behavior.
- The Stitch mockup is the main design input:
  - `Design/Mockups/stitch/DESIGN.md`
  - `Design/Mockups/stitch/code.html`
  - `Design/Mockups/stitch/screen.png`
- `code.html` is the structural reference and `screen.png` is the visual confirmation.
- When Stitch and `Design/GENERAL-DESIGN.md` disagree, keep the repo-root general design as the product-intent authority.
- Preserve the overlay as a compact private cockpit rather than a general-purpose window.

## Implementation Home

Keep task-owned artifacts under `Tracking/Task-0002/`.

Implement the UI work in the existing desktop app under `app/`.

## Acceptance Criteria

- The current dashboard surface is refactored to visually and semantically follow the Stitch overlay composition.
- The overlay remains hotkey-first and dismissible immediately, including via the existing close path and `Escape` if already supported.
- The UI keeps the dark operator-console feel and clearly separates title, interval controls, summary metrics, main chart, status, and footer actions.
- The main chart remains a single token-velocity chart, not a multi-pane browser or transcript view.
- The general-design requirements still appear where they are part of the product intent: weekly burn/redline context, budget editing, startup toggle, and advisory Codex context.
- The task stays bounded to interface implementation in `app/` and does not expand into ingest, persistence, or backend rewrites.

## References

- `Design/GENERAL-DESIGN.md`
- `Design/Mockups/stitch/DESIGN.md`
- `Design/Mockups/stitch/code.html`
- `Design/Mockups/stitch/screen.png`
- `Tracking/Task-0001/TASK.md`
- `Tracking/Task-0001/HANDOFF.md`
