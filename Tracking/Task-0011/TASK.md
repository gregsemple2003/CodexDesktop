# Task 0011

## Title

Design and build the dashboard `Review` tab as the canonical intake surface for incoming asks.

## Summary

CodexDashboard now has two different human jobs to support:

- supervise committed work
- review incoming asks before they become committed work

`Task-0009` owns the first job through the `Tasks` tab.

This task owns the second job through a new `Review` tab.

The review job is not narrow Dream-email triage.

It includes multiple ask pipelines that want explicit human judgment:

- Dream-generated option tasks
- interface-review findings
- QA or adversarial bug-finding batches
- general-design or first-principles product asks
- approval requests
- runtime anomalies that need acknowledgement, routing, or intervention

Email can still notify the human that those asks exist.

Email should not be the canonical place where the human manages them.

The intended `Review` tab is not a raw inbox and not a second `Tasks` tab.

It is a trustworthy intake surface that lets the human answer:

- what came in
- why it arrived
- what evidence supports it
- what action is being asked for
- what happens if I promote, approve, route, defer, or dismiss it

The surface must keep provisional asks separate from committed work while still making the next step cheap and explicit.

Any `Promote to Task` action exposed on this surface must be a client of the single backend promotion contract owned by [Task-0010](../Task-0010/TASK.md).

This task does not define a second promotion mechanism.

## Writeup Type

Concrete implementation task.

The product split is now concrete enough to define a separate human-facing review surface instead of continuing to overload the `Tasks` tab or email.

## Burden Being Reduced

The human is currently forced to reconstruct incoming asks from too many places:

- email digests
- packet folders
- task-local design notes
- remembered context about which asks were already reviewed
- local judgment about which asks became real work and which ones remained provisional

The deeper burden is not just inconvenience.

It is review fatigue and trust erosion.

If incoming asks arrive as scattered side-band material, the human has to do too much reconstruction work before any real decision can happen.

That creates a second-order failure:

- the human starts avoiding the very review activities that would improve the system

## Current Truth

The repo has durable task artifacts, a backend-backed `Jobs` tab, and a planned `Tasks` surface for committed work.

What it does not have is one canonical in-product review surface for incoming asks.

That means the current product truth is fragmented:

- committed work and active runs are heading toward one surface
- incoming asks still spill across email, packets, docs, and memory
- candidate work and approval work risk being confused with already-committed tasks

## Target Truth

The target truth is a dedicated `Review` tab that lets a cold human understand incoming asks quickly enough to make trustworthy decisions without packet hunting.

After this task succeeds, the product should have one canonical intake surface that can honestly communicate:

- what new asks arrived
- which pipeline produced each ask
- what action each ask wants
- what evidence or provenance supports it
- what is still unresolved
- what happens next if the human acts

## Causal Claim

If CodexDashboard gains a humane `Review` tab with:

- grouped incoming asks
- source-aware detail
- explicit ask types
- clear dispositions
- truthful downstream actions

then incoming work will stop dying inside email and packet artifacts, and the human will no longer have to reconstruct intake state from scattered surfaces before acting.

## Evidence

- [Task-0009](../Task-0009/TASK.md) established the need for a humane committed-work surface and exposed the risk of overloading it with intake concerns.
- [Design/HUMAN-NEED-AND-REVIEW-TAB-DIRECTION.md](./Design/HUMAN-NEED-AND-REVIEW-TAB-DIRECTION.md) captures the human-needs argument for a separate `Review` tab.
- [Design/REVIEW-AND-INTAKE-SURFACE-DIRECTION.md](./Design/REVIEW-AND-INTAKE-SURFACE-DIRECTION.md) defines the concrete product direction that demotes email to notification rather than canonical state.
- [Research/2026-04-23-REVIEW-SURFACE-SURVEY.md](./Research/2026-04-23-REVIEW-SURFACE-SURVEY.md) surveys related product patterns and concludes that a distinct review surface is stronger than pushing incoming asks directly into the work queue.
- [Task-0007](../Task-0007/TASK.md) and [Task-0010](../Task-0010/TASK.md) prove the repo already produces richer follow-on asks than the human should manage through ad hoc packet inspection.

## Why This Mechanism

The right first intervention is not:

- more email polish
- putting candidate asks directly into the `Tasks` tab
- telling the human to keep opening packet folders

The right intervention is a separate review surface because incoming asks and committed work are different human jobs.

`Tasks` answers:

- what real work exists now
- what is running
- what is blocked
- what is safe to dispatch

`Review` answers:

- what just came in
- what it wants
- whether it should become durable work or take some other route

## Scope Rationale

This task intentionally owns the human-facing intake surface, not the backend generators or the committed-work surface.

That split is necessary because:

- [Task-0009](../Task-0009/TASK.md) owns the `Tasks` tab for committed work
- [Task-0010](../Task-0010/TASK.md) owns Dream generation and the single backend promotion contract
- [Task-0008](../Task-0008/TASK.md) owns dispatch and durable execution-state semantics

This task should stop at:

- intake presentation
- ask grouping
- ask detail and provenance
- human decision affordances that call the right downstream contracts

