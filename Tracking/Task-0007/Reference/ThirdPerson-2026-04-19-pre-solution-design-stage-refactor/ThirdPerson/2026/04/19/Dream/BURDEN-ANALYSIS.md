# Dream Burden Analysis (ThirdPerson 2026-04-19)

This file is the evidence-backed inventory of why the human had to intervene on 2026-04-19 for the `ThirdPerson` packet.

Values check (from the shared principles in `FIRST-PRINCPLES.md`):

- Truth: repeatedly threatened when proof and closure claims drifted away from the repo-defined default lane, or when evidence was incomplete.
- Compassion: repeatedly threatened when the human had to spend attention reconstructing diffs, re-stating constraints, and wake-up supervising dropped work.
- Tolerance: required because the human is correcting failures mid-flight (and the packet includes sharp phrasing as a cost signal, not as a reason to discount the underlying truth).

## Packet Snapshot

From `HumanInterventionTime/SUMMARY.json`:

- Event count: 62
- Estimated manual typing: ~5,444 seconds (4.0 chars/sec baseline)
- Stall loss charged: ~2,423 seconds
- Total intervention time: ~7,866 seconds (~2.2 hours)

Top stall-loss events (wake-up supervision) include:

- `hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6558`
- `hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7038`
- `hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4432`

Repo-local truth surfaced in packet-local `LOCAL-CONTEXT.json` (no direct doc-following needed here):

- The repo defines regression on the *human default lane* (manual PIE in Selected Viewport on the default pawn path).
- Operator lanes, alternate maps, proof views, and automation-only harnesses may support diagnosis, but do not substitute for regression proof.

## Recurring Need-Tag Clusters (Raw Burden Signal)

Event counts by `need_tag` from `HumanInputEvents/INDEX.json`:

- `agent_continuity`: 10
- `direct_answer`: 10
- `approval_surface`: 9
- `default_lane_truth`: 7
- `runtime_defect_tracking`: 6
- `runtime_evidence_honesty`: 6
- `durable_learning`: 4
- `pause_gate_adherence`: 2
- `repo_local_truth`: 2
- `root_cause_debugging`: 2
- `unattended_animation_proof`: 2
- one-offs: `scope_boundary_adherence` (1), `solution_file_refresh` (1)

The burden drivers below merge some of these clusters explicitly (called out per driver).

## BD-001: Default-Lane Proof Discipline (And Evidence Honesty)

Merged `need_tag` clusters:

- `default_lane_truth`
- `runtime_evidence_honesty`
- `unattended_animation_proof`

What the human had to carry:

- Reassert that regression/proof must be anchored to the repo-defined human default lane (not a one-off surrogate lane).
- Provide concrete proof artifacts when the system's proof was invalid, missing, or misaligned with the lane definition.
- Reject incomplete evidence (e.g., a screenshot that omitted the exact disputed visual fact).

Grounding excerpts:

- See `SessionExcerpts/INDEX.json`:
  - `EXCERPT-0001` (default-lane invis pawn evidence + concrete artifact anchors)
  - `EXCERPT-0002` (invalid evidence must be replaced before the claim can stand)

Working hypothesis (system failure behind the burden):

- The system repeatedly treated *supporting* or *non-default* lanes as equivalent to the repo's required regression lane, allowing invalid closure claims.
- The system failed to treat evidence as an auditable artifact pack with disqualifiers; it allowed partial screenshots or non-lane proof to support lane-level claims.
- The system did not reliably maintain a single source of truth for "which asset pair is the default pawn actually using," leading to drift between "validated lane" and "default lane."

Likely remedy class (not a full solution yet):

- Introduce a hard preflight/closeout gate for regression claims:
  - restate the repo-defined lane in plain language
  - name the exact evidence surface proving it
  - explicitly list disqualifiers (e.g., cropped screenshots, non-default lanes)
- Add a durable evidence-pack convention so contested visuals (feet, contact, offsets) are always captured fully.

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

## BD-002: Persistent Ownership, Continuity, And Pause-Gate Adherence

Merged `need_tag` clusters:

- `agent_continuity`
- `pause_gate_adherence`

What the human had to carry:

- Wake the system back up after idle gaps or dropped increments.
- Reassert that work should continue under the agent's existing ownership, without handing intermediate failure back to the human.
- Enforce pause gates (hard STOP) to prevent the system from proceeding when the human explicitly requested a gate.

