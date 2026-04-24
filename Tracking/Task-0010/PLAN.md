# Task 0010 Plan

## Planning Verdict

This plan is ready for explicit human approval.

Human gate:

- do not start implementation until the human agrees that Dream output should become a daily product workflow rather than only a manual research exercise

## Planning Basis

This plan is grounded in:

- the Dream research and packet work completed under [Task-0007](../Task-0007/TASK.md)
- the backend recurring-job path from [Task-0005](../Task-0005/TASK.md)
- the future `Review`, `Tasks`, and dispatch split in [Task-0011](../Task-0011/TASK.md), [Task-0009](../Task-0009/TASK.md), and [Task-0008](../Task-0008/TASK.md)

## PASS-0000 Lock The Daily Dream Product Contract

### Objective

Define the exact daily output contract before implementing the automation.

### Implementation Notes

- define the Dream run trigger and cadence
- define the digest email contents and collapse model
- define the single promotion contract shared by email links and dashboard clients
- define the generated task-skeleton contract and seeded provenance

### Verification

- the daily output shape is explicit enough to implement without inventing behavior ad hoc

### Exit Bar

- the task no longer risks becoming `some email about Dream` or two competing promotion flows

## PASS-0001 Run Dream Through The Backend Job Path

### Objective

Turn Dream into a real recurring repo job with durable output.

### Implementation Notes

- add the job spec and backend wiring
- ensure Dream output lands in the canonical report home
- preserve run identity and failure status

### Verification

- a real backend-managed run produces the expected Dream artifacts

### Exit Bar

- Dream daily execution is no longer a manual ritual

## PASS-0002 Build The Digest Email

### Objective

Send a human-usable daily Dream digest with enough context to triage proposed work quickly.

### Implementation Notes

- render collapsible option-task sections
- include concise packet context
- preserve readable fallback content
- make candidate status unmistakable

### Verification

- a real digest email renders with useful context and readable collapse behavior

### Exit Bar

- the email is genuinely useful for triage, not just a proof artifact

## PASS-0003 Add Promotion To Real Tasks

### Objective

Let a selected option task become a real repo task with durable provenance.

### Implementation Notes

- define the promotion endpoint or action
- create the new `Tracking/Task-<id>/` skeleton
- preserve links back to the source packet and option task

### Verification

- promoting a candidate produces a usable real task skeleton
- provenance survives the promotion

### Exit Bar

- candidate work can join the normal task lifecycle without manual copy-paste

## PASS-0004 Expose Candidate And Promoted Work To Review And Tasks Surfaces

### Objective

Make candidate asks visible to `Review` and promoted work visible and enqueueable through the future `Tasks` UI.

### Implementation Notes

- connect candidate state to the model used by [Task-0011](../Task-0011/TASK.md) and promoted-task state to the model used by [Task-0009](../Task-0009/TASK.md)
- keep candidate versus promoted versus dispatched state explicit across both surfaces

### Verification

- the UI can distinguish candidate and promoted work honestly across `Review` and `Tasks`

### Exit Bar

- Dream output can flow through `Review` into real dispatch without losing its origin story

## Watchouts

- do not let email become the only durable record
- do not auto-dispatch promoted tasks silently
- do not reopen winner selection during promotion
- do not lose source packet provenance
