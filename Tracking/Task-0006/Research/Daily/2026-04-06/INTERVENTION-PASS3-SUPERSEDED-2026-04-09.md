# INTERVENTION-PASS3

Source day: `2026-04-06`

## Source scope analyzed

- PASS2 artifact: `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-06\INTERVENTION-PASS2.md`
- Candidate ids in scope: `C01` through `C26`
- PASS3 emphasis:
  - all PASS2 likely accepted incidents
  - non-incident support events `C05`, `C09`, `C10`, `C11`, `C15`, `C16`, and `C21` only where they sharpen a principle boundary
- Supporting corpus contract reread: `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\README.md`
- Transcript windows reopened during PASS3: none

## PASS2 artifact used

- Artifact used: `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-06\INTERVENTION-PASS2.md`
- Boundary facts carried forward from PASS2:
  - `C10` is only a gate-clear event on the cited refs, not honest support for the broader "infer likely human meaning" coaching theme.
  - `C12` and `C13` stay split because they capture two different misses: wrong closure model, then failure to convert the corrected model into concrete work.
  - `C23` and `C24` stay split because the first closeout order did not fully stick and the human had to intervene again to stop churn.
- Strongest PASS2 human-model signals carried into this pass:
  - `C02` and `C04`: a Home progress bar means time-to-done.
  - `C08`: a user-reported regression after the assistant's change is first-class evidence and should be handled with charity.
  - `C12` and `C13`: the human is not the operator; "done" for a long-running backend means a usable always-on service lane remains running.
  - `C14`: meaning for a human-facing control surface belongs on the surface itself.
  - `C20`: when a bounded proof becomes a concrete defect, durable task state must be updated before more debugging.
  - `C22`, `C23`, and `C24`: once the option space or closure bar is clear, analysis drift should stop.
  - `C25` and `C26`: debugging should stay on the earliest remaining deterministic divergence upstream.

## Candidate clusters considered

### `CL01` Surface meaning and form-factor truth

- Shared decision failure: the assistant let docs, planning prose, or backend truth carry meaning that the visible surface itself still miscommunicated.
- Supporting events: `C02`, `C03`, `C04`, `C14`
- Result: kept as `P01`

### `CL02` Human-world closure and operator-boundary truth

- Shared decision failure: the assistant treated proof, partial symptom relief, or a hidden extra handoff step as equivalent to a real delivered fix.
- Supporting events: `C05`, `C06`, `C07`, `C12`, `C13`, `C15`
- Result: kept as `P02`

### `CL03` Bounded execution, durable-state-first, and closeout discipline

- Shared decision failure: once the honest slice, failure state, or closure bar was already known, the assistant kept pausing, researching, narrating, or delaying instead of taking the bounded next action and writing truthful durable state first.
- Supporting events: `C10`, `C16`, `C17`, `C18`, `C20`, `C22`, `C23`, `C24`
- Result: kept as `P03`

### `CL04` Evidence-grounded problem framing and upstream debugging

- Shared decision failure: the active problem story drifted away from authoritative evidence by minimizing user regression signals, keeping stale blockers alive, guessing tool behavior, or stopping at a downstream explanation.
- Supporting events: `C08`, `C09`, `C11`, `C19`, `C21`, `C25`, `C26`
- Result: kept as `P04`

### `CL05` Active-task reprioritization

- Shared decision failure: stale task priority survived after the human had already changed the real workstream.
- Supporting events: `C01`
- Result: not kept as a standalone principle for this day; the support is too single-event and too queue-specific relative to the stronger recurring rules above.

## Final kept principles

### `P01`

