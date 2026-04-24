# Task 0010

## Title

Run Dream daily, email the results, and promote option tasks into real tasks that the dashboard can review and enqueue.

## Summary

Dream should not stay trapped inside manual packet inspection.

The repo now has a much stronger Dream pipeline than it had before:

- first-principles burden analysis
- antagonistic solution generation
- audited winner synthesis
- richer winner-task writeups

But the output still requires too much manual harvesting.

This task turns that pipeline into a daily product bridge:

- run Dream on a schedule
- produce a digest email with enough context to review quickly
- show collapsible option-task sections in the email
- provide a `Promote to Task` action
- turn promoted options into real task skeletons
- make candidate asks visible to the `Review` tab and promoted tasks available to the `Tasks` tab for enqueue and later dispatch

This task is not the same as the intake UI, dispatch runtime, or the committed-work `Tasks` UI. It bridges Dream output into the normal task lifecycle.

The canonical promotion path is frozen for this task:

- [Task-0010](../Task-0010/TASK.md) owns the single backend promotion mechanism and provenance rules
- email `Promote to Task` links and any future `Review` tab `Promote to Task` button are both clients of that same promotion contract
- the product must not grow separate promotion semantics in email and dashboard

## Writeup Type

Concrete implementation task.

## Burden Being Reduced

Dream output currently produces useful solution and task material, but the human still has to harvest that output manually.

The human burden includes:

- noticing that a new packet exists
- opening the packet manually
- scanning candidate tasks by hand
- deciding what to copy into a real task
- preserving provenance manually

The deeper burden is that promising follow-on work can stay trapped inside reports instead of entering the normal task lifecycle.

## Current Truth

The repo has a much stronger Dream pipeline than before, but that pipeline still behaves like a manual research and packet workflow rather than a daily product loop.

Current truth is:

- Dream can produce useful candidate tasks
- the candidate tasks are durable in packet artifacts
- there is no daily digest and no first-class promotion path into `Tracking/Task-<id>/`

## Target Truth

The target truth is a daily Dream product flow that:

- runs through the intended backend path
- produces a useful digest email
- exposes candidate tasks clearly
- lets the human promote one into a real task with durable provenance
- makes candidate asks visible to the future `Review` surface and real tasks available to the future `Tasks` surface

## Causal Claim

If Dream output becomes a daily reviewed flow with:

- digest email
- collapsible candidate sections
- promote-to-task
- durable provenance

then useful proposed work will stop dying inside packets and can enter the normal task lifecycle with much less manual harvesting.

## Evidence

- [Task-0007](../Task-0007/TASK.md) preserved the research, packet work, and Dream-quality improvements that make daily Dream output worth operationalizing
- the shared Dream docs under `.codex` now define a richer solution and task-generation workflow than before
- [Task-0005](../Task-0005/TASK.md) already established the backend path for recurring jobs in this repo
- [Task-0011](../Task-0011/TASK.md), [Task-0009](../Task-0009/TASK.md), and [Task-0008](../Task-0008/TASK.md) define the product surfaces that should eventually consume candidate and promoted work

## Why This Mechanism

The right first intervention is not more manual guidance telling the human to inspect packets every day.

The right intervention is a product bridge:

- run Dream automatically
- summarize it humanely
- let the human promote good work directly

This mechanism is chosen because the current failure is one of adoption and flow, not only one of raw generation quality.

## Scope Rationale

This task intentionally owns the bridge from Dream output into the real task lifecycle, not the runtime dispatch layer and not the `Tasks` tab itself.

That split is necessary because:

- [Task-0011](../Task-0011/TASK.md) owns the intake and review surface for candidate asks
- [Task-0009](../Task-0009/TASK.md) owns the high-level task surface
- [Task-0008](../Task-0008/TASK.md) owns dispatch and supervision

This task should stop at:

- candidate generation
- digest review
- promotion into real tasks

## Human Relief If Successful

If this task succeeds, the human should get:

- a daily Dream review loop that does not require packet hunting
- enough context in email to triage quickly
- one-step promotion into a real task skeleton
- less copy-paste and less provenance loss

## Remaining Uncertainty

- Gmail-compatible collapse behavior may need compromise between HTML richness and fallback safety
- how much candidate context should be copied versus linked into the promoted task still needs product judgment

## Falsifier

