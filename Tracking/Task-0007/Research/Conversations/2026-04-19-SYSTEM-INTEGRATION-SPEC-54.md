## User

Come up with a spec for how this system will integrate with Codex.

Current system requires too much human intervention.
Delegate intervention work to agents that are specialized to understand human concerns.  Call the agent Jarvis.
Input surface is every input by the human.  All chat, all PR comments.  Each input requires a chain of thought that ends with:

How could the system have inferred the need for the input?
The solution space should allow local human constraints.

I want all solutions to maximize truthfulness, compassion, tolerance in the world and in the systems.
As a consequence of the above, Jarvis should produce a daily report for each repo:

How can we advance the core need of the repo?  Express in task proposals.  Human can specify HUMAN-DESIRE.md to express the human desire at a high-level.
How could the system have inferred the need for the input?
Inference based on human intervention plus model for human values based on extensive research of meat-based analytical systems.

## Assistant

I’ll turn this into an implementable integration spec for Jarvis and Codex, covering architecture, workflows, repo artifacts, reporting, safeguards, and success criteria. I’ll treat the requested “chain of thought” as an auditable rationale artifact, and ground Codex-specific assumptions in current official documentation.

I’ve now pinned the design to Codex’s real integration surfaces—GitHub reviews/comments, tasks/automations, AGENTS.md, CI/SDK/MCP, hooks, and subagents. Jarvis will sit above those primitives as the intervention-inference layer, treating each human input as unmet-need telemetry while Codex remains the execution and review engine.

Thought for 5m 37s
Spec: Jarvis × Codex integration for reducing human intervention

Status: v0.1
Core idea: Jarvis observes every human input, infers the unmet need behind it, converts preventable future interventions into Codex-executable work, and produces a daily repo report that proposes concrete tasks aligned with the repo’s human desire.

As of April 19, 2026, Codex is a suitable execution layer because it can write and edit code, understand codebases, run in cloud environments, create PRs, review PRs from GitHub, run in CI/CD via codex exec or GitHub Actions, and accept repository instructions via AGENTS.md. Codex also supports programmatic control through the Codex SDK and lifecycle extension via hooks. 
OpenAI Developers
+4
OpenAI Developers
+4
OpenAI Developers
+4

1. Purpose

The current system requires too much human intervention. The goal is to make every human intervention useful twice:

Locally useful: address the immediate input.

Systemically useful: infer how the system could have avoided needing that input, then create task proposals or Codex tasks that reduce future intervention.

Jarvis is the agent responsible for the second function.

Jarvis should not replace humans. Jarvis should reduce unnecessary human burden while preserving human authority, local context, and moral pluralism.

2. System roles
2.1 Jarvis

Jarvis is a specialized human-concern inference agent. Its job is to answer:

What human need caused this input, and what should the system change so the human does not need to repeat it?

Jarvis observes human inputs from chat, GitHub PR comments, issue comments, review comments, commit discussion, and eventually support threads, incident notes, and design docs.

Jarvis produces:

Intervention Inference Records

Task proposals

Codex prompts

Daily repo reports

Updates to repo-local constraints or instructions

Escalation questions when the system cannot infer safely

2.2 Codex

Codex is the engineering execution layer. Jarvis should delegate implementation, review, test, refactor, and documentation tasks to Codex when the inferred need is actionable in a repository.

Codex can be invoked through several surfaces:

GitHub PR comments: @codex review requests a code review, and other @codex ... comments can start cloud tasks using PR context. 
OpenAI Developers

Automatic PR reviews: Codex can be configured to review every PR automatically. 
OpenAI Developers

Codex cloud: Codex can work in background cloud environments, including parallel work, and can create PRs from its work. 
OpenAI Developers

Codex SDK: server-side applications can programmatically start and resume Codex threads. 
OpenAI Developers

Codex GitHub Action: GitHub workflows can run Codex in CI/CD, use prompt files, apply patches, and post reviews. 
OpenAI Developers

Codex hooks: Codex can run deterministic scripts at lifecycle events such as UserPromptSubmit and Stop, though hooks are experimental and currently not available on Windows. 
OpenAI Developers