- `principle_id`: `P01`
- `principle_statement`: `Make human-facing status and control surfaces communicate the intended human meaning on the surface itself; if a cue would read as time-to-done or otherwise imply a stronger promise than reality, change the surface rather than explaining it elsewhere.`
- `decision_point`: When choosing a status widget, control label, explanatory copy, or plan wording for a human-facing surface.
- `failure_signature`: `letting docs or prose carry meaning that the visible surface still misstates`
- `why_this_is_durable`: The same miss recurred across design explanation, plan grounding, and dashboard control semantics. Humans read the visible surface first, so this rule survives beyond one card or one repo.
- `supporting_events`: `C02`, `C03`, `C04`, `C14`
- `supporting_human_model_signals`: `C02` and `C04` state that a progress bar means time-to-done; `C03` says the explanation must remove follow-up burden rather than preserve ambiguity; `C14` says docs or CLI explanations are not enough for a human-facing dashboard surface.
- `counterfactual_prevention_claim`: If applied earlier, this likely would have prevented the Home progress-bar ambiguity, the follow-up "dropped ball" correction, the plan re-grounding, and the later dashboard-surface intervention.
- `scope_and_non_goals`: This applies to human-facing status and control surfaces. It is not a rule for internal debug-only affordances, and it does not require one exact widget so long as the chosen surface honestly matches the promised meaning.
- `pre_action_question`: `If a human saw only this surface, what would they think it promises, and is that promise actually true?`
- `operational_check`: `Before calling the surface acceptable, state the human reading in one sentence. If that reading is stronger than the real promise, redesign the surface or copy before closure.`
- `confidence`: `strong`

### `P02`

- `principle_id`: `P02`
- `principle_statement`: `Do not call work fixed, done, or handed off until the real human-world operating state is true without hidden operator labor; distinguish future prevention from current-state repair, and own activation steps the human cannot realistically perform.`
- `decision_point`: When deciding whether a fix is complete, what proof counts, or whether a remaining step can be pushed onto the human.
- `failure_signature`: `calling it done from partial proof, symptom relief, or hidden human handoff`
- `why_this_is_durable`: This same miss appeared in live phone repair, long-running backend closure, and dashboard activation. The durable rule is about real operating state, not any one platform.
- `supporting_events`: `C05`, `C06`, `C07`, `C12`, `C13`, `C15`
- `supporting_human_model_signals`: `C06` distinguishes preventing future bad rows from healing the current visible state; `C07` says stalled uploads mean the phone is not repaired; `C12` says the human is not the operator and the long-running backend must leave scheduled jobs running; `C13` converts that corrected model into a required always-on service lane; `C15` says not to hand the human a restart step they have no affordance to perform.
- `counterfactual_prevention_claim`: If applied earlier, this likely would have prevented the source-only stopping point, the premature "phone is repaired" claim, the teardown-based Task-0005 closure model, and the restart handoff burden.
- `scope_and_non_goals`: This is not a ban on bounded proof runs or validation lanes. Those remain useful, but they must be labeled honestly and cannot be confused with the delivered steady state.
- `pre_action_question`: `If I left right now, would the human actually have the intended working state without extra undocumented steps from them?`
- `operational_check`: `Before saying fixed, done, or handoff-ready, name the live operating state and the remaining human actions. If any essential action still sits on the human without an affordance, or if only recurrence prevention is proven, do not close.`
- `confidence`: `strong`

### `P03`

