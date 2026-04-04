# Stitch Jobs Tab 0001

## Status

User-approved as the current design direction for the `Jobs` surface.

Imported source artifacts:

- `Tracking/Task-0004/Design/STITCH-JOBS-TAB-0001/DESIGN.md`
- `Tracking/Task-0004/Design/STITCH-JOBS-TAB-0001/code.html`
- `Tracking/Task-0004/Design/STITCH-JOBS-TAB-0001/screen.png`

Original capture source:

- `C:\Users\gregs\Downloads\stitch_codex_token_velocity_overlay`

## What Is Approved

- Keep the existing token cockpit as the default `Usage` lane.
- Add a sibling `Jobs` lane for declarative Codex jobs and reconciliation state.
- Keep the jobs surface dense, dark, and operational.
- Keep summary counts, last reconciled time, per-job health, and plain-language drift reasons visible without opening details.
- Keep `Reconcile` and `Refresh` as bounded top-level actions.
- Keep row-level `Details` as the reveal point for raw Windows plumbing.

## Carry-Forward Decisions

- Treat the imported mockup as the visual direction for `Task-0004`, not as a pixel-perfect contract.
- Treat the `Jobs` tab as an additive surface on top of the current single-overlay implementation in `app/codex_dashboard/ui.py`.
- Preserve the current product hierarchy:
  - `Usage` remains primary
  - `Jobs` is the new secondary machine-state lane
- Keep the default visible layer human-facing:
  - use `Desired / observed`
  - use `Kind` or `Mechanism`
  - use `Drift status`
- Keep raw job mechanisms, Windows identifiers, paths, command lines, session ids, kernel-like labels, and similar diagnostics behind details or a later debug surface.
- Keep `unknown` and `blocked` as distinct visible states rather than folding them into `missing`.

## Scope Notes

- The mockup currently shows `Usage`, `Jobs`, `Logs`, and `Terminal` in the shell.
- For `Task-0004`, only `Jobs` is in scope as a new implemented surface.
- `Logs` and `Terminal` are design placeholders only unless later task artifacts explicitly pull them into scope.
- The bottom footer telemetry shown in the mockup should only survive into implementation if it maps to a real, useful product signal.

## Follow-Up Consulting Round

The approved mockup was reviewed again through:

- general-designer
- interface-designer

Main outcomes:

- the concept is approved and should not be reopened
- visible labels should stay more human-facing than operator-facing
- the scope of `Reconcile` must be explicit in the eventual UI copy
- the design must not let decorative or debug-like shell elements crowd the core jobs message
