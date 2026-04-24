# Task 0008 Plan

## Planning Verdict

This plan was explicitly approved for implementation on `2026-04-24`.

Implementation starts with `PASS-0000`.

The approved scope is the backend-only dispatch runtime and durable execution-state contract, including:

- the backend readback, freshness, action-gating, and deep-context capabilities the future `Tasks` tab consumes
- an exclusive backend-owned repo checkout or equivalent isolated repo lane for simple task execution instead of the human's shared primary worktree
- explicit restore-commit semantics so unit proof and execution cleanup can reset that owned checkout to a recorded useful commit baseline
- the explicit human-state envelope, meaningful-progress rules, full wait contract, dispatch-readiness rule, multi-run story-selection rule, and attention-priority shape described below
- a doc-to-backend ingest and snapshot model that preserves live runtime truth across git rollback without making ad hoc agent HTTP calls the required primary mechanism
- first proof allowed through Codex or direct backend interactions rather than frontend work

## Planning Basis

This plan is grounded in:

- the backend and Temporal baseline from [Task-0005](../Task-0005/TASK.md)
- the `Tasks` tab product-surface direction from [Task-0009](../Task-0009/TASK.md)
- the `Tasks`-tab dependency expectations in [Task-0009 Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md](../Task-0009/Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md)
- the first-screen action and row contract in [Task-0009 Design/STITCH-PROMPT.md](../Task-0009/Design/STITCH-PROMPT.md)
- the task-local execution-state note at [Design/DURABLE-EXECUTION-STATE-CONTRACT.md](./Design/DURABLE-EXECUTION-STATE-CONTRACT.md)

## State Ownership And Mutation Model

Task-0008 must keep the state split explicit:

- task-definition state lives in task-owned markdown artifacts such as:
  - `TASK.md`
  - `PLAN.md`
  - `HANDOFF.md`
  - testing, bug, and research artifacts
- live execution state lives in the backend-owned task-run runtime contract

The backend-owned live flow must be backed by durable Temporal state so the system does not `forget` that a run is waiting, blocked, sleeping, interrupted, or still progressing after:

- backend process restarts
- quiet periods
- client disconnects
- dashboard or Codex restarts

The UI and readback layer are derivative views over backend truth, not the source of truth.

That means:

- dashboard surfaces may summarize or map state for humans
- direct backend callers may inspect or exercise the contract
- neither the UI nor a client session may become the canonical owner of wait, progress, interrupt, or repo-reset state

Local per-run artifacts such as logs, transcripts, JSONL event captures, or final-message files remain supporting evidence.

They do not become the canonical durable record of:

- whether the run is waiting
- why silence is legitimate or suspicious
- what progress freshness window applies
- which cleanup or reset step is pending

Repo-owned checkout and reset state also belong in the backend-owned runtime contract.

The runtime must persist that state explicitly rather than inferring it later from ad hoc filesystem observation of whatever checkout happens to exist on disk.

## Declared Doc Truth And Runtime Snapshot Model

Task-0008 must treat these git-authored task docs as co-owned declared task/process truth:

- `TASK.md`
- `PLAN.md`
- `HANDOFF.md`
- `TASK-STATE.json`

These remain the declared task-definition truth at the current git revision.

The declared parsing scope must be named explicitly with:

- `declared_worktree_root`
  - the checkout or worktree root the backend parsed
- `declared_task_root`
  - the specific task directory the backend read, such as `Tracking/Task-0008`

The backend runtime does not replace them.

Instead, the backend or orchestrator must durably ingest or snapshot the relevant task-definition state into Temporal-backed run state when it matters, especially:

- before dispatch
- when a backend-owned reconcile decides the declared task definition changed materially
- when later readback needs to explain why a run's captured task contract differs from current git truth

The primary model is doc-to-backend ingest and backend-owned reconcile.

Preferred day-to-day workflow:

1. agents continue writing durable local task docs:
   - `TASK.md`
   - `PLAN.md`
   - `HANDOFF.md`
   - `TASK-STATE.json`
2. backend learns declared task truth by reading and parsing task docs on task-API boundaries within:
   - `declared_worktree_root`
   - `declared_task_root`
