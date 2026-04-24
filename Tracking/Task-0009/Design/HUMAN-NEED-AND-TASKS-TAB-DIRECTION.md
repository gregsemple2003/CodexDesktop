# Human Need And Tasks Tab Direction

## Why This Document Exists

This note is intentionally long.

It exists because this repo is at risk of doing something very common and very bad:

- building a technically capable orchestration system
- then leaking that orchestration shape directly into the human-facing product surface

That failure mode is especially likely here because the repo already has real infrastructure:

- local session telemetry
- task artifacts
- shared orchestration docs
- intervention packets
- Dream analysis output
- a Temporal-backed jobs control plane

All of that can tempt the product into becoming a system-inspector for insiders.

That is not the human need.

The human need is much simpler and much more demanding:

- help one person keep a growing body of partly autonomous work legible, steerable, and recoverable without turning supervision into a second full-time job

This document is task-local because the shape is not yet stable enough to promote into shared repo design truth.

Scope note as of 2026-04-24:

- canonical intake and review ownership has been split into [Task-0011](../../Task-0011/TASK.md)
- discussion below of Dream candidates on the `Tasks` surface is historical design pressure and provenance context, not current ownership
- the committed-work `Tasks` surface should only show real tasks, including work promoted out of `Review` once it becomes committed

It is a working product brief for [Task-0009](../TASK.md), grounded in:

- [Task-0005](../../Task-0005/TASK.md)
- [Task-0007](../../Task-0007/TASK.md)
- [Design/GENERAL-DESIGN.md](../../../Design/GENERAL-DESIGN.md)
- [GENERAL-DESIGNER.md](../../../../../Users/gregs/.codex/Orchestration/Prompts/GENERAL-DESIGNER.md)
- [INTERFACE-DESIGNER.md](../../../../../Users/gregs/.codex/Orchestration/Prompts/INTERFACE-DESIGNER.md)

It is not a final visual spec.

It is the argument for what the product is for, what the `Tasks` tab must do, what it must not become, and how the repo should think about human need before implementation calcifies around the wrong abstractions.

## Primary Read

If someone only reads one page of this document, the most important point is this:

CodexDashboard is gradually becoming a local supervision cockpit for long-running AI work.

The human does not need:

- more status labels
- more raw logs
- more backend nouns
- more proof that there is a sophisticated system under the hood

The human needs:

- one place that tells the truth about what is happening
- one place that makes the next useful action obvious
- one place that notices when work quietly died
- one place that preserves enough context to recover quickly

That one place is the future `Tasks` tab.

The `Tasks` tab should become the beating heart of the product because tasks are where human intent, agent work, durable artifacts, and future automation all meet.

The dashboard already has:

- `Usage`, which tells the human how hard the system is burning tokens
- `Jobs`, which tells the human whether recurring automation exists and is alive

What it lacks is:

- the place where actual work becomes reviewable, dispatchable, monitorable, interruptible, and recoverable

That is the product hole this task is filling.

## The Human Need This Repo Fulfills

### The Plain-Language Promise

The plain-language promise of this repo should be:

- `I can trust one local dashboard to show what my Codex work is doing, what needs me, and what I should do next.`

That promise is better than:

- `a dashboard for token telemetry`
- `a job scheduler GUI`
- `a transcript browser`
- `a control plane`

Those narrower descriptions name pieces of the system.

They do not name the human job.

### The Human Job

The human job is not `inspect orchestration`.

The human job is:

- supervise meaningful work without constant babysitting

That supervision job contains a cluster of smaller jobs:

- notice when something needs attention
- decide what work to start next
- see what is already running
- spot when something fell asleep
- understand why something is blocked
- recover a working thread quickly
- decide whether a Dream-generated proposal deserves promotion
- keep the whole system pointed at reality instead of at internal convenience

### The Hard Truth

Without a good product surface, AI task work becomes psychologically expensive.

It becomes expensive because the human has to remember too much:

- which tasks exist
- which ones are real versus provisional
- which ones are blocked
- which ones are still safe to enqueue
- where a given thread lives
- whether a silent agent is fine or has quietly dropped the ball

That memory burden is the product problem.

The product is not done when the backend can technically answer those questions.

The product starts to become useful when the human no longer has to reconstruct those answers from many surfaces.

### The Emotional Part Matters

This repo is not just reducing clicks.

It is trying to reduce a specific kind of mental drag:

- low-grade anxiety that the system is quietly doing the wrong thing
- fear that silence means failure
- frustration from having to keep reloading context by hand
- reluctance to dispatch more work because recovery feels too expensive

That matters because the human eventually stops using flows that feel fragile, even if they are powerful.

If dispatch or monitoring feels like babysitting, the product is failing.

If the human avoids starting tasks because checking on them later will be a mess, the product is failing.

If the human cannot tell whether a quiet task is fine or asleep, the product is failing.

### Truth-Seeking And Compassion

Two values matter here and should stay explicit:

- truth-seeking
- compassion for the human

Truth-seeking means:

- the surface must not pretend a task is healthy when it is not
- the surface must not collapse important state distinctions just to feel clean
- the surface must not label candidate work as approved work
- the surface must not present silence as progress

Compassion means:

- the human should not need to remember system history from yesterday to act correctly today
- the default screen should explain itself
- the system should notice likely failure before the human has to
- recovery should start from one obvious action, not a scavenger hunt

These two values work together.

Truth without compassion becomes a cold operator console.

Compassion without truth becomes a soothing lie.

The product needs both.

## What CodexDashboard Is Becoming

### The Old Story

The oldest easy story about the repo is:

- local token telemetry overlay for Codex sessions

That story was true at the beginning.

It is no longer enough.

### The Current Story

The repo is now three things at once:

1. local session telemetry and burn visibility
2. a backend-backed local automation and jobs surface
3. an emerging supervision product for task-oriented AI work

The third part is the one that needs the strongest product definition now.

