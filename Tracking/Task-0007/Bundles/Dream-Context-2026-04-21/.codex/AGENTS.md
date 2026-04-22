# Global Codex Preferences

## File Link Formatting

- In Codex responses, format local filesystem links as Markdown using a leading slash before the drive letter, for example `[label](/c:/path/to/file.cpp#L13)`.
- In Codex responses, prefer absolute Windows paths in that same format so VSCodium link extensions can open them reliably.
- In Codex responses, when a line reference is useful, append `#L<number>` using the same Markdown link format.
- In durable Markdown docs stored on disk, prefer relative Markdown links such as `./PLAN.md` or `../Testing/REGRESSION-RUN-0001.md` rather than absolute `/c:/...` links.
- For durable Markdown docs, prefer file links or section links over `#L<number>` line anchors when possible, because editor handling for line anchors in source-mode Markdown is less reliable than in chat responses.

## .codex Backup Hygiene

- When working anywhere under `C:\Users\gregs\.codex`, review both `C:\Users\gregs\.codex\.gitignore` and `C:\Users\gregs\.codex\SETUP.md` after adding or changing skills, scripts, toolchains, caches, reports, or other folders.
- Keep user-authored assets tracked by default: handwritten config, docs, rules, prompts, reports, schedules, and skill source files unless the user says otherwise.
- Ignore reinstallable or generated artifacts by default: virtualenvs, `node_modules`, `__pycache__`, temp/log dirs, sandboxes, runtime sqlite files, and machine-local secrets.
- Every ignored path must stay documented in `C:\Users\gregs\.codex\SETUP.md` as one of: auto-regenerated, recreated by re-authentication, or manually rebuilt with concrete commands.
- If a tracked config references an ignored path, keep the rebuild commands for that path current in `C:\Users\gregs\.codex\SETUP.md`.
- If a new path is ambiguous, prefer asking before ignoring something that may contain user-authored work.

## Shared Docs

- `C:\Users\gregs\.codex\AGENTS.md` is the shared front door for glossary, task artifact structure, and documentation precedence across repos.
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md` defines the shared task lifecycle workflows.
- Prefer promoting generic workflow requirements to `C:\Users\gregs\.codex` instead of re-documenting them separately in each repo.
- Use repo-local docs for repo-specific commands, environment quirks, product layout, operator lanes, artifact locations, and exemplar links.
- Repo-local docs may clarify how the shared workflow maps onto a repo, but they should not redefine the shared glossary or the shared task structure.
- If a local doc change would be useful across repos, move the durable rule into shared `.codex` docs first, then trim the repo-local doc down to the repo-specific delta or a short pointer.
- Avoid duplicating long shared-process explanations in repo-local `AGENTS.md`, `TESTING.md`, or `DEBUGGING.md` when the same rule already lives in `.codex`.
- Shared `.codex` workflow docs should describe true north process, not project history.
- Project history, pass-by-pass status, bug narratives, audits, and regression-run results belong in repo task artifacts, not in shared workflow docs.
- Repo-local generic docs should follow the same split: stable process and structure at the root, historical narrative in task-owned files.

## Orchestration Directory Intent

Treat the major shared orchestration directories as different storage classes with different allowed contents.

- `C:\Users\gregs\.codex\Orchestration\Reports\`
  - promoted output artifacts only
  - examples: canonical day packets, generated digests, durable report packets
  - do not store shared workflow docs, shared prompt bundles, or reusable policy here
- `C:\Users\gregs\.codex\Orchestration\Processes\`
  - shared normative process and workflow docs
  - examples: lifecycle rules, packet structure rules, testing workflow, debugging workflow
  - do not store generated report outputs or task history here
- `C:\Users\gregs\.codex\Orchestration\Prompts\`
  - shared reusable prompt templates and prompt bundles
  - examples: agent roles, dispatch notes, pass-oriented prompt sets
  - do not store report outputs, packet-local invocation records, or task history here
- `C:\Users\gregs\.codex\Orchestration\Exemplars\`
  - reference examples of intended document shape
  - use them as examples, not as current task state or current policy history
- `Tracking/Task-<id>/` inside a repo
  - task-owned working history, evidence, pass artifacts, audits, bug narratives, and current status
  - do not promote cross-repo true-north process here

When deciding where something belongs, choose by purpose first:

- if it is a reusable rule, it belongs under `Processes/`
- if it is a reusable agent prompt or prompt bundle, it belongs under `Prompts/`
- if it is a promoted generated artifact, it belongs under `Reports/`
- if it is task history or current task evidence, it belongs under `Tracking/Task-<id>/`

If a file seems to fit more than one home, prefer the more reusable home for the rule and let the less reusable home only point to it.

## Shared Entry Points

Use these shared docs intentionally:

- `C:\Users\gregs\.codex\AGENTS.md`
  - glossary
  - task artifact structure
  - shared doc precedence
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
  - task lifecycle workflows
  - pass sequencing
  - workflow handoff points
- `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-REPORTS.md`
  - shared structure and promotion rules for canonical intervention day packets
- `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md`
  - shared standards for writing or rewriting `TASK.md`
  - writeup types such as concrete implementation, consensus, and research
  - specificity and falsifiability rules for enqueue-ready tasks
- `C:\Users\gregs\.codex\Orchestration\FILE-NAMING.md`
  - shared naming rules for orchestration docs and task-owned artifacts
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md`
  - shared contract for machine-readable task state
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json`
  - machine-checkable schema for `Tracking/Task-<id>/TASK-STATE.json`
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.md`
  - shared contract for machine-readable pass-closeout state
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.schema.json`
  - machine-checkable schema for `Testing/PASS-<NNNN>-CHECKLIST.json`
- `C:\Users\gregs\.codex\Orchestration\AUDIT-RESULT.md`
  - shared contract for machine-readable auditor results
- `C:\Users\gregs\.codex\Orchestration\AUDIT-RESULT.schema.json`
  - machine-checkable schema for `Testing/PASS-<NNNN>-AUDIT.json`
- `C:\Users\gregs\.codex\Orchestration\Prompts\README.md`
  - shared orchestration prompt-template index
- `C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\README.md`
  - shared intervention packet prompt bundles
- `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
  - shared testing artifact and naming rules
- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
  - shared debugging artifact and investigation rules
  - when a task-owned proof or regression failure turns into concrete defect tracking, open or update `Tracking/Task-<id>/BUG-<NNNN>.md` immediately
- `C:\Users\gregs\.codex\Orchestration\Processes\AUTOIMPROVEMENT.md`
  - shared daily self-improvement loop for orchestration problems, solution research, and recommendation briefs
- `C:\Users\gregs\.codex\Orchestration\Exemplars\README.md`
  - exemplar index
- `C:\Users\gregs\.codex\Orchestration\Exemplars\TASK.md`
  - intended `TASK.md` shape
- `C:\Users\gregs\.codex\Orchestration\Exemplars\TESTING.md`
  - intended repo-local `TESTING.md` shape
- `C:\Users\gregs\.codex\Orchestration\Exemplars\REGRESSION.md`
  - intended repo-root `REGRESSION.md` shape

## Workflow Precedence

Use these layers intentionally:

- shared `.codex` docs:
  - cross-repo true north for orchestration workflow, terminology, and task artifact structure
- repo-root docs:
  - repo-specific true north
  - product layout, commands, environment notes, operator lanes, and release expectations
- task-owned artifacts under `Tracking/Task-<id>/`:
  - history
  - active status
  - pass-by-pass notes
  - audits
  - bug narratives
  - executed regression-run results

If shared and repo-local docs overlap, repo-local docs should defer to `.codex`. If workflow prose and task artifacts overlap, task artifacts own the history and current evidence.

Authoritative-by-domain rule:

- shared `.codex/Orchestration/*` docs are authoritative for lifecycle, workflow, artifact roles, glossary, prompt responsibilities, and doc-precedence rules
- repo-root canonical docs are authoritative for repo-specific truth such as regression lanes, regression case definitions, commands, operator flows, and environment-specific expectations
- task-owned artifacts are authoritative for current task scope, current status, executed evidence, and which higher-level repo cases matter to the task
- task-owned artifacts must not redefine, relax, or substitute shared or repo-root canonical rules

Regression-specific rule:

- repo-root `REGRESSION.md` is the sole source of truth for regression lanes and regression pass criteria in that repo
- task-owned artifacts may reference required `REGRESSION.md` case ids and record current status, but they must not redefine what lane counts as regression or substitute supporting proof for regression closure

## Canonical Task Structure

