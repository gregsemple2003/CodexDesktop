# Dream Solution Design (ThirdPerson 2026-04-19)

This pass-`2A` artifact converts the pass-`1` burden inventory into an explicit problem set and option set.

It was regenerated under the hard antagonistic solution-generation contract:

- `2A` is the solution-writing lane
- `2B` is a separate attacking audit lane
- the writer leaves behind durable design intent that must survive a hostile reread

Repo-local constraints were taken from the packet repo, not the current workspace repo:

- [AGENTS.md](../../../../../../../../../Agent/ThirdPerson/AGENTS.md)
- [REGRESSION.md](../../../../../../../../../Agent/ThirdPerson/REGRESSION.md)
- [TESTING.md](../../../../../../../../../Agent/ThirdPerson/TESTING.md)

Under the current merge discipline, every major burden driver remains a `separate_problem_row`.

## Burden-To-Problem Mapping

| Burden Driver | Mapping | Problem Ref | Why Not Merged |
| --- | --- | --- | --- |
| `BD-001 Default-Lane Proof Discipline (And Evidence Honesty)` | `separate_problem_row` | `P-001` | Lane identity and evidence validity share one enforcement boundary: any claimed human-default-lane regression closeout. |
| `BD-002 Persistent Ownership, Continuity, And Pause-Gate Adherence` | `separate_problem_row` | `P-002` | Wake-up supervision and ignored STOP or approval gates share one execution-state boundary. |
| `BD-003 Direct Answer Discipline` | `separate_problem_row` | `P-003` | This is a conversational answer-shape failure, not a proof, debugging, or task-state failure. |
| `BD-004 Approval-Ready Review Surface` | `separate_problem_row` | `P-004` | Human approval friction sits at a distinct review-surface boundary: changed artifacts, pass framing, and links must be legible before approval is requested. |
| `BD-005 Root-Cause Debugging Discipline (And Defect-Tracking Honesty)` | `separate_problem_row` | `P-005` | The burden is a forced mode switch from tweak mode to value-traced root-cause debugging. |
| `BD-006 Durable Constraint Retention (And Repo-Local Truth Retention)` | `separate_problem_row` | `P-006` | Task-specific hard constraints and repo-local truth retention need a durable capture-and-consult mechanism. |

No burden drivers were dropped.

## P-001 Default-Lane Proof And Evidence Gate

### Problem Boundary

Prevent ThirdPerson closeout from turning non-default, supporting-only, cropped, or otherwise invalid evidence into a human-default-lane regression claim.

### Source Burden Drivers

- `BD-001 Default-Lane Proof Discipline (And Evidence Honesty)`

### Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3325",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3649",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3431",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3473",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3618",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3702",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3814",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4202",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4114",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4432",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4909",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6451",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5040",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5050",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5100"
]
```

### Why This Is One Problem

Default-lane substitution, supporting-lane overclaiming, and cropped or invalid visuals all fail at the same moment: when the system turns evidence into a human-default-lane claim.

### Rejected Nearby Splits Or Merges

- Do not split lane identity from evidence validity. The packet shows both must fail closed together.
- Do not merge with `P-005`. Debugging can still be honest while closeout proof is invalid.

### Option A: Repo-Local Regression Claim Gate With Evidence Disqualifiers

- `Option Label`: `Option A`
- `Summary`: Make ThirdPerson closeout fail closed unless a claimed regression run names the repo lane, cites the required case ids, links full-view runtime evidence for the disputed fact, and labels any non-default lane as supporting-only.
- `Writeup Type`: `concrete implementation`
- `Implementation Home`: repo-local [REGRESSION.md](../../../../../../../../../Agent/ThirdPerson/REGRESSION.md), shared `Processes/TESTING.md`, and task-owned `Tracking/Task-<id>/Testing/REGRESSION-RUN-<NNNN>.md`.
- `Implementation Home Rationale`: Shared docs can own generic proof-quality claim-shape fields, but the decisive lane truth and disqualifiers remain repo-local.
- `Enforcement Boundary`: any auditable ThirdPerson regression claim or closure citation that says `TP-REG-001`, `TP-REG-002`, or `TP-REG-003` passed.
- `Acceptance Test`: A future ThirdPerson task cannot honestly claim regression passing with only lane `C`, proof-view, or cropped evidence; the artifact itself blocks that overclaim.
- `Falsifier`: The system still closes a task while the only proof is non-default or supporting-only evidence, or a disputed visual claim survives after the cited screenshot is invalidated.
- `Edge Cases Or Failure Modes`: If the repo default lane changes, the gate must follow `REGRESSION.md` rather than hard-coded map names.
- `What This Option Avoids`: Generic UE wording in shared docs, operator-lane substitution, and "artifact exists therefore lane passed" reasoning.
- `Why This Wins`: The burden was not lack of lane doctrine; it was lack of a fail-closed repo-local proof gate that still absorbs generic shared proof-quality claim fields.
- `Scope Level`: `repo_local`
- `Why Not Narrower Scope`: A single task-local checklist would not protect the next ThirdPerson task from the same substitute-lane failure.
- `Why Not Broader Scope`: Durable packet evidence shows the user rejected generic shared UE lane wording in orchestration docs.
- `Promotion Or Demotion Trigger`: Promote only if other repos show the same failure shape after their own repo-local regression contracts are already explicit.
- `Reversal Path`: Demote to a task-local checklist if ThirdPerson later stops needing repo-specific lane exceptions.

### Option B: Shared Proof-Claim Contract In Shared Testing Docs

- `Option Label`: `Option B`
- `Summary`: Put the generic proof-quality claim shape into shared testing guidance so every repo declares claimed lane, evidence surface, and disqualifiers, while repo-specific lane truth stays local.
- `Writeup Type`: `consensus`
- `Implementation Home`: shared testing process docs.
- `Implementation Home Rationale`: One shared contract would reduce reinvention across repos.
- `Enforcement Boundary`: any repo-level regression claim across the orchestration system.
- `Acceptance Test`: Multiple repos adopt the same proof-quality claim fields without needing repo-specific rewrites.
- `Falsifier`: The shared rule either drifts into ThirdPerson-specific language again or remains too vague to block substitute-lane claims in practice.
- `Edge Cases Or Failure Modes`: A broad shared rule may overfit Unreal specifics or under-specify repo-local lane truth.
- `What This Option Avoids`: Per-repo duplicated prose.
- `Why This Option Is Not The Winner Yet`: The packet contains explicit durable evidence that the human objected to putting this repo's lane semantics into generic orchestration text, so this option is only the generic seam.
- `Scope Level`: `shared`
- `Why Not Narrower Scope`: Lane-provenance failures recur outside one task or repo.
- `Why Not Broader Scope`: Shared docs still should not own repo-specific default-lane semantics.
- `Promotion Or Demotion Trigger`: Keep live if several repos need the same claim fields after their own repo-local docs are explicit.
- `Reversal Path`: Replace if a stronger shared testing contract subsumes this comparison cleanly.

## P-002 Ownership Continuity And Pause-Latch Control

### Problem Boundary

Prevent wake-up supervision, ambiguous stop states, and ignored STOP or approval gates by treating ownership and pause state as explicit durable control flow.

### Source Burden Drivers

- `BD-002 Persistent Ownership, Continuity, And Pause-Gate Adherence`

### Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3447",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3692",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3929",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4015",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4590",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4919",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5108",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6558",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5272",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5300",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5310",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7038"
]
```

### Why This Is One Problem

Wake-up supervision, dropped ownership, and ignored STOP or approval gates all come from the same missing execution-state contract around who owns progress and when action is latched off.

### Rejected Nearby Splits Or Merges

- Do not merge with `P-006`. Both touch state, but `P-002` is about stop or resume ownership semantics, while `P-006` is about remembering specific constraints and truths.

### Option A: Shared Execution-State Gate With Pause Latch