This task should be considered wrong or incomplete if, after implementation:

- Dream still requires manual packet harvesting to become actionable
- promoted tasks lose source provenance
- candidate work is visually or structurally confused with approved or dispatched work
- the email remains too thin for real triage

## Internal Mechanism Map

### Mechanism 1: Daily Dream Execution

Failure reduced:

- Dream output only appears when manually run

Mechanism:

- schedule Dream through the repo's intended recurring job path

### Mechanism 2: Digest Email

Failure reduced:

- useful Dream output stays buried in packet files

Mechanism:

- send a triage-friendly daily digest with collapsible candidate sections

### Mechanism 3: Promotion To Real Task

Failure reduced:

- promising option tasks do not enter the normal backlog without copy-paste

Mechanism:

- create real task skeletons from promoted candidates while preserving provenance

### Mechanism 4: Visibility To Review And Tasks Surfaces

Failure reduced:

- candidate and promoted work still feel detached from the normal product workflow

Mechanism:

- make candidate asks visible to `Review` and promoted tasks available to `Tasks` and later dispatch flow

## Goals

- Run the Dream pipeline as a real daily product workflow.
- Preserve the output durably in the canonical intervention and Dream homes.
- Send a daily human-readable digest email with enough context to triage the proposed work.
- Use collapsible sections so the digest remains skimmable while still carrying useful context.
- Add a `Promote to Task` action for option-task proposals.
- When promoted, create a real task home with durable task artifacts rather than leaving the option only inside the packet.
- Make candidate asks visible in `Review` and promoted tasks visible and enqueueable through the future `Tasks` tab.
- Preserve provenance from promoted task back to:
  - source day packet
  - source problem
  - source winner
  - source option task

## Non-Goals

- Replacing Dream with only email summaries.
- Letting email alone count as durable task creation.
- Dispatching promoted tasks automatically without the task lifecycle and UI controls to support that honestly.
- Reopening winner selection inside the promotion step.
- Turning every Dream option into a task automatically.

## Rival Explanations Considered

- `The real problem is only that Dream needs better solutions, not better delivery.`
  - rejected because even stronger Dream output will stay underused if it remains hard to discover and promote
- `The real problem is only missing task UI.`
  - rejected because even with a good `Tasks` tab, manual packet harvesting would still be too expensive

## Rival Mechanisms Considered

- manual daily packet review with no email bridge
  - rejected because it exports too much vigilance work to the human
- email summary with no promotion path
  - rejected because it still leaves manual task creation as the bottleneck
- auto-promote every option task
  - rejected because candidate work must remain reviewable rather than becoming silent backlog inflation

## Constraints And Baseline

- [Task-0007](../Task-0007/TASK.md) and the shared `.codex` Dream process already define the current packet-generation baseline.
- [Task-0005](../Task-0005/TASK.md) already established the intended backend path for recurring jobs in this repo.
- [Task-0011](../Task-0011/TASK.md) owns the `Review` tab surface that should consume candidate asks.
- [Task-0009](../Task-0009/TASK.md) owns the `Tasks` tab surface that should consume promoted tasks.
- [Task-0008](../Task-0008/TASK.md) owns dispatch runtime and durable task-run state after a task is promoted and enqueued.
- The digest must stay truthful:
  - candidate work is not the same as approved work
  - promoted work is not the same as dispatched work
- The promotion flow must preserve enough source context that the human can review what was promoted and why.

## Tradeoffs

- Richer email context improves triage but risks overload if sections are not collapsible and well-structured
- Easy promotion improves throughput but risks noise if provenance and approval boundaries are weak
- Tight provenance improves trust but may make task creation more verbose

## Shared Substrate

- [Task-0005](../Task-0005/TASK.md) recurring-job infrastructure
- [Task-0007](../Task-0007/TASK.md) Dream and packet research baseline
- [Task-0011](../Task-0011/TASK.md) future review-surface consumer
- [Task-0009](../Task-0009/TASK.md) future task-surface consumer
- [Task-0008](../Task-0008/TASK.md) future dispatch consumer
- shared `.codex` Dream process and canonical intervention report homes

## Not Solved Here

- live task dispatch and supervision
- the main `Review` tab layout
- the main `Tasks` tab layout
- broader Dream quality improvements beyond this delivery bridge

## Expected Resolution

