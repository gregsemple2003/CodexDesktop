# File Naming

This file defines the shared naming rules for orchestration docs and task-owned artifacts.

Use it together with:

- `C:\Users\gregs\.codex\AGENTS.md` for shared glossary, task structure, and precedence
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md` for lifecycle workflows
- `C:\Users\gregs\.codex\Orchestration\Exemplars\README.md` for concrete examples

## Scope

Apply these rules to:

- shared orchestration docs under `C:\Users\gregs\.codex\Orchestration\`
- task-owned docs and artifacts under `Tracking/Task-<id>/`

Do not apply these rules blindly to:

- repo implementation code or package/module names
- third-party or mirrored reference trees under `Research/`
- tool-managed folders such as `.git`, `node_modules`, or `__pycache__`
- immutable history outputs such as archived session transcripts

## Directory Names

Use single-cap first directory names for human-owned workflow folders:

- `Orchestration/`
- `Processes/`
- `Exemplars/`
- `Tracking/Task-0001/`
- `Design/`
- `Testing/`
- `Research/`

Use `Task-0001`, not `task-0001` or `Task_0001`.

## Canonical Singleton Docs

Use all-caps base filenames for canonical singleton docs:

- `AGENTS.md`
- `ORCHESTRATION.md`
- `Processes/TASK-CREATE.md`
- `FILE-NAMING.md`
- `TASK-STATE.md`
- `TASK.md`
- `PLAN.md`
- `HANDOFF.md`
- `RESEARCH-ANALYSIS.md`
- `TESTING.md`
- `DEBUGGING.md`
- `REGRESSION.md`

## Structured State Files

Use these names for shared task-state files:

- `CURRENT-TASK.json`
- `CURRENT-TASK.schema.json`
- `TASK-STATE.json`
- `TASK-STATE.schema.json`
- `PASS-CHECKLIST.json`
- `PASS-CHECKLIST.schema.json`
- `AUDIT-RESULT.json`
- `AUDIT-RESULT.schema.json`

Rules:

- keep the base name in all caps with `-` separators
- use `snake_case` for JSON keys
- use `snake_case` for enum values
- prefer a standard `.schema.json` suffix for machine-checkable schemas

## Compound Filenames

When a filename has multiple words, separate words with `-`, not `_`.

Examples:

- `GENERAL-DESIGN.md`
- `RESEARCH-PLAN.md`
- `RESEARCH-ANALYSIS-PROBLEM-0001.md`
- `AUTOIMPROVEMENT-PROBLEMS.md`
- `AUTOIMPROVEMENT-PROBLEM-0001.md`
- `AUTOIMPROVEMENT-BRIEF.md`
- `PASS-SCAFFOLD-PLAN.md`

## Ordinal Artifacts

When an artifact has an ordinal, use:

- all-caps words
- `-` between words
- a four-digit ordinal

Examples:

- `PASS-0003-AUDIT.md`
- `PASS-TEST-0003.md`
- `REGRESSION-RUN-0001.md`
- `BUG-0001.md`
- `PASS-0001-CHECKLIST.json`
- `PASS-0001-AUDIT.json`

For pass-owned closeout artifacts, make the pass id the stable key and put the artifact type after it.

Examples:

- `PASS-0003-AUDIT.md`
- `PASS-0003-AUDIT.json`
- `PASS-0003-CHECKLIST.json`

## Reference Updates

When renaming a shared orchestration doc or task-owned artifact:

- update live docs and active artifacts in the same change
- do not leave old path references behind in current guidance
- preserve historical context, but do not let current docs point at stale names

## Markdown Links

For durable Markdown docs stored on disk:

- prefer relative Markdown links computed from the current file, such as `./PLAN.md`, `../INDEX.md`, or `./Task-Candidates/INDEX.md`
- use `/c:/...` absolute Markdown links in Codex responses, not in authored repo docs
- prefer file links or heading links inside durable docs over `#L<number>` line anchors when possible, because editor support for line anchors is less reliable in source-mode Markdown

## Windows Note

On Windows, case-only renames may require an intermediate temp name or temp holding folder. Prefer the safe route over assuming the filesystem will accept a direct case-only rename.
