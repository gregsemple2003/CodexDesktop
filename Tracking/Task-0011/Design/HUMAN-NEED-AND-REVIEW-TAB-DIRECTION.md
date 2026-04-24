# Human Need And Review Tab Direction

## Why This Document Exists

This note is intentionally long.

It exists because the repo is now at risk of making a second major product mistake.

The first possible mistake was:

- turning the dashboard into a thin operator console for active tasks and jobs

The second possible mistake is:

- treating incoming asks as side-band clutter and making the human manage them through email, memory, scattered packets, and ad hoc review rituals

That second mistake is now realistic because the repo is producing, or will soon produce, multiple streams of candidate work and decision requests:

- Dream runs that propose new tasks
- one-off interface review passes that produce design fixes
- QA or adversarial sweeps that produce defect candidates
- general-design or first-principles passes that produce product gaps
- explicit approvals that block or release downstream work
- runtime anomalies that need human judgment instead of silent drift

All of that material wants something from the human.

That means the product no longer has only one human-facing center.

It has at least two:

- `Tasks` for committed work
- `Review` for incoming asks that still need human judgment

This document is the human-needs-rooted general design brief for the future `Review` tab.

It is task-local because the shape is now being pressure-tested inside [Task-0011](../TASK.md).

It is grounded in:

- [Task-0011](../TASK.md)
- [PLAN.md](../PLAN.md)
- [Task-0009](../../Task-0009/TASK.md)
- [HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md](../../Task-0009/Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md)
- [REVIEW-AND-INTAKE-SURFACE-DIRECTION.md](./REVIEW-AND-INTAKE-SURFACE-DIRECTION.md)
- [2026-04-23-REVIEW-SURFACE-SURVEY.md](../Research/2026-04-23-REVIEW-SURFACE-SURVEY.md)
- [Design/GENERAL-DESIGN.md](../../Design/GENERAL-DESIGN.md)
- [GENERAL-DESIGNER.md](../../../../Users/gregs/.codex/Orchestration/Prompts/GENERAL-DESIGNER.md)
- [INTERFACE-DESIGNER.md](../../../../Users/gregs/.codex/Orchestration/Prompts/INTERFACE-DESIGNER.md)
- [FIRST-PRINCIPLES.md](../../../../Users/gregs/.codex/Orchestration/Processes/FIRST-PRINCIPLES.md)

It is not a visual spec.

It is the argument for:

- what the `Review` tab is for
- which human job it must reduce
- what kinds of asks belong there
- what the screen must and must not become
- why email should be demoted to delivery instead of being the canonical state home
- how the product should integrate multiple ask pipelines without flattening them into one noisy inbox

## Primary Read

If someone only reads one page of this document, the most important point is this:

CodexDashboard is becoming a local cockpit for AI work that does not only run.

It also:

- proposes work
- requests approvals
- surfaces anomalies
- emits improvement ideas
- produces follow-on asks that the human must review

The human does not need those asks to arrive as:

- loose email prose
- scattered markdown artifacts
- hidden packet folders
- backend events that only make sense to an operator

The human needs:

- one place where new asks arrive in a trustworthy way
- one place where each ask clearly states what it wants
- one place where provenance and evidence stay visible
- one place where the next safe decision is obvious
- one place where reviewed work can become real work without the human carrying queue state in memory

That place is the future `Review` tab.

The `Review` tab should become the product’s human decision surface.

The `Tasks` tab should remain the product’s committed-work supervision surface.

The `Jobs` tab should remain the recurring automation surface.

Those are three different jobs.

The product gets stronger when it treats them as three related surfaces with different meanings, not as one overloaded console.

## The Human Need This Repo Fulfills

### The Plain-Language Promise

The plain-language promise of the repo should evolve into:

- `I can trust one local dashboard to show what work is real, what new asks are waiting, and what I should do next.`

That is a better promise than:

- `a token dashboard`
- `a jobs panel`
- `a review-email generator`
- `a local control plane`

Those narrower descriptions point at subsystems.

They do not name the human job.

### The Human Job

The human job here is not:

- monitor infrastructure
- read every packet by hand
- remember yesterday’s proposed work
- translate internal runtime state into actual decisions

The human job is:

- supervise, review, and steer a growing body of partly autonomous work without turning that supervision into a second full-time job

That job naturally splits into two major lanes:

1. `Committed Work Supervision`
   - what is real
   - what is running
   - what is blocked
   - what is sleeping
   - what is ready to dispatch
2. `Incoming Ask Review`
   - what new asks arrived
   - what kind of asks they are
   - what evidence they carry
   - what action they want
   - whether they should become real work, trigger approval, route to a bug, or be dismissed

The `Tasks` tab owns the first lane.

The `Review` tab should own the second lane.

### The Hard Truth

Without a good review surface, the system will quietly externalize several burdens onto the human:

- remembering what new asks already came in
- remembering which asks were already partially reviewed
- remembering whether a candidate became a task or was just discussed
- remembering where the supporting packet lives
- remembering what promotion or approval would actually do
- remembering which source pipeline produced the ask and what level of trust to assign it

That memory burden is not a small annoyance.

It is a product failure.

The backend can technically preserve all this state and the product can still fail if the human has to reconstruct it manually every day.

### The Emotional Part Matters

This repo is not only reducing clicks.

It is reducing a more specific kind of hidden cost:

- low-grade anxiety that useful asks are getting lost in side channels
- reluctance to start more automation because review will become messy later
- fear that a pile of incoming suggestions will dilute focus rather than sharpen it
- frustration from having to re-read too much before making one simple decision
- distrust created when the system treats proposed work and committed work as if they are the same thing