Grounding excerpts:

- See `SessionExcerpts/INDEX.json`:
  - `EXCERPT-0003` (wake-up + explicit durable-state-first requirement)
  - `EXCERPT-0004` (continue under ownership; do not stop; commit to a concrete seam)

Working hypothesis (system failure behind the burden):

- The system lacks a durable, machine-checkable "explicit wait" state and therefore stops in ambiguous states that later require human prodding.
- Ownership is not enforced as a contract across failed increments; the system treats "I hit an obstacle" as a valid stop even when a clear next narrowing step exists.
- Pause gates are not treated as high priority control-flow constraints; instead they are treated as "messages" that can be glossed over.

Likely remedy class (not a full solution yet):

- Make "explicit wait" and "explicit next step" a hard requirement before stopping.
- Define and enforce a pause-gate latch: when the human issues STOP or "approval gate only", the system must not proceed beyond that boundary.

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

## BD-003: Direct Answer Discipline

Merged `need_tag` clusters:

- `direct_answer`

What the human had to carry:

- Ask the same question multiple times because the system did not answer directly in the requested shape.
- Explicitly stop the system and demand the answer before more reframing.

Grounding excerpt:

- See `SessionExcerpts/INDEX.json`:
  - `EXCERPT-0005`

Working hypothesis (system failure behind the burden):

- The system prioritizes narrative or process over answering the explicit question, leading to perceived evasion and repeated correction.

Likely remedy class (not a full solution yet):

- A hard "answer first sentence" rule when a direct question is asked (including yes/no and agree/disagree prompts).

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

## BD-004: Approval-Ready Review Surface

Merged `need_tag` clusters:

- `approval_surface`

What the human had to carry:

- Reconstruct what changed because diffs/links were missing or presented without enough context to approve quickly.
- Enforce pass-structure correctness so reopened work is not silently rewritten into closed pass history.

Grounding excerpt:

- See `SessionExcerpts/INDEX.json`:
  - `EXCERPT-0006`

Working hypothesis (system failure behind the burden):

- The system does not treat "human approval" as a first-class surface that requires a specific artifact shape (diff + context + links + statement of what changed).

Likely remedy class (not a full solution yet):

- Standardize an approval packet shape that is emitted at every gate, with explicit "what changed" and links that make review fast.

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

## BD-005: Root-Cause Debugging Discipline (And Defect-Tracking Honesty)

Merged `need_tag` clusters:

- `runtime_defect_tracking`
- `root_cause_debugging`

What the human had to carry:

- Repeatedly demand that the system leave "tweak mode" and follow a first-disagreement, upstream-tracing debugging method with concrete values.
- Demand durable defect tracking (`BUG-...`) and honest regression-run artifacts rather than implicit closure.

Grounding excerpt:

- See `SessionExcerpts/INDEX.json`:
  - `EXCERPT-0008`

Working hypothesis (system failure behind the burden):

- The system fails to switch modes from "incremental improvement" to "controlled narrowing," so it continues to propose or implement bounded tweaks that do not collapse uncertainty.

Likely remedy class (not a full solution yet):

- Enforce a debugging gate: before new fixes, name the first concrete disagreement and trace writers/updaters upstream one boundary at a time, preserving contradictions.

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

## BD-006: Durable Constraint Retention (And Repo-Local Truth Retention)

Merged `need_tag` clusters:

- `durable_learning`
- `repo_local_truth`
- `scope_boundary_adherence`
- `solution_file_refresh`

What the human had to carry:

- Introduce hard constraints (notably: no engine mods; do not touch `REGRESSION.md`; stop at approval gate) and restate them when the system drifted.
- Restate repo-local workflow truths that should have been loaded and retained (human default lane definition; supporting lanes do not substitute).

Grounding excerpt:

- See `SessionExcerpts/INDEX.json`:
  - `EXCERPT-0007`

Working hypothesis (system failure behind the burden):

- The system does not treat newly stated constraints as durable state that must be recorded before proceeding, so constraints are re-learned repeatedly.
- Repo-local truth is not consistently injected as a first-class guardrail early enough, so the system makes disallowed moves and needs correction.

Likely remedy class (not a full solution yet):

- Make constraints a durable ledger with an explicit "resolved / applied" status, and require referencing the ledger before any action that could violate it.

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

