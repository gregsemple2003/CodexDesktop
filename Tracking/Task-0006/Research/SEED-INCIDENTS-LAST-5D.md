# Task 0006 Seed Incidents From The Last Five Days

## Purpose

This note names the five starter incidents recommended for `Task-0006`.

The set is intentionally mixed. It is meant to prove that incident capture can preserve usability misses, UI-language misses, and broader misunderstandings instead of only one class of problem.

This seed note predates the newer `why_chains` contract.

Read the detailed bullets below as candidate cues for what each incident is about, not as the current schema shape. When these candidates are promoted into real incidents, they should be rewritten into:

- concrete `expected_state`
- concrete `actual_state`
- `human_intervention`
- one or more ordered `why_chains`

These are candidate seeds, not automatically qualified incidents. Under the revised contract, each record must still prove:

- a grounded pre-correction state existed before the intervention
- the human corrected that outcome directly
- the correction delta is preserved

Any candidate that cannot clear that gate should be replaced during backfill rather than weakened into a generic bug note.

## Incident 0001 Home Upload State Story Mismatch

- source date: April 4, 2026
- coverage: usability, truthful status communication
- source evidence:
  - [Split Crystallize bug tasks transcript](/c:/Users/gregs/.codex/sessions/2026/04/04/rollout-2026-04-04T00-44-55-019d56ce-f0d4-7360-af4c-cd1dcca105ec.jsonl#L127)
  - [Task-0022](/c:/Agent/Crystallize/Tracking/Task-0022/TASK.md)

Expected incident content:

- `title`: Home says clips are safe while still rendering unfinished progress
- `primary_family`: `usability_status_truth`
- `surface` (`observed`): the card says `Safe on server` while also showing incomplete progress treatment.
  Falsifiers: the same state does not actually show both elements together; the progress element is explicitly labeled as a different metric.
- `human_goal` (`inferred`): the card should let a human feel "upload is complete and safe," but instead it suggests unfinished upload work and erodes trust.
  Falsifiers: the copy or structure clearly explains the bar as local-retention or cleanup work rather than upload progress.
- `task_goal` (`inferred`): a completed safe state should use completed-state grammar and tell one calm authoritative story.
  Falsifiers: the approved contract or design anchor explicitly allows safe state plus partially filled upload-style progress.
- `producer_goal` (`hypothesized`): the producer likely optimized for preserving visibility into retained local work or background sync detail even when the state was already safe.
  Falsifiers: transcript or implementation evidence shows the producer understood the contradiction and the mismatch came from a temporary placeholder or separate bug.
- `workflow_goal` (`hypothesized`): no review step forced "one state story per card," so a semantically mixed composition survived because it looked informative.
  Falsifiers: the task, prompt, or review already encoded that exact rule and the miss was instead a one-off execution defect.
- `primitive` (`hypothesized`): completed safety states must not reuse in-progress visual grammar unless the secondary metric is explicitly named and clearly subordinate.
  Falsifiers: a good counterexample where a safe state legitimately carries a clearly labeled secondary progress metric without confusing the main story.
- `likely_durable_fix_layer`: Android state-mapping and copy treatment on the Home upload card

## Incident 0002 People Surface Behaves Like A Setup Detour

- source date: April 4, 2026
- coverage: information architecture, naming, CTA framing
- source evidence:
  - [Split Crystallize bug tasks transcript](/c:/Users/gregs/.codex/sessions/2026/04/04/rollout-2026-04-04T00-44-55-019d56ce-f0d4-7360-af4c-cd1dcca105ec.jsonl#L127)
  - [Task-0025](/c:/Agent/Crystallize/Tracking/Task-0025/TASK.md)

Expected incident content:

- `title`: People route still reads like onboarding instead of direct person management
- `primary_family`: `ia_and_cta_mismatch`
- `surface` (`observed`): the surface uses labels and actions like `Trusted cast`, `Refresh me and dad setup`, and `Open onboarding`.
  Falsifiers: the visible route title and primary CTA are actually direct people-management actions.
- `human_goal` (`inferred`): the route should help a human manage recurring people directly, but it instead reads like a setup detour.
  Falsifiers: the list structure and primary action clearly communicate direct people management despite the legacy copy.
- `task_goal` (`inferred`): the route should behave like `People`, support direct actions such as `Add person`, and keep starter defaults from hardening into the only model.
  Falsifiers: the approved route contract is explicitly onboarding-first rather than management-first.
- `producer_goal` (`hypothesized`): the producer likely overfit to the existing enrollment scaffolding and the `Me` / `Dad` starter narrative, treating bootstrap structure as the route's enduring IA.
  Falsifiers: transcript evidence shows the producer intentionally chose a management-first model and the mismatch was instead limited to unfinished copy replacement.
- `workflow_goal` (`hypothesized`): the task producer or review flow failed to turn the bug report into a sharper route contract, leaving setup scaffolding unchallenged as the default frame.
  Falsifiers: the task artifacts already captured the management-first rule clearly and the miss survived despite explicit consensus.
- `primitive` (`hypothesized`): a human-facing route should be named and actioned around the ongoing human job it performs, not around the bootstrap path that originally created its contents.
  Falsifiers: true onboarding or recovery routes where the setup path is in fact the ongoing primary job.
- `likely_durable_fix_layer`: People surface information architecture, route naming, and CTA design

## Incident 0003 Jobs Tab Row Taxonomy Is Too Operator-Heavy

- source date: April 4, 2026
- coverage: UI semantics, terminology, scan-language mismatch
- source evidence:
  - [Review Jobs-tab mockup transcript](/c:/Users/gregs/.codex/sessions/2026/04/04/rollout-2026-04-04T11-34-46-019d5921-e4e1-7930-b39d-8c582fc9fd22.jsonl#L66)
  - [Task-0004](/c:/Agent/CodexDashboard/Tracking/Task-0004/TASK.md)

Expected incident content:

- `title`: Jobs surface exposes raw operator jargon where the default UI should be human-readable
- `primary_family`: `ui_language_mismatch`
- `surface` (`observed`): labels such as `WINDOWS_BOOT`, `CRON_DAEMON`, and `State (D/O)` appear in the main row taxonomy.
  Falsifiers: those labels exist only behind disclosure or diagnostic reveals rather than on the default scan path.
- `human_goal` (`inferred`): the screen should let a human understand job state quickly, but it instead asks the reader to parse internal mechanism names and acronyms first.
  Falsifiers: the same rows already provide plain-language equivalents with the raw terms clearly demoted.
- `task_goal` (`inferred`): the default Jobs surface should stay human-readable, with raw plumbing pushed behind details and labels like `Desired / Observed` used in the primary reading path.
  Falsifiers: the approved surface contract explicitly targets trained operators and intentionally uses raw mechanism taxonomy by default.
- `producer_goal` (`hypothesized`): the producer likely optimized for accurate exposure of internal runtime truth and assumed terminology fidelity was safer than translation.
  Falsifiers: transcript evidence shows the producer was already aiming for plain language and the remaining jargon was simply unfinished replacement work.
- `workflow_goal` (`hypothesized`): interface review or task framing did not force a hard distinction between default human-facing semantics and deeper diagnostic truth.
  Falsifiers: the review contract already encoded that split explicitly and the miss was only an implementation follow-through failure.
- `primitive` (`hypothesized`): default human-facing surfaces should name the human-meaningful state first; raw runtime taxonomy belongs behind disclosure unless the product explicitly targets trained operators.
  Falsifiers: surfaces whose explicit job is low-level operational debugging for trained operators.
- `likely_durable_fix_layer`: Jobs surface copy rules and interface disclosure policy

## Incident 0004 Producer Left The Resolution Too Loose

- source date: April 4, 2026
- coverage: general misunderstanding, orchestration, task-definition quality
- source evidence:
  - [Clarify usability task rules transcript](/c:/Users/gregs/.codex/sessions/2026/04/04/rollout-2026-04-04T01-52-00-019d570c-5a4e-7280-a118-6e17e120f60c.jsonl#L163)

Expected incident content:

- `title`: producer stopped at honest problem capture and failed to drive the likely agreed resolution
- `primary_family`: `orchestration_consensus_miss`
- `surface` (`observed`): the task wording left core human-facing design choices unresolved in the main task body.
  Falsifiers: the task already names the expected resolution shape clearly and limits uncertainty to a short bounded section.
- `human_goal` (`inferred`): the task should be reviewable by a human without reopening the core design problem, but instead it leaves the next agent to rediscover the intended outcome.
  Falsifiers: another agent could execute the task without reopening the core design question and still land a predictable human-facing result.
- `task_goal` (`inferred`): a producer-owned usability task should already reflect likely design-owner consensus and leave only truly residual uncertainty open.
  Falsifiers: the producer's contract for this class of task explicitly allows open-ended design resolution by downstream implementers.
- `producer_goal` (`hypothesized`): the producer likely believed "capture the problem honestly" was enough and treated consensus-building as downstream work rather than producer responsibility.
  Falsifiers: transcript evidence shows the producer did perform that synthesis and the unresolved wording came from a different constraint.
- `workflow_goal` (`hypothesized`): the task-production workflow lacked a hard rule that usability tasks must synthesize interface and general-design input into a reviewable expected-resolution section.
  Falsifiers: the shared producer workflow already encoded that rule clearly before the miss.
- `primitive` (`hypothesized`): for narrow human-facing tasks, the producer must deliver a falsifiable expected resolution, not just a sincere problem statement.
  Falsifiers: exploratory research tasks whose explicit job is to leave the design question open.
- `likely_durable_fix_layer`: task-producer workflow, prompt rules, and task artifact expectations

## Incident 0005 Proof-View Experiment Drifted Away From The Real Defect

- source date: April 4, 2026
- coverage: general misunderstanding, proxy-proof drift, debugging discipline
- source evidence:
  - [Lead Task-0002 workflow transcript](/c:/Users/gregs/.codex/sessions/2026/04/04/rollout-2026-04-04T11-18-43-019d5913-3060-70f1-a7d7-44a771efaef8.jsonl#L750)

Expected incident content:

- `title`: proof-only camera helper work started expanding around the wrong seam
- `primary_family`: `proxy_truth_misread`
- `surface` (`observed`): the workflow was moving toward richer proof-view automation instead of staying bounded to one discriminating check.
  Falsifiers: the proof-view work was already explicitly capped and directly tied to a go-or-no-go defect decision.
- `human_goal` (`inferred`): the work should answer the human-facing question of whether the pawn is actually visible, but proxy evidence risked delaying that answer.
  Falsifiers: the helper work was the only realistic path to answering the visibility question and did not expand beyond that.
- `task_goal` (`inferred`): the task should either prove the real gameplay view works or quickly justify promotion to the actual render or visibility defect path.
  Falsifiers: the task contract explicitly prioritized generalized automation helpers over a bounded discriminating proof.
- `producer_goal` (`hypothesized`): the producer likely optimized for better proof tooling because tool quality felt like forward progress and local uncertainty made the defect seam seem ambiguous.
  Falsifiers: transcript evidence shows the producer already understood the real seam and was only finishing one bounded helper check.
- `workflow_goal` (`hypothesized`): the debugging workflow lacked a strong enough pass-boundary rule for when proxy proof must stop and a real defect path must begin.
  Falsifiers: the workflow already encoded that decision boundary and the drift happened despite explicit enforcement.
- `primitive` (`hypothesized`): proof helpers should stay bounded to the smallest discriminating check; when they fail to answer the human-facing question, escalate to the real defect path instead of enriching the proxy.
  Falsifiers: cases where broader proof tooling is itself the primary deliverable rather than a temporary diagnostic seam.
- `likely_durable_fix_layer`: debugging workflow rules and pass-boundary discipline

## Coverage Note

This seed set intentionally spans:

- one direct usability-truth incident
- one information-architecture incident
- one explicit UI-language incident
- two broader misunderstanding incidents at the orchestration and proof-strategy layers

It also spans different upstream goal-failure families:

- state-story mismatch
- route-model mismatch
- default-surface semantics mismatch
- producer-consensus failure
- proxy-proof workflow drift

That mix is the point. `Task-0006` should learn from all five kinds of misses, not just the most photogenic UI ones.
