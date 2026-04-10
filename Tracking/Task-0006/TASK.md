# Task 0006

## Title

Capture human-facing orchestration incidents as durable evidence.

## Summary

The higher-level orchestration goal is to catch dropped balls, misunderstandings, and process bumps before they have to be escalated manually.

This task is intentionally narrower than that larger goal. It is only about capturing the moments where the human had to step in and disagree with an AI-produced or AI-directed outcome in order to protect the intended human-facing result.

Longer-horizon, this capture work is meant to support a future counterpoint agent that can say some version of `I think the human will want X` as a check against the producer, with the goal of reducing unnecessary human escalation. This task does not build that agent yet; it builds the evidence base honestly enough that such an agent could later be grounded in real human interventions and real human-interest statements.

That future agent is not meant to be a default veto-heavy critic. The intended default stance is advisory:

- warn when a produced outcome appears to violate a protected human interest
- explain the likely conflict concretely
- let the producer decide in most cases
- reserve stronger intervention or escalation for later workflow rules and higher-confidence conflicts

For this task, an `incident` means:

- a divergence from the intended human-facing outcome
- where an AI-produced or AI-directed outcome existed before the intervention
- where the human explicitly intervened to reject, correct, or redirect that outcome
- and where that intervention should become durable evidence instead of disappearing into chat history

The seed research for this task shows that these divergences are not all the same. Some are perceptual or interface-read failures, some are human-world or dignity failures, some are taste failures, and some come from orchestration or prompt gaps that let the miss through. This task should capture the incident first so later work can learn from it honestly.

## Goals

- Define a durable incident contract for human course correction incidents.
- Preserve the human's correction path, not just the final corrected artifact.
- Preserve the concrete pre-correction state, including active course when the incident is about drift rather than just a static artifact snapshot.
- Capture the human-facing expected state, the actual state, and why the gap mattered.
- Trace each incident upward through one or more ordered `why_chains`, where each chain is internally linear and each later entry answers why the prior entry mattered to the human.
- Distinguish the main incident layers when useful, such as:
  - perception or interface-read failure
  - human-world or dignity failure
  - taste or stylistic mismatch
  - orchestration, prompt, or workflow failure
- Make incidents concrete enough that later prompts, skills, evals, or training work can learn from them.
- Keep incidents tied to tasks, artifacts, screenshots, prompts, or concrete outputs rather than vague complaints.
- Separate `fix the artifact now` from `encode the durable process or prompt repair for next time`.
- Allow a quick first-pass incident capture, followed by a second-pass root-cause refinement once the underlying mechanism is better grounded.
- Make accepted incident JSONs heavyweight enough to carry their own verbatim human timeline and transcript context.
- Preserve enough surrounding daily-message context that later work can learn not only from explicit corrections, but also from non-corrective statements that reveal durable human interests, human limits, or human-world constraints.
- Preserve enough evidence that later work can distinguish recurring protected-interest families such as:
  - state-story truth
  - real-world done
  - human-facing form factor
  - control-boundary ownership

## Non-Goals

- Building the full autonomous critic stack, reviewer swarm, or "Jarvis" loop in this task.
- Solving every captured incident automatically.
- Replacing bug tracking for defects that do not involve explicit human disagreement over an AI-produced or AI-directed outcome.
- Turning this task into a general taste-library or model-training task.
- Treating every user preference or local polish comment as an incident.
- Building the counterpoint agent itself in this task.

That said, do not flatten screenshot-visible mockup distortions into `mere polish` when the human is clearly correcting product-read quality such as typography, spacing, icon fidelity, hierarchy, clipping, or similar human-facing semantic distortion.

## Constraints And Baseline

- The trigger for an incident is explicit human disagreement over an AI-produced or AI-directed outcome, not generic model uncertainty.
- The daily review corpus may still contain non-corrective human-interest statements that never qualify as incidents but are worth preserving for later counterpoint-agent work.
- A record is not a qualified incident unless it preserves:
  - the grounded pre-correction state
  - the human intervention itself
- The outcome bar is the proper human-facing result in the ordinary human sense, not a weaker technical proxy.
- The seed research in [Scaling-Agent-Orchestration.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Scaling-Agent-Orchestration.md) surfaces at least four recurring miss families:
  - perceptual or visual failures
  - human-world and dignity failures
  - taste or stylistic failures
  - orchestration or prompt failures
