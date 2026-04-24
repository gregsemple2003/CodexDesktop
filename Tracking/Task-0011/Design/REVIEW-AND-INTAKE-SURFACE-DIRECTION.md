# Review And Intake Surface Direction

## Why This Document Exists

The existing `Task-0009` brief correctly argues that CodexDashboard needs a strong `Tasks` surface for committed work.

That is still true.

But it is no longer sufficient.

The product is now growing a second major human job:

- review incoming asks from multiple pipelines before they become committed work

That includes more than Dream candidate tasks.

It includes:

- Dream-generated option tasks
- one-off interface reviews
- QA bug-finding runs
- general-design or first-principles ideation
- approval requests
- runtime exceptions that need human judgement

Email can help notify the human that this material exists.

Email should not be the canonical place where the human manages it.

This document defines the humane product direction for that need.

It now lives under `Task-0011` because the review-and-intake surface has been split out from the committed-work `Tasks` tab owned by [Task-0009](../../Task-0009/TASK.md).

## Primary Human Need

The human needs one trustworthy place to answer:

- what new asks came in
- what kind of ask each one is
- what evidence or provenance supports it
- what action is actually being requested from me
- what happens if I approve, promote, defer, route, or dismiss it

This is a different job from supervising active work.

Supervising active work asks:

- what is running
- what is sleeping
- what is blocked
- what is ready to dispatch

Reviewing incoming asks asks:

- what is this
- why did it arrive
- what does it want from me
- is it good enough to become durable work

That difference is important enough to deserve a separate surface.

## Product Decision

Add a top-level `Review` tab as the canonical surface for incoming asks.

Keep `Tasks` as the canonical surface for committed work.

Use `email` only as:

- notification
- digest
- reminder
- link back into `Review`

## Naming

### Human-Facing Name

Use:

- `Review`

Do not use:

- `Intake` as the primary label

Reason:

- the human job is review and decision-making
- `intake` sounds like plumbing or a backend ingestion layer
- the surface will include more than brand-new work

It also handles:

- approvals
- reroutes
- exceptions
- duplicates
- stale or reopened asks

### Internal Model Name

Internally, it is still fine to call the backend layer:

- `intake`
- `review queue`
- `incoming asks`

But the screen the human clicks should be named after the human job.

## Core Product Split

### Review

`Review` is for asks that are not yet committed work.

Examples:

- promote this Dream candidate into a real task
- approve this task plan
- route this issue to a bug task
- approve this closeout
- inspect this QA-found defect cluster
- review this design proposal before it hardens into implementation work

### Tasks

`Tasks` is for work that is already durable and committed.

Examples:

- pending but real tasks
- ready tasks
- running tasks
- blocked tasks
- sleeping tasks
- waiting-on-human tasks

### Jobs

`Jobs` remains the surface for recurring automation and backend job state.

### Why The Split Matters

If `Review` and `Tasks` are collapsed together:

- proposed work will look too real
- approvals will look like execution work
- the human will lose confidence in what is actually committed

If email is treated as the home instead:

- review state will fragment
- dedupe will be weak
- batch handling will be poor
- the product will offload too much memory burden to the human

## The Review Object Model

The product should normalize everything entering the surface into three concepts:

1. `batch`
2. `review item`
3. `disposition`

### Batch

A `batch` is one source run or one source package of related asks.

Examples:

- one Dream daily run for one repo
- one interface review session
- one QA exploratory pass
- one general-design ideation packet

Recommended batch fields:

- source pipeline
- repo
- created timestamp
- freshness status
- ask counts by type
- severity summary
- source artifact links
- run summary

### Review Item

A `review item` is one actionable ask inside a batch.

Recommended fields:

- stable id
- ask type
- title
- one-paragraph summary
- burden or problem being addressed
- rationale
- provenance
- supporting artifacts
- recommended action
- allowed actions
- current disposition

### Disposition

A disposition records what happened after review.

Minimum useful states:

- `new`
- `needs_review`
- `deferred`
- `duplicate`
- `dismissed`
- `approved`
- `promoted_to_task`
- `routed_to_bug`
- `routed_to_design`

This is better than a vague unread/read model because the human job is not just reading.

It is deciding.

## Ask Types

The first version should explicitly support multiple ask types instead of pretending everything is a candidate task.

Suggested starting taxonomy:

- `candidate_task`
- `approval_plan`
- `approval_task`
- `approval_closeout`
- `bug_or_defect`
- `design_gap`
- `runtime_exception`
- `follow_up_research`

This protects the UI from becoming one ambiguous bucket.

## Source Pipelines

The review surface should be source-aware.

Likely first sources:

- `Dream`
- `Interface Review`
- `QA Sweep`
- `General Design`
- `Runtime Supervision`

Source awareness matters because the next action may differ:

- a Dream candidate may be promoted
- a QA finding may be routed to a bug task
- a runtime exception may need poke, interrupt, or acknowledgement
- an approval item may need yes/no with rationale

## The Right UI Shape

### Overall Layout

The best likely desktop layout is:

- top header row
- summary cards
- three-column body

Three-column body:

- left: source and batch rail
- center: grouped review-item queue
- right: persistent detail pane

This is better than a single flat list because the human needs both:

- batch-level orientation
- item-level decisions

### Top Header

