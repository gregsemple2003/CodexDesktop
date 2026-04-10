# INTERVENTION-PASS3

Source day: `2026-04-02`

## Source scope analyzed

- Zero-context backfill using only local artifacts derived from this source day.
- PASS2 artifact analyzed: `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-02\INTERVENTION-PASS2.md`
- Event set in scope: `C01` through `C18`, exactly as bounded in PASS2.
- PASS3 worked from PASS2 sections on boundary corrections, repeated cluster hints, strongest human-model signals, likely accepted incidents, likely non-incident interventions, and the specific event records most relevant to cluster merges: `C01`, `C03`, `C06`, `C09`, `C11`, `C14`, `C15`, and `C18`.
- No parent-thread summaries, accepted-incident JSONs, or non-day artifacts were consulted.

## PASS2 artifact used

- `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-02\INTERVENTION-PASS2.md`

## Candidate clusters considered

- `CL01` - Reality over proxy closure.
  Events considered: `C04`, `C05`, `C07`, `C08`, `C09`, `C10`, `C13`, `C16`, `C17`.
  Shared failure: declaring progress from workflow edits, narrated fixes, partial symptom relief, or unproven launch paths instead of the real running path.
  Outcome: kept as `P01`.
- `CL02` - Literal human boundaries and named scope.
  Events considered: `C01`, `C02`, `C11`, `C12`.
  Shared failure: silently widening scope or choosing a convenience path after the human already named the boundary, target, lane, or control condition.
  Outcome: kept as `P02`.
- `CL03` - Actor-owned guidance and supervision liveness.
  Events considered: `C01`, `C04`, `C06`, `C14`.
  Shared failure: responsibility living in the wrong layer or supervision ending before the delegated actor reached a real stop condition.
  Outcome: kept as `P03`.
- `CL04` - Human-readable state and causal adequacy.
  Events considered: `C03`, `C15`, `C18`.
  Shared failure: surfaces that look informative but do not say the real missing condition, real scale, or real cause in ordinary human terms.
  Outcome: kept as `P04`.
- `CL05` - Self-verification before asking the human again.
  Events considered: `C07`, `C10`, `C17`.
  Shared failure: handing proof burden back to the human before the assistant has proven the local path.
  Outcome: merged into `P01` because it is a narrower expression of proxy-vs-reality failure.
- `CL06` - Investigation must answer root cause, not summary.
  Events considered: `C02`, `C08`, `C18`.
  Shared failure: stopping at local summaries or partial slices when the human asked for the trigger, mechanism, and avoidable mistake.
  Outcome: merged into `P04` because it is the investigation-specific form of human-readable state and causal adequacy.

## Final kept principles

### P01 - Verify the real operating state before claiming closure or handing proof back

- `principle_id`: `P01`
- `principle_statement`: Do not call work done, fixed, or ready for human retest from workflow edits, code diffs, narrated fixes, partial symptom relief, or tool output alone; verify the actual running path and user-relevant behavior first.
- `decision_point`: Before declaring closure, reopening downward, saying a fix landed, or asking the human to validate the next step.
- `failure_signature`: calling it done from proxy evidence
- `why_this_is_durable`: The same miss appears across UI passes, stale-build review, runtime debugging, launch-path verification, and technically-working-but-invalid fixes. The surface changes, but the decision failure is stable.
- `supporting_events`: `C04` workflow hardening presented as closure; `C05` Sync shell gaps still visible after pass close; `C07` old installed build still running; `C08` solved slice treated like practical closure while the live client problem remained; `C09` "that last one worked" despite being wrong for networked Mover behavior; `C10` relaunch narration while wrapper/path error stayed visible; `C13` visible bottom-bar `Home` affordance was dead; `C16` claimed layout fix while overlay was still clipped; `C17` human had to demand local self-test before another retest request.
- `supporting_human_model_signals`: `C04` return with a working emulator and all issues fixed; `C05` the shell bars should always be in view; `C07` if the human still sees the old UI, the fix does not count; `C09` a fix is not acceptable just because it worked once if it violates the real runtime model; `C16` remove the bottom bar and leave room for the graph; `C17` test it yourself before making me test again.
- `counterfactual_prevention_claim`: If this rule had been applied early, it likely would have prevented most of the day's reopenings that came from proxy closure, stale proof, or premature handoff back to the human.
- `scope_and_non_goals`: This is not a ban on interim status updates. It applies when the assistant is making adequacy claims or shifting proof burden. Incremental notes are still fine if clearly labeled as not-yet-proven.
- `pre_action_question`: Am I about to call this done or ask the human to verify something that I have not yet proven on the real running path?
- `operational_check`: Closure gate or audit question: "What exact live surface, running build, or end-to-end path did I personally verify, and what visible behavior proves it?"
- `confidence`: `strong`

