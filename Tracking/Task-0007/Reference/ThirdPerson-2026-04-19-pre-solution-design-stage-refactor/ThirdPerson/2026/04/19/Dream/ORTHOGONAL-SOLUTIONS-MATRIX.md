# Orthogonal Solutions Matrix (ThirdPerson 2026-04-19 Dream)

This matrix turns the burden drivers in `BURDEN-ANALYSIS.md` into a problem set with concrete options and task-shaped proposals.

## Burden-to-Problem Mapping

All major burden drivers from `BURDEN-ANALYSIS.md` are accounted for exactly once, as separate problem rows.

- `BD-001` -> `separate_problem_row` -> `PROBLEM-0001`
- `BD-002` -> `separate_problem_row` -> `PROBLEM-0002`
- `BD-003` -> `separate_problem_row` -> `PROBLEM-0003`
- `BD-004` -> `separate_problem_row` -> `PROBLEM-0004`
- `BD-005` -> `separate_problem_row` -> `PROBLEM-0005`
- `BD-006` -> `separate_problem_row` -> `PROBLEM-0006`

## PROBLEM-0001: Default-Lane Proof Gate And Evidence Pack

### Source Burden Drivers

`BD-001`

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

### Mechanism Boundary

The regression-claim boundary:

- The system cannot claim "regression passed" or "closeout proof is sufficient" unless it restates the repo-defined lane and attaches an auditable evidence pack with explicit disqualifiers.

### Acceptance Test

- For the next regression-closeout attempt on the repo-defined human default lane, the produced `REGRESSION-RUN-<NNNN>.md` (or equivalent lane proof artifact) includes:
  - Claimed lane (in ordinary language)
  - Actual flow exercised
  - Why this counts for the claimed lane
  - Disqualifiers / limitations (including evidence-quality disqualifiers like cropped screenshots)
  - A durable evidence pack that includes the contested visuals fully in-frame (e.g., feet fully visible when feet are under dispute)
- A human reviewer can validate the claim without reconstructing missing context from raw logs.

### Falsifier

- The proof gate reports "PASS" (or "sufficient") but the human default lane still visibly fails or the evidence is still invalid (cropped, wrong lane, or missing the disputed surface).

### Options

- Option A (Winner): [`./Option-Tasks/PROBLEM-0001-OPTION-A.md`](./Option-Tasks/PROBLEM-0001-OPTION-A.md)
  - Add a shared proof-quality gate plus linting for regression-claim artifacts and evidence packs.
- Option B: [`./Option-Tasks/PROBLEM-0001-OPTION-B.md`](./Option-Tasks/PROBLEM-0001-OPTION-B.md)
  - Build an unattended default-lane runner that produces a durable evidence pack by construction.

## PROBLEM-0002: Explicit Wait State And Continuity Enforcement

### Source Burden Drivers

`BD-002`

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

### Mechanism Boundary

The stop/resume boundary:

- When the system is not actively executing, it must be in an explicit, durable "wait gate" state (with a blocking question or a blocking external dependency), or it must persist a concrete next step and resumption plan.

### Acceptance Test

- After implementing the chosen option, "wake-up" interventions are not needed to resume dropped work for the next comparable multi-hour pass.
- Any stop is accompanied by a durable stop reason and one of:
  - a single blocking question the human must answer, or
  - a named external dependency and how to detect its resolution.

### Falsifier

- Wake-up supervision remains necessary (the human still has to say "continue / do not stop / resume") even when the system has enough context to proceed.

### Options

- Option A (Winner): [`./Option-Tasks/PROBLEM-0002-OPTION-A.md`](./Option-Tasks/PROBLEM-0002-OPTION-A.md)
  - Add explicit wait-state and continuity fields to durable task state and enforce them in prompts and closeout rules.
- Option B: [`./Option-Tasks/PROBLEM-0002-OPTION-B.md`](./Option-Tasks/PROBLEM-0002-OPTION-B.md)
  - Add a watchdog that detects dropped work and triggers an automatic resume or escalates to a clear operator notification.

## PROBLEM-0003: Direct Answer Gate For Explicit Questions

### Source Burden Drivers

`BD-003`

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

### Mechanism Boundary

The response-shape boundary:

- For explicit questions (especially yes/no, agree/disagree, or "is X true?"), the first sentence must be a direct answer in the requested shape.

### Acceptance Test

- For a sample of 20 direct questions, 19+ have a direct first-sentence answer and do not require a follow-up "STOP, answer my question" correction.

### Falsifier

- Direct-answer interventions continue at roughly the same rate even after the chosen mechanism is applied.

### Options

- Option A (Winner): [`./Option-Tasks/PROBLEM-0003-OPTION-A.md`](./Option-Tasks/PROBLEM-0003-OPTION-A.md)
  - Add an answer-first gate to shared prompts plus a lightweight output lint check.
- Option B: [`./Option-Tasks/PROBLEM-0003-OPTION-B.md`](./Option-Tasks/PROBLEM-0003-OPTION-B.md)
  - Add a dashboard surface that detects unanswered questions and blocks "continue" until answered.

## PROBLEM-0004: Approval Packet And Diff Reconstruction

### Source Burden Drivers

`BD-004`

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

### Mechanism Boundary

The approval gate boundary:

- Before asking for plan/pass approval, the system must emit an approval packet with diffs and contextual links that a human can review quickly.

### Acceptance Test

