# Task 0011 Handoff

## Current Status

`Task-0011` is newly created and currently pending planning approval.

This task owns the human-facing `Review` tab for CodexDashboard:

- incoming ask visibility
- grouped intake
- source-aware detail
- bounded review dispositions
- provenance-preserving handoff into downstream surfaces

This task does not own:

- the committed-work `Tasks` tab in [Task-0009](../Task-0009/TASK.md)
- the Dream backend generation or promotion contract in [Task-0010](../Task-0010/TASK.md)
- dispatch runtime and interruption semantics in [Task-0008](../Task-0008/TASK.md)

## Current Objective

Lock the first humane product surface for `Review` before the intake workflow hardens around email, packet hunting, or `Tasks`-tab overload.

The design and research work for that now lives in:

- [Design/HUMAN-NEED-AND-REVIEW-TAB-DIRECTION.md](./Design/HUMAN-NEED-AND-REVIEW-TAB-DIRECTION.md)
- [Design/REVIEW-AND-INTAKE-SURFACE-DIRECTION.md](./Design/REVIEW-AND-INTAKE-SURFACE-DIRECTION.md)
- [Research/2026-04-23-REVIEW-SURFACE-SURVEY.md](./Research/2026-04-23-REVIEW-SURFACE-SURVEY.md)

## Current Baseline

The repo currently has:

- the token-velocity overlay
- the backend-backed `Jobs` tab
- a planned `Tasks` tab for committed work
- no canonical review surface for incoming asks

That missing intake surface is why this task now exists.

## Next Recommended Step

Approve the product split and the first-release review contract:

- `Review` for incoming asks
- `Tasks` for committed work
- email as notification and digest, not canonical state

## Watchouts

- do not let the `Review` tab become a noisy inbox
- do not let `Review` become a second `Tasks` tab
- do not let candidate promotion semantics fork from [Task-0010](../Task-0010/TASK.md)
- do not make the human open packet folders to understand the default screen

## References

- [TASK.md](./TASK.md)
- [PLAN.md](./PLAN.md)
- [Task-0009](../Task-0009/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
- [Task-0008](../Task-0008/TASK.md)
- [Design/GENERAL-DESIGN.md](../../Design/GENERAL-DESIGN.md)
