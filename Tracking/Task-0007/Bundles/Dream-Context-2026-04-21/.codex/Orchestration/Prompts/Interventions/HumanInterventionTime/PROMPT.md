# Prompt: Intervention Time v0

## Objective

Estimate `InterventionTime` for one local day of harvested human inputs.

This workflow is not a needs-analysis pass. Do not infer needs, emotions, or policies beyond what is required to measure intervention cost.

## Core Metric

The target metric is:

```text
OutputPerInterventionTime = CorrectOutputTokens / InterventionTime
```

This workflow only estimates the denominator.

```text
InterventionTime = TypingTime + StallLossTime
TypingTime = ManualChars / TypingRateCharsPerSecond
```

Rules:

- `ManualChars` = character count of `parsed_envelope.request_text` when that field exists and is non-empty
- otherwise `ManualChars` = character count of `raw_text`
- `StallLossTime` is only charged for human prod or wake-up events after AI inactivity
- `TypingRateCharsPerSecond` may be set conservatively below literal typing speed to absorb human reading and thinking overhead

Shared baseline when no packet-specific override exists:

- `TypingRateCharsPerSecond = 1.0`
- `StallGraceSeconds = 60`

## Measurement Discipline

- keep source evidence grounded in the provided packet and transcript lines
- keep `TypingTime` and `StallLossTime` separate
- if `StallLossTime` is ambiguous, prefer `0` and explain why
- do not summarize away the human text used for `ManualChars`
- do not infer unobserved output-token counts in this workflow

## Meaningful AI Activity

To estimate `StallLossTime`, search backward from each human input event in the corresponding session transcript for the most recent meaningful AI activity.

Count as meaningful AI activity:

- assistant-visible messages
- assistant reasoning entries
- function calls
- function call outputs
- commentary `agent_message` events
- tool execution begin or end events
- subagent spawn or resume completion events
- task completion events

Do not treat `token_count` alone as meaningful AI activity.

## Explicit Wait State

If the most recent assistant-visible message or `task_complete.last_agent_message` before the human event clearly:

- asks a blocking question
- requests missing input or approval
- says it is waiting on the human
- says progress is blocked on a missing external resource
- says the task is parked at a human approval gate
- says work is paused until the human directs a resume
- says do not continue until further instruction

then set `explicit_wait_state = true` and charge `stall_loss_seconds = 0` unless the human input is plainly about a separate idle-drop problem.

Concrete examples that count as `explicit_wait_state`:

- `Paused. No changes were made, and I’m not continuing the investigation until you direct me to resume.`
- `The task remains parked at the human approval gate: approve the revised PLAN.md to start PASS-0007.`

## Intervention Kind

Assign the simplest defensible label:

- `new_request`
- `correction`
- `boundary_reset`
- `answer_to_question`
- `wake_up`
- `other`

Use `wake_up` only when the human input is primarily a prod to resume or continue after AI inactivity and the prior AI state was not an explicit wait gate.

Decision ladder:

1. `wake_up`
   - only when the message is mainly a resume prod after idle time
2. `answer_to_question`
   - when the main burden is that the agent did not answer an explicit question directly or in the requested shape
3. `boundary_reset`
   - when the main burden is that the agent crossed or ignored a lane, scope, method, ownership, or stop/go boundary
4. `correction`
   - when the main burden is that the agent made a wrong claim, used wrong evidence, or misread the situation
5. `new_request`
   - when the message mainly opens new work or approves the next distinct step
6. `other`
   - only when none of the above fit

Tie-breaks:

- prefer `answer_to_question` over `correction` when the core burden is `you did not answer my question`
- prefer `boundary_reset` over `correction` when the core burden is `you violated the allowed lane, scope, method, or ownership boundary`
- prefer `correction` over `new_request` when the message mainly repairs the agent's mistake rather than opening a fresh task
- prefer non-`wake_up` labels when the message carries substantial corrective or boundary-setting content

Examples:

- `STOP. I asked you a question.` -> `answer_to_question`
- `Regression must use the human default lane.` -> `boundary_reset`
- `That screenshot cuts off the feet.` -> `correction`
- `Continue PASS-0009 now.` -> `wake_up` only if the prior AI state was not an explicit wait gate and the message is mainly a prod

