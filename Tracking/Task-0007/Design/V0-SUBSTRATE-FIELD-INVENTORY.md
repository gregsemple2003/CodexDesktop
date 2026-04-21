# Task 0007 V0 Substrate Field Inventory

## Purpose

Define the smallest defensible field inventory for storing human-input events and linking later analysis artifacts without freezing the analysis structure too early.

This doc is intentionally narrower than a full Jarvis schema. It defines the substrate that should still make sense even if the analyzer is not Codex, even if the upper-level taxonomy changes, and even if the daily report format evolves.

## Boundary

This doc defines only:

- identity
- provenance
- raw source capture or source references
- repo, thread, scope, and task linkage
- lifecycle status
- artifact references for later analysis

This doc does not define:

- inference taxonomy
- desire taxonomy
- deliberation structure
- confidence or warrant enums
- routing policy
- daily report semantics
- analyzer-specific reasoning fields

Those belong in attached artifacts that can evolve without breaking the base record contract.

## Design Rules

- Prefer stable facts over interpretive fields.
- Allow unknowns instead of fabricated values.
- Preserve raw source material or a durable source reference.
- Keep linkage explicit so later artifacts can be traced back to the triggering input.
- Let interpretation live in versioned artifacts, not in the substrate.
- Promote new base fields only after repeated concrete use shows they are durable.

## Common Conventions

- IDs are opaque strings and only need to be stable within the local system.
- Timestamps use ISO 8601 datetime strings when sub-day precision exists in the source system.
- Date-only timestamp values are a backfill fallback only and should not be the default for raw `jsonl` intake.
- If a value is unknown, omit it or set it to `null`; do not invent it.
- Paths, URLs, and file-line references should be stored as durable references, not prose descriptions.

## ArtifactRef

`ArtifactRef` is the reusable pointer type for source captures, analysis notes, reports, tasks, and later outputs.

### Fields

| Field | Required | Meaning |
| --- | --- | --- |
| `uri` | yes | Durable pointer to the artifact. This can be a local file path, file-line reference, URL, issue link, or other stable locator. |
| `kind` | yes | Broad storage kind such as `markdown`, `json`, `html`, `pdf`, `issue`, `pull_request`, or `url`. |
| `relation` | yes | Why this artifact is linked, such as `source_capture`, `analysis_note`, `report`, `task`, `proposal`, or `audit`. |
| `label` | no | Human-readable short name. |
| `content_hash` | no | Optional integrity hint for immutable captures. |
| `captured_at` | no | When this artifact was captured or created, if known. |

### Notes

- `relation` is linkage metadata, not analysis structure.
- `ArtifactRef` should be sufficient to attach future JSON, Markdown, or other artifacts without changing the event or record shape.

## HumanInputEvent

`HumanInputEvent` is the substrate record for one human-originating input or human-signaled action.

Examples:

- a chat message
- a PR comment
- a review comment
- an issue comment
- a manual `@codex` invocation
- a human approval, dismissal, or rerun request when those are meaningful signals

### Fields

| Field | Required | Meaning |
| --- | --- | --- |
| `event_id` | yes | Stable local identifier for the event. |
| `repo_ref` | yes | Repository identifier or durable local repo path. |
| `source` | yes | Input surface such as `chat`, `github_pr_comment`, `github_review_comment`, `issue_comment`, or `manual_delegation`. |
| `author_ref` | yes | Stable author identity as available locally. |
| `captured_at` | yes | Full capture timestamp for this event. |
| `occurred_at` | no | Full source timestamp for when the event actually occurred, if known. |
| `thread_ref` | no | Durable thread or conversation locator. |
| `parent_event_id` | no | Upstream event if this is a reply or child event. |
| `raw_text` | no | Verbatim human text when the source is text and local retention allows it. |
| `source_artifact_refs` | yes | One or more `ArtifactRef` values pointing at the raw capture, transcript, webhook payload, or source page. |
| `scope_refs` | no | Repo subtree, file path, or other narrow scope references attached to the event. |
| `task_refs` | no | Task-owned artifacts or task IDs that currently own follow-up work. |
| `status` | yes | Lifecycle state such as `captured`, `redacted`, `superseded`, or `invalid`. |

