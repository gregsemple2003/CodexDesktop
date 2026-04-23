# PROBLEM-0004 OPTION-A: Approval-Gate Packet Contract (Reviewable Without Reconstruction)

## Title

Require a durable approval packet shape at every human approval gate so the human can approve or reject without reconstructing diffs or pass-history effects.

## Summary

The packet shows a repeated approval-surface failure:

- approval was requested without a usable diff summary
- links were missing, or present without context (for example: links to headers with no "what changed" summary)
- pass structure drift occurred (new work folded into older pass history), forcing the human to explicitly demand new passes and corrections

This winner proposes a shared approval-gate contract: every approval ask must point to a task-owned "approval packet" that states what changed, which artifacts changed, how pass history is affected, and provides contextual review links.

The mechanism is at the approval request itself. The human should not have to open files just to infer what they are approving.

## Writeup Type

Concrete implementation task (burden-reduction proposal).

## Source Event IDs

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

## Burden Being Reduced

The human currently has to do repeated review reconstruction work:

- asking for diffs because "approve" was requested without any statement of what changed
- asking for links because the approval ask does not provide a reviewable surface
- policing pass-history integrity ("put all new work under a new pass") because reopened work is being folded into closed history

This is exported labor: the agent did the edits but left the human to do the mapping from "what did you change?" to "what do I approve?"

## Current Truth

Today, the shared lifecycle docs describe planning and pass structure, but they do not define a required approval packet shape that makes approval reviewable at the gate itself.

As a result, approval asks can devolve into:

- raw path dumps
- header-only links with no summary
- missing disclosure about whether closed pass history was modified

The packet also shows a real-world hazard: the human’s editor surface can hide changes due to local unsaved edits (`...-4581`). That means "just look at the file" is not a reliable approval contract; the approval ask itself must carry a self-contained summary of what changed.

## Target Truth

Every time a leader requests human approval for:

- `PLAN.md`
- a pass structure correction / reopen
- a closeout gate

the request must point to a task-owned approval packet that includes:

- what changed (plain language, short)
- which artifacts changed (with direct links)
- the pass-history effect (did this add a new pass? did this modify closed pass history?)
- the review links needed to approve without reconstruction

The human should be able to approve or reject by reading the approval packet first, without having to ask for basic context.

## Causal Claim

If the shared lifecycle contract requires an explicit approval packet at the gate, and the packet is stored durably in task-owned artifacts (not only in chat), then:

- approval becomes faster and less adversarial because the human is not forced into reconstruction labor
- pass-history integrity improves because the approval packet requires explicit disclosure of pass changes
- the review surface becomes resilient to local editor quirks (stale buffers, local edits) because the summary is not "hidden in the diff"

## Evidence

