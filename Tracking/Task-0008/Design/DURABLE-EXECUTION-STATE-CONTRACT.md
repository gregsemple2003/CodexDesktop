# Durable Execution-State Contract

## Purpose

The future `Tasks` surface cannot supervise work honestly unless the runtime records more than `started` and `done`.

This note defines the first durable contract that backend implementation should honor.

This task is backend-only.

The first honest operator path can be Codex or direct backend interaction against this contract before any frontend wiring exists.

Chosen runtime shape:

- implement task dispatch as a separate backend-owned task-run workflow and API contract
- do not model task dispatch as another Git-tracked recurring job spec
- allow recurring jobs to create task runs later if needed, but keep task-run state as its own runtime concept

The contract exists so the system can answer, durably:

- what work was dispatched
- what the active contract is
- what the latest meaningful progress was
- whether silence is legitimate
- when a poke is allowed
- what interruption means
- what deeper working context belongs to the run
- which exclusive repo lane the run owns for execution
- which recorded commit that repo lane may reset to for proof or cleanup

## Task-0009 Consumer Obligations

The first implementation proof for this task can still be Codex or direct backend interaction.

That does not change the downstream consumer bar.

The future `Tasks` tab from [Task-0009](../../Task-0009/TASK.md) already assumes this task will expose enough backend truth to support:

- grouped task rows tied to the right active run
- a detail pane that can show:
  - current state
  - why that state is true
  - what changed recently
  - the next expected step
  - whether the state is fresh enough to trust
  - which actions are safe now
- direct `Open Thread` or equivalent working-context launch
- honest enable or disable behavior for:
  - `Dispatch`
  - `Poke`
  - `Interrupt`

This task does not need to return final UI copy.

It does need to expose enough machine-readable detail that the UI does not have to guess from raw phase names or weak status labels.

The backend contract therefore needs to preserve or compute enough durable truth for:

- one minimal task-level meaning field
- state classification
- reason inputs
- next-step inputs
- freshness and staleness inputs
- action-availability inputs
- deep-context launch metadata
- owned-checkout identity
- restore-commit identity and reset status
- one explicit human-facing state envelope
- meaningful-progress truth
- wait-contract truth
- story-run selection truth
- attention-priority truth

## Canonical State Split

This task keeps two different state classes on purpose:

- task-definition state
  - task-owned markdown artifacts such as `TASK.md`, `PLAN.md`, `HANDOFF.md`, and task evidence docs
- live execution state
  - backend-owned task-run runtime state

The live execution flow must be backed by durable Temporal state so the runtime does not forget a run's waiting, progress, interrupt, or cleanup state after:

- process restart
- quiet periods
- dashboard disconnect
- Codex or other client disconnect

The UI and any readback client are derivative consumers of that runtime truth.

They may summarize or map it for humans, but they do not own the canonical wait or progress state.

Local run artifacts such as logs, transcripts, JSONL event captures, and final-message files are supporting evidence, not the canonical record of:

- wait state
- progress freshness
- interrupt status
- repo-reset status

Repo-lane ownership and restore-commit state must also be persisted directly in the backend-owned runtime contract, not rediscovered later from ad hoc filesystem inspection.

## Declared Doc Truth And Runtime Ingest Model

These git-authored task docs are co-owned declared task/process truth:

- `TASK.md`
- `PLAN.md`
- `HANDOFF.md`
- `TASK-STATE.json`

At the current git revision, they remain the declared task-definition truth.

The declared parsing scope must be named explicitly with:

- `declared_worktree_root`
  - the checkout or worktree root the backend parsed
- `declared_task_root`
  - the specific task directory the backend read

The backend runtime does not replace them.

Instead, the backend or orchestrator must ingest or snapshot the relevant task-definition state into Temporal-backed run state when needed, especially:

- at dispatch
- during backend-owned reconcile when declared task docs changed materially
- when readback needs to explain divergence between current git docs and a run's captured task snapshot

Primary model:

- doc-to-backend ingest
- backend-owned reconcile

Preferred day-to-day model:

1. agents continue writing durable local task docs:
   - `TASK.md`
   - `PLAN.md`
   - `HANDOFF.md`
   - `TASK-STATE.json`
2. task APIs learn declared task truth by reading and parsing current task docs on demand within:
   - `declared_worktree_root`
   - `declared_task_root`