### P02 - Treat explicit human boundaries as literal constraints unless renegotiated

- `principle_id`: `P02`
- `principle_statement`: Treat the human's named scope, control boundary, repro condition, and action lane as literal task constraints; do not widen them or route around them without explicit re-approval.
- `decision_point`: When interpreting an instruction, choosing a control path, deciding what to touch, or deciding what repo/path/surface is in scope.
- `failure_signature`: helpful widening past a named boundary
- `why_this_is_durable`: The same literalism failure shows up in documentation seams, debugging controls, repo scope, and emergency freeze moments. The recurring problem is not misunderstanding the domain, but silently overriding a boundary the human already made explicit.
- `supporting_events`: `C01` ownership rule spread into `AGENTS.md` instead of staying in the governing prompt layer; `C02` human said do not touch anything and requested forensics first; `C11` background-safe, `CombatMap`, no-focus debugging contract had to be restated repeatedly; `C12` `push .codex` widened into a staged-scope repo commit.
- `supporting_human_model_signals`: `C01` keep `AGENTS.md` light and avoid free-floating responsibility; `C02` no touch actions while the human cannot see what is burning tokens; `C11` the backgrounded editor and no-focus condition are part of the repro contract; `C12` when i say push `.codex` it means just `.codex`.
- `counterfactual_prevention_claim`: If this rule had been applied early, the assistant likely would have avoided the wrong doc seam, avoided extra churn during token-loss diagnosis, respected the no-focus repro boundary, and avoided the widened `.codex` commit scope.
- `scope_and_non_goals`: This is not a ban on proposing alternatives. It only says the assistant must not silently switch boundaries. Alternatives require explicit renegotiation before action.
- `pre_action_question`: What exact boundary did the human name here, and am I widening scope or changing the control path without asking?
- `operational_check`: Plan-review check: "List the literal repo/path/window/control constraints the human set, then confirm the planned action stays inside them."
- `confidence`: `strong`

### P03 - Put corrective responsibility on the actor that must execute it, and keep it live until a real endpoint

- `principle_id`: `P03`
- `principle_statement`: When a miss is caused by delegation or workflow structure, attach the corrective rule to the actor and layer that actually controls the behavior, and keep supervisory ownership alive until that actor reaches a real stop condition.
- `decision_point`: After a delegated role misses, when deciding where a new rule should live, and when deciding whether supervision may end.
- `failure_signature`: free-floating responsibility or premature supervision exit
- `why_this_is_durable`: The day shows the same governance error in two forms: rules landing in the wrong layer, and supervisors assuming a status update is equivalent to a finished delegated endpoint. Both come from losing track of who still owns the next real action.
- `supporting_events`: `C01` task-leader responsibility belonged in the task-leader prompt, not front-door docs; `C04` prompt hardening was treated as enough while the actual task still needed to be carried through; `C06` the implementation role needed self-audit hardening against the exact missed gaps; `C14` supervision stopped before the delegated task leader was actually done.
- `supporting_human_model_signals`: `C01` responsibility should not be left as a free-floating constraint; `C04` take the ball for awhile and come back with the working emulator; `C06` harden the prompt around the gaps you are witnessing now; `C14` continue working until the task leader is done.
- `counterfactual_prevention_claim`: If this rule had been applied early, the assistant likely would have put the right obligation in the right prompt, hardened the delegated role after the first miss, and avoided dropping supervisory ownership before the child leader reached a real endpoint.
- `scope_and_non_goals`: This is not a command to supervise forever. A real blocker, explicit human gate, or completed endpoint still ends supervision. The rule is only against treating summaries or indirect ownership as sufficient.
- `pre_action_question`: Does the actor who must perform the next real action actually carry this rule, and if I am supervising, has that actor truly reached a stop condition yet?
- `operational_check`: Audit question: "Which prompt, role, or owner now contains the corrective rule, and what concrete endpoint marks supervision as legitimately over?"
- `confidence`: `strong`

### P04 - Make human-facing state and investigation output answer the real human question in plain terms

