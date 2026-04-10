# INTERVENTION-PASS3

Source day: `2026-04-06`

Canonical note: promoted on `2026-04-09` from the explanatory rerun; the earlier day-local PASS3 is archived as `INTERVENTION-PASS3-SUPERSEDED-2026-04-09.md`.

## Source scope analyzed

- Source day: `2026-04-06`
- PASS2 artifact used: `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-06\INTERVENTION-PASS2.md`
- Contract/context inputs used for this pass: `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\README.md` and `C:\Users\gregs\.codex\Orchestration\Prompts\INTERVENTION-PASS3.md`
- Existing April 6 PASS3 material was not used as evidence for this rerun.

## Candidate clusters considered

| Cluster | Events | PASS3 resolution |
| --- | --- | --- |
| `status-signal-meaning` | `C01`, `C02`, `C03`, `C04` | Kept as the core of `P01`. |
| `control-surface` / `form-factor-protection` | `C18`, `C19` | Split across `P01` and `P05`: on-surface meaning stays with `P01`; restart/last-mile ownership stays with `P05`. |
| `user-report-evidence` / `principle-of-charity` / `repeated-teaching` | `C03`, `C05`, `C11` | Kept as `P02`. |
| `autonomy-persistence` / `wrong-seam-debugging` | `C07`, `C08` | Kept as `P03`. |
| `boundedness` / `false-blocker` / `local-grounding` / `decision-discipline` | `C06`, `C09`, `C11`, `C13`, `C14`, `C15` | Merged into `P04`. |
| `proof-scope-control` / `durable-state-first` / `closure-truth` / `operator-boundary` | `C10`, `C12`, `C16`, `C17`, `C19` | Merged into `P05`. |

## Final kept principles

### `P01` - Design human-facing state and controls to the ordinary reading on the surface

- `principle_statement`: On a human-facing surface, only use a status or control metaphor when the surface itself truthfully communicates the ordinary human meaning; put the meaning on the surface rather than in docs, backend lore, or implied caveats.
- `decision_point`: Choosing progress/status UI, visible fix framing, or quick-use control wording on the main human-facing surface.
- `failure_signature`: Proxy or backend semantics presented as if they were self-evident user-facing truth.
- `why_this_is_durable`: The strongest April 6 UI coaching is not about one widget; it is about ordinary human reading. A Home progress bar was read as time-to-done, not generic activity. A dashboard control surface was expected to carry its own meaning, not require docs or CLI interpretation. The same adequacy rule applies across status indicators and quick-use controls.
- `supporting_events`: `C01`, `C02`, `C03`, `C04`, `C18`
- `supporting_human_model_signals`: A Home progress bar means `how long until done`; if that promise is not true, do not use the bar. If the dashboard is the human-facing control surface, meaning must live there first. When the visible user-facing state will not change yet, say that plainly instead of implying otherwise.
- `counterfactual_prevention_claim`: If applied earlier, this would likely have prevented the repeated Home-bar reteaching, the partial-fix confusion around the blocked banner, and the dashboard dropped-ball around hidden control meaning.
- `scope_and_non_goals`: This is not a ban on progress bars, docs, or operator tools. It applies when a surface is meant to be read quickly by a human without auxiliary interpretation.
- `pre_action_question`: If the human only saw this surface, would they correctly understand what is happening and whether work is done?
- `operational_check`: Audit question: `What ordinary human promise does this surface element make, and is that promise true end-to-end?`
- `confidence`: `strong`

### `P02` - Treat repeated human teaching and lived regression reports as first-class evidence

- `principle_statement`: Treat explicit human follow-up burden and user-reported regressions as first-class evidence; do not explain them away before recording them and updating the active model.
- `decision_point`: When a human says they still had to follow up, says a fix made reality worse, or repeats a blocker correction that should already have propagated.
- `failure_signature`: Rebutting or flattening the human report into optimization noise while the human is still teaching the same point.
- `why_this_is_durable`: The April 6 explanatory-labor rerun makes clear that repeated human teaching is not background chatter. It is supervision cost and evidence that the prior correction did not stick. The regression-report event is explicit, and the repeated follow-up events show the same meta-rule from a different angle.
- `supporting_events`: `C03`, `C05`, `C11`
- `supporting_human_model_signals`: When the human says they still had to follow up, that burden is itself evidence of a miss. A report like `it was faster before your fix` is real regression evidence, not something to recast as mere optimization without comparative proof. A blocker correction that has to be repeated across threads is also evidence that the lesson did not propagate.
- `counterfactual_prevention_claim`: If applied earlier, this would likely have caused the regression report to be preserved and framed charitably at once, and it would have turned repeated follow-up into a plan-changing signal rather than a second correction on the same theme.
- `scope_and_non_goals`: This does not require blindly accepting every human claim as the final causal diagnosis. It does require preserving the claim as evidence and treating repeated teaching as a real miss.
- `pre_action_question`: Has the human just spent effort correcting my model or reporting worse lived behavior, and have I elevated that to evidence instead of rebuttal?
- `operational_check`: Plan-review check: any repeated correction or user-reported regression must appear in notes, state, or bug tracking before further interpretation proceeds.
- `confidence`: `strong`