If it is not defined, the repo will drift into a bundle of adjacent tools instead of a coherent product.

### The Better Story

The better story is:

- CodexDashboard is a Windows-first local cockpit for supervising Codex work across tasks, jobs, and intervention-driven improvement loops.

That story still needs one more simplification for human value:

- it is the place where the human can see what matters, dispatch what is ready, and recover what is drifting

### Why Tasks Are The Right Center

Tasks are the right center because they sit at the intersection of:

- human intent
- durable scope
- agent execution
- audits
- handoffs
- Dream-generated suggestions
- eventual dispatch state

Token burn matters, but burn is not the unit of meaning.

Jobs matter, but jobs are not the unit of work the human actually reasons about.

Tasks are the unit the human already uses to understand whether progress is real.

That is why the `Tasks` tab should become the product heart.

## Human Burden Taxonomy

The repo is easiest to design well when the human burden is named concretely.

The `Tasks` tab should not be designed only around features.

It should be designed around the burdens the human is paying today.

### Burden 1: Reconstructing State From Scattered Surfaces

Right now the human often has to combine:

- task markdown
- handoff notes
- local memory
- session history
- backend/job state

to answer one simple question:

- what is really happening with this work

That reconstruction tax is expensive because it is paid repeatedly.

The `Tasks` tab should directly reduce it by making one trustworthy summary surface exist.

### Burden 2: Carrying Silent Contracts In Memory

Humans currently have to remember things like:

- this task is probably waiting, not sleeping
- that task usually goes quiet for a while
- this other task tends to stall unless nudged

That memory should not live primarily in the human.

The product should carry as much of that contract as the system can honestly encode.

### Burden 3: Paying Recovery Cost Every Time Something Goes Quiet

If the human cannot tell what silence means, then every quiet task becomes risky.

The cost of recovery includes:

- finding the right artifact
- remembering what the task was doing
- finding the last relevant transcript
- deciding whether to nudge or interrupt

The `Tasks` tab should reduce that cost by keeping the state, the reason, and the launch path nearby.

### Burden 4: Confusing Suggestion With Commitment

Dream output is valuable, but it creates a new burden:

- the human now has more proposed work to evaluate

If the product does not separate candidate work from real work clearly, the human pays in confusion and mistrust.

The `Tasks` tab must therefore absorb candidate work without flattening the difference between:

- suggestion
- promotion
- dispatch

### Burden 5: Needing To Read Too Much Before Acting

Rich task markdown is good.

It is not the same as a humane default screen.

When the human has to read too much before they can:

- dispatch
- recover
- review
- promote

the product is offloading comprehension work that it should have absorbed.

### Burden 6: Fear Of Breaking The Flow

Humans avoid touching systems that feel brittle.

If interrupting, poking, or opening a task feels dangerous or unclear, the human will hesitate to use the product at the exact moment it is supposed to help.

The interface should reduce that fear by making actions:

- specific
- reversible when possible
- clearly justified by visible state

### Burden 7: Ambiguous Priority

A long task list is not the same as a priority surface.

When the human cannot tell what actually matters now, they either:

- over-read everything
- or act on whatever is freshest, loudest, or most familiar

Neither is good product behavior.

The `Needs Attention` region and grouped stream should reduce this burden explicitly.

### Burden 8: Weak Longitudinal Trust

The human's opinion of the product is not formed by a single success.

It is formed by a repeated feeling:

- when I come back later, does the dashboard help me understand reality quickly

That means the product must optimize for repeated re-entry, not only for first demo value.

## The Product Jobs The Tasks Tab Must Support

The `Tasks` tab exists to support a small number of recurring human jobs.

If it tries to support everything equally, it will become muddy.

### Job 1: Rapid Re-Entry

The human returns after time away and wants a fast answer:

- what changed
- what matters
- what should I do first

The screen should be optimized for this job above almost all others.

### Job 2: Safe Dispatch

The human sees a ready task and wants confidence that dispatch is warranted.

They should not need a full artifact dive to answer:

- what this task is
- why it is ready
- what will happen when I dispatch it

### Job 3: Recover A Quiet Task

The human notices or is told that a task went quiet.

The product should make the recovery path obvious:

- understand the silence
- choose a bounded action
- get to deeper context fast

### Job 4: Review Proposed Work

The human wants to review Dream candidates without confusing them for active commitments.

The product should help them answer:

- is this proposal meaningful
- where did it come from
- should it be promoted

### Job 5: Spot Systemic Drift

The human wants to notice when the whole system is developing a pattern:

- too many sleeping tasks
- too many waiting-on-human tasks
- too few truly ready tasks

The summary layer should help here without becoming a BI dashboard.

## Screen States In Detail

The `Tasks` surface should not be designed only for the happy populated state.

Its credibility depends on how it behaves when the world is incomplete or inconvenient.

### State: First Meaningful Empty

This is the state where the repo has no current tasks or no current data source has produced any yet.

The screen should not look broken.

It should explain:

- there are no tasks to show yet
- what kinds of work will appear here later
- what the first useful next step is

Possible copy:

- `No tracked tasks yet. Promoted Dream work and authored tasks will appear here once they exist.`

### State: Loading

Loading needs to feel active but not noisy.

The product should prefer skeleton structure that matches the real screen rather than a spinner on an empty void.

That teaches the screen composition even before the data arrives.

### State: Fresh Populated

This is the normal working state.

The product should feel calm, structured, and immediately usable.

The strongest cues should go to:

- attention items
- selected task meaning
- next safe action

### State: Slightly Stale

This is a subtle but important state.

The backend may be reachable soon, but the last known snapshot is not brand new.

The product should say this gently:

- `Showing the last trusted snapshot from 8 minutes ago.`

That is enough to preserve honesty without sounding alarming.

### State: Stale Enough To Distrust

Past a certain point, the product should become more explicit:

- `Live task status is stale. The last trusted snapshot is from 3:12 PM.`