- `principle_id`: `P04`
- `principle_statement`: For human-facing surfaces, name the real missing condition, real scale, or real cause in ordinary language; ambiguous labels, impossible metrics, and summary-only output are not acceptable substitutes.
- `decision_point`: When choosing status copy, presenting numbers, or shaping investigation output for a human consumer.
- `failure_signature`: informative-looking output that still leaves the human asking what this actually means
- `why_this_is_durable`: The same adequacy rule spans UI copy, dashboard metrics, and investigation reports. In each case the assistant produced something structurally plausible that still failed the human's actual interpretive question.
- `supporting_events`: `C03` `Connected` and `Not ready` language was too vague and in the wrong semantic domain; `C15` the `8M` budget number did not make sense relative to the quota story; `C18` investigation output summarized sessions instead of explaining the operator action, mechanism, and avoidable mistake.
- `supporting_human_model_signals`: `C03` vague status language is too open for interpretation; `C15` if the budget number does not make sense, semantic truth outranks cosmetic polish; `C18` the tool should answer what the human did to cause the burst and give root cause, no bullshit.
- `counterfactual_prevention_claim`: If this rule had been applied early, the assistant likely would have avoided shipping ambiguous status copy, avoided polishing around impossible numbers, and aimed the investigation product at the causal answer the human actually wanted.
- `scope_and_non_goals`: This is not a demand for maximum verbosity. Short labels and concise briefs are still fine if they directly answer the human's interpretive question and do not hide the real condition or cause.
- `pre_action_question`: If a cold human saw this output, would they know what is missing or what caused the problem without having to ask a follow-up meaning question?
- `operational_check`: Prompt rule or review gate: "Every human-facing state/report must answer one of these explicitly: what is missing, what changed, what caused it, or what to avoid next time."
- `confidence`: `strong`

## Rejected or merged principle candidates and why

- `candidate_statement`: Self-test the launch or debug path yourself before making the human retry it.
  `status`: `merged`
  `reason`: This is a narrower instance of the stronger P01 rule against proxy closure and proof handoff without real-path verification.
  `merged_into`: `P01`
- `candidate_statement`: Never accept a fix just because the last run worked.
  `status`: `merged`
  `reason`: This is the runtime-specific form of P01's broader rule that live adequacy outranks narrated or one-off success signals.
  `merged_into`: `P01`
- `candidate_statement`: Keep `AGENTS.md` light.
  `status`: `rejected`
  `reason`: Too doc-local on its own. The durable rule is either literal boundary respect (`P02`) or placing responsibility on the governing actor layer (`P03`), not this one document by itself.
- `candidate_statement`: `push .codex` means just `.codex`.
  `status`: `merged`
  `reason`: Strong evidence, but still a specific instance of the broader literal-boundary principle.
  `merged_into`: `P02`
- `candidate_statement`: Background-safe PIE on `CombatMap` with no focus theft.
  `status`: `merged`
  `reason`: Durable only as a concrete example of respecting explicit human control and repro boundaries.
  `merged_into`: `P02`
- `candidate_statement`: Continue supervising until the task leader is done.
  `status`: `merged`
  `reason`: This is the supervisory form of the stronger P03 rule about keeping responsibility with the real actor until a true endpoint.
  `merged_into`: `P03`
- `candidate_statement`: Investigation output must give root cause, not session summary.
  `status`: `merged`
  `reason`: This is the investigation-specific expression of P04's broader human-readable state and causal adequacy rule.
  `merged_into`: `P04`
- `candidate_statement`: Remove the bottom bar and leave room for the graph.
  `status`: `rejected`
  `reason`: Too local to one screen fix. The durable rule is P01's real-surface verification requirement, not the exact layout remedy used that night.

## The smallest recommended principle set for this scope

- Recommended minimum: `4` principles.
- Keep `P01`, `P02`, `P03`, and `P04` as the smallest honest set that still preserves distinct decision points:
  `P01` covers adequacy and proof against the real operating state.
  `P02` covers literal boundary respect.
  `P03` covers actor placement and supervision liveness.
  `P04` covers human-readable state and causal adequacy.
- Collapsing below four would force either `P02` and `P03` together or `P01` and `P04` together. Either merge would blur an important operational distinction and weaken later auditability.

## Principles still too weak and needing more days or more events

- Possible future candidate: freeze action and pivot to local forensics when the human reports unexplained quota loss plus missing local history.
  Current support: mainly `C02`.
  Why still weak: strong local signal, but only one clear day-level example here.
- Possible future candidate: keep debugging centered on the still-live failure even after one sub-cause is explained.
  Current support: `C08` as a lead-in and `C09` as the stronger accepted event.
  Why still weak: on this day it is better treated as part of `P01` than as its own principle.
- Possible future candidate: split no-focus control boundaries from wrong-map repro boundaries.
  Current support: bundled together inside `C11`.
  Why still weak: PASS2 already noted that this day does not require a split.

## Transcript windows reopened during PASS3 and why

- No raw transcript windows were reopened during PASS3.
- PASS2 already preserved the necessary event boundaries and explicit human-model signals for principle extraction.
- PASS3 only reread selected PASS2 event records to resolve merges and keep the principle set minimal.