Should show:

- `Review`
- freshness sentence
- repo or source filter
- quick scope controls

Example:

- `Showing asks from 5 fresh batches across 3 repos.`

### Summary Cards

The first summary row should answer:

- what needs me now
- what kind of work it is
- whether the queue is growing

Suggested first cards:

- `Needs review`
- `Approvals`
- `New work`
- `Exceptions`
- `Deferred`

If later needed:

- `Stale`
- `Duplicates`

### Batch Rail

The left rail should help the human choose the right packet before reading every item.

Each batch row should show:

- source
- repo
- short summary
- freshness
- counts by ask type

Examples:

- `Dream · CodexDashboard · 7 candidate tasks`
- `QA Sweep · CodexDashboard · 3 defects, 1 duplicate`
- `Interface Review · CodexDashboard · 5 UI asks`

### Review Queue

The center column should show the items inside the selected scope.

Group by meaning, not raw chronology.

Suggested groups:

- `Needs review now`
- `Approvals`
- `New work proposals`
- `Exceptions and anomalies`
- `Deferred or snoozed`

Each row should show:

- human-readable title
- short summary
- ask type
- provenance
- recommended action
- freshness
- one reason line

Examples of good reason lines:

- `Dream thinks this should become a real task because the burden is recurring and under-specified.`
- `QA found a reproducible defect with screenshot proof and no matching open bug task.`
- `This approval blocks a task from becoming dispatchable.`

### Detail Pane

The detail pane must absorb the first minute of comprehension.

It should show:

- `Summary`
- `Why this arrived`
- `Source and provenance`
- `Evidence`
- `Recommended action`
- `What happens if you approve or promote`
- `Related task or bug links`
- `Actions`

## Action Model

The action set must depend on ask type.

Do not force every row into the same buttons.

Likely actions:

- `Promote to Task`
- `Approve`
- `Decline`
- `Route to Bug`
- `Route to Design`
- `Merge with Existing Task`
- `Defer`
- `Dismiss`
- `Open Source Packet`

### Important Rule

The product should show the consequence of an action before click.

Examples:

- `Promote to Task`
  - creates a real task skeleton and removes this item from the candidate queue
- `Approve Plan`
  - marks the task as ready for the next lifecycle step
- `Route to Bug`
  - creates or links a bug-tracking task

## Truth-Seeking And Compassion

This surface needs the same principles as the rest of orchestration.

### Truth

The screen must not blur:

- candidate work and committed work
- notifications and durable asks
- evidence-rich asks and weak suggestions
- real exceptions and ordinary waiting

It must preserve provenance and uncertainty honestly.

### Compassion

The human should not have to remember:

- where the ask came from
- what the ask wants
- what action is safe
- whether it was already seen yesterday

The system should hold that memory for them.

## Why Email Should Be Demoted

Email is valuable for awareness.

It is weak as the canonical home because it is bad at:

- item state
- dedupe
- merge and reroute
- batch comparison
- longitudinal queue management
- preserving one obvious next action

The humane role for email is:

- tell me there is new review work
- summarize the batch
- deep-link me into the exact place where I should act

That is enough.

## Relationship To Existing Task-0009 Direction

This does not make the existing `Tasks` work wrong.

It clarifies that `Tasks` is not the only heart of the product.

The stronger product story is:

- `Review` is the human decision surface for incoming asks
- `Tasks` is the human supervision surface for committed work
- `Jobs` is the automation runtime surface

This is better than overloading one tab with all three jobs.

## Example Review Flow

### Dream Batch Flow

1. Open `Review`.
2. See a fresh Dream batch for `CodexDashboard`.
3. Select the batch.
4. Scan its seven candidate tasks grouped under `New work proposals`.
5. Open one item in the detail pane.
6. Read the burden summary, provenance, and recommended action.
7. Click `Promote to Task` or `Dismiss`.

### QA Flow

1. Open `Review`.
2. Select the latest QA batch.
3. See three defect asks and one duplicate.
4. Open the top defect.
5. Read proof, reproduction summary, and related artifacts.
6. Click `Route to Bug`.

### Approval Flow

1. Open `Review`.
2. See two approval items under `Approvals`.
3. Open the first item.
4. Read what approval authorizes next.
5. Click `Approve` or `Decline`.

## First Slice Recommendation

The best first slice is not full perfection.

It is:

1. canonical `Review` tab
2. Dream batches and approval asks first
3. clear batch and item model
4. explicit dispositions
5. email only as summary plus deep-link

After that:

- add QA and design-review sources
- add dedupe and merge flows
- add richer routing actions

## Open Questions

- Should `Review` become the default launch tab when there are outstanding asks?
- Which ask types should be in the first release versus a later release?
- Should `Approvals` be a first-class filter, a group, or both?
- How much batch summary should be visible before the screen feels crowded?
- What is the cleanest deep-link contract from email into a specific batch or review item?

## Working Conclusion

CodexDashboard should not treat email as the canonical home for candidate tasks.

The better product direction is:

- build a first-class `Review` surface for incoming asks
- keep `Tasks` clean for committed work
- keep `Jobs` clean for recurring automation
- use email as a delivery channel back into `Review`

That gives the human one place to decide what new work means, while preserving `Tasks` as a trustworthy picture of work that is already real.