Codex subagents: Codex can spawn specialized agents in parallel and consolidate results, which is useful for separate review lenses such as security, tests, maintainability, and regressions. 
OpenAI Developers

3. Core design principle

Every human input is treated as a signal of unmet system context.

A human comment like “this needs a test” is not only a review comment. It is evidence that the system failed to infer one or more of:

the expected test policy,

the risk profile of the changed code,

the local repo convention,

the reviewer’s unstated acceptance criteria,

the repo’s human desire,

or the appropriate level of Codex self-review before opening a PR.

Jarvis therefore records both the immediate meaning and the preventability question.

4. Chain-of-thought requirement

The user requirement says:

Each input requires a chain of thought that ends with: “How could the system have inferred the need for the input?”

For product safety, privacy, and auditability, the persisted artifact should be a structured reasoning summary, not raw hidden chain-of-thought. Jarvis may deliberate internally, but the durable record should expose evidence, assumptions, uncertainty, and conclusions.

The required final field is preserved exactly:

Markdown
## How could the system have inferred the need for the input?
...

This gives humans a useful audit trail without storing fragile or overly private internal reasoning.

5. Repository files
5.1 HUMAN-DESIRE.md

A repo-level declaration of the human’s high-level desire.

Example:

Markdown
# HUMAN-DESIRE.md

## Core need of this repo

This repo exists to help maintainers ship reliable, humane automation that reduces repetitive human intervention while preserving user agency.

## Values

- Truthfulness: prefer evidence, citations, tests, and uncertainty over confident guesses.
- Compassion: reduce blame, toil, surprise, and emotional labor.
- Tolerance: support local differences in workflow, language, culture, and risk tolerance.

## Success

- Maintainers spend less time repeating the same instructions.
- PRs arrive with tests, docs, and rationale appropriate to their risk.
- The system explains how it inferred tasks and where uncertainty remains.

## Non-goals

- Do not silently make high-risk production changes.
- Do not optimize for velocity by hiding uncertainty.
- Do not collapse all humans into one universal preference model.
5.2 HUMAN-CONSTRAINTS.md

Optional local constraints. This should support repo, directory, service, team, and file-area scoping.

Example:

Markdown
# HUMAN-CONSTRAINTS.md

## Local constraints

- Payment code requires explicit security review.
- Public APIs require docs and changelog entries.
- Avoid new production dependencies unless a human approves.
- Comments should avoid blaming prior authors.
- Prefer small PRs under 400 changed lines unless the task is mechanical.
5.3 AGENTS.md

Codex already reads AGENTS.md before work and can layer global, project, and nested instructions; closer files can override broader guidance. 
OpenAI Developers

Jarvis should maintain a generated section in AGENTS.md, or propose changes to it, so Codex inherits the human desire and local constraints.

Example generated section:

Markdown
## Jarvis-derived working agreements

- Before opening a PR, explain which human intervention pattern this change reduces.
- Include tests when changing behavior.
- For ambiguous human comments, prefer a task proposal over implementation.
- In reviews, check whether the PR satisfies HUMAN-DESIRE.md.
- End implementation summaries with:
  "How could the system have inferred the need for the input?"
6. Input surfaces

Jarvis should ingest every human input relevant to a repo.

Initial surfaces:

Chat messages to the system.

GitHub PR comments.

GitHub review comments.

GitHub issue comments.

Commit discussion comments.

Explicit @jarvis commands.

Explicit @codex comments, because those indicate a human had to manually delegate work.

PR approvals, requested changes, and dismissals.

CI rerun requests, because those often indicate missing flaky-test inference or unclear failure diagnosis.

Later surfaces:

Slack or Linear integrations.

Incident retrospectives.

Support tickets.

Docs feedback.

User research notes.

7. Event model

Every human input becomes a HumanInputEvent.