The screen should remain usable for artifact review and orientation, but it should down-rank control confidence.

### State: Backend Unavailable

The dashboard should not collapse into emptiness.

It should show:

- live refresh failed
- last trusted snapshot time if available
- what the human can still inspect safely

Actions like `Dispatch` may need to be visibly disabled with reasons.

### State: Candidate-Only

There may be moments where no real tasks are active but Dream has surfaced candidates.

That should not feel like the system is empty.

It should feel like:

- there is no active queue, but there is suggested work worth review

### State: Everything Quiet But Healthy

This is a subtle design challenge.

Sometimes the best truthful answer is:

- nothing needs you right now

That state should feel reassuring, not barren.

The summary and recent completions can help make the system feel alive without inventing drama.

## Section-Level Interaction Rules

### Needs Attention Section

This section should always stay expanded by default when it has content.

The product should not hide urgent work behind a collapsed group.

### Running Section

This section can be more compact because the main question is:

- is progress happening

The row reason line matters a lot here.

`Running` without a clear recent-progress story can quickly turn into false comfort.

### Ready To Dispatch Section

This section should feel actionable and slightly optimistic.

It should invite useful action without pressuring the human into clicking blindly.

### Dream Candidates Section

This section can be visually a little softer, because it is suggestion space, not active burden space.

But it should never look like dead or irrelevant content.

## Selection Behavior

The selected task should be obvious.

The selection should not rely only on a faint border.

Good selection behavior:

- clear background shift
- strong detail pane synchronization
- stable scroll position when moving between rows

The product should preserve orientation.

If selecting a row causes the whole list to jump unpredictably, the screen will feel brittle.

## Sorting And Grouping Rules

Grouping should come first.

Sorting should happen within groups.

A likely first rule set:

- `Needs Attention`: sort by severity then recency
- `Running`: sort by freshest recent progress
- `Ready`: sort by strategic importance or recent qualification
- `Dream Candidates`: sort by recency and review value
- `Done`: sort by completion recency

This is better than one universal sort because different groups answer different human questions.

## Provenance Must Stay Visible

Provenance matters because the same-looking task titles can mean very different things.

The product should show provenance labels like:

- `Authored`
- `Promoted`
- `Candidate`
- `Follow-on`

These do not need to dominate the row, but they should be discoverable at a glance.

That small amount of visible provenance prevents a lot of misreading.

## Why A Rich Task Surface Helps The Whole Repo

This is not just a UI convenience task.

It helps the whole repo because it puts pressure on the underlying system to become clearer.

Once the product has to show:

- sleeping
- waiting
- provenance
- next expected step
- last meaningful progress

the backend and task-writing standards are forced to get more honest too.

That is good.

It means the UI is not only a consumer of truth.

It is also a forcing function for cleaner truth.

## Additional Copy Examples

Useful microcopy examples:

- `Ready to start. No active blocker is recorded.`
- `Waiting on your answer about the dispatch contract.`
- `Promoted from Dream on Apr 23. Not dispatched yet.`
- `Blocked by missing backend support for durable run state.`
- `Sleeping: no recent progress and no valid wait reason.`
- `Open the working thread to continue from the latest context.`

Bad microcopy examples:

- `Operationally unavailable`
- `Inactive`
- `Out of sync`
- `Paused`

Those are too abstract unless the product is very explicit about what they mean.

## The Wrong Product Shapes To Avoid

### Wrong Shape 1: The Database Browser

This is the surface that shows:

- task id
- status
- phase
- updated_at
- some filters

and calls it a task dashboard.

That shape fails because it makes the human do the translation.

The human still has to answer:

- does this need me
- is this safe to dispatch
- why is it blocked
- can I trust silence here
- where do I go deeper

### Wrong Shape 2: The Mini IDE

This is the surface that tries to prove power by showing too much:

- logs
- raw markdown
- raw JSON
- console output
- diff views
- multiple cramped panes

That shape fails because it front-loads complexity instead of revealing it progressively.

It may help an expert in a pinch.

It is not the right default surface for a human who wants one quick trustworthy answer.

### Wrong Shape 3: The Control Plane Cosplay

This is the surface that talks like infrastructure:

- runtime
- orchestration
- workflow execution
- enqueue state
- desired versus observed
- trigger kind

Some of those concepts are real and necessary underneath.

They are not the first language the human should have to read on the default screen.

### Wrong Shape 4: The Dream Inbox

This is the surface that over-rotates toward suggestions:

- lots of option tasks
- lots of proposals
- lots of generated work
- weak separation between proposals and approved tasks

That shape fails because it makes the system feel noisy and overeager.

Dream output matters.

But the dashboard heart must stay grounded in real work and real commitments first.

### Wrong Shape 5: The Comforting Lie

This is the surface that makes everything look fine:

- green status pills everywhere
- vague labels like `active` or `ready`
- no clear exposure of sleeping or silent work
- no visible human waiting state

That shape fails because it teaches the human not to trust the product.

Once trust is lost, even good automation feels risky.

## The Right Product Shape

### One-Sentence Product Shape

The right shape is:

- a calm high-level task cockpit that tells the truth, makes the next action obvious, and lets the human go deeper only when needed

### What The Default Surface Must Answer

The default `Tasks` screen should let the human answer, within a few seconds:

- what needs me right now
- what is actively moving
- what is sleeping
- what is blocked
- what is ready to dispatch
- what Dream proposed recently

If the screen cannot answer those six questions quickly, it is not yet the product heart.

### What The Default Surface Should Not Require

The default screen should not require:

- reading raw task markdown
- remembering yesterday's context
- knowing backend implementation terms
- checking multiple tabs to reconstruct one task story
- guessing whether silence is healthy

## Primary Human Promise

The primary human promise of the future product should be:

- `You can leave work running, come back later, and understand the real state quickly enough to act with confidence.`

That promise matters more than whether the screen is dense or sparse.

