# INTERVENTION-PASS3

Source day: `2026-04-02`

Canonical note: promoted on `2026-04-09` from the April 2 rerun; the earlier day-local PASS3 is archived as `INTERVENTION-PASS3-SUPERSEDED-2026-04-09.md`.

## Source scope analyzed

- PASS2 artifact:
  - `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-02\INTERVENTION-PASS2.md`
- Incident corpus contract:
  - `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\README.md`
- PASS3 prompt contract:
  - `C:\Users\gregs\.codex\Orchestration\Prompts\INTERVENTION-PASS3.md`
- Raw transcript windows reopened during PASS3:
  - none; the rerun PASS2 boundary notes and human-model signals were sufficient for cluster and merge decisions

## PASS2 artifact used

- `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-02\INTERVENTION-PASS2.md`

## Candidate clusters considered

### CL01 - Visible structure and mockup/product-read fidelity

- Events: `C03`, `C05`, `C08`, `C10`, `C13`
- Shared decision failure:
  - the surface technically existed, but shell structure, card composition, iconography, or ordering still changed how the product read to a human

### CL02 - Typography, spacing, clipping, and color readability

- Events: `C11`, `C12`, `C16`, `C17`
- Shared decision failure:
  - the surface rendered, but density, rhythm, clipping, or color treatment still made it hard for a human to read

### CL03 - State semantics and truthful explanation on human-facing surfaces

- Events: `C04`, `C15`, `C18`, `C19`
- Shared decision failure:
  - the surface showed labels or summaries without delivering the human-meaningful truth or causal answer the human actually needed

### CL04 - Closure truth on the actual running surface

- Events: `C06`, `C08`, `C09`, `C14`
- Shared decision failure:
  - progress or fixes were narrated from proxy evidence instead of the actual running, installed, or behaviorally equivalent target

### CL05 - Real owner, real seam, real cause

- Events: `C01`, `C02`, `C05`, `C07`, `C18`, `C19`
- Shared decision failure:
  - action or guidance was routed to an easier adjacent layer instead of the layer that actually owned the decision or cause

## Cluster merge outcome

- `CL01`, `CL02`, and the surface-facing part of `CL03` merge into one kept principle because the human standard is the same: if the surface does not communicate the intended truth readably, it is still wrong.
- `CL04` stays separate because the decision point is closure and proof, not surface design.
- `CL05` stays separate because the decision point is where to intervene or investigate, not how the final surface reads.

## Final kept principles

### P01

1. `principle_id`: `P01`
2. `principle_statement`: `Treat human-facing surfaces as incomplete until they communicate the truthful, human-meaningful answer in a readable form; visible fidelity, typography, clipping, copy, iconography, and hierarchy are correctness, not polish.`
3. `decision_point`: When shipping or revising any UI, dashboard, status card, or other human-facing readout.
4. `failure_signature`: `The surface technically exists, but a human still asks what it means, cannot read it comfortably, or loses trust because the layout or readout still reads wrong.`
5. `why_this_is_durable`: The same miss recurred across Sync, Home, and CodexDashboard. The shared human standard was not pixel matching; it was that structure, wording, hierarchy, and readability determine whether the human can understand or trust the surface.
6. `supporting_events`:
   - `C03` - hero control lost its intended composed shape
   - `C04` - `Connected` status card meaning was unclear
   - `C05` - peer-card parity, copy clarity, and icon drift were treated as real misses
   - `C08` - shell structure omissions made a done-claim visibly false
   - `C10` - post-build iconography, alignment, and cramped-button issues still changed the screen's read
   - `C11` - typography and chip sizing were corrected as legibility failures
   - `C12` - spacing and hierarchy still left the summary cluttered
   - `C13` - card separation and order carried product meaning on Home
   - `C15` - visible budget math had to become truthful before polish work continued
   - `C16` - footer clipping and control crowding made the dashboard hard to read
   - `C17` - repo-mode color treatment was unreadable for a human
7. `supporting_human_model_signals`:
   - PASS2 explicitly carried that screenshot-visible typography, spacing, iconography, card shell, and clipping can materially distort the intended human-facing outcome.
   - PASS2 explicitly carried that status cards should explain human-meaningful state directly.
   - PASS2 explicitly carried that dashboards must tell the truth before surrounding polish matters.
   - Inference from PASS2: mockup fidelity matters here when it changes product read or readability on the real surface.
