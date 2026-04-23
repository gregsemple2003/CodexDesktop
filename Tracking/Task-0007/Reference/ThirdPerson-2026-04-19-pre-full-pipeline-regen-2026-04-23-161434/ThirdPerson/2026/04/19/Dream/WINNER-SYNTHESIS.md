# Dream Winner Synthesis (ThirdPerson 2026-04-19)

This pass-`2B` artifact was refreshed during pass `2C` winner-task audit.

It still freezes the smallest honest winner set from [SOLUTION-DESIGN.md](./SOLUTION-DESIGN.md), but it now also records the task-audit corrections that were required before the winner set could be treated as closeout-ready.

## Pass-2B Audit Outcome

- `Verdict`: `ready`
- `Problem Rows Reviewed`: `P-001`, `P-002`, `P-003`, `P-004`, `P-005`, `P-006`
- `Blocking Questions`:
  - `P-002`: what exact `TASK-STATE` fields and `ORCHESTRATION` gate rules create an owner or wait or STOP-latched state instead of capture-only bookkeeping; resolved by the `2A` rewrite to exact shared lifecycle and state homes
  - `P-003`: which exact shared durable file owns the direct-answer-first rule; resolved by narrowing the home to shared `AGENTS.md`
  - `P-004`: what exact approval-packet fields are required at the human gate, and where the requirement is enforced on the outgoing approval ask; resolved by narrowing to the outgoing approval packet and adding `pass-history effect`
- `Strengthening Questions`:
  - `P-001`: whether full-view visual evidence should be required only for disputed visual claims or for all visual regression claims
  - `P-006`: whether repo-local truths should be stored as copied prose or as canonical doc refs plus a short task-local reminder

## Pass-2C Task Audit Outcome

- `Verdict`: `ready`
- `Tasks Reviewed`:
  - [Option-Tasks/PROBLEM-0001-OPTION-A.md](./Option-Tasks/PROBLEM-0001-OPTION-A.md)
  - [Option-Tasks/PROBLEM-0002-OPTION-A.md](./Option-Tasks/PROBLEM-0002-OPTION-A.md)
  - [Option-Tasks/PROBLEM-0003-OPTION-A.md](./Option-Tasks/PROBLEM-0003-OPTION-A.md)
  - [Option-Tasks/PROBLEM-0004-OPTION-A.md](./Option-Tasks/PROBLEM-0004-OPTION-A.md)
  - [Option-Tasks/PROBLEM-0005-OPTION-A.md](./Option-Tasks/PROBLEM-0005-OPTION-A.md)
  - [Option-Tasks/PROBLEM-0006-OPTION-A.md](./Option-Tasks/PROBLEM-0006-OPTION-A.md)
- `Primary Findings`:
  - All six tasks are already in `TASK-CREATE.md` burden-reduction (or `consensus`) shape with explicit boundaries, homes, and falsifiers; the audit does not depend on hidden thread context.
  - Three concrete weak-closure / ambiguity hazards were corrected:
    - `P-001` proof-plan fixtures now have an explicit durable home (task-owned `Tracking/Task-<id>/Testing/`), rather than an implementer-chosen scratch location.
    - `P-002` now avoids "two truths" about `status` by naming `execution.run_state` (not `execution.status`) and by defining explicit consistency rules with existing task-level `status/phase/current_gate`.
    - `P-004` proof-plan template work now points explicitly to the shared exemplars (`Exemplars/PLAN.md` and `Exemplars/HANDOFF.md`) and removes the "pick any template home" escape hatch.
- `Blocking Clarifying Questions`:
  - none
- `Strengthening Clarifying Questions`:
  - `P-001`: whether the repo should later require full-surface evidence for all visual claims or only disputed ones
  - `P-003`: whether multipart direct questions need one carve-out or several narrower carve-outs once the consensus task is run
- `Required Rewrites`:
  - none remaining after the on-disk corrections recorded above
