# Task Create

This file defines the shared workflow for creating or rewriting `Tracking/Task-<id>/TASK.md`.

Use it together with:

- `C:\Users\gregs\.codex\AGENTS.md`
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\FILE-NAMING.md`
- `C:\Users\gregs\.codex\Orchestration\Exemplars\TASK.md`

## Purpose

The point of task creation is not to name a problem area.

The point is to produce a task writeup that lets a later human or agent answer three questions honestly:

1. What exact outcome is wanted?
2. What exact change is proposed?
3. What would prove the task is done?

If a task still allows multiple materially different implementations, including one the human would reject, the task is not written tightly enough yet.

## Task Writeup Types

Not every task needs the same level of specificity.

Choose one of these writeup types before writing `TASK.md`.

### 1. Concrete Implementation Task

Use this when the human is deciding whether to enqueue real work now.

This is the default for solution proposals, follow-up fixes, concrete workflow improvements, and other tasks that are meant to become executable backlog items.

Required bar:

- the main solution shape is already chosen
- the task names the concrete files, artifacts, fields, sections, checks, or behaviors that will change
- a later agent should not need to invent the main approach during implementation

### 2. Consensus Task

Use this when the real blocker is an unresolved product, policy, or workflow choice.

Required bar:

- the task names the exact decision that must be made
- the task names the options or tensions that must be resolved
- the output is a decision artifact, not disguised implementation

Do not write a fake implementation task when the real work is still choosing between solution shapes.

### 3. Research Task

Use this when key facts are still missing and those facts could change the design.

Required bar:

- the task names the specific questions that must be answered
- the task names the sources or code areas that must be checked
- the output is a bounded research artifact that should collapse uncertainty

Do not hide open research inside mushy implementation wording like `investigate`, `improve`, or `handle`.

## Base `TASK.md` Sections

Every new or rewritten `TASK.md` must include:

- `## Title`
- `## Summary` or `## Context`
- `## Goals`
- `## Non-Goals`
- `## Implementation Home`
- `## Acceptance Criteria`

Optional sections are allowed when they make the task harder to misread, such as:

- `## Constraints And Baseline`
- `## Proposed Changes`
- `## Expected Resolution`
- `## What Does Not Count`
- `## Open Questions`
- `## Proof Plan`
- `## References`

## Extra Requirements By Writeup Type

### Concrete Implementation Task

In addition to the base sections, concrete implementation tasks should usually include:

- `## Proposed Changes`
- `## Expected Resolution` when the task touches a human-facing surface or workflow
- `## What Does Not Count` when there is a real risk of proxy closure or degraded fallback closure
- `## Proof Plan` when the proof path is easy to fake or misread

Concrete implementation tasks must name:

- the chosen solution shape
- the concrete files, docs, prompts, schemas, generators, or artifacts that will be added or edited
- any exact fields, sections, reply modes, manifests, or gates that must exist afterward
- the exact boundary where the new behavior will be enforced

Concrete implementation tasks must not:

- leave two or three equally live solution shapes inside one task
- say only `add support`, `improve`, `handle`, `address`, or `tighten` without naming the resulting mechanism
- push the main design choice into implementation time

### Consensus Task

Consensus tasks should include:

- `## Decision To Make`
- `## Inputs`
- `## Options To Compare`
- `## Decision Output`

Consensus tasks must name what artifact will hold the decision once made.

### Research Task

Research tasks should include:

- `## Questions`
- `## Sources To Check`
- `## Output Artifacts`
- `## Exit Bar`

Research tasks must say how the research result will change the next lifecycle step.

## Concrete Implementation Workflow

When writing a concrete implementation task, use this order.

### Step 1. Ground From Durable Evidence

Read the repo docs, task artifacts, bugs, regression runs, or intervention packets that justify the task.

Prefer durable local artifacts over chat memory.

### Step 2. Name The Real Human Or Operator Outcome

Say what a person should be able to do, understand, or avoid after the task is done.

Do this before describing the mechanism.

### Step 3. State The Current Truth

Say what happens now and why that is not good enough.

If the current state is only a proxy, fallback, or partial proof, say that plainly.

### Step 4. Choose One Solution Shape

Pick the solution you are actually proposing.

Do not write a task that is a quantum superposition of several possible designs.

If the main choice is still open, stop and write a consensus or research task instead.

### Step 5. Name The Concrete Changes

List the concrete changes that will happen if the task is executed.

For example:

- exact shared docs to edit
- exact prompt files to edit
- exact schemas to extend
- exact task-owned artifact names to add
- exact fields or sections to add
- exact gate or validator behavior to enforce

This is the section that lets a human decide whether to enqueue the task.

### Step 6. Separate Goals From Non-Goals

Write what the task will do and what it will deliberately not do.

This keeps the task small enough to be real.

### Step 7. Write Falsifiable Acceptance Criteria

Every criterion should be pass, fail, or unknown.

Good criteria verify the solution you chose, not just the problem area.

Bad:

- `improves approval workflow`
- `handles direct questions better`
- `supports continuity`

Good:

- `IMPLEMENTATION-LEADER.md` requires a task-owned approval packet before asking for plan approval
- `TASK-STATE.schema.json` validates the new ownership and stop-state fields
- `DEBUG-LEADER.md` blocks root-cause claims unless bug artifacts include the required disagreement and writer-chain sections

### Step 8. Say What Does Not Count

When a weak completion path is easy to imagine, list it explicitly.

Examples:

- stronger wording with no generator or gate
- a hidden heuristic with no durable artifact
- off-lane proof relabeled as closure
- vague blocker text with no blocker evidence

### Step 9. Add A Proof Plan When Needed

If the task is likely to be gamed through weak evidence, say what proof will count and what will not.

### Step 10. Run The Fuzziness Check

Reject the draft if any of these are true:

- a reviewer could not tell what exact files or artifacts will change
- two later agents could implement materially different solutions and both claim `done`
- the acceptance criteria only restate the area, not the solution
- the task still hides the main decision behind words like `improve`, `support`, or `handle`
- the task could be closed through a degraded fallback the human would reject

## Human-Facing Resolution Rule

For narrow usability, copy, affordance, IA, or operator-flow fixes:

- collapse the expected resolution into concrete claims a human could agree or disagree with before implementation starts
- make the task concrete enough that someone could sketch the intended result without inventing the main UX choice later
- avoid negative-only claims such as `it will no longer show ...` without stating what replaces the bad state

If you cannot state the expected resolution honestly yet, the task is not ready for concrete implementation mode.

## Concrete Task Review Checklist

Before considering a concrete implementation task ready, confirm:

- the writeup type is really `concrete implementation`
- the main solution shape is chosen
- the concrete changes are named
- the implementation home is named
- the acceptance criteria test the chosen solution
- the weak closure paths are listed under `What Does Not Count` when needed
- the task is specific enough that a human can decide whether to enqueue it or leave it aside as speculative