If reviewing incoming asks feels like triaging a messy inbox and reconstructing the product’s intent from scraps, the product is failing.

If the human starts avoiding Dream, QA, or interface-review passes because their output is annoying to process, the product is failing.

If the human cannot quickly tell what a proposed item wants and what would happen if they act on it, the product is failing.

### Truth, Compassion, And Tolerance

The shared first-principles frame applies directly here.

#### Truth

Truth means:

- the review surface must not pretend a weak suggestion is a ready task
- the surface must not hide missing evidence behind polished UI
- the surface must not collapse different ask types into one generic queue label
- the surface must not imply that reading a message equals resolving the underlying ask

#### Compassion

Compassion means:

- the human should not have to remember state the system could carry
- the product should explain what an ask wants without forcing an artifact hunt
- recovery from interruption should be easy
- repeated human decision work is real cost, not neutral traffic

#### Tolerance

Tolerance means:

- incoming asks will vary in clarity and quality
- some source pipelines will be cleaner than others
- some asks will be partial, overlapping, or ambiguous
- the product should support imperfect inputs without making the human become an archaeologist

The design challenge is not to require perfect pipeline output.

It is to present imperfect but meaningful asks in a way that still supports good human judgment.

## What CodexDashboard Is Becoming

### The Old Story

The repo started as a local token telemetry and hotkey overlay tool.

That story was narrow and coherent:

- tell the human how hard Codex is burning tokens

### The Current Story

The repo is now clearly more than that.

It is becoming:

1. a local session telemetry surface
2. a local automation and scheduling surface
3. a task supervision surface
4. a review-and-improvement surface for incoming asks

The fourth part is the new pressure.

If the product does not define it well, the dashboard will drift into one of two bad shapes:

- a task monitor plus an email habit
- or a single overloaded work console that mixes committed work and provisional asks into one muddy stream

### The Better Story

The better story is:

- CodexDashboard is a Windows-first local supervision cockpit for Codex work across active tasks, incoming asks, and recurring automation.

That becomes even clearer when each top-level surface has a clean job:

- `Usage`
  - how hard the system is burning
- `Jobs`
  - whether recurring automation is alive and healthy
- `Review`
  - what new asks have arrived and what they want from the human
- `Tasks`
  - what committed work is active, blocked, sleeping, or ready

That is a better product story than:

- one giant surface that tries to expose every noun in the orchestration layer

### Why Review Needs Its Own Center

Incoming asks are not a mere subset of tasks.

They have their own meaning, cadence, and risk.

An incoming ask can be:

- a suggested task
- an approval request
- a runtime exception
- a design proposal
- a bug candidate
- a follow-on research request

Those are not the same as:

- real tasks
- dispatchable tasks
- active runs
- sleeping tasks

If the UI forces them into the same conceptual bucket too early, the human loses the ability to trust the screen at a glance.

That is why `Review` deserves to be a first-class surface instead of a subfilter buried inside `Tasks`.

## Why Email Is The Wrong Center

### Email Is Good At Delivery

Email is useful for:

- notifying the human that new review work exists
- sending a daily digest
- summarizing a batch
- linking back into the dashboard

That is real value.

The dashboard should keep using email for those jobs.

### Email Is Bad At Canonical Review State

Email is weak for:

- dedupe
- provenance collapse
- batch comparison
- routing
- merge decisions
- explicit dispositions
- partially reviewed state
- reopening
- source-specific actions

Email also blurs a dangerous line:

- reading something looks too much like resolving it

But the human job is usually not:

- read message and forget it

It is:

- make an explicit decision with durable consequences

That difference is too important to leave in the inbox.

### Email Increases Hidden Memory Burden

Once email becomes the home, the human has to remember:

- whether this ask has already been handled in the app
- whether this email represents a task that is already real
- whether the same ask arrived through another channel
- whether the supporting packet said anything that was not included in the message body
- whether this digest item is stale or still actionable

That is exactly the sort of hidden state the product should absorb instead of exporting.

### Email Is A Great Doorbell

The right metaphor is:

- email is the doorbell
- `Review` is the front door and the foyer where the human actually decides what to do

The system can ring the bell.

It should not ask the human to live in the bell.

## Human Burden Taxonomy For Incoming Asks

The easiest way to design `Review` well is to name the human burdens it should reduce.

The tab should not be designed only around data structures.

It should be designed around the burdens the human is currently paying or will soon pay.

### Burden 1: Losing Track Of What Is New

When asks come in through packets, docs, and email digests, the human has to reconstruct:

- what arrived today
- what arrived since the last review session
- what is still unresolved
- what is stale versus genuinely fresh

The `Review` tab should reduce this by keeping incoming asks in one canonical surface with explicit freshness and explicit unresolved state.

### Burden 2: Confusing Suggestion With Commitment

This is the core semantic risk.

If proposed work looks too much like real work:

- the human may believe something is already part of the active system when it is not
- the human may assume it was already approved
- the product may overstate its own reliability

The review surface must preserve the difference between:

- proposed
- approved
- promoted
- routed
- committed

### Burden 3: Reading Too Much Before Making One Decision

A good review surface should absorb the first sixty seconds of comprehension.

If the human must open three markdown files and two emails to answer:

- what is this item
- why did it arrive
- what happens if I click approve

then the product is not doing enough of the comprehension work.

### Burden 4: Losing Provenance

Without visible provenance, the human has to remember:

- whether this came from Dream
- whether QA found it
- whether this was generated from a design exercise
- whether it belongs to this repo or another one
- what batch it came from

Provenance needs to be visible enough that the human can trust the item quickly.

### Burden 5: No Durable Disposition Model

Without explicit review state, all the human really has is:

- unread
- read
- remembered
- forgotten

That is not enough.

The human needs state that tracks what actually happened:

- deferred
- duplicate
- promoted
- dismissed
- approved
- routed

That is not metadata polish.

That is the durable memory of the review workflow.

### Burden 6: Batch Blindness

A daily Dream run for one repo may produce seven candidate tasks.

A QA sweep may produce four defect asks.

An interface review may produce a visually coherent cluster of five UI asks.

If the product flattens everything into a single chronological list, the human loses the context that made the items intelligible in the first place.

The `Review` tab must preserve:

- which asks belong together
- which run or pass produced them
- what the batch summary was

### Burden 7: Weak Action Semantics

The human should not have to guess:

- what `Promote` means
- whether `Approve` creates or releases something
- whether `Dismiss` destroys history or simply marks an item closed
- whether `Route to Bug` creates a new task or links an existing one

Each action needs explicit consequence framing.

### Burden 8: Duplicate And Overlap Confusion

In the real system, similar asks will arrive through different pipelines.

Example:

- Dream proposes a task
- QA finds a symptom that points at the same missing need
- interface review notices the same missing affordance in the product surface

If the system cannot help the human see that these may be related, the human pays a large dedupe and merge tax.

### Burden 9: Interrupted Review Sessions

Humans do not process all incoming asks in one perfect uninterrupted block.

They get interrupted.

They return later.

They forget which items were only skimmed versus actually resolved.

The `Review` tab must support re-entry with clear state, not just a recency feed.

### Burden 10: Fear Of Creating More Work Than The System Can Absorb

If incoming asks are surfaced badly, the human will start avoiding the very activities that improve the system:

- Dream runs
- QA sweeps
- interface audits
- design ideation

That means the review surface needs to make incoming work feel governable, not explosive.

### Burden 11: Hidden Decision Cost

Every review item silently asks the human to spend:

- attention
- working memory
- trust
- appetite for future automation

The product should therefore make each decision cheaper without becoming dishonest.

### Burden 12: Weak Longitudinal Trust

The product will win trust if, day after day, the human feels:

- the system remembers what came in
- the system remembers what I already did
- the system keeps related things together
- the system does not make me reconstruct old meaning from scratch

That is the long-term job.

## The Product Jobs The Review Tab Must Support

The `Review` tab should be designed around recurring human jobs, not just source pipelines.

### Job 1: Rapid Morning Orientation

The human opens the app after some time away and wants fast answers:

- what new asks came in
- how many matter now
- which source produced them
- whether anything urgent needs immediate judgment

This is likely the most common and important job.

### Job 2: Process A Dream Batch

The human wants to review a batch of candidate tasks from one repo.

They need to answer:

- which candidates seem strong
- which ones are duplicates or weak
- which should become real tasks

The screen must support evaluating a batch as a batch, not just as isolated rows.

### Job 3: Approve Or Reject A Gating Ask

An ask may block downstream work:

- a plan approval
- a closeout approval
- a task approval

The human needs to know:

- what the approval authorizes
- what artifacts changed
- what happens next if they approve

This is not the same as promoting a candidate.

### Job 4: Route A Defect Or Design Gap

A QA or design review item may not want direct promotion into a task.

It may want:

- route to bug
- route to design
- merge with existing task

The review surface needs to support those actions as first-class outcomes.

### Job 5: Defer Or Snooze Without Losing State

Sometimes the right decision is:

- not now

That should be handled as a safe, explicit disposition, not as leaving the item unread and hoping memory holds.

### Job 6: Revisit A Partially Reviewed Batch

The human may skim a batch, act on two items, defer three, and want to revisit later.

The surface should make that progression obvious.

### Job 7: Detect Pattern-Level Drift

The human may need to notice broader trends:

- too many asks are unresolved
- too many runtime exceptions are arriving
- one repo is producing a flood of low-confidence Dream suggestions
- approvals are piling up and slowing throughput

The review surface should support this without becoming a business intelligence tool.

## The Review Tab Must Not Become

Before defining the right shape, it helps to define the wrong ones.

### Wrong Shape 1: The Canonical Email Ritual

This shape says:

- the real queue is in email
- the app is just a later convenience

That fails because the human ends up juggling:

- inbox state
- product state
- markdown artifacts
- memory

### Wrong Shape 2: The Giant Catch-All Inbox

This shape says:

- every incoming thing goes into one list
- the human can filter later

That fails because it forces the human to:

- infer meaning from mixed ask types
- manually reconstruct relatedness
- mentally separate committed work from provisional asks

### Wrong Shape 3: The Operator Queue Console

This shape talks in system terms:

- event type
- pipeline id
- trigger kind
- target entity
- workflow status

Those may be valid backend nouns.

They are not humane default UI language.

### Wrong Shape 4: The Fake Review Feed

This shape looks active but has no durable dispositions.

It tells the human:

- you can read this
- maybe archive it
- maybe click through

But it does not actually track what review decision happened.

That is weak and untrustworthy.

### Wrong Shape 5: The Premature Task Generator

This shape treats every ask like a task candidate and pushes direct promotion by default.

That fails because some asks should instead:

- become a bug
- become a design note
- become an approval
- merge into an existing task
- be dismissed

### Wrong Shape 6: The Comforting But Vague Dashboard