- `principle_id`: `P03`
- `principle_statement`: `When the smallest honest slice, blocker, or closure bar is already clear, stop research and analysis churn: take the bounded next action, and if the lane failed, update durable task state before any more debugging or retrying.`
- `decision_point`: When a task has narrowed to a specific implementation slice, a concrete failed proof, a final fork, or an already-satisfied closeout condition.
- `failure_signature`: `continuing research, narration, or delay after the next honest bounded move is already known`
- `why_this_is_durable`: The day showed the same failure mode in several forms: unnecessary gate pauses, open-ended research, unsafe continuation after a bounded proof failed, stale task artifacts during debugging, and delayed closeout after success was already visible. The underlying miss is failure to collapse onto the bounded next move.
- `supporting_events`: `C10`, `C16`, `C17`, `C18`, `C20`, `C22`, `C23`, `C24`
- `supporting_human_model_signals`: `C10` removes avoidable gate chasing; `C17` explicitly bans open-ended SDK and doc research once PASS-0001 is bounded; `C18` treats an out-of-bounds proof as a containment event, not normal continuation; `C20` requires `TASK-STATE.json`, `HANDOFF.md`, and a bug record before more debugging; `C22` says to choose the final bounded retry or blocked closeout now; `C23` and `C24` say that once proof has cleared the bar, the only honest move is durable closeout with minimal response churn.
- `counterfactual_prevention_claim`: If applied earlier, this likely would have reduced the plan-gate pause, the silent drift after setup, the PASS-0001 research widening, the post-failure state lag, the late fork choice, and the repeated closeout interventions after proof success.
- `scope_and_non_goals`: This is not anti-research and not anti-observation. Early broad discovery is still valid when the slice is not known yet. The rule activates once the honest next move is already bounded by the evidence.
- `pre_action_question`: `Has the option space already collapsed to one bounded next move or closeout path, and if so have I done that and written truthful durable state first?`
- `operational_check`: `Ask three questions in order: "Is the next honest move already bounded?" "Did the last step fail in a way that requires durable-state updates first?" "Has proof already cleared the closure bar?" Let the answers decide between bounded action, durable checkpointing, or closeout-only mode.`
- `confidence`: `strong`

### `P04`

- `principle_id`: `P04`
- `principle_statement`: `Ground problem framing and debugging in live authoritative evidence: preserve user-reported regressions, use the real local tool and environment behavior, and chase the earliest still-live upstream cause instead of a stale, minimizing, or downstream story.`
- `decision_point`: When reframing a blocker, selecting the next debug seam, or choosing the next retry in a failing lane.
- `failure_signature`: `problem story drifting away from live evidence or stopping at a downstream explanation`
- `why_this_is_durable`: Several April 6 misses came from story-first reasoning: minimizing a user-observed regression, retaining a false blocker, relying on guessed CLI semantics, or stopping at the first plausible downstream explanation. The corrective rule is evidence-first problem framing.
- `supporting_events`: `C08`, `C09`, `C11`, `C19`, `C21`, `C25`, `C26`
- `supporting_human_model_signals`: `C08` says a user-reported regression after the assistant's change is real evidence and should be handled with charity; `C09` says quantify it when logs already exist; `C11` and `C19` say not to keep a false spend blocker in an already-authenticated environment and to realign on actual runtime and PATH truth; `C21` injects the installed binary's real help semantics; `C25` and `C26` say debugging should stay on the earliest remaining deterministic divergence and the exact upstream structural cause.
- `counterfactual_prevention_claim`: If applied earlier, this likely would have prevented the "optimization problem" minimization, the stale spend-blocker story, the guessed executor-flag change, and the downstream debug stopping points in the hierarchy-mismatch lane.
- `scope_and_non_goals`: This does not mean every user report is automatically the root cause or that every issue needs exhaustive measurement before any action. It means the active explanation and next step must stay anchored to the strongest available evidence.
- `pre_action_question`: `What authoritative evidence source am I ignoring or overriding right now: user report, local logs, local binary semantics, actual environment state, or the earliest live divergence?`
- `operational_check`: `Before reframing or retrying, list the evidence sources that control the decision. Reject any next-step story that ignores one of those sources without saying why, and in debugging prefer the earliest still-live upstream cause over the first plausible downstream producer.`
- `confidence`: `medium`

## Rejected or merged principle candidates and why

### Candidate `R01`

- `candidate_statement`: `Treat any Home progress bar as a time-to-done promise.`
- `status`: `merged`
- `reason`: Too narrow on its own. The stronger recurring rule also has to cover ambiguous dashboard controls and explanation adequacy, not just one widget family.
- `merged_into`: `P01`

### Candidate `R02`

- `candidate_statement`: `Put dashboard meaning on the surface instead of in docs or CLI.`
- `status`: `merged`
- `reason`: Same recurring adequacy rule as the progress-bar events. Keeping a single stronger surface-meaning principle is cleaner than splitting UI semantics by component.
- `merged_into`: `P01`

### Candidate `R03`