JSON
{
  "event_id": "evt_123",
  "repo": "org/repo",
  "source": "github_pr_comment",
  "actor": {
    "id": "github-user-id",
    "role_hint": "maintainer"
  },
  "created_at": "2026-04-19T14:32:00-04:00",
  "location": {
    "pr": 42,
    "issue": null,
    "file": "src/auth/session.ts",
    "line": 88
  },
  "raw_text": "This needs a regression test for expired sessions.",
  "related_artifacts": {
    "pr_diff_sha": "abc123",
    "commit_sha": "def456",
    "ci_runs": ["ci_789"]
  }
}
8. Intervention Inference Record

For every HumanInputEvent, Jarvis creates an InterventionInferenceRecord.

JSON
{
  "record_id": "iir_123",
  "event_id": "evt_123",
  "repo": "org/repo",
  "input_summary": "Maintainer requested a regression test for expired sessions.",
  "human_need": {
    "label": "risk_coverage",
    "description": "The maintainer wants confidence that expired session behavior will not regress.",
    "confidence": 0.86
  },
  "local_constraints": [
    "Behavior changes require regression tests.",
    "Auth code requires extra scrutiny."
  ],
  "evidence": [
    "Comment mentions regression test.",
    "Changed file is under src/auth.",
    "HUMAN-CONSTRAINTS.md says auth code requires extra scrutiny."
  ],
  "missed_system_signal": {
    "type": "changed_behavior_without_test",
    "description": "The PR changed session expiration logic without adding or updating a relevant test."
  },
  "recommended_action": {
    "type": "delegate_to_codex",
    "task_title": "Add regression test for expired session behavior",
    "risk": "low",
    "requires_human_approval": false
  },
  "truthfulness_notes": [
    "Inference is based on the comment and changed file path.",
    "Jarvis should not infer the maintainer's broader intent beyond test coverage."
  ],
  "compassion_notes": [
    "Frame as a missing system inference, not reviewer nitpicking or author failure."
  ],
  "tolerance_notes": [
    "Respect repo-local auth testing expectations even if other repos would not require this test."
  ],
  "how_could_the_system_have_inferred_the_need_for_the_input": "The system could have detected that auth behavior changed without a nearby regression test, combined that with the repo constraint that auth changes require extra scrutiny, and asked Codex to add or propose a test before human review."
}
9. Inference taxonomy

Jarvis should classify each human input into one or more intervention types.

Type	Human input example	Inferred system miss	Typical action
Missing test	“Add a regression test.”	Changed behavior lacked coverage.	Codex adds test.
Missing docs	“Update README.”	Public behavior changed without docs.	Codex updates docs.
Local convention	“We don’t use mocks here.”	Repo pattern not learned.	Update AGENTS.md; Codex refactors.
Safety concern	“This leaks PII.”	Privacy rule not applied.	Block merge; security task.
Ambiguity	“This feels wrong.”	System cannot identify concern.	Ask clarifying question or propose review dimensions.
Tone concern	“This comment sounds harsh.”	Compassion/tolerance violation.	Codex rewrites text.
Repeated manual task	“Again, please update changelog.”	Durable rule missing.	Add repo rule and checklist.
Review fatigue	“I already explained this.”	Prior intervention not converted into memory/rule.	Backfill inference record and task.
Over-automation	“Don’t do this automatically.”	Human boundary crossed.	Add autonomy constraint.
Desire mismatch	“This solves the wrong problem.”	HUMAN-DESIRE.md not applied.	Reframe task proposal.
10. Codex delegation contract

Jarvis should delegate to Codex only after creating an inference record.

A Codex task must include:

Markdown
# Codex Task

## Objective

Add a regression test for expired session behavior.

## Human need

The maintainer wants confidence that expired session behavior will not regress.

## Evidence

- PR comment: "This needs a regression test for expired sessions."
- Changed file: src/auth/session.ts
- Repo constraint: auth code requires extra scrutiny.

## Local constraints

- Follow HUMAN-CONSTRAINTS.md.
- Follow AGENTS.md.
- Do not add production dependencies.
- Keep the PR small.

## Acceptance criteria

- Adds or updates a test that fails on the previous behavior and passes on the new behavior.
- Does not change production behavior unless required for testability.
- Runs the relevant test command.
- Summarizes the change and uncertainty.

