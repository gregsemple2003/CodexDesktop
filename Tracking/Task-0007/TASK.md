# Task 0007

## Title

Bootstrap a task-owned home for Jarvis intervention analysis and conversation capture.

## Summary

This task exists to hold analysis and conversation artifacts for a proposed `Jarvis` layer that treats direct human input as evidence that the current system failed to infer something it should have inferred.

The recurring question behind that model is:

- `How could the system have inferred the need for this input?`

The immediate need is not implementation. The immediate need is to preserve the framing, tensions, open questions, and repo-fit implications in durable task artifacts instead of letting them disappear into chat history.

## Goals

- Create a durable task-owned home for ongoing Jarvis analysis and conversation notes.
- Preserve the current framing that human intervention is a system-insufficiency signal worth analyzing.
- Keep the recurring inference question explicit in the task definition.
- Define the research space for daily per-repo reports and task proposals.
- Leave explicit room for a future `HUMAN-DESIRE.md` contract.
- Keep local human constraints in scope instead of flattening all repos into one global policy.
- Keep the desired value targets explicit:
  - truthfulness
  - compassion
  - tolerance
- Require later work to separate explicit statements, strong implication, and speculation.
- Keep the promotion boundary explicit so shared orchestration rules are only promoted after the local analysis is stable.

## Non-Goals

- Implementing `Jarvis`, autonomous critics, or daily job execution in this task.
- Claiming the system can infer every new or novel human request.
- Treating every human message as equally inferable or equally blameworthy.
- Finalizing a cross-repo workflow contract before the analysis is grounded.
- Recreating [Task-0006](/c:/Agent/CodexDashboard/Tracking/Task-0006/TASK.md) as a separate active incident-capture lane. `Task-0006` is now superseded historical context.

## Constraints And Baseline

- The current request is archival and analytical, not execution-focused.
- The model is cross-repo in spirit, but this task is incubating the analysis locally first.
- [Task-0006](/c:/Agent/CodexDashboard/Tracking/Task-0006/TASK.md) holds earlier adjacent incident-capture artifacts, but that separate task is now closed as superseded by this broader research-and-reference line.
- Later inference work must represent local human constraints honestly rather than hallucinating a universal human model.
- Any later report, schema, or agent output should distinguish:
  - explicitly stated human desire
  - strongly implied desire from repeated interventions
  - speculative inference
- Cross-repo durable rules belong in `C:\Users\gregs\.codex\Orchestration\` only after the task-local shape is stable enough to promote.

## Expected Resolution

- `Task-0007` becomes the default home for this analysis thread.
- The task leaves behind a clear framing for a daily repo report that can answer:
  - what the repo's core need is
  - what high-level human desire is in force
  - what interventions happened
  - how the system could have inferred the need for each intervention
  - what task proposals would reduce recurrence
- The task leaves explicit room for a future `HUMAN-DESIRE.md` file without pretending that contract is already finished.
- The task preserves an inference-honesty rule so later work cannot blur explicit, implied, and speculative claims together.

## What Does Not Count

- A broad slogan about reducing human intervention without stored evidence or bounded questions.
- Immediate implementation of `Jarvis` without first grounding the report shape, failure taxonomy, and uncertainty boundaries.
- Treating every human message as equally inferable.
- Using `truthfulness`, `compassion`, or `tolerance` as a license to invent preferences the human did not actually express.
- Promoting cross-repo workflow rules into `.codex` before the local task has a stable analytical baseline.

## Implementation Home

Keep task-owned artifacts under `Tracking/Task-0007/`.

Use `Tracking/Task-0007/Research/Conversations/` for dated conversation captures and `Tracking/Task-0007/Research/Analysis/` for synthesized notes.

If this task later produces durable cross-repo workflow rules, report formats, or a stable `HUMAN-DESIRE.md` contract, promote those shared artifacts into `C:\Users\gregs\.codex\Orchestration\`.

## Acceptance Criteria

- `Tracking/Task-0007/` exists with `TASK.md`, `PLAN.md`, `HANDOFF.md`, and a valid `TASK-STATE.json`.
- `Tracking/Task-0007/Research/` contains explicit homes for conversation capture and analysis synthesis.
- The task definition preserves the core Jarvis framing and the recurring inference question.
- The task definition keeps local human constraints and the target humane values explicit.
- The task definition explicitly separates explicit, strongly implied, and speculative inference.
- The task definition keeps implementation of Jarvis and recurring automation out of scope for now.
- The next honest phase after task creation is research and analysis rather than immediate system changes.

## Open Questions

- What is the smallest durable shape of `HUMAN-DESIRE.md`?
- What counts as a reasonable inference failure versus genuine novelty?
- What should a daily per-repo Jarvis report look like in practice?
- How should this task relate to the preserved [Task-0006](/c:/Agent/CodexDashboard/Tracking/Task-0006/TASK.md) incident records and any future shared intervention canon?

## References

- [Task-0006](/c:/Agent/CodexDashboard/Tracking/Task-0006/TASK.md)
- [2026-04-16-SWE-Orchestration.md](/c:/Users/gregs/.codex/reports/2026-04-16-SWE-Orchestration.md)