3. dispatch captures a task-definition snapshot into Temporal-backed run state
4. task read APIs do read-through reconcile for active runs:
   - backend reads current task docs from `declared_task_root` inside `declared_worktree_root`
   - backend reads active Temporal run state
   - if current docs differ materially from the last captured snapshot for that active run, backend records a new doc snapshot or reconcile event into Temporal-backed run state
5. Temporal owns only non-rewindable runtime facts
6. manual re-ingest or reconcile remains a repair or debugging fallback, not the normal workflow

What does not count as the primary mechanism:

- requiring ordinary agents to keep task state truthful mainly by making direct HTTP calls

A file watcher may be added later if implementation proves it useful, but it is not required by this plan.

The contract should therefore say, explicitly:

- task APIs must parse current task docs on demand
- task APIs must know which `declared_worktree_root` and `declared_task_root` they are reconciling
- for active runs, task APIs must reconcile doc changes into Temporal-backed run snapshots
- this must not rely on a file watcher as the normal path
- this must not rely on direct agent HTTP calls as the normal path

The durable invariant is:

- latest git docs remain the declared task-definition truth
- live runtime truth remains preserved in Temporal-backed run state even if git later rewinds

Rollback is therefore a designed case, not an exception.

If git rewinds after dispatch, the runtime must not forget:

- the captured task-definition snapshot the run started under
- live wait state
- interrupts
- cleanup/reset state
- other in-flight execution truth already recorded durably

Read-through reconcile is scoped to the declared worktree and declared task root, not merely to a coarse repo family or repo identifier.

## Backend Capabilities Task-0009 Expects From Task-0008

Task-0009 is intentionally a UI task, but its durable design brief already assumes specific backend/runtime support from this task.

Task-0008 must leave behind a backend contract strong enough that the later `Tasks` tab can consume truth rather than invent it.

The minimum consumer-facing capabilities are:

- task and active-run readback that can power both:
  - grouped task rows
  - a selected-task detail pane
- durable state distinctions and reason inputs for:
  - `ready`
  - `running`
  - `waiting_for_human`
  - `blocked`
  - `sleeping_or_stalled`
  - `interrupted`
  - `completed`
  - `failed`
- enough durable explanation to render honest reason lines and next-step summaries without making the UI infer meaning from raw phase or status labels alone
- freshness and trust signals for:
  - last meaningful progress
  - last trusted backend snapshot
  - when silence becomes suspicious
  - when live state is stale enough to down-rank control confidence
- dispatch readiness and post-dispatch readback so the UI can tell the human:
  - why a task is ready
  - what contract it will use
  - that a run now exists
  - what the next expected signal is
- bounded action gating for `Dispatch`, `Poke`, and `Interrupt`, including durable reasons when an action is not currently allowed
- deep-context launch provenance strong enough for `Open Thread` or equivalent working-context launch without transcript hunting
- durable provenance that keeps a real task and its active run connected without making the UI invent which run or context belongs to which task
- an execution-lane model where simple tasks run inside an exclusive backend-owned checkout rather than the human's shared worktree
- restore-baseline semantics that name which commit the backend may reset that owned checkout to during unit proof or cleanup

Task-0008 does not need to ship the final human-facing copy or layout.

It does need to expose enough durable machine-readable truth that Task-0009 can show:

- task meaning
- current state
- reason
- freshness
- next expected step
- valid actions
- deeper context

## Explicit Contract Decisions For Approval

### 1. Human-Readable State Envelope

Every human-facing task or run readback must expose one explicit durable envelope with at least:

- `state`
- `reason_code`
- `state_summary`
- `evidence_refs`
- `next_owner`
- `next_expected_event`
- `suspicious_after`
- `action_block_reasons`

This envelope is the canonical backend input for reason lines, next-step summaries, and action availability in the planned Task-0008 backend or readback contract.

The UI may map or style it, but it must not invent a weaker explanation model.

This naming belongs to the planned Task-0008 runtime or readback contract rather than the shared `.codex` `TASK-STATE.json` schema.

`state_summary` must be:

- short
- plain-language
- present-tense
- non-jargony
- understandable without opening artifacts first

`evidence_refs` is a typed launchable structure, not loose prose:

- `Array<object>` where each object contains:
  - `type`
    - `task_artifact`
    - `run_artifact`
    - `transcript`
    - `api_resource`
    - `owned_checkout`
    - `commit`
    - `test_result`
  - `label`
  - `uri`
- optional `line_range`
- optional `created_at`