- [VISION-HARVESTER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/VISION-HARVESTER.md) pushes this task to preserve the correction path and the right durable repair layer instead of only repainting the local artifact.
- The seed research also explicitly recommends climbing from raw feedback toward more general lessons, but this task stops earlier: it should preserve grounded `why_chains` before any later primitive distillation work.
- [GENERAL-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/GENERAL-DESIGNER.md) pushes this task to capture first-value, dignity, hidden-assumption, and human-burden failures rather than flattening them into "design taste."
- [INTERFACE-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/INTERFACE-DESIGNER.md) pushes this task to capture literal communication, scan path, hierarchy, clipping, typography, spacing, icon fidelity, control meaning, placeholder truth, and mockup-fidelity distortion when the incident is interface-facing.
- Generic workflow, prompt, and schema changes should be promoted into `C:\Users\gregs\.codex\Orchestration\` when they are cross-repo and durable.
- Task-local history, examples, and evidence should stay under `Tracking/Task-0006/`.

## Expected Resolution

- The task leaves behind a clear incident definition with scope boundaries.
- A durable incident record shape exists for cases where the human had to step in.
- The process intent is explicit: the corpus is meant to support a future advisory counterpoint agent, not just a bug scrapbook.
- The durable process explicitly allows two-step incident maturation:
  - first-pass capture to avoid losing the correction event
  - second-pass root-cause refinement to tighten the record against the actual mechanism
- The accepted incident JSON itself carries heavyweight verbatim evidence:
  - chronological human timeline
  - contiguous transcript windows
  - original date-time prefixes
  - embedded line wrapping rather than split turns
- That record shape preserves, at minimum:
  - the concrete event-level expected state
  - the concrete event-level actual outcome that triggered disagreement
  - the concrete pre-correction state, including active course when relevant
  - the human intervention summary and evidence
  - why the gap mattered to the human
  - the incident layer or layers involved
  - one or more verbatim `verbatim_transcript_windows`
  - explicit `why_chains` that:
    - keep each rationale path linear
    - allow multiple sibling paths when the human expressed multiple upstream whys
    - progress from the immediate target state in this incident
    - to the broader target state the human was protecting
    - to the highest-level reason the human explicitly stated in the correction evidence for that path
  - a schema-constrained principle category for each entry in each chain
  - evidence references
- The storage contract is explicit:
  - schema and example artifacts stay at `Tracking/Task-0006/Research/`
  - task-local `Tracking/Task-0006/Research/Daily/` is the bootstrap contract and backfill workspace while this task is still refining the shape
  - the promoted shared workflow home lives under `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\`
  - task-local day folders may carry transcript-first `INTERVENTION-PASS1.md`, `INTERVENTION-PASS2.md`, and `INTERVENTION-PASS3.md` analysis artifacts while the day is still being investigated
  - narrower accepted-evidence packaging, if reintroduced later, should be derived from the pass stack rather than treated as the primary workflow
- The system can later use those incident records as input to vision harvesting, humane design review, interface review, prompt repair, or eval generation without pretending those later steps were already solved here.
- The task framing explicitly preserves non-corrective human-interest statements because the future counterpoint agent will need more than objections alone to model what the human wants.
- The task framing also allows a derived daily brief so a human or outside model can review one source day without reconstructing the corpus structure first.

## What Does Not Count

- A vague statement that something "felt off" without preserving the expected state, actual state, and human cost.
- A flat interpretation like "the UI was confusing" without preserving the grounded event and the upward `why_chains`.
- A silent fix where the human correction never becomes a durable incident record.
- A generic bug report that does not preserve the grounded pre-correction state, the human disagreement, and the human-facing outcome gap together.
- A human-reported product defect that never required course-correcting an AI-produced outcome.
- A one-off taste preference note with no evidence that the human had to step in to protect the intended outcome.
- A broader plan for critics, training, or automation that never first captures the incident itself.

## Implementation Home

Keep task-owned artifacts under `Tracking/Task-0006/`.

Promote shared workflow rules, prompt updates, canon docs, eval design, and related durable orchestration docs into `C:\Users\gregs\.codex\Orchestration\`.

If repo-local helpers or experiments are needed for CodexDashboard-specific task flow, keep them repo-local and explicitly justify why they are not shared orchestration material.

Within `Tracking/Task-0006/Research/`:

- keep contract files and example incidents at the top level
- keep bootstrap day-owned capture artifacts under `Tracking/Task-0006/Research/Daily/YYYY-MM-DD/` while the contract is still being refined
- promote stable shared intervention artifacts into `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\`

## Acceptance Criteria

- `Task-0006` defines `incident` explicitly as a human-corrected divergence from the intended human-facing outcome.
- The task scope is clearly limited to capture and classification, not full autonomous resolution.
- The minimum durable incident record shape is defined, including:
  - event-level expected state
  - event-level actual state
  - grounded pre-correction state, including active course when relevant
  - human intervention summary and evidence
  - human cost or reason the divergence mattered
  - layer or classification
  - `why_chains` with linear entries that progressively generalize the target state being protected
  - schema-constrained principle categories for those entries
  - evidence references
- The task owns a concrete task-local incident schema and at least two validated example incidents.
- The task definition distinguishes incidents from:
  - ordinary code bugs
  - pass audits
  - regression runs
  - taste-library entries
  - later prompt or training assets
- The task framing preserves the concern families surfaced by the seed research and the named review prompts:
  - perceptual or interface-read failures
  - human-world and dignity failures
  - taste failures
  - orchestration or workflow failures
- The task framing requires incidents to preserve one or more progressively generalized `why_chains` rather than only a symptom description, and each chain must stop before unsupported inference.
- The task framing makes daily capture review explicit:
  - transcript-first intervention-pass artifacts may exist per day while a day is under review
  - one day folder per incident date
  - accepted incident instances stored in the reports folder matching their `source_date`
- The task framing also distinguishes:
  - strict incidents, which require explicit human correction
  - broader intervention or human-model evidence, which may stay in task-local pass artifacts even when it is not promoted into the accepted incident set
- The task framing also distinguishes first-pass incident capture from later root-cause refinement, so incident opening does not wait on perfect diagnosis.
- The implementation home makes the shared-versus-repo-local split explicit.
- The next honest phase after task creation is research on incident shape, storage, and workflow fit rather than immediate implementation guesswork.

## Open Questions

- Should incidents live only as task-owned markdown artifacts, or also as a shared normalized index or schema-backed store?
- How much historical backfill belongs in this task versus a later follow-on task?
- How long should an individual `why` path usually be before it stops being useful and starts becoming philosophy?
- When should a repeated `why` path pattern be promoted into a reusable primitive or shared rule?

## References

- [Scaling-Agent-Orchestration.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Scaling-Agent-Orchestration.md)
- [VISION-HARVESTER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/VISION-HARVESTER.md)
- [GENERAL-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/GENERAL-DESIGNER.md)
- [INTERFACE-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/INTERFACE-DESIGNER.md)