- `Optional Strengthening`:
  - `P-004`: name the exact canonical header(s) that the approval packet must use in `PLAN.md` and `HANDOFF.md`, so cold review can find the packet without guessing.
  - `P-001`/`P-005`: name the exact shared proof fields (in `Processes/TESTING.md` and `Processes/DEBUGGING.md`) that are considered mandatory for a claim to be "auditable", so repo-local adapters can reference them directly.
- `Relevant Rule Or Exemplar`:
  - `C:\Users\gregs\.codex\Orchestration\Processes\TASK-AUDIT.md`
  - `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md`
  - `C:\Users\gregs\.codex\Orchestration\Exemplars\PLAN.md`
  - `C:\Users\gregs\.codex\Orchestration\Exemplars\HANDOFF.md`
- `What Must Change Before Enqueue`:
  - enqueue from the audited task set only
  - treat any older pass-2C notes that claim large rewrites were still needed as superseded by the current on-disk tasks

## Winner Set

### W-001 ThirdPerson Regression Claim Gate

- `Winner Label`: `W-001 ThirdPerson Regression Claim Gate`
- `Source Problem Refs`: `P-001`
- `Source Option Refs`: `P-001 / Option A`, with the generic proof-quality seam from `Option B` merged in
- `Final Boundary`: auditable ThirdPerson regression claims must cite repo-local lane truth from `REGRESSION.md` and a task-owned `REGRESSION-RUN` that names claimed lane, actual flow, why it counts, why the evidence covers the disputed surface, and which disqualifiers keep a run from proving more
- `Why This Is The Right First Intervention Boundary`: the burden happens when supporting or invalid evidence is turned into a default-lane pass claim; this is the first fail-closed point
- `Audited Task File`: [Option-Tasks/PROBLEM-0001-OPTION-A.md](./Option-Tasks/PROBLEM-0001-OPTION-A.md)
- `Pass-2C Audit Corrections`: tightened the proof plan to require "bad/good" `REGRESSION-RUN` fixtures under the implementing task's own `Tracking/Task-<id>/Testing/` home (no implementer-chosen scratch location)
- `What Reopens The Choice`: multiple repos with explicit repo-local regression docs still suffer the same claim-shape failure

### W-002 Shared Execution-State Pause/Ownership Gate

- `Winner Label`: `W-002 Shared Execution-State Pause/Ownership Gate`
- `Source Problem Refs`: `P-002`
- `Source Option Refs`: `P-002 / Option A` after audit rewrite
- `Final Boundary`: shared lifecycle state must encode owner, execution run-state, continuation latch, and resume condition so stop, idle, and approval-gated transitions are explicit durable control flow
- `Why This Is The Right First Intervention Boundary`: wake-up supervision and ignored STOPs both occurred at execution-state boundaries, not at repo-specific gameplay seams
- `Audited Task File`: [Option-Tasks/PROBLEM-0002-OPTION-A.md](./Option-Tasks/PROBLEM-0002-OPTION-A.md)
- `Pass-2C Audit Corrections`: renamed `execution.status` to `execution.run_state` and added explicit consistency rules with existing task-level `status/phase/current_gate`, so the task cannot close with two conflicting "pause/ownership" truths
- `What Reopens The Choice`: prompt or workflow enforcement alone eliminates wake-up and STOP failures without new durable state

### W-003 Shared Direct-Answer-First Rule

- `Winner Label`: `W-003 Shared Direct-Answer-First Rule`
- `Source Problem Refs`: `P-003`
- `Source Option Refs`: `P-003 / Option A` after exact-home narrowing
- `Final Boundary`: explicit yes or no, agree or disagree, short-answer, and `what is still wrong` questions should receive a direct first-sentence answer before framing, but the exact shared carve-outs still require a bounded decision artifact before enqueue-ready implementation
- `Why This Is The Right First Intervention Boundary`: the burden appears before repo content matters; the failure is answer shape itself
- `Audited Task File`: [Option-Tasks/PROBLEM-0003-OPTION-A.md](./Option-Tasks/PROBLEM-0003-OPTION-A.md)
- `Pass-2C Audit Corrections`: none required in this audit beyond confirming the writeup type is honestly `consensus` and the decision output is durable
- `What Reopens The Choice`: the failure proves tied to one repo-specific wrapper rather than shared interaction behavior