Separate from the run or state envelope, task readback must also expose one minimal task-level meaning field:

- `meaning_summary`

`meaning_summary` is:

- the primary task-level summary of why the task exists or why the human should care
- derived from task-definition docs
- distinct from `state_summary`, which remains the current run or state summary

This field stays intentionally narrow so the future `Tasks` tab can reduce markdown reconstruction burden without inventing a large family of overlapping task-level summary fields.

### 2. Meaningful Progress

`Meaningful progress` is narrower than `any output`.

It counts only when durable evidence shows real movement in the task contract, for example:

- a durable state transition
- wait contract creation, update, or clearance
- a successful or failed bounded tool step that changes what can happen next
- a material artifact mutation in the owned execution lane
- a persisted test result
- a git commit or reset to a recorded restore target
- an agent-authored checkpoint that changes the next expected step, wait contract, or supervisor understanding

It does not include:

- raw log noise
- heartbeats with no contract change
- repeated identical status writes
- filesystem touches that do not change the task's real execution story

Freshness must vary by state category, not one global timeout:

- `running`
  - ages against expected progress cadence and becomes suspicious fastest
- `waiting_for_human`
  - ages against the explicit wait stale deadline for the requested human action
- `blocked`
  - ages more slowly and becomes suspicious when the blocker has not been revisited by the promised review point
- terminal states
  - do not use freshness to imply active progress expectations

### 3. Full Wait Contract

`Waiting` must not be a junk drawer.

Any run claiming a healthy wait must persist:

- who or what it is waiting on
- why progress cannot continue without that dependency
- what event or condition would resume work
- whether human action is required
- where the human should click, open, or act when human action is required
- who owns the next move
- when the wait itself becomes stale or suspicious

`waiting_for_human` is valid only when those fields are explicit enough that the human does not have to play detective.

`human_action_target` must be explicit and typed, not buried in prose.

Minimum structure:

- `human_action_target.kind`
  - `task_artifact`
  - `run_artifact`
  - `transcript`
  - `api_resource`
  - `owned_checkout`
  - `approval_action`
- `human_action_target.label`
- `human_action_target.uri`
- optional `human_action_target.line_range`

### 4. Dispatch Readiness Rule

A task is dispatch-ready only when the backend contract can prove all required preconditions, at minimum:

- the current plan is approved
- no unresolved blockers prevent execution
- no still-owning active run already controls the task's live story unless the new dispatch is an explicit superseding action
- an accepted exclusive execution lane is acquired or reserved
- `declared_worktree_root` is resolved
- `declared_task_root` is resolved
- a valid restore baseline is recorded for that lane
- a task-definition snapshot is captured from the current declared docs
- the expected first signal is recorded
- the first suspiciousness deadline is recorded

Dispatch readiness is about safe dispatch, not merely enabling a button.

### 5. Multiple Runs Per Task

The backend must choose one current human-facing story explicitly rather than leaving later clients to infer it.

The rule is:

- at most one run may own the current live story for a task at a time
- an active nonterminal run that has not been superseded owns that story
- retries or superseding runs must explicitly displace the older run
- terminal runs stay in history and may remain visible as recent evidence, but they do not keep owning the task once task-level state has moved back to `ready`, `blocked`, or another non-run state
- when no run currently owns the live story, task readback must say so explicitly rather than pretending an older run is still the present tense

Any superseding dispatch must also record:

- `supersede_reason`
- `superseded_by_actor`
- disposition of the older run:
  - `left_in_history`
  - `interrupted`
  - `marked_superseded`
  - `cleanup_blocked`

### 6. Attention-Priority Shape

The backend contract should expose explicit attention-priority inputs so Task-0009 does not have to invent fragile heuristics.

The minimum durable shape is:

- `attention_level`
- `attention_reason`
- `attention_sort_key`

This lets the later `Needs Attention` region stay derivative from backend truth while still being fast and humane on re-entry.

Rough precedence should be:

- `urgent`
  - sleeping or stalled runs
  - waits that went stale
  - cleanup or reset failures
- `needs_attention`
  - waiting for human with an actionable target
  - blocked states needing review
  - dispatch-ready tasks intentionally surfaced for action
- `watch`
  - healthy running work or non-stale blocked work worth monitoring
- `none`
  - terminal history or states that do not currently compete for attention

