# HIE/HIT To Dream Burden Loss Structure

## Objective

Determine exactly how burden signal is lost between:

1. `HumanInputEvents`
2. `HumanNeeds`
3. `HumanInterventionTime`
4. `Dream`

for the current rebuilt April 19 packet.

## Verdict

The loss is real, and it is staged.

The current packet does not simply have "weaker Dream analysis."

It loses burden depth at four different points:

1. `HumanNeeds` sampling and prioritization
2. `HumanInterventionTime` measurement under fidelity gaps
3. Dream burden decomposition
4. task proposal concretization

The dominant pattern is:

- lane-proof burden survives
- answer-shape burden survives
- approval burden survives but is demoted
- ownership/restart-supervision burden is weakened
- debugging/root-cause burden is weakened or dropped
- durable-learning burden disappears

## Stage 0: Raw HIE Still Contains More Burden Than Dream Admits

Raw `HumanInputEvents/INDEX.json` still contains distinct evidence for at least six burden families.

### 1. Lane / proof-surface burden

Examples in `HumanInputEvents/INDEX.json`:

- `...-3325` invisible-pawn regression failure and explicit time-loss complaint
- `...-3431` human-default-lane regression rule
- `...-4909` invalid screenshot / invalid evidence surface
- `...-7038` explicit instruction to remain on the runtime human default lane

See line anchors:

- `...-3325` at lines `23-27`
- `...-3431` at lines `47-51`
- `...-4909` at lines `295-299`
- `...-7038` at lines `479-483`

### 2. Approval-surface burden

Examples:

- `...-4379` no diff for `plan.md`
- `...-4399` asks for links explicitly
- `...-4512` describes the deeper failure: no human-suitable diff, no context, bad links
- `...-4919` explicit approval gate

See:

- `...-4379` at lines `199-203`
- `...-4919` at lines `271-275`

### 3. Ownership / restart-supervision burden

Examples:

- `...-4015` says not to take over subagent work and to be generous with time
- `...-5272` “Why are you messaging me?”
- `...-5281` “What is a real stop point?”
- `...-5290` and `...-5300` reject weak stop reasoning
- `...-5310` “keep working without throwing your homework at my feet”
- `...-6558` explicit wake-up
- `...-7038` explicit continue-under-existing-ownership instruction

See:

- `...-6558` at lines `407-411`
- `...-5272` at lines `439-443`
- `...-5310` at lines `471-475`
- `...-7038` at lines `479-483`

### 4. Answer-shape burden

Examples:

- `...-3392` “STOP. I asked you a question.”
- `...-3431` agree/disagree question
- `...-3473` short-answer demand
- `...-5073` “Now answer my questions.”
- `...-5090` “Stop evading the question.”

Representative anchors:

- `...-3431` at lines `47-51`
- `...-3392` is adjacent in the same early cluster

### 5. Debugging / root-cause burden

Examples:

- `...-5340` explicit root-cause method
- `...-7446` first-disagreement tracing instruction

See:

- `...-5340` at lines `487-491`

### 6. Durable-learning / repeated-correction burden

Examples:

- `...-3463` says the lack of judgment required course correction and asks for a process fix
- `...-3482` asks for a durable fix
- `...-3549` and `...-3585` correct over-specific or bad process wording
- `...-4512` says “for later analysis” and names the dropped ball explicitly

These are not just one-off corrections. They are lessons the packet is trying to promote.

## Stage 1: HumanNeeds Keeps Some Burdens But Already Starts Narrowing

`HumanNeeds/PACKET-TRIAGE.json` still carries five packet-level needs:

- `NEED-01` trustworthy default-lane proof
- `NEED-02` hard stop/go semantics
- `NEED-03` direct answers
- `NEED-04` low-friction approval surfaces
- `NEED-05` root-cause debugging discipline

See lines:

- `NEED-01` through `NEED-05` at lines `7`, `13`, `19`, `25`, `31`

So the first major loss is not that `HumanNeeds` forgets everything.

The first loss is narrower:

- it still centers `NEED-01` through the highest weight in `PACKET-RECORD.json`
- it preserves the other needs, but only as lower-weight neighbors around that center
- it still does not preserve durable learning as a separate need

After this analysis, the `HumanNeeds` contract was changed to a single weighted `needs` array precisely to remove the stronger `selected_need` collapse.

Even with that improvement, the structural risk remains:

- lane-proof can still become the main lens
- the lower-weight needs can still be treated as supporting frictions instead of coexisting burden families

See:

- `PACKET-RECORD.json` `needs` weights in the current packet

### Representative-event loss

The selected representative events are:

- `3325`
- `3431`
- `3463`
- `4379`
- `4909`
- `5340`
- `3692`
- `4919`

See `REPRESENTATIVE-EVENTS.json` lines:

