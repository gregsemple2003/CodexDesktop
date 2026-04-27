# Task 0009 Plan

## Planning Verdict

This plan was explicitly approved for implementation on `2026-04-26`.

Implementation starts with `PASS-0000`.

The approved scope is the committed-work `Tasks` tab while [Task-0011](../Task-0011/TASK.md), [Task-0008](../Task-0008/TASK.md), and [Task-0010](../Task-0010/TASK.md) own the adjacent intake and backend subsystems.

## Planning Basis

This plan is grounded in:

- the current repo design anchor at [Design/GENERAL-DESIGN.md](../../Design/GENERAL-DESIGN.md)
- the backend-and-dashboard runtime split delivered by [Task-0005](../Task-0005/TASK.md)
- the intervention and Dream analysis preserved in [Task-0007](../Task-0007/TASK.md)
- the review/intake split preserved in [Task-0011](../Task-0011/TASK.md)
- the task-local design brief at [Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md](./Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md)
- the visual style reference at [Design/STITCH-PROMPT.md](./Design/STITCH-PROMPT.md)
- the current generated Stitch mockup at [Mockup/stitch_task_tab/screen.png](./Mockup/stitch_task_tab/screen.png) and [Mockup/stitch_task_tab/code.html](./Mockup/stitch_task_tab/code.html)

## Visual Style Standard

Use [Design/STITCH-PROMPT.md](./Design/STITCH-PROMPT.md) and the generated [Stitch mockup](./Mockup/stitch_task_tab/screen.png) as the visual style reference standard for implementation.

The Stitch prompt is not an exhaustive pixel or component contract. It does not explicitly show every required state, row, control, or backend-dependent affordance. Do not overfit absent details or treat missing mockup details as permission to skip acceptance criteria from [TASK.md](./TASK.md) or this plan.

The implemented `Tasks` tab should stay visually true to the mockup's style, tone, hierarchy, density, and interaction feel:

- dark Windows-first supervision cockpit
- high-contrast operational hierarchy
- grouped task stream plus persistent detail pane
- crafted cards and rows rather than default widget styling
- explicit state color semantics for needs-you, running, ready, sleeping, failed, and blocked states
- concrete action labels and calm status copy

Where TASK.md or PLAN.md requires states or behavior not explicitly shown in the Stitch prompt, extend the mockup language consistently rather than inventing an unrelated surface.

The current generated mockup contains known semantic drift that must be overridden during implementation:

- Do not show unpromoted `Candidates` as normal work on the `Tasks` surface. Candidate and intake work belongs to `Review` and [Task-0011](../Task-0011/TASK.md) until it is promoted into a real committed task.
- Do not label a committed task as `Candidate` or `Prov: Candidate`. For committed tasks that came from an upstream review or Dream path, use durable source provenance such as `Promoted from Dream`, `Promoted from Review`, or equivalent backend-provided provenance.
- Do not represent long-running AI task progress as a progress bar. Prefer elapsed time, last meaningful progress, current state, and next expected event. Progress bars imply false precision for this kind of work unless a specific backend contract can prove real bounded progress.

## Validation And Regression Lane Isolation

Do not modify or depend on the human's live lane.

All Task-0009 validation and regression must use isolated agent-owned lanes:

- use agent-owned data, not the human's dashboard app data
- use agent-owned backend or app ports, not the human's service-lane ports
- use a validation or regression lane that can be started, inspected, and torn down without disturbing the human's always-on dashboard service
- keep screenshots, smoke outputs, fixtures, and proof data task-owned or otherwise disposable

The human's dashboard app data is off-limits unless the human gives explicit later instructions.

Regression testing must be carried out on the task's isolated lane and data, not the human's service lane or app data. If an isolated lane cannot be brought up, record the blocker honestly in task-owned testing artifacts instead of falling back to the human's lane.

## Run Control, Human Instructions, And Deep Context Standard

Preserve the owner split between [Task-0008](../Task-0008/TASK.md) and this task:

- Task-0008 owns backend run-control semantics, Temporal signals or updates, backend run state, cleanup or restore behavior, action gating, and deep-context readback.
- Task-0009 owns the humane `Tasks` tab presentation and the explicit launch controls that consume that backend truth.

When a task is running, the human-facing stop or hold path is a bounded `Pause` action, not a UI-side process killer. If the Task-0008 backend contract still names the durable transition `interrupt` internally, Task-0009 should keep that backend/runtime truth intact while translating the visible action, row copy, and detail-pane language to `Pause` where the human is choosing to stop or hold a running task.

The `Tasks` tab should only enable `Pause` when backend readback says the underlying run-control action is valid. Invoking it should call the backend task-run control contract, then render the resulting durable state, cleanup or restore status, follow-up review requirement, human-action target, and action-block reasons from backend readback.