- `Option Label`: `Option A`
- `Summary`: Treat lifecycle state as a real execution-state gate: stop, idle, approval, and resume transitions must encode owner, wait or block reason, and STOP or approval latch state before continuation is allowed.
- `Writeup Type`: `concrete implementation`
- `Implementation Home`: shared `ORCHESTRATION.md`, shared `TASK-STATE.md`, shared `TASK-STATE.schema.json`, and task-owned `TASK-STATE.json` plus `HANDOFF.md`.
- `Implementation Home Rationale`: Ownership and gating semantics need both durable state fields and lifecycle rules that treat those fields as binding control flow rather than passive bookkeeping.
- `Enforcement Boundary`: any stop, idle period, approval gate, STOP instruction, or resume transition during a live task pass.
- `Acceptance Test`: After an idle gap or STOP, durable state already shows owner, waiting or block reason, and whether continuation is latched off, and later sessions can tell whether work may proceed.
- `Falsifier`: The human still needs to send a wake-up or ownership reminder, or a STOP or approval gate is ignored in practice.
- `Edge Cases Or Failure Modes`: Supporting agents and leader-owned tasks need consistent state ownership; long-running external tools still need a durable blocked or waiting state rather than silence.
- `What This Option Avoids`: Ambiguous silence, chat-only pause semantics, and ownership drifting back to the human after a failed increment.
- `Why This Wins`: The largest quantified stall-loss events were wake-up supervision events, and the correct home is a shared execution-state gate that combines lifecycle rules with durable state.
- `Scope Level`: `shared`
- `Why Not Narrower Scope`: A ThirdPerson-only fix would not address the same continuity failure in other repos.
- `Why Not Broader Scope`: The fix belongs in task lifecycle state, not in a new global memory system.
- `Promotion Or Demotion Trigger`: Demote only if shared lifecycle enforcement is abandoned and the work is forced back into local prose.
- `Reversal Path`: Keep only repo-local handoff guidance if the shared task-state contract proves too heavy.

### Option B: Shared Workflow Enforcement Without Schema Expansion

- `Option Label`: `Option B`
- `Summary`: Enforce explicit owner, wait, and STOP behavior only through shared workflow rules and prompt discipline, without adding new machine-checkable state fields.
- `Writeup Type`: `consensus`
- `Implementation Home`: shared `ORCHESTRATION.md` and shared prompt rules.
- `Implementation Home Rationale`: This is the strongest cheaper rival because it acts directly on the worker without requiring schema expansion.
- `Enforcement Boundary`: shared stop, wait, approval, and resume behavior.
- `Acceptance Test`: Wake-up supervision and ignored STOPs disappear through workflow and prompt discipline alone, even though task state remains structurally unchanged.
- `Falsifier`: Later sessions and tools still cannot tell whether continuation is allowed, or STOP drift returns because no durable state captures it.
- `Edge Cases Or Failure Modes`: A wording-only gate can look strong in the moment while still leaving later sessions unable to inspect whether work should proceed.
- `What This Option Avoids`: Shared schema churn.
- `Why This Option Is Not The Winner Yet`: It is a serious rival, but it still leaves too much ambiguity for later sessions and tools.
- `Scope Level`: `shared`
- `Why Not Narrower Scope`: The burden is lifecycle-wide rather than task-local.
- `Why Not Broader Scope`: This option intentionally avoids new state fields or background infrastructure.
- `Promotion Or Demotion Trigger`: Promote only if workflow and prompt enforcement alone eliminate wake-up and STOP failures.
- `Reversal Path`: Drop if durable state remains necessary to preserve inspectable ownership truth.

## P-003 Direct-Answer-First Contract

### Problem Boundary

Force direct questions to receive a direct first-sentence answer in the requested shape before framing, narration, or supporting analysis.

### Source Burden Drivers

- `BD-003 Direct Answer Discipline`

### Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3392",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5001",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5010",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5020",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5030",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5073",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5090",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5263",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5281",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5290"
]
```

### Why This Is One Problem

All cited events show the same failure mode: the answer shape itself was wrong before any repo-specific content mattered.

### Rejected Nearby Splits Or Merges

- Do not merge with `P-004`. A review packet can be perfect and still evade a direct yes or no or short-answer question.

### Option A: Shared Answer-First Rule

- `Option Label`: `Option A`
- `Summary`: Add a shared interaction rule that explicit yes or no, agree or disagree, short-answer, and `what is still wrong?` questions must be answered in the first sentence before any framing or analysis.
- `Writeup Type`: `consensus`
- `Implementation Home`: shared [AGENTS.md](../../../../../../../../AGENTS.md).
- `Implementation Home Rationale`: The burden is conversational and cross-repo, and `AGENTS.md` is the narrowest exact shared durable home already governing interaction behavior.
- `Enforcement Boundary`: any explicit question whose requested shape is short or direct.
- `Acceptance Test`: Future transcripts no longer require restated prompts like `STOP. I asked you a question` or `stop evading the question`.
- `Falsifier`: The user still has to repeat the question because the first sentence does not answer it directly.
- `Edge Cases Or Failure Modes`: Uncertainty must be answered directly as uncertainty, not buried after framing.
- `What This Option Avoids`: Process-first narration and semantic drift away from the asked question.
- `Why This Wins`: The packet evidence is general conversational burden, not repo-local tooling burden.
- `Scope Level`: `shared`
- `Why Not Narrower Scope`: A repo-local clause would incorrectly localize a generic failure mode.
- `Why Not Broader Scope`: This is a prompt or behavior contract, not a UI or platform rewrite.
- `Promotion Or Demotion Trigger`: Demote to repo-local guidance only if the failure can be shown to come from a specific repo workflow wrapper rather than base interaction behavior.
- `Reversal Path`: Remove from shared guidance and keep only task-specific reminders if false positives prove too costly.

### Option B: Repo-Local Direct-Answer Clause In ThirdPerson Docs

- `Option Label`: `Option B`
- `Summary`: Add a ThirdPerson-specific instruction that runtime and debugging questions must be answered directly first.
- `Writeup Type`: `consensus`
- `Implementation Home`: repo-local [AGENTS.md](../../../../../../../../../Agent/ThirdPerson/AGENTS.md).
- `Implementation Home Rationale`: It is the narrowest durable home if the goal is zero shared-doc churn.
- `Enforcement Boundary`: ThirdPerson-only operator conversations.
- `Acceptance Test`: ThirdPerson sessions improve even if the shared agent contract stays unchanged.
- `Falsifier`: The same answer-shape failure still appears in other repos or in non-ThirdPerson work.
- `Edge Cases Or Failure Modes`: It duplicates a general interaction rule in a repo-local home and may still miss non-ThirdPerson recurrence.
- `What This Option Avoids`: Shared-doc edits.
- `Why This Option Is Not The Winner Yet`: The packet gives no durable reason to believe this burden is specific to ThirdPerson.
- `Scope Level`: `repo_local`
- `Why Not Narrower Scope`: A single task-local reminder would not address repeated repo sessions.
- `Why Not Broader Scope`: This option intentionally tests a narrow home against the shared rule.
- `Promotion Or Demotion Trigger`: Promote only if several tasks in the repo benefit while shared guidance remains untouched.
- `Reversal Path`: Drop if cross-repo recurrence appears immediately.

## P-004 Approval-Gate Review Packet

### Problem Boundary

Make human approval requests reviewable without forcing the human to reconstruct diffs, pass framing, or changed files manually.

### Source Burden Drivers

- `BD-004 Approval-Ready Review Surface`

### Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3494",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4379",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4399",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4421",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4755",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4512",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4522",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4571",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4581"
]
```

### Why This Is One Problem

Missing diffs, links without context, and invisible pass-structure changes all burden the same surface: the human approval request itself.

### Rejected Nearby Splits Or Merges

- Do not split links from pass framing. The packet shows the human needed both context and location together.
- Do not merge with `P-006`. Remembered constraints still need a human-usable approval packet.

### Option A: Shared Approval-Gate Packet Contract

- `Option Label`: `Option A`
- `Summary`: Require every outgoing approval request to carry a minimal packet: what changed, which artifacts changed, the pass-history effect, and contextual links that make review possible without reconstruction.
- `Writeup Type`: `concrete implementation`
- `Implementation Home`: shared `ORCHESTRATION.md` plus task-owned `PLAN.md` and `HANDOFF.md`.
- `Implementation Home Rationale`: Approval gates are lifecycle-wide, but the outgoing review packet still points into task-owned artifacts.
- `Enforcement Boundary`: any outgoing human approval request for plan, pass, or reopen changes.
- `Acceptance Test`: The human can approve or reject without asking for links, diff context, or whether reopened work changed pass history.
- `Falsifier`: An approval request still forces manual reconstruction of what changed or silently rewrites closed pass history.
- `Edge Cases Or Failure Modes`: Large binary changes still need summarized semantic context rather than raw path lists.
- `What This Option Avoids`: Header-only links, raw path dumps, opaque pass-history rewrites, and approval asks that assume the human will reconstruct the diff.
- `Why This Wins`: The burden is at the approval surface before the human opens artifacts, so the contract must exist at the gate itself.
- `Scope Level`: `shared`
- `Why Not Narrower Scope`: The review-shape burden is not specific to ThirdPerson.
- `Why Not Broader Scope`: It belongs in workflow docs and task artifacts, not in a platform-level IDE feature request.
- `Promotion Or Demotion Trigger`: Demote to repo-local guidance only if other repos do not exhibit approval-surface failures after a shared contract exists.
- `Reversal Path`: Keep only repo-local approval templates if shared guidance becomes too broad.