This shape keeps everything visually clean by hiding distinctions:

- no clear ask type
- no clear batch relationship
- no clear recommended action
- no clear consequence of acting

That is compassionate-looking but not truthful.

## The Right Product Shape

### One-Sentence Product Shape

The right shape is:

- a calm, explicit review cockpit that groups incoming asks by source and meaning, preserves provenance, and makes the next human decision obvious without confusing provisional asks for committed work

### Review As Human Decision Surface

The tab should feel like:

- the place where the human handles new requests on the system’s behalf

It should not feel like:

- a passive notification center
- a mail client
- a task table with extra rows

### Primary Product Promise Of The Tab

The plain-language promise of the `Review` tab should be:

- `When the system wants something from me, I can come here, understand it quickly, and act without losing context.`

That promise is specific.

It is also demanding.

It requires:

- truthful distinctions
- visible provenance
- durable dispositions
- clear actions
- preserved batch context

## The Object Model The Human Will Feel

The backend may call things many names.

The human should feel a simpler model.

### The Human Sees Batches

A batch is:

- one coherent packet of asks from one source

Examples:

- `Dream · CodexDashboard · 7 candidate tasks`
- `QA Sweep · CodexDashboard · 3 defects`
- `Interface Review · CodexDashboard · 5 UI asks`

The human should feel that these belong together.

### The Human Sees Items

Each item is:

- one ask inside the batch

Each item should answer:

- what it is
- why it arrived
- what evidence supports it
- what action is being requested

### The Human Sees Dispositions

The human should feel that an item can be:

- unresolved
- deferred
- dismissed
- merged
- promoted
- approved
- routed

That is a more humane and accurate model than:

- unread
- read

## Review Item Taxonomy

The first version of the review surface should recognize that not all asks are the same.

### Candidate Task

A proposal that may deserve promotion into a real task.

Common sources:

- Dream
- general design
- first-principles passes

Typical actions:

- promote to task
- defer
- dismiss
- merge with existing task

### Approval Ask

A request for the human to explicitly authorize a next step.

Common sources:

- task planning
- task closeout
- reopen requests

Typical actions:

- approve
- decline
- defer

### Bug Or Defect Ask

A candidate defect or quality problem that likely wants durable tracking.

Common sources:

- QA sweeps
- runtime anomalies
- interface review

Typical actions:

- route to bug
- merge with existing bug task
- dismiss
- defer

### Design Gap Ask

A product or interface gap that likely wants design ownership rather than immediate code-task ownership.

Common sources:

- interface review
- general designer
- first-principles analysis

Typical actions:

- route to design
- promote to task
- merge with existing task
- dismiss

### Runtime Exception Ask

A runtime observation that may need intervention, acknowledgement, or follow-on work.

Common sources:

- supervision runtime
- task-dispatch monitoring

Typical actions:

- acknowledge
- poke related run
- interrupt related run
- route to task
- dismiss as expected

### Follow-Up Research Ask

A recommendation to investigate or survey before implementation.

Typical actions:

- promote to research task
- merge with existing task
- defer
- dismiss

The UI should not force the human to learn this taxonomy all at once.

But the system must know it, because action semantics depend on it.

## Batch Design

### Why Batch Matters

Batch context often explains why an item exists.

Without batch context, the human loses:

- source confidence
- comparative context
- recency meaning
- batch-level summaries that make the items easier to judge

### What A Batch Row Should Communicate

A batch row should quickly answer:

- what source produced this
- which repo it belongs to
- how fresh it is
- how many asks it contains
- whether anything inside needs urgent review

### Good Batch Row Examples

- `Dream · CodexDashboard · 7 candidate tasks · fresh`
- `QA Sweep · CodexDashboard · 3 defects, 1 duplicate · needs review`
- `Approvals · CodexDashboard · 2 blocking asks · waiting on you`

### Bad Batch Row Examples

- `pipeline_run_3187`
- `review event batch`
- `4 items`

Those are too abstract and force too much inference.

### Batch-Level Actions

The first version of the surface may not need heavy batch actions.

But useful later actions could include:

- mark all viewed
- expand all unresolved
- open source packet
- defer batch until tomorrow

Batch actions must remain careful.

They should not make it too easy to accidentally flatten item-level review.

## The Core Screen Layout

### Overall Layout

The strongest first-direction desktop layout is:

- top header row
- summary cards directly below
- main three-column body

Three-column body:

- left: batch rail and scope controls
- center: grouped review queue
- right: persistent detail pane

This layout works because the review job needs three things visible at once:

- the broad packet landscape
- the item-level queue
- the selected item’s explanation and actions

### Why Three Columns Win

A single-column feed is too flattening.

A two-column layout helps, but often forces batch context to become hidden or secondary.

Three columns allow:

- persistent source context
- persistent queue context
- persistent detail context

That lowers orientation cost.

### Header Row

The header should show:

- `Review`
- a freshness sentence
- scope filters
- maybe a lightweight `Open latest digest` or `Refresh` action

The freshness sentence matters because incoming asks are time-sensitive in a different way than tasks.

Example:

- `Showing 14 unresolved asks from 5 fresh batches across 3 repos.`

### Summary Cards

The summary cards should answer:

- what kind of review burden exists now
- where the human’s attention should go first

Suggested first cards:

- `Needs review`
- `Approvals`
- `New work`
- `Exceptions`
- `Deferred`

Potential later cards:

- `Duplicates`
- `Stale`
- `Resolved today`

### The Batch Rail

The batch rail should likely be the left-most structure.

It should support:

- source filter
- repo filter
- freshness filter
- selected batch

The batch rail should not feel like a file tree.

It should feel like a high-signal list of incoming packets.

### The Queue Pane

The center pane is where actual review work lives.

It should show items within the selected scope, grouped by meaning.

Likely groups:

- `Needs review now`
- `Approvals`
- `New work proposals`
- `Exceptions and anomalies`
- `Deferred`

This grouping supports the human job better than one flat date-sorted feed.

### The Detail Pane

The detail pane is where the product absorbs the first minute of comprehension.

It should be rich enough that the human can often decide without leaving the tab.

But it should not become a full markdown reader.

The split should be:

- detail pane = first minute of understanding
- source artifacts = deeper evidence and history

## What The Human Should See First

The first glance matters more than almost anything else.

The product wins if the human can open `Review` and immediately answer:

- how much incoming work exists
- what kind of work it is
- whether anything urgent is blocking real progress
- what one thing should be reviewed first

That means the default visual emphasis should go to:

- unresolved count
- approvals
- urgent exceptions
- fresh Dream or QA batches
- explicit reason lines

Not:

- raw timestamps everywhere
- ids
- pipeline internals

## Screen States In Detail

The review surface has to be truthful across more than just the happy path.

### State: No Review Work Yet

The screen should explain:

- no incoming asks exist yet
- what kinds of asks will appear here later
- what causes them to appear

Possible copy:

- `No review items yet. Dream batches, approvals, QA findings, and design asks will appear here when the system produces them.`

### State: Loading

Use structural skeletons that match:

- summary cards
- batch rail
- queue rows
- detail pane

This teaches the composition instead of showing a meaningless spinner.

### State: Fresh Populated

This is the ideal working state.

The screen should feel:

- calm
- dense enough to be useful
- clear enough to act from immediately

### State: Slightly Stale

The product should preserve honesty:

- `Showing the last trusted review snapshot from 7 minutes ago.`

This keeps the human oriented without sounding like an error.

### State: Stale Enough To Distrust

At a certain age, the surface should explicitly lower confidence:

- `Review data is stale. The last trusted snapshot is from 2:14 PM. Some asks may have changed elsewhere.`

Actions that depend on live state may need reduced prominence or disabled treatment.

### State: Backend Unavailable

The surface should not collapse to empty.

It should show:

- that live refresh failed
- the age of the last trusted snapshot
- what the human can still inspect safely

### State: Candidate-Heavy But No Active Tasks

This state should not look like:

- nothing is happening

It should look like:

- no committed work needs review here, but there is fresh proposed work worth triaging

### State: Approval-Heavy Day

Some days may have few candidate tasks but several approvals.

The screen should make this obvious.

It should not hide approvals beneath a generalized unresolved count.

### State: Exception-Heavy Day

If runtime exceptions or QA findings cluster, the screen should say so clearly.

This may need stronger color and stronger top-summary treatment than ordinary candidate-task review.

### State: Everything Resolved

When the human has cleared the review queue, the screen should feel:

- finished
- trustworthy
- calm

Possible copy:

- `No unresolved asks right now. New review work will appear here when a source pipeline needs your judgment.`

## Section-Level Semantics

### Needs Review Now

This group should always be expanded by default when it has content.

These are the asks most likely to deserve immediate judgment.

### Approvals

Approvals are special because they often gate other work.

They should be visible enough that the human can tell:

- these are not speculative asks
- these are permission checkpoints

### New Work Proposals

This group includes:

- candidate tasks
- some design proposals
- some follow-up research asks

This group should feel slightly more exploratory than approvals, but still concrete.

### Exceptions And Anomalies

This group should feel more serious.

It often represents:

- a system state that may need intervention
- a defect cluster
- a runtime mismatch

### Deferred

Deferred items should remain visible but lower-pressure.

The product must not make deferral equal disappearance.

## Row Design

Each row should support glanceable understanding.

### Required Row Fields

At minimum, each row should show:

- human-readable title
- short one-line summary
- ask type or source badge
- provenance label
- freshness
- recommended action
- one reason line

### Why Reason Lines Matter

The reason line is one of the highest-value parts of the row.

It should answer:

- why am I seeing this
- why does this matter

Examples:

- `Dream thinks this should become a real task because the burden is recurring and still unresolved.`
- `QA found a reproducible defect with proof and no matching open bug task.`
- `This approval is blocking a task from becoming dispatchable.`
- `This exception has no valid wait reason and likely needs intervention.`

### Provenance Treatment

The row should keep provenance visible with labels like:

- `Dream`
- `QA`
- `Interface Review`
- `General Design`
- `Approval`
- `Runtime`

This should not dominate the row, but it should never be hidden behind a tooltip.

### Recommended Action Treatment

The row can hint at the most likely next action:

- `Promote`
- `Approve`
- `Route to Bug`
- `Review`
- `Open Source`

But the surface should not prematurely perform the action.

The detail pane should remain where the human sees the consequence clearly.

## Detail Pane Design

The detail pane must make one item legible without opening five artifacts.

### Detail Pane Sections

Recommended sections:

- `Summary`
- `Why this arrived`
- `Source and provenance`
- `Evidence`
- `Recommended action`
- `What happens next`
- `Related work`
- `Actions`

### Summary

This should be the fastest plain-English explanation of the ask.

The human should understand:

- what this is
- why it exists

without reading the source packet immediately.

### Why This Arrived

This section should answer:

- what source produced the item
- what pattern or burden triggered it
- whether it is routine or unusual

### Source And Provenance

