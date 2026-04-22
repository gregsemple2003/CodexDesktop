# Problem 0002 Option C Plan 0001

## Intent

Route hard questions into a short clarification step before answering, but only when the question is materially ambiguous.

## What Changes

- Add an ambiguity check before answering.
- Ask one narrow clarification question when the system cannot answer honestly without choosing between competing meanings.
- Skip the clarification step when the question is already direct enough to answer.

## Files Or Artifact Types That Move

- Shared question triage rules.
- Prompt examples that separate answerable questions from truly ambiguous ones.
- Review notes that explain why a clarification was required.

## Rollout

1. Define the ambiguity threshold.
2. Limit the clarification step to one short question.
3. Test it against packet examples to ensure obvious direct questions do not trigger it.
4. Keep the route available for the small set of cases where truth would otherwise suffer.

## Success Check

- The system asks clarifying questions only when a direct answer would be misleading.
- Obvious direct questions are still answered directly.
- Clarification use stays rare.

## Burden Reduction Under Directional Context

`Truth`: it prevents false certainty when the meaning is truly split.

`Compassion`: it tries to avoid longer repair loops caused by a wrong initial answer.

`Tolerance`: it treats imperfect phrasing as normal, but still protects against real ambiguity.
