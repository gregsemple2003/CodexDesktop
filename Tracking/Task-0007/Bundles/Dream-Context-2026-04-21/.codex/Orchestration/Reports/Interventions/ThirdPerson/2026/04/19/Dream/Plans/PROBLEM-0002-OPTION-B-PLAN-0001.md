# Problem 0002 Option B Plan 0001

## Intent

Add a strict short-answer mode that activates when the human explicitly asks for short answers, small words, or one-fact replies.

## What Changes

- Detect short-answer instructions in the active turn.
- Constrain the next answer to one short sentence or a yes/no plus one fact.
- Delay any elaboration until the human asks for it.

## Files Or Artifact Types That Move

- Shared prompt instructions for response length control.
- Test prompts and review examples for short-answer mode.
- Optional task-state notes that record when short-answer mode is active.

## Rollout

1. Define the trigger phrases for short-answer mode.
2. Add strict output caps for the next reply.
3. Test the mode on packet cases that used "short answers" and "small words."
4. Review whether the mode helps without causing harmful under-answering.

## Success Check

- Explicit short-answer requests produce visibly shorter replies.
- Overlong explanatory replies after those requests disappear.
- The mode can be turned off cleanly after the direct answer is delivered.

## Burden Reduction Under Directional Context

`Truth`: it prevents hedging by forcing a clear small answer.

`Compassion`: it cuts the effort needed to strip away extra framing.

`Tolerance`: it respects the human's preferred answer shape once stated, even if the wording is abrupt.