This section should provide:

- source pipeline
- repo
- batch
- creation time
- related source artifacts

### Evidence

Evidence should be strong enough to support trust, but not so broad that the detail pane becomes a transcript browser.

Examples:

- burden summary
- screenshots
- reproduction summary
- related docs
- excerpted rationale

### Recommended Action

This section explains the likely best next move and why.

It should not be a vague AI opinion.

It should connect to the human’s actual options.

### What Happens Next

This is where the product explains action consequences before click.

Examples:

- `Promote to Task` will create a new durable task skeleton and remove this item from the candidate queue.
- `Approve` will unblock the next lifecycle step for the linked task.
- `Route to Bug` will create or link a bug-tracking task and close this review item as routed.

### Related Work

This section should surface:

- existing tasks
- related bugs
- related design notes
- existing duplicates or merges

That helps keep the human from recreating work unnecessarily.

## Action Model

The action model is meaning-bearing.

It should be designed as carefully as any status taxonomy.

### Actions Must Be Ask-Type Aware

Do not show one universal action row for every ask type.

Different ask types deserve different choices.

### Likely First Actions

- `Promote to Task`
- `Approve`
- `Decline`
- `Route to Bug`
- `Route to Design`
- `Merge with Existing Task`
- `Defer`
- `Dismiss`
- `Open Source Packet`

### Defer

`Defer` is not a weak action.

It is an important humane action because:

- not all good asks deserve immediate work
- humans need a safe way to preserve state without pretending to resolve

### Dismiss

`Dismiss` should be explicit and slightly heavier than defer.

The product should clarify whether dismissal means:

- not worth acting on
- duplicate
- weak or unsupported

If possible, dismissal should preserve why it was dismissed.

### Merge With Existing Task

This becomes valuable once multiple pipelines start surfacing overlapping asks.

The product should help the human avoid accidental duplication.

### Source Packet And Artifact Actions

The product still needs deeper context paths:

- `Open Source Packet`
- `Open Related Task`
- `Open Design Note`

Those should be close at hand, but secondary to the main decision.

## Dispositions

The review tab needs a real disposition model.

### Why Dispositions Matter

Dispositions are durable memory.

Without them, the review surface degrades into:

- message viewer
- semi-random list of things the human vaguely remembers touching before

### Minimum Useful Dispositions

- `new`
- `needs_review`
- `deferred`
- `duplicate`
- `dismissed`
- `approved`
- `promoted_to_task`
- `routed_to_bug`
- `routed_to_design`

### How Dispositions Should Feel

They should feel like:

- trustworthy outcomes

Not:

- abstract workflow internals

### Visible Versus Hidden

The default surface likely does not need to show every raw disposition label on every row.

But it should be able to communicate:

- unresolved
- deferred
- approved
- routed
- promoted
- dismissed

in a humane, legible way.

## Source Pipeline Integration

The review tab should accept that not all pipelines are equal.

### Dream

Dream is likely the first and most visible source.

Dream asks often come in batches and often want:

- promote
- dismiss
- merge
- defer

Dream rows should feel exploratory but not flimsy.

### Interface Review

Interface-review asks often care about:

- UI truth
- humane copy
- component semantics
- affordance clarity

These asks may want:

- route to design
- promote to task
- merge with existing task

### QA Sweep

QA asks often care about:

- reproducible defects
- strange edge cases
- runtime anomalies

These asks often want:

- route to bug
- merge with existing bug task
- dismiss if invalid

### General Design

General-design asks often care about:

- missing product needs
- weak product framing
- hidden human burden

These asks may want:

- route to design
- promote to task
- follow-up research

### Runtime Supervision

Runtime supervision asks may represent:

- run anomalies
- sleeping runs
- invalid wait contracts
- state mismatches

These asks may want:

- acknowledge
- poke
- interrupt
- create follow-on task

The product should not pretend one generic action vocabulary fits all of these.

## Trust Model

Trust is the main job of the `Review` tab.

### What Builds Trust

Trust comes from:

1. clear distinction between provisional asks and committed work
2. clear provenance
3. explicit action consequences
4. visible freshness
5. durable review state
6. easy path to source evidence

### What Breaks Trust

Trust breaks when:

- email and app disagree
- items disappear without a clear disposition
- a promoted item still looks provisional
- duplicate asks flood the queue with no merge path
- weak suggestions are styled like high-confidence approvals
- the product makes a decision look done when the underlying ask is still unresolved

### Freshness

Freshness must be visible and calm.

The surface should say when it is:

- fresh
- slightly stale
- stale enough to distrust
- unavailable

### Provenance And Confidence

If the system can eventually estimate confidence or evidence strength, that can be useful.

But it must not replace provenance.

The human should always know:

- where the ask came from

before they know:

- how some model scored it

## Humane Error And Recovery Model

The review surface should prevent confusion when possible and explain it well when not possible.

### Preventable Errors

The best errors are prevented.

Examples:

- do not show `Promote to Task` on an item already promoted
- do not show `Approve` on an item that is only a suggestion
- do not show `Route to Bug` if the item is already routed and linked
- do not let the human accidentally confuse a batch view with a task list

### When Error Is Still Needed

Error copy should answer:

- what happened
- what it means
- what to do next

Good examples:

- `Review items could not refresh from the backend. The last trusted snapshot is from 3:24 PM.`
- `This ask cannot be promoted because it has already been converted into Task-0012.`
- `This approval item is stale because the related task changed after the approval request was created.`

### Recovery Paths

Major recovery paths should include:

- retry refresh
- open source packet
- reopen related task
- view related duplicate
- re-run source batch later

### Interrupted Review Sessions

The product should make it safe to stop halfway through review and return later.

That means:

- unresolved items remain clearly unresolved
- deferred items are visible
- partially reviewed batches remain legible

## First-Run And Recurring-Use Flows

### First Encounter With Review

The human opens the tab for the first time.

The product should explain:

- what this tab is for
- what kinds of asks appear here
- how it differs from `Tasks`

The first empty state is therefore important.

### Normal Daily Review Flow

Ideal daily flow:

1. Open `Review`.
2. Read the summary strip.
3. Check `Needs review now`.
4. Select the most important batch or item.
5. Read the detail pane.
6. Act.
7. Repeat until the queue is calm.

This should take minutes, not a long spelunking session.

### Dream Review Flow

Ideal flow:

1. Open a Dream batch.
2. Scan all proposed items together.
3. Compare them within their batch context.
4. Promote the strongest items.
5. Dismiss or defer the weaker ones.

### Approval Flow

Ideal flow:

1. Open `Approvals`.
2. Select the top item.
3. Read what the approval authorizes.
4. Read which artifact changes matter.
5. Approve or decline.

### QA Flow

Ideal flow:

1. Open a QA batch.
2. Review the top defect ask.
3. Read evidence and reproduction summary.
4. Route to bug or merge with existing bug task.

## Scenario Walkthroughs

### Scenario 1: A Dream Daily Run Produces Seven Candidates

The worst experience is:

- seven emails or seven undifferentiated rows

The right experience is:

- one batch visible in the left rail
- one batch summary row
- seven grouped items in the center queue
- one detail pane that explains each candidate without packet hunting

The human can move through the set coherently and finish with:

- two promotions
- three dismissals
- two deferrals

The batch then becomes a durable reviewed artifact, not a vague memory.

### Scenario 2: QA And Dream Surface The Same Underlying Gap

The surface should help the human notice:

- these are different asks with shared causal roots

The human may then:

- merge the new ask into an existing task
- route only one of them forward

Without this support, the human pays a duplicate-detection tax every day.

### Scenario 3: Approval Requests Pile Up

The review surface should not bury approvals among speculative proposals.

Approvals often gate real work.

The UI should therefore let the human answer:

- what is blocking downstream execution right now

### Scenario 4: The Human Returns After Two Busy Days

The product should make re-entry humane.

The human should not need to recall:

- which asks were new yesterday
- which were already deferred
- which were already promoted

The surface should already know.

### Scenario 5: The Backend Is Partly Unavailable

The screen should remain useful for orientation and review of last-known asks.

It should not pretend:

- there is no queue

when the queue is merely temporarily stale.

### Scenario 6: One Ask Is Weak And Unsupported

The product should let the human dismiss or defer it safely.

It should not pressure the human into promoting weak asks simply because the UI only supports forward motion.

### Scenario 7: A Runtime Exception Needs Human Judgment

The item should not look like:

- a task candidate

It should look like:

- an operational ask that needs judgement now

Its actions may be:

- acknowledge
- poke related run
- interrupt
- open related task

### Scenario 8: The Human Wants To Clear The Queue Calmly

The surface should support a satisfying, bounded feeling of progress.

As items are promoted, routed, or dismissed, the summary should update in a way that reinforces:

- the queue is becoming more truthful and more manageable

## Copy System

### Tone

The tone should be:

- direct
- calm
- precise
- non-theatrical

It should not feel like:

- a chat assistant
- a ticketing bureaucracy
- an infrastructure console

### Good Nouns

Prefer:

- review
- ask
- candidate
- approval
- exception
- batch
- evidence
- promote
- route
- defer
- dismiss

Be careful with:

- intake
- workflow
- trigger
- entity
- pipeline event

Those are sometimes useful internally, but weak as the default human-facing nouns.

### Good Secondary Sentences

Examples:

- `Dream proposes this as real work because the burden is recurring and still unresolved.`
- `This approval blocks the related task from becoming dispatchable.`
- `QA found a reproducible defect with proof and no matching open bug task.`
- `This item was deferred yesterday and still has no related real task.`
- `Promotion will create a new task skeleton and remove this item from the candidate queue.`

### Bad Secondary Sentences

- `Pending review event`
- `Awaiting action`
- `Item unresolved`
- `Workflow state active`
- `Suggested for next steps`

These are too generic to reduce human burden.

## Interface Semantics

### Visual Atmosphere

The visual tone should feel:

- serious
- local
- trustworthy
- operational without becoming cold

It should not feel like:

- a SaaS inbox clone
- a generic kanban
- a flat admin template

### Color Semantics

Likely semantic treatments:

- warm amber for `Needs review now`
- muted blue or cyan for `Approvals`
- softer olive or green for `New work`
- coral or orange-red for `Exceptions`
- subdued slate for `Deferred`

These colors should help orient, not dominate.

### Typography

Typography should make role obvious:

- strong title for the tab
- clear group headers
- readable item titles
- legible reason lines
- restrained monospace only for ids or machine-like metadata

The screen should never feel like a pile of default widgets.

### Icons

Icons should reinforce meaning:

- clipboard-check or review mark for approvals
- spark or wand very carefully for Dream candidates
- bug or warning mark for defects and exceptions
- branch or route mark for routing actions
- clock for deferred

But icons should never be the sole carrier of meaning.

## The Relationship To Other Tabs

### Review

`Review` is where incoming asks wait for human judgment.

### Tasks

`Tasks` is where committed work is supervised.

