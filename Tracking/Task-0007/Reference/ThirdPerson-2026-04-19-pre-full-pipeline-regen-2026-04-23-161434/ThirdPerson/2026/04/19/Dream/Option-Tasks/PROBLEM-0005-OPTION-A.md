# PROBLEM-0005 OPTION-A: Repo-Local First-Disagreement Debugging Gate (Honest Runtime Defect Narrowing)

## Title

Before another fix attempt on a reopened ThirdPerson runtime defect, require `BUG-<NNNN>.md` and `REGRESSION-RUN-<NNNN>.md` to record the first concrete disagreement (with values) and the next upstream writer boundary.

## Summary

The packet shows a repeated debugging failure mode: the work drifts into "tweak mode" (bounded changes without decisive narrowing), forcing the human to repeatedly demand root-cause debugging and durable defect tracking.

The winner proposes a repo-local gate, enforced at the artifact boundary ThirdPerson already uses for runtime truth:

- task-owned `BUG-<NNNN>.md` for the narrowing narrative
- task-owned `REGRESSION-RUN-<NNNN>.md` for executed lane evidence

Before a second (or later) fix attempt, these artifacts must preserve:

- the first concrete disagreement (with values) observed on the human default lane
- the next upstream writer boundary to trace (also concrete, not category-level)

## Writeup Type

Concrete implementation task (burden-reduction proposal).

## Source Event IDs

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

## Burden Being Reduced

The human is currently forced to repeatedly perform "debugging leadership" work:

- demanding root-cause analysis instead of superficial tweaks
- demanding durable defect tracking (explicit `BUG-...` notes) rather than implicit closure
- demanding that evidence remain anchored to the real human default runtime lane, not proxy lanes or export-only artifacts

This is exported cognitive labor. Without a hard gate, the system will continue to propose or implement changes that do not collapse uncertainty, and the human must intervene to restore the narrowing method.

## Current Truth

ThirdPerson already has:

- a repo-local debugging adapter at [ThirdPerson/DEBUGGING.md](/c:/Agent/ThirdPerson/DEBUGGING.md)
- a shared debugging method that explicitly requires narrowing from the "first concrete disagreement" (see [../../../../../../../../Processes/DEBUGGING.md](../../../../../../../../Processes/DEBUGGING.md))
- task-owned artifact homes for runtime bug tracking and regression reruns:
  - `Tracking/Task-<id>/BUG-<NNNN>.md`
  - `Tracking/Task-<id>/Testing/REGRESSION-RUN-<NNNN>.md`

Despite this, the packet shows the human still had to repeatedly demand that the system:

- stop at the first concrete disagreement
- trace writers/updaters upstream one boundary at a time
- preserve contradictions instead of narrating category-level causes

So the current truth is: the method exists, but it is not reliably enforced at the repo’s actual evidence artifact boundary.

## Target Truth

When a runtime defect is reopened or survives one attempted fix, the next work loop must not proceed as "another tweak" until the task-owned artifacts preserve:

- **First Concrete Disagreement**:
  - the exact observed disagreement between expected and actual behavior, including concrete values (frames, world Z clearances, transforms, state flags, etc.) on the human default lane
- **Next Upstream Writer Boundary**:
  - the next upstream component that writes/advances the bad state (animation clip writer, retarget op, pawn transform writer, mover update, etc.), named concretely enough to guide a discriminating next check

These truths must be durable so later sessions do not re-open broad search or re-litigate already narrowed seams.

## Causal Claim

If ThirdPerson enforces a first-disagreement gate inside its repo-local debugging adapter and requires the gate’s outputs to exist in `BUG-...` and `REGRESSION-RUN-...` artifacts before the next fix attempt, then:

- the system will be forced into controlled narrowing rather than tweak iteration
- the human will have fewer interventions demanding root-cause debugging
- debugging progress will survive session boundaries because the narrowed seam and writer chain are durably recorded

## Evidence

