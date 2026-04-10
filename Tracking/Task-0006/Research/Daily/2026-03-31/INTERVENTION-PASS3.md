# INTERVENTION-PASS3 Principle Report

Source date: `2026-03-31`

Pass scope: principle distillation only. This pass stayed bounded to the March 31 PASS2 artifact plus a small set of reread transcript windows needed to settle cluster boundaries. No incident JSON files or shared workflow docs were changed in this pass.

## 1. Source scope analyzed

- PASS3 prompt: [INTERVENTION-PASS3.md](/c:/Users/gregs/.codex/Orchestration/Prompts/INTERVENTION-PASS3.md)
- PASS2 artifact used: [INTERVENTION-PASS2.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-03-31/INTERVENTION-PASS2.md)
- Event set in scope from PASS2: `C01`, `C02`, `C03`, `C04`
- PASS3 method note: PASS2 was treated as authoritative for event selection and first-pass analysis. Raw transcripts were reopened only where the principle boundary was still ambiguous, especially around `C02` and `C04`.

## 2. Candidate clusters considered

- `CL01` `real workflow truth over proxy success`:
  `C02`, `C03`, `C04`
  Kept as `P01`.
- `CL02` `explicit ownership or burden boundary`:
  `C01`, `C04`
  Kept as `P02`.
- `CL03` `standalone diagnostic-readiness rule`:
  `C02`
  Not kept separately. Merged into `P01` because the broader recurring rule is about whether the actual retest path is still ready-to-use.
- `CL04` `standalone headless-auth rule`:
  `C03`
  Not kept separately. Merged into `P01` because popup-mediated success is one form of relying on proxy workflow truth.
- `CL05` `standalone leader-owned git narrative rule`:
  `C04`
  Not kept separately. Its readability side folds into `P01`; its burden-split side folds into `P02`.

## 3. Final kept principles

### P01

- `principle_id`: `P01`
- `principle_statement`: Do not advance or call a workflow step acceptable until the real human or operator path is ready-to-use, readable, and automation-safe; proxy availability or process completion is not enough.
- `decision_point`: Before asking for a retest, treating a tool step as automation-safe, or creating a human-facing checkpoint artifact.
- `failure_signature`: `calling it done from proxy evidence`
- `why_this_is_durable`: The same miss appears across debugging, automation, and checkpoint communication whenever the assistant treats an underlying capability as sufficient even though the human-facing path has degraded. This is broader than any one tool or repo.
- `supporting_events`: `C02` ([L1539](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L1539), [L1559](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L1559), [L1582](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L1582)); `C03` ([L163](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L163), [L214](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L214)); `C04` ([L411](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L411), [L426](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L426))
- `supporting_human_model_signals`: `C02` says the path "was working pretty smoothly before" and asks for callstacks to be re-enabled before the next retest. `C03` states that if a popup click is required, the git path "would break automation later." `C04` says the commit comment should be something a human can understand that declares intent for why the commit exists.
- `counterfactual_prevention_claim`: If followed earlier, this principle would likely have caused the assistant to surface the degraded default callstack path before asking for another repro, reject a popup-backed push as non-headless, and refuse process-shaped checkpoint wording that lacked a human-readable why.
- `scope_and_non_goals`: Applies to human-facing and operator-facing workflow steps. It does not require every internal mechanism to be default-on or polished; it requires honesty about whether the actual next human or automation step is ready under real operating conditions.
- `pre_action_question`: Am I treating this as done because the underlying mechanism exists, or because the actual human or operator path works the way the next step assumes?
- `operational_check`: `closure gate`
  Before advancing a retest, automation step, or checkpoint commit, name the real operator path and explicitly note any remaining manual click, extra setup, unreadable output, or default-off switch that still makes the path not ready.
- `confidence`: `strong`

### P02

- `principle_id`: `P02`
- `principle_statement`: When the human defines a control boundary or burden split, keep the solution on that ownership seam and do not make the forbidden side the default path.
- `decision_point`: Before proposing an implementation seam or assigning responsibility for a human-facing workflow artifact.
- `failure_signature`: `solving on the wrong ownership seam`
- `why_this_is_durable`: Explicit ownership seams recur in both code and process. Once the human states who owns a surface or which side of a boundary is allowed, crossing that seam usually creates maintenance burden, review friction, or rework even if the technical idea itself is plausible.
- `supporting_events`: `C01` ([L264](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L264), [L271](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L271), [L274](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L274)); `C04` ([L411](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L411), [L426](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L426))
- `supporting_human_model_signals`: `C01` states that project code versus engine code is a real control boundary and that engine files are out of scope. `C04` explicitly asks whether leaders should own git comments so workers are not burdened with that process obligation.
- `counterfactual_prevention_claim`: If followed earlier, this principle would likely have prevented the engine-edit proposal in `C01` from becoming the primary recommendation and would have pushed checkpoint-comment ownership to the leader side sooner in `C04`.
- `scope_and_non_goals`: Applies when the human has made the seam explicit or strongly implied it through correction. It is not a blanket ban on engine changes, worker-authored text, or cross-boundary work when the human explicitly requests those paths.
- `pre_action_question`: Has the human already told me which seam or role owns this surface, and does my proposed path stay on that side?
- `operational_check`: `plan-review check`
  Before presenting a fix or workflow policy, name the allowed seam or owner explicitly and confirm that the default proposal does not cross it.