It matters more than whether the backend is elegant.

It matters more than whether the product can technically enumerate every possible task field.

It is the promise the human will judge the product by.

## The Primary Human Goal

The primary human goal for the `Tasks` tab is:

- review and steer work without becoming the system's missing memory

That is the actual human relief the repo should provide.

## The Human Persona This Surface Is For

This product is not for a random consumer.

It is also not for a trained NOC operator.

It is for a very specific person:

- a technically capable human who does not want to burn attention on unnecessary orchestration trivia

This person:

- is comfortable making product and engineering judgments
- can read task artifacts when needed
- does not want the default screen to behave like a low-level system inspector
- values truth and specificity
- gets annoyed when the system hides real failure
- also gets annoyed when the system makes them babysit preventable details

That means the humane bar is not `make it childish`.

The humane bar is:

- make it direct
- make it truthful
- make it easy to scan
- make it fast to recover context

## First-Run And Recurring Use

### First-Run Expectation

On first meaningful use of the `Tasks` tab, the human should immediately understand:

- this is where work lives
- these are real tasks
- these are proposed tasks
- these tasks are safe to start
- these tasks need attention
- this is where to click to go deeper

The human should not need onboarding copy that explains architecture.

The screen itself should teach the mental model.

### Recurring-Use Expectation

On recurring use, the human should be able to:

1. open the tab
2. scan the top summary and attention list
3. decide whether to:
   - recover something
   - dispatch something
   - review a candidate
   - open a task
4. leave again

That is the bar.

Not:

- study a dashboard
- perform a system audit
- read a runbook

## Why The Tasks Tab Should Be The Beating Heart

### Because Meaning Lives Here

The heart of this product should not be:

- token output
- scheduler internals
- a raw daily report

The heart should be the place where the human can see meaningful work moving.

That is tasks.

### Because It Connects The Other Surfaces

`Usage` tells the human that something is expensive.

`Jobs` tells the human that recurring automation exists and whether it is alive.

`Tasks` tells the human whether that activity is worth anything, whether it is drifting, and what to do next.

Without `Tasks`, the other surfaces remain partial truths.

### Because It Is The Bridge Between Manual And Autonomous

Today the system still depends on human direction often.

Tomorrow it should depend on it less.

The place where that transition becomes visible is the `Tasks` surface:

- authored work
- generated candidates
- enqueued work
- running work
- waiting work
- completed work

That is where the human can feel whether the system is getting more trustworthy.

## The Proposed Mental Model

The future `Tasks` tab should teach a very small mental model:

1. work starts as a task or candidate
2. tasks become ready, blocked, or waiting
3. ready tasks can be dispatched
4. dispatched tasks become runs with durable state
5. runs can need attention, stall, or finish
6. Dream can suggest more work, but suggestion is not commitment

That is enough.

Anything that expands this model should appear only when the human chooses to go deeper.

## Information Architecture

### Top-Level Regions

The first-release `Tasks` surface should have five major regions:

1. header and high-level status
2. `Needs Attention` summary strip
3. task list or grouped task stream
4. detail pane for the selected task
5. bounded action area

That is the core anatomy.

### Header

The header should answer:

- what this tab is
- whether the data is fresh enough to trust
- whether the system is reachable

It should not become a toolbar graveyard.

Recommended visible content:

- `Tasks`
- a one-line status summary
- `Last updated`
- one primary refresh action

### Needs Attention Strip

This should be the first content block because it protects human attention.

It should summarize counts and make the urgent state visible.

Recommended first categories:

- sleeping
- waiting on you
- blocked
- failed
- candidate tasks awaiting review

These should not all be equal visually.

Sleeping and waiting-on-you should probably be strongest because they are the easiest to lose.

### Task Stream

The main list should be grouped by meaning, not only sorted by timestamp.

The first grouping model should likely be:

- Needs attention
- Running
- Ready to dispatch
- Waiting or blocked
- Dream candidates
- Recently completed

This group structure matches human questions better than:

- all tasks in one flat list

### Detail Pane

The detail pane should let the human understand one selected task without opening markdown immediately.

It should explain:

- what the task is
- why it exists
- current state
- next expected step
- what changed recently
- what artifacts matter
- what actions are safe now

The detail pane is where the system earns trust.

If it only repeats labels from the list, it is wasted space.

### Action Area

Actions should be explicit and bounded.

Good examples:

- `Dispatch`
- `Open Task`
- `Open Thread`
- `Poke`
- `Interrupt`
- `Review Candidate`
- `Promote`

Bad examples:

- `Manage`
- `Control`
- `Run Action`
- `Continue`

The action names should tell the human what the button will do, not make them guess.

## Row Anatomy

### What A Task Row Must Communicate

Each task row should let the human answer:

- what this work is
- why it matters
- what state it is in
- whether it needs me
- whether I can act on it now

### Recommended Row Fields

Each row should include:

- task title
- short plain-language burden summary
- state pill
- one provenance label
  - authored
  - promoted from Dream
  - candidate
- one freshness signal
  - last progress
  - last update
- one next-step or blocker summary

Optional but likely useful:

- repo area or subsystem tag
- owning thread/run indicator

### What Should Not Be In The Default Row

The default row should not show:

- raw task ids as the most prominent text
- long path strings
- backend workflow ids
- raw timestamps only without human interpretation
- more than one secondary sentence of detail

### Row Copy Rule

The secondary text in a row should answer:

- why should I care about this row right now

Examples:

- `Waiting on human approval for the plan.`
- `No progress since 2:14 PM and no wait contract recorded.`
- `Ready to dispatch after Dream promotion.`
- `Blocked by missing backend execution-state contract.`

Not:

- `phase: planning`
- `status: in_progress`
- `gate: audit`

Those fields may exist underneath.

They are not the right default language.

## Status System

### Human-Visible Status Vocabulary

The first visible vocabulary should be small and blunt:

- `Needs you`
- `Running`
- `Sleeping`
- `Blocked`
- `Ready`
- `Candidate`
- `Done`
- `Failed`

This is not the full backend vocabulary.

It is the humane summary vocabulary.

### Why This Small Vocabulary Works

Each label answers a human question:

- `Needs you`: do I have to act
- `Running`: is it alive and progressing
- `Sleeping`: did it likely drop the ball
- `Blocked`: is progress impossible right now
- `Ready`: can I start it safely
- `Candidate`: is this suggestion not commitment
- `Done`: can I stop thinking about it
- `Failed`: did it end badly

### Visible State Details

The label alone is not enough.

Every state should pair with a human-readable reason line.

Examples:

- `Needs you` + `Waiting on approval of the implementation plan.`
- `Sleeping` + `Last progress was 4 hours ago and no wait reason was recorded.`
- `Blocked` + `Dispatch runtime contract is not implemented yet.`
- `Candidate` + `Dream proposes this as follow-on work from April 19 interventions.`

### Color Semantics

The color system should support meaning without carrying it alone.

Proposed direction:

- amber or warm gold for `Needs you`
- cool cyan for `Running`
- muted coral or red-orange for `Sleeping` and `Failed`
- muted slate or steel for `Blocked`
- olive or calm green for `Ready`
- violet is explicitly not required and should not become the lazy default accent

The product should feel operational, not generic SaaS.

## The Detail Pane In Depth

### Purpose

The detail pane is not just a bigger row.

It is the place where the human decides whether to trust the state and what to do next.

### Proposed Sections

The first detail pane should include:

- `Summary`
- `Why this task exists`
- `Current state`
- `What changed recently`
- `Next expected step`
- `Artifacts`
- `Actions`

### Summary

The summary should restate the task in plain language.

It should not just mirror the title.

It should answer:

- what job this task is trying to do for the human

### Why This Task Exists

This section matters because tasks otherwise lose meaning and become queue entries.

It should summarize the burden being reduced.

The text can come from `TASK.md` or a derived preview, but the product should present it as human meaning, not as markdown sections.

### Current State

This section should say:

- where the work actually is
- what the durable run or task state says
- whether the state is trusted
- what makes it trusted or untrusted

### What Changed Recently

This section is the human memory saver.

It should protect the human from having to remember:

- the last pass
- the last transition
- the last intervention
- the last meaningful run activity

It should likely show a short event list rather than raw logs.

### Next Expected Step

This is the section that makes the next action obvious.

It should answer:

- what the system expects to happen next
- who owns that next step
- when silence becomes suspicious

That last part is critical.

The detail pane should explicitly say when a task can be quiet and when it should not be quiet.

### Artifacts

This section should make the deeper documents accessible without forcing them into the default visible surface.

The likely artifact links are:

- `TASK.md`
- `PLAN.md`
- `HANDOFF.md`
- relevant testing or research artifacts
- transcript or thread context if a run exists

The human should be able to open these quickly, but should not have to open them for every task.

### Actions

Actions should live next to the context needed to make them feel safe.

The action section should make clear:

- what the action does
- whether it is safe now
- why it may be unavailable

## Deep-Context Launch

### Why It Matters

This is one of the most important product outcomes in the whole design.

The human will sometimes need to go deeper.

When that happens, the worst possible product behavior is:

- make them manually search for the right transcript, task folder, or thread

That converts a moment of necessary deeper work into friction and distrust.

### What The Product Should Do

The product should offer a direct `Open Thread` or `Open Working Context` action that launches the best available deep surface.

Possible targets:

- the task folder in VSCodium
- the relevant transcript
- a session-specific workspace
- the strongest local deep-link available at implementation time

### What The Product Should Say

The label should name the user goal:

- `Open Thread`
- `Open Task Files`
- `Open Working Context`

Not:

- `Launch`
- `Attach`
- `Open Session Artifact`

### Why This Connects To Human Need

The human does not need one more button.

The human needs the cost of recovery to be low enough that dispatch feels safe.

Deep-context launch lowers that recovery cost.

## Dispatch From The Tasks Surface

### What Dispatch Means To A Human

To a human, `Dispatch` should mean:

- start this work under the agreed contract and show me where it goes

It should not mean:

- call some backend endpoint and hope for the best

### Preconditions Must Be Visible

Before the human dispatches, the surface should tell them:

- whether the task is ready
- why it is ready
- what contract or plan it is using
- what will happen next

If any of that is unknown, the task is not actually ready.

### After Dispatch

After dispatch, the row and detail pane should change immediately enough that the action feels causally connected.

The human should see:

- that a run exists
- what its starting state is
- what the next expected signal will be

That protects trust.

## Sleeping And Stall Recovery

### Why This Is Central

A large part of the repo's long-term value is reducing supervision burden.

That value does not come mainly from:

- launching work

It comes from:

- knowing when launched work quietly stopped doing what it should

### How The UI Should Express Sleeping

`Sleeping` should be treated as a first-class visible state, not a hidden backend concept.

The UI should communicate:

- when the last meaningful progress happened
- why silence is suspicious
- what the human can do next

Good copy:

- `No progress since 1:42 PM. This task did not record a valid wait reason.`

Bad copy:

- `Run stale`
- `Heartbeat missing`
- `Workflow timeout`

### Recovery Actions

The likely visible actions are:

- `Poke`
- `Open Thread`
- `Interrupt`

These actions must feel understandable and bounded.

`Poke` should not sound magical.

It should sound like:

- ask the task to continue or justify its silence

`Interrupt` should not feel like a dangerous admin move.

It should feel like:

- stop this run and return control cleanly

## Waiting On Human

### This State Must Be Precious

`Waiting on human` is not just another status.

It is one of the main reasons the human opens the tab.

The product should make this state unmistakable.

### What It Should Show

It should show:

- what the task is waiting for
- why the task cannot proceed without it
- what artifact or decision the human should inspect

