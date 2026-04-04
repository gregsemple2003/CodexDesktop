# Task 0004

## Title

Declarative Codex jobs registry, Windows reconciliation, and Jobs tab for CodexDashboard.

## Summary

CodexDashboard currently succeeds at one thing: a hotkey-first token-usage cockpit. It does not yet own the related machine-state problem.

Today, Codex-related Windows state is spread across one-off mechanisms:

- the dashboard is launched at sign-in through a Startup-folder `.cmd` file
- scheduled digest jobs live in Windows Task Scheduler and call scripts under `C:\Users\gregs\.codex\scheduled-digests\`
- the app itself still models startup as a special-case preference instead of a first-class managed job

That forces the user to inspect Startup-folder entries, Task Scheduler, and local scripts by hand. This task introduces one coherent model:

- a declarative registry of Codex-related Windows jobs
- a reconciler that compares desired state to actual durable Windows state
- a dedicated `Jobs` tab in the dashboard so drift and health are visible without polluting the token-usage cockpit

## Goals

- Define a tracked, user-authored declarative registry for local Codex-related jobs.
- Support an initial bounded set of Windows job kinds:
  - Startup-folder launchers
  - Scheduled Task entries
- Reconcile desired state against actual Windows durable state idempotently and classify drift honestly.
- Preserve `Usage` as the default hotkey-summoned surface and add a separate `Jobs` tab for machine-state visibility.
- Show at-a-glance summary counts plus per-job health, desired vs observed state, and plain-language drift reasons.
- Replace the dashboard startup special case with a first managed job entry under the jobs model.
- Keep raw Windows plumbing available behind explicit details instead of on the default surface.
- Make the current manually managed Codex jobs on this machine visible to the product instead of leaving them off-screen.

## Non-Goals

- Building a general-purpose Windows automation console.
- Supporting cross-platform job management in this task.
- Replacing the token-usage overlay as the app's primary job.
- Reworking ingest, token aggregation, normalized-metric semantics, or bucket investigation behavior.
- Building a freeform in-app command editor or Task Scheduler clone in the first pass.
- Managing unrelated non-Codex Windows tasks.
- Turning the overlay into a multi-pane admin shell.
- Adding `run now` execution controls for arbitrary jobs in the first slice.
- Treating the mockup's `Logs` and `Terminal` tabs as committed scope for this task.
- Adding decorative system footer telemetry unless it is backed by a real product signal.

## Constraints And Baseline

- The repo-level design anchor still defines CodexDashboard as a hotkey-first private cockpit whose primary use case is quick token inspection.
- The current visible app surface is still a single compact overlay under `app/codex_dashboard/ui.py`.
- The current Windows startup integration is implemented through a Startup-folder launcher in `app/codex_dashboard/startup.py`, with matching config state in `app/codex_dashboard/config.py`.
- This machine already has real Codex-adjacent durable state that proves the problem is not hypothetical:
  - `C:\Users\gregs\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\CodexDashboard.cmd`
  - `Codex Daily Agentic SWE Digest`
  - `Codex Daily Physical Agents Digest`
  - `Codex Daily UE Determinism Digest`
- The three existing scheduled digest tasks currently report real run history and failure state, so the product must surface health and drift, not just registration.
- The registry should live in tracked user-authored state under `C:\Users\gregs\.codex\`, not in opaque app-only storage.
- Unknown, unreadable, or partially verified Windows state must remain visible as unknown or blocked rather than being silently treated as healthy.
- The approved Stitch concept for the Jobs tab is now captured under `Tracking/Task-0004/Design/STITCH-JOBS-TAB-0001/`.

## Expected Resolution

- The overlay keeps the existing token cockpit as the default `Usage` tab.
- A sibling `Jobs` tab becomes the machine-state view for Codex-related Windows jobs.
- The first implementation uses a declarative file-backed registry plus reconciliation logic rather than more ad hoc toggles.
- The first supported job kinds are Startup-folder launchers and Scheduled Tasks.
- Existing supported Windows objects can be bootstrapped into the registry so the dashboard startup launcher and current digest tasks are not stranded outside the model.
- The default Jobs surface stays glanceable:
  - summary counts
  - per-job status
  - desired vs observed state
  - last reconciliation time
  - plain-language drift reason
- The default Jobs surface uses human-facing labels on the visible layer:
  - `Desired / observed`
  - `Kind` or `Mechanism`
  - `Drift status`
- Raw paths, task names, command lines, and other Windows-specific internals stay behind an explicit details reveal.
- The result feels like a bounded reconciliation surface, not a generic admin console.

## Implementation Home

Keep task-owned artifacts under `Tracking/Task-0004/`.

Implement dashboard, registry, discovery, and reconciliation code under `app/codex_dashboard/`.

Store the user-authored jobs registry in a durable tracked location under `C:\Users\gregs\.codex\` chosen by implementation, rather than burying it inside opaque application state.

## Acceptance Criteria

- The product defines a file-backed declarative registry for local Codex-related jobs in a durable tracked location under `C:\Users\gregs\.codex\`.
- The first implementation supports at least two Windows durable-state mechanisms:
  - Startup-folder launchers
  - Scheduled Tasks
- The reconciler can compare desired vs observed state for supported job kinds and classify at least:
  - `in sync`
  - `missing`
  - `drifted`
  - `disabled`
  - `unknown` or `blocked`
- Reconcile/apply behavior for supported job kinds is idempotent.
- The existing dashboard startup launcher is represented as a managed job instead of a special-case product control.
- The current scheduled digest tasks can be bootstrapped into, or otherwise brought under, the new jobs model without requiring the user to recreate them manually from scratch.
- The hotkey overlay defaults to the existing token `Usage` surface and adds a separate `Jobs` tab rather than mixing jobs into the chart area.
- The Jobs tab shows summary counts, per-job rows, last reconciliation time, and plain-language drift reasons.
- The default Jobs surface keeps raw Windows plumbing hidden unless the user explicitly opens details.
- The default Jobs surface uses plain visible labels for job semantics rather than operator acronyms such as `State (D/O)`.
- If implementation keeps the mockup's `Logs` or `Terminal` shell affordances visible, they must be clearly treated as inactive future surfaces and not implied to be part of this task's delivered scope.
- The primary `Reconcile` action has explicit UI copy that makes its scope clear.
- `unknown` and `blocked` remain first-class visible states rather than being collapsed into `missing`.
- The Jobs tab exposes bounded actions for this slice:
  - refresh state
  - reconcile/apply supported drift
  - enable or disable supported jobs
- Focused unit coverage passes for the new registry, discovery, and reconciliation behavior.
- The repo-root app-surface regression proof is updated or extended honestly to cover the new Jobs surface.

## References

- `Design/GENERAL-DESIGN.md`
- `Tracking/Task-0001/TASK.md`
- `Tracking/Task-0001/HANDOFF.md`
- `Tracking/Task-0002/HANDOFF.md`
- `Tracking/Task-0003/HANDOFF.md`
- `app/codex_dashboard/ui.py`
- `app/codex_dashboard/startup.py`
- `app/codex_dashboard/config.py`
- `C:\Users\gregs\.codex\scheduled-digests\`
- `Tracking/Task-0004/Design/STITCH-JOBS-TAB-0001/code.html`
- `Tracking/Task-0004/Design/STITCH-JOBS-TAB-0001/screen.png`