### Notes

- `raw_text` is optional because some systems may only retain a source ref plus a redacted payload.
- `captured_at` is required even when `occurred_at` is unknown.
- For raw `jsonl` intake, `captured_at` should normally be a full datetime, not a date-only value.
- This type intentionally does not include inferred need, confidence, or preventability.

## InterventionRecord

`InterventionRecord` is the durable container that ties one or more human input events to the later artifact trail around that intervention.

It is intentionally a linkage and lifecycle record, not an analysis schema.

### Fields

| Field | Required | Meaning |
| --- | --- | --- |
| `record_id` | yes | Stable local identifier for the intervention record. |
| `repo_ref` | yes | Repository identifier or durable local repo path. |
| `input_event_ids` | yes | One or more `HumanInputEvent` identifiers aggregated by this intervention record. |
| `anchor_event_id` | no | Optional canonical entrypoint when one input event is the trigger or best entrypoint for the intervention thread. |
| `thread_ref` | no | Durable thread or conversation locator when the intervention spans a known thread. |
| `scope_refs` | no | Repo subtree, file path, or other narrow scope references. |
| `task_refs` | no | Task-owned artifacts or task IDs that own the intervention thread. |
| `artifact_refs` | yes | Linked `ArtifactRef` values for source captures, analysis notes, reports, proposals, audits, and other outputs. |
| `created_at` | yes | Full local creation timestamp for this record. |
| `created_by` | yes | Human or system actor that created the record. |
| `updated_at` | no | Last local update time, if tracked. |
| `status` | yes | Lifecycle state such as `open`, `closed`, `merged`, `superseded`, or `redacted`. |
| `supersedes_record_id` | no | Prior record displaced by this one, if any. |

### Notes

- `input_event_ids` is the direct reference set for the human inputs that this record aggregates.
- `anchor_event_id` is only an entrypoint convenience; it is not a substitute for the full aggregated input set.
- `artifact_refs` is where later analysis attaches. That is the main escape hatch that lets the upper-level schema evolve.
- `InterventionRecord` does not require a one-shot decision about inference taxonomy or deliberation format.
- Default usage can still be one record per human input if `input_event_ids` contains only one value, but the substrate also supports later clustering without changing the type.

## Deferred On Purpose

These fields may become useful later, but they are intentionally excluded from the base substrate for now:

- inferred need
- candidate interpretations
- epistemic warrant
- confidence
- preventability or difficulty
- local constraints applied
- recommended action
- selected route
- outcome summary
- daily report ranking

Those belong in attached analysis artifacts until they prove stable enough to promote.

## Validation Against Current Corpus

### Example A: HumanInputEvent from System Integration Spec 52