## Human Relief If Successful

If this task succeeds, the human should be able to:

- review new asks without packet hunting
- distinguish speculative asks from real work immediately
- process Dream batches without turning the `Tasks` tab into noise
- approve, promote, defer, dismiss, or route work with clear consequence lines
- re-enter the system and understand what still needs human judgment

## Remaining Uncertainty

- the exact first-release ask taxonomy still needs narrowing
- the first release may need a staged rollout between read-only intake and live dispositions
- some ask actions depend on backend contracts delivered by [Task-0010](../Task-0010/TASK.md) and [Task-0008](../Task-0008/TASK.md)

## Falsifier

This task should be considered wrong or incomplete if, after implementation:

- the human still has to open email or packet folders first to understand new asks
- the surface cannot distinguish incoming asks from committed work semantically and visually
- a user cannot tell what a disposition such as `Promote`, `Approve`, `Route`, `Defer`, or `Dismiss` will actually do
- the tab behaves like a noisy inbox instead of a calm decision surface

## Internal Mechanism Map

### Mechanism 1: Grouped Intake Stream

Failure reduced:

- incoming asks arrive as an undifferentiated stream

Mechanism:

- group asks by source, ask type, and urgency instead of dropping them into one flat list

### Mechanism 2: Source-Aware Detail Pane

Failure reduced:

- the human has to open side artifacts just to understand one ask

Mechanism:

- show provenance, evidence, requested action, and downstream consequence in one persistent detail surface

### Mechanism 3: Explicit Dispositions

Failure reduced:

- the surface can show asks but not tell the human what decision is actually available

Mechanism:

- expose bounded, truthful dispositions such as `Promote`, `Approve`, `Route`, `Defer`, or `Dismiss` only when the backing contract exists

### Mechanism 4: Downstream Surface Handoff

Failure reduced:

- promoted or approved work loses continuity after review

Mechanism:

- preserve enough provenance that reviewed items can flow into the right downstream surface without losing their origin story

## Goals

- Make `Review` the canonical human-facing intake surface for incoming asks in CodexDashboard.
- Keep `Review` and `Tasks` semantically separate:
  - `Review` for provisional asks
  - `Tasks` for committed work
- Support multiple ask sources without pretending they are all the same kind of item.
- Preserve provenance so the human can see where an ask came from and why it exists.
- Expose bounded dispositions that make the next action explicit.
- Keep email as notification and digest, not as the canonical state home.
- Produce task-local design and research artifacts strong enough to guide real implementation.

## Non-Goals

- Replacing the `Tasks` tab as the home for committed work.
- Defining the Dream backend promotion contract here instead of in [Task-0010](../Task-0010/TASK.md).
- Defining dispatch-state semantics or runtime interruption rules here instead of in [Task-0008](../Task-0008/TASK.md).
- Building every future pipeline producer before the first humane intake surface exists.
- Turning `Review` into a generic inbox, BI dashboard, or transcript browser.
- Auto-promoting all candidate work into tasks.

## Rival Explanations Considered

- `The real problem is only that Dream email needs better formatting.`
  - rejected because the review burden is broader than Dream and broader than email rendering
- `The real problem is only that the Tasks tab needs to show more things.`
  - rejected because incoming asks and committed work are different human jobs with different truth semantics

## Rival Mechanisms Considered

- keep email as the canonical home for candidate tasks and other incoming asks
  - rejected because it exports too much state reconstruction work to the human
- surface incoming asks directly inside the `Tasks` tab
  - rejected because it blurs the line between proposed work and committed work
- rely on packet folders and markdown links as the intake workflow
  - rejected because it makes the review job too expensive and too easy to avoid

## Constraints And Baseline

- [Task-0009](../Task-0009/TASK.md) remains the committed-work surface.
- [Task-0010](../Task-0010/TASK.md) remains the owner of Dream generation and the single backend promotion path.
- [Task-0008](../Task-0008/TASK.md) remains the owner of dispatch and durable run-state semantics.
- The first release must stay truthful:
  - a candidate ask is not yet committed work
  - an approval request is not the same as a task proposal
  - an anomaly ask is not necessarily a dispatchable task
- The review surface must respect the repo's human-facing outcome bar:
  - low reconstruction cost
  - clear consequences
  - honest uncertainty

## Tradeoffs

- A separate top-level tab improves clarity but increases navigation complexity slightly.
- A richer ask model improves truthfulness but risks noise if grouping and summarization are weak.
- More visible dispositions improve throughput but can become dangerous if consequence lines are not explicit.

## Shared Substrate

- [Task-0009](../Task-0009/TASK.md) committed-work surface
- [Task-0010](../Task-0010/TASK.md) Dream generation and promotion backend contract
- [Task-0008](../Task-0008/TASK.md) dispatch and anomaly-action semantics
- [ui.py](../../app/codex_dashboard/ui.py)
- the task-local design briefs and review-surface survey

## Not Solved Here

- backend-owned Dream scheduling
- packet generation quality
- dispatch runtime control semantics
- broader multi-pipeline backend ingestion contracts that still need separate design