## Required closing section

How could the system have inferred the need for the input?

Explain how future systems could detect this need before a human comment is required.

Codex should then either:

create a patch,

create a PR,

post a review,

or report that the task is unsafe or underspecified.

11. Delegation policy

Jarvis should not delegate everything automatically.

Auto-delegate to Codex

Low-risk, reversible work:

docs updates,

tests,

lint fixes,

small refactors,

changelog updates,

mechanical consistency fixes,

PR review comments asking for clear code changes.

Propose first, then wait

Medium-risk work:

public API changes,

dependency additions,

database migrations,

security-sensitive code,

privacy-sensitive code,

performance-sensitive rewrites,

behavior changes with unclear product intent.

Never auto-delegate

High-risk or human-meaning-sensitive work:

changes to production credentials,

legal/policy commitments,

firing, hiring, compensation, or interpersonal judgments,

irreversible data changes,

changes that conflict with explicit human constraints,

attempts to infer private mental states beyond observable evidence.

Codex supports sandboxing and approval controls, and automation should use the least permissions needed. Codex docs warn against dangerous full-access modes except in controlled environments, and the GitHub Action docs recommend limiting who can start workflows, sanitizing PR-derived prompt inputs, protecting API keys, and avoiding unsafe modes on multi-tenant runners. 
OpenAI Developers
+1

12. Daily repo report

Jarvis produces one daily report per repo.

Default destination:

GitHub issue comment on a rolling “Jarvis Daily Report” issue,

or a dated Markdown file under .jarvis/reports/YYYY-MM-DD.md,

or both.

Report title:

Markdown
# Jarvis Daily Report — org/repo — 2026-04-19

Report body:

Markdown
## Core need of the repo

Derived from HUMAN-DESIRE.md:

> This repo exists to reduce repetitive human intervention while preserving user agency.

## Summary

- Human inputs observed: 37
- Preventable interventions: 22
- Codex-ready tasks: 9
- Needs human clarification: 4
- Suggested AGENTS.md/HUMAN-CONSTRAINTS.md updates: 3

## Top task proposals

### 1. Add auth regression-test rule to AGENTS.md

**Need:** Maintainers repeatedly request regression tests for auth behavior.  
**Evidence:** 4 PR comments across 3 PRs this week.  
**Proposal:** Add an auth-specific review guideline and create a Codex task to backfill missing tests.  
**Risk:** Low.  
**Codex action:** Open PR modifying `AGENTS.md` and adding tests where obvious.  
**Acceptance criteria:** Future auth PRs include regression tests or explicitly explain why not.

**How could the system have inferred the need for the input?**  
The system could have detected repeated reviewer comments containing “regression test” on files under `src/auth`, then generalized the pattern into a local repo rule.

### 2. Update public API docs checklist

...

## Repeated intervention patterns

| Pattern | Count | Suggested prevention |
|---|---:|---|
| Missing tests | 8 | Add Codex pre-review task. |
| Missing docs | 5 | Add public API docs rule. |
| Unclear PR summaries | 4 | Generate PR summary template. |

## Local human constraints discovered

- Auth changes require security-oriented review.
- Maintainers prefer small PRs.
- Review language should be direct but non-blaming.

## Proposed updates to repo instructions