- `22`, `30`, `38`, `46`, `62`, `78`, `86`, `94`

What is missing from that representative set:

- `6558` wake-up
- `7038` ownership restart
- `5272` “Why are you messaging me?”
- `5310` “throwing your homework at my feet”

Those missing events are exactly the strongest ownership / restart-supervision signals in raw `HIE`.

So the first concrete structural loss is:

`HumanNeeds` keeps a stop/go concept, but it drops the stronger ownership / wake-up cluster from the representative surface.

## Stage 2: HIT Quantifies The Wrong Burdens

The current `HumanInterventionTime/SUMMARY.json` shows:

- `stall_event_count = 0`
- `new_request = 38`
- `correction = 9`
- `wake_up = 0`
- `ownership = 13`
- `learning = 4`

See lines:

- `event_count`, `stall_event_count`, `typing_rate_chars_per_second`, `intervention_time_seconds` at lines `6-14`
- `new_request`, `correction`, `boundary_reset`, `answer_to_question`, `wake_up` at lines `17-21`
- `ownership`, `learning` at lines `29-30`

The file itself explains why:

- the session transcript for `46/62` events is missing
- stall attribution is conservative and often zeroed

See line `35`.

This creates the second major structural loss:

- ownership burden is visible only as a tag, not as charged time
- learning burden is nearly erased (`4`)
- wake-up burden is completely erased (`0`)
- the day is re-described as mostly `new_request`

That is a major distortion of the source packet.

In practical terms, it turns:

- "the system stopped supervising and made the human resume ownership"

into:

- "the human made many new requests"

## Stage 3: Dream Collapses What Remains Into Three Neat Problems

The current `Dream/BURDEN-ANALYSIS.md` says the packet is driven by only three primary failure modes:

1. proof-lane mismatch
2. stop/go and approval gate drift
3. answer-shape failures

See line `3`.

That means the Dream burden layer has already collapsed:

- approval into stop/go
- ownership into stop/go
- debugging into a supporting truth note
- durable learning into nothing

The same collapse shows up in the matrix:

- `PROBLEM-0001` proof lane
- `PROBLEM-0002` stop/pause and approval gates
- `PROBLEM-0003` direct questions

See `ORTHOGONAL-SOLUTIONS-MATRIX.md` lines `3-24`.

So the third structural loss is:

- `HumanNeeds` had five candidate needs
- Dream turns them into three problems

The lost distinctions are:

1. `approval` no longer stands on its own
2. `debugging` no longer stands on its own
3. `ownership/restart supervision` is weakened into pause etiquette
4. `durable learning` disappears entirely

## Stage 4: Option Tasks Preserve Shape Better Than Before, But Still Flatten Substance

The new option-task format is better than previous thin winner tasks, but the actual proposals are still flattened.

All three top option tasks reuse almost the same burden-reduction contract:

- burden being reduced: repeated human corrections about lane/proof/boundaries/answers
- causal claim: lack of durable gates and templates
- human relief: fewer STOP/answer/proof interventions

See:

- `PROBLEM-0001-OPTION-A.md` lines `25-34`
- `PROBLEM-0002-OPTION-A.md` lines `24-33`
- `PROBLEM-0003-OPTION-A.md` lines `24-33`

This is the fourth loss:

the task format now has more fields, but the actual burden stories are still being collapsed into one generic story.

That generic story is too coarse to recover:

- continuity state as a durable contract
- approval packet generation as a separate burden
- intervention lesson promotion as a separate burden
- root-cause verification as a separate burden

## Definitive Loss Structure

The burden-loss structure is:

1. **Raw HIE contains at least six burden families.**
2. **HumanNeeds keeps five candidate needs, but the representative-event set drops the strongest ownership/wake-up cluster.**
3. **HIT under-measures ownership and learning because the missing transcript zeros out wake-up/stall signal and reclassifies too much into `new_request`.**
4. **Dream compresses five-plus burden families into three problems.**
5. **Option tasks preserve formatting better than before, but they inherit the compressed burden map and therefore remain generic.**

This means the deepest current bottleneck is not task formatting.

It is the loss of burden distinction between `HIE/HIT` and `Dream`.

## What Survives And What Dies

### Preserved

- lane / proof-surface burden
- answer-shape burden

### Partially preserved

- approval burden
- stop/go burden

### Lost or badly weakened

- ownership / restart-supervision burden
- debugging / root-cause burden as a first-class problem
- durable-learning / repeated-correction burden

## Practical Implication

If the next correction target is chosen correctly, it should not start at task wording.

It should start by forcing Dream to answer:

1. What burden families are present in raw `HIE`?
2. Which of those are preserved in `HumanNeeds`?
3. Which of those are preserved or distorted in `HIT`?
4. Which of those disappear before the matrix is written?

Without that audit, task quality will keep regressing because the task layer will only ever receive a thinned burden model.
