# Task 0004 Research Summary

## Research Verdict

Task research is planning-ready.

## Inputs And Evidence

- The durable product anchor still defines CodexDashboard as a hotkey-first private cockpit centered on token usage.
- The current implementation has only one visible surface in `app/codex_dashboard/ui.py`; it does not yet have a second machine-state view.
- Windows startup is currently implemented as a special case through `app/codex_dashboard/startup.py` plus config state in `app/codex_dashboard/config.py`.
- Real Codex-adjacent durable state already exists on this machine:
  - a Startup-folder launcher for CodexDashboard
  - three daily Codex digest Scheduled Tasks
- The three digest tasks currently report `LastTaskResult = 1`, which means the surface must show health and drift, not just whether something is registered.
- Local digest automation already has durable on-disk scripts and runtime folders under `C:\Users\gregs\.codex\scheduled-digests\`, so the problem space already includes real file-backed job definitions and run artifacts.
- The user-approved Stitch concept for the Jobs surface is now imported under `Tracking/Task-0004/Design/STITCH-JOBS-TAB-0001/`.

## Harvested Reviews

This task definition was pressure-tested through four delegated reviews:

- task-harvester
- vision-harvester
- general-designer
- interface-designer

All four converged on the same core move:

- stop modeling machine-state behavior as one-off controls
- introduce desired state vs actual state vs reconcile
- keep token usage as the default cockpit
- add a separate jobs-oriented surface rather than polluting the existing chart view

## Key Decisions

- Treat startup and scheduled jobs as managed machine state, not as preferences.
- Keep `Usage` as the default overlay tab and add a separate `Jobs` tab.
- Use a file-backed declarative registry in tracked user-authored state under `C:\Users\gregs\.codex\`.
- Support a bounded first set of Windows job kinds:
  - Startup-folder launchers
  - Scheduled Tasks
- Make bootstrap of the current dashboard startup launcher and digest tasks part of the first slice so the product can take ownership of real existing jobs.
- Keep the default Jobs surface glanceable and human-readable.
- Push raw Windows plumbing into explicit details rather than default-visible UI.
- Refuse a generic automation-console expansion in the first slice.
- Treat `Logs` and `Terminal` in the approved mockup as non-binding placeholders unless later tasks explicitly pull them into scope.
- Keep visible surface labels human-facing, even when the underlying state model is Windows-specific.

## Why This Is The Honest Repo Fit

- The current repo already owns the dashboard shell, background polling, config, and Windows startup integration.
- The user pain is now broader than the old startup toggle: the problem is fragmented Codex-related durable Windows state.
- Adding more one-off controls would deepen the current drift; a first-class jobs model is the cleanest next abstraction.
- A separate `Jobs` tab keeps the original token-spike workflow intact while giving the product a second honest responsibility.

## Carry-Forward Constraints

- Preserve hotkey-first summon and immediate dismiss behavior.
- Keep token usage as the primary default view.
- Do not widen into a general-purpose admin console.
- Keep unknown or blocked Windows state visibly unknown; do not fake confidence.
- Do not change ingest, aggregation, or investigation semantics as part of this task.
- The repo-level design anchor will need an explicit update during implementation because it still names a startup toggle and still frames the product as a single-surface overlay.
- The approved mockup is directional, not a literal scope promise for every shell element it shows.

## Approved Mockup Follow-Up

After the Stitch capture was imported, the design received a second consulting round through the general-designer and interface-designer roles.

What they validated:

- the approved design stays in the right product lane
- the Jobs surface is glanceable enough for the intended hotkey-overlay workflow
- the summary strip and per-job table structure are strong implementation anchors

What they tightened:

- visible labels should prefer plain language over internal telemetry or acronyms
- the product needs explicit status taxonomy for `needs attention`, `unknown`, and `blocked`
- `Reconcile` needs honest scope in final UI copy
- footer telemetry, session ids, kernel-like labels, and similar debug-looking shell elements should not survive by default unless they map to real product signals

## Planning Recommendation

Plan the work in three passes:

- `PASS-0000`: registry schema, Windows discovery, reconciliation model, and bootstrap/import of current supported jobs
- `PASS-0001`: `Jobs` tab, per-job details, bounded actions, and focused unit coverage
- `PASS-0002`: app-surface regression proof, handoff, and closure