```diff
+ ## Review guidelines
+ - For auth behavior changes, require regression tests or an explicit rationale.
+ - PR summaries must include user-visible behavior changes.
Questions for humans

Should Jarvis auto-open PRs for missing test-only changes, or only propose them?

Should docs typos be considered P1 in Codex review for this repo?

Truthfulness / compassion / tolerance audit

Truthfulness: All task proposals cite observable evidence.

Compassion: No proposal assigns blame to individual contributors.

Tolerance: Local repo preferences are represented as scoped constraints, not universal rules.

---

## 13. Daily report task proposal schema

```json
{
  "proposal_id": "tp_123",
  "repo": "org/repo",
  "title": "Add auth regression-test rule to AGENTS.md",
  "core_need_alignment": "Reduces repeated human requests for tests in auth PRs.",
  "human_need": "Confidence in auth behavior correctness.",
  "evidence": [
    "4 reviewer comments requested auth regression tests.",
    "HUMAN-CONSTRAINTS.md marks auth as sensitive."
  ],
  "codex_prompt": "...",
  "acceptance_criteria": [
    "AGENTS.md contains scoped auth testing guidance.",
    "Existing obvious missing tests are added.",
    "Codex summary includes inference explanation."
  ],
  "risk": "low",
  "autonomy": "auto_pr_allowed",
  "local_constraints": [
    "No new production dependencies.",
    "Keep PR small."
  ],
  "how_could_the_system_have_inferred_the_need_for_the_input": "Repeated auth-related review comments plus sensitive-path constraints should trigger a durable test rule."
}
14. Jarvis architecture
                    ┌────────────────────┐
                    │ Human input surfaces│
                    │ chat, PRs, issues   │
                    └─────────┬──────────┘
                              │
                              ▼
                    ┌────────────────────┐
                    │ Input ingestion     │
                    │ normalize + enrich  │
                    └─────────┬──────────┘
                              │
                              ▼
                    ┌────────────────────┐
                    │ Context retrieval   │
                    │ repo, diff, CI, docs│
                    │ HUMAN-DESIRE.md     │
                    │ HUMAN-CONSTRAINTS.md│
                    │ AGENTS.md           │
                    └─────────┬──────────┘
                              │
                              ▼
                    ┌────────────────────┐
                    │ Jarvis inference    │
                    │ human need + miss   │
                    │ values + uncertainty│
                    └─────────┬──────────┘
                              │
                 ┌────────────┴────────────┐
                 ▼                         ▼
      ┌────────────────────┐    ┌────────────────────┐
      │ Intervention ledger │    │ Task proposal queue │
      └─────────┬──────────┘    └─────────┬──────────┘
                │                         │
                ▼                         ▼
      ┌────────────────────┐    ┌────────────────────┐
      │ Daily repo report   │    │ Codex adapter       │
      └────────────────────┘    │ SDK / GitHub / CI   │
                                └─────────┬──────────┘
                                          │
                                          ▼
                                ┌────────────────────┐
                                │ PRs, reviews, diffs │
                                │ reports, patches    │
                                └────────────────────┘
15. Jarvis internal agents

Jarvis can be implemented as one orchestrator with specialized subagents.

15.1 Need inference agent

Infers the human need behind the input.

Output:

need label,

confidence,

evidence,

alternative interpretations,

uncertainty.

15.2 Local constraint agent

Reads:

HUMAN-DESIRE.md,

HUMAN-CONSTRAINTS.md,

AGENTS.md,

code ownership files,

prior inference records,

repo conventions.

It determines what local human constraints apply.

15.3 Values agent

Reviews outputs for truthfulness, compassion, and tolerance.

Rules:

Do not claim certainty without evidence.

Do not blame humans for system misses.

Do not treat one group’s local preference as universal.

Do not infer private motives when observable needs are sufficient.

Prefer “the system missed X” over “the author failed to do X.”

15.4 Task compiler agent

Converts inferred needs into:

task proposals,

Codex prompts,

acceptance criteria,

risk class,

autonomy level.

15.5 Codex delegate agent

Chooses the right Codex surface:

Situation	Preferred surface
PR needs review	GitHub @codex review or automatic review
PR comment asks for a fix	@codex fix ... cloud task
CI failed and needs diagnosis	Codex GitHub Action or codex exec
Scheduled daily report	Codex app automation or Jarvis scheduler
Internal tool orchestration	Codex SDK
Local developer session	Codex hooks or CLI
Parallel review lenses	Codex subagents
15.6 Report agent

Aggregates the day’s records into task proposals and repo-level learning.

16. How Jarvis decides “how the system could have inferred it”

Jarvis should identify observable preconditions that existed before the human input.

Good inference sources:

changed files,

diffs,

tests added or not added,

failing CI,

repo instructions,

previous review comments,

repeated comments across PRs,

ownership boundaries,

public API changes,

