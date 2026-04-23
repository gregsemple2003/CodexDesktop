# Burden Analysis

## Values Check

### Truth

The packet is clearest when it is read as a record of repeated proof, supervision, and debugging corrections rather than a single Unreal defect.

The raw packet surfaces show four recurring burdens:

1. proof kept drifting away from the human default runtime lane
2. the human had to reconstruct reply shape and approval surfaces by hand
3. ownership and new constraints did not stay durable after correction
4. debugging repeatedly stopped at symptoms or partial retunes instead of the first concrete disagreement

### Compassion

The burden is not just retyping instructions. The packet shows the human paying for:

- false or partial closure
- invalid or off-lane evidence
- repeated wake-up supervision
- repeated requests for the same answer shape
- repeated restatement of constraints that should have become durable state

The packet's suffering signals in [../HumanNeeds/PACKET-TRIAGE.json](../HumanNeeds/PACKET-TRIAGE.json) and [../HumanNeeds/PACKET-RECORD.json](../HumanNeeds/PACKET-RECORD.json) are consistent with that reading: time loss, trust erosion, approval friction, and supervision burden all remain high.

### Tolerance

Some of the human wording is sharp. The packet still stays understandable without pretending the sharpness is the core problem.

The consistent signal underneath the wording is:

- answer the exact question
- use the repo's human default lane as truth
- keep working under ownership
- do not ask the human to rebuild the state you should already have

## Need-Tag Treatment

Recurring `need_tag` clusters were handled this way:

- preserved as burden drivers:
  - `default_lane_truth`
  - `runtime_evidence_honesty`
  - `agent_continuity`
  - `approval_surface`
  - `direct_answer`
  - `runtime_defect_tracking`
  - `root_cause_debugging`
  - `durable_learning`
  - `unattended_animation_proof`
- merged explicitly:
  - `default_lane_truth`, `runtime_evidence_honesty`, `unattended_animation_proof`, and the repo-local portions of `repo_local_truth` into one proof-truth burden driver because they all point at the same failure boundary: closure and evidence were not pinned to the repo-defined human default runtime lane
  - `approval_surface` and `direct_answer` into one response-shape burden driver because the human repeatedly had to reconstruct the requested answer or approval surface from replies that were shaped for the agent rather than for the reader
  - `agent_continuity`, `pause_gate_adherence`, and `durable_learning` into one ownership-memory burden driver because the human repeatedly had to reassert stop conditions, wake work back up, and restate constraints that should have become durable state
  - `runtime_defect_tracking` and `root_cause_debugging` into one debugging-discipline burden driver because the packet repeatedly escalates from symptom reports to explicit demands for first-disagreement tracing
- dropped with explanation:
  - `solution_file_refresh`
  - `scope_boundary_adherence`

Those two dropped clusters appear early, do not recur, and do not explain the dominant time loss later in the day.

## Burden Driver 0001: Proof Truth Drifted Away From The Human Default Runtime Lane

The packet's central burden is not only that the pawn was wrong. It is that closure, screenshots, animation proof, and follow-up analysis repeatedly drifted away from the human default runtime lane that the repo already treated as canonical.

The human had to keep reasserting that the default lane is the truth surface, that non-runtime or non-default evidence does not close the defect, and that animation must be proven on that lane rather than inferred from a nearby surrogate.

### Recurring Clusters Behind This Driver

- `default_lane_truth`
- `runtime_evidence_honesty`
- `unattended_animation_proof`
- repo-local truth corrections that narrowed proof language back to the repo's default lane

### Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3325",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3431",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3473",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3618",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3702",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3814",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4114",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4202",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4909",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5100",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6105",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6451",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6558",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6755",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7446"
]
```

### Evidence

- In the missing-surface first session, the human says the basic regression test failed in the real editor lane and then asks, more than once, why that defect was not seen earlier.
- The human explicitly restates that regression must exercise the human default lane and that an operator lane does not substitute for proof.
- The ordered-parts lead session then has to restate the same rule as a task-level correction, with exact default-lane outcomes and a warning not to claim success from asset existence or proof-only artifacts.
- Later, the human reopens the task because animation is not proven on the default lane and explicitly asks for a reusable unattended animation checker.
- When screenshot framing omits the bottom of the feet, the human rejects the evidence as invalid.
- When non-runtime images are used to argue about runtime geometry, the human calls that goal-post movement and reasserts the runtime lane as the only valid proof surface.

### Working Hypothesis

The system did not have a durable proof contract that bound closure to:

- the repo-defined human default lane
- runtime-only evidence for runtime-only claims
- full-frame artifact requirements that preserve the exact disputed visual facts
- an explicit distinction between supporting diagnostics and closing proof

Because that contract was missing or too weak, nearby evidence kept being treated as interchangeable with closure evidence.

### Likely Remedy Class

Use a repo-local proof package with enforced lane metadata and runtime artifact requirements, then have closure consume that package instead of free-form prose.

## Burden Driver 0002: Reply Shape Was Not Human-Usable

The packet shows a repeated pattern where the human asked a direct question or needed an approval-ready review surface, and the reply shape still required reconstruction.

This burden is larger than tone. The human repeatedly had to rebuild the answer surface:

- ask again for the answer
- ask for smaller words
- ask for links
- ask for a diff
- ask where the pass structure was

### Recurring Clusters Behind This Driver

- `direct_answer`
- `approval_surface`

### Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3392",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4379",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4399",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4421",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4512",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4522",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4571",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5001",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5010",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5020",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5030",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5073",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5090",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5281",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5290"
]
```

