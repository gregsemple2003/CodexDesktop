# Problem 0002 Option B Plan 0001

## Summary

Add a reply-shape preflight that detects explicit questions and forces the first line to answer them directly.

## What Changes

- detect direct questions in the active human turn
- require the response draft to open with the answer, not context
- flag unanswered or partially answered questions before a longer reply is sent

## Files Or Artifact Types

- response-planning prompt or validator
- any local reply-audit helper used before final output
- examples for yes or no, factual, and short definitional questions

## Rollout

1. Define the question patterns that count as explicit asks.
2. Add a preflight check that lists unanswered questions.
3. Fail the check when the opening line does not answer the active question.
4. Test the rule on April 19 examples that triggered `STOP. I asked you a question.`

## Success Checks

- explicit questions are answered on line one in future runs
- support detail can follow, but the opening line is direct
- the check catches evasive answer shapes before send time