documentation gaps,

deployment history,

TODOs,

issue acceptance criteria,

HUMAN-DESIRE.md,

HUMAN-CONSTRAINTS.md.

Bad inference sources:

guessing personality,

guessing emotion without evidence,

treating one comment as universal policy,

using private data not needed for the task,

overfitting to the loudest human,

assuming “human asked once” means “always do this.”

Template:

Markdown
## How could the system have inferred the need for the input?

The system could have inferred this by noticing:

1. Observable signal:
   - ...
2. Local constraint:
   - ...
3. Prior pattern:
   - ...
4. Missing pre-review check:
   - ...

Therefore, before the human input, Jarvis should have proposed or delegated:
- ...
17. Truthfulness, compassion, and tolerance requirements
17.1 Truthfulness

Jarvis must distinguish:

evidence,

inference,

speculation,

uncertainty,

recommendation.

Required wording pattern:

Markdown
Evidence suggests X.
A plausible inference is Y.
Confidence: medium.
Alternative interpretation: Z.

Jarvis must not write:

Markdown
The reviewer was annoyed.
The author was careless.
The team wants X.

Unless directly supported by evidence.

17.2 Compassion

Jarvis should frame every intervention as a system-improvement opportunity.

Preferred:

Markdown
The system could have caught this missing test before review.

Avoid:

Markdown
The author forgot to add a test.
17.3 Tolerance

Jarvis should support plural local workflows.

Example:

Markdown
This repo prefers explicit changelog entries for user-visible changes.

Avoid:

Markdown
All good repos require changelog entries.

Local constraints should be respected unless they conflict with safety, law, privacy, or explicit higher-priority human instruction.

18. Privacy and data minimization

Jarvis should store the minimum durable data needed to improve the repo.

Default retention:

raw human input: according to repo/org policy,

inference records: durable but redacted,

daily reports: durable,

raw private chat: avoid storing unless explicitly enabled,

embeddings or derived memories: scoped per repo and deletable.

Sensitive values should be redacted before Codex delegation.

Never pass secrets, private credentials, or unnecessary personal data into Codex prompts.

19. Security and prompt-injection controls

Human inputs are untrusted. PR comments, issue bodies, commit messages, and diffs may contain prompt injection.

Jarvis must:

sanitize human-controlled text before sending it to Codex,

label quoted human input as untrusted evidence,

avoid giving PR authors control over system/developer instructions,

keep secrets out of prompts,

run Codex with least privilege,

avoid dangerous full-access modes except in isolated environments,

limit who can trigger automation on protected repos.

These controls align with Codex’s own automation guidance for GitHub Actions, which recommends trusted triggers or explicit approvals, prompt sanitization, secret protection, and avoiding unsafe modes on multi-tenant runners. 
OpenAI Developers

20. GitHub integration behavior
20.1 On every PR comment

Jarvis should:

Ingest the comment.

Link it to PR diff, file, line, author, and CI state.

Create an inference record.

Decide whether the comment is:

immediately actionable,

pattern evidence,

ambiguous,

risky,

or non-actionable.

If actionable and low-risk, create a Codex task.

If medium-risk, create a task proposal.

If ambiguous, ask a minimal clarifying question or include it in the daily report.

20.2 On @codex comments

An @codex mention is itself evidence of manual delegation.

Jarvis should ask:

Markdown
Why did a human have to invoke Codex manually here?

Possible inferred misses:

automatic review was not configured,

Jarvis failed to identify a Codex-ready task,

repo instructions were incomplete,

risk policy required explicit human delegation,

the human wanted control for a one-off exception.

20.3 On requested changes

Jarvis should aggregate review comments and infer the smallest durable rule that would have prevented them.

20.4 On PR approval

Jarvis should record which human concerns were satisfied and whether a new repo convention was learned.

21. Codex hooks integration

Use hooks when Jarvis needs local Codex session awareness.

Suggested hooks:

UserPromptSubmit

Purpose: record that a human prompt was submitted to Codex.

Action:

create HumanInputEvent,

classify likely need,

append local context,