## Expected Resolution

- CodexDashboard gains a real `Review` tab for incoming asks.
- The product gets a clean split between provisional intake and committed work.
- Dream candidates, approvals, review findings, and anomalies can be reviewed in one canonical surface without being mislabeled as already-approved tasks.
- Downstream actions preserve provenance instead of forcing the human to reconstruct why an item existed.

## What Does Not Count

- a better email template with no in-product review surface
- a raw chronological feed with no ask types or dispositions
- candidate work shown inside `Tasks` with only cosmetic labels
- a `Promote` button whose effect is unclear or whose semantics fork from [Task-0010](../Task-0010/TASK.md)

## Implementation Home

Primary product home:

- `app/codex_dashboard/`

Likely implementation surfaces:

- `app/codex_dashboard/ui.py`
- `app/codex_dashboard/review_tab.py`
- `app/codex_dashboard/review_backend.py`

Task-owned design and research home:

- `Tracking/Task-0011/`

Supporting follow-on homes owned elsewhere:

- [Task-0010](../Task-0010/TASK.md) for Dream generation and promotion
- [Task-0009](../Task-0009/TASK.md) for committed-work display after review decisions become real tasks
- [Task-0008](../Task-0008/TASK.md) for dispatch and runtime actions

## Implementation Home Rationale

This belongs primarily in `app/codex_dashboard/` because the missing boundary is a human-facing product surface, not only a backend rule.

The failure being reduced is that incoming asks do not yet land in one trustworthy review surface.

The task-local design and research artifacts belong under `Tracking/Task-0011/` because the intake model is still being refined before it hardens into broader product truth.

## Proposed Changes

- update [ui.py](../../app/codex_dashboard/ui.py) so the dashboard exposes a `Review` tab alongside the existing surfaces
- add `app/codex_dashboard/review_tab.py` to own the `Review` tab composition, grouping, row rendering, detail pane, and bounded ask actions
- add `app/codex_dashboard/review_backend.py` to consume ask, provenance, and disposition readback APIs from the backend side
- move the review-surface design and research docs out of [Task-0009](../Task-0009/TASK.md) into this dedicated task home
- render a top summary strip for high-signal intake states such as:
  - needs review
  - approvals
  - new work
  - anomalies
  - deferred
- render a grouped ask stream that distinguishes source and ask type
- render a persistent detail pane with:
  - summary
  - source
  - requested action
  - evidence and provenance
  - likely downstream consequence
  - bounded actions
- expose `Promote to Task` only through the single backend promotion mechanism owned by [Task-0010](../Task-0010/TASK.md)
- expose other actions such as approval, routing, or anomaly handling only when the backing contract exists and is explainable

## Acceptance Criteria

- CodexDashboard renders a real `Review` tab that is selectable from the main dashboard shell.
- The `Review` tab shows a top summary strip for high-signal incoming-ask states.
- The `Review` tab shows a grouped incoming-ask stream and does not collapse all review material into one flat inbox.
- Selecting an ask updates a persistent detail pane that includes:
  - summary
  - source
  - requested action
  - provenance
  - bounded actions
- The surface can honestly represent at least these states:
  - empty
  - loading
  - populated
  - stale
  - backend unavailable
- The surface distinguishes incoming asks from committed work visually and semantically.
- `Promote to Task` appears only for review items that are true task candidates and only through the single promotion mechanism owned by [Task-0010](../Task-0010/TASK.md).
- The surface exposes other dispositions only when the relevant backing contract exists and the human-visible consequence is explicit.
- A real UI proof bundle exists showing mixed incoming asks, a high-attention review state, and at least one candidate-task detail view.
- Repo regression proof exists for the real dashboard surface after the tab lands.

## Proof Plan

- review the implemented screen against the task-local design briefs and review-surface survey
- add unit coverage for any new ask-state mapping or grouping logic
- capture real UI proof for:
  - empty or loading state
  - mixed-source populated state
  - candidate-task detail
  - approval or anomaly attention state
- run the real dashboard regression lane once the surface is implemented

## Open Questions

- Which ask types should ship in the very first release versus later follow-on work?
- Which dispositions should be live first, and which ones should remain read-only or disabled until backend contracts exist?
- Should approvals always sort above speculative candidate work, or only when they block downstream execution?
- What is the cleanest way to represent batch-produced asks without turning the surface into a noisy queue?

## References

- [Task-0009](../Task-0009/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
- [Task-0008](../Task-0008/TASK.md)
- [Design/HUMAN-NEED-AND-REVIEW-TAB-DIRECTION.md](./Design/HUMAN-NEED-AND-REVIEW-TAB-DIRECTION.md)
- [Design/REVIEW-AND-INTAKE-SURFACE-DIRECTION.md](./Design/REVIEW-AND-INTAKE-SURFACE-DIRECTION.md)
- [Research/2026-04-23-REVIEW-SURFACE-SURVEY.md](./Research/2026-04-23-REVIEW-SURFACE-SURVEY.md)
- [Design/GENERAL-DESIGN.md](../../Design/GENERAL-DESIGN.md)