### Option B: Task-Local Approval Template Only

- `Option Label`: `Option B`
- `Summary`: Standardize a reusable approval section in `PLAN.md` and `HANDOFF.md` for this task lane without changing shared workflow docs.
- `Writeup Type`: `concrete implementation`
- `Implementation Home`: task-owned `PLAN.md` and `HANDOFF.md`.
- `Implementation Home Rationale`: It is the smallest local change and preserves flexibility.
- `Enforcement Boundary`: ThirdPerson task approval messages that point into the task folder.
- `Acceptance Test`: ThirdPerson plan or pass approvals become faster even if shared docs remain untouched.
- `Falsifier`: The human still needs to ask for links or context because the approval request itself is not standardized.
- `Edge Cases Or Failure Modes`: Different agents can ignore the local template; other repos receive no benefit.
- `What This Option Avoids`: Shared orchestration edits.
- `Why This Option Is Not The Winner Yet`: The packet burden is clearly about the approval gate shape, not just the contents of already-open docs.
- `Scope Level`: `task_local`
- `Why Not Narrower Scope`: A one-off chat message does not create a durable approval surface.
- `Why Not Broader Scope`: This option intentionally tests the narrowest honest home.
- `Promotion Or Demotion Trigger`: Promote if several tasks in the repo need the same local approval section before shared rollout is ready.
- `Reversal Path`: Drop if the shared approval packet contract absorbs the burden directly.

## P-005 First-Disagreement Debugging Gate

### Problem Boundary

Force repeated runtime-defect work to switch from tweak mode to value-traced root-cause debugging, with honest bug and regression artifacts updated around that narrowing path.

### Source Burden Drivers

- `BD-005 Root-Cause Debugging Discipline (And Defect-Tracking Honesty)`

### Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4768",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6105",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5119",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5129",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5154",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6755",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5340",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7446"
]
```

### Why This Is One Problem

The packet is consistent that the defect burden came from staying in tweak mode instead of switching to value-based first-disagreement tracing with honest bug and regression artifacts.

### Rejected Nearby Splits Or Merges

- Do not split defect tracking from debugging method. The packet repeatedly ties honest bug and regression artifacts to the same mode-switch requirement.
- Do not merge with `P-001`. Proof validity is downstream of the debugging mode choice.

### Option A: Repo-Local First-Disagreement Gate In Bug And Regression Artifacts

- `Option Label`: `Option A`
- `Summary`: Once a runtime defect reopens or survives one attempted fix, require `BUG` and `REGRESSION-RUN` artifacts to name the first concrete disagreement with values and trace its upstream writers before further fix or closeout claims.
- `Writeup Type`: `concrete implementation`
- `Implementation Home`: repo-local debugging docs plus task-owned `BUG-<NNNN>.md` and `REGRESSION-RUN-<NNNN>.md`.
- `Implementation Home Rationale`: The shared debugging method already exists, and ThirdPerson already has the repo-local adapter plus artifact homes for runtime evidence, bugs, and regression reruns.
- `Enforcement Boundary`: any reopened runtime defect or any pass where one fix attempt failed and the task is about to try another.
- `Acceptance Test`: A future reopened runtime defect records a concrete value disagreement on the human default lane and an upstream writer chain before the next code change is treated as justified.
- `Falsifier`: The task still proceeds with category-level causes like `single-node seam` or with bounded tweaks absent concrete disagreement values and traced writers.
- `Edge Cases Or Failure Modes`: Some defects will end in a bounded contradictory branch rather than one root cause; the gate must allow that if the contradiction is preserved honestly.
- `What This Option Avoids`: Tweak mode, symptom-only iteration, and bug narratives that outrun the actually narrowed evidence.
- `Why This Wins`: It uses an already-documented shared debugging method but places enforcement where this repo actually records runtime truth.
- `Scope Level`: `repo_local`
- `Why Not Narrower Scope`: A single-task reminder would not fix the next ThirdPerson runtime debugging pass.
- `Why Not Broader Scope`: The exact evidence surfaces and artifact paths are repo-specific.
- `Promotion Or Demotion Trigger`: Promote only if multiple repos need the same artifact-level first-disagreement fields with comparable runtime evidence surfaces.
- `Reversal Path`: Fall back to the shared debugging doc alone if ThirdPerson later removes repo-local runtime artifact discipline.

### Option B: Automation-Owned Runtime Checker With Quantitative Fail Thresholds

- `Option Label`: `Option B`
- `Summary`: Build or harden a `GameAutomation` checker that quantitatively fails foot-height, lateral offset, temporal continuity, or roll proxies when the runtime lower-body read is wrong.
- `Writeup Type`: `research`
- `Implementation Home`: repo-local automation code and proof scripts.
- `Implementation Home Rationale`: If the repo wants unattended runtime evidence, the automation substrate is the natural local home.
- `Enforcement Boundary`: any unattended runtime proof lane that claims believable lower-body motion on the default lane.
- `Acceptance Test`: The checker emits concrete disagreement values and rejects a run whose foot, capsule, or pose-state proxies exceed a justified envelope.
- `Falsifier`: The checker still passes a human-obviously-wrong lower-body run, or the packet cannot justify stable thresholds from durable evidence.
- `Edge Cases Or Failure Modes`: The packet does not contain enough durable evidence to freeze trustworthy numeric thresholds; capture-only automation can still miss the disputed human-visible fact.
- `What This Option Avoids`: Manual value extraction in every future debugging pass.
- `Why This Option Is Not The Winner Yet`: The packet proves the current checker family can pass wrong behavior, and the durable evidence here is insufficient to freeze trustworthy thresholds.
- `Scope Level`: `repo_local`
- `Why Not Narrower Scope`: A one-task checker would not help repeated runtime debugging inside the repo.
- `Why Not Broader Scope`: The emitted values and thresholds are tightly product-specific.
- `Promotion Or Demotion Trigger`: Promote only if stable quantitative seams emerge across repeated packets.
- `Reversal Path`: Keep deferred if honest first-disagreement tracing reduces the burden without a quantitative checker.

## P-006 Durable Constraint Ledger

### Problem Boundary

Make newly stated hard constraints and repo-local truths durable state that must be recorded and consulted before more work proceeds.

### Source Burden Drivers

- `BD-006 Durable Constraint Retention (And Repo-Local Truth Retention)`

### Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3094",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3197",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3463",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3482",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3549",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3585",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4294",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4684"
]
```

### Why This Is One Problem

The packet's constraint failures share one mechanism need: newly stated hard constraints and repo-local truths must become durable state before more work proceeds.

### Rejected Nearby Splits Or Merges

- Do not split transient hard constraints from repo-local truth retention. The same `record before proceeding` mechanism serves both.
- Do not merge with `P-002`. Remembering a rule is different from honoring pause ownership semantics.

### Option A: Task-Local Active-Constraints Ledger In Current-State Artifacts

- `Option Label`: `Option A`
- `Summary`: Add an explicit active-constraints section to the task's visible current-state artifacts and require it to record new hard constraints plus canonical repo-truth refs before more work proceeds.
- `Writeup Type`: `concrete implementation`
- `Implementation Home`: task-owned `HANDOFF.md` and `TASK-STATE.json`.
- `Implementation Home Rationale`: The packet repeatedly asked for immediate durable task-state updates, and these are already the task's human-readable and machine-readable current-state surfaces.
- `Enforcement Boundary`: any edit, build, or plan transition after a new hard constraint, gate, or repo-local truth correction is issued.
- `Acceptance Test`: `no engine mods`, `do not touch REGRESSION.md`, `stop at the approval gate`, and similar constraints appear durably before work resumes, while stable repo truths are stored as canonical doc refs rather than copied policy prose.
- `Falsifier`: A stated constraint must be repeated later or is violated despite already being given durably.
- `Edge Cases Or Failure Modes`: Overcapturing ephemeral chatter would create noise; only actionable constraints and repo-local truth corrections should enter the ledger.
- `What This Option Avoids`: Re-learning constraints, losing repo-local wording rules, silently moving past approval gates, and copying repo policy into task state when a doc ref is safer.
- `Why This Wins`: The burden is task-scoped and human-visible first, and startup rereads alone cannot absorb mid-task corrections.
- `Scope Level`: `task_local`
- `Why Not Narrower Scope`: Chat-only memory is exactly what failed.
- `Why Not Broader Scope`: The concrete constraints in this packet were task-scoped, and the repo-local truths already have canonical homes in ThirdPerson docs.
- `Promotion Or Demotion Trigger`: Promote only if many tasks repeatedly need the same active-constraints shape and leaders want it standardized in the shared task-state contract.
- `Reversal Path`: Keep only `HANDOFF.md` if maintaining machine-readable task-state parity proves too costly.