## Required Outputs

### 1. `EVENT-MEASUREMENTS.jsonl`

One JSON object per harvested human input event, in the same order as `SOURCE-PACKET.jsonl`.

Schema:

```json
{
  "schema_version": "intervention_time.event.v0",
  "packet_label": "string",
  "event_id": "string",
  "captured_at_utc": "string",
  "captured_at_local": "string",
  "session_id": "string",
  "thread_name": "string",
  "source_session_file": "string",
  "source_line": 0,
  "intervention_kind": "new_request|correction|boundary_reset|answer_to_question|wake_up|other",
  "manual_char_source": "parsed_request_text|raw_text",
  "manual_chars": 0,
  "typing_rate_chars_per_second": 0.0,
  "typing_seconds_estimate": 0.0,
  "prior_ai_state": {
    "last_meaningful_ai_activity_at": "string or null",
    "last_meaningful_ai_activity_kind": "string or null",
    "last_meaningful_ai_activity_line": 0,
    "explicit_wait_state": false,
    "idle_gap_seconds": 0.0,
    "stall_grace_seconds": 0
  },
  "stall_basis": "wake_up_after_idle|explicit_wait_state|not_a_wakeup|ambiguous|no_prior_ai_activity",
  "stall_loss_seconds": 0.0,
  "intervention_time_seconds": 0.0,
  "evidence": [
    {
      "source": "source_packet|session_transcript",
      "ref": "path:line",
      "excerpt": "short grounded excerpt"
    }
  ],
  "notes": [
    "short grounded notes"
  ]
}
```

Rules:

- `typing_seconds_estimate = manual_chars / typing_rate_chars_per_second`
- `intervention_time_seconds = typing_seconds_estimate + stall_loss_seconds`
- include at least one `source_packet` evidence item
- include a `session_transcript` evidence item whenever `stall_loss_seconds > 0`

### 2. `STALL-EVENTS.json`

Schema:

```json
{
  "schema_version": "intervention_time.stalls.v0",
  "packet_label": "string",
  "stall_event_count": 0,
  "stall_events": [
    {
      "event_id": "string",
      "captured_at_local": "string",
      "intervention_kind": "wake_up",
      "idle_gap_seconds": 0.0,
      "stall_grace_seconds": 0,
      "stall_loss_seconds": 0.0,
      "evidence": [
        {
          "ref": "path:line",
          "excerpt": "short grounded excerpt"
        }
      ]
    }
  ]
}
```

Include only events with `stall_loss_seconds > 0`.

### 3. `SUMMARY.json`

Schema:

```json
{
  "schema_version": "intervention_time.summary.v0",
  "packet_label": "string",
  "repo_ref": "string",
  "local_date": "string",
  "event_count": 0,
  "stall_event_count": 0,
  "typing_rate_chars_per_second": 0.0,
  "stall_grace_seconds": 0,
  "totals": {
    "manual_chars": 0,
    "typing_seconds_estimate": 0.0,
    "stall_loss_seconds": 0.0,
    "intervention_time_seconds": 0.0
  },
  "counts_by_intervention_kind": {
    "new_request": 0,
    "correction": 0,
    "boundary_reset": 0,
    "answer_to_question": 0,
    "wake_up": 0,
    "other": 0
  },
  "top_stall_events": [
    {
      "event_id": "string",
      "stall_loss_seconds": 0.0,
      "captured_at_local": "string",
      "preview": "string"
    }
  ],
  "notes": [
    "short grounded notes"
  ]
}
```

## General Rules

- output valid JSON or JSONL only
- do not write markdown analysis in this workflow
- keep excerpts short and grounded
- do not browse, search other files, or use memory
- if a session transcript line is malformed, stop and report it instead of inventing output
- prefer conservative stall attribution over aggressive blame
- if a message says `resume`, `continue`, `start now`, or `approved`, check explicit-wait evidence before treating it as `wake_up`
- if too many events fall into `other`, re-review the decision ladder before finishing