Do not implement arbitrary local process termination from the dashboard UI as the normal stop behavior. If backend interruption or cleanup is blocked, the tab should show that blocked state honestly rather than pretending the run stopped.

A paused task should be visible as durable paused, waiting-for-human, or pause-review state from backend readback, with:

- an explicit reason
- an explicit `human_action_target` or equivalent target for the next human action
- visible next-owner and next-expected-event copy
- clear action availability for whether the same run can continue

When the human needs to provide custom instructions before continuing, the primary product path is not a dashboard text box. The `Tasks` tab should show enough context for the human to understand why input is needed, then launch the exact live thread or working context where the human can talk to the running Codex session directly.

For paused or waiting-for-human runs, prefer these launch actions when backend readback exposes the target:

- `Open Live Thread`
- `Open Thread`
- `Open Working Context`

Use direct dashboard instruction entry only as a fallback or later enhancement when the backend explicitly supports it and no better live-thread path exists. Do not prioritize a separate dashboard form that makes the human re-enter instructions outside the thread that is already doing the work.

Continuing a paused or waiting run should be a bounded backend action, visibly named `Resume` or `Continue`. It should send the existing Temporal workflow a backend-owned signal or update with either:

- the instruction payload, or
- a durable reference to the updated task or run artifact that holds the instruction delta

If a live thread target exists and the run is waiting on the human, typing directly in that thread is the preferred way to continue. If Task-0008 needs a Temporal resume signal after that human response, Task-0008 owns bridging the live-thread response back into backend workflow state. Task-0009 owns exposing the launch target, honest state, and clear copy about where the human should respond.

The backend, not the UI, owns whether the existing run can safely continue. Readback should record and Task-0009 should display:

- who or what resumed the run
- what instruction delta was supplied or referenced
- what next expected event is now active
- whether the current run still owns the task story

If the existing run cannot safely continue, the `Tasks` tab should say that plainly and offer a new dispatch or superseding run only when backend readback says that action is valid. Do not silently turn `Resume` into a new unrelated dispatch.

VSCodium or editor access should be exposed as explicit deep-context launch targets such as:

- `Open Task`
- `Open Working Context`
- `Open Live Thread`, `Open Thread`, or `Open Transcript`

Those actions should use backend `deep_context` or equivalent launch-target readback from Task-0008 when a run exists, and task artifact paths such as the task folder, `TASK.md`, `PLAN.md`, and `HANDOFF.md` when only declared task context exists. The intended experience is direct context recovery from the selected task, not manual markdown hunting or transcript search by the human. If no live thread target exists, the UI may fall back to opening the task or working context and should explain that direct thread continuation is unavailable.

## PASS-0000 Lock The Product Surface

### Objective

Turn the task-local design brief into a stable product-surface contract for the first `Tasks` tab slice.

### Implementation Notes

- decide the first-release information architecture
- decide the primary summary regions and task list model
- decide the default row status taxonomy and action affordances
- keep backend and orchestration leakage out of the default human-facing copy

### Verification

- the task-local design brief and `Stitch` prompt agree on the core surface
- the first-release screen structure is explicit enough to implement

### Exit Bar

- the team can build the first `Tasks` tab without rediscovering what the screen is for

## PASS-0001 Build The Read-Only Task Surface

### Objective

Add the first real `Tasks` tab UI with trustworthy read-only task visibility.

### Implementation Notes

- add the new tab shell and navigation entry
- render the chosen summary regions
- render the task list and detail surface
- support the main non-destructive states:
  - empty
  - loading
  - populated
  - stale
  - backend unavailable

### Verification

- the dashboard can render the `Tasks` tab without dispatch actions enabled
- the tab communicates useful task state without requiring markdown-first investigation

### Exit Bar

- the `Tasks` tab is a real high-level monitoring surface even before active dispatch controls land

## PASS-0002 Add Dispatch And Recovery Actions

### Objective

Integrate the humane task surface with the durable execution-state and dispatch layer.

### Implementation Notes

- consume the APIs and contracts delivered by [Task-0008](../Task-0008/TASK.md)
- add bounded task actions such as:
  - dispatch
  - poke
  - pause
  - open live thread for human instructions
  - resume or continue
  - open thread or working context
- keep each action specific enough that the human knows what will happen before clicking

### Verification