### W-004 Approval-Gate Packet Contract

- `Winner Label`: `W-004 Approval-Gate Packet Contract`
- `Source Problem Refs`: `P-004`
- `Source Option Refs`: `P-004 / Option A` after boundary narrowing
- `Final Boundary`: every human approval ask for plan, pass, or reopen changes must point to a task-owned `Approval Packet` that includes what changed, which artifacts changed, the pass-history effect, and review links that make approval possible without reconstruction
- `Why This Is The Right First Intervention Boundary`: the burden is at the approval request itself; the human should not need to open files just to infer what they are approving
- `Audited Task File`: [Option-Tasks/PROBLEM-0004-OPTION-A.md](./Option-Tasks/PROBLEM-0004-OPTION-A.md)
- `Pass-2C Audit Corrections`: tightened the proof plan to use the shared exemplars (`Exemplars/PLAN.md` and `Exemplars/HANDOFF.md`) as the canonical template home (no "pick any template home" escape hatch)
- `What Reopens The Choice`: a shared approval contract proves too blunt and local templates solve the burden without recurrence

### W-005 Repo-Local First-Disagreement Debugging Gate

- `Winner Label`: `W-005 Repo-Local First-Disagreement Debugging Gate`
- `Source Problem Refs`: `P-005`
- `Source Option Refs`: `P-005 / Option A`
- `Final Boundary`: before another code change or closure claim on a reopened or surviving human-default-lane runtime defect, `BUG-<NNNN>.md` and `REGRESSION-RUN-<NNNN>.md` must preserve the first concrete disagreement with values and the next upstream writer boundary
- `Why This Is The Right First Intervention Boundary`: the transcript explicitly turns the problem from tweak mode into first-disagreement narrowing plus honest bug and regression artifacts
- `Audited Task File`: [Option-Tasks/PROBLEM-0005-OPTION-A.md](./Option-Tasks/PROBLEM-0005-OPTION-A.md)
- `Pass-2C Audit Corrections`: none required in this audit beyond confirming the repo-local enforcement boundary and required bug/rerun fields are explicit
- `What Reopens The Choice`: repeated packets yield stable, trustworthy numeric seams that justify a stronger automation-owned checker path

### W-006 Task-Local Active-Constraints Ledger

- `Winner Label`: `W-006 Task-Local Active-Constraints Ledger`
- `Source Problem Refs`: `P-006`
- `Source Option Refs`: `P-006 / Option A` after truth-ref narrowing
- `Final Boundary`: before acting after a new hard constraint, approval gate, or repo-truth correction, task-owned current-state artifacts must record the active constraint and point to canonical repo truth instead of copying policy text
- `Why This Is The Right First Intervention Boundary`: the burden arrived mid-task, not just at startup, and repeatedly required immediate durable capture before more work
- `Audited Task File`: [Option-Tasks/PROBLEM-0006-OPTION-A.md](./Option-Tasks/PROBLEM-0006-OPTION-A.md)
- `Pass-2C Audit Corrections`: none required in this audit beyond confirming the ledger entry shape, status semantics, and repo-truth reference rule are explicit
- `What Reopens The Choice`: many tasks need the same active-constraints shape and shared task state wants to standardize it

## Final Winner Set And Rollout Order

1. `W-001 ThirdPerson Regression Claim Gate`
2. `W-005 Repo-Local First-Disagreement Debugging Gate`
3. `W-006 Task-Local Active-Constraints Ledger`
4. `W-004 Approval-Gate Packet Contract`
5. `W-002 Shared Execution-State Pause/Ownership Gate`
6. `W-003 Shared Direct-Answer-First Rule`

This remains the smallest honest winner set after both winner synthesis and winner-task audit.
