# Orchestration Workflow

This file defines the shared task lifecycle across repos under `C:\Agent`.

Use it together with:

- `C:\Users\gregs\.codex\AGENTS.md` for shared glossary, task artifact structure, and doc precedence
- `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md` for task writeup types and the canonical `TASK.md` specificity standard
- `C:\Users\gregs\.codex\Orchestration\FILE-NAMING.md` for shared naming rules
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md` for the machine-readable task-state contract
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json` for machine validation of `TASK-STATE.json`
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.md` for machine-readable pass-closeout state
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.schema.json` for machine validation of `PASS-<NNNN>-CHECKLIST.json`
- `C:\Users\gregs\.codex\Orchestration\AUDIT-RESULT.md` for machine-readable auditor results
- `C:\Users\gregs\.codex\Orchestration\AUDIT-RESULT.schema.json` for machine validation of `PASS-<NNNN>-AUDIT.json`
- `C:\Users\gregs\.codex\Orchestration\Prompts\README.md` for reusable orchestration prompt templates
- `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md` for shared testing artifact rules
- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md` for shared debugging artifact rules
- `C:\Users\gregs\.codex\Orchestration\Processes\AUTOIMPROVEMENT.md` for the shared daily orchestration self-improvement loop
- `C:\Users\gregs\.codex\Orchestration\Exemplars\README.md` for exemplar entry points
- `C:\Users\gregs\.codex\Orchestration\Exemplars\TASK.md` for the intended `TASK.md` shape
- `C:\Users\gregs\.codex\Orchestration\Exemplars\TESTING.md` for the intended repo-local `TESTING.md` shape
- `C:\Users\gregs\.codex\Orchestration\Exemplars\REGRESSION.md` for the intended repo-root `REGRESSION.md` shape

## Lifecycle Overview

The default task lifecycle is:

1. Task Creation
2. Task Research
3. Task Breakdown Into Passes
4. Task Implementation Pass
5. Task Unit Testing
6. Task Handoff
7. Task Regression Testing
8. Task Debugging
9. Task Closure

This is not a strictly one-way sequence.

- `Task Implementation Pass`, `Task Unit Testing`, and `Task Handoff` usually repeat until the planned passes are complete.
- `Task Debugging` usually branches off `Task Regression Testing` and loops back into regression reruns.
- `Task Research` can be reopened when a task exposes new unknowns, but it should not silently replace the main task artifacts.
- when a task uses `Tracking/Task-<id>/TASK-STATE.json`, update it only at durable state transitions rather than on every thought or intermediate experiment
- when a leader persists a durable `TASK-STATE.json` transition that changes `phase`, `current_pass`, or `current_gate`, commit that checkpoint and push it upstream before proceeding

## Leader-Owned Git Narrative

Leaders own git history for the workflow states they persist.

Rules:

- the leader that owns the durable transition also owns the final commit creation and commit message wording for that transition
- workers may change code, tests, docs, or task artifacts inside their assigned scope, but they should not own the git narrative by default
- do not burden workers with final commit-message process; at most they may hand a short human-readable change summary back to the leader when useful
- prefer one leader-owned commit that closes a meaningful workflow step over multiple tiny commits whose subjects only describe bookkeeping
- if bookkeeping can be bundled honestly into the nearest substantive leader-owned commit, do that instead of creating a process-only commit
- avoid low-signal subjects such as `checkpoint`, `start pass`, `sent toast`, `update state`, or `docs update` unless the commit truly contains only that durable action and no more meaningful human summary is possible

Commit subject contract:

- use a subject that includes both:
  - a task-locating prefix
  - a one-sentence human-readable summary of the intent or delivered outcome
- preferred shape:
  - `<task-id> <locator>: <human-readable summary>`
- common locators:
  - `research`
  - `plan`
  - `pass-<NNNN>`
  - `regression`
  - `debug`
  - `closure`
  - `checkpoint` only when no stronger locator fits

Commit body contract:

- when practical, include short fields that make the durable state easy for a human to recover:
  - `Phase:`
  - `Pass:`
  - `Gate:`
  - `Why:`
  - `Proof:`
- the body can omit fields that are genuinely not applicable, but it should still make the reason for the commit legible to a later reader

Summary wording rules:

- describe the work or decision, not just the mechanical action that happened afterward
- prefer intent-and-outcome wording such as `add transcript text search and a minimal browsing UI`
- avoid process-only wording such as `close checklist`, `send toast`, or `enter implementation` when a more human-readable explanation exists
- for leader-owned checkpoints that are mostly state transitions, explain the transition in human terms such as `approve the clip-first plan and enter implementation`

Examples:

- `task-0007 research: narrow the first transcript consumer to clip-first inspection and search`
- `task-0007 plan: split transcript consumer work into listing/inspect and search/UI passes`
- `task-0007 pass-0000: add transcript session and clip inspection endpoints`
- `task-0007 pass-0001: add transcript text search and a minimal browsing UI`
- `task-0007 closure: finalize transcript consumer completion after passing server validation`

## Leader Continuity And User Notification

These rules apply across task-owning leaders and any top-level orchestrator that delegates to them.

- when the human gives a standing umbrella instruction such as `finish the remaining passes`, `continue until complete`, or `keep going unless blocked`, that instruction survives delegated-leader rotation and pass boundaries until one of these happens:
  - the human explicitly changes direction
  - a real human gate is reached
  - a real blocker is reached
  - the task honestly reaches its requested finish line
- do not treat a delegated leader stopping at its designed handoff boundary as permission to pause the whole task if the standing user instruction still implies continued work
- when the human explicitly asks for a subagent or delegated leader to be run, interpret that as an instruction to actively monitor it until it reaches a terminal state or a real human gate; do not `spawn and forget`
- active monitoring means entering a real `wait_agent` polling loop with a maximum cadence of 60 seconds between checks until that delegated work reaches a terminal state or a real human gate
- do not claim to be `watching` delegated work unless you are actively in that polling loop right now; if you are no longer polling, say that you are not actively monitoring at this moment
- status questions, clarification questions, or brief side conversations from the human do not implicitly cancel that monitoring obligation
- when a delegated leader is still active, answer those side questions in commentary and then resume the polling loop immediately unless the human explicitly changes direction
- do not send a `final` answer, task-complete closeout, or equivalent end-of-turn summary while a delegated leader is still active unless the human explicitly says to stop supervising or to switch into status-only mode
- when any spawned subagent reaches a terminal state, report that milestone to the human immediately
- if additional repo inspection or verification is useful, do that after the immediate milestone update rather than delaying the user-visible status report
- immediate milestone reporting does not remove the obligation to verify durable task state before claiming more than the completed milestone actually proves

## Default Prompt Entry Points

The lifecycle phases above are the source-of-truth workflow model.

The default operational entry points for that model are:

- use `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md` when one agent should supervise the whole post-creation lifecycle in serial
- use `C:\Users\gregs\.codex\Orchestration\Prompts\RESEARCH-LEADER.md` to own `Task Research`
- let `RESEARCH-LEADER.md` own local problem decomposition, per-problem grounding, and synthesis by default
- use `RESEARCH-BRIEFER.md` only when the human explicitly wants an external critique packet after local research already exists
- use `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md` as the default lead entry point for both `Task Breakdown Into Passes` and one `Task Implementation Pass` at a time
- use `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md` to own `Task Regression Testing` when the regression phase needs task-level orchestration
- use `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md` to own `Task Debugging` when the debugging phase needs branch coordination
- inside that implementation flow, treat `IMPLEMENTER.md`, `UNIT-TESTER.md`, and `AUDITOR.md` as per-pass supporting gates rather than separate top-level lead entry points

Operationally, that means:

- `TASK-LEADER.md` is the thin top-level lifecycle supervisor after `Task Creation` when a single orchestrator is desired
- research is usually run by `RESEARCH-LEADER.md`
- local task research is the default; optional external critique is a later add-on, not the main research path
- planning plus explicit human `PLAN.md` approval is usually Stage A of `IMPLEMENTATION-LEADER.md`
- pass execution is usually Stage B of `IMPLEMENTATION-LEADER.md`
- each closed implementation pass should normally rotate to a fresh `IMPLEMENTATION-LEADER.md` instance before the next pass begins
- regression and debugging can later be dispatched by `TASK-LEADER.md` when those phases are needed
- regression is usually owned by `REGRESSION-LEADER.md` when the task has entered a substantial regression phase
- debugging is usually owned by `DEBUG-LEADER.md` when the task has an active multi-branch bug investigation
- `Task Unit Testing` and `Task Handoff` still exist as lifecycle phases, but in the default implementation flow they are nested gates inside each approved pass loop rather than separate task-owning leader launches

When `TASK-LEADER.md` is not in use, the specialized leaders above can still be launched directly from the relevant phase.

## Required UI Surface Gate

When a pass creates or materially changes a visible user-facing screen, card, control, placeholder, icon set, or navigation shell, treat interface review as a required gate rather than optional polish.

Rules:

- require an `INTERFACE-DESIGNER.md` review before pass closeout for materially changed visible UI
- that review must compare the approved design reference against the implemented surface as rendered in the emulator or on-device, not just against code or mockup prose
- when paired HTML and screenshot mockups exist, treat the HTML or source composition as the primary structural truth and the screenshot as supporting visual evidence
- the interface review must call out both semantic drift and structural drift, including:
  - ambiguous card meaning
  - vague empty-state or placeholder wording
  - icon-family substitutions
  - missing card backgrounds, wrong visual-weight parity, or missing rounded-container treatment
  - composite controls that should be reconstructed from the source composition instead of being approximated with unrelated stock widgets
- a UI pass is not ready for closeout when the implementation still visibly miscommunicates the mockup's structure or meaning, even if the code is technically functional

## Human-Facing Outcome Gate

When a task, pass, regression lane, or debug branch touches a human-facing surface, leaders must keep the target human experience explicit instead of letting a narrower technical proxy silently take over.

Rules:

- restate in plain terms what a human should perceive when the work is actually done
- do not let easier-to-prove proxies stand in for that outcome just because they are easier to instrument, such as:
  - state existence
  - possession or control without credible presentation
  - geometry or triangles present without a believable visible result
  - debug-only or proof-only screenshots
  - placeholder, fallback, or reference-pose states that are visibly below the intended product bar
- treat degraded fallbacks as checkpoints or diagnostic evidence unless the human explicitly accepts them as the real end state
- when the current implementation truth is a fallback rather than the intended experience, preserve that gap explicitly in task artifacts instead of smoothing it over in plan, regression, bug, or closure wording
- if the human correction path is effectively `closer, but not root-level`, treat that as a signal to run the `VISION-HARVESTER` workflow and encode the smallest durable fix at the right shared or local layer
- closure preflight for human-facing work must be able to answer both:
  - why a human would now consider the surface fixed
  - what human-visible behavior is still degraded, if any

## Delegated Prompt Invocation Rule

The prompt templates under `C:\Users\gregs\.codex\Orchestration\Prompts\` are operational workflow entry points, not passive reference material.

Rules:

- if a task has already entered the orchestrated lifecycle and the human asks to continue, implement, keep going, or finish that task, default to launching the appropriate delegated leader rather than merely reading its prompt and then doing the work ad hoc in the main thread
- when the next phase is not yet certain, default to launching `TASK-LEADER.md` as a delegated subagent so it can choose the next honest phase from durable task evidence
- when the next phase is already clear from durable artifacts, launch the corresponding specialized leader as a delegated subagent rather than treating its prompt as advisory prose
- reading a leader prompt for understanding does not count as following that workflow
- if you choose to bypass the delegated workflow and act locally in the main thread, treat that as an explicit workflow bypass and do it only when:
  - the human explicitly asks to bypass orchestration, or
  - there is a documented shared-workflow exception that clearly applies
- do not claim to have run the task through `TASK-LEADER`, `IMPLEMENTATION-LEADER`, `REGRESSION-LEADER`, `DEBUG-LEADER`, or `RESEARCH-LEADER` unless that leader was actually launched as a delegated subagent
- do not mark a task complete from main-thread ad hoc execution if the delegated workflow would still require another owned phase such as regression or debugging

Default continuation rule:

- for an existing task with `TASK.md` already on disk, unqualified requests such as `continue`, `implement this task`, `keep going`, or `finish it` should default to delegated orchestration
- unless the human explicitly says otherwise, that means launching `TASK-LEADER.md` to supervise the remaining lifecycle rather than continuing locally by personal discretion

## Durable Product Design Anchor

When a repo has one repo-local design document that carries the enduring product direction across many later tasks, name that file:

- `Design/GENERAL-DESIGN.md`

Treat that document as the holder of the product's general design soul:

- long-lived product intent
- north-star workflow or user value
- durable architecture direction
- major later-stage capabilities that current tasks are still marching toward

Working rules:

- keep `GENERAL-DESIGN.md` durable and forward-looking rather than turning it into pass history
- keep it repo-local by default rather than nesting it under one task unless the repo explicitly documents a different structure
- task-owned pass plans, audits, bugs, and regression runs still belong in their usual task artifacts
- when task creation, task research, or task harvesting needs the repo's high-level direction, read `GENERAL-DESIGN.md` early when it exists
- when newer task artifacts intentionally supersede part of `GENERAL-DESIGN.md`, call out that override explicitly instead of silently ignoring the design anchor

## Task Creation

### When It Starts

Use this workflow when a new workstream needs durable ownership, scope, and task artifacts.

### Inputs

- the initial problem statement
- current constraints and assumptions
- known implementation home if already decided

### Outputs

- `Tracking/Task-<id>/TASK.md`
- the initial task folder scaffold under `Tracking/Task-<id>/`
- optional initial `Tracking/Task-<id>/TASK-STATE.json` when structured orchestration state is in use

### Task Writeup Standard

Use `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md` as the canonical standard for:

- task writeup types such as concrete implementation, consensus, and research
- required base `TASK.md` sections
- extra sections for human-facing or easy-to-game tasks
- specificity and falsifiability rules
- the rule that enqueue-ready concrete tasks must name the chosen solution shape and the concrete changes that will happen

Task creation should not leave the main solution in a fuzzy superposition of multiple designs.
If the work is not yet concrete enough for implementation mode, write an honest consensus or research task instead.

### Exit Bar

- a task id exists
- `TASK.md` exists and follows `TASK-CREATE.md`
- the task context, goals, non-goals, and acceptance criteria are written down
- the expected implementation home is named

### Usually Followed By

- `Task Research`
- optional launch of `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md` to supervise the task after creation

## Task Research

### Default Lead Prompt

The default task-owning lead for this phase is:

- `C:\Users\gregs\.codex\Orchestration\Prompts\RESEARCH-LEADER.md`

When a top-level lifecycle supervisor is in use, it usually dispatches that research leader rather than replacing it.

Use it when one agent should own the research phase end to end, including:

- local grounding from `TASK.md`, `HANDOFF.md`, code, and repo docs
- creating or refreshing `RESEARCH-PLAN.md`
- creating or refreshing `RESEARCH-ANALYSIS.md`
- writing bounded per-problem research artifacts under `Tracking/Task-<id>/Research/`
- distilling the local findings into `RESEARCH.md`
- optionally launching `RESEARCH-BRIEFER.md` later only when the human explicitly wants external critique after local research exists

If older docs or habits still point at the retired single-name research prompt, migrate them to `RESEARCH-LEADER.md` as the default research owner and reserve `RESEARCH-BRIEFER.md` for optional external critique packaging only.

### When It Starts

Use this workflow when the task needs baseline reading, reference gathering, constraint discovery, or design comparison before implementation planning is trustworthy.

### Inputs

- the current `TASK.md`
- existing repo notes, code, and task context
- any relevant external references only when the task or human explicitly calls for them

### Outputs

- task-owned research artifacts under `Tracking/Task-<id>/Research/`
- task-root research coordination docs:
  - `RESEARCH-PLAN.md`
  - `RESEARCH-ANALYSIS.md`
  - `RESEARCH.md`

### Exit Bar

- the major unknowns are narrowed enough to plan implementation honestly
- important constraints and assumptions are preserved in task-owned artifacts
- the problem list is bounded to questions that can materially change architecture, pass planning, or debugging direction
- the per-problem research artifacts compare canonical, frontier, and repo-fit options rather than becoming a literature dump
- the research artifacts are focused on task guidance rather than becoming the canonical task definition
- any task-root research summary stays consistent with the deeper `Research/` material it is distilling

### Usually Followed By

- `Task Breakdown Into Passes`

## Task Breakdown Into Passes

### Default Lead Prompt

The default task-owning lead for this phase is:

- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`