Evidence is concentrated in [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-005`) and the stable runtime-defect and root-cause-debugging events in [../../HumanInputEvents/INDEX.json](../../HumanInputEvents/INDEX.json), including:

- explicit demand for root cause and "do not stop without root cause" (`...-4768`)
- explicit instruction to reopen the task for a new defect (`...-6105`)
- explicit instruction to adopt the shared debugging method and record first concrete disagreement values and writer chain (`...-7446`)

## Why This Mechanism

This proposal chooses an artifact-level gate because:

- "remember to debug properly" is not enough; the packet shows it was already in play and still required intervention
- ThirdPerson already has durable artifact homes (`BUG` and `REGRESSION-RUN`) that are the natural enforcement surface for runtime truth

The mechanism is fail-closed: if the required narrowing output is missing, the next fix attempt is not considered justified.

## Scope Rationale

Repo-local scope is earned because:

- the evidence surfaces and artifact paths for runtime proof are repo-specific
- ThirdPerson already has a repo-local debugging adapter intended to express the local delta without rewriting the shared method

This task does not attempt to standardize the full gate across all repos (that would require additional multi-repo evidence).

## Goals

- Prevent "tweak mode" continuation after a failed fix attempt on a runtime defect.
- Make the first disagreement and upstream writer boundary durable so future work is anchored.
- Increase honesty of bug and regression artifacts so closure claims cannot outrun narrowed evidence.

## Non-Goals

- Building an automation-owned runtime checker with stable numeric thresholds (explicitly deferred).
- Adding new shared debugging schema or changing shared debugging workflow docs in this slice.
- Guaranteeing that every bug has a single root cause; some will end in a bounded contradictory branch.

## Implementation Home

- Repo-local enforcement doc:
  - [ThirdPerson/DEBUGGING.md](/c:/Agent/ThirdPerson/DEBUGGING.md)
- Task-owned defect narrative:
  - `Tracking/Task-<id>/BUG-<NNNN>.md`
- Task-owned executed proof reruns:
  - `Tracking/Task-<id>/Testing/REGRESSION-RUN-<NNNN>.md`

## Implementation Home Rationale

- The shared debugging method already exists; ThirdPerson should not duplicate it verbatim.
- ThirdPerson’s repo-local `DEBUGGING.md` is the correct adapter home for enforcing "when the gate triggers" and "which task artifacts must carry the required fields."
- `BUG` and `REGRESSION-RUN` artifacts are where runtime truth and rerun evidence already live; enforcing the gate anywhere else would be easier to bypass or forget.

## Constraints And Baseline

- Keep all narrowing anchored to the repo’s human default lane as defined by [ThirdPerson/REGRESSION.md](/c:/Agent/ThirdPerson/REGRESSION.md) unless the repo explicitly approves a different lane as equivalent.
- Preserve the shared debugging method language as the authoritative process; this task adds a repo-local enforcement gate, not a new method.
- Keep the gate small enough to be used consistently: first disagreement + next upstream writer boundary, not a full essay.

## Proposed Changes

These are the concrete, reviewable surfaces this winner changes.

1. **Add a repo-local "First-Disagreement Gate" section**
   - File: [ThirdPerson/DEBUGGING.md](/c:/Agent/ThirdPerson/DEBUGGING.md)
   - Add a new section, for example `## First-Disagreement Gate (Runtime Defects)`, that defines:
     - trigger: a runtime defect is reopened, or one fix attempt has failed and another is about to be attempted
     - gate: before another fix attempt, the task must update `BUG-...` and `REGRESSION-RUN-...` with the required fields below
     - enforcement boundary: "do not implement the next fix until the artifacts contain the first disagreement + next writer boundary"
2. **Require a stable subsection inside `BUG-<NNNN>.md` when the gate triggers**
   - Artifact: `Tracking/Task-<id>/BUG-<NNNN>.md`
   - Require:
     - `### First Concrete Disagreement`:
       - expected vs observed behavior on the human default lane
       - at least one concrete value disagreement (e.g. foot/toe world Z vs ground/capsule bottom; pelvis roll; component-space transform deltas)
       - stable evidence anchor (screenshot name, log line, CSV row, artifact id)
     - `### Next Upstream Writer Boundary`:
       - name the next upstream writer/updater to trace (asset writer, retarget op, anim evaluator, gameplay component writer)
       - state the discriminating next check that will prove or falsify that writer as the cause
3. **Require a minimal cross-link and lane disclosure inside `REGRESSION-RUN-<NNNN>.md` for reruns**
   - Artifact: `Tracking/Task-<id>/Testing/REGRESSION-RUN-<NNNN>.md`
   - Require:
     - link to the active `BUG-<NNNN>.md`
     - explicitly state which lane/case ids were rerun and what was observed at the first disagreement seam

## Acceptance Criteria

- When a runtime defect is reopened or survives one fix attempt, the next fix attempt does not proceed until:
  - `BUG-<NNNN>.md` contains the first concrete disagreement with values and the next upstream writer boundary
  - any rerun evidence is recorded in `REGRESSION-RUN-<NNNN>.md` with a link back to the bug note
- A later session can resume by reading the bug note and immediately knowing:
  - what exact disagreement is being traced
  - what upstream boundary is next
  - what lane and evidence anchor is used

## Expected Resolution

Human-facing outcome:

- Fewer "stop tweaking; find root cause" interventions.
- Debugging progress persists across sessions and subagents because the narrowing seam is durable.
- Closure and pass narratives become more honest because they are forced to match the preserved disagreement seam.

## Human Relief If Successful

- Less repeated demand for root cause.
- Less time wasted on bounded tweaks that do not collapse uncertainty.
- More trustworthy bug and regression artifacts for review and handoff.

## Internal Mechanism Map

1. Define a repo-local gate: when a second fix attempt is about to happen, require artifact updates first.
2. Make the required artifact sections explicit and stable.
3. Force the work loop to proceed only after the first disagreement seam is recorded durably.

## Rival Explanations Considered

- "The issue is aesthetic; narrow values aren’t needed."
  - Rejected: the packet explicitly treats the defects as runtime-visible failures requiring evidence and controlled narrowing.
- "We already have shared debugging docs; that’s enough."
  - Rejected: the packet contains explicit instruction to apply the shared method and still needed repeated human intervention; enforcement at the repo artifact boundary is missing.

## Rival Mechanisms Considered

- `P-005 / Option B` (automation-owned runtime checker with quantitative thresholds):
  - Deferred: it can reduce manual extraction later, but the packet evidence is insufficient to freeze trustworthy numeric thresholds today.
  - It also does not replace the need to preserve first disagreement and writer chain in a durable bug narrative when a defect is reopened.

## Tradeoffs

- Adds friction before "just try another fix."
  - Intentional: the friction buys narrowing and prevents time waste.
- Requires some value extraction work.
  - Also intentional: until trustworthy automation thresholds exist, the first disagreement still needs concrete values.

## Shared Substrate

- Shared debugging method: [../../../../../../../../Processes/DEBUGGING.md](../../../../../../../../Processes/DEBUGGING.md)
- ThirdPerson repo debugging adapter: [ThirdPerson/DEBUGGING.md](/c:/Agent/ThirdPerson/DEBUGGING.md)
- ThirdPerson regression lane truth (default lane anchoring): [ThirdPerson/REGRESSION.md](/c:/Agent/ThirdPerson/REGRESSION.md)

## Not Solved Here

- Building an unattended numeric checker that enforces thresholds automatically.
- Solving regression claim validity (`P-001`) or approval packet shape (`P-004`).
- Capturing task-specific constraints (`P-006`).

## What Does Not Count

- Category-level diagnoses like "single-node seam" without the first concrete disagreement values.
- A bug note that asserts a cause without naming the upstream writer boundary that wrote the bad state.
- Rerun claims that do not disclose lane/case ids and do not link to the active bug note.

## Remaining Uncertainty

- The smallest universal "value set" that should be required for ThirdPerson runtime defects (foot clearance, pelvis roll, mesh mount offsets, etc.) versus letting the gate accept any concrete-value seam.

## Falsifier

This proposal is falsified if, after implementation:

- the task proceeds to a second (or later) fix attempt on a runtime defect without a preserved first concrete disagreement (with values) and a named next upstream writer boundary in durable artifacts

## Proof Plan

1. Update [ThirdPerson/DEBUGGING.md](/c:/Agent/ThirdPerson/DEBUGGING.md) with the gate definition and required artifact sections.
2. Create a small "before/after" example inside a real task folder:
   - a reopened defect with a bug note missing the required sections (pre-gate)
   - the same defect after adding `First Concrete Disagreement` and `Next Upstream Writer Boundary` (post-gate)
3. Run a rehearsal debugging loop and confirm the gate forces narrowing before action.

## Open Questions

- Should the gate trigger after exactly one failed fix attempt, or only when a defect is explicitly reopened by the human?
- Should the gate require the first disagreement to be backed by a file artifact (CSV/log/screenshot path) when available, or is a durable numeric excerpt sufficient?

## References

- Burden driver: [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-005`)
- Designed options: [../SOLUTION-DESIGN.md](../SOLUTION-DESIGN.md#p-005-first-disagreement-debugging-gate)
- Frozen winner boundary: [../WINNER-SYNTHESIS.md](../WINNER-SYNTHESIS.md#w-005-repo-local-first-disagreement-debugging-gate)
- Final matrix row: [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md#p-005-first-disagreement-debugging-gate)

