# INTERVENTION-PASS2

Source day: `2026-04-03`

Primary input artifact: [INTERVENTION-PASS1.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-03/INTERVENTION-PASS1.md)

## Source scope analyzed

- PASS1 candidate source: [INTERVENTION-PASS1.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-03/INTERVENTION-PASS1.md)
- Incident corpus contract: [README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/README.md)
- Incident schema snapshot: [INCIDENT.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/INCIDENT.schema.json)
- Raw transcript rereads:
- `T10` = [rollout-2026-04-03T10-11-16-019d53af-14f1-7a80-987d-cda22ee69505.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/03/rollout-2026-04-03T10-11-16-019d53af-14f1-7a80-987d-cda22ee69505.jsonl)
- `T11` = [rollout-2026-04-03T11-35-01-019d53fb-c3db-7e61-8adf-a4ff171c0de9.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/03/rollout-2026-04-03T11-35-01-019d53fb-c3db-7e61-8adf-a4ff171c0de9.jsonl)
- `T12` = [rollout-2026-04-03T12-28-10-019d542c-6ab7-7dc2-833c-5acc5a904748.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/03/rollout-2026-04-03T12-28-10-019d542c-6ab7-7dc2-833c-5acc5a904748.jsonl)
- `T20` = [rollout-2026-04-03T20-51-56-019d55f9-a32a-7241-bb7b-2fa9afd3b849.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/03/rollout-2026-04-03T20-51-56-019d55f9-a32a-7241-bb7b-2fa9afd3b849.jsonl)
- `T22` = [rollout-2026-04-03T22-52-48-019d5668-49a1-7563-88c0-47b562ff8120.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/03/rollout-2026-04-03T22-52-48-019d5668-49a1-7563-88c0-47b562ff8120.jsonl)
- Directly implicated durable docs reread for expected-state checks:
- Crystallize task surface: [TASK.md](/c:/Agent/Crystallize/Tracking/Task-0019/TASK.md#L23), [PLAN.md](/c:/Agent/Crystallize/Tracking/Task-0019/PLAN.md#L18), [HANDOFF.md](/c:/Agent/Crystallize/Tracking/Task-0019/HANDOFF.md#L17), [GENERAL-DESIGN.md](/c:/Agent/Crystallize/Design/GENERAL-DESIGN.md#L1008)
- Shared design prompts: [GENERAL-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/GENERAL-DESIGNER.md#L7), [INTERFACE-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/INTERFACE-DESIGNER.md#L7)
- Shared debug prompts and process docs: [DEBUG-WORKER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/DEBUG-WORKER.md#L54), [DEBUG-LEADER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/DEBUG-LEADER.md#L76), [DEBUGGING.md](/c:/Users/gregs/.codex/Orchestration/Processes/DEBUGGING.md#L151)
- ThirdPerson task surface: [TASK.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/TASK.md#L7), [PLAN.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/PLAN.md#L3), [HANDOFF.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/HANDOFF.md#L28)
- Task-0021 icon concept: [ICON-COMP-BEACON-FACET-0001.png](/c:/Agent/Crystallize/Tracking/Task-0021/Design/ICON-COMP-BEACON-FACET-0001.png)

## Candidate ids analyzed

- `C01`
- `C02`
- `C03`
- `C04`
- `C05`
- `C06`
- `C07`
- `C08`
- `C09`
- `C10`
- `C11`
- `C12`
- `C13`

## Candidate boundary corrections relative to PASS1

- No candidate ids were added or dropped. The pass still analyzes `C01` through `C13`.
- `C02` stays separate from `C01`, but only as the second half of the same sandbox correction arc. I kept it separate because the human changes the adequacy bar from "do not block prompts" to "keep one shared `.codex` and still test elevated self-sufficiency."
- `C03` stays one long contract-correction event running from the first zero-touch question through the later "dropping the ball" complaint. `C04` starts only when the human moves from repo/task contract to the deeper claim that Codex lacks a human model.
- `C10` stays one event rather than three. The turns at `T22:L2708`, `T22:L3060`, `T22:L3302`, and `T22:L3315` all correct the same implementation drift: wrong ownership boundaries and CMC/Mover boundaries failing to stick in code.

## Per-event analysis records

### C01 - Deleting the Windows sandbox block created a blocker instead of a safe Windows answer

- `event_id`: `C01`
- `title`: Deleting the Windows sandbox block created a blocker instead of a safe Windows answer
- `session_or_thread`: `019d53af-14f1-7a80-987d-cda22ee69505` in `c:\Agent\Crystallize`
- `transcript_path`: `T10`
- `primary_refs`: `T10:L214`, `T10:L228`, `T10:L235`, `T10:L238`
- `ai_course`: The assistant diagnosed the helper/UAC prompt as caused by `[windows] sandbox = "elevated"`, removed that block from user config, verified the line was gone, and reported that future sessions should stop asking for the helper.
- `human_intervention`: The human came back immediately after restoring the block and said the change had blocked later prompts, then explicitly asked for research "instead of running me into a blocker."
- `adequate_outcome`: Prove the actual Windows sandbox contract first, or use a one-off read-only sandbox test, before changing a machine-level config that can stop Codex from accepting prompts.
- `event_boundary_notes`: Kept local to the remove-config then restore-config arc. The later mode-selection discussion is tracked as `C02`.
- `human_model_signal`: Explicit signal: "do some research here instead of running me into a blocker." The human is treating continued tool usability as the first safety bar for any config advice.
- `failure_family_hypothesis`: `verification_proof` by way of acting on an unproven causal theory in a high-impact configuration surface
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Immediate workflow blockage, manual rollback work, and trust damage on future config-changing advice
- `local_lesson_hypothesis`: Do not mutate critical runtime config on a guessed explanation when a reversible read-only probe can test the theory first.
- `cluster_hints`: `closure-truth`, `operator-boundary`, `real-world completion mismatch`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: The earlier reasoning path before `T10:L214` was not reread in full; the event diagnosis rests on the actual change, the blocker report, and the immediate correction arc.

### C02 - The repaired sandbox advice still optimized for the wrong Windows mode

- `event_id`: `C02`
- `title`: The repaired sandbox advice still optimized for the wrong Windows mode
- `session_or_thread`: `019d53af-14f1-7a80-987d-cda22ee69505` in `c:\Agent\Crystallize`
- `transcript_path`: `T10`
- `primary_refs`: `T10:L349`, `T10:L383`, `T10:L406`, `T10:L430`, `T10:L458`, `T10:L465`, `T10:L474`
- `ai_course`: After learning that deleting the block was wrong, the assistant pivoted toward `unelevated` as the cleaner Windows answer, demonstrated that `unelevated` could still see the shared `C:\Users\gregs\.codex` state, and treated that as the practical direction.
- `human_intervention`: The human clarified that one shared `.codex` home mattered because scripts depend on it, then explicitly redirected the test target: "No run the elevated one. I will probably want that level of self-sufficiency on this machine."
- `adequate_outcome`: Optimize for the human's actual workflow contract: shared `C:\Users\gregs\.codex`, no silent split home, and honest evidence about whether elevated mode remains viable on this machine.
- `event_boundary_notes`: Kept separate from `C01` because the adequacy bar changes. `C01` is "do not block me." `C02` is "test the mode I actually care about."
- `human_model_signal`: Explicit signal: one shared `.codex` matters to workflow because scripts consume those logs and sessions, and elevated self-sufficiency still matters enough to test directly.
- `failure_family_hypothesis`: `human_world` via optimizing for the technically neat alternative instead of the human's stated workflow constraints
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Repeated clarification burden and drift toward the wrong long-term operating mode
- `local_lesson_hypothesis`: Once the human states the real operating constraint, test that target directly instead of optimizing for the cleaner substitute.
- `cluster_hints`: `operator-boundary`, `shared-state contract`, `wrong-target repair`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: This may collapse into `C01` in a later accepted-incident pass because both belong to the same sandbox/UAC arc.

### C03 - Zero-touch self-hosting had to be forced into the durable Crystallize surface

- `event_id`: `C03`
- `title`: Zero-touch self-hosting had to be forced into the durable Crystallize surface
- `session_or_thread`: `019d53fb-c3db-7e61-8adf-a4ff171c0de9` in `c:\Agent\Crystallize`
- `transcript_path`: `T11`
- `primary_refs`: `T11:L275`, `T11:L337`, `T11:L357`, `T11:L432`, `T11:L452`, `T11:L597`
- `ai_course`: The assistant could answer that the recommended deployed shape was always-on API plus always-on worker, but the task and repo surfaces initially only implied that expectation. It then patched the durable task docs in stages as the human kept sharpening the requirement.
- `human_intervention`: The human first asked whether "deploy once, then don't touch the box" was actually clear in the task surface, then required the expectation to be put in durable docs, then added the Windows always-on requirement, and finally said there was "some dropping the ball" because the supported lane still was not being modeled as "run a single script" simplicity.
- `adequate_outcome`: The durable task and product surfaces should make the supported self-hosting promise explicit: deploy once, leave one always-on API plus one always-on worker running, avoid routine manual draining, and target one supported action per human-operated job. That is now explicit in [TASK.md](/c:/Agent/Crystallize/Tracking/Task-0019/TASK.md#L27), [PLAN.md](/c:/Agent/Crystallize/Tracking/Task-0019/PLAN.md#L23), [HANDOFF.md](/c:/Agent/Crystallize/Tracking/Task-0019/HANDOFF.md#L27), and [GENERAL-DESIGN.md](/c:/Agent/Crystallize/Design/GENERAL-DESIGN.md#L1013).
- `event_boundary_notes`: Kept as one multi-turn contract-correction arc because all turns are about making the same product/task promise explicit and durable. `C04` starts only when the human moves up a level and diagnoses the missing human model itself.
- `human_model_signal`: Explicit signals: "My goal is to not touch that machine except for the original deployment" and later "I don't want to do anything more complicated than run a single script."
- `failure_family_hypothesis`: `workflow_orchestration` because the durable task surface under-modeled the intended operator lane and left too much room for babysitting
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Repeated clarification work, risk that the product would ship a burdensome self-hosting lane, and lost trust that "hands-off" wording really meant hands-off operation
- `local_lesson_hypothesis`: If the human keeps restating the same "don't make me babysit this box" requirement, encode the operator contract explicitly in durable docs rather than leaving it as implied product intent.
- `cluster_hints`: `operator-boundary`, `real-world completion mismatch`, `one-action baseline`, `human-step-away`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: No major factual ambiguity. The only real judgment call is how much of the long arc to keep with `C03` versus treating the later meta-level diagnosis as `C04`.

### C04 - The humane-design fix had to move up to an explicit human model

- `event_id`: `C04`
- `title`: The humane-design fix had to move up to an explicit human model
- `session_or_thread`: `019d53fb-c3db-7e61-8adf-a4ff171c0de9` in `c:\Agent\Crystallize`
- `transcript_path`: `T11`
- `primary_refs`: `T11:L607`, `T11:L641`, `T11:L667`, `T11:L677`
- `ai_course`: The assistant initially repaired the design prompts by layering more embodied-use and burden-budget heuristics into the shared prompt set, then showed those edits back to the human.
- `human_intervention`: The human said the deeper problem is that "codex doesn't understand people by default," then pushed for a distilled description of "what humans are like" so designers can infer click-path, babysitting, and hidden-state expectations from first principles rather than from an ever-growing checklist.
- `adequate_outcome`: Put the compact human baseline at the top of the shared design prompt layer and let downstream heuristics derive from it. That now exists in [GENERAL-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/GENERAL-DESIGNER.md#L7), [GENERAL-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/GENERAL-DESIGNER.md#L9), and is inherited by [INTERFACE-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/INTERFACE-DESIGNER.md#L7).
- `event_boundary_notes`: Separate from `C03` because the intervention changes level. `C03` is about the Crystallize operator contract. `C04` is about the shared model of humans used to derive many such contracts.
- `human_model_signal`: Explicit signal: "codex doesn't understand people by default" and "distill 'what humans are like'." The human also supplied the candidate model: shortest path, minimal clicks, minimal manual intervention, and a preference for systems that "just work."
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Hours of repeated explanation, prompt bloat from piling on symptoms, and continued misses on human-facing surfaces if the root model stayed absent
- `local_lesson_hypothesis`: When many humane-design rules are really consequences of the same underlying human baseline, encode that baseline explicitly instead of accreting downstream reminders forever.
- `cluster_hints`: `human-model baseline`, `operator-boundary`, `real-world completion mismatch`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: No major factual ambiguity. The remaining open question is only how widely a later clustering pass will generalize this human-baseline pattern across other days.

### C05 - Rustfire debugging was redirected from intuition-first tuning to ruthless causal narrowing

- `event_id`: `C05`
- `title`: Rustfire debugging was redirected from intuition-first tuning to ruthless causal narrowing
- `session_or_thread`: `019d542c-6ab7-7dc2-833c-5acc5a904748` in `c:\EHG_GregS_main`
- `transcript_path`: `T12`
- `primary_refs`: `T12:L429`, `T12:L436`
- `ai_course`: The active Rustfire branch had just finished a bandwidth-cap experiment and was still operating at the level of broad configuration hypotheses and "disprove one global cause" style narrowing.
- `human_intervention`: The human redirected the method explicitly: reveal information that narrows the bug space in the most efficient way, preferably with callstacks, breakpoints, or verbose logging that trace the cause upstream from the concrete evidence until root cause is reached.
- `adequate_outcome`: The debug path should anchor on the first concrete disagreement seam and choose the highest-information check near that seam before trying more global explanations.
- `event_boundary_notes`: Short single-turn correction arc. `C06` is separate because it applies the same principle one layer later to the shared debug docs themselves.
- `human_model_signal`: Explicit signal: root-cause work should prioritize the highest-information evidence path, not broad speculation. The human named callstacks, breakpoints, and verbose logging as the preferred tools.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Wasted debug cycles on low-information hypotheses and slower convergence on the actual fault source
- `local_lesson_hypothesis`: When a live bug already exposes a concrete log or behavior seam, start from the seam with the best causal information gain rather than from broad environmental guesses.
- `cluster_hints`: `wrong-seam debugging`, `diagnostic_scope`, `root-cause first`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: The pre-`L429` debug exploration was only reread in bounded form, but the correction itself is explicit and unambiguous.

### C06 - The updated debug docs were still one level too vague

- `event_id`: `C06`
- `title`: The updated debug docs were still one level too vague
- `session_or_thread`: `019d53fb-c3db-7e61-8adf-a4ff171c0de9` in `c:\Agent\Crystallize`
- `transcript_path`: `T11`
- `primary_refs`: `T11:L1330`, `T11:L1395`, `T11:L1405`, `T11:L1414`
- `ai_course`: The assistant dry-ran the newly updated debug method against the Rustfire handoff and got a better answer than before, but the result still stopped at abstract "write path vs reader path" reasoning rather than tracing the named variables in the fault message itself.
- `human_intervention`: The human said "that's better" but still not right, asked whether the stronger inference was already available from code-read, and then requested an explicit rule to trace the fault/assert/log message's named values back through their writers while deleting lower-level predicates that were only derived restatements.
- `adequate_outcome`: Shared debug guidance should explicitly tell the worker to trace named values through the code and state that write or advance them. That rule now exists in [DEBUG-WORKER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/DEBUG-WORKER.md#L59), [DEBUG-LEADER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/DEBUG-LEADER.md#L76), and [DEBUGGING.md](/c:/Users/gregs/.codex/Orchestration/Processes/DEBUGGING.md#L166).
- `event_boundary_notes`: Kept as one event from the first "still one level too vague" complaint through the value-tracing rewrite and re-test. It is distinct from `C05` because this is about the prompt contract, not just one live debug branch.
- `human_model_signal`: Explicit signal: instruction following can pessimize if too many lower-level predicates pile up, and the correct rule for named fault values is provenance tracing rather than generic branch wandering.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: A shared debug prompt that still lands one level too vague would keep producing slower, lower-information debug branches across tasks
- `local_lesson_hypothesis`: If a fault or warning names concrete variables, encode a direct provenance-tracing rule and prune redundant heuristics that distract from that seam.
- `cluster_hints`: `wrong-seam debugging`, `named-value tracing`, `prompt-overconstraint`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: No major factual ambiguity. The main later question is whether `C05` and `C06` cluster as one debugging-method family or remain distinct incident records.

### C07 - Task-0001 planning was rejected as too narrow for the real automation goal

- `event_id`: `C07`
- `title`: Task-0001 planning was rejected as too narrow for the real automation goal
- `session_or_thread`: `019d5668-49a1-7563-88c0-47b562ff8120` in `c:\Agent\ThirdPerson`
- `transcript_path`: `T22`
- `primary_refs`: `T22:L135`, `T22:L171`, `T22:L174`
- `ai_course`: The first planning package expanded detail, UnLua analysis, and API shape, but still modeled `Task-0001` too much like a narrow observability-first vertical slice.
- `human_intervention`: The human said the plan was still too narrow, clarified that a rich automation layer was needed so the human could step away, and made scripting, input, navigation or execution, observability, capture, and queryability non-negotiable.
- `adequate_outcome`: The durable task and plan should define `Task-0001` as a broad automation-foundation task rather than a thin slice. That correction is now explicit in [TASK.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/TASK.md#L13), [TASK.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/TASK.md#L79), [PLAN.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/PLAN.md#L3), and [PLAN.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/PLAN.md#L20).
- `event_boundary_notes`: Kept separate from `C08` because `C07` is the scope reset itself. `C08` is the later architecture-boundary tightening after that broader scope had already landed.
- `human_model_signal`: Explicit signal: the automation layer should be rich enough that "the human can step away." The human also named the concrete foundation surfaces that must exist.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: A too-small first task would re-open the same missing-foundation work later and keep the human in the loop for runtime automation chores
- `local_lesson_hypothesis`: When a task is supposed to reduce future orchestration drag, optimize for a reusable foundation, not for the smallest demo that happens to be easy to plan.
- `cluster_hints`: `automation-foundation`, `human-step-away`, `scope-too-narrow`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: No major factual ambiguity. The only open synthesis question is how broadly this "foundation, not thin slice" pattern repeats across repos.

### C08 - The broadened Task-0001 plan still had to be corrected to Mover-first, minimal, and script-owned

- `event_id`: `C08`
- `title`: The broadened Task-0001 plan still had to be corrected to Mover-first, minimal, and script-owned
- `session_or_thread`: `019d5668-49a1-7563-88c0-47b562ff8120` in `c:\Agent\ThirdPerson`
- `transcript_path`: `T22`
- `primary_refs`: `T22:L246`, `T22:L297`, `T22:L338`
- `ai_course`: After the scope broadening landed, the plan still left room for base `Character` and `CharacterMovementComponent` assumptions, treated snapshots too much like a first-class C++ surface, and framed per-pass UnLua regression more rigidly than the human wanted.
- `human_intervention`: The human required Mover instead of base Characters, clarified that snapshots should stay optional and script-composed where possible, and softened the earlier per-pass regression insistence into "do it if helpful, otherwise keep it simple while still covering the key areas."
- `adequate_outcome`: The plan should keep the C++ API minimal, use Mover-backed execution only, let UnLua compose snapshots from primitives, and treat early proof as honest pass-local evidence rather than rigid ritual. Those boundaries are now explicit in [PLAN.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/PLAN.md#L140), [PLAN.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/PLAN.md#L169), [PLAN.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/PLAN.md#L289), [PLAN.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/PLAN.md#L304), and summarized in [HANDOFF.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/HANDOFF.md#L33).
- `event_boundary_notes`: Separate from `C07` because the broader scope correction had already landed. This event is about which substrate and ownership boundaries are acceptable inside that broader scope.
- `human_model_signal`: Explicit signals: "Anything that could go in script should go in script" and "I see snapshots as optional." The human is pushing for minimal core surface plus script-side composition.
- `failure_family_hypothesis`: architecture-boundary drift toward heavy C++ ownership and the wrong movement stack, closest schema fit `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Risk of calcifying the foundation around the wrong movement substrate, too much C++ surface area, and overly ceremonial proof rules
- `local_lesson_hypothesis`: After scope is broadened, immediately re-check movement stack, ownership boundaries, and API minimality so the foundation does not silently get heavier and less script-friendly.
- `cluster_hints`: `architecture-boundary drift`, `Mover-only`, `script-vs-core boundary`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: No major factual ambiguity. The only open question is whether later clustering keeps this distinct from `C10` or treats it as the planning half of the same architecture-boundary family.

### C09 - Repeated git account-picker popups forced an auth-path detour

- `event_id`: `C09`
- `title`: Repeated git account-picker popups forced an auth-path detour
- `session_or_thread`: `019d5668-49a1-7563-88c0-47b562ff8120` in `c:\Agent\ThirdPerson`
- `transcript_path`: `T22`
- `primary_refs`: `T22:L1673`, `T22:L1685`, `T22:L1688`, `T22:L1704`
- `ai_course`: The assistant was doing normal commit and push work during PASS-0003/PASS-0004 transition, and the networked Git path kept invoking Git Credential Manager on an HTTPS remote, which surfaced a Windows account picker each time.
- `human_intervention`: The human stopped implementation and said the repeated popup was "obviously not scalable," forcing the assistant to pause feature work and diagnose the git auth path instead.
- `adequate_outcome`: The repo should have a deterministic push/auth path that does not require repeated human account selection during normal agent work.
- `event_boundary_notes`: Short interruption event. Kept separate from surrounding implementation work because the human explicitly suspended task progress to deal with the auth burden.
- `human_model_signal`: Explicit signal: repeated manual account selection inside a normal automation loop is not acceptable operator burden.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Repeated popup fatigue, stalled implementation, and hidden dependence on manual account picking
- `local_lesson_hypothesis`: If a support command repeats inside the normal work loop, stabilize its auth path before continuing mainline task work.
- `cluster_hints`: `operator-boundary`, `hidden-interactive-step`, `supporting-tool friction`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: This is clearly a real intervention, but it is less obviously part of the durable human-course-correction corpus than the stronger scope, proof, and human-model events.

### C10 - PASS-0004 implementation drifted back into pawn/class ownership and CMC assumptions

- `event_id`: `C10`
- `title`: PASS-0004 implementation drifted back into pawn/class ownership and CMC assumptions
- `session_or_thread`: `019d5668-49a1-7563-88c0-47b562ff8120` in `c:\Agent\ThirdPerson`
- `transcript_path`: `T22`
- `primary_refs`: `T22:L2708`, `T22:L3060`, `T22:L3302`, `T22:L3315`
- `ai_course`: During PASS-0004 implementation and local experiments, the assistant kept sliding back toward `GameAutomation` owning a pawn structure, toward a parallel `CharacterMovementComponent` path in `GameAutomation`, and then toward a project-side Character-style wrapper in `ThirdPerson`.
- `human_intervention`: The human repeatedly reasserted the architecture boundary: `GameAutomation` should not enforce actor or pawn structure, it may assume required components such as `UMoverComponent`, `GameAutomation` does not support CMC in v1, the project module should not use CMC either, and the probe pawn should be renamed to the game-owned `AThirdPersonPawn`.
- `adequate_outcome`: Implementation should stay Mover-only, leave actor ownership to the game module, and keep future extension script-first. The later durable task surface reflects that in [HANDOFF.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/HANDOFF.md#L130), [HANDOFF.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/HANDOFF.md#L148), and [HANDOFF.md](/c:/Agent/ThirdPerson/Tracking/Task-0001/HANDOFF.md#L182).
- `event_boundary_notes`: Treated as one multi-stage event because all four corrections point at the same underlying miss: the planning-stage architecture boundaries were not sticking during code work.
- `human_model_signal`: Explicit signals: "`GameAutomation` should not enforce an actor or pawn structure," "`GameAutomation` doesn't support CMC in v1," and "only Mover" at project level too.
- `failure_family_hypothesis`: architecture-boundary drift toward the wrong ownership layer and unsupported movement stack, closest schema fit `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Rework during implementation, increased risk that the foundation would calcify around the wrong substrate, and evidence that earlier planning corrections were not reliably constraining code decisions
- `local_lesson_hypothesis`: Once architecture boundaries are stated as hard task constraints, every implementation experiment should be checked against them before the code path grows or gets named as durable surface.
- `cluster_hints`: `architecture-boundary drift`, `Mover-only`, `plan-to-implementation nonstickiness`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: A later accepted-incident writeup may want a wider reread of the concrete code diffs to show exactly how the drift manifested in files, not just in chat corrections.

### C11 - The Android launcher icon was visibly too zoomed in against the approved concept

- `event_id`: `C11`
- `title`: The Android launcher icon was visibly too zoomed in against the approved concept
- `session_or_thread`: `019d55f9-a32a-7241-bb7b-2fa9afd3b849` in `c:\Agent\ModelGeneration`
- `transcript_path`: `T20`
- `primary_refs`: `T20:L645`, `T20:L648`
- `ai_course`: The assistant had just declared the approved `Beacon Facet` launcher icon implemented from source assets.
- `human_intervention`: The human rejected the actual phone result, said it looked "way too zoomed in," and required screenshot-backed fixing so the phone icon would match the concept PNG's proportions against the icon border.
- `adequate_outcome`: The proof bar for launcher icon geometry is the real phone render compared against [ICON-COMP-BEACON-FACET-0001.png](/c:/Agent/Crystallize/Tracking/Task-0021/Design/ICON-COMP-BEACON-FACET-0001.png), not the source asset alone.
- `event_boundary_notes`: Kept separate from `C12` and `C13` because this first event is specifically about proportional crop mismatch on the phone.
- `human_model_signal`: No broad principle stated beyond the concrete adequacy bar: the phone icon should match the concept PNG's proportions versus the border.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Visible device-level defect on the actual surface the human uses
- `local_lesson_hypothesis`: For launcher art, on-device screenshot comparison is the closure bar; source-asset correctness is only intermediate evidence.
- `cluster_hints`: `proof-must-reach-device`, `human-facing surface mismatch`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: This may remain ordinary iterative design review rather than a durable incident, depending on how later passes distinguish normal visual iteration from stronger closure-truth misses.

### C12 - A claimed icon update still did not look different on the actual phone

- `event_id`: `C12`
- `title`: A claimed icon update still did not look different on the actual phone
- `session_or_thread`: `019d55f9-a32a-7241-bb7b-2fa9afd3b849` in `c:\Agent\ModelGeneration`
- `transcript_path`: `T20`
- `primary_refs`: `T20:L981`, `T20:L1018`, `T20:L1021`, `T20:L1050`
- `ai_course`: After changing the icon highlight and app label, rebuilding, and reinstalling, the assistant was still narrating the update as if the launcher surface should now show the new result.
- `human_intervention`: The human said the icon "doesn't look any different" on the home page, which forced the assistant to stop inferring from the APK and instead pull the actual screenshot, identify launcher caching, and then choose the safe cache-refresh path that did not touch app data.
- `adequate_outcome`: Visible launcher changes should be validated against the live phone surface, and cache-refresh fixes should preserve data when possible.
- `event_boundary_notes`: Kept separate from `C11` because the specific miss here is not the initial crop issue but a false visible-completion story after an update had supposedly landed.
- `human_model_signal`: Implicit signal reinforced by the later safe-path choice: visible home-screen state outranks repo assumptions, and no-data repair matters.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: False closure on a visible device change and risk of unnecessary destructive remediation if the cache theory had not been tested carefully
- `local_lesson_hypothesis`: When the user says the live surface still looks unchanged, stop narrating from build outputs and inspect the actual device state before choosing the repair path.
- `cluster_hints`: `proof-must-reach-device`, `closure-truth`, `safe-state repair`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: The event is real, but it may stay below accepted-incident threshold as ordinary device iteration unless later clustering emphasizes premature visible-closure claims.

### C13 - The on-device product name still surfaced as `MemoryRecorder`

- `event_id`: `C13`
- `title`: The on-device product name still surfaced as `MemoryRecorder`
- `session_or_thread`: `019d55f9-a32a-7241-bb7b-2fa9afd3b849` in `c:\Agent\ModelGeneration`
- `transcript_path`: `T20`
- `primary_refs`: `T20:L1301`, `T20:L1304`
- `ai_course`: The assistant had already treated application-level label changes as evidence that the app now surfaced as `Crystallize`.
- `human_intervention`: The human checked search on the actual device and reported that the app still only appeared as `MemoryRecorder`, which forced the assistant to investigate the launcher-activity label rather than relying on application-level metadata alone.
- `adequate_outcome`: The searchable and launchable on-device surface should actually present `Crystallize`, not just the application label inside the APK.
- `event_boundary_notes`: Kept separate from `C12` because the core miss changed from cached home-screen shortcut state to the deeper question of what the launcher activity itself advertises to the system.
- `human_model_signal`: No explicit generalized principle beyond the concrete human-facing bar: the phone should surface `Crystallize`, not `MemoryRecorder`.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Misleading product identity on the actual device surface the human uses to find and launch the app
- `local_lesson_hypothesis`: Renaming is not complete until the launchable/searchable device surface changes, not merely until repo-side manifest or APK metadata looks correct.
- `cluster_hints`: `proof-must-reach-device`, `real-world completion mismatch`, `human-facing surface mismatch`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: This may remain ordinary device polish unless later clustering treats it as part of a broader "repo proof versus actual device surface" family with `C11` and `C12`.

## Likely accepted incidents

- `C01` Deleting the Windows sandbox block created a blocker instead of a safe Windows answer
- `C03` Zero-touch self-hosting had to be forced into the durable Crystallize surface
- `C04` The humane-design fix had to move up to an explicit human model
- `C05` Rustfire debugging was redirected from intuition-first tuning to ruthless causal narrowing
- `C06` The updated debug docs were still one level too vague
- `C07` Task-0001 planning was rejected as too narrow for the real automation goal
- `C08` The broadened Task-0001 plan still had to be corrected to Mover-first, minimal, and script-owned
- `C10` PASS-0004 implementation drifted back into pawn/class ownership and CMC assumptions

## Likely non-incident but still important intervention events

- `C09` Repeated git account-picker popups forced an auth-path detour
- `C11` The Android launcher icon was visibly too zoomed in against the approved concept
- `C12` A claimed icon update still did not look different on the actual phone
- `C13` The on-device product name still surfaced as `MemoryRecorder`

## Repeated cluster hints noticed across the analyzed set

- `operator-boundary` appears repeatedly in `C01`, `C02`, `C03`, `C04`, and `C09`: the human keeps rejecting hidden or repeated operator work and keeps asking for one clear supported action rather than babysitting.
- `real-world completion mismatch` appears in `C01`, `C03`, `C11`, `C12`, and `C13`: the assistant's current story about "done" or "good enough" is weaker than the real machine, device, or workflow state the human actually has to live with.
- `wrong-seam debugging` appears in `C05` and `C06`: the human keeps pushing the method down to the first concrete disagreement seam and away from broad theory.
- `architecture-boundary drift` appears in `C08` and `C10`: even after the plan says Mover-only and script-minimal, implementation pressure keeps trying to reintroduce heavier or wrong-layer ownership.
- `human-step-away` appears in `C03`, `C04`, `C07`, and `C08`: the human repeatedly defines adequacy as a state where the product or foundation works without constant human supervision.

## Strongest human-model signals worth carrying into a later clustering or principle pass

- `C04`: "codex doesn't understand people by default" and the proposed "what humans are like" baseline are the clearest direct statement of the missing model.
- `C03`: "My goal is to not touch that machine except for the original deployment" and later "I don't want to do anything more complicated than run a single script" give a concrete adequacy bar for self-hosting.
- `C05`: Debugging should reveal the most information in the most efficient way, preferably by tracing concrete evidence upstream with debugger-grade or targeted logging evidence.
- `C06`: When a fault names values, the right next step is to trace those named values through their writers or updaters; too many derived predicates can pessimize instruction following.
- `C07`: The foundation should be rich enough that the human can step away; do not optimize for a tiny slice when the real need is durable automation substrate.
- `C08`: "Anything that could go in script should go in script" and "snapshots are optional" are crisp statements of the desired ownership boundary.
- `C01` and `C02`: Keep Codex usable, keep one shared `.codex` state, and test the mode the human actually intends to use rather than the cleaner substitute.

## Events that still need a wider reread

- No wider reread is required to support this local analysis pass.
- A later accepted-incident writeup should probably widen `C02` if it needs to decide whether that event survives separately from `C01`.
- A later accepted-incident writeup for `C10` should probably reread the surrounding code diffs and task artifacts, not just the transcript turns, so the implementation drift is evidenced in both chat and code.
- If a later pass wants to promote any of `C11` through `C13`, it should widen from the chat windows into the actual screenshot and APK-artifact trail to decide whether those are ordinary device-iteration events or stronger proof/closure incidents.