When `TASK-LEADER.md` is in use, it usually dispatches `IMPLEMENTATION-LEADER.md` for this phase rather than planning locally.

In the current shared model, pass planning is usually Stage A of `IMPLEMENTATION-LEADER.md`, not a separate planning-only lead workflow.

That leader owns:

- requiring `RESEARCH.md` before planning starts
- creating or refreshing `PLAN.md`
- pausing for explicit human approval of the current `PLAN.md`
- updating `TASK-STATE.json` to reflect the durable planning baseline when structured state tracking is in use

### When It Starts

Use this workflow when the task scope is understood well enough to sequence implementation into bounded passes.

### Inputs

- `TASK.md`
- relevant `Research/` artifacts
- current codebase baseline

### Outputs

- `Tracking/Task-<id>/PLAN.md`

### Exit Bar

- the pass order is explicit
- each pass has a clear objective
- each pass has a practical verification plan
- each pass has an exit bar that can be checked honestly

### Usually Followed By

- `Task Implementation Pass`

## Task Implementation Pass

### Default Lead Prompt

The default task-owning lead for this phase is:

- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`

When `TASK-LEADER.md` is in use, it usually dispatches `IMPLEMENTATION-LEADER.md` for this phase rather than executing passes locally.

In the current shared model, pass execution is usually Stage B of `IMPLEMENTATION-LEADER.md`.

Each implementation-leader instance should normally close one execution pass and then hand control back rather than continuing across multiple closed passes in one long-lived thread.

That leader owns:

- selecting the next approved pass from `PLAN.md`
- updating `TASK-STATE.json` before durable pass transitions when structured state tracking is in use
- spawning bounded per-pass workers such as `IMPLEMENTER.md`, `UNIT-TESTER.md`, and `AUDITOR.md`
- integrating their results
- closing the pass with handoff, commit, push, and final toast

### When It Starts

Use this workflow when executing one planned pass from `PLAN.md`.

### Inputs

- `TASK.md`
- `PLAN.md`
- `HANDOFF.md`
- relevant code and task artifacts

### Outputs

- the code and docs for the current pass
- the pass audit under `Tracking/Task-<id>/Testing/`
- an updated task `HANDOFF.md`
- optional structured companions when pass-closeout state is being tracked explicitly:
  - `Tracking/Task-<id>/Testing/PASS-<NNNN>-CHECKLIST.json`
  - `Tracking/Task-<id>/Testing/PASS-<NNNN>-AUDIT.json`

### Required Sequence

1. Expansion: implement the next pass.
2. Contraction: solidify it with unit tests.
3. Audit: write the pass audit under `Tracking/Task-<id>/Testing/`.
4. Handoff: update the relevant task `HANDOFF.md`.
5. Commit: create a git commit for the completed pass.
6. Push: push the completed pass to `upstream`.
7. Notify: send a Windows toast using the `windows-toast-notification` skill as the final step.

Do not leave a pass half-closed. A pass is only complete after the audit, handoff update, commit, push, and final toast are done.

When structured pass-closeout state is in use:

- use `PASS-<NNNN>-CHECKLIST.json` to record which closeout gates are complete
- use `PASS-<NNNN>-AUDIT.json` to record the auditor verdict and structured findings

### Usually Followed By

- another `Task Implementation Pass`
- `Task Regression Testing` once planned passes are complete or the task reaches a release-ready checkpoint

## Task Unit Testing

### Default Supporting Prompt

The default supporting prompt for this phase is:

- `C:\Users\gregs\.codex\Orchestration\Prompts\UNIT-TESTER.md`

`Task Unit Testing` remains a real lifecycle phase, but in the default implementation flow it is usually a per-pass gate inside `IMPLEMENTATION-LEADER.md`, not a separate task-owning lead launch.

### When It Starts

Use this workflow during each implementation pass to prove pass-local correctness.

### Inputs

- the current pass scope
- changed code paths
- existing relevant unit coverage

### Outputs

- unit tests added or updated as needed
- unit-test evidence recorded in the pass audit

### Exit Bar

- new or affected unit tests pass
- the pass audit records what unit coverage was actually exercised
- pass-local correctness is proven without pretending task-level regression has already happened

### Usually Followed By

- the rest of the current `Task Implementation Pass`
- `Task Handoff`

## Task Handoff

### Default Ownership

In the default implementation flow, the task-owning leader that ran the current pass also owns the handoff update.

That is usually:

- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md` during implementation passes
- `C:\Users\gregs\.codex\Orchestration\Prompts\RESEARCH-LEADER.md` only when research materially changes the next recommended step or watchouts