- `candidate_statement`: `Do not say a device is repaired when you only prevented future bad rows.`
- `status`: `merged`
- `reason`: Strong local rule, but it is one sub-case of the broader human-world closure principle that also covers always-on services and restart ownership.
- `merged_into`: `P02`

### Candidate `R04`

- `candidate_statement`: `A long-running backend is not done unless the always-on service lane is left running.`
- `status`: `merged`
- `reason`: Correct, but still a domain-specific instance of the more general rule that completion must match the delivered operating state without hidden operator labor.
- `merged_into`: `P02`

### Candidate `R05`

- `candidate_statement`: `Do the restart or activation step yourself when the human has no affordance to perform it.`
- `status`: `merged`
- `reason`: Important, but it belongs under the broader operator-boundary part of the real-world completion rule rather than standing alone.
- `merged_into`: `P02`

### Candidate `R06`

- `candidate_statement`: `After a failed bounded proof, update TASK-STATE, HANDOFF, and BUG artifacts before more debugging.`
- `status`: `merged`
- `reason`: Strong and explicit, but it is one necessary branch of the larger bounded-execution rule, which also covers research drift, fork choice, and immediate closeout after success.
- `merged_into`: `P03`

### Candidate `R07`

- `candidate_statement`: `Once the smallest honest implementation slice is known, stop researching and ship it.`
- `status`: `merged`
- `reason`: This is one frequent manifestation of the same boundedness failure covered by `P03`.
- `merged_into`: `P03`

### Candidate `R08`

- `candidate_statement`: `After proof success is already clear, stop narrating and switch to closeout-only behavior.`
- `status`: `merged`
- `reason`: Valuable, but still the closeout branch of the stronger bounded-execution principle.
- `merged_into`: `P03`

### Candidate `R09`

- `candidate_statement`: `Quantify user-reported regressions when logs already exist.`
- `status`: `merged`
- `reason`: Useful, but too narrow as a standalone rule. It fits better as one operational branch of evidence-first problem framing.
- `merged_into`: `P04`

### Candidate `R10`

- `candidate_statement`: `Use the installed binary help before changing executor flags.`
- `status`: `merged`
- `reason`: Correct but tool-specific. It is better retained as one concrete operationalization of the stronger live-evidence rule.
- `merged_into`: `P04`

### Candidate `R11`

- `candidate_statement`: `Infer likely human meaning so the human intervenes less.`
- `status`: `rejected`
- `reason`: PASS2 explicitly says the cited `C10` refs do not support this broader coaching theme. April 6 does not ground this principle honestly without reopening different transcript windows or adding more days.

### Candidate `R12`

- `candidate_statement`: `Immediately rewrite the active-task model whenever the human changes direction.`
- `status`: `rejected`
- `reason`: Probably real, but April 6 provides only one clear support event (`C01`), and the rule remains too queue-specific to justify a standalone principle in this small set.

## The smallest recommended principle set for this scope

- `P01` Make human-facing meaning truthful on the surface.
- `P02` Define completion against the real human-world operating state.
- `P03` Once the honest slice or closeout bar is clear, act boundedly and put durable state first.
- `P04` Ground debugging and blocker stories in authoritative live evidence.

Four principles is the smallest honest set for this day. Folding `P01` into the closure or boundedness rules would erase the distinct human-facing semantics standard, and folding `P04` into `P03` would collapse specific evidence and debugging discipline into a vague reminder to stay grounded.

## Principles still too weak and need more days or more events

- `Infer likely human meaning so the human intervenes less.` This likely exists as a broader coaching theme, but April 6 does not ground it on the cited refs and PASS3 did not reopen wider windows.
- `Rewrite task priority immediately on human rescope.` `C01` supports it locally, but the day does not yet show enough repetition to keep it as a durable standalone principle.
- `Consult real tool help before every executor or CLI flag change.` Real and useful, but still too narrow as its own principle and better carried as a sub-case of `P04` unless it repeats across more days.

## Transcript windows reopened during PASS3 and why

- None. PASS2 already provided the needed boundary corrections, repeated cluster hints, and strongest human-model signals, so the PASS3 reread triggers were not met.
