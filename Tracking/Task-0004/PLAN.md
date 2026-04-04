# Task 0004 Plan

## Planning Verdict

This plan is approved for execution.

The approved registry location for managed jobs is:

- `C:\Users\gregs\.codex\Orchestration`

## Pass Order

### PASS-0000 Declarative Registry, Windows Discovery, Reconciliation, And Bootstrap

Objective:

- add the declarative jobs registry under `C:\Users\gregs\.codex\Orchestration`
- replace the startup special case with managed job data
- discover current supported Windows objects and classify desired vs observed state
- bootstrap the existing dashboard startup launcher and current digest tasks into the jobs model

Expected work:

- define the file-backed schema and storage path for tracked Codex jobs under `C:\Users\gregs\.codex\Orchestration`
- add startup-folder and Scheduled Task discovery readers
- add a bounded reconciliation model with honest states such as `in sync`, `missing`, `drifted`, `disabled`, and `unknown` or `blocked`
- add idempotent apply or reconcile behavior for supported job kinds
- remove or narrow the old startup-specific config or plumbing so the jobs model becomes the product abstraction

Verification:

- run `python -m unittest discover -s tests -p "test_*.py" -v`
- run `python -m app.codex_dashboard --scan-once --print-summary`

Exit bar:

- the registry is file-backed, durable, and user-authored under `C:\Users\gregs\.codex\Orchestration`
- startup launchers and Scheduled Tasks are both represented in the supported jobs model
- the reconciler compares desired vs observed state honestly for the supported job kinds
- the current dashboard launcher and existing digest tasks can be imported or represented without manual recreation from scratch
- focused unit proof for the new registry and reconciliation logic passes

### PASS-0001 Jobs Tab UI, Bounded Actions, And UI-Fidelity Proof

Objective:

- add the `Jobs` tab without disturbing `Usage` as the default hotkey surface
- render glanceable job summaries, per-job state, details reveals, and bounded actions
- satisfy the required UI-fidelity gate against the approved Stitch design reference

Expected work:

- add the tab shell and default-to-`Usage` behavior in the overlay
- add summary counts, per-job rows, last reconciliation time, desired vs observed state, and plain-language drift reasons
- keep raw Windows plumbing behind explicit details reveals
- add bounded actions for refresh, reconcile or apply, and enable or disable supported jobs
- keep any visible `Logs` or `Terminal` affordances clearly inactive placeholders if they remain on screen
- run the required UI review against `Tracking/Task-0004/Design/STITCH-JOBS-TAB-0001/`

Verification:

- run `python -m unittest discover -s tests -p "test_*.py" -v`
- compare the implemented Jobs surface against `Tracking/Task-0004/Design/STITCH-JOBS-TAB-0001/code.html` and `screen.png`

Exit bar:

- the overlay still opens on the token `Usage` surface by default
- the Jobs tab shows the required summary and per-job data without exposing raw Windows details by default
- visible labels stay human-facing instead of operator acronyms
- bounded actions are present and scoped honestly
- pass-local unit proof passes
- the UI review confirms the implemented Jobs tab matches the approved design direction closely enough to close the pass

### PASS-0002 App-Surface Regression, Handoff, And Closure

Objective:

- execute the repo-root desktop app regression lane against the delivered Jobs surface
- close the task only if the required app-surface proof and task artifacts are honest

Expected work:

- run repo-root `REG-001 Desktop Overlay Launch And Data Smoke`
- extend the exercised flow to cover the new Jobs tab and its visible job-state surface
- write the task-owned regression artifact and any required bug artifact if the lane fails or blocks
- finalize `HANDOFF.md` and the closure state when the regression evidence supports it

Verification:

- run the repo-root regression lane from `REGRESSION.md`

Exit bar:

- the regression artifact names the exact lane run, the flow exercised, why it counts, and what it does not prove
- the Jobs tab is included honestly in the app-surface validation story
- any failed or blocked required regression produces a task-owned bug artifact instead of a forced close
- `HANDOFF.md` and `TASK-STATE.json` reflect the real final state