The action should be direct:

- `Open plan`
- `Review candidate`
- `Approve`
- `Open task`

### Why This Matters

When `waiting on human` is vague, the human becomes a detective.

When it is explicit, the product feels collaborative instead of needy.

## Dream Candidates

### Why Candidates Belong On The Tasks Surface

Dream-generated option tasks are relevant work suggestions.

They belong on the `Tasks` surface because the human should not have to visit a separate hidden universe to discover new work.

### Why They Must Be Distinct

They must also remain visually distinct from real tasks because:

- they are suggestions
- they are not yet commitments
- they may not be safe to dispatch

### Proposed Candidate Treatment

Candidate work should appear in a dedicated group:

- `Dream Candidates`

Each candidate row should show:

- title
- short burden summary
- provenance
- recency
- one clear action
  - `Review`
  - or `Promote`

The row should not show `Dispatch`.

That would confuse suggestion with commitment.

### Candidate Detail Pane

The detail pane for a candidate should make it easy to understand:

- what burden it is reducing
- which Dream packet produced it
- why it was proposed
- what promotion will do

Promotion should feel like:

- turning this into a real task

Not:

- starting work immediately

## The Human Need Around Dream

The real need is not:

- receive more generated tasks

The real need is:

- review proposals with enough context to decide whether they deserve durable status and future execution

That is why promotion must be visible and conservative.

## Summary Metrics That Actually Matter

### The Top Summary Should Exist

The top summary matters because it lets the human orient fast.

### The Right Metrics

Useful metrics:

- tasks needing attention now
- tasks waiting on human
- tasks currently running
- sleeping tasks
- ready-to-dispatch tasks
- Dream candidates awaiting review

Potentially useful later:

- completed today
- interrupted today
- average time since last meaningful progress for active work

### The Wrong Metrics

Bad top-summary metrics:

- total tasks in repo without context
- raw pass counts
- raw transcript counts
- average task markdown size
- total backend workflows

Those can be interesting somewhere else.

They are not the human job.

## What Should Be Automatic, Visible, Or Hidden

### Should Be Automatic

- freshness checks
- stale or sleeping detection once the durable contract exists
- candidate provenance capture
- promoted-task provenance wiring
- sensible task grouping and default sorting

### Should Be Visible

- task meaning
- status and reason
- last meaningful progress
- next expected step
- whether the human is needed
- candidate versus real task distinction

### Should Be Hidden By Default

- raw run ids
- backend workflow ids
- detailed runtime taxonomy
- raw task-state JSON
- full logs

Those can live behind a `Details` or deeper context path when they are useful.

## Trust Model

### What Builds Trust Here

Trust comes from four things:

1. the screen tells the truth about uncertainty
2. the screen notices when silence is suspicious
3. actions feel causally connected to visible state changes
4. deeper context is easy to open when needed

### What Breaks Trust

- vague green statuses
- silent failures
- unclear difference between candidate and approved work
- actions that do something without the surface reflecting it quickly
- stale data with no visible freshness warning

### Data Freshness

Freshness needs visible treatment.

The product should say when the data is:

- fresh
- slightly stale
- stale enough to distrust
- unavailable because the backend is down

This should be calm and clear, not panicky.

## Humane Error And Recovery Model

### The Product Should Prefer Prevention

The best task-related error message is often a prevented confusion.

Examples:

- do not show `Dispatch` when a task is not ready
- do not show `Poke` when a valid wait contract exists
- do not show `Promote` on a task that is already real

### When Error Is Still Needed

Error copy should answer:

- what failed
- what it means for the human
- what to do next

Good examples:

- `Tasks could not refresh from the backend. The last trusted snapshot is from 4:12 PM.`
- `This task cannot be dispatched yet because the plan is still awaiting approval.`
- `This run cannot be poked because it is explicitly waiting on human approval.`

### Recovery Paths

Every major failure mode should have a calm recovery path:

- backend unavailable
  - retry
  - show last trusted snapshot
- task sleeping
  - poke
  - open thread
  - interrupt
- task blocked
  - open blocking artifact
- candidate uncertain
  - review source packet

## Interaction Flows

### Morning Review Flow

This is likely the most important recurring flow.

The human opens the app or tab and wants to know:

- what happened since last check
- what needs them now
- what they should do first

Ideal flow:

1. Open `Tasks`.
2. See top summary with urgent counts.
3. Scan `Needs Attention`.
4. Select the top sleeping or waiting task.
5. Read one clear reason line in the detail pane.
6. Click the obvious next action.

This flow should take seconds, not minutes.

### Dispatch Flow

Ideal flow:

1. Open `Ready` group.
2. Select a task.
3. Read what it does and why it is ready.
4. Click `Dispatch`.
5. Watch the state change to `Running`.
6. Leave confidently.

### Recovery Flow

Ideal flow:

1. Open `Tasks`.
2. See one task marked `Sleeping`.
3. Select it and read why silence is suspicious.
4. Choose `Poke` or `Open Thread`.
5. Resume work without hunting through session history.

### Candidate Review Flow

Ideal flow:

1. Open `Dream Candidates`.
2. Select a candidate.
3. Read burden summary and provenance.
4. Decide whether to ignore, review source, or promote.
5. Promote if warranted.
6. See it become a real task, not a still-ambiguous suggestion.

## Scenario Walkthroughs

### Scenario 1: The Human Returns After Half A Day

The human was away for hours.

Several things may have happened:

- a run completed
- a run stalled
- a Dream packet generated candidates
- a task started waiting on approval

The product must compress that reality into a trustworthy first glance.

The first glance should not require comparing timestamps mentally.

It should say something like:

- `2 tasks need you, 1 task is sleeping, 3 tasks are running, 2 Dream candidates await review.`

That one sentence creates orientation.

Then the grouped stream turns orientation into action.

### Scenario 2: The Human Is Skeptical About Silence

The human sees a task row with no updates for a while.