- `confidence`: `medium`

## 4. Merged or rejected principle candidates and why

### Candidate A

- `candidate_statement`: Preserve diagnostic capture defaults before asking for another repro.
- `status`: `merged`
- `reason`: This is a narrower debugging instance of `P01`. The stronger recurring rule is to verify that the actual retest path is still ready-to-use rather than relying on the existence of the underlying capability.
- `merged_into`: `P01`

### Candidate B

- `candidate_statement`: Validate headless auth behavior before using git push inside an unattended workflow.
- `status`: `merged`
- `reason`: This is a transport-specific form of `P01`. The core mistake is still treating a technically successful step as automation-safe even though the real path required human intervention.
- `merged_into`: `P01`

### Candidate C

- `candidate_statement`: Checkpoint commit text must explain to a human why the commit exists, not just which process marker fired.
- `status`: `merged`
- `reason`: This is the human-readable artifact face of `P01`. The stronger general rule is to judge checkpoint artifacts by real human readability instead of process-shaped proxy signals.
- `merged_into`: `P01`

### Candidate D

- `candidate_statement`: Leaders should own checkpoint git comments so workers are not burdened with git narrative.
- `status`: `merged`
- `reason`: This is a specific role-allocation application of `P02`. The more durable rule is to honor explicit burden splits once the human defines them.
- `merged_into`: `P02`

### Candidate E

- `candidate_statement`: Never reach into engine code when project code might be able to solve the problem.
- `status`: `merged`
- `reason`: As written, this is too local to the March 31 camera-fix event. `P02` keeps the durable part of the rule: respect the human's declared ownership seam rather than turning one project-versus-engine correction into a blanket engine taboo.
- `merged_into`: `P02`

### Candidate F

- `candidate_statement`: Be more careful about workflow ergonomics.
- `status`: `rejected`
- `reason`: Too vague to guide a future decision. It does not tell a later agent what to check before acting, and it collapses distinct misses into a non-operational slogan.

## 5. The smallest recommended principle set for this scope

The smallest honest set for `2026-03-31` is two principles:

- `P01` covers the recurring adequacy rule behind `C02`, `C03`, and the readability side of `C04`.
- `P02` covers the ownership-seam rule in `C01` and the burden-split side of `C04`.

Any smaller set would blur two materially different decision points:

- whether a workflow step is actually usable by the human or automation path that must consume it
- whether the proposed seam or owner is allowed at all

## 6. Principles still too weak to keep separately without more days or more events

- `diagnostic paths should remain ready-to-use by default before the next repro`
  This may become its own principle later, but on March 31 it is still mostly `C02` and is better carried as a sub-case of `P01`.
- `leader-owned git narrative should always be a standalone principle`
  `C04` states this very explicitly, but one day of evidence is still better treated as the burden-split facet of `P02` unless more days show the same failure pattern.
- `automation-safe git transport deserves its own top-level principle`
  `C03` is strong, but for this scope it is still a specific operator-lane form of `P01`.

## 7. Transcript windows reopened during PASS3 and why

- `C01` reread: [rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L264](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L264) through [L281](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L281)
  Reason: confirm that the correction was fundamentally about a forbidden ownership seam, not about the technical quality of the camera fix.
- `C02` reread: [rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L1539](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L1539) through [L1582](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L1582)
  Reason: decide whether `C02` supports a broader workflow-readiness rule or only a local tool-state restore. The reread supported the broader readiness interpretation, but only as a sub-case of `P01`.
- `C03` reread: [rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L163](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L163) through [L217](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L217)
  Reason: confirm that the push failure mode was not merely inconvenience but a direct contradiction of headless automation assumptions, anchoring `P01`.
- `C04` reread: [rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L411](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L411) through [L496](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L496)
  Reason: settle whether `C04` should stand alone or be split across the two stronger principles. The reread confirmed it genuinely contains both a readability rule and a burden-split rule.

No other transcript windows were reopened during PASS3.