So `Task Handoff` remains a distinct lifecycle phase, but it is usually materialized by the current task-owning leader rather than by launching a separate handoff-only lead role.

### When It Starts

Use this workflow whenever a pass ends, a task pauses, or the next session needs a reliable resume point.

### Inputs

- the current task baseline
- the latest pass outcome or debugging state

### Outputs

- `Tracking/Task-<id>/HANDOFF.md`

### Exit Bar

- the current baseline is written down honestly
- the next step is explicit
- active watchouts, blockers, or bug pointers are visible
- a later session can resume without rediscovering the state

### Usually Followed By

- another `Task Implementation Pass`
- `Task Regression Testing`
- `Task Closure`

## Task Regression Testing

### Default Orchestration

When a top-level lifecycle supervisor is in use, `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md` usually dispatches `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md` for this phase.

The default task-owning lead for a substantial regression phase is:

- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`

That leader may use `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md` as a bounded worker for one executed regression run or focused rerun slice.

### When It Starts

Use this workflow after the planned passes for a task are implemented, or when the task reaches a stable checkpoint that is ready for end-user or operator validation.

### Inputs

- the task baseline
- the repo's shared `REGRESSION.md`
- the built app or service in a runnable state

### Outputs

- `Tracking/Task-<id>/Testing/REGRESSION-RUN-<NNNN>.md`
- `Tracking/Task-<id>/BUG-<NNNN>.md` when the required regression lane fails or is blocked and the task needs debugging or an explicit carry-forward explanation

### Exit Bar

- the real app or operator flow has been exercised honestly for the intended lane
- anything not run is marked `NOT RUN` or `BLOCKED`
- the result artifact states what was actually exercised and what that proves
- the result artifact explicitly names the claimed repo-root regression lane, the exact flow exercised, why that run counts, and any disqualifiers or limitations
- if the required regression lane failed or blocked, the task has an active `BUG-<NNNN>.md` that preserves the issue before the task leaves regression

### Important Rule

`Task Regression Testing` is task-level proof. It should not be implied by pass-local unit coverage.

Repo-root `REGRESSION.md` is authoritative for what lane counts as regression and what counts as passing regression for that repo.

Regression must use the human default lane for the repo under test.

An operator lane may be used for setup, diagnosis, implementation, or supporting evidence, but it does not substitute for regression proof.

Task-owned artifacts may:

- reference which `REGRESSION.md` case ids matter for the current task
- record what was actually run
- record blockers, failures, and limitations honestly

Task-owned artifacts may not:

- redefine the regression lane
- relax the regression pass bar
- substitute supporting proof for a required regression lane without an explicit human change to the repo-root regression matrix

When the required regression lane fails or is blocked:

- record the executed run in `REGRESSION-RUN-<NNNN>.md`
- create or update `BUG-<NNNN>.md` even if the immediate resolution is an upstream human, environment, or operator prerequisite
- route the task into `Task Debugging` or an explicit human gate instead of letting the failure live only inside the regression note

### Usually Followed By

- `Task Closure` if the task-level proof is acceptable
- `Task Debugging` if regression exposes a real issue

## Task Debugging

### Default Orchestration

When a top-level lifecycle supervisor is in use, `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md` usually dispatches `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md` for this phase.

The default task-owning lead for a substantial debugging phase is:

- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`