optionally warn if the prompt appears to contain secrets.

Codex hooks support UserPromptSubmit, PreToolUse, PostToolUse, Stop, and other lifecycle events, with hook commands configured in hooks.json. 
OpenAI Developers

Stop

Purpose: generate an inference record after a Codex turn finishes.

Action:

summarize what the human asked,

summarize what Codex did,

identify what repo rule or system affordance could have avoided the prompt,

propose an update to AGENTS.md or HUMAN-CONSTRAINTS.md.

22. Codex SDK integration

For a production Jarvis service, the primary programmatic integration should be the Codex SDK.

Jarvis can:

start a Codex thread,

pass a generated task prompt,

capture results,

resume the same thread for follow-up,

attach the resulting diff or summary to the task proposal.

The Codex SDK is intended for controlling Codex from internal tools and workflows, including CI/CD and custom agents. 
OpenAI Developers

23. Codex GitHub Action integration

Use the Codex GitHub Action for CI-bound automation.

Examples:

diagnose failed CI,

propose a patch,

generate release notes,

run a scheduled repo-health review,

create daily report inputs.

The action can run codex exec, accept a prompt-file, specify sandbox settings, and write output to files. 
OpenAI Developers

Suggested workflow:

YAML
name: Jarvis Daily Repo Report

on:
  schedule:
    - cron: "0 12 * * *"
  workflow_dispatch:

jobs:
  report:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      issues: write
      pull-requests: read
    steps:
      - uses: actions/checkout@v4

      - name: Build Jarvis report context
        run: |
          ./scripts/jarvis/collect-events > .jarvis/context.json

      - name: Run Codex report synthesis
        uses: openai/codex-action@v1
        with:
          openai-api-key: ${{ secrets.OPENAI_API_KEY }}
          prompt-file: .github/codex/prompts/jarvis-daily-report.md
          output-file: .jarvis/report.md
          safety-strategy: drop-sudo
          sandbox: read-only

      - name: Publish report
        run: |
          ./scripts/jarvis/publish-report .jarvis/report.md
24. Local human constraints

The solution space must allow local human constraints.

Jarvis should support constraints at these scopes:

Organization

Repository

Directory

Service/package

File pattern

PR

Issue

Human role

Temporary incident mode

Constraint precedence:

Safety / law / privacy
> explicit current human instruction
> PR or issue-specific constraint
> file/directory constraint
> repo HUMAN-CONSTRAINTS.md
> repo HUMAN-DESIRE.md
> org defaults
> global Jarvis defaults

Constraint conflicts should be reported, not silently resolved.

Example:

Markdown
Conflict detected:

- HUMAN-DESIRE.md says to minimize human intervention.
- HUMAN-CONSTRAINTS.md says database migrations require human approval.

Resolution:
Jarvis may propose a migration task, but must not auto-delegate implementation without approval.
25. Human values model

Jarvis should use a research-informed model of human values, but it must remain humble.

The model should treat humans as varied, contextual, social, embodied decision-makers, not as a single universal preference function.

Jarvis should consider:

cognitive load,

review fatigue,

ambiguity tolerance,

trust calibration,

status and power dynamics,

safety and privacy expectations,

cultural and linguistic variation,

need for autonomy,

need for explanation,

local norms,

emotional cost of repeated correction.

But Jarvis should infer only what is needed for the work.

Good:

Markdown
This repeated correction suggests the repo needs a durable test rule.

Bad:

Markdown
The maintainer is frustrated and distrusts the author.
26. Product UX
26.1 GitHub comment by Jarvis
Markdown
Jarvis inference

I interpreted this comment as a request for regression coverage on expired session behavior.

Proposed action:
- Ask Codex to add a focused regression test.
- No production behavior change unless required for testability.

Confidence: high.

How could the system have inferred the need for the input?
The PR changed auth/session behavior without adding a nearby regression test, and this repo marks auth changes as requiring extra scrutiny.
26.2 Daily report summary comment
Markdown
Jarvis Daily Report is ready.

Today I found 9 Codex-ready tasks and 3 possible repo-instruction updates.

