# Task 0009 Handoff

## Current Status

`Task-0009` has explicit human approval to begin implementation.

Current lifecycle state:

- phase: `implementation`
- current pass: `PASS-0001`
- last completed pass: `PASS-0000`
- current gate: `implementation`

This task owns the human-facing `Tasks` tab for committed work in CodexDashboard:

- the high-level view
- committed-task triage
- dispatch intent
- monitoring
- recovery affordances
- deep-context launch

This task does not own intake review, the promotion contract, the dispatch runtime, or the daily Dream automation.

Those are split out into:

- [Task-0011](../Task-0011/TASK.md)
- [Task-0008](../Task-0008/TASK.md)
- [Task-0010](../Task-0010/TASK.md)

If the `Tasks` tab later shows work that was promoted out of `Review`, it should consume durable provenance from the single promotion contract owned by [Task-0010](../Task-0010/TASK.md). This task does not define candidate review or promotion semantics.

## Current Objective

Execute `PASS-0001`, building the first read-only `Tasks` tab surface from the locked product contract.

The design work for that lives in:

- [Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md](./Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md)
- [Design/STITCH-PROMPT.md](./Design/STITCH-PROMPT.md)
- [Mockup/stitch_task_tab/screen.png](./Mockup/stitch_task_tab/screen.png)
- [Mockup/stitch_task_tab/code.html](./Mockup/stitch_task_tab/code.html)

The review-surface design and research now live under [Task-0011](../Task-0011/HANDOFF.md).

The product split is now explicit:

- `Review` for incoming asks
- `Tasks` for committed work

## Binding Implementation Constraints

Validation and regression for this task must not touch the human's live lane.

Use an isolated Task-0009 validation/regression lane with agent-owned data and ports. The human's dashboard service lane, service ports, and app data are off-limits unless the human gives explicit later instructions.

If isolated validation or regression cannot run, record the blocker in task-owned testing artifacts instead of falling back to the human's lane.

Implementation must stay aligned with the active generated mockup and style references while overriding the known semantic drift in [PLAN.md](./PLAN.md):

- no unpromoted candidates on `Tasks`
- no `Candidate` / `Prov: Candidate` provenance for committed work
- promoted work uses durable source provenance such as `Promoted from Dream` or `Promoted from Review`
- user-facing stop or hold action is `Pause`, even if backend internals use `interrupt`
- `Open Live Thread` is the preferred instruction path when available
- no AI task-run progress bars unless a future backend contract proves trustworthy bounded progress

## Completed Passes

### PASS-0000 Lock The Product Surface

Completed in this checkpoint.

Primary output:

- [Design/PASS-0000-PRODUCT-SURFACE-CONTRACT.md](./Design/PASS-0000-PRODUCT-SURFACE-CONTRACT.md)

Audit:

- [Testing/PASS-0000-AUDIT.md](./Testing/PASS-0000-AUDIT.md)

Result:

- the first-release information architecture is explicit
- the active mockup is accepted as a style reference, with candidate/progress-bar drift overridden
- action copy and instruction flow are frozen for implementation
- validation and regression lane isolation is explicit

## Current Baseline

The repo currently has:

- the token-velocity overlay
- the backend-backed `Jobs` tab
- task-owned markdown artifacts under `Tracking/`

What it does not yet have is one humane surface that ties together:

- real tasks
- dispatch state
- sleeping runs
- waiting-for-human committed work
- promoted-task provenance

That missing surface is why this task exists.

## Next Recommended Step

Start `PASS-0001` by implementing the read-only `Tasks` tab shell, grouped stream, selected detail pane, state/provenance rendering, and loading/empty/stale/backend-unavailable states against isolated validation data.

## Watchouts

- do not let the `Tasks` tab become a raw orchestration inspector
- do not hide task provenance
- do not let intake-review material leak back into `Tasks`; that belongs in [Task-0011](../Task-0011/TASK.md)
- do not make the human read markdown to understand the default screen
- do not use the human's dashboard lane, ports, or app data for validation/regression
- do not copy generated mockup candidate labels or AI progress bars into implementation

## References

- [TASK.md](./TASK.md)
- [PLAN.md](./PLAN.md)
- [Task-0005](../Task-0005/TASK.md)
- [Task-0007](../Task-0007/TASK.md)
- [Task-0008](../Task-0008/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
- [Task-0011](../Task-0011/TASK.md)
- [Design/GENERAL-DESIGN.md](../../Design/GENERAL-DESIGN.md)