- a dispatched task reflects live durable state changes on the `Tasks` tab
- the surface can distinguish sleeping, blocked, running, and waiting-for-human honestly
- a running task exposes visible `Pause` copy while preserving backend run-control truth
- paused or waiting-for-human runs prefer `Open Live Thread` or equivalent thread launch for human instructions when backend readback exposes that target
- if no live thread target exists, the UI falls back to a precise task or working-context target and says direct thread continuation is unavailable
- `Resume` or `Continue` is only available when backend readback says the existing run can continue, and superseding dispatch is only offered when backend readback says it is valid
- thread or session launch works from the task surface

### Exit Bar

- the `Tasks` tab is now useful for monitoring, pause/resume control, and deep-context recovery, not only for read-only review

## PASS-0003 Integrate Reviewed And Promoted Work Provenance

### Objective

Show reviewed and promoted work on the committed-work surface without turning `Tasks` into the intake queue.

### Implementation Notes

- consume the promotion provenance delivered by [Task-0010](../Task-0010/TASK.md) and the intake split delivered by [Task-0011](../Task-0011/TASK.md)
- separate:
  - authored tasks
  - promoted tasks with durable provenance
  - enqueued tasks
  - active dispatched runs
- make promoted-task lineage and enqueue state legible without surfacing still-pending review items as committed work
- treat mockup labels such as `Candidates` or `Prov: Candidate` as visual/mockup drift, not acceptance guidance

### Verification

- a human can tell the difference between an authored task and a promoted committed task at a glance
- promoted tasks appear on the `Tasks` surface without ambiguous provenance or intake-state leakage
- unpromoted candidates remain off the committed-work `Tasks` surface
- committed promoted work uses source provenance such as `Promoted from Dream` or `Promoted from Review`, not `Candidate`

### Exit Bar

- the `Tasks` tab becomes the genuine top-level committed-work surface across authored and promoted tasks

## PASS-0004 Polish, Audit, And Regression

### Objective

Make the tab trustworthy under real human use instead of only technically complete.

### Implementation Notes

- run the interface against the task-local design brief
- tighten copy, hierarchy, placeholder truth, and visible status semantics
- capture repo-root regression proof for the new human-facing lane

### Verification

- unit tests cover any new task-state mapping logic
- regression proves the real dashboard surface on isolated Task-0009 validation/regression lane data and ports, not on the human's service lane or app data
- task-owned audit confirms the tab still behaves like a humane high-level surface

### Exit Bar

- the first `Tasks` tab slice can be used as the product heart without apologizing for it

## PASS-0005 Pin The Human-Facing Dashboard Frontend

### Objective

Fix [BUG-0002](./BUG-0002.md) by making the desktop dashboard frontend use an
explicit pinned release path instead of launching from the mutable repo checkout.

### Implementation Notes

- publish dashboard source into `%LOCALAPPDATA%\CodexDashboard\dashboard-releases\`
- write `%LOCALAPPDATA%\CodexDashboard\dashboard-current-release.json`
- install a stable runtime launcher under `%LOCALAPPDATA%\CodexDashboard\dashboard-launcher\`
- register startup to invoke the runtime launcher rather than `C:\Agent\CodexDashboard`
- add proof scripts and tests that report release id, git commit, source dirty
  status, hash validation, launcher path, startup path, and running process identity
- keep backend service-lane worktree-root proof in scope when it blocks honest
  `Tasks` tab verification

### Verification

- unit tests cover the dashboard release scripts and startup command
- script parsing and release hash validation pass
- backend service-lane release isolation still validates
- the human-facing dashboard can be restarted from a pinned frontend release
- proof shows the visible `Tasks` surface is backed by both a pinned frontend and
  a backend that reads the repo-root `Tracking` directory

### Exit Bar

- no claim about the human-visible dashboard surface is made from backend-only
  proof
- the frontend has release evidence comparable to the backend service lane

## Task-Level Validation

Expected validation before closure:

- task-surface unit coverage where mapping or formatting logic exists
- repo-root regression for the real dashboard surface
- task-owned audit against the design brief and copy semantics

## Watchouts

- do not let the UI own the dispatch runtime
- do not let raw backend taxonomy become the default visible language
- do not hide uncertainty by collapsing sleeping, blocked, and waiting-for-human into one vague label
- do not make the `Tasks` tab depend on transcript reading for basic comprehension
- do not expose backend-internal `interrupt` copy as the primary human-facing stop/hold action when the intended product action is `Pause`
- do not let `Resume` or `Continue` create a new unrelated run unless backend readback explicitly routes the human to a valid superseding dispatch
- do not make a dashboard instruction text box the main path when the intended interaction is for the human to continue in the existing live Codex thread
- do not copy the generated mockup's `Candidates` or `Prov: Candidate` labels into the committed-work implementation
- do not use progress bars for AI task-run progress unless a future backend contract provides trustworthy bounded progress semantics
- do not use the human's dashboard service lane, ports, or app data for validation or regression