Evidence is concentrated in [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-004`) and the approval-surface events in [../../HumanInputEvents/INDEX.json](../../HumanInputEvents/INDEX.json), including:

- explicit complaint about lacking a plan diff / approval context (`...-4379`)
- explicit demand for links (`...-4399`)
- explicit demand that new work be placed under a new pass (`...-4421`)
- explicit postmortem: "links without context" and "no indication of what changed" (`...-4512`)

## Why This Mechanism

This mechanism is chosen because the burden is at the approval boundary itself:

- Even perfect artifacts are not helpful if the approval ask does not make the change set legible.
- Approval is a human-facing surface. It must be designed as such, not treated as an afterthought.

The approval packet contract is a durable interface: it defines what the human receives at a gate, not just what files exist somewhere on disk.

## Scope Rationale

Shared scope is earned because:

- approval gates exist across repos and workflows
- the burden is not ThirdPerson-specific; it is about the shared approval surface shape

Task-owned storage is still required because:

- the outgoing packet must point into the actual task’s changed artifacts
- the pass-history effect is task-specific

## Goals

- Make human approval reviewable without reconstruction.
- Make pass-history effects explicit at the approval gate.
- Reduce repeated "give me links" and "what changed?" interventions.

## Non-Goals

- Solving editor UX problems (stale buffers) directly.
- Replacing human approval with automation.
- Writing new repo-local approval templates only; the burden is shared and requires a shared contract.

## Implementation Home

- Shared lifecycle contract:
  - [../../../../../../../../ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
- Task-owned artifacts that carry the approval packet:
  - `Tracking/Task-<id>/PLAN.md`
  - `Tracking/Task-<id>/HANDOFF.md`

## Implementation Home Rationale

- The rule "what an approval gate must contain" is shared lifecycle truth and belongs in shared `ORCHESTRATION.md`.
- The approval packet itself must point to the actual task’s artifacts and therefore must live in task-owned docs.
- `PLAN.md` and `HANDOFF.md` are already the durable, human-readable task surfaces used for planning and resume; they are the right homes for an embedded approval packet.

## Constraints And Baseline

- Do not rely on "the human can diff locally" as the approval contract.
- Preserve the pass-history rule: executed history belongs in pass audits; reopened changes must not silently rewrite closed pass narratives.
- Keep the approval packet small, structured, and review-oriented (not a narrative dump).

## Proposed Changes

These are the concrete, reviewable surfaces this winner changes.

1. **Define the approval packet contract in shared lifecycle docs**
   - File: [../../../../../../../../ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
   - Add a new section, for example `## Approval-Gate Packet Contract`, that requires every approval ask to include:
     - `Approval ask type`: plan approval, reopen/pass-structure approval, closeout approval
     - `What changed`: 3-10 short bullets describing the semantic changes, not just filenames
     - `Artifacts changed`: direct links to the changed task artifacts (and any repo/shared docs when applicable)
     - `Pass-history effect`: explicit statement of:
       - whether any closed pass history changed (should be "no" by default)
       - whether new work was moved under a new pass (and which pass id)
     - `Review instructions`: what the human should look for, and what approval authorizes next
2. **Embed an "Approval Packet" section in task-owned `PLAN.md`**
   - Artifact: `Tracking/Task-<id>/PLAN.md`
   - Add a stable section (for example `## Approval Packet`) that can be updated when requesting plan approval, containing the required packet fields.
3. **Embed an "Approval Packet" section in task-owned `HANDOFF.md` for reopen/pass changes**
   - Artifact: `Tracking/Task-<id>/HANDOFF.md`
   - Add a stable section (for example `## Approval Packet`) used when requesting approval for reopen/pass-structure corrections or closeout gate approval.
4. **Refresh shared exemplars so implementers have a canonical template**
   - Files:
     - [../../../../../../../../Exemplars/PLAN.md](../../../../../../../../Exemplars/PLAN.md)
     - [../../../../../../../../Exemplars/HANDOFF.md](../../../../../../../../Exemplars/HANDOFF.md)
   - Add a short exemplar `## Approval Packet` section to each, showing the intended minimal shape.

## Acceptance Criteria

- When requesting approval, the approval ask points to a task-owned approval packet (in `PLAN.md` and/or `HANDOFF.md`) that includes:
  - what changed
  - artifacts changed (with links)
  - pass-history effect disclosure
  - review instructions
- The human can approve or reject without first asking for "diff", "links", or "where is pass N".
- Reopened work no longer silently rewrites closed pass history; any pass-structure change is explicitly disclosed in the approval packet and visible in the task artifact.

## Expected Resolution

Human-facing outcome:

- Approvals become fast: the human reads one structured packet and can decide.
- Review becomes less adversarial because the approval ask is no longer forcing reconstruction labor.
- Pass history becomes more trustworthy because the approval packet makes structural changes explicit.

## Human Relief If Successful

- Less repeated demand for links and diffs.
- Less need to police pass structure in real time.
- Less risk of approving something the human did not intend because the approval ask was underspecified.

## Internal Mechanism Map

1. Shared lifecycle doc defines "approval packet" as a required gate output.
2. Task artifacts embed a durable packet section for plan and reopen/closeout approvals.
3. When requesting approval, the agent links to that packet and (optionally) repeats the short summary inline.
4. The packet explicitly discloses pass-history effects, preventing silent history rewrites.

## Rival Explanations Considered

- "The human just didn’t notice the changes."
  - Rejected: the packet includes explicit complaints about missing diffs/links/context, and also shows an editor-surface hazard that makes "just look at the file" unreliable.
- "This is just about tone."
  - Rejected: the burden is structural; it is about missing review affordances and pass-history disclosure.

## Rival Mechanisms Considered

- `P-004 / Option B` (task-local approval template only):
  - Rejected as winner: without a shared lifecycle contract, the burden recurs across tasks and repos; the approval gate is lifecycle-wide.
- "Rely on IDE diff and editor UI":
  - Rejected: editor surfaces can be stale or misleading; the approval contract must be durable and self-contained.

## Tradeoffs

- More work at approval time for the agent:
  - Intentional: it replaces exported human reconstruction labor.
- Risk of verbosity:
  - Managed by requiring a minimal structured packet with bounded bullets, not narrative dumping.

## Shared Substrate

- Shared lifecycle and gate model: [../../../../../../../../ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
- Shared task artifact roles (plan vs handoff vs pass audits): [../../../../../../../../../AGENTS.md](../../../../../../../../../AGENTS.md)
- Exemplar artifacts (templates to copy from): [../../../../../../../../Exemplars/PLAN.md](../../../../../../../../Exemplars/PLAN.md), [../../../../../../../../Exemplars/HANDOFF.md](../../../../../../../../Exemplars/HANDOFF.md)

## Not Solved Here

- STOP/ownership continuity fields (`P-002`).
- Regression proof/lane gates (`P-001`).
- Direct-answer-first conversational contract (`P-003`).

## What Does Not Count

- A list of filenames with no statement of what changed.
- Header-only links with no summary of the semantic changes.
- Requesting approval while silently modifying closed pass history.
- Treating "the human can diff" as sufficient when the approval ask itself is not reviewable.

## Remaining Uncertainty

- The tightest minimal set of required fields that stays usable across many repos without growing into a full change log.
- Whether the approval packet should require a short "before/after" excerpt for certain fields (for example, plan pass list changes) or whether a bullet summary is sufficient.

## Falsifier

This proposal is falsified if, after implementation:

- the human still has to ask for links/diff/context in order to approve
- or closed pass history is still being silently rewritten during reopen/approval flows

## Proof Plan

1. Update `ORCHESTRATION.md` with the approval packet contract.
2. Update the `PLAN.md` and `HANDOFF.md` exemplars to include a minimal `## Approval Packet` section.
3. On a real task, perform a rehearsal approval request:
   - change a plan in a reviewable way (for example: add a new pass)
   - write the approval packet in `PLAN.md`
   - ask for approval by linking directly to the packet
4. Confirm the human can approve/reject without additional reconstruction questions.

## Open Questions

- Should the approval packet be required for *all* approval asks, or only for plan/reopen/closeout approvals?
- Should the approval packet explicitly state "No closed pass history was modified" as a default required sentence to prevent ambiguity?

## References

- Burden driver: [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-004`)
- Designed options: [../SOLUTION-DESIGN.md](../SOLUTION-DESIGN.md#p-004-approval-gate-review-packet)
- Frozen winner boundary: [../WINNER-SYNTHESIS.md](../WINNER-SYNTHESIS.md#w-004-approval-gate-packet-contract)
- Final matrix row: [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md#p-004-approval-gate-review-packet)

