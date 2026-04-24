# Task 0009

## Title

Design and build the dashboard `Tasks` tab as the primary human surface for committed task dispatch and monitoring.

## Summary

CodexDashboard needs a high-level committed-work surface that tells a human, at a glance:

- what matters now
- what is running
- what is stuck
- what is waiting on the human
- which real tasks were recently authored or promoted
- what can safely be dispatched next

The current dashboard has useful but fragmented surfaces:

- the overlay explains token burn
- the `Jobs` tab explains backend job state

What is missing is the product heart for committed work:

- a humane task-dispatch and task-monitoring surface

This task owns that surface.

The intended `Tasks` tab is not a raw database browser, not a transcript explorer, and not the canonical intake surface for new asks.
That intake surface belongs to [Task-0011](../Task-0011/TASK.md).

It is a trustworthy command surface for long-running committed work:

- queue review
- dispatch intent
- active run monitoring
- stuck-run recovery
- task readiness
- task provenance
- task-detail drilldown
- thread launch for deeper work when needed

The tab should make the repo feel more like one system with durable memory and supervision, and less like a loose pile of tasks, session transcripts, job specs, and review artifacts.

When the tab surfaces work promoted out of `Review`, it must show durable provenance from the single promotion contract owned by [Task-0010](../Task-0010/TASK.md) and the intake split owned by [Task-0011](../Task-0011/TASK.md).

This task does not define promotion semantics or intake review.

## Writeup Type

Concrete implementation task.

The product direction is already clear enough to define the primary human surface, even though some supporting backend capability will land in follow-on tasks.

## Burden Being Reduced

The human is currently forced to reconstruct the state of meaningful work from too many places:

- task markdown
- handoff notes
- session memory
- backend job surfaces
- local context about which items are real committed tasks, which asks are still under review, and which tasks have quietly gone cold

The deeper burden is not only inconvenience.

It is supervision anxiety.

Without one humane high-level task surface, the human has to carry too much hidden system state in their head just to answer:

- what matters now
- what needs me
- what can safely be started
- what likely fell asleep

## Current Truth

The repo already has useful but partial surfaces:

- the token-usage overlay
- the backend-backed `Jobs` tab
- rich task-owned markdown under `Tracking/`

What it does not have is a unified human-facing task cockpit.

That means the current product truth is fragmented:

- work meaning lives in task docs
- automation truth lives in jobs and backend state
- intervention and Dream follow-on suggestions live in separate artifacts
- the human still has to reconstruct the actual task picture manually

## Target Truth

The target truth is a `Tasks` tab that lets a cold human understand the real work picture quickly enough to act with confidence.

After this task succeeds, the product should have one high-level surface that can honestly communicate:

- what needs attention
- what is running
- what is blocked
- what is ready
- which committed tasks were recently promoted or authored
- where to go deeper

## Causal Claim

If CodexDashboard gains a humane `Tasks` tab with:

- grouped task visibility
- explicit reason lines
- real attention prioritization
- bounded actions
- easy deep-context launch

then the human will no longer need to reconstruct task state from scattered artifacts for ordinary supervision work.

## Evidence

The current repo shape already shows the gap:

- [ui.py](../../app/codex_dashboard/ui.py) exposes `usage` and `jobs`, but no durable task surface
- [Task-0005](../Task-0005/TASK.md) proved the dashboard is moving toward backend-backed operational truth
- [Task-0007](../Task-0007/TASK.md) and its research and Dream artifacts proved the repo now generates richer promotable task material than a human should have to harvest manually
- the task-local design brief at [Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md](./Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md) makes the missing human-facing product center explicit

## Why This Mechanism

The right first intervention is a dedicated `Tasks` tab, not:

- patching a few more labels into `Jobs`
- telling the human to read more markdown
- relying on external memory of which tasks matter

This mechanism is chosen because the missing burden is a product-surface problem:

- there is no one high-level place where task meaning, status, and next action meet

## Scope Rationale

This task intentionally owns the human-facing committed-work surface, not the whole runtime stack.

That split is necessary because the surface must be designed and audited as a humane product artifact, while the runtime, promotion, and intake machinery belong in:

- [Task-0008](../Task-0008/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
- [Task-0011](../Task-0011/TASK.md)

Keeping the UI task separate avoids two common failures:

- backend convenience deciding the screen before the human need is defined
- UI work pretending it already owns runtime truth it cannot yet prove

## Human Relief If Successful

If this task succeeds, the human should be able to:

- re-enter the system faster after time away
- see what needs attention without detective work
- dispatch work with less fear of losing context
- recover a quiet task with less searching
- see promoted-task provenance without confusing it for still-pending review work

## Remaining Uncertainty

- the exact first-release screen layout still needs visual iteration
- the best staged rollout between read-only visibility and live dispatch controls is still open
- some visible states depend on backend capability from [Task-0008](../Task-0008/TASK.md)

## Falsifier

This task should be considered wrong or incomplete if, after implementation:

- the human still has to open markdown first to understand what to do next
- the screen cannot distinguish sleeping, blocked, waiting-on-human, and ready in a trustworthy way
- the default surface still behaves like a raw operator console rather than a humane task cockpit

## Internal Mechanism Map

### Mechanism 1: Attention Summary

Failure reduced:

- the human cannot quickly tell what matters now

Mechanism:

- add a top-level summary that makes urgent state visible first

### Mechanism 2: Grouped Task Stream

Failure reduced:

- the human has to infer priority from a flat task list

Mechanism:

- group the task stream by human meaning instead of raw task metadata

### Mechanism 3: Rich Detail Pane

Failure reduced:

- the human must read raw artifacts just to understand one task

Mechanism:

- provide a selected-task pane with summary, state, next step, artifacts, and bounded actions

### Mechanism 4: Deep-Context Launch

Failure reduced:

- recovery becomes expensive because the human must hunt for the right working context

Mechanism:

- keep a direct path from task surface to deeper working context

## Goals

- Make `Tasks` the high-level, always-useful home for task dispatch and task monitoring in CodexDashboard.
- Give the human one place to answer:
  - what should happen next
  - what is actively happening now
  - where intervention is needed
  - whether a task is safe to dispatch
- Keep the surface humane:
  - low interpretation cost
  - obvious next actions
  - calm failure handling
  - minimal hidden state
- Show the relationship between:
  - repo tasks
  - active agent work
  - durable execution state
  - promoted-task provenance
- Support fast task triage without forcing the human to read raw markdown first.
- Let the human click into a task and open the deeper working context when needed.
- Keep the first version legible on desktop without turning the app into a sprawling multi-pane IDE clone.
- Ground the surface in [GENERAL-DESIGNER.md](../../../../Users/gregs/.codex/Orchestration/Prompts/GENERAL-DESIGNER.md) and [INTERFACE-DESIGNER.md](../../../../Users/gregs/.codex/Orchestration/Prompts/INTERFACE-DESIGNER.md), not only in backend convenience.
- Produce a durable task-local design brief and a reusable `Stitch` prompt so the UI direction can be iterated visually before or during implementation.

## Non-Goals

- Building the dispatch runtime itself in this task.
- Defining Temporal orchestration internals here instead of in [Task-0008](../Task-0008/TASK.md).
- Implementing the daily Dream job or digest generation here instead of in [Task-0010](../Task-0010/TASK.md).
- Owning the canonical intake and review surface for new asks; that belongs in [Task-0011](../Task-0011/TASK.md).
- Replacing task-owned markdown artifacts with only UI state.
- Turning the `Tasks` tab into a raw transcript browser, log tail, or arbitrary file explorer.
- Solving every future task-flow need before the first humane high-level surface exists.
- Using task dispatch as a pretext to leak raw backend or agent jargon into the default UI.

## Rival Explanations Considered

- `The real problem is only missing backend state, not a missing UI surface.`
  - rejected because the human burden exists before full runtime sophistication: the current product still lacks a coherent task-level reading surface
- `The real problem is only missing documentation.`
  - rejected because the burden is repeated re-entry and supervision, not lack of static explanation

## Rival Mechanisms Considered

- expand the `Jobs` tab to carry task meaning
  - rejected because jobs and tasks are not the same unit of human reasoning
- rely on task-folder or transcript browsing from the file system
  - rejected because it exports too much reconstruction work to the human
- build a separate standalone task web UI outside the dashboard
  - rejected because the dashboard should remain the local supervision cockpit

## Constraints And Baseline

- The repo already has a backend-backed `Jobs` tab from [Task-0005](../Task-0005/TASK.md); this task should build on that direction instead of recreating local scheduler ownership.
- Task dispatch is a separate concern and belongs in [Task-0008](../Task-0008/TASK.md), even though the `Tasks` tab must consume it.
- Daily Dream automation and option-task promotion belong in [Task-0010](../Task-0010/TASK.md), and candidate review belongs in [Task-0011](../Task-0011/TASK.md), even though the `Tasks` tab must surface the resulting committed work.
- The current dashboard implementation lives under:
  - [ui.py](../../app/codex_dashboard/ui.py)
  - [jobs.py](../../app/codex_dashboard/jobs.py)
- The enduring repo-level design anchor is still [Design/GENERAL-DESIGN.md](../../Design/GENERAL-DESIGN.md):
  - compact
  - operational
  - high-signal
  - low-friction
- The tab must respect the human-facing outcome rule from shared `.codex` docs:
  - do not optimize for machine convenience if the result becomes hard to trust or hard to scan
- Task dispatch and monitoring are human-facing surfaces even when they control agent work.
- The first release must prefer visible status and honest incompleteness over fake autonomy.

## Tradeoffs

- A richer default surface improves comprehension but risks noise if grouping and hierarchy are weak.
- A calmer screen reduces overload but becomes dishonest if it hides sleeping or waiting state.
- Showing promoted-task provenance increases trust but can dilute focus if provenance overwhelms the default committed-work view.

## Shared Substrate

- [Task-0008](../Task-0008/TASK.md) for dispatch state and supervision
- [Task-0010](../Task-0010/TASK.md) for candidate-task promotion and promotion provenance rules
- [Task-0011](../Task-0011/TASK.md) for the canonical intake surface that feeds committed work into `Tasks`
- [ui.py](../../app/codex_dashboard/ui.py)
- `app/codex_dashboard/tasks_tab.py`
- `app/codex_dashboard/tasks_backend.py`
- the task-local design brief and `Stitch` prompt

## Not Solved Here

- backend-owned task dispatch runtime
- durable execution-state semantics
- daily Dream scheduling
- canonical review-surface UX for incoming asks
- candidate promotion backend flow

## Expected Resolution

- CodexDashboard gains a real `Tasks` tab with a stable information architecture and humane interaction model.
- The tab becomes the default high-level surface for task review and dispatch rather than forcing the human to bounce between task folders, chat, and scattered status tools.
- The surface can distinguish:
  - ready-to-dispatch tasks
  - running tasks
  - sleeping or stalled tasks
  - blocked tasks
  - waiting-for-human tasks
  - authored versus promoted committed work
- The human can reach the deeper working context from a task row without hunting through local transcripts manually.
- The task leaves behind a design brief strong enough to guide both UI implementation and future audit.

## What Does Not Count

- A list box of task ids and status labels with no humane prioritization or action design.
- A surface that requires the human to remember hidden queue state from earlier sessions.
- A pseudo-IDE screen that dumps raw backend fields instead of communicating task meaning.
- A `Tasks` tab that cannot explain why a task is stuck, asleep, blocked, or ready.
- A surface that only works if the human already knows internal orchestration terminology.
- A dispatch UI that depends on the human opening markdown first just to know what clicking a button would do.

## Implementation Home

Primary product home:

- `app/codex_dashboard/`

Likely implementation surfaces:

- `app/codex_dashboard/ui.py`
- a new task-surface module or modules under `app/codex_dashboard/`
- backend-client glue for task and dispatch state

Task-owned design and planning home:

- `Tracking/Task-0009/`

Supporting follow-on homes owned elsewhere:

- [Task-0008](../Task-0008/TASK.md) for dispatch runtime and durable execution state
- [Task-0010](../Task-0010/TASK.md) for daily Dream generation and option-task promotion

## Implementation Home Rationale

This belongs primarily in `app/codex_dashboard/` because the main intervention boundary is the human-facing product surface.

The failure being reduced is not only that task truth exists in the wrong file.

The failure is that the dashboard does not yet present task truth in one humane supervisory surface.

The task-local design brief belongs under `Tracking/Task-0009/` because the design direction is still being refined before it hardens into repo-root design truth.

## Proposed Changes

- update [ui.py](../../app/codex_dashboard/ui.py) so the dashboard exposes a `Tasks` tab alongside the existing surfaces
- add `app/codex_dashboard/tasks_tab.py` to own the `Tasks` tab view composition, grouping, row rendering, detail pane, and task action affordances
- add `app/codex_dashboard/tasks_backend.py` to consume task, run, and promoted-task provenance readback APIs from the backend side
- render a top summary strip for urgent and high-signal task states
- render a grouped task stream that distinguishes:
  - needs you
  - running
  - sleeping
  - blocked
  - ready
- render a persistent detail pane with:
  - summary
  - why this task exists
  - current state
  - what changed recently
  - next expected step
  - provenance
  - artifacts
  - bounded actions
- add direct `Open Task` and `Open Thread` affordances that use durable provenance from [Task-0008](../Task-0008/TASK.md)
- show authored versus promoted provenance for committed work without reintroducing still-pending review items onto the `Tasks` surface

## Acceptance Criteria

- CodexDashboard renders a real `Tasks` tab that is selectable from the main dashboard shell.
- The `Tasks` tab shows a top summary strip for:
  - needs attention
  - sleeping
  - running
  - ready
  - blocked
- The `Tasks` tab shows a grouped task stream and does not collapse all work into one flat table.
- Selecting a task updates a persistent detail pane that includes:
  - summary
  - current state
  - next expected step
  - artifact links
  - bounded actions
- The surface can honestly represent at least these states:
  - loading
  - populated
  - stale
  - backend unavailable
  - empty-but-healthy
- The surface distinguishes authored, promoted, and actively running tasks visually and semantically without showing still-pending review items as committed work.
- The surface can link back to upstream review provenance for promoted tasks without redefining the promotion flow owned by [Task-0010](../Task-0010/TASK.md).
- The surface exposes `Dispatch`, `Poke`, `Interrupt`, and deep-context launch only when the backing contract from [Task-0008](../Task-0008/TASK.md) says those actions are valid.
- A real UI proof bundle exists showing the populated screen, an attention state, and committed-task provenance.
- Repo regression proof exists for the real dashboard surface after the tab lands.

## Proof Plan

- review the implemented screen against the task-local design brief and `Stitch` prompt
- add unit coverage for any new task-state mapping or grouping logic
- capture real UI proof for:
  - empty or loading state
  - populated state
  - committed-task provenance
  - attention state
- run the real dashboard regression lane once the surface is implemented

## Open Questions

- How much of the first `Tasks` surface should work read-only before dispatch APIs are ready?
- Should the first release use a split-pane layout, stacked cards, or a hybrid master-detail pattern?
- How should promoted work from `Review` retain visible provenance without crowding the default task view?
- What is the cleanest way to express sleeping versus blocked versus waiting-for-human in one glanceable status system?
- Which summary metrics help trust and prioritization, and which ones would just be decorative noise?
- How much session or transcript detail should be visible in the default surface before it becomes clutter?

## References

- [Task-0005](../Task-0005/TASK.md)
- [Task-0007](../Task-0007/TASK.md)
- [Task-0008](../Task-0008/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
- [Task-0011](../Task-0011/TASK.md)
- [Design/GENERAL-DESIGN.md](../../Design/GENERAL-DESIGN.md)
- [GENERAL-DESIGNER.md](../../../../Users/gregs/.codex/Orchestration/Prompts/GENERAL-DESIGNER.md)
- [INTERFACE-DESIGNER.md](../../../../Users/gregs/.codex/Orchestration/Prompts/INTERFACE-DESIGNER.md)
- [backend/orchestration/README.md](../../backend/orchestration/README.md)
- [ui.py](../../app/codex_dashboard/ui.py)
- [investigation.py](../../app/codex_dashboard/investigation.py)
