# Orthogonal Solutions Matrix

## Problems
- PROBLEM-0001: Regression proof is not reliably constrained to the human default lane; off-lane evidence can be mislabeled as regression closure.
- PROBLEM-0002: STOP/PAUSE and approval gates are not treated as hard execution state; drift forces repeated boundary resets.
- PROBLEM-0003: Direct questions are not answered directly; human must re-ask to get the needed shape.

## Options
### PROBLEM-0001
- OPTION-A: Default-lane regression proof gate (contract + checklist; refuse closure without lane id + artifact).
- OPTION-B: Default-lane runner standardization (tooling that produces canonical artifacts and makes lane overrides explicit).

### PROBLEM-0002
- OPTION-A: Explicit STOP/PAUSE protocol as a durable execution state (no tools while paused).
- OPTION-B: Gate-aware approval surface template (diff + links + one explicit approval question).

### PROBLEM-0003
- OPTION-A: Direct Answer First rule for agree/disagree and yes/no questions.
- OPTION-B: Question capture checklist for multi-question messages (answer every extracted question before proceeding).

## Winners And Rollout Order
1. PROBLEM-0001 OPTION-A then OPTION-B. Contract first; tooling second.
2. PROBLEM-0002 OPTION-A then OPTION-B. Stop/go correctness first; approval friction second.
3. PROBLEM-0003 OPTION-A now; OPTION-B only if question omissions persist.

## Option Tasks
- [Option-Tasks/INDEX.md](./Option-Tasks/INDEX.md)