- Dream runs daily through the intended repo job path.
- The human receives a digest email that is actually useful for task triage.
- Candidate option tasks can be promoted into real tasks without manual copy-paste.
- Candidate asks can appear on the `Review` surface, and promoted tasks carry durable provenance onto the `Tasks` surface for later enqueue and dispatch.

## What Does Not Count

- A daily email with only links and no context.
- A promotion action that loses the source packet or option-task provenance.
- A promotion flow that creates vague task stubs with no durable task skeleton.
- A system that confuses candidate tasks with real approved work in the UI.

## Implementation Home

Primary product homes:

- `backend/orchestration/` for the recurring job path and promotion endpoints
- `app/codex_dashboard/` for future UI consumption

Durable output homes:

- canonical Dream packet output under `.codex/Orchestration/Reports/Interventions/`
- task-owned task skeletons under `Tracking/Task-<id>/`

Task-owned planning home:

- `Tracking/Task-0010/`

## Implementation Home Rationale

This belongs primarily in `backend/orchestration/` because the first intervention boundary is recurring product flow and promotion wiring, not only a UI affordance.

The digest, scheduling, and promotion actions should exist durably even if the dashboard UI is not open.

The UI should consume the results later, but it should not be the only home of the daily Dream bridge.

## Proposed Changes

- add a backend-owned daily Dream job path, including the tracked job spec under `.codex/Orchestration/Jobs/specs/`
- add digest generation that summarizes candidate tasks with collapsible sections and readable fallback
- add one canonical backend promotion contract for Dream candidates, with email links and dashboard clients both targeting that same mechanism
- create a promoted-task writer that allocates the next repo task id by scanning existing `Tracking/Task-*` directories and writing the next task skeleton atomically
- generate these artifacts for every promoted task:
  - `Tracking/Task-<id>/TASK.md`
  - `Tracking/Task-<id>/PLAN.md`
  - `Tracking/Task-<id>/HANDOFF.md`
  - `Tracking/Task-<id>/TASK-STATE.json`
- seed promoted tasks with durable provenance back to:
  - source packet
  - source problem
  - source winner
  - source option task
- expose candidate visibility to the `Review` surface and promoted-task visibility to the `Tasks` surface once [Task-0011](../Task-0011/TASK.md) and [Task-0009](../Task-0009/TASK.md) are ready

## Acceptance Criteria

- A backend-managed daily Dream job exists through the intended repo job path and produces the expected Dream packet outputs.
- A real digest email is emitted for a daily Dream run and includes collapsible candidate sections plus readable non-collapsible fallback content.
- The email's `Promote to Task` action and any dashboard-side candidate promotion from `Review` both use the same backend promotion contract and do not fork promotion semantics.
- Promoting a candidate creates a real task skeleton with:
  - `TASK.md`
  - `PLAN.md`
  - `HANDOFF.md`
  - `TASK-STATE.json`
- The generated `TASK-STATE.json` starts with:
  - `status: pending`
  - `phase: planning`
  - `plan_approved: false`
  - `current_pass: null`
  - `last_completed_pass: null`
  - `current_gate: planning`
  - `latest_audit_verdict: unknown`
- The promoted task records durable provenance back to the source packet, problem, winner, and option-task artifact.
- Candidate asks can be surfaced to the `Review` tab model, and the resulting promoted task can be surfaced to the `Tasks` tab model, without losing source lineage.
- Real proof exists for one full flow:
  - daily Dream run
  - digest emission
  - candidate promotion
  - generated task skeleton

## Proof Plan

- prove a real backend-managed daily Dream run
- capture a real digest email example with readable collapse behavior and fallback
- prove promotion of one candidate into a real task skeleton with durable provenance
- prove the candidate ask can surface in `Review` and the promoted task can later surface to the `Tasks` tab model without losing source lineage

## Open Questions

- How much of the option-task body should be copied versus linked when building the new task skeleton?
- How should the email render collapsible sections across Gmail clients while staying readable in plain text fallback?

## References

- [Task-0005](../Task-0005/TASK.md)
- [Task-0007](../Task-0007/TASK.md)
- [Task-0011](../Task-0011/TASK.md)
- [Task-0009](../Task-0009/TASK.md)
- [Task-0008](../Task-0008/TASK.md)
- [HANDOFF.md](../Task-0007/HANDOFF.md)