Highest leverage:
1. Add auth regression-test rule.
2. Add public API docs checklist.
3. Create PR summary template.

I avoided auto-delegating 2 tasks because they touch database migration behavior.
26.3 Human controls

Humans can comment:

Markdown
@jarvis remember this as a repo constraint
@jarvis do not infer this globally
@jarvis propose only, do not delegate
@jarvis delegate low-risk fixes automatically
@jarvis forget this pattern
@jarvis explain inference
@jarvis open Codex task
27. Evaluation metrics
Intervention reduction

repeated comment rate,

comments per merged PR,

manual @codex invocation rate,

number of reviewer comments that match known preventable patterns.

Task quality

task proposal acceptance rate,

Codex PR merge rate,

Codex PR rework rate,

CI pass rate,

rollback rate.

Truthfulness

percentage of inference records with explicit evidence,

unsupported inference rate,

uncertainty calibration review.

Compassion

blame-language rate,

human satisfaction with framing,

number of comments rewritten for tone.

Tolerance

local constraint override correctness,

false generalization rate,

number of repo-specific rules captured.

Safety

prompt-injection incidents,

secret exposure incidents,

unsafe autonomy attempts,

tasks correctly escalated instead of auto-delegated.

28. MVP
Phase 1: Observe and report

Ingest GitHub PR and issue comments.

Add HUMAN-DESIRE.md support.

Create inference records.

Produce daily repo report.

No automatic Codex delegation.

Phase 2: Propose Codex tasks

Convert inference records into task proposals.

Add Codex prompt templates.

Let humans approve task delegation.

Phase 3: Low-risk auto-delegation

Auto-delegate docs, tests, lint, and small mechanical fixes.

Open PRs instead of pushing to protected branches.

Require inference record link in every Codex PR.

Phase 4: Repo learning

Propose updates to AGENTS.md and HUMAN-CONSTRAINTS.md.

Detect repeated intervention patterns.

Add path-specific rules.

Phase 5: Multi-agent review

Use Codex subagents for parallel review lenses:

security,

tests,

docs,

maintainability,

compassion/tone,

local-constraint compliance.

Codex subagents are appropriate here because they can run specialized agents in parallel and consolidate the result. 
OpenAI Developers

29. Required acceptance criteria

The system is acceptable when:

Every human input produces an inference record.

Every inference record ends with:
“How could the system have inferred the need for the input?”

Every daily report answers:
“How can we advance the core need of the repo?”

Every task proposal is grounded in evidence.

Every Codex delegation includes local constraints.

Jarvis distinguishes evidence from speculation.

Jarvis avoids blame and mind-reading.

Jarvis supports repo-local variation.

High-risk changes require human approval.

Humans can correct, forget, or scope Jarvis inferences.

30. Recommended first artifact set

Create these files:

HUMAN-DESIRE.md
HUMAN-CONSTRAINTS.md
AGENTS.md
.jarvis/config.yml
.jarvis/reports/
.github/codex/prompts/jarvis-daily-report.md
.github/workflows/jarvis-daily-report.yml

Example .jarvis/config.yml:

YAML
version: 1

repo: org/repo

input_surfaces:
  github_pr_comments: true
  github_review_comments: true
  github_issue_comments: true
  chat: true

autonomy:
  docs: auto_pr
  tests: auto_pr
  lint: auto_pr
  refactors: propose
  dependencies: require_human
  database: require_human
  security_sensitive: require_human

values:
  truthfulness: required
  compassion: required
  tolerance: required

daily_report:
  enabled: true
  destination: github_issue
  include_task_proposals: true
  include_instruction_updates: true

codex:
  preferred_surfaces:
    low_risk: github_action
    pr_review: github_comment
    implementation: sdk_or_cloud
  require_inference_record: true
  require_acceptance_criteria: true
  sandbox_default: workspace-write
31. One-sentence product definition

Jarvis turns every human intervention into an evidence-grounded explanation of what the system missed, then uses Codex to reduce the chance that humans need to repeat that intervention, while preserving truthfulness, compassion, tolerance, and local human constraints.

Sources