8. `counterfactual_prevention_claim`: If applied earlier, this principle would likely have prevented most of the repeated UI-fidelity, typography, clipping, and product-read interventions because those misses would have been treated as acceptance failures rather than late polish.
9. `scope_and_non_goals`: This is not a demand for pixel-perfect mimicry or a ban on experimentation. It applies when visible structure, wording, hierarchy, or styling changes what a human can understand, read, or trust on the actual surface.
10. `pre_action_question`: `If the human only saw this surface, would they immediately understand the truthful state and read it comfortably without decoding, squinting, or asking what it means?`
11. `operational_check`: `Before claiming UI progress, run a surface-read review that checks self-evident meaning, truthful metrics or copy, readable typography and contrast, intact composition, and no clipping or crowding.`
12. `confidence`: `strong`

### P02

1. `principle_id`: `P02`
2. `principle_statement`: `Do not call work done, fixed, or improved from source edits, prompt or doc motion, or stale builds; verify the actual running, installed, or behaviorally equivalent target first.`
3. `decision_point`: Before claiming completion, requesting review, or narrating a fix.
4. `failure_signature`: `Success is narrated from code or artifact progress while the real surface still shows old behavior, missing structure, or no user-visible effect at all.`
5. `why_this_is_durable`: This recurred in UI closure, build verification, and real-machine regression. The human kept reasserting the same standard: closure depends on the actual target state, not on nearby evidence that sounds encouraging.
6. `supporting_events`:
   - `C06` - prompt-only progress did not fix the real emulator-side loop
   - `C08` - a done-claim collapsed under visible shell omissions
   - `C09` - the emulator was still showing the old build
   - `C14` - the real hotkey still did nothing on the actual machine
   - `C07` - a workaround changed the target behavior instead of proving the real fix
7. `supporting_human_model_signals`:
   - PASS2 explicitly carried that if the human cannot see the claimed fix on the actually installed or running surface, the fix is not yet real.
   - PASS2 explicitly carried that prompt and doc changes do not count if the real visible workflow is still broken.
   - PASS2 explicitly carried that real-machine behavior outranks supporting artifacts when closing a desktop feature.
   - PASS2 explicitly carried that a workaround that changes the target behavior is not an adequate fix.
8. `counterfactual_prevention_claim`: If applied earlier, this principle would likely have blocked the false closure attempts on Sync and CodexDashboard and would have forced live-build or real-machine verification before asking the human to trust the result.
9. `scope_and_non_goals`: Internal progress can still be reported, but it must be labeled as unverified when the actual target has not been checked. This principle is about closure and review claims, not about banning intermediate work.
10. `pre_action_question`: `What actual running surface or equivalent scenario have I verified, and can the human reproduce the claimed improvement there right now?`
11. `operational_check`: `Require closure evidence from the live build, real machine, or behaviorally equivalent target; otherwise mark the claim as unverified and keep the task open.`
12. `confidence`: `strong`

### P03

1. `principle_id`: `P03`
2. `principle_statement`: `Investigate and route fixes at the real causal seam: inspect the full operating state before acting, and do not move responsibility, debugging, or explanation to an easier adjacent layer that leaves the real cause untouched.`
3. `decision_point`: When choosing an action, owner, or debugging path after a problem report or investigation request.
4. `failure_signature`: `A fast action, doc patch, workaround, or summary changes something nearby while the real owner, mechanism, or causal action remains unexamined.`
5. `why_this_is_durable`: This recurred in docs and ownership routing, quota triage, Rustfire debugging, and burst investigation. The human kept correcting the assistant away from shallow or shifted framings and back toward the real owner or cause.
6. `supporting_events`:
   - `C01` - UI-fidelity guidance was moved into bloated front-door docs instead of the real owner seam
   - `C02` - a shallow quota read almost triggered premature action
   - `C05` - repeated fidelity misses forced creation of an explicit interface-design owner lane
   - `C07` - the standalone workaround changed the scenario instead of solving the networked cause
   - `C18` - investigation stopped at session summaries instead of root cause
   - `C19` - the improved investigation still failed to identify the action that caused the burst