### Jobs

`Jobs` is where recurring automation and runtime health are tracked.

### Usage

`Usage` is where token burn or system activity is summarized.

### Why These Must Stay Distinct

When tabs do different jobs, the human can build reliable habits:

- open `Review` when I need to process asks
- open `Tasks` when I need to supervise real work
- open `Jobs` when I need to check automation

That is a cleaner mental model than one overloaded workspace.

## Relationship To Existing Task-0009 Direction

The original `Task-0009` framing around `Tasks` is still valuable.

This document does not erase it.

It clarifies something the shorter review note already suggested:

- the dashboard likely needs both a strong `Tasks` tab and a strong `Review` tab

That means the product center is not one screen.

It is one coherent family of high-level surfaces:

- `Review` for incoming asks
- `Tasks` for committed work

## Product Boundaries

### What Belongs In The First Review Tab

- batch visibility
- explicit ask types
- grouped item review
- provenance
- evidence summary
- consequence-aware actions
- durable dispositions
- email deep-link integration

### What Does Not Need To Be In The First Review Tab

- full transcript browsing
- arbitrary query language
- bulk admin tooling
- complete audit history for every item inline
- every possible source pipeline from day one

### What Can Stay Elsewhere

- full markdown artifacts
- raw runtime ids
- backend logs
- deep bug investigation details

Those can remain one click deeper.

## What Should Be Automatic, Visible, Or Hidden

### Should Be Automatic

- batch grouping
- freshness evaluation
- disposition persistence
- provenance capture
- deep-link resolution from email into the correct batch or item

### Should Be Visible

- ask type
- source
- batch
- freshness
- recommended action
- consequence of acting
- whether the item is unresolved, deferred, routed, promoted, or dismissed

### Should Be Hidden By Default

- raw event ids
- raw workflow ids
- transport metadata
- internal orchestration nouns
- low-level backend payloads

## Anti-Patterns

### Anti-Pattern: Treating Email As State

This creates two truths.

That is bad enough on its own.

It is worse when decisions become durable.

### Anti-Pattern: Only Two Outcomes

If the UI only supports:

- promote
- ignore

then it is lying about the variety of real asks.

### Anti-Pattern: Source-Blind Review

If the human cannot tell where something came from, the queue will feel muddy and untrustworthy.

### Anti-Pattern: No Batch Context

If each item appears alone, the product throws away meaning it already has.

### Anti-Pattern: Read Equals Resolved

Reading is not the same as deciding.

The UI should never behave as if it were.

### Anti-Pattern: Review Surface As Ticketing Bureaucracy

The goal is not to create more ritual.

The goal is to make necessary judgment cheap and clear.

### Anti-Pattern: Hiding Consequences Behind Clicks

The human should not have to guess what `Promote`, `Approve`, or `Route` will do.

### Anti-Pattern: Pushing Everything Into Tasks

If `Review` becomes merely a feeder with no independent design bar, the product will quickly leak too much noise into `Tasks`.

## Design Review Checklist

Before calling any `Review` surface good enough, the team should ask:

- Can a human tell what kind of ask they are looking at in five seconds?
- Can a human tell where it came from without opening another artifact?
- Can a human tell whether it is still unresolved?
- Can a human tell what action is being requested?
- Can a human tell what will happen after they act?
- Can a human revisit a half-processed batch without memory work?
- Can a human distinguish proposals, approvals, and exceptions immediately?
- Does the surface feel like a decision cockpit rather than a mail client or ops console?

If any of those answers is no, the surface is not ready.

## Open Tensions

### Tension 1: Richness Versus Calm

Incoming asks can be rich.

The screen still needs to stay calm.

The answer is:

- summarize by default
- reveal evidence progressively

### Tension 2: Batch Context Versus Queue Simplicity

Too much batch detail can clutter the left rail.

Too little batch detail flattens meaning.

This likely needs real visual iteration.

### Tension 3: Flexibility Versus Action Clarity

Different ask types need different actions.

The product should preserve that without making the action system feel sprawling.

### Tension 4: Review Versus Tasks Overlap

Some items will sit near the boundary between:

- incoming ask
- already-real task

The product will need explicit lifecycle semantics so that this boundary stays honest.

### Tension 5: Early Value Versus Full Runtime Integration

The first version can likely deliver value before every source pipeline is fully normalized.

That is healthy as long as the UI is honest about which sources and actions are real in that slice.

## Working Conclusion

The repo’s next human need is not:

- `send me candidate tasks by email`

It is:

- `give me one humane place to review everything the system now wants from me before it becomes real work`

The `Review` tab should be that place.

If this tab lands well:

- Dream becomes easier to use
- QA and interface review become easier to absorb
- approvals become easier to process
- the product gains a cleaner split between incoming asks and committed work
- the human carries less hidden queue state in memory

If this tab lands badly:

- incoming asks will sprawl across email, packets, and memory
- `Tasks` will become noisy or misleading
- the human will avoid useful ask-generating workflows because the review burden feels too high

That is why this tab matters.

That is why the review surface should be treated as a human-centered product feature, not as a thin inbox or a dumping ground for pipeline output.

## Immediate Design Direction

For implementation planning and future UI iteration, the current best direction is:

- a first-class `Review` tab
- three-column desktop layout
- visible batch context
- grouped review queue
- persistent detail pane
- explicit ask types
- explicit dispositions
- consequence-aware actions
- email only as digest and deep-link transport
- strong separation from `Tasks`, which remains the surface for committed work

That is the direction the next visual and implementation passes should pressure-test.
