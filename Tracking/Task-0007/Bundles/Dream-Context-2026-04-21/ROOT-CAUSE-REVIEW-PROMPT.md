# Root Cause Review Prompt

Use this prompt with a large model that does not have local file access.
Give it this bundle as context.

## Prompt

You are reviewing a self-contained context bundle about one intervention-report
packet and a later regeneration of that same packet.

Your job is to determine the root causes of why the current regenerated packet
produced worse Dream solutions than the preserved reference packet.

Do not just compare wording quality. Find the causal failures in:

- process design
- pass boundaries
- prompt design
- scoring logic
- burden framing
- intervention-time measurement
- task-generation handoff
- loss of important distinctions from the source packet

The bundle contains:

- a current canonical packet under `.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/`
- a preserved stronger reference packet under `CodexDashboard/Tracking/Task-0007/Reference/ThirdPerson-2026-04-19-pre-durable-replay/ThirdPerson/2026/04/19/`
- the shared process and prompt docs that shaped both the packet rebuild and the
  Dream regeneration
- supplemental research notes that influenced the later prompt revisions

## Review Goal

Explain why the current packet still produced worse solution quality than the
reference, even after the Dream process was split into multiple passes and
updated with stronger first-principles language.

## Important Frame

Use these principles while reasoning:

- truth: reality outranks the model's story
- compassion: human frustration and repeated direction are evidence of burden
  exported by the system
- tolerance: imperfect human phrasing is normal and should not be blamed

The optimization target is reducing human input burden.
Treat repeated human input as failure telemetry by default unless the packet
shows genuine novelty, required approval, or hidden external state.

Do not blame the human for "changing the plan" unless the evidence shows an
actual contradiction. Prefer the frame that the system failed to infer or
retain an important boundary.

## Required Reading Order

Read in this order:

1. `.codex/Orchestration/Processes/FIRST-PRINCPLES.md`
2. `.codex/Orchestration/Processes/INTERVENTION-REPORTS.md`
3. `.codex/Orchestration/Processes/DREAMING.md`
4. `.codex/Orchestration/Prompts/Dream/WORKFLOW.md`
5. `.codex/Orchestration/Prompts/Dream/PROMPT-PASS1.md`
6. `.codex/Orchestration/Prompts/Dream/PROMPT-PASS2.md`
7. `.codex/Orchestration/Prompts/Dream/PROMPT-PASS3.md`
8. `.codex/Orchestration/Prompts/Dream/PROMPT-PASS4.md`
9. `.codex/Orchestration/Prompts/Dream/PROMPT-PASS5.md`
10. `.codex/Orchestration/Prompts/Interventions/HumanInputEvents/WORKFLOW.md`
11. `.codex/Orchestration/Prompts/Interventions/HumanInputEvents/PROMPT.md`
12. `.codex/Orchestration/Prompts/Interventions/HumanNeeds/WORKFLOW.md`
13. `.codex/Orchestration/Prompts/Interventions/HumanNeeds/PROMPT.md`
14. `.codex/Orchestration/Prompts/Interventions/HumanInterventionTime/WORKFLOW.md`
15. `.codex/Orchestration/Prompts/Interventions/HumanInterventionTime/PROMPT.md`
16. `.codex/Orchestration/Prompts/Interventions/HumanInterventionTime/DISPATCH-PATTERN.md`
17. the reference packet
18. the current packet
19. the two supplemental research notes

## Questions To Answer

Answer all of these:

1. What are the top 5 root causes of the quality drop from reference to current
   packet?
2. Which of those root causes are upstream of Dream and which are inside Dream?
3. Where did the process lose the human-time signal?
4. Where did the process flatten or rename an important burden in a way that
   made the final solutions worse?
5. Which prompt lines or workflow rules most directly caused the wrong priority
   order?
6. Which parts of the current process are improvements and should be kept?
7. What exact changes would most likely recover or exceed the reference quality
   on the next rerun?

## Output Format

Produce:

### 1. Executive Verdict

- one short paragraph

### 2. Root Causes

- exactly 5 items
- for each item include:
  - `Root cause`
  - `Why it happened`
  - `Where it appears`
  - `How it degraded the final solutions`

### 3. Priority Failure Analysis

- explain why the current packet ranked solutions worse than the reference
- explicitly address:
  - restart-supervision burden
  - wrong default-lane or wrong-surface closure
  - invalid proof of done

### 4. Measurement Failure Analysis

- explain whether `InterventionTime` and stall accounting were used correctly
- say where the time model distorted the ranking

### 5. Keep / Change

- `Keep`: flat list
- `Change now`: flat list

### 6. Patch Plan

- the minimum durable process and prompt changes needed before the next clean
  rerun

## Constraints

Do not:

- give generic advice
- blame the human
- hide behind "more data needed" if the bundle is enough
- treat the current packet as acceptable if the reference is clearly better

Be concrete and cite the exact files that support your claims.