### 7. Completed Means Truly Done

`completed` must mean no further human judgment, review, or approval is required for the task or run to be considered done in the human-facing system.

If review or approval is still required, the state must not present as `completed`.

Instead it must surface as:

- `waiting_for_human`
- with a review or approval reason code
- and `human_action_target.kind = approval_action` or an equivalent explicit approval target

This rule exists to protect the trust model:

- the system must not look finished while still quietly waiting on human signoff

## Additional Hardening Decisions

### 7. Restore-Target Lifecycle

A later restore target is allowed only when:

- the runtime records an explicit checkpoint event that justifies it
- the actor who approved it is persisted durably
- the checkpoint is specific enough to count as a known-good restore point for proof or cleanup

Useful checkpoint examples:

- a persisted test result bundle
- a committed known-good code checkpoint
- an explicit supervisor-approved reset point

If cleanup cannot reach the chosen restore target, the runtime must:

- report the failure durably
- preserve the older fallback restore commit if one exists
- mark cleanup as blocked rather than pretending reset succeeded

### 8. Git Rollback And Divergence Model

Task-definition docs and runtime state have different jobs:

- current git docs are the declared task-definition truth
- Temporal-backed run state is the durable record of live execution truth and side effects

At dispatch, the backend must capture a task-definition snapshot from the declared docs, including enough metadata to identify:

- the task docs revision used for dispatch
- the task-definition fields the run is executing against

If git later rewinds or changes materially, task readback must be able to show:

- the current declared task revision
- the run's captured task snapshot revision
- whether they match
- a divergence status or equivalent explanation when they do not

This divergence model exists so rollback does not silently erase runtime truth.

The system should report mismatch honestly rather than pretending:

- the run always matches current git
- or the old runtime snapshot replaced the latest declared docs

### 9. Contract Must Never Allow

The contract must never allow:

- a git rewind to silently erase recorded wait, interrupt, or cleanup truth
- current git inspection alone to replace the run's captured task-definition snapshot
- ad hoc agent HTTP calls to be the required primary mechanism for keeping task state honest
- task readback to hide that current git and the run snapshot differ
- runtime side effects to be forgotten just because the repo moved backward
- `state_summary` that only repeats backend jargon or raw phase labels
- `evidence_refs` as untyped free-form text
- `waiting_for_human` with no actionable target
- one global freshness timeout applied equally to running, blocked, and waiting states
- superseding a run without recording why and what happened to the older run
- promoting a later restore target without a durable known-good checkpoint
- claiming cleanup succeeded when the owned checkout did not actually reach a recorded restore commit

## PASS-0000 Encode The Chosen Task-Run Contract

### Objective

Encode the chosen runtime shape for task dispatch before implementation spreads ambiguity:

- a separate backend-owned task-run workflow and API contract
- not a Git-tracked recurring job spec

### Implementation Notes

- define the run lifecycle states
- define legal wait reasons
- define what counts as sleeping or stalled
- define poke and interrupt rules
- define session and thread provenance fields
- define the exact backend endpoints and durable run fields that later passes will implement
- define the exclusive repo-ownership model for simple execution:
  - one backend-owned checkout or equivalent isolated repo lane per run or per reusable owned lane
  - no ordinary simple-task execution in the human's shared primary worktree
- define restore-commit semantics:
  - capture the dispatch baseline commit
  - define when a later commit becomes an allowed restore target
  - define which reset operations are allowed during unit proof and cleanup
- define the explicit human-readable state envelope:
  - field names
  - required presence rules
  - how task-level and run-level readback use it
  - `state_summary` writing rule
  - typed `evidence_refs`
- define what counts as meaningful progress and what does not
  - including freshness by state category
- define the full wait contract and make `waiting_for_human` non-ambiguous
  - including typed `human_action_target`
- define the dispatch-readiness rule and its preconditions
- define how task readback chooses the one run that owns the current live story
  - including superseding-dispatch recording rules
- define the attention-priority shape so later UI grouping stays derivative from backend truth
  - including rough precedence rules
- define restore-target approval and failure behavior
- define the doc-to-backend ingest and reconcile model:
  - which declared docs are authoritative at current git
  - what task-definition snapshot is captured at dispatch
  - how rollback or git drift is represented without forgetting runtime truth
  - why direct agent HTTP calls are not the required primary mechanism
  - task APIs parse current docs on demand
  - active-run task reads perform read-through reconcile into Temporal-backed snapshots
  - manual re-ingest stays a fallback path, not the normal workflow
  - reconcile scope is explicitly keyed by `declared_worktree_root` and `declared_task_root`