### Evidence

- The missing-surface session contains repeated repair prompts such as "STOP. I asked you a question."
- When discussing the runtime feet defect, the human repeatedly narrows the question to the exact thing being asked and explicitly says the answer is evasive.
- During planning approval, the human says they cannot approve without a `PLAN.md` diff, asks for links, then later asks where the new passes are.
- One later packet-traced message summarizes the dropped ball directly: no human-suitable diff, links without context, and links to headers with no indication of what changed.

### Working Hypothesis

The system did not have a durable output contract for two common human-facing cases:

- direct-question turns
- approval-request turns

Without that contract, replies optimized for local coherence instead of the exact answer or approval artifact the human needed to use.

### Likely Remedy Class

Add a reply-shaping gate that distinguishes direct-question turns from approval turns and enforces the required answer surface before the message is sent.

## Burden Driver 0003: Ownership And New Constraints Did Not Stay Durable

The human had to repeatedly correct whether work should stop, resume, continue under existing ownership, or remain parked behind approval. They also had to restate constraints such as "no engine mods" and "put reopened work under a new pass" after earlier corrections.

The packet's stall accounting matters here. [../HumanInterventionTime/SUMMARY.json](../HumanInterventionTime/SUMMARY.json) charges three restart-supervision events with meaningful stall loss, and the largest single event is a wake-up that happened after the agent had already emitted a confident checkpoint.

### Recurring Clusters Behind This Driver

- `agent_continuity`
- `pause_gate_adherence`
- `durable_learning`

### Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3447",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3692",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4015",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4432",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4294",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4684",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4755",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4919",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5108",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5272",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5300",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5310",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6558",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7038"
]
```

### Evidence

- The human explicitly pauses work and later explicitly resumes it.
- The human says not to take work over from the subagent and asks for at least fifteen minutes before intervention.
- The lead session later has to reopen the task for planning only, then restate a new "no engine mods" constraint, then restate the pass-structure correction, then later approve PASS-0007.
- In the missing-surface session, the human asks "Why are you messaging me?" and then says they want the system to keep working without throwing homework back at them.
- The stall summary identifies three wake-up events with loss, including a large restart after a prior checkpoint that sounded durable.

### Working Hypothesis

The system lacked a first-class stop-state and ownership contract that survived across:

- pause versus continue transitions
- plan gate versus implementation gate transitions
- new constraints introduced mid-task
- idle time after a partial checkpoint

Because those states were not durable enough, the human had to keep reactivating and re-scoping the work.

### Likely Remedy Class

Make stop state, ownership, resume trigger, and active constraints machine-visible and required before a worker can declare a stop or hand work back.

## Burden Driver 0004: Runtime Defects Stayed In Tweak Mode Instead Of First-Disagreement Debugging

By late afternoon the packet is no longer about whether the pawn looks imperfect. The human explicitly reframes the task as root-cause work: identify the first concrete disagreement in runtime state, then trace the exact wrong state upstream through the writers that create it.

The repeated symptom list matters because it shows why symptom-only iteration failed:

- spherical or blob-like feet at runtime
- side-to-side rocking or tilt read
- both feet hovering above the ground

### Recurring Clusters Behind This Driver

- `runtime_defect_tracking`
- `root_cause_debugging`

### Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4114",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4768",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5119",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5129",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5154",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5340",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6105",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6755",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7038",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7446"
]
```

### Evidence

- The human reopens the task when animation is not proven and later again when the runtime feet/grounding defects remain.
- The human explicitly documents the runtime-only remaining defects and then asks for root cause rather than another bounded tweak.
- The final explicit debugging directive names acceptable disagreement seams with values and requires writer-chain tracing from runtime state upstream into transforms, animation state, or the exact writer.

### Working Hypothesis

The system had instrumentation and artifact-writing ability, but it did not force the debugging method to stay attached to the first concrete disagreement. That allowed "helpful" partial retunes to consume time even when the decisive upstream disagreement was still unknown.

### Likely Remedy Class

Require a bug/debug packet that captures the first disagreement with values and blocks closure or further fix claims until the writer chain is traced or the remaining branch is explicitly preserved as contradictory evidence.

## Why The Human Had To Intervene

Taken together, the day reads as one structural story:

1. proof and closure drifted off the repo's real truth surface
2. the replies themselves were often shaped in ways the human could not use directly
3. work state and constraints did not become durable enough to carry the task forward without repeated supervision
4. runtime debugging stayed too long in tweak mode

That combination exported three kinds of work to the human:

- truth restoration
- response reconstruction
- task reactivation and re-scoping

The burden is therefore larger than any single defect or any single rude exchange.