7. `supporting_human_model_signals`:
   - PASS2 explicitly carried that durable guidance should live with the actor that must apply it.
   - PASS2 explicitly carried that quota fears should not trigger action until the whole state is inspected.
   - PASS2 explicitly carried that a workaround that changes the target behavior is not an adequate fix.
   - PASS2 explicitly carried that investigation tooling should recover the actionable human-world cause, not prettier summaries.
8. `counterfactual_prevention_claim`: If applied earlier, this principle would likely have prevented the Boyle stop impulse, the Rustfire wrong-seam fix, the AGENTS.md bloat response, and the summary-only burst investigation repairs.
9. `scope_and_non_goals`: This does not require exhaustive forensics for every trivial edit. It applies when acting at the wrong layer could hide the real cause, create false closure, or keep the real decision owner blind.
10. `pre_action_question`: `Am I changing or explaining the layer that actually owns the problem, or am I taking an easier action on a nearby proxy?`
11. `operational_check`: `For debugging and investigation plans, require an explicit statement of the real owner, the full-state evidence checked, preserved scenario equivalence, and the causal action or mechanism to be explained.`
12. `confidence`: `strong`

## Rejected or merged principle candidates and why

1. `candidate_statement`: `Keep a separate principle for mockup-visible shell, iconography, and card-composition fidelity.`
   - `status`: `merged`
   - `reason`: Those misses share the same decision point and human standard as the typography, clipping, and product-read events: the surface still reads wrong to a human.
   - `merged_into`: `P01`
2. `candidate_statement`: `Keep a separate principle for typography, spacing, clipping, and color readability.`
   - `status`: `merged`
   - `reason`: Readability defects were not a distinct behavioral rule from the other surface-read failures; they are the legibility branch of the same human-facing surface principle.
   - `merged_into`: `P01`
3. `candidate_statement`: `Keep a separate principle for status and dashboard truth before polish.`
   - `status`: `merged`
   - `reason`: Ambiguous state labels and false-visible metrics are the semantic branch of the same surface-truth rule kept in `P01`.
   - `merged_into`: `P01`
4. `candidate_statement`: `Prompt or doc edits do not count as progress.`
   - `status`: `merged`
   - `reason`: This is one proxy-closure pattern inside the broader `verify the actual target first` rule.
   - `merged_into`: `P02`
5. `candidate_statement`: `Require real-machine or live-build proof before claiming a fix.`
   - `status`: `merged`
   - `reason`: This is the concrete closure gate expressed by `P02`, not a distinct principle.
   - `merged_into`: `P02`
6. `candidate_statement`: `Create dedicated fidelity-owner lanes when visible misses recur.`
   - `status`: `merged`
   - `reason`: Owner-lane creation is an operationalization of routing responsibility to the real decision seam, not a separate durable rule.
   - `merged_into`: `P03`
7. `candidate_statement`: `Preserve the exact hero-control composition by rebuilding it from HTML instead of treating it like a loose icon.`
   - `status`: `rejected`
   - `reason`: Real evidence for `P01`, but too local to one control and one design moment to keep as its own durable principle.
8. `candidate_statement`: `Home must always use three distinct cards in the order Recent Activity -> People Setup -> Waiting to upload.`
   - `status`: `rejected`
   - `reason`: The event shows that ordering can carry product meaning, but this exact statement is too surface-specific to keep as a durable cross-task principle.

## The smallest recommended principle set for this scope

- `P01` - human-facing surfaces must communicate truthful meaning readably
- `P02` - closure claims must be proven on the actual target
- `P03` - investigate and fix at the real causal seam

Smallest honest count: `3`

## Principles still too weak and need more days or more events

- A standalone `owner-lane creation` principle still looks too operational and is better treated as one implementation of `P03` until more days show the same pattern.
- A standalone `information architecture and card ordering` principle still looks too close to ordinary design iteration without more repeated evidence beyond `C13`.
- A standalone `hero control composition` principle still looks too local to one surface; `C03` is better preserved as supporting evidence inside `P01`.

## Transcript windows reopened during PASS3 and why

- none; the rerun PASS2 artifact provided enough boundary and human-signal detail to merge aggressively without reopening raw April 2 transcripts
