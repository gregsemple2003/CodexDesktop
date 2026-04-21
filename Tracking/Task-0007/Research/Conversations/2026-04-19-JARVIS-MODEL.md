# 2026-04-19 Jarvis Model Capture

## Context

This note captures the task-creation conversation that motivated `Task-0007`.

## Human Framing

The human direction was:

- treat every input to Codex as a failure signal
- delegate intervention work to an agent specialized in understanding human concerns
- call that agent `Jarvis`
- require each intervention analysis to end with:
  - `How could the system have inferred the need for the input?`
- allow solutions to respect local human constraints
- maximize truthfulness, compassion, and tolerance in the world and in the systems
- produce a daily report for each repo that asks:
  - how to advance the repo's core need
  - how the system could have inferred the need for the intervention
  - what task proposals follow
- allow a future `HUMAN-DESIRE.md` to express high-level human desire

## Initial Analytical Read

The immediate strength of the model is that it forces every intervention to become learning material instead of disappearing into chat history.

The immediate risk is that it can collapse distinct cases together:

- forgotten preference
- weak context modeling
- missing initiative
- genuine novelty
- changed human desire

## Follow-On Questions

- What evidence standard should `Jarvis` use before claiming it knows the human's need?
- What belongs in a daily repo report versus a task-local note?
- How should a future `HUMAN-DESIRE.md` constrain the agent instead of becoming a vague values slogan?