3. dispatch captures a task-definition snapshot into Temporal-backed run state
4. task read APIs perform read-through reconcile for active runs:
   - read current task docs from `declared_task_root` inside `declared_worktree_root`
   - read active Temporal run state
   - if current docs differ materially from the active run's last captured snapshot, record a new doc snapshot or reconcile event into Temporal-backed run state
5. Temporal owns only non-rewindable runtime facts
6. manual re-ingest or reconcile remains a repair or debugging fallback, not the normal workflow

Not required as the primary mechanism:

- ad hoc agent HTTP calls to keep task state truthful

A file watcher may be added later if implementation proves it useful, but it is not required by this contract.

The contract therefore requires:

- task APIs must parse current task docs on demand
- task APIs must know which `declared_worktree_root` and `declared_task_root` they are reconciling
- for active runs, task APIs must reconcile doc changes into Temporal-backed run snapshots
- file watching is optional, not the normal path
- direct agent HTTP calls are optional helpers at most, not the normal path

The durable invariant is:

- latest git docs remain the declared task-definition truth
- live runtime truth already recorded in Temporal-backed state is not forgotten if git later rewinds

## Human-Facing State Envelope

Every task-level or run-level human-facing readback must expose one explicit durable envelope with at least:

- `state`
- `reason_code`
- `state_summary`
- `evidence_refs`
- `next_owner`
- `next_expected_event`
- `suspicious_after`
- `action_block_reasons`

Field intent:

- `state`
  - the current human-visible state the backend stands behind
- `reason_code`
  - a stable machine-readable reason identifier
- `state_summary`
  - one concise sentence a human can read directly
- `evidence_refs`
  - durable references that justify the current state or summary
- `next_owner`
  - who owns the next move
- `next_expected_event`
  - what concrete event should happen next
- `suspicious_after`
  - when silence or lack of change becomes suspicious
- `action_block_reasons`
  - explicit block reasons keyed by action rather than hidden UI guesses

The UI may style or group this envelope.

It must not replace it with a weaker client-invented explanation layer.

These field names belong to the planned Task-0008 runtime or readback contract rather than the shared `.codex` `TASK-STATE.json` schema.

`state_summary` writing rule:

- short
- plain-language
- present-tense
- non-jargony
- understandable without opening artifacts first

`evidence_refs` must be a typed launchable structure:

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

Separate from the run or state envelope, task readback must expose one minimal task-level meaning field:

- `meaning_summary`

`meaning_summary` is:

- the primary task-level summary of why the task exists or why the human should care
- derived from declared task docs
- distinct from `state_summary`, which remains the current run or state summary

This field stays intentionally narrow so later clients can reduce markdown reconstruction burden without introducing a large family of overlapping task-level summary fields.

## Meaningful Progress

`Meaningful progress` is a durable change that moves the run's real contract forward.

It includes:

- a durable state transition
- creation, update, or clearance of a wait contract
- a successful or failed bounded tool step that changes what can happen next
- a material artifact mutation in the owned execution lane
- a persisted test result
- a git commit
- a reset to a recorded restore target
- an agent-authored checkpoint note that changes the next expected step, wait contract, or supervisor understanding

It does not include:

- raw stdout or stderr noise
- a heartbeat with no contract change
- repeated identical status writes
- filesystem touches with no contract meaning

This is the basis for `last_progress_at`, `last_progress_summary`, freshness, and sleeping detection.

Freshness policy must vary by state category:

- `running`
  - ages against expected progress cadence and becomes suspicious fastest
- `waiting_for_human`
  - ages against the stored wait stale deadline
- `blocked`
  - ages against the promised blocker-review cadence, not the running cadence
- terminal states
  - do not imply active progress expectations

## Repo Ownership Model For Simple Execution

For simple task execution in this task's scope, `own a repo entirely` is tightened to:

- the runtime must mutate only an exclusive backend-owned repository checkout or equivalent isolated repo lane
- that owned lane must not be the human's shared primary worktree
- destructive git cleanup operations are allowed only inside that owned lane

This can be implemented with a separate clone, a dedicated worktree, or another isolated local checkout model.

The important contract is ownership and safety, not one specific provisioning mechanism.

What does not count:

- reusing the human's current dirty checkout as the normal execution lane
- assuming cleanup is safe because the runtime promises to be careful later

## Restore-Commit Semantics