That leader may use `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-WORKER.md` as a bounded worker for one hypothesis branch, repro lane, crash-analysis lane, or fix-confirmation lane.

### When It Starts

Use this workflow when task-level regression exposes a failure, flaky result, or suspicious behavior against the established task baseline.

### Inputs

- the failing regression slice
- the current handoff and task baseline
- the relevant artifacts and evidence

### Outputs

- `Tracking/Task-<id>/BUG-<NNNN>.md`
- additional `REGRESSION-RUN-<NNNN>.md` artifacts for actual reruns

### Exit Bar

- the current hypothesis is preserved
- exact evidence is written down
- the next debugging or verification step is explicit
- the bug is either narrowed, verified fixed, or left in an honest carry-forward state

### Usually Followed By

- `Task Regression Testing`
- `Task Closure`

## Task Closure

### Default Ownership

When a top-level lifecycle supervisor is in use, `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md` usually owns the final closure check and any task-level summary updates that are still needed after the specialized phases finish.

### When It Starts

Use this workflow when the task appears complete and needs final confirmation rather than another implementation pass.

### Inputs

- `TASK.md`
- `HANDOFF.md`
- completed pass audits
- regression and debugging artifacts, if any

### Outputs

- a final task state reflected in `HANDOFF.md`

### Exit Bar

- the task acceptance criteria are satisfied or any intentional gaps are explicit
- planned passes are complete or intentionally deferred with explanation
- required unit-test evidence exists
- required regression testing has been done honestly
- the repo-root regression lane required for closure has passed, or the task remains explicitly open
- active bugs are either closed or clearly carried forward
- the task can be handed off or considered complete without hidden state

### Closure Regression Claim

Before a task may be marked closed, the closing leader must be able to name:

- the exact repo-root regression lane or case ids required for closure
- the exact `REGRESSION-RUN-<NNNN>.md` artifact that satisfies that lane
- why that run counts for the required lane
- what that run does not prove

For human-facing work, that claim must anchor on a run of the repo's human default lane. Supporting operator-lane or override-lane evidence may be cited in addition, but it cannot replace the default-lane regression artifact.

If that claim cannot be made honestly from durable artifacts, the task is not closure-ready.