The worst product response is ambiguity.

The right response is explicitness:

- `Waiting on approval since 2:05 PM.`
- or
- `No progress since 2:05 PM and no wait reason recorded.`

Those are radically different meanings.

The product must not blur them together.

### Scenario 3: A Candidate Looks Valuable

Dream surfaced a candidate that seems useful.

The human wants enough context to judge it without opening five docs.

The candidate detail pane should give:

- a one-paragraph summary
- the burden being reduced
- where it came from
- why it was proposed
- the action to promote

The human can still open deeper artifacts.

But the product has already done the first layer of comprehension work.

### Scenario 4: A Task Is Blocked On Another Task

This is a common future shape.

The product should express dependency honestly:

- `Blocked by Task-0008 dispatch contract work.`

Then it should make the blocker reachable.

If the human cannot click through to the blocker, the dashboard forces more mental bookkeeping than it should.

### Scenario 5: The Backend Is Down

The product should not go blank and pretend nothing exists.

It should show:

- that live refresh failed
- the age of the last trusted snapshot
- what the human can still do safely

This is a trust issue.

If the product acts like no tasks exist when the backend is merely unreachable, the human learns the wrong lesson.

### Scenario 6: The Human Wants To Stop Something

Interrupt is not an edge case.

It is part of supervising autonomous work.

The product should let the human stop a run without fear that doing so will lose all context.

That means:

- clear state
- clear reason
- durable record of interruption
- easy path back to task context

### Scenario 7: A Task Has Rich Markdown But The UI Must Stay Simple

This is where progressive disclosure matters.

The UI should summarize first and link deeper second.

It should not collapse into a markdown reader because the tasks are rich.

The richness of the docs is a strength.

The job of the UI is to make the first layer of that richness glanceable.

### Scenario 8: The Human Wants One Quick Decision

Sometimes the human only wants to know:

- should I dispatch something right now

The `Ready` group should make that easy.

The detail pane should tell them:

- what this task is
- what it will affect
- whether there is any known blocker
- why it is considered ready

That is enough for a quick decision.

## Copy System

### Tone

The tone should be:

- direct
- calm
- specific
- non-theatrical

The dashboard should not sound apologetic, chatty, or mystical.

### Good Nouns

Prefer:

- task
- candidate
- run
- waiting on you
- sleeping
- blocked
- ready
- open thread

Avoid as the default surface language:

- workflow
- activity
- trigger kind
- reconcile
- desired state
- execution artifact

These may still exist in deeper views or backend docs.

### Good Secondary Sentences

Examples of secondary sentences that do real work:

- `Ready to dispatch after plan approval and no active blockers.`
- `Waiting on your review of the Dream-generated proposal.`
- `Blocked until the durable execution-state contract exists.`
- `Sleeping: no progress for 3 hours and no wait reason recorded.`

### Bad Secondary Sentences

- `Status synchronized.`
- `Task is in progress.`
- `Awaiting next action.`
- `Workflow in pending state.`

Those say very little.

## Interface Semantics

### Icons

Icons should reinforce meaning but not carry it alone.

Likely useful icon families:

- play or arrow-forward for dispatch
- pulse or activity for running
- pause-hand or hourglass-human hybrid metaphor for waiting on human
- moon, dim pulse, or sleep icon for sleeping
- stop-square for interrupt
- thread/link/file icon for open thread or open task
- wand or spark only carefully for Dream candidates

Avoid icon choices that feel decorative or generic.

### Typography

The hierarchy should feel operational and intentional.

Recommendations:

- a strong, slightly condensed heading face or operational sans
- a separate monospace only for ids, timestamps, and machine-like fields
- clear type-role differences between:
  - screen title
  - group headers
  - task titles
  - secondary burden summaries
  - status pills
  - metadata

The product should not feel like default Tk widgets painted dark.

### Density

The product can be dense, but not cramped.

Dense is good when:

- the scan path is obvious
- the most important text stays readable
- the row structure is consistent

Cramped is bad when:

- every row competes equally
- text wraps unpredictably
- controls feel too close to status text

## The Role Of The Detail Pane Versus Markdown

The detail pane should absorb the first sixty seconds of comprehension.

Markdown should absorb the rest.

That is the right split.

If the human has to open markdown for the first sixty seconds of comprehension, the UI is too thin.

If the UI tries to replace markdown completely, the UI becomes bloated.

## The Relationship To Task-0008

The `Tasks` tab should not fake runtime sophistication before [Task-0008](../../Task-0008/TASK.md) exists.

This means the product should be honest about staged capability.

Possible progression:

1. read-only surface with task grouping and provenance
2. real dispatch once the backend contract exists
3. real sleeping detection once durable execution-state exists
4. real poke and interrupt once supervision exists

That staged rollout is healthy as long as each phase is explicit.

The UI should not pretend to know more than the backend can currently prove.

## The Relationship To Task-0010

Dream output is extremely relevant.

But candidate work should arrive on the `Tasks` surface with clear labels and provenance once [Task-0010](../../Task-0010/TASK.md) exists.

The surface should never imply:

- a candidate is already approved
- a promoted task is already dispatched

Three states must stay visually distinct:

- candidate
- promoted real task
- active dispatched run

## The Relationship To Task-0005

[Task-0005](../../Task-0005/TASK.md) already taught one important lesson:

- the dashboard should be a client of durable backend truth, not a local scheduler pretending to be the system

That lesson applies here too.

The `Tasks` tab should consume durable state.

It should not become the only place where task-run truth exists.

## Product Boundaries

### What Belongs In The First Tasks Tab

- grouped task review
- attention-oriented summary
- task meaning
- task state
- deep-context launch
- bounded actions

### What Does Not Need To Be In The First Tasks Tab

- full timeline diffing
- arbitrary transcript browsing
- custom query building
- bulk admin operations
- every task field ever stored

Those may become useful later.

They are not what makes the first version valuable.

## Implementation Taste