The backend must record at least one useful restore commit for an owned execution lane:

- `baseline_commit`
  - the commit captured before the run begins mutating the owned lane

It may also record later allowed restore targets when the runtime intentionally creates them, for example:

- `approved_restore_commit`
  - a later known-good commit that proof or cleanup may reset to instead of the original baseline

The contract must make explicit:

- which reset operations are allowed
- which commit each reset operation may target
- whether reset is currently pending, complete, or blocked

This is how unit proof and execution cleanup avoid hand-wavy `clean up afterward` language.

Later restore targets are allowed only when:

- a durable checkpoint event justifies them
- the approving actor is recorded
- the checkpoint is specific enough to count as a known-good restore point

If cleanup cannot reach the chosen restore target, the runtime must:

- report the failure durably
- preserve a fallback recorded restore commit when one exists
- mark cleanup as blocked rather than claiming success

## Dispatch Readiness Rule

A task is dispatch-ready only when all required durable preconditions are satisfied, at minimum:

- the current plan is approved
- no unresolved blockers prevent execution
- no nonterminal still-owning run already controls the task's live story unless the new dispatch is an explicit superseding action
- an accepted exclusive execution lane is acquired or reserved
- `declared_worktree_root` is resolved
- `declared_task_root` is resolved
- `baseline_commit` is recorded for that lane
- a task-definition snapshot is captured from the declared task docs
- the first expected event is recorded
- the first suspiciousness deadline is recorded

Dispatch readiness is a backend contract decision.

It is not merely a UI enablement choice.

## Core Principles

- Silence is not automatically healthy.
- A run is only allowed to be quiet when its durable state says why.
- `Waiting` must always name what the run is waiting on.
- A run without recent progress and without a valid wait contract is a supervision problem.
- The human should not have to infer provenance by cross-referencing transcripts manually.
- Every dispatched run should preserve enough context that Codex or a future client can open deeper context.

## Required Run Fields

Each durable task run should eventually record at least:

- `run_id`
- `task_id`
- `meaning_summary`
- `task_title_snapshot`
- `task_revision` or equivalent snapshot identity
- `declared_worktree_root`
- `declared_task_root`
- `captured_task_snapshot`
- `captured_task_revision`
- `declared_task_revision_at_dispatch`
- `dispatch_source`
  - manual
  - promoted_candidate
  - scheduled
  - automated_follow_on
- `status`
- `state_envelope`
- `owned_repo_root` or equivalent owned-lane identifier
- `repo_checkout_mode`
- `wait_reason`
- `blocking_reason`
- `interrupt_reason`
- `latest_contract_summary`
- `last_progress_at`
- `last_progress_summary`
- `last_supervisor_touch_at`
- `last_human_touch_at`
- `next_expected_step`
- `next_expected_step_owner`
- `expected_progress_by` or equivalent freshness deadline
- `expected_first_event`
- `expected_first_event_by`
- `dispatch_ready`
- `dispatch_block_reasons`
- `attention_level`
- `attention_reason`
- `attention_sort_key`
- `story_role`
- `supersede_reason`
- `superseded_by_actor`
- `superseded_run_disposition`
- `superseded_by_run_id`
- `superseded_at`
- `baseline_commit`
- `approved_restore_commit` when one exists
- `cleanup_restore_commit`
- `cleanup_state`
- `doc_runtime_divergence_status`
- `doc_runtime_divergence_summary`
- `session_id`
- `session_transcript_path`
- `launch_target`
- `created_at`
- `updated_at`

## Status Vocabulary

The first durable vocabulary should include:

- `queued`
- `dispatching`
- `running`
- `waiting_for_human`
- `blocked`
- `sleeping_or_stalled`
- `interrupted`
- `completed`
- `failed`

### Notes

- `running` means the run is actively progressing or recently progressed.
- `waiting_for_human` means the silence is justified by a named human dependency.
- `blocked` means the run cannot proceed because an external condition is unmet and that condition is named durably.
- `sleeping_or_stalled` means the system expected progress but did not get it, and there is no valid wait contract that excuses the silence.
- `completed` means no further human judgment, review, or approval is required for the task or run to count as done in the human-facing system.

If review or approval is still required, the state must not present as `completed`.

It must instead present as:

- `waiting_for_human`
- with a review or approval reason code
- and `human_action_target.kind = approval_action` or an equivalent explicit approval target