- define a short `must never allow` block for implementation guardrails
- freeze the consumer-facing readback fields Task-0009 will later depend on:
  - current durable state
  - reason inputs
  - next expected step and owner
  - freshness and staleness inputs
  - valid actions and action-block reasons
  - deep-context launch metadata
  - owned-checkout identity and restore-baseline readback
  - story-run selection inputs
  - attention-priority inputs

### Verification

- the contract can explain when a silent run is healthy and when it deserves intervention
- the contract can support both backend decisions and later client summaries
- the contract is specific enough that Task-0009 will not need to guess why a task is ready, blocked, sleeping, or waiting
- the contract is specific enough that later implementation does not have to guess which repo lane it owns or which commit cleanup may reset to
- the contract is specific enough that later clients do not have to invent their own reason-envelope, progress, wait, story-run, or attention heuristics
- the contract is specific enough that git rollback cannot silently erase live runtime truth or blur declared-doc truth with captured run truth

### Exit Bar

- the system no longer has to guess what `running` means or which runtime model owns task dispatch

## PASS-0001 Build Task Dispatch And Run Persistence

### Objective

Add durable backend-owned task dispatch and task-run persistence.

### Implementation Notes

- create the backend task-run model
- add dispatch entrypoints
- persist state transitions durably
- preserve task and run identity across retries or reconnects
- provision or bind an exclusive backend-owned checkout or equivalent isolated repo lane for simple task execution
- capture the dispatch baseline commit before the run mutates that owned checkout
- record any later allowed restore target when the runtime intentionally creates one
- ingest the declared task docs and persist a dispatch-time task-definition snapshot in run state
- persist the declared parse scope for that snapshot:
  - `declared_worktree_root`
  - `declared_task_root`
- make `GET /api/v1/tasks` and `GET /api/v1/tasks/{task_id}` capable of returning task-level readback with the latest relevant run snapshot, freshness inputs, next-step inputs, and action-availability inputs
- make `GET /api/v1/task-runs/{run_id}` capable of returning durable run detail for later task-detail drilldown
- make run readback explicit about:
  - which owned checkout the run may mutate
  - which commit it will restore to for proof or cleanup
  - whether the owned checkout is currently clean, dirty, or reset in progress
- make task and run readback explicit about:
  - the human-state envelope
  - meaningful-progress timestamps and summary
  - wait-contract fields
  - attention-priority fields
  - whether the run currently owns the task's live story
  - which declared task-doc snapshot the run captured
  - which declared worktree and task root that snapshot came from
  - whether current git task docs diverge from that captured snapshot

### Verification

- a dispatched task produces a durable run record
- the record survives client restarts and reconnects
- backend readback can identify which durable run currently owns the task's live execution story
- simple execution can start without touching the human's shared primary worktree
- backend readback can explain, in one durable envelope, why the task is in its current state and what should happen next
- after a git rewind, backend readback still preserves the run's prior wait, interrupt, cleanup, and captured task-snapshot truth

### Exit Bar

- task dispatch is a real runtime capability, not only a future UI intention

## PASS-0002 Add Supervision, Poke, And Interrupt

### Objective

Make the runtime recoverable when work goes silent, drifts, or needs to be stopped.

### Implementation Notes

- add supervision logic for sleeping or stalled detection
- add poke behavior that is conditioned on the durable contract
- add interrupt behavior and reflected state transitions
- preserve human-readable reason fields for intervention actions
- expose durable action-eligibility answers and rejection reasons so later clients only enable `Poke` or `Interrupt` when the contract says they are valid
- define cleanup and interruption behavior for the owned checkout:
  - when the runtime must reset to the baseline commit
  - when a later allowed restore target is preferred
  - when repo cleanup is blocked and must be surfaced honestly
- drive sleeping detection from the explicit meaningful-progress rule and stale wait-contract rule rather than from raw output absence alone
- keep `waiting_for_human` valid only when the stored wait contract says what the human should do and when that wait becomes stale

### Verification