Task-owned planning, research, audit, regression-run, and handoff artifacts should live under `Tracking/Task-<id>/`.

Preferred task structure:

- `Tracking/Task-<id>/TASK.md`
- `Tracking/Task-<id>/PLAN.md`
- `Tracking/Task-<id>/HANDOFF.md`
- recommended structured companion state file when orchestration state is being tracked explicitly:
  - `Tracking/Task-<id>/TASK-STATE.json`
- optional task-root research briefs or summaries when the task uses them:
  - `Tracking/Task-<id>/RESEARCH-PLAN.md`
  - `Tracking/Task-<id>/RESEARCH-ANALYSIS.md`
  - `Tracking/Task-<id>/RESEARCH.md`
- `Tracking/Task-<id>/Design/`
- `Tracking/Task-<id>/Testing/`
  - optional structured companions when pass-closeout state is being tracked explicitly:
    - `Tracking/Task-<id>/Testing/PASS-<NNNN>-CHECKLIST.json`
    - `Tracking/Task-<id>/Testing/PASS-<NNNN>-AUDIT.json`
- `Tracking/Task-<id>/Research/`
  - optional per-problem research artifacts such as:
    - `Tracking/Task-<id>/Research/RESEARCH-ANALYSIS-PROBLEM-<NNNN>.md`

Keep real implementation in product directories, not under `Tracking/`.

Do not recreate top-level `Testing/` or `Research/` folders outside task-owned areas unless a repo explicitly documents an exception.

A repo may explicitly reserve repo-root `Design/GENERAL-DESIGN.md` as a product-scoped design anchor. When that exception is documented by the repo:

- treat repo-root `Design/GENERAL-DESIGN.md` as the durable definition of what the product is trying to be
- keep task-owned `Design/` folders for task-scoped plans, design notes, and pass artifacts
- do not duplicate the same long-lived product intent in both places without an explicit reason

## Glossary

These terms are normative across repos.

- `human-facing surface`:
  - any workflow touched directly by a person
  - includes in-app UI, onboarding, setup, install, deploy, status, maintenance, recovery, and debug surfaces when a human is expected to perform the steps
- `operator flow` or `operator lane`:
  - a human-facing workflow for installing, deploying, configuring, validating, maintaining, or debugging a system outside the normal end-user UI
  - it is still a human-facing surface and should meet the same humane-design bar unless a repo explicitly says the lane targets trained operators
  - do not treat `operator` as shorthand for `allowed to be clunky`
- `trained operator`:
  - a human role explicitly expected to tolerate specialist tools, repeated administrative work, and deeper system knowledge
  - do not assume this role by default
- `unit test`:
  - any scripted verification that does not drive the real app UI
  - includes direct API calls, CLI-driven checks, and other automation against server/runtime artifacts
  - some of these might be called integration or smoke tests elsewhere, but in this workflow they still belong on the `unit test` side of the line rather than the `regression test` side
- `regression test`:
  - strictly in-app testing functionality
  - starts from the real app surface and exercises visible user behavior
- `end-to-end regression`:
  - a regression test that continues past the app into downstream systems or artifacts
  - for recorder/server style work, this means app or emulator or device -> upload -> downstream persistence, queue state, and when applicable worker output
- `server-only smoke`:
  - a quick scripted sanity check for packaged server entrypoints
  - useful supporting coverage, but not a `regression test`

When a request simply says `regression`, default to an in-app flow.

When a request is about scripted API or CLI proof without the app, call it `unit test` or `server-only smoke`, not `regression test`.

## Human-Facing Outcome Rule

For any task that touches a human-facing surface:

- optimize for the human-perceived outcome, not the weakest technical proxy that can be proven
- if a task says `visible`, `working`, `playable`, `readable`, `usable`, `legible`, `trustworthy`, or similar experiential language, interpret that in its ordinary human sense unless the human explicitly approves a narrower bar
- do not silently narrow the goal into lower-level claims such as:
  - the object exists in state
  - possession or input works
  - the render path emits geometry
  - a debug-only or proof-only view can see the object
  - fallback or placeholder content technically occupies the surface
- degraded fallbacks and diagnostic proofs are useful evidence, but they do not satisfy closure on their own unless the task or human explicitly says that fallback is an acceptable end state
- when current implementation truth is only a fallback, record that gap explicitly as remaining debt, open bug, or incomplete acceptance rather than treating the task as finished