## Wait Contract

If a run wants to stop making progress without being treated as sleeping, it must store a durable wait contract with:

- `wait_reason`
- `wait_started_at`
- `waiting_on`
- `blocked_why`
- `resume_condition`
- `resume_signal`
- `resume_owner`
- `human_action_required`
- `human_action_target`
- `wait_stale_after`
- `poke_policy`
- `evidence_refs`

`human_action_target` must be typed:

- `kind`
  - `task_artifact`
  - `run_artifact`
  - `transcript`
  - `api_resource`
  - `owned_checkout`
  - `approval_action`
- `label`
- `uri`
- optional `line_range`

These wait-contract fields are not only for supervision logic.

They are also the durable inputs the later `Tasks` surface will use to explain:

- what the task is waiting for
- why progress cannot continue
- who owns the next move
- where the human should act when human action is required
- when the wait itself becomes stale
- why `Poke` is or is not allowed

`waiting_for_human` is only valid when:

- `human_action_required = true`
- `human_action_target` is explicit enough to open or inspect
- `wait_stale_after` is recorded
- the human can tell why their action is needed

Examples of valid wait reasons:

- waiting_for_human_answer
- waiting_for_tool_result
- waiting_for_external_runtime
- waiting_for_approval

Invalid wait behavior:

- a silent run with no durable reason
- a run that says `waiting` but never says on what
- a run that says `waiting_for_human` but cannot say where the human should act
- a run that stops after promising a next step without updating its state

## Sleeping Or Stalled Detection

A run should be eligible for `sleeping_or_stalled` when all are true:

- it is not `completed`, `failed`, or `interrupted`
- it does not have a valid active wait contract
- its last meaningful progress is older than the allowed freshness window for its current contract

The exact time thresholds can evolve later.

The important product rule is structural:

- the system should decide from durable contract and progress evidence
- not from vibes

## Attention Priority Model

The backend should expose explicit attention-priority fields so later clients do not invent fragile urgency heuristics.

The minimum durable shape is:

- `attention_level`
- `attention_reason`
- `attention_sort_key`

Suggested `attention_level` values:

- `none`
- `watch`
- `needs_attention`
- `urgent`

The UI may group or style these values, but the backend remains the source of truth for why a task belongs in a `Needs Attention` region and how it should sort there.

Rough precedence should be:

- `urgent`
  - sleeping or stalled runs
  - waits that have gone stale
  - cleanup or reset failures
- `needs_attention`
  - waiting-for-human with an actionable target
  - blocked states needing review
  - dispatch-ready tasks intentionally surfaced for action
- `watch`
  - healthy running work
  - non-stale blocked work
- `none`
  - terminal history and non-competing background states

## Poke Rule

`Poke` is allowed when:

- the run is `sleeping_or_stalled`
- or the run is `running` but exceeded its promised progress freshness window
- and there is no valid wait contract that says silence is expected

`Poke` should record durably:

- who initiated the poke
- why it was allowed
- when it happened
- what the expected next progress signal is

Backend readback should also make the current poke eligibility obvious, either by returning an action-availability view or by returning enough durable fields that the caller can derive the same answer without heuristics.

## Interrupt Rule

Interrupt must be a first-class state transition, not a hidden side effect.

Interrupt should record durably:

- who interrupted
- why
- what previous state was interrupted
- whether the run is resumable
- what follow-on state the task should show

Backend readback should expose whether interrupt is currently allowed and, if not, why not.

If interruption or cleanup requires a repo reset, backend readback should also expose:

- which commit the owned lane will reset to
- whether that reset succeeded
- or the exact reason reset is blocked

Candidate interrupt outcomes:

- interrupted_pending_review
- interrupted_and_requeueable
- interrupted_terminal

Those can stay collapsed under `interrupted` in the first visible UI if the backend still stores the subtype or reason durably.

## Thread And Session Provenance

Each run should preserve enough deep-context information that Codex or a future client can open the right context directly.

The first contract should preserve:

- the primary session id if one exists
- the local transcript path if one exists
- the preferred launch target
  - transcript file
  - VSCodium task folder
  - VSCodium thread/session workspace
  - future richer deep-link type if implementation proves it
- the launch command or launch parameters needed to open the context
- enough repo-lane context that proof and debugging can identify the owned checkout without confusing it with the human's primary worktree

The product goal is:

- one backend readback path to useful deeper context

Not:

- a breadcrumb trail the human must manually decode

## Task Readback Shape

Task-0009's grouped task view means the backend cannot stop at run-only truth.

The backend readback also needs to answer, for a task:

- what the task means in one narrow human-readable summary
- whether there is an active or most relevant run
- whether that run currently owns the live human-facing story
- what durable state the task should currently show
- why that state is true
- what the next expected step is
- when the current view becomes stale enough to distrust
- which actions are valid right now
- which owned checkout is associated with the active run
- which recorded commit cleanup or proof may restore it to
- what attention priority the task currently has
- what declared worktree root is being read
- what declared task root is being read
- what declared task revision is current in git
- what task snapshot revision the active run captured
- whether those two differ materially

## Git Rollback And Runtime Divergence Model

Current git docs and live runtime state have different jobs:

- current git docs remain the declared task-definition truth
- Temporal-backed run state remains the durable record of live execution truth and side effects already observed

If git later rewinds or changes materially, the system must preserve:

- the run's captured task-definition snapshot
- recorded wait state
- interrupts
- cleanup/reset state
- other runtime side effects already persisted

Task readback must represent mismatch explicitly, at minimum through:

- `declared_worktree_root`
- `declared_task_root`
- current declared task revision
- captured task snapshot revision
- `doc_runtime_divergence_status`
- `doc_runtime_divergence_summary`

This exists so rollback does not make the system forget prior runtime truth while still allowing latest git to remain the declared task definition.

Read-through reconcile is scoped to the declared worktree and declared task root, not merely to a coarse repo family.

Read-through reconcile is the normal way that declared-doc changes flow into active runtime context.

Manual re-ingest exists for repair and debugging only.

## Multiple Runs Per Task

At most one run may own the current live human-facing story for a task at a time.

Selection rules:

- an active nonterminal run that has not been superseded owns the live story
- a retry or superseding run must explicitly displace the older run
- superseded runs remain queryable history but stop owning the present-tense story
- a terminal run may remain visible as recent evidence, but it does not keep owning the task once task-level readback has moved back to `ready`, `blocked`, `waiting_for_human`, or another non-run state
- when no run currently owns the live story, task readback must say so explicitly rather than pretending an older run is still current

Any superseding dispatch must also record:

- `supersede_reason`
- `superseded_by_actor`
- `superseded_run_disposition`
  - `left_in_history`
  - `interrupted`
  - `marked_superseded`
  - `cleanup_blocked`

This is how later clients avoid guessing from timestamps alone.

This can be implemented as task-level readback, run-level readback, or a thin composition of both.

What does not count is forcing the UI to rediscover the live task story by stitching together weakly related task markdown, raw run ids, and local transcript search.

## Contract Must Never Allow

The contract must never allow:

- `state_summary` that only repeats backend jargon or raw phase labels
- `evidence_refs` as untyped free-form text
- `waiting_for_human` with no actionable target
- one global freshness timeout applied equally to running, blocked, and waiting states
- superseding a run without recording why and what happened to the older run
- promoting a later restore target without a durable known-good checkpoint
- claiming cleanup succeeded when the owned checkout did not actually reach a recorded restore commit
- git rollback to silently erase recorded wait, interrupt, cleanup, or captured task-snapshot truth
- current git inspection alone to replace the run's captured task-definition snapshot
- ad hoc agent HTTP calls as the required primary source-of-truth mechanism for task state
- task readback to hide that current git and the run snapshot differ

## Future UI Implications

This contract should let the future `Tasks` tab say, simply:

- `Ready`
- `Running`
- `Waiting on you`
- `Blocked`
- `Sleeping`
- `Interrupted`
- `Done`
- `Failed`

while still keeping the backend detail needed for honest control actions.

It should also let the future `Tasks` tab explain, truthfully:

- why the task is in that state
- when the last meaningful progress happened
- what signal should happen next
- whether the state is still fresh enough to trust
- why an action is enabled or disabled

For operators and proof lanes, it should also let the backend or direct caller explain:

- whether execution happened in an exclusive owned repo lane
- what baseline commit anchored that lane
- what commit the runtime restored during cleanup

## Promotion Boundary

If this contract proves stable and becomes useful across repos, promote the durable rule into shared `.codex` orchestration docs.

Until then, this note is the task-local baseline for implementation.