The visual direction should feel like:

- a serious private cockpit
- warm enough to feel human
- sharp enough to feel operational

It should not feel like:

- a generic web admin
- a spreadsheet
- a toy kanban board

Good reference qualities:

- high contrast
- purposeful typography
- grouped blocks with clear hierarchy
- visible breathing room around urgent content
- subtle texture or atmosphere instead of flat sameness

## Anti-Patterns

### Anti-Pattern: Everything Is A Pill

Do not reduce the whole surface to:

- title
- status pill
- small metadata

That shape is too thin for meaningful supervision.

### Anti-Pattern: Hidden Reasons

Do not make the reason for a state invisible until hover or secondary click.

The reason line is part of the default truth.

### Anti-Pattern: One Giant Activity Feed

Feeds are tempting.

They are bad at answering the human's core questions unless heavily structured.

The default should be grouped tasks, not an infinite event feed.

### Anti-Pattern: A Sortable Table As The Main Surface

Tables are useful for operators.

They are weak at expressing meaning, urgency, and action in a humane way.

The product may use some table logic underneath.

It should not feel like a table-first product.

## Concrete Screen Proposal

### Overall Layout

The strongest first-direction layout is probably:

- top header row
- attention summary cards directly below
- main split body
  - left: grouped task stream
  - right: persistent detail pane

This gives the human:

- fast scan
- stable focus
- no context loss when selecting tasks

### Why Split Body Wins

A split layout likely wins over stacked full-page drilldown because:

- the human can keep the wider queue in view
- selecting a task does not feel like leaving the screen
- comparison between rows stays easy

### Mobile Is Not The Priority

This repo is Windows-first.

The desktop surface can therefore use a denser and wider split confidently.

But it should still behave reasonably when the window narrows.

## Example Default Screen Story

Imagine the human opens the app and sees:

- `Tasks`
- `Updated 20s ago`
- summary cards:
  - `Needs you: 2`
  - `Sleeping: 1`
  - `Running: 3`
  - `Ready: 4`
  - `Candidates: 2`
- a grouped stream:
  - `Needs Attention`
  - `Running`
  - `Ready to Dispatch`
  - `Dream Candidates`
- one selected task detail pane:
  - `Summary`
  - `Why this task exists`
  - `Current state`
  - `What changed recently`
  - `Next expected step`
  - `Artifacts`
  - `Actions`

From that one screen, the human can decide:

- rescue the sleeping task
- approve something waiting
- dispatch something ready
- review a candidate

That is what product coherence feels like.

## Example Good Row Set

### Needs Attention Row

Title:

- `Build task dispatch layer and durable execution-state contract`

Secondary line:

- `Blocked until the execution-state vocabulary is locked for the backend.`

Status:

- `Needs you`

Meta:

- `Task-0008`
- `Last changed 45m ago`

### Running Row

Title:

- `Regenerate April 19 intervention packet`

Secondary line:

- `Running with fresh progress 3 minutes ago.`

Status:

- `Running`

### Candidate Row

Title:

- `Approval packet headers as a shared contract`

Secondary line:

- `Dream candidate from April 19 intervention analysis.`

Status:

- `Candidate`

Action:

- `Review`

## Example Bad Screen Story

Imagine the human opens the app and sees:

- a toolbar with ten buttons
- a table of task ids
- columns for phase, gate, audit, and updated_at
- generic green and gray pills
- no explanation of sleeping versus waiting
- no clear distinction between candidates and tasks

That screen may be technically complete.

It fails the human need completely.

The human still has to reconstruct the story from system terms.

## Design Review Checklist

Before calling any `Tasks` surface good enough, the team should ask:

- Can a human tell what needs attention in five seconds?
- Can a human distinguish candidate work from real work immediately?
- Can a human tell why a task is sleeping, blocked, or waiting?
- Can a human open deeper context quickly?
- Can a human decide whether dispatch is safe from the detail pane alone?
- Does the screen still make sense if the backend is stale?
- Does the screen feel like a task cockpit rather than an orchestration console?

If any of those answers is no, the surface is not ready.

## Open Tensions

### Tension 1: Richness Versus Calm

The tasks are rich.

The screen should stay calm.

The answer is not to hide truth.

The answer is to put the first truth in the default layer and the rest one click deeper.

### Tension 2: Early Read-Only Value Versus Full Dispatch Ambition

The UI can deliver real value before the full dispatch layer exists.

That is good.

But the read-only phase should be framed as a real phase, not as the whole product.

### Tension 3: Candidate Visibility Versus Noise

Dream candidates are valuable.

Too much candidate emphasis will make the dashboard feel noisy and speculative.

The solution is grouping, provenance, and clear promotion boundaries.

### Tension 4: Operational Seriousness Versus Friendliness

The app should feel serious.

It should not feel harsh.

The balance is:

- strong structure
- precise words
- calm explanations

not:

- cute language
- neon chaos
- fake friendliness

## Working Conclusion

The repo's emerging human need is not merely:

- `show me what Codex is doing`

It is:

- `let me trust, direct, and recover real AI work from one local surface without carrying the system in my head`

The `Tasks` tab is where that need becomes visible and useful.

If this tab lands well, the repo stops feeling like several adjacent tools and starts feeling like one coherent supervision product.

If this tab lands badly, the repo will still have useful internals, but the human will continue paying unnecessary cognitive rent.

That is why this task matters.

That is why the `Tasks` tab should become the beating heart.

## Immediate Design Direction

For implementation planning, the current best direction is:

- a split-pane desktop `Tasks` surface
- grouped by human meaning, not raw task phase
- attention summary first
- task meaning and reason lines visible in the list
- deep-context launch always close at hand
- candidate work clearly separated
- sleeping and waiting-on-human treated as first-class states
- detail pane that explains current truth and next action without forcing markdown-first reading

That is the direction the `Stitch` prompt should visualize and the implementation should pressure-test next.