- the backend can distinguish healthy waiting from unhealthy silence
- poke and interrupt actions change durable state honestly
- the backend can explain why an action is available or unavailable without relying on UI-only heuristics
- interrupt and cleanup paths leave the owned checkout at a recorded useful commit or report the exact block honestly
- sleeping versus running is anchored to meaningful-progress truth instead of raw output noise
- runtime supervision continues from Temporal-backed truth even if git state changes or clients disconnect

### Exit Bar

- the system can supervise long-running work instead of merely launching it

## PASS-0003 Add Context Readback And Operator Contract

### Objective

Connect durable task runs to deeper working context and prepare the API shape that Codex and later clients will consume.

### Implementation Notes

- preserve thread, transcript, or session provenance for each run
- define context-ready fields or launch-target readback
- expose the minimal backend contract needed by Codex first and [Task-0009](../Task-0009/TASK.md) later
- expose enough detail for Task-0009 to render:
  - `What changed recently`
  - `Next expected step`
  - `Open Thread` or equivalent deep-context launch
- expose enough execution-lane context for operators and proof to know:
  - which checkout was used
  - what baseline commit anchored it
  - what commit the runtime can reset it to safely
- expose enough run-history context that task readback can point to the one run currently owning the task story while still linking older runs as history
- expose enough divergence context that readback can say when current declared task docs differ from the run's captured snapshot

### Verification

- Codex or a direct backend caller can recover a useful deeper working context from durable state
- later clients no longer have to infer provenance from brittle local heuristics
- Task-0009's detail-pane needs can be satisfied from backend truth rather than markdown-first reconstruction
- backend proof can show the owned checkout and restore baseline without relying on hand-written cleanup notes
- backend proof can show why one run owns the present-tense story and why older runs do not
- backend proof can show declared-doc snapshot capture and honest git-vs-runtime divergence reporting after rollback

### Exit Bar

- deep-context recovery becomes a backend capability instead of manual detective work

## PASS-0004 Audit And Regression

### Objective

Prove the runtime is durable enough for real supervision and later humane UI consumption.

### Implementation Notes

- add focused tests for state transitions and supervision rules
- capture task-owned proof for poke and interrupt behavior
- run the relevant backend proof lanes honestly
- capture proof that the backend readback exposes the freshness, next-step, action-gating, and deep-context fields the later `Tasks` tab depends on
- capture proof that simple task execution and unit cleanup can reset the owned checkout to a recorded useful commit baseline
- capture proof that:
  - meaningful progress advances freshness
  - invalid waits go stale honestly
  - dispatch is rejected until readiness preconditions are satisfied
  - superseded or terminal runs stop owning the current task story
  - attention-priority fields stay aligned with the durable state contract
  - git rewind preserves prior runtime truth while readback reports the declared-doc snapshot mismatch honestly

### Verification

- tests prove the state machine and supervision rules
- task-owned evidence proves at least one real dispatch, poke, and interrupt path through backend interaction
- task-owned evidence proves the backend consumer contract is strong enough for later `Tasks`-tab state explanation and action gating
- task-owned evidence proves repo cleanup is a concrete reset-to-commit operation, not an ambiguous shared-worktree cleanup promise
- task-owned evidence proves the backend, not the client, decides the state envelope, meaningful progress, wait validity, story-run selection, and attention priority
- task-owned evidence proves declared git truth and captured runtime truth can diverge without the system forgetting either one

### Exit Bar

- the dispatch layer is safe to connect to the first interactive `Tasks` tab slice

## Watchouts

- do not collapse all inactive states into `waiting`
- do not let frontend or Codex client memory become the source of truth
- do not add `poke` without a durable justification rule
- do not make context recovery depend on the human reconstructing provenance by hand
- do not make Task-0009 infer human-visible reason lines from opaque backend labels with no durable explanation inputs
- do not let ordinary simple-task execution depend on mutating or cleaning the human's shared primary worktree
- do not say the runtime can `roll back later` unless the recorded restore commit and allowed reset path are explicit
- do not let `running` mean only `stdout happened`
- do not let `waiting_for_human` become a junk drawer with no explicit action target or stale deadline
- do not let the UI choose the current story run or attention order from ad hoc heuristics when the backend contract can expose it directly
- do not let git rollback erase previously captured runtime truth
- do not make ad hoc agent HTTP calls the required primary source-of-truth mechanism for task state