This event uses the opening user request in [2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md#L1).

This is a backfilled example from a markdown transcript capture, not a raw `jsonl` event log, so the example keeps date-only precision to avoid inventing an exact time. A real `jsonl`-backed record should carry a full ISO datetime here.

```json
{
  "event_id": "evt-task0007-spec52-user-0001",
  "repo_ref": "c:/Agent/CodexDashboard",
  "source": "chat",
  "author_ref": "user",
  "captured_at": "2026-04-19",
  "thread_ref": "/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md#L1",
  "raw_text": "Come up with a spec for how this system will integrate with Codex.",
  "source_artifact_refs": [
    {
      "uri": "/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md#L1",
      "kind": "markdown",
      "relation": "source_capture",
      "label": "System Integration Spec 52"
    }
  ],
  "scope_refs": [
    "Tracking/Task-0007"
  ],
  "task_refs": [
    "/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md"
  ],
  "status": "captured"
}
```

Why this is useful:

- it preserves the raw human input
- it links to the durable transcript
- it does not fabricate an exact `occurred_at`
- it does not force any analysis fields into the event

### Example B: HumanInputEvent from System Integration Spec 54

This event uses the opening user request in [2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md#L1).

This is also a backfilled example from a markdown transcript capture, so the date-only value is a precision fallback rather than the recommended format for live ingestion.

```json
{
  "event_id": "evt-task0007-spec54-user-0001",
  "repo_ref": "c:/Agent/CodexDashboard",
  "source": "chat",
  "author_ref": "user",
  "captured_at": "2026-04-19",
  "thread_ref": "/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md#L1",
  "raw_text": "Come up with a spec for how this system will integrate with Codex.",
  "source_artifact_refs": [
    {
      "uri": "/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md#L1",
      "kind": "markdown",
      "relation": "source_capture",
      "label": "System Integration Spec 54"
    }
  ],
  "scope_refs": [
    "Tracking/Task-0007"
  ],
  "task_refs": [
    "/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md"
  ],
  "status": "captured"
}
```

Why this is useful:

- it shows the same substrate works across another capture without changing the base type
- it stores the event faithfully without deciding how it should later be analyzed

### Example C: InterventionRecord Linking Later Analysis

This example shows how the substrate can link the two conversation captures to later artifacts without embedding the analysis structure into the record itself.

As above, the example uses a date-only `created_at` because it is a backfilled design example. Live records should use full creation timestamps.

```json
{
  "record_id": "ir-task0007-jarvis-integration-0001",
  "repo_ref": "c:/Agent/CodexDashboard",
  "input_event_ids": [
    "evt-task0007-spec52-user-0001",
    "evt-task0007-spec54-user-0001"
  ],
  "anchor_event_id": "evt-task0007-spec52-user-0001",
  "thread_ref": "jarvis-integration-spec",
  "scope_refs": [
    "Tracking/Task-0007"
  ],
  "task_refs": [
    "/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md"
  ],
  "artifact_refs": [
    {
      "uri": "/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md#L1",
      "kind": "markdown",
      "relation": "source_capture",
      "label": "System Integration Spec 52"
    },
    {
      "uri": "/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md#L1",
      "kind": "markdown",
      "relation": "source_capture",
      "label": "System Integration Spec 54"
    },
    {
      "uri": "/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-01-INPUT-EVENT-INGESTION.md",
      "kind": "markdown",
      "relation": "analysis_note",
      "label": "Topic 01 Input and Event Ingestion"
    },
    {
      "uri": "/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-02-INTERVENTION-RECORD-SHAPE.md",
      "kind": "markdown",
      "relation": "analysis_note",
      "label": "Topic 02 Intervention Record Shape"
    },
    {
      "uri": "/c:/Agent/CodexDashboard/Tracking/Task-0007/Design/V0-SUBSTRATE-FIELD-INVENTORY.md",
      "kind": "markdown",
      "relation": "design_note",
      "label": "V0 Substrate Field Inventory"
    }
  ],
  "created_at": "2026-04-19",
  "created_by": "codex",
  "status": "open"
}
```

Why this is useful:

- it links triggering inputs to later artifacts
- it allows clustering or later merge without changing the event type
- it avoids freezing fields like `confidence`, `epistemic_warrant`, or `selected_route` into the base record

## Immediate Next Use

If this substrate is accepted, the next honest move is not a grand unified schema. It is to instantiate a few real `HumanInputEvent` and `InterventionRecord` examples against the current corpus and see which additional fields actually recur.

Only after that should any analysis field be promoted into the base contract.

## References

- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md)
- [TOPIC-01 input and event ingestion](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-01-INPUT-EVENT-INGESTION.md)
- [TOPIC-02 intervention record shape](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-02-INTERVENTION-RECORD-SHAPE.md)
- [TOPIC-04 deliberation record shape](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-04-DELIBERATION-RECORD-SHAPE.md)
- [2026-04-19 System Integration Spec 52](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md)
- [2026-04-19 System Integration Spec 54](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md)