### Option B: Repo-Local Preflight Loader For Stable Truth Only

- `Option Label`: `Option B`
- `Summary`: Strengthen ThirdPerson startup and preflight guidance so the agent must reread `AGENTS`, `REGRESSION`, and `TESTING` before touching regression or runtime-debugging work.
- `Writeup Type`: `consensus`
- `Implementation Home`: repo-local [AGENTS.md](../../../../../../../../../Agent/ThirdPerson/AGENTS.md), [REGRESSION.md](../../../../../../../../../Agent/ThirdPerson/REGRESSION.md), and [TESTING.md](../../../../../../../../../Agent/ThirdPerson/TESTING.md).
- `Implementation Home Rationale`: Stable repo truth already belongs in repo-root docs and should not be duplicated elsewhere unless necessary.
- `Enforcement Boundary`: start of ThirdPerson regression or debugging tasks.
- `Acceptance Test`: Default-lane truth and supporting-lane limitations are loaded before work starts without needing later correction.
- `Falsifier`: Task-specific hard constraints still drift because repo-root docs cannot carry per-task `no engine mods` or approval-gate state.
- `Edge Cases Or Failure Modes`: Helps stable truths but not transient gates or task-specific prohibitions.
- `What This Option Avoids`: Shared memory sprawl and task-artifact schema churn.
- `Why This Option Is Not The Winner Yet`: Stable repo truth was only half the burden; the packet repeatedly needed task-specific hard constraints recorded immediately.
- `Scope Level`: `repo_local`
- `Why Not Narrower Scope`: Stable regression truth should not live only in one task.
- `Why Not Broader Scope`: The packet already shows the danger of over-generalizing repo-specific wording into shared docs.
- `Promotion Or Demotion Trigger`: Promote to shared startup guidance only if multiple repos need the same `always reread repo docs first` contract.
- `Reversal Path`: Reduce to task-specific links if repo-root preflight becomes redundant.

## Consequential Row Divergence Notes

- `P-001`
  - current preferred option: `Option A`
  - closest dangerous rival: `Option B`
  - orthogonal family collapse: automation-only proof-pack linting collapsed because it cannot adjudicate the disputed human-visible fact or repo-specific lane semantics by itself
- `P-002`
  - current preferred option: `Option A`
  - closest dangerous rival: `Option B`
  - orthogonal family collapse: notification-only stall detection collapsed because it reports a dropped lane after the human has already paid the wake-up cost
- `P-003`
  - current preferred option: `Option A`
  - closest dangerous rival: `Option B`
  - orthogonal family collapse: post-hoc answer auditing collapsed because it detects evasion only after the human has already repeated the question
- `P-004`
  - current preferred option: `Option A`
  - closest dangerous rival: `Option B`
  - orthogonal family collapse: IDE-diff reliance collapsed because the local editor surface can hide the real change and therefore cannot be the approval contract
- `P-005`
  - current preferred option: `Option A`
  - closest dangerous rival: `Option B`
  - orthogonal family collapse: shared doc-only reminders collapsed because the packet already had the shared debugging contract in play and still needed repo-local runtime evidence discipline
- `P-006`
  - current preferred option: `Option A`
  - closest dangerous rival: `Option B`
  - orthogonal family collapse: a global cross-repo memory ledger collapsed because the packet needs task-local hard constraints and canonical repo-truth refs, not broad permanent promotion

## Pass-2A Conformance Check

- All six major burden drivers from `BURDEN-ANALYSIS.md` are accounted for exactly once as `separate_problem_row`.
- Each row keeps one mechanism boundary and one proof bar. Shared substrate alone was not used as a merge reason.
- Each row includes a provisional preferred option, a closest dangerous rival, and an orthogonal family collapse reason.
- Exact implementation homes were kept on durable surfaces already evidenced by the packet when possible.
- This file stops at `2A`: no winner synthesis, no rollout order, and no task drafting.
