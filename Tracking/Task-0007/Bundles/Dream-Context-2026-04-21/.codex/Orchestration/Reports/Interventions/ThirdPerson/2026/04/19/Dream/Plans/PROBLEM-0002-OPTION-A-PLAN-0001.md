# Problem 0002 Option A Plan 0001

## Intent

Add an answer-first response gate so direct questions get a direct answer in the first line before any extra detail.

## What Changes

- Detect direct questions in the active turn.
- Force the first response line to answer the question directly.
- Allow supporting detail only after the first-line answer is complete.

## Files Or Artifact Types That Move

- Shared prompt rules for operator or coding agents.
- Response-shape tests or examples that cover yes/no, one-fact, and short-form answers.
- Review artifacts that flag when the answer-first rule was or was not followed.

## Rollout

1. Define the direct-question patterns that trigger the gate.
2. Write a minimal answer format for yes or no, factual, and scoped short-answer requests.
3. Add the gate to the shared response workflow used for active task work.
4. Test it against packet examples such as direct accountability questions and "small words" requests.
5. Keep the longer explanation path available only after the answer line lands first.

## Success Check

- Questions like "agree or disagree" and "what is still wrong" are answered directly in line one.
- The number of answer-shape repair turns drops.
- The rule still allows follow-up detail when needed.

## Burden Reduction Under Directional Context

`Truth`: the system commits to a clear answer instead of hiding behind prose.

`Compassion`: it removes repeated "I asked you a question" repairs.

`Tolerance`: it handles blunt or partial phrasing without demanding a more formal question first.
