# Orchestration Prompt Templates

This directory contains reusable prompt templates for multi-agent orchestration.

When a prompt folder represents a multi-pass prompt set:

- keep `WORKFLOW.md` as the coordinator for the set
- number the agent-facing prompts as `PROMPT-PASS1.md`, `PROMPT-PASS2.md`, and so on
- note any allowed parallelism or pass dependencies inside `WORKFLOW.md`

Each prompt template now has four parts:

- `When To Use`
  - when this role is the right fit
- `Context To Inject`
  - only the values that must be injected or attached at launch because the agent cannot reliably derive them from stable local artifacts
- `Launch Notes`
  - orchestrator-facing notes such as suggested role, context breadth, write ownership, and whether the agent should be read-only
- `Runtime Prompt`
  - the exact agent-facing block to send at launch

The `Runtime Prompt` is the part that should be communicated to the agent.
The `Context To Inject` section is a launcher checklist, not an automatic transport.
If a value from that section is not injected or otherwise made available at launch, the agent does not have it.
Prefer runtime prompts that tell the agent to discover stable local context from `CURRENT-TASK.json`, the current task directory, and the standard repo docs.
Reserve `Context To Inject` for true interpolation points such as explicit pass ids, diffs, external-model responses, scope overrides, or budget constraints.
The other sections are for the human or future orchestration framework.

Current prompt templates:

- `TASK-LEADER.md` for the task-owning post-creation orchestrator that supervises the lifecycle from research through closure by dispatching specialized leaders in serial, rotating to a fresh implementation leader after each closed pass, and selectively retrieving shared intervention packet context for delegated work
- `TASK-HARVESTER.md` for mining repo and task artifacts to recommend the best next task candidates without auto-creating a new task, with candidate writeups phrased in plain English, tied to user-facing outcomes when applicable, and concrete enough to be falsifiable rather than broad area labels
- `INCIDENT-HARVESTER.md` is retired from the active workflow and kept only as a light archival marker; active intervention packet recipes now live under `C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\`
- `DOCUMENTATION-HARVESTER.md` for harvesting repo truth into user-facing docs such as `Documentation/USERS-GUIDE.md`; when used in a durable docs flow, pair its draft with the shared `CHALLENGER.md` prompt before final acceptance
- `GENERAL-DESIGNER.md` for defining the humane in-app product surface for ordinary users, including screens, entry points, tap paths, visible information, actions, and states, while steering flows away from command-line, jargon, and self-debugging assumptions
- `INTERFACE-DESIGNER.md` for pressure-testing concrete interface communication such as labels, chips, placeholders, iconography, affordance meaning, mockup-faithful component semantics, and required mockup-vs-emulator discrepancy review after the broader product flow is already known
- `VISION-HARVESTER.md` for harvesting product-vision deltas from visible misses by reconstructing what the producing agent seemed to optimize for, separating artifact repair from process repair, and encoding the smallest durable fix at the right shared or local layer
- `AUTOIMPROVEMENT-HARVESTER.md` for the daily orchestration self-improvement sweep that mines recent Codex sessions and relevant artifacts, writes a bounded problem set, researches the top problems locally, and distills a recommendation brief without silently mutating shared workflow docs
- `IMPLEMENTATION-LEADER.md` for the task-owning agent that starts only after `RESEARCH.md` exists, creates or refreshes `PLAN.md`, waits for explicit human plan approval, and then executes one approved pass with supporting agents before handing control back
- `RESEARCH-LEADER.md` for the task-owning research orchestrator that grounds the task locally, decomposes the blocking questions into a small set of decision-shaping problems, writes per-problem research artifacts, and distills a recommended direction into `RESEARCH.md`
- `RESEARCH-BRIEFER.md` for the optional research-brief worker that assembles an external-critique packet only when the human explicitly wants outside feedback after local research already exists
- `REGRESSION-LEADER.md` for the task-owning regression orchestrator that decides what regression proof is still needed, coordinates bounded regression runs, and routes the task toward closure or debugging
- `DEBUG-LEADER.md` for the task-owning debugging orchestrator that keeps the canonical bug narrative, coordinates bounded debug branches, and decides when to rerun regression or carry the task forward
- `MAKER.md` for planning and proposal work such as drafting task or plan direction
- `IMPLEMENTER.md` for a single-pass expansion worker
- `UNIT-TESTER.md` for a single-pass unit-proof and pass-audit worker
- `REGRESSION-TESTER.md` for a bounded regression worker that executes one task-level end-to-end regression run or focused rerun slice
- `DEBUG-WORKER.md` for a bounded debug worker that investigates one task-scoped branch from an active `BUG-<NNNN>.md`
- `Dream/WORKFLOW.md` for the optional intervention Dream annex flow; that prompt set now runs in ordered passes `PROMPT-PASS1.md` through `PROMPT-PASS5.md`
- `SUPERBRAIN-DEBUG-BRIEF.md` for having a local agent gather a curated bug brief such as `BUG-<NNNN>-BRIEF-<NNNN>.md` from repo evidence for a stronger downstream model
- `SUPERBRAIN-DEBUG-ANALYZER.md` for sending a prepared bug brief such as `BUG-<NNNN>-BRIEF-<NNNN>.md` to a stronger downstream model for analysis
- `SUPERBRAIN-DEBUG-FOLLOWUP.md` for having a local agent test the analyzer's hypothesis against the real repo and continue the debug path
- `CHALLENGER.md` for the agent independently stress-testing the current direction
- `RESEARCHER.md` for bounded discovery and recommendation work
- `AUDITOR.md` for the pre-commit read-only audit gate
- `INTEGRATOR.md` for reconciling outputs from other agents into one canonical result

Use these prompts together with:

- `C:\Users\gregs\.codex\AGENTS.md`
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-REPORTS.md`
- `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md`
- `C:\Users\gregs\.codex\Orchestration\Exemplars\README.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\README.md`

Prompt templates are operational assets. They should stay consistent with the shared orchestration rules, but they are not the source of truth for those rules.