### `P03` - On deterministic debugging tasks, do not stop at the first plausible story

- `principle_statement`: On deterministic debugging tasks, keep tracing the current highest-signal seam until the exact structural cause or an honest blocker is explicit; do not stop at the first plausible upstream story.
- `decision_point`: Deciding whether a debug explanation is sufficient or which remaining seam deserves the next investigation step.
- `failure_signature`: Premature satisfaction with an upstream producer or continued work on a stale seam after the remaining disagreement has narrowed elsewhere.
- `why_this_is_durable`: The human stated the autonomy contract directly: `pursue root cause deterministically` means exact structural cause, not plausible producer. The related wrong-seam redirection shows the same rule at a narrower layer.
- `supporting_events`: `C07`, `C08`
- `supporting_human_model_signals`: Deterministic root-cause pursuit means do not stop at the first plausible upstream producer. Once the remaining disagreement is narrower, move to that exact seam instead of polishing an older story.
- `counterfactual_prevention_claim`: If applied earlier, this would likely have prevented the premature handoff in `C07` and the extra time spent on the stale Head-hierarchy story in `C08`.
- `scope_and_non_goals`: This is not a demand to debug forever. Stopping is valid once the exact structural cause is explicit or a real blocker is documented honestly.
- `pre_action_question`: Have I named the exact structural cause at the current decisive seam, or only the earliest plausible explanation?
- `operational_check`: Debugging heuristic: before handoff, state the exact remaining disagreement and why the current cause claim actually closes it.
- `confidence`: `strong`

### `P04` - Collapse to the real blocker from concrete local facts, then take the bounded next branch

- `principle_statement`: Collapse to the real blocker from concrete local/runtime facts, then take the bounded next branch; do not let open-ended research, assumed environment limits, or analysis drift replace that decision.
- `decision_point`: Gating execution, choosing between more research and implementation, or deciding whether to retry, declare blocked, or close out.
- `failure_signature`: Invented approval or spend blockers, stale environment inference, or prolonged analysis after the branch set is already small.
- `why_this_is_durable`: This cluster is broader than one false blocker. Across April 6, the human repeatedly had to collapse speculation into the actual blocker, the exact local facts, and a bounded next branch. The same discipline also appears in the anti-research correction and the forced decision/closeout corrections.
- `supporting_events`: `C06`, `C09`, `C11`, `C13`, `C14`, `C15`
- `supporting_human_model_signals`: Once a bounded slice and live runtime exist, stop open-ended research. A pre-authenticated local hosted tool is not automatically an `additional money` blocker. Use exact local paths and actual CLI semantics when the machine can reveal them. Once the acceptable branch set is explicit, decide now rather than staying in analysis. After proof success, close out directly.
- `counterfactual_prevention_claim`: If applied earlier, this would likely have cut the repeated spend-gate clarification, the environment-guesswork teaching, and the late-task churn between analysis, retry choice, and closeout.
- `scope_and_non_goals`: This is not anti-research in general, and it does not forbid honest blocker declarations. It applies once machine facts and task bounds are already sufficient to collapse the decision.
- `pre_action_question`: What is the actual blocker proven by current machine facts, and what bounded action follows from it right now?
- `operational_check`: Plan-review check: every proposed next step must name the actual blocker, its evidence source, and the exit condition for this branch.
- `confidence`: `medium`

### `P05` - Treat live-state truth and unattended operability as part of done