- For the next approval gate:
  - a single "approval packet" artifact exists
  - it lists what changed, why, and provides usable diff context (not raw path dumps)
  - the human can approve/reject without re-opening many files to reconstruct context

### Falsifier

- The human still reports that links/diffs are unusable or missing and has to reconstruct the state by hand.

### Options

- Option A (Winner): [`./Option-Tasks/PROBLEM-0004-OPTION-A.md`](./Option-Tasks/PROBLEM-0004-OPTION-A.md)
  - Add a shared approval-packet generator for task artifacts (plan, handoff, state, bug notes) plus a standard shape.
- Option B: [`./Option-Tasks/PROBLEM-0004-OPTION-B.md`](./Option-Tasks/PROBLEM-0004-OPTION-B.md)
  - Build an interactive review UI (dashboard) that renders diffs and contextual links automatically.

## PROBLEM-0005: First-Disagreement Debugging Gate (With Durable Bug Narrative)

### Source Burden Drivers

`BD-005`

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

### Mechanism Boundary

The "debugging vs tweaking" boundary:

- After a defect survives one fix attempt (or after a regression failure is confirmed), switch to the shared debugging method:
  - define the first concrete disagreement with values
  - trace those exact values upstream through writers/updaters one boundary at a time
  - preserve contradictions in a durable bug narrative

### Acceptance Test

- For the next comparable runtime-defect iteration:
  - a `BUG-<NNNN>.md` exists (or is updated) immediately when the defect is confirmed
  - the bug note contains a first disagreement with values, a narrowing path timeline, and references to the executed regression slice artifacts
  - a fix is not claimed without the corresponding default-lane evidence

### Falsifier

- The work continues to oscillate between bounded tweaks without collapsing uncertainty, and the human still has to demand "trace the disagreement upstream" explicitly.

### Options

- Option A (Winner): [`./Option-Tasks/PROBLEM-0005-OPTION-A.md`](./Option-Tasks/PROBLEM-0005-OPTION-A.md)
  - Add a debugging gate template plus enforcement in prompts and closeout checks.
- Option B: [`./Option-Tasks/PROBLEM-0005-OPTION-B.md`](./Option-Tasks/PROBLEM-0005-OPTION-B.md)
  - Add a runtime disagreement capture harness (feet world Z vs ground, pelvis roll, etc.) to make narrowing cheaper and more repeatable.

## PROBLEM-0006: Durable Constraint Ledger And Repo-Truth Injection

### Source Burden Drivers

`BD-006`

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

### Mechanism Boundary

The constraint-application boundary:

- Before acting, the system must:
  - capture new constraints durably when they appear, and
  - reference the current constraint ledger explicitly when making a move that could violate it.

### Acceptance Test

- For the next comparable multi-step task:
  - new constraints (no engine mods, do not touch `REGRESSION.md`, stop at approval gate) are written durably within one checkpoint
  - the system stops making moves that violate previously recorded constraints
  - the human does not need to restate the same hard constraint more than once

### Falsifier

- Constraints continue to be re-stated by the human because they were not recorded durably or not consulted before acting.

### Options

- Option A (Winner): [`./Option-Tasks/PROBLEM-0006-OPTION-A.md`](./Option-Tasks/PROBLEM-0006-OPTION-A.md)
  - Add a task-local constraint ledger artifact plus prompt enforcement.
- Option B: [`./Option-Tasks/PROBLEM-0006-OPTION-B.md`](./Option-Tasks/PROBLEM-0006-OPTION-B.md)
  - Add a global constraint memory store that auto-injects constraints into prompts for continuity across sessions.

## Winners And Rollout Order

Recommended rollout order prioritizes gating the highest-cost truth failures first, then reducing review friction and debugging churn, then reducing supervision cost:

1. `PROBLEM-0006 OPTION-A` (constraints become durable immediately): [`./Option-Tasks/PROBLEM-0006-OPTION-A.md`](./Option-Tasks/PROBLEM-0006-OPTION-A.md)
2. `PROBLEM-0001 OPTION-A` (proof-quality gate): [`./Option-Tasks/PROBLEM-0001-OPTION-A.md`](./Option-Tasks/PROBLEM-0001-OPTION-A.md)
3. `PROBLEM-0005 OPTION-A` (first-disagreement debugging gate): [`./Option-Tasks/PROBLEM-0005-OPTION-A.md`](./Option-Tasks/PROBLEM-0005-OPTION-A.md)
4. `PROBLEM-0004 OPTION-A` (approval packet generator): [`./Option-Tasks/PROBLEM-0004-OPTION-A.md`](./Option-Tasks/PROBLEM-0004-OPTION-A.md)
5. `PROBLEM-0002 OPTION-A` (explicit wait-state + continuity): [`./Option-Tasks/PROBLEM-0002-OPTION-A.md`](./Option-Tasks/PROBLEM-0002-OPTION-A.md)
6. `PROBLEM-0003 OPTION-A` (answer-first gate): [`./Option-Tasks/PROBLEM-0003-OPTION-A.md`](./Option-Tasks/PROBLEM-0003-OPTION-A.md)

Longer-term follow-ups (recommended only after the gates above reduce rework):

- `PROBLEM-0001 OPTION-B` (unattended default-lane runner)
- `PROBLEM-0002 OPTION-B` (watchdog/resume service)
- `PROBLEM-0004 OPTION-B` (review UI)
- `PROBLEM-0005 OPTION-B` (runtime disagreement capture harness)
- `PROBLEM-0006 OPTION-B` (global constraint memory store)