- `principle_statement`: For live/runtime and backend tasks, treat scope fences, truthful durable state, and unattended operability as part of done; the human is not the default operator.
- `decision_point`: Before starting live hosted work, during recovery from proof-scope failures, and when declaring runtime/backend work complete.
- `failure_signature`: Proof-only success treated as completion, unsafe runtime started without scope fencing, or hidden last-mile action handed back to the human.
- `why_this_is_durable`: This cluster carries the strongest real-world completion coaching in the rerun. The same underlying adequacy rule appears in unsafe worker startup, durable-state-first recovery, rejection of proof-only closure for a `long-running backend`, creation of the always-on service lane, and rejection of the restart handoff.
- `supporting_events`: `C10`, `C12`, `C16`, `C17`, `C19`
- `supporting_human_model_signals`: Prove scope fences before live worker-backed proofs. After a proof-scope failure, durable-state truth is part of recovery. `Long-running backend` implies a usable unattended operating model. The human is not the operator by default. If the assistant can perform the final restart or activation, it should own that last mile unless explicitly told otherwise.
- `counterfactual_prevention_claim`: If applied earlier, this would likely have prevented the unsafe worker startup, forced truthful task state immediately after the proof-scope failure, and blocked proof-only closure or restart handoff as acceptable end states.
- `scope_and_non_goals`: This does not ban explicit trained-operator workflows when the task says so. It applies when the requested outcome implies an unattended local service or an assistant-owned last-mile step.
- `pre_action_question`: If I stop now, will the system actually remain in the claimed running state without hidden operator work from the human?
- `operational_check`: Closure gate: for any live-runtime task, verify running-state evidence, truthful durable artifacts, and ownership of any final activation or restart step.
- `confidence`: `strong`

## Rejected or merged principle candidates and why

| Candidate statement | Status | Reason | Merged into |
| --- | --- | --- | --- |
| `If Home shows a progress bar, it must mean time-to-done.` | `merged` | Too narrow; it is one vivid instance of the broader human-surface semantics rule. | `P01` |
| `Put dashboard control meaning on the dashboard, not in docs or CLI.` | `merged` | Same decision failure as the progress-bar events: meaning was left off the primary surface. | `P01` |
| `When the human says this is a dropped ball, treat that follow-up burden as evidence.` | `merged` | Better kept as part of the broader repeated-teaching and regression-evidence rule. | `P02` |
| `Apply a principle of charity to user-reported regressions.` | `merged` | Strong and explicit, but it is best modeled as the clearest sub-case of treating human correction as first-class evidence. | `P02` |
| `After a bounded slice exists, stop open-ended research.` | `merged` | Durable, but on this day it is better treated as one manifestation of collapsing to the real blocker and bounded next branch. | `P04` |
| `A pre-authenticated local CLI is not an additional-money blocker.` | `merged` | Important repeated teaching, but too local to keep as a standalone principle once merged into the broader blocker-grounding rule. | `P04` |
| `After proof success, close out immediately.` | `merged` | Strong locally, but best treated as the closeout end of the same bounded-branch discipline. | `P04` |
| `Prove scope fences before launching worker-backed runtime.` | `merged` | This is the pre-execution face of the stronger live-state and unattended-operability rule. | `P05` |
| `Long-running backend means actually left running.` | `merged` | Explicit and central, but still a narrower expression of the full operator-boundary and real-world-done rule. | `P05` |
| `Do not hand the restart back to the human.` | `merged` | Too specific to keep separately once last-mile ownership is absorbed into the real-world-done principle. | `P05` |

## Smallest recommended principle set for this scope

- The honest minimum for April 6 is `5` principles: `P01` through `P05`.
- Collapsing below `5` would start to blur meaningfully different decision points:
  - surface semantics on the human-facing UI
  - human correction and regression reports as evidence
  - deterministic debug persistence
  - blocker-grounding and bounded branch choice
  - real-world runtime closure and operator-boundary ownership
- If forced to compress further, `P04` and `P05` could be merged into one broader execution/closure rule, but that would make pre-execution blocker discipline and post-execution operating-model truth harder to audit separately.

## Principles still too weak to keep separately

- No additional kept principle is marked `weak`.
- The following ideas may be durable later, but April 6 alone does not justify them as standalone principles after aggressive merging:
  - `anti-research drift as its own independent principle`
  - `pre-authenticated tooling is not a spend gate`
  - `immediate closeout after proof success`

## Transcript windows reopened during PASS3 and why

- None.
- PASS2 already preserved the analyzed event set, cluster hints, and strongest human-model signals with enough clarity to resolve the principle boundaries without another raw transcript reread.
