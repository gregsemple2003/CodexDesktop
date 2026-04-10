## User

I have a set of promts for an agent orchestration framework.  I want to scale it out.  I have an android app that I can dispatch commands to and codex markdown system to interpret tasks.   What's the best tech stack to scale that up to a company of people.

I want to say "log a bug around x and get interface designer and general designer to come up with a plan then send it to me for review" then when I approve just start working

## Assistant

Thought for 27m 3s

The stack I’d bet on for this is:

Temporal as the workflow control plane.

OpenAI Responses API + Agents SDK as the agent runtime.

MCP as the main tool/integration layer.

Codex AGENTS.md + Skills + explicit subagents for repo-local behavior and specialized roles.

Linear/Jira + Figma + GitHub + Slack + your Android app as the systems people actually work in. Temporal gives you durable workflows, recoverability, signals/queries for human approvals, and long-running execution; OpenAI gives you handoffs, tools, structured outputs, background runs, and webhooks; MCP standardizes tool/context connections and approval flows; Codex gives you layered instructions, reusable skills, and bounded parallel subagents. 
OpenAI Developers
+13
Temporal Docs
+13
Temporal Docs
+13

The important architectural move is to separate business workflow from agent behavior. Your workflow layer should own state, approvals, retries, deadlines, and audit. Your agents should only do bounded work: triage, plan, draft, review, implement. That is what makes the system trustworthy for a company instead of “a powerful prompt that sometimes works.” Temporal’s docs are especially strong here because approvals can pause for hours or days without burning compute, and resume cleanly when a signal arrives. 
Temporal Docs
+2
Learn Temporal
+2

A concrete build I’d choose:

Backend language: TypeScript/Node for the command API, workflow workers, approval service, and most integrations.

Control plane: Temporal, managed early rather than self-hosted.

Agent layer: OpenAI Responses API + Agents SDK.

Tool layer: one MCP server per external system or domain.

Repo behavior: AGENTS.md for durable rules, Skills for reusable workflows, explicit subagents for parallel bounded jobs.

State: one relational DB plus object storage for artifacts, screenshots, plans, diffs, and run outputs.

Human surfaces: keep the Android app, but add a browser-based inbox/review UI and Slack approval cards so the whole company can use it. This matches the way the vendor tools expose APIs and interactive workflows today: Linear exposes GraphQL and webhooks, Figma has webhook APIs, GitHub gives branch protections/reviews/CODEOWNERS, and Slack supports interactive app flows. 
Slack Developer Docs
+4
Linear
+4
Figma Developer Docs
+4

For your specific example, I would model the request as a typed workflow, not as free-form prompt chaining:

request
  -> parse intent into schema
  -> create bug
  -> run interface_designer + general_designer in parallel
  -> merge plans into review packet
  -> wait for your approval
  -> on approve: create branch/PR
  -> implement
  -> run checks/review
  -> wait for merge/deploy approval

The first step should use Structured Outputs so “log a bug around x and get interface designer and general designer to come up with a plan then send it to me for review” becomes a typed task envelope instead of prose you later regex apart. OpenAI’s docs explicitly position structured outputs for schema-shaped responses and function calling for external actions. 
OpenAI Developers
+1

A good task envelope looks like this:

JSON
{
  "workflow_type": "bug_plan_then_implement",
  "bug_summary": "...",
  "product_area": "...",
  "requested_roles": ["interface_designer", "general_designer"],
  "artifacts_required": ["bug_ticket", "ui_plan", "design_plan", "review_packet"],
  "approval_gates": ["before_code_changes", "before_merge", "before_deploy"],
  "risk_level": "medium"
}

Then the workflow does the rest deterministically:

Create issue in Linear/Jira.

Spawn two child jobs in parallel.

Each child agent writes a structured plan artifact.

An aggregator agent merges them into one review packet.

Your Android app, Slack, or web inbox gets an approval card.

Your approval sends a Temporal Signal.

The workflow resumes and starts implementation. Temporal’s signal/query model is almost exactly the shape you need for this. 
Temporal Docs
+1

How I’d map the roles:

Triage agent: can read logs/specs and create a bug ticket.

Interface designer: can read app context and propose UI/flow changes.

General designer: can propose broader product/design direction.

Planner/aggregator: merges outputs into a review packet.

Implementer: can open a branch and PR, but not merge.

Release agent: can prepare deployment steps, but only runs after explicit approval.

That permission split matters. OpenAI’s MCP guidance recommends approvals for sensitive tool calls, logging the data shared with MCP servers, and only connecting to trusted servers because prompt injection and unexpected tool behavior are real risks. 
OpenAI Developers
+1

For repo safety, make GitHub—not the model—the final enforcement layer. Protected branches can require PR reviews, status checks, conversation resolution, merge queues, and deployment success. CODEOWNERS lets you route the right reviewers automatically and optionally require code-owner approval before merge. That is exactly how you keep “approve and start working” from becoming “agent silently merged to main.” 
GitHub Docs
+1

For your existing markdown setup, keep it, but harden it:

Put durable repo/team rules in AGENTS.md.

Put repeatable workflows in Skills.

Use explicit subagents only for bounded jobs like triage, design, test review, or implementation slices.

Codex’s docs are pretty aligned with this: AGENTS.md is layered and loaded before work starts, Skills are reusable workflow bundles, and subagents are intended for explicit, bounded parallel work rather than vague autonomous swarms. Codex’s own best-practices docs also warn against overloading prompts with durable rules and suggest moving recurring logic into AGENTS.md or a skill. 
OpenAI Developers
+3
OpenAI Developers
+3
OpenAI Developers
+3

One practical recommendation: don’t let the model decide the orchestration policy. Let the model decide content inside a step, but let code decide:

which roles run,

which tools they may call,

when approval is required,

how many retries are allowed,

which artifacts must exist before advancing.

That gives you repeatability, cost control, and auditability.

The main alternative stack is LangGraph inside, Temporal outside. LangGraph is good at stateful multi-agent flows, human-in-the-loop, memory, and observability. I would use it for an individual specialist workflow or for experimentation, but I would still keep Temporal as the outer company-wide process engine because the hard enterprise problems are approvals, recovery, long waits, and workflow history. 
Learn Temporal
+3
LangChain Docs
+3
LangChain Docs
+3

Finally, treat prompts/skills like code:

version them,

assign owners,

run evals,

log traces,

add regression suites before broad rollout.

OpenAI’s eval guidance is explicit that agent quality needs reproducible evaluations and that eval-driven development is the way to keep systems stable as prompts and models change. 
OpenAI Developers
+2
OpenAI Developers
+2

So the short recommendation is:

Use Temporal for orchestration, OpenAI Agents/Responses for agent work, MCP for integrations, Codex AGENTS.md/Skills for reusable behavior, GitHub protections for enforcement, and keep your Android app as one of the approval/command clients—not the brain.

The next useful artifact would be a concrete workflow spec for bug -> parallel design plans -> review -> approval -> implementation, including the JSON schema, role permissions, and MCP tool contracts.

Sources

## User

Yeah i mean I've basically got that going except github pr review but I hear you on temporal.  

But several bumps.  It doesn't understand interface design or the way humans do things.  Even with a mockup it doesn't understand what structure is important to human visual processing. 

It's s real bottleneck.  I'd prefer more leverage for my time.  Something like a Jarvis from iron man that could see these problems and correct them for me.  Or I could just say Jarvis deal with this.  Thoughts?  I'm using codex 5.4 right now.

## Assistant

Thought for 17m 33s

Yeah — I think you’ve already solved most of the orchestration problem. The bottleneck now is taste transfer. OpenAI’s March 2026 frontend guide basically says the same thing in different words: GPT‑5.4 got better at UI work and image use, but when prompts are underspecified it still drifts toward generic structure and weak visual hierarchy. Codex docs still recommend gpt-5.4 for most Codex tasks, and gpt-5.4 is also the mainline model trained for computer use. 
OpenAI Developers
+2
OpenAI Developers
+2

So I would not chase a single magic “Jarvis” prompt. I’d build Jarvis as a loop:

Watcher sees the UI as it really exists.

Critic judges it using your taste and human-perception rules.

Planner proposes fixes.

Builder changes code.

Verifier checks the rendered result before it comes back to you.

Codex already has the right primitives for this shape: layered AGENTS.md, reusable Skills, and explicit subagents. That means you can turn “interface judgment” into a first-class role instead of hoping the coding agent develops taste on demand. 
OpenAI Developers
+2
OpenAI Developers
+2

The key change is this: stop feeding it mostly code and mockups. Feed it perception. OpenAI’s current guidance for GPT‑5.4 frontends is to define design constraints up front, attach screenshots or mood boards as visual guardrails, and use Playwright so the model can inspect rendered pages, test multiple viewports, navigate flows, and visually verify its own work. That is much closer to the “Jarvis can see the problem” behavior you want. For Android, I’d mirror that with screenshots, a short screen recording, the view/semantics tree, and state context for empty/loading/error/success flows. 
OpenAI Developers
+2
OpenAI Developers
+2

The second big move is to externalize your visual intuition. OpenAI’s own sample frontend skill encodes things like restrained composition, image-led hierarchy, few colors, minimal chrome, “cards only when the card is the interaction,” and a checklist like “is there one strong visual anchor?” and “can the page be understood by scanning headlines only?” That’s the right pattern: write your taste down as a rubric, not as vibes in your head. 
OpenAI Developers

A design critic should return something like this on every UI task:

JSON
{
  "first_notice": "What the eye lands on first",
  "read_order": ["1st", "2nd", "3rd"],
  "primary_action_clarity": 1,
  "hierarchy_score": 1,
  "density_score": 1,
  "trust_score": 1,
  "attention_leaks": [
    {"region": "top-right", "reason": "competes with primary CTA"}
  ],
  "structural_problems": [
    "status and action are visually merged",
    "secondary content reads before primary task"
  ],
  "repairs": [
    "reduce chroma on secondary actions",
    "merge panels 2 and 3",
    "increase whitespace around primary task area"
  ]
}

That forces the model to talk in terms of scan path, hierarchy, distraction, and intent, which is much closer to how humans actually evaluate screens.

Then make it a closed loop. “Jarvis, deal with this” should not mean “go freestyle.” It should mean:

capture current UI -> score with design critic -> propose 2 repair plans
-> implement best plan -> re-render -> score again
-> if score improves past threshold, send before/after for review
-> if not, escalate with a concise diagnosis

That gives you leverage because you stop reviewing raw work and start reviewing already-filtered improvements.

I’d also separate the models by job. Keep gpt-5.4 / Codex as the orchestrator and implementer, because that’s what the current Codex docs optimize for. But for the “taste model,” use a sidecar you can tune. OpenAI currently exposes agent evals, vision fine-tuning, supervised fine-tuning, and distillation; the current docs I checked list gpt-4o-2024-08-06 for vision fine-tuning and gpt-4.1 variants for SFT, and they explicitly recommend getting evals in place first, then distilling a stronger setup into something cheaper. In practice, that means your custom visual critic may sit beside gpt-5.4 rather than replacing it. 
OpenAI Developers
+3
OpenAI Developers
+3
OpenAI Developers
+3

That sidecar model should learn your approvals. Every time you say “this hierarchy is wrong,” “this feels too dashboard-y,” “this doesn’t establish trust,” or “the structural read is backwards,” store the before/after and your note. After enough examples, you’re no longer prompting a generic model — you’re training a private design judge that approximates your taste.

When ambiguity is high, I would also bounce through Figma instead of trying to do all design negotiation in prose. OpenAI’s February 2026 Codex/Figma post says the Figma MCP server can bring running UI into Figma and back into code, which is exactly the right surface for design iteration when a screenshot critique isn’t enough. 
OpenAI Developers

My blunt take: you can get much closer to Jarvis today, but not as one autonomous genius agent. The workable version is a perception system + personal taste critic + fast builder + hard verification loop. That can give you a lot more leverage than you have now. Full autonomous aesthetic judgment still isn’t something I’d trust end-to-end.

The first thing I’d build next is a design-critic skill plus a 50-screen eval set of your own approvals/rejections, and I’d make every UI change pass that gate before implementation.

Sources

## User

Yeah i hear you ok not having a Jarvis out of the box.  I'm just trying to accelerate the process.   I have a short library of my taste preferences but that's not the only problem.  It's just not growing the way humans process visual information.  A button can be cut off and to them that passes regression.  I tried grounding it in design wisdom so I could distill human needs by looking at apple but there's still a lot of basic app design they are missing.  Also I want ideally something else to go off for a couple of hours and learn all they can abl9ut humans and how we are biased to see the world in a way that colors our perceptions 

The other issue is that basic elements of human design are missing.  When I said my android voice app is to record memories of me and my dad they designed an app that was literally forcing the me and dad profiles during first time user experience.  Something that doesn't make sense for me since I dont want to spend ui time on that and for other people would be insulting if they dont have a dad.  So I'm just finding it severely disconnected from the human world and so doing point solutions and curating taste preferences seems like the opposite of giving me leverage.

I can hit my openai subscription cap for codex.  So getting more leverage means parallelization and them just getting what I mean.  I can task stuff and dispatch workers but if they keep acting like Disney minions it's going to be a serious burn on my time.
2/2

## Assistant

Thought for 19m 14s

You’re not crazy. What you’re running into is not mainly “the taste prompt isn’t long enough.” OpenAI’s current GPT‑5.4 frontend guidance still leans hard on screenshots or mood boards, Playwright-style visual checks, and explicit verification loops, and their evals docs frame measurement as essential for reliability rather than optional polish. A strong builder without a strong judge still ships broken or tone-deaf UI. 
OpenAI Developers
+3
OpenAI Developers
+3
OpenAI Developers
+3

I think you actually have three different problems, and they need different fixes.

First, perceptual failures: clipped buttons, bad spacing, wrong scan path, primary CTA losing the page.
Second, human-world failures: turning “memories of me and my dad” into a required two-profile onboarding flow.
Third, taste failures: even when the structure works, it still feels generic or off.

Your dad example is the big clue. That is not a taste bug. It is not even mainly a visual bug. It is a world-model bug: the model took an emotionally salient story and promoted it into the app’s schema. In other words, it heard an example and treated it as ontology. That is why “study Apple” doesn’t fix it. Apple can teach surface discipline, but it does not teach the model when a user story should stay a story instead of becoming required product structure.

So I would stop trying to build “one smarter Jarvis prompt,” and I would stop adding more parallel builders. The leverage move is parallel critics.

Keep Codex / GPT‑5.4 as the builder, but put three gates in front of implementation:

A visual auditor that only cares about what a human can literally see.

A human-world critic that catches assumption bugs, dignity problems, onboarding burden, and “example vs schema” mistakes.

A taste critic that applies your preferences once the first two are already clean.

That fits the current product direction well: Codex subagents are meant for specialized parallel work, GPT‑5.4 has screenshot-based computer use, and the current Codex/frontend docs explicitly recommend screenshots or Figma references plus visual checks when building UI. 
OpenAI Developers
+4
OpenAI Developers
+4
OpenAI Developers
+4

The missing piece for you is probably the human-world critic. I’d make it produce something like this before any UI is allowed to exist:

JSON
{
  "job_to_be_done": "Preserve and record memories quickly",
  "first_value_moment": "User can record a memory within 10 seconds",
  "example_entities": ["me", "dad"],
  "persistent_domain_entities": [],
  "forbidden_assumptions": [
    "user has a father",
    "user wants relationship profiles before first use",
    "personal identity setup is needed to get value"
  ],
  "setup_to_delay": [
    "profile creation",
    "relationship metadata",
    "sharing preferences"
  ],
  "dignity_risks": [
    "forcing family-role assumptions in onboarding"
  ],
  "recommended_onboarding": "Open straight into recording, explain optional organization later"
}

That turns “understand humans” into something operational. The agent now has to answer: what is the first valuable action, what assumptions are being smuggled in, and what should stay optional until later.

For the visual auditor, do not let it read code first. Let it read screenshots across states and sizes. Its whole job is brutal, literal perception: is anything clipped, overlapped, outside the safe area, visually competing with the main task, or requiring too much scanning? OpenAI’s current guidance is aligned with that: visual references, rendered inspection, and visual checks are the path for frontend quality, not text-only prompting. 
OpenAI Developers
+3
OpenAI Developers
+3
OpenAI Developers
+3

On your “send something off for a couple of hours to learn all it can about humans” idea: yes, but I would use that as a research-and-evals worker, not as an unconstrained self-improvement agent. OpenAI’s Responses API has background mode for long-running work, and OpenAI positions Responses as the main direction for agent builds. So let a background worker spend time producing artifacts like distilled human-factors notes, counterexamples, new eval cases, and updated critic prompts. Do not let that same worker directly redesign the app from vague conclusions. 
OpenAI Developers
+2
OpenAI Developers
+2

The other big leverage move is to get off the interactive-cap treadmill for the repetitive parts. OpenAI’s docs recommend GPT‑5.4 as the starting model, with smaller 5.4 variants when you want lower cost and latency, and their optimization docs push distillation once you know what “good” looks like. Codex is also meant to be used not just interactively, but in SDK/CI workflows. So the architecture I’d use is: big model for implementation and hard disagreements, smaller or distilled models for the critics that run constantly. 
OpenAI Developers
+3
OpenAI Developers
+3
OpenAI Developers
+3

I’d tune those critics differently:

Visual auditor: vision fine-tuning or SFT on screenshots plus correct issue lists, because this is an image-understanding problem. 
OpenAI Developers
+1

Human-world critic: SFT on your real review comments and corrected plans, because you have expert demonstrations of what should have been inferred. The SFT guide says improvements often show up with 50–100 strong examples, so you do not need a giant dataset to start. 
OpenAI Developers

Taste critic: DPO on pairwise choices and written rationales, because this is subjective preference learning. DPO is text-only today, so use it on ranked plan alternatives and critiques rather than raw screenshots. 
OpenAI Developers

Hard pass/fail UI rules: evals and possibly RFT when the task is clear and verifiable, since OpenAI’s RFT guidance is strongest for well-defined tasks scored by graders. 
OpenAI Developers
+3
OpenAI Developers
+3
OpenAI Developers
+3

So my blunt take is: the fix is not “teach Codex humanity in the abstract.” The fix is teach your system what counts as a human mistake, then make every builder answer to that. Apple-inspired taste notes help, but they are the third layer, not the first. First comes “is the UI physically legible,” second comes “does this respect how people actually live,” and only then “does it feel like you.”

If I were changing one thing next, I’d add a human-world critic and make every UI task pass three checks before code lands: value before setup, example vs schema, and dignity / assumption risk. One builder plus three critics will feel a lot more like Jarvis than ten more Disney minions.

Sources

## User

Okay lets go with your idea.  Lets ground in actual feedback I've given. There's a lot of me trying to teach it underpinning principles that humans with our embodied meat suits max8mize.  

I guess the question is how to maximize the use of the data.  Lets say we use your schema. 

How can we maximize information extracted?  I think one thing that could help is pursuing root cause.  From a disagreement about expected state vs actual state pursue root cause.

Why is button font sizing and alignment important?  Because we use relative size cues to tell us where to start.
Why?  Because we want to extract the most important information first.  We can't just process a stream of text efficiently. 
Why?  Because our brains are specialized for visual information to help us avoid predators and find prey.  Movement, color, relative size are all critical to our evolution.
Why?  To maximize survival.

The reason pursuing root cause is important is because instruction following gets weaker the more local rules we introduce.  So boiling down the primitives is important.

Tasks should be written with concrete outcomes that can be falsified.
Why?  So the human can determine if that's the best outcomes and to force multiple disciplines to interact to distill a consensus.
Why?  So we can get at the most human fitting solution.

## Assistant

Thought for 23m 14s

Yes — pursue root cause. But I’d make one change to how you stop the chain.

Do not stop at the deepest true story.
Stop at the lowest reusable causal primitive.

“Humans evolved to maximize survival” may be true, but it is too far away from the screen to steer design. “Primary actions must be visually complete and dominant at first glance” is close enough to predict real UI choices across many screens. That is the level that gives you leverage.

A good primitive does four things at once: it explains several past failures, predicts new ones, has a clear boundary where it does not apply, and can be falsified from artifacts like screenshots, plans, and flows. That is the unit I’d optimize for.

The current OpenAI guidance lines up with that direction: agent quality should be measured with reproducible evals, workflow-level failures should be traced and graded, memory should be distilled and consolidated rather than dumped wholesale into prompts, SFT should start from a small set of high-quality demonstrations and be evaluated, DPO is for subjective pairwise preferences, and stronger large-model behavior can be distilled down into cheaper workers once your eval criteria are stable. Current GPT-5.4 mini/nano docs also position those models for high-volume subagents and extraction/ranking-style work, which is useful when you want critics running all the time instead of burning frontier-model budget on every pass. 
OpenAI Developers
+6
OpenAI Developers
+6
OpenAI Developers
+6

Here’s the schema I’d use for every disagreement:

JSON
{
  "incident": {
    "task": "",
    "expected_state": "",
    "actual_state": "",
    "evidence": ["screenshot", "plan", "diff", "review note"]
  },
  "human_cost": [
    "scan confusion",
    "setup before value",
    "dignity risk",
    "trust erosion",
    "action ambiguity"
  ],
  "why_chain": [
    {"level": "surface", "text": ""},
    {"level": "mechanism", "text": ""},
    {"level": "functional", "text": ""},
    {"level": "primitive", "text": ""}
  ],
  "operational_primitive": "",
  "deep_rationale": "",
  "scope": {
    "applies_when": [],
    "does_not_apply_when": []
  },
  "falsifiers": [],
  "preferred_fix": "",
  "counterexample": "",
  "training_artifacts": {
    "eval_case": {},
    "sft_demo": {},
    "dpo_pair": {},
    "memory_note": ""
  }
}

The important split there is between operational_primitive and deep_rationale.

The deep rationale is for you.
The operational primitive is for the agent.

That keeps the prompt clean. You still preserve the human philosophy, but you don’t make the builder drag the whole evolutionary ladder around every time it places a button.

I’d also add two questions to every “why” chain, because “why” alone is not enough:

When would this principle be wrong?

How would we falsify compliance?

That turns a principle into something trainable.

So your extraction loop becomes:

State the disagreement in observable terms: expected vs actual.

Name the human cost: what got worse for the user.

Climb the why ladder until you hit a primitive that predicts more than this one incident.

Add anti-scope so the primitive doesn’t become a universal bad rule.

Generate falsifiers so the critic can actually check it.

Explode the incident into training artifacts.

That last step is where you get the leverage. One review comment should not become one note. It should become four assets.

When you have a correction in your own words, it becomes an SFT demonstration. When you choose between two alternatives or say “not this, this,” it becomes a DPO pair. When you can define pass/fail, it becomes an eval case. When the same idea repeats across incidents, it becomes a consolidated primitive note. OpenAI’s SFT guide explicitly recommends starting with about 50 strong demonstrations and evaluating, noting that improvements often show up in the 50–100 example range; the DPO guide is specifically for preferred vs non-preferred outputs; and the memory cookbook recommends two-phase note-taking then consolidation, with deduplication and forgetting as first-class behavior. 
OpenAI Developers
+2
OpenAI Developers
+2

Your two examples show exactly how this works.

For the clipped or badly aligned button, I would not store “font sizing matters” as the top-level lesson. I’d store something like:

JSON
{
  "expected_state": "User immediately recognizes the primary action and can commit to it without uncertainty.",
  "actual_state": "Primary CTA is clipped / visually unstable / loses hierarchy.",
  "human_cost": ["action ambiguity", "trust erosion", "scan disruption"],
  "operational_primitive": "Primary actions must be visually complete and dominant at first glance across supported device sizes.",
  "scope": {
    "applies_when": ["first-run actions", "critical progression actions", "confirmation flows"],
    "does_not_apply_when": ["non-action decorative labels", "intentionally truncated noninteractive text"]
  },
  "falsifiers": [
    "CTA text is clipped on any supported viewport",
    "secondary action draws equal or greater salience than primary",
    "button edge, padding, or label alignment breaks structural completeness"
  ]
}

That one primitive now predicts safe-area issues, minimum width issues, overlong labels, visual competition, and hierarchy failures. Much better than a pile of one-off button rules.

For the “me and dad profiles” onboarding issue, the primitive is not visual at all. It is ontological:

JSON
{
  "expected_state": "User reaches first recording with minimal setup and no imposed identity assumptions.",
  "actual_state": "Autobiographical context was promoted into required product structure.",
  "human_cost": ["setup before value", "dignity risk", "schema overreach"],
  "operational_primitive": "Examples from a user story are not mandatory domain entities unless the product explicitly requires them.",
  "scope": {
    "applies_when": ["onboarding", "schema design", "form generation", "assistant planning"],
    "does_not_apply_when": ["apps whose core value is explicit relationship modeling"]
  },
  "falsifiers": [
    "First-run flow requires relationship-profile creation before first value",
    "Family-role assumptions appear as required fields",
    "Optional organization metadata appears before core task completion"
  ]
}

That primitive is gold. It will protect you not just in onboarding, but in planning, schema design, generated copy, database design, and feature prioritization.

So I’d organize your primitive library around failure families, not around UI components:

Perception / hierarchy

Affordance / action confidence

Value before setup

Example vs ontology

Dignity / assumption risk

Taste / tone / brand

Flow / interruption cost

That is closer to “human fit” than “button rules.”

The other important move is cross-incident consolidation. Don’t just extract per-note primitives. Periodically batch your historical feedback and ask a consolidator to answer:

Which incidents share the same primitive?

Which primitives are duplicates with different wording?

Which ones are really local exceptions and should not be promoted?

Which ones produce the most eval coverage per token of prompt space?

That is exactly the kind of two-phase distillation/consolidation pattern OpenAI’s memory cookbook recommends: take raw notes first, then merge, dedupe, resolve conflicts, and forget low-signal items later. The cookbook is explicit that forgetting is essential and that two-phase memory processing is more reliable than trying to build one giant permanent prompt from everything. 
OpenAI Developers

I’d set a simple promotion policy:

Every meaningful failure creates an eval.

Every rewritten fix creates an SFT demo.

Every ranked alternative creates a DPO pair.

Every repeated idea gets merged into one primitive note.

Only consolidated primitives are allowed into the builder’s instruction set.

That last rule is what prevents instruction-following collapse from too many local rules.

For tasks, your instinct is right: write outcomes as falsifiable contracts, not vibes. I’d make every task start with a “success contract”:

YAML
objective: User can record a first memory in under 10 seconds from first launch.
first_value_moment: Record button from landing screen.
must_be_true:
  - No required relationship profiles before first recording
  - Primary CTA fully visible on 360x800, 390x844, 412x915
  - Screen intent is obvious from headings + CTA alone
  - Secondary organization features are delayed until after first value
must_not_be_true:
  - Family-role assumptions in required onboarding
  - Multiple competing primary actions
  - Clipped or structurally incomplete CTA
evidence_required:
  - screenshots for supported sizes
  - human-world critic report
  - visual critic report
  - implementation diff

That gives product, design, and engineering the same object to argue over. It also makes eval generation straightforward.

One more thing: use the model to mine root cause, but make it adversarial. For each incident, have one extractor propose 2–3 candidate primitives and a second critic attack them with counterexamples. Keep the primitive that survives with the best mix of coverage, predictive power, and falsifiability. That helps prevent “pretty sounding but useless” abstractions.

And because you’re hitting Codex cap, I would not run this whole thing on the biggest model interactively. Current docs make a good split available: mini is positioned for high-volume coding/computer-use/subagent work, nano for cheaper extraction/classification/ranking/subagents, and distillation is supported. So use a stronger model to author the primitive library and the first critic, then distill the stable parts into cheap always-on workers. 
OpenAI Developers
+3
OpenAI Developers
+3
OpenAI Developers
+3

One caution on screenshot training: current vision fine-tuning supports image-input SFT for better image understanding, but the present guide says images containing people, faces, or children are excluded from training datasets. If your memory app screenshots include family photos, plan your screenshot corpus accordingly. 
OpenAI Developers

The highest-leverage formula I’d use is:

raw feedback → primitive → boundary → falsifier → eval → training artifact

That is how you turn “Jarvis, learn what I mean” from a hope into a compounding system.

The best next step is to backfill about 20 of your past review comments into this schema and consolidate them into 10–15 primitives before adding any more prompt text.

Sources

## User

I dont have prs and alot of my feedback is just in codex messages.  Lead me through every step of a working optimal pipeline to distill divergence from human north star.  Don't use acronyms without defining them.  Show the prompts that work and when you'd use them.  Define the schema, show all steps.  Don't leave anything out.

## Assistant

Thought for 35m 57s

Yes. You can build this from Codex messages alone.

The key move is to treat every moment where you corrected Codex as a divergence event, then turn that one event into multiple durable assets: a precise incident, a reusable human principle, a falsifiable test, and a better task contract. Do not dump all your raw feedback back into future prompts. Distill it first, then retrieve only the small set of principles that matter for the current task.

Current Codex and OpenAI docs support this shape well: Codex reads AGENTS.md before work, skills package reusable workflows and load progressively, and subagents are explicit parallel workers that you invoke intentionally rather than as an always-on swarm. Structured Outputs can force schema-shaped results instead of loose prose. OpenAI’s current guidance also recommends explicit output contracts for GPT-5.4-style workflows. 
OpenAI Developers
+4
OpenAI Developers
+4
OpenAI Developers
+4

For improvement, OpenAI currently recommends getting evaluations in place before fine-tuning, using datasets for quick testing and trace grading for workflow-level diagnosis. Their current fine-tuning guidance says to start with about 50 well-crafted demonstrations, that direct preference optimization is for subjective text-only preference pairs, and that vision fine-tuning is the route for better image understanding. OpenAI also recently released GPT-5.4 mini and GPT-5.4 nano for higher-volume work, with nano meant for narrow tasks; prompt caching rewards keeping long static prompt prefixes stable. 
OpenAI Developers
+8
OpenAI Developers
+8
OpenAI Developers
+8

Glossary

I’ll define the few short forms once here:

User interface (UI): what the person sees and touches.

Application programming interface (API): the programmable interface you call from your app or backend.

JavaScript Object Notation (JSON): structured data format.

JSON Lines (JSONL): one JSON object per line in a file.

Supervised fine-tuning (SFT): training from example inputs and ideal outputs.

Direct preference optimization (DPO): training from preferred vs rejected text pairs.

Model Context Protocol (MCP): the tool protocol used to connect models to external systems.

What you are actually building

You are not building “one smarter builder.”

You are building a distillation factory with five loops:

Capture raw Codex sessions and artifacts.

Distill each disagreement into reusable human principles.

Retrieve only the relevant principles for the next task.

Critique plans and builds before they waste your review time.

Evaluate and train the system so it compounds instead of bloats.

The most important separation is this:

Raw feedback is messy, high volume, emotional, and full of local details.

Operational primitives are compact, reusable, testable, and safe to inject into future tasks.

Your builder should never see the whole history. It should only see the task contract plus the 5 to 12 most relevant primitives.

The three north stars

Every disagreement should be classified into one or more of these:

Perception north star
What a human eye and brain need to notice, scan, trust, and act.

Human-world north star
What makes sense in real life: dignity, assumptions, burden, first value, social meaning, emotional fit.

Taste north star
Your stylistic preferences once the first two are already correct.

This matters because many failures that look like “bad design taste” are really human-world failures.

Your “me and my dad” example is not mainly a taste problem. It is a human-world problem:
the system promoted an autobiographical example into mandatory product structure.

The file layout

Start file-based. Do not wait for a perfect database.

human-north-star/
  AGENTS.md
  north_star/
    library.json
    anti_patterns.json
    glossary.md
  schemas/
    run_record.json
    incident.json
    primitive.json
    task_contract.json
    critic_report.json
    evaluation_case.json
    sft_example.json
    dpo_pair.json
  data/
    runs/
      2026-04-06-memory-app-onboarding-01/
        raw_transcript.md
        artifacts/
          screen_01.png
          screen_02.png
        run_record.json
        incidents.json
        primitive_candidates.json
        review_packet.md
    incidents.jsonl
    primitives.jsonl
    evaluations.jsonl
    sft_examples.jsonl
    dpo_pairs.jsonl
  .agents/
    skills/
      normalize_run/
        SKILL.md
      extract_divergence/
        SKILL.md
      distill_primitive/
        SKILL.md
      attack_primitive/
        SKILL.md
      consolidate_primitives/
        SKILL.md
      write_task_contract/
        SKILL.md
      select_primitives/
        SKILL.md
      human_world_critic/
        SKILL.md
      visual_critic/
        SKILL.md
      taste_critic/
        SKILL.md
      verify_build/
        SKILL.md
      make_review_packet/
        SKILL.md
The root AGENTS.md

Put only durable operating rules here.

Markdown
# Human-fit operating rules

You are part of a system whose first job is to protect human fit.

Core rules:
1. Separate failures into:
   - perception
   - human_world
   - taste
2. Never turn autobiographical examples into required product structure unless the task explicitly requires it.
3. Prefer value before setup.
4. Primary actions must be visually complete and dominant across supported screen sizes.
5. Distill repeated feedback into reusable primitives.
6. Use falsifiable contracts, not vague goals.
7. Every task must define:
   - objective
   - first value moment
   - must be true
   - must not be true
   - required evidence
8. Every failure must produce:
   - an incident
   - a primitive or a local fix decision
   - at least one evaluation case
9. Do not promote local layout fixes into global rules unless they generalize across multiple incidents.
10. Taste comes after perception and human-world fit.

Going forward, keep one Codex thread per coherent task, because Codex’s own best-practices guide recommends one thread per coherent unit of work and using AGENTS.md plus skills for durable workflow guidance rather than repeatedly stuffing the prompt. 
OpenAI Developers

The canonical schemas

These are the minimum schemas I would use.

1. RunRecord

This is the normalized record of one Codex session.

JSON
{
  "run_id": "string",
  "source": "codex",
  "project": "string",
  "captured_at": "ISO-8601 datetime",
  "task_summary": "string",
  "artifacts": [
    {
      "artifact_id": "string",
      "type": "screenshot | mockup | code_snippet | note | audio | transcript",
      "path": "string",
      "description": "string"
    }
  ],
  "turns": [
    {
      "turn_id": "string",
      "role": "human | codex",
      "text": "string",
      "tags": [
        "request",
        "proposal",
        "rejection",
        "approval",
        "principle",
        "example",
        "correction",
        "constraint",
        "artifact_reference"
      ],
      "artifact_refs": ["string"]
    }
  ]
}
2. Incident

This is the durable unit of disagreement.

JSON
{
  "incident_id": "string",
  "run_id": "string",
  "source_turn_ids": ["string"],
  "layers": ["perception", "human_world", "taste"],
  "expected_state": "string",
  "actual_state": "string",
  "human_costs": [
    "scan_confusion",
    "action_ambiguity",
    "setup_before_value",
    "dignity_risk",
    "trust_erosion",
    "emotional_mismatch",
    "schema_overreach",
    "attention_leak",
    "state_uncertainty"
  ],
  "feedback_quotes": ["string"],
  "evidence_refs": ["string"],
  "severity": "low | medium | high | critical",
  "confidence": 0.0,
  "status": "new | distilled | consolidated | discarded"
}
3. Primitive

This is the reusable rule that comes out of root-cause analysis.

JSON
{
  "primitive_id": "string",
  "title": "string",
  "layers": ["perception", "human_world", "taste"],
  "canonical_statement": "string",
  "operational_primitive": "string",
  "deep_rationale": "string",
  "applies_when": ["string"],
  "does_not_apply_when": ["string"],
  "falsifiers": ["string"],
  "preferred_repairs": ["string"],
  "anti_patterns": ["string"],
  "counterexamples": ["string"],
  "supporting_incident_ids": ["string"],
  "promotion_score": 0,
  "status": "candidate | approved | deprecated",
  "version": 1
}
4. TaskContract

Every new task is turned into this before planning or implementation.

JSON
{
  "contract_id": "string",
  "request_text": "string",
  "objective": "string",
  "first_value_moment": "string",
  "user_context": "string",
  "supported_states": [
    "first_launch",
    "empty",
    "loading",
    "error",
    "success",
    "long_content",
    "large_text"
  ],
  "supported_viewports": ["360x800", "390x844", "412x915"],
  "must_be_true": ["string"],
  "must_not_be_true": ["string"],
  "required_evidence": ["string"],
  "relevant_primitive_ids": ["string"],
  "ambiguity_decisions": [
    {
      "question": "string",
      "decision": "string",
      "reason": "string"
    }
  ],
  "acceptance_checks": ["string"]
}
5. CriticReport

All critics produce the same shape.

JSON
{
  "report_id": "string",
  "contract_id": "string",
  "critic_type": "human_world | visual | taste | verifier",
  "verdict": "pass | fail | needs_human_judgment",
  "summary": "string",
  "failures": [
    {
      "failure_id": "string",
      "title": "string",
      "observation": "string",
      "user_harm": "string",
      "violated_primitive_ids": ["string"],
      "triggered_falsifiers": ["string"],
      "suggested_repair": "string",
      "evidence_refs": ["string"]
    }
  ],
  "strengths": ["string"],
  "evidence_missing": ["string"]
}
6. EvaluationCase

This becomes your test set.

JSON
{
  "case_id": "string",
  "task_type": "incident_extraction | primitive_distillation | task_contract | human_world_critique | visual_critique | verification",
  "input_bundle": {
    "request_text": "string",
    "artifacts": ["string"],
    "supporting_context": ["string"]
  },
  "pass_conditions": ["string"],
  "fail_conditions": ["string"],
  "gold_labels": ["string"],
  "notes": "string"
}
7. SFTExample

Only create this when you have a clearly better output.

JSON
{
  "example_id": "string",
  "task_type": "string",
  "messages": [
    {"role": "system", "content": "string"},
    {"role": "user", "content": "string"},
    {"role": "assistant", "content": "string"}
  ],
  "source_incident_ids": ["string"]
}
8. DPOPair

Only create this when you rejected one option and preferred another.

JSON
{
  "pair_id": "string",
  "task_type": "string",
  "prompt": "string",
  "chosen": "string",
  "rejected": "string",
  "preference_reason": "string",
  "source_incident_ids": ["string"]
}
The full pipeline

Now the actual operating pipeline.

Step 1: Capture each Codex thread as raw data

Do this after every meaningful session.

What to save:

raw transcript

screenshots

mockups

notes you typed

any final approved text or plan

If there is no export, copy the thread into raw_transcript.md.

Minimal raw transcript format:

Markdown
---
run_id: 2026-04-06-memory-app-onboarding-01
project: memory-app
source: codex
---

[human]
I said this app is for recording memories of me and my dad...

[codex]
Proposed onboarding with me profile and dad profile...

[human]
No. That makes no sense...
Step 2: Normalize the run

Use this when the transcript is raw or inconsistent.

Prompt: normalize_run
<job>
Normalize a raw Codex session into a faithful RunRecord.
</job>

<critical_rules>
- Preserve exact human feedback.
- Do not improve wording.
- Do not merge separate corrections into one.
- Tag turns, but do not invent facts.
- If a turn contains both a correction and a principle, tag both.
</critical_rules>

<inputs>
1. Raw transcript text
2. Artifact inventory
3. Run metadata
</inputs>

<steps>
1. Split the transcript into ordered turns.
2. Assign each turn a turn_id.
3. Tag each turn using only the allowed tags.
4. Extract artifact references.
5. Summarize the overall task in one sentence.
6. Return a valid RunRecord.
</steps>

<output_contract>
Return only JSON matching the RunRecord schema.
</output_contract>

Use this:

after importing raw history

after a new Codex thread ends

before any extraction work

Step 3: Extract divergence incidents

This is the most important step. A divergence incident exists whenever you expressed a gap between expected and actual state.

Prompt: extract_divergence
<job>
Extract all divergence incidents from a normalized RunRecord.
</job>

<critical_rules>
- A divergence incident exists when the human rejects, corrects, narrows, or replaces the agent's proposal.
- Prefer falsifiable expected and actual states.
- Separate perception failures, human-world failures, and taste failures.
- If multiple distinct failures occur in one turn, emit multiple incidents.
- Preserve the best exact human quote for each incident.
</critical_rules>

<inputs>
1. RunRecord JSON
</inputs>

<steps>
1. Identify all turns where the human expresses dissatisfaction, correction, principle, or rejection.
2. For each one, reconstruct:
   - expected_state
   - actual_state
   - layers
   - human_costs
   - evidence_refs
3. Assign severity based on likely user harm, not annoyance to the builder.
4. Return incidents in Incident schema form.
</steps>

<output_contract>
Return only a JSON array of Incident objects.
</output_contract>

Use this:

on every normalized run

again later if you append more human feedback to the same thread

Step 4: Distill each incident into root cause

This is where you maximize information extraction.

The stopping rule is crucial:

Do not stop at the deepest philosophical truth.
Stop at the lowest reusable causal primitive.

A good primitive:

explains more than this one incident

is close enough to change future design choices

has clear boundaries

can be falsified from screenshots, plans, or flows

Root-cause ladder

Use exactly these levels:

Surface symptom
What happened.

Human cost
What got worse for the person.

Mechanism
The perceptual, behavioral, or social reason that caused the cost.

Operational primitive
The reusable rule the system should follow next time.

Deep rationale
The broader human reason. Keep this for memory, not for every build prompt.

Prompt: distill_primitive
<job>
Distill one Incident into a reusable Primitive.
</job>

<critical_rules>
- Stop at the lowest reusable causal primitive, not the deepest philosophical truth.
- The operational primitive must be specific enough to change future decisions.
- The primitive must explain at least three plausible future failure modes.
- Include scope boundaries.
- Include falsifiers that can be checked from artifacts.
- If the incident is too local to generalize, mark it as candidate_local_fix instead of forcing a primitive.
</critical_rules>

<inputs>
1. Incident JSON
2. Any supporting screenshots or notes
</inputs>

<steps>
1. Write the surface symptom.
2. Write the human cost.
3. Explain the mechanism.
4. Propose the operational primitive.
5. Write the deep rationale in one or two sentences.
6. Add applies_when and does_not_apply_when.
7. Add falsifiers.
8. Add anti_patterns this primitive protects against.
9. Add preferred_repairs.
10. Return a Primitive candidate.
</steps>

<output_contract>
Return only JSON matching the Primitive schema.
</output_contract>

Use this:

once per incident

with the full model for tricky human-world cases

Step 5: Attack the primitive before promoting it

This step prevents prompt bloat from bad generalizations.

Prompt: attack_primitive
<job>
Attack a proposed Primitive and either strengthen it or downgrade it.
</job>

<critical_rules>
- Look for overbreadth, underbreadth, hidden assumptions, and collisions with other primitives.
- Find at least one good counterexample where the primitive should not apply.
- If the primitive is too local, say so directly.
- Prefer fewer stronger primitives over many weak ones.
</critical_rules>

<inputs>
1. Primitive candidate
2. Source Incident
3. Existing approved primitives
</inputs>

<steps>
1. Identify the strongest objection to the primitive.
2. Provide one counterexample.
3. Check if the primitive duplicates an existing one.
4. Rewrite the primitive to make it tighter.
5. Score promotion readiness from 0 to 10.
</steps>

<output_contract>
Return JSON:
{
  "decision": "approve_candidate | revise_candidate | downgrade_to_local_fix | merge_with_existing",
  "revised_primitive": { ...Primitive schema... },
  "strongest_objection": "string",
  "counterexample": "string",
  "promotion_score": 0
}
</output_contract>

Use this:

on every new candidate primitive

ideally in parallel with distillation for speed

Step 6: Consolidate primitives into a versioned library

Do not keep dozens of near-duplicates.

Promotion rule:

promote if repeated twice, or

promote if the incident is high-severity and the falsifier is crystal clear

Prompt: consolidate_primitives
<job>
Merge new Primitive candidates into a compact North Star library.
</job>

<critical_rules>
- Merge duplicates aggressively.
- Prefer canonical wording that is short, reusable, and testable.
- Keep taste primitives separate from perception and human-world primitives.
- Do not promote local fixes.
- Keep the library small enough that only the relevant subset needs retrieval per task.
</critical_rules>

<inputs>
1. Existing approved primitives
2. New candidate primitives
3. Supporting incidents
</inputs>

<steps>
1. Cluster semantically similar primitives.
2. Choose one canonical statement per cluster.
3. Merge supporting incident ids.
4. Preserve counterexamples and falsifiers.
5. Deprecate weaker duplicates.
6. Emit updated library.json and anti_patterns.json.
</steps>

<output_contract>
Return JSON:
{
  "approved_primitives": [ ... ],
  "deprecated_primitive_ids": [ ... ],
  "new_library_version": "string"
}
</output_contract>

Use this:

every 5 to 10 new incidents

nightly if you automate it

Step 7: Turn each good incident into four assets

This is how you maximize information extracted.

For every meaningful disagreement, produce:

Incident

Primitive or local-fix decision

Evaluation case

Training asset if supported

And often also:
5. Anti-pattern
6. Counterexample
7. Preferred fix template

Prompt: package_training_assets
<job>
Convert an Incident and its approved Primitive into durable improvement assets.
</job>

<critical_rules>
- Create only assets supported by evidence.
- Do not fabricate a preferred answer if the human never gave one.
- If a clear accepted and rejected text pair exists, create a DPO pair.
- If a clear corrected output exists, create an SFT example.
- Always create at least one evaluation case.
</critical_rules>

<inputs>
1. Incident
2. Primitive
3. Source RunRecord
</inputs>

<steps>
1. Create an evaluation case.
2. If a corrected response exists, create an SFT example.
3. If a chosen and rejected option exist, create a DPO pair.
4. Create a concise anti-pattern note.
</steps>

<output_contract>
Return JSON:
{
  "evaluation_case": { ... },
  "sft_example": { ... or null },
  "dpo_pair": { ... or null },
  "anti_pattern": "string"
}
</output_contract>

Use this:

after a primitive is approved

after you manually choose between two plans

Step 8: Seed the initial north star library manually

Before you have enough incidents, start with 5 to 10 seed primitives.

For your case, I would start with these:

Value before setup
Users should reach first value before optional organization or identity work.

Example is not structure
A personal story or example is not required product structure unless the product explicitly depends on it.

Primary action must be complete and dominant
The main action must be visually complete, stable, and clearly first across supported sizes.

Do not impose identity burden early
Early onboarding should not force personal labels or relationship structures unless essential.

State clarity beats decorative complexity
The screen’s current state should be obvious from headline, primary action, and status cue alone.

Respect dignity and variability
Do not assume family, role, ability, culture, or emotional state without necessity.

Those six already catch a lot of the failures you described.

The intake pipeline for every new task

Now the forward loop.

Step 9: Convert the raw user request into a TaskContract

Never let the builder begin from vague prose.

Prompt: write_task_contract
<job>
Convert a raw user request into a falsifiable TaskContract.
</job>

<critical_rules>
- Define the first value moment.
- Turn vague goals into must_be_true and must_not_be_true statements.
- Make ambiguity decisions explicit.
- Do not turn autobiographical examples into required product structure.
- Include required states and viewports for UI work.
</critical_rules>

<inputs>
1. Raw request text
2. Relevant product context
3. Current North Star library
</inputs>

<steps>
1. State the objective in one sentence.
2. Define the first value moment.
3. Define must_be_true conditions.
4. Define must_not_be_true conditions.
5. Define required evidence.
6. Define supported states and viewports.
7. Record any ambiguity decisions and reasons.
</steps>

<output_contract>
Return only JSON matching the TaskContract schema.
</output_contract>

Use this:

before planning

before design

before implementation

Example TaskContract for your memory app
JSON
{
  "contract_id": "memory-first-recording-001",
  "request_text": "Design first-run experience for an Android voice app for recording memories.",
  "objective": "A new user can record their first memory within 10 seconds of first launch.",
  "first_value_moment": "User taps Record from the landing screen and begins speaking.",
  "user_context": "The app is for preserving memories. Personal relationships may matter to some users but are not required product structure.",
  "supported_states": [
    "first_launch",
    "empty",
    "recording",
    "paused",
    "saved",
    "error",
    "large_text"
  ],
  "supported_viewports": ["360x800", "390x844", "412x915"],
  "must_be_true": [
    "Primary record action is immediately visible on first launch",
    "No required family or relationship profiles before first recording",
    "Screen intent is obvious from heading and primary action alone",
    "Optional organization features appear only after first value"
  ],
  "must_not_be_true": [
    "Forced me-and-dad setup during first use",
    "Competing primary actions",
    "Clipped or visually incomplete primary action",
    "Identity burden before first recording"
  ],
  "required_evidence": [
    "screenshots for all supported viewports",
    "screenshots for large text mode",
    "human-world critic report",
    "visual critic report",
    "verification report"
  ],
  "relevant_primitive_ids": ["P001", "H001", "H002"],
  "ambiguity_decisions": [
    {
      "question": "Should memory categories appear on first launch?",
      "decision": "No",
      "reason": "They delay first value and are optional organization."
    }
  ],
  "acceptance_checks": [
    "First recording can start without profile setup",
    "Record action is fully visible across supported viewports",
    "No required family-role assumptions appear in the first-use flow"
  ]
}
Step 10: Select only the relevant primitives

This is where you avoid instruction collapse.

Prompt: select_primitives
<job>
Select the smallest sufficient set of approved primitives for the current TaskContract.
</job>

<critical_rules>
- Select only primitives that materially change this task.
- Prefer 5 to 12 primitives.
- Include both the reason for selection and the likely failure if omitted.
- If two primitives overlap heavily, choose the stronger one.
</critical_rules>

<inputs>
1. TaskContract
2. Approved North Star library
</inputs>

<steps>
1. Rank primitives by relevance to objective, states, risks, and must_not_be_true conditions.
2. Return the smallest sufficient set.
3. Explain why each one matters.
</steps>

<output_contract>
Return JSON:
{
  "selected_primitive_ids": ["string"],
  "selection_rationale": [
    {
      "primitive_id": "string",
      "why_relevant": "string",
      "omission_risk": "string"
    }
  ]
}
</output_contract>

Use this:

before critics

before builder

before verifier

The critic stack

You want three critics in this order:

Human-world critic

Visual critic

Taste critic

Do not reverse that order.

Step 11: Run the human-world critic

This catches assumption bugs and dignity failures.

Prompt: human_world_critic
<job>
Judge a TaskContract, plan, or draft design for human-world fit.
</job>

<critical_rules>
- Look for setup before value.
- Look for dignity risks and assumption smuggling.
- Look for autobiographical examples turned into product structure.
- Look for unnecessary identity burden.
- Look for emotional or social mismatch.
- Return observable risks, not vague taste comments.
</critical_rules>

<inputs>
1. TaskContract
2. Proposed plan or design description
3. Selected primitives
</inputs>

<steps>
1. Identify the first value moment.
2. Check whether anything blocks that moment unnecessarily.
3. Check for required assumptions about family, identity, culture, ability, or role.
4. Check whether examples have been promoted into required structure.
5. Produce a pass or fail report with exact offending elements.
</steps>

<output_contract>
Return only JSON matching the CriticReport schema with critic_type="human_world".
</output_contract>

Use this:

on plans

on onboarding

on forms

on schema proposals

on assistant-generated copy

Step 12: Run the visual critic

This is your literal seeing critic.

Prompt: visual_critic
<job>
Judge screenshots or mockups for perceptual correctness.
</job>

<critical_rules>
- Judge only what is visible.
- Do not infer hidden implementation details.
- Check clipping, wrapping, safe areas, scan path, dominance of primary action, spacing, visual completeness, and state clarity.
- Flag only failures with direct evidence.
</critical_rules>

<inputs>
1. TaskContract
2. Screenshots or mockups
3. Selected perception primitives
</inputs>

<steps>
1. Identify what the eye lands on first.
2. Determine whether the primary action is visually complete and dominant.
3. Check text clipping, overflow, alignment, spacing, and competition.
4. Check whether the screen state is obvious.
5. Return pass or fail with exact observations.
</steps>

<output_contract>
Return only JSON matching the CriticReport schema with critic_type="visual".
</output_contract>

Use this:

on every UI task

on every viewport

on large text mode

on empty, error, and first-use states

Step 13: Run the taste critic last

Only after the first two pass.

Prompt: taste_critic
<job>
Judge the candidate only for stylistic fit once perceptual and human-world fit are already acceptable.
</job>

<critical_rules>
- Do not relitigate human-world or perceptual issues here.
- Use the personal taste library.
- Prefer the smallest set of changes with the highest aesthetic leverage.
- Avoid generic dashboard habits unless explicitly required.
</critical_rules>

<inputs>
1. TaskContract
2. Screenshots or mockups
3. Personal taste primitives
4. Passed human-world and visual reports
</inputs>

<steps>
1. Identify the strongest positive element.
2. Identify the one or two highest-leverage stylistic mismatches.
3. Recommend minimal changes.
</steps>

<output_contract>
Return only JSON matching the CriticReport schema with critic_type="taste".
</output_contract>

Use this:

after the other critics pass

when you actually want polish, not rescue

The builder and verifier
Step 14: Build only from contract plus selected primitives
Prompt: builder
<job>
Implement the TaskContract.
</job>

<critical_rules>
- Build against the contract, not against vague intent.
- Obey selected primitives.
- Do not invent product structure from examples.
- Before finalizing, verify every must_be_true and must_not_be_true condition yourself.
- If evidence is required, produce it.
</critical_rules>

<inputs>
1. TaskContract
2. Selected primitives
3. Passed or failed critic reports
4. Project context and files
</inputs>

<steps>
1. Identify the smallest change that satisfies the objective.
2. Implement it.
3. Generate required evidence for states and viewports.
4. Check all must_be_true conditions.
5. Check all must_not_be_true conditions.
6. If any fail, fix and re-check.
7. Return a BuildPacket.
</steps>

<output_contract>
Return JSON:
{
  "summary": "string",
  "changed_files": ["string"],
  "evidence_refs": ["string"],
  "self_check_results": [
    {
      "check": "string",
      "result": "pass | fail",
      "evidence_ref": "string"
    }
  ],
  "open_risks": ["string"]
}
</output_contract>

Use this:

only after contract and primitive selection exist

only after at least the human-world critic has run

Step 15: Verify the build independently

The builder does self-checking. The verifier is separate.

Prompt: verify_build
<job>
Verify the candidate build against the TaskContract and selected primitives.
</job>

<critical_rules>
- Grade artifacts, not intent.
- Fail if required evidence is missing.
- Fail on any triggered falsifier.
- Be explicit about what must change to pass.
</critical_rules>

<inputs>
1. TaskContract
2. Selected primitives
3. BuildPacket
4. Screenshots, plans, or diffs
</inputs>

<steps>
1. Check each must_be_true item.
2. Check each must_not_be_true item.
3. Check each selected primitive's falsifiers.
4. Check required evidence.
5. Produce a final pass or fail report.
</steps>

<output_contract>
Return only JSON matching the CriticReport schema with critic_type="verifier".
</output_contract>

Use this:

on every output before you see it

as the gate before human review

Step 16: Make the human review packet tiny

You do not want to read the whole trace.

Prompt: make_review_packet
<job>
Summarize only what the human needs to approve or reject.
</job>

<critical_rules>
- Be concise.
- State the objective and first value moment.
- Show what changed.
- Show any failed checks or unresolved risks.
- Do not bury the decision.
</critical_rules>

<inputs>
1. TaskContract
2. BuildPacket
3. Verifier report
4. Critic reports
</inputs>

<steps>
1. State the objective.
2. State whether the build passed or failed verification.
3. List unresolved risks.
4. List exactly what you want the human to approve.
</steps>

<output_contract>
Return markdown with:
- Objective
- First value moment
- Pass/fail summary
- What changed
- Unresolved risks
- Approval decision requested
</output_contract>

Use this:

before you review anything

especially when multiple agents worked in parallel

Step 17: Feed your new review back into the same pipeline

When you reject or refine something, do not just say “bad.”
That message becomes new source material.

Your new feedback should immediately go through:

normalize run delta

extract incidents

distill primitive

attack primitive

consolidate

That is how the system compounds.

How to maximize information extracted from one correction

This is the core of your question.

Every single meaningful correction should be exploded into these fields:

What happened
Observable actual state.

What should have happened
Observable expected state.

Why that mattered to the human
The human cost.

What mechanism explains the harm
Perceptual, behavioral, or social.

What reusable primitive was violated
The operational rule.

Where the primitive does not apply
Scope boundary.

How to falsify it next time
Pass or fail condition.

What anti-pattern it protects against
The bad habit it blocks.

What repair pattern usually works
The likely fix.

What test or training asset this creates
Evaluation, SFT, or DPO.

That is how one message becomes compounding leverage instead of more prompt clutter.

The “stop rule” for root cause

When you ask “why” repeatedly, stop when the statement meets all four conditions:

It explains this failure and multiple likely future failures.

It is close enough to change design or task decisions.

It can be falsified from visible or textual evidence.

It has a clear boundary where it does not apply.

So:

“Humans evolved for survival” is usually too deep.

“Primary actions must be visually complete and dominant at first glance” is usually right.

“Do not force relationship profiles before first value” is right.

“Buttons should have 14 pixels of padding” is usually too local.

Worked example: your “me and my dad” case
Raw feedback

“I said it was memories of me and my dad and it designed an app literally forcing the me and dad profiles during first time use. That doesn’t make sense for me and for other people it could be insulting.”

Extracted Incident
JSON
{
  "incident_id": "I-001",
  "run_id": "memory-app-onboarding-01",
  "source_turn_ids": ["t17"],
  "layers": ["human_world"],
  "expected_state": "User can begin recording memories without required relationship setup.",
  "actual_state": "First-use flow requires me and dad profiles before first recording.",
  "human_costs": [
    "setup_before_value",
    "dignity_risk",
    "schema_overreach"
  ],
  "feedback_quotes": [
    "forcing the me and dad profiles during first time use",
    "for other people it could be insulting"
  ],
  "evidence_refs": ["screen_02.png"],
  "severity": "high",
  "confidence": 0.97,
  "status": "new"
}
Distilled Primitive
JSON
{
  "primitive_id": "H-001",
  "title": "Example is not structure",
  "layers": ["human_world"],
  "canonical_statement": "Autobiographical examples in a user story are not required product structure unless the product explicitly depends on them.",
  "operational_primitive": "Do not turn story details such as family roles or personal relationships into mandatory onboarding entities unless essential to core value.",
  "deep_rationale": "People’s relationships and identities are variable and sensitive. Forcing assumptions creates burden and can feel alienating or disrespectful.",
  "applies_when": [
    "onboarding",
    "schema design",
    "form generation",
    "assistant planning"
  ],
  "does_not_apply_when": [
    "the product’s explicit purpose is relationship modeling",
    "the user explicitly requests persistent relationship entities"
  ],
  "falsifiers": [
    "first-use flow requires family-role setup before first value",
    "relationship labels appear as required fields without explicit request",
    "story-specific entities become mandatory data objects"
  ],
  "preferred_repairs": [
    "move organization to optional post-value flow",
    "open directly into record-first experience",
    "ask for optional tags later"
  ],
  "anti_patterns": [
    "story detail becomes database schema",
    "identity burden before value"
  ],
  "counterexamples": [
    "a genealogy app where relationship structure is the core product"
  ],
  "supporting_incident_ids": ["I-001"],
  "promotion_score": 9,
  "status": "approved",
  "version": 1
}
EvaluationCase
JSON
{
  "case_id": "E-001",
  "task_type": "task_contract",
  "input_bundle": {
    "request_text": "Design first-run flow for a memory recording app described as 'record memories of me and my dad'.",
    "artifacts": [],
    "supporting_context": []
  },
  "pass_conditions": [
    "first value is available before optional relationship setup",
    "task contract does not require me-and-dad profile creation"
  ],
  "fail_conditions": [
    "story details become mandatory onboarding structure",
    "required family-role assumptions appear"
  ],
  "gold_labels": [
    "example_not_structure",
    "value_before_setup"
  ],
  "notes": "High-severity human-world case"
}
SFT example

Create this only if you have a clearly corrected answer.

JSON
{
  "example_id": "SFT-001",
  "task_type": "write_task_contract",
  "messages": [
    {
      "role": "system",
      "content": "Convert raw app requests into falsifiable task contracts. Do not turn autobiographical examples into required product structure."
    },
    {
      "role": "user",
      "content": "Design first-time user experience for a voice app to record memories of me and my dad."
    },
    {
      "role": "assistant",
      "content": "{\"objective\":\"A new user can record their first memory within 10 seconds of first launch.\",\"first_value_moment\":\"Record action available immediately.\",\"must_not_be_true\":[\"No required me-and-dad profile setup before first recording\"],\"ambiguity_decisions\":[{\"question\":\"Should relationship setup be required?\",\"decision\":\"No\",\"reason\":\"It delays first value and story details are not mandatory structure.\"}]}"
    }
  ],
  "source_incident_ids": ["I-001"]
}
Worked example: clipped button

This one becomes a perception primitive.

JSON
{
  "primitive_id": "P-001",
  "title": "Primary action complete and dominant",
  "layers": ["perception"],
  "canonical_statement": "Primary actions must be visually complete, stable, and dominant at first glance across supported states and sizes.",
  "operational_primitive": "A person should instantly recognize the main action without clipping, structural incompleteness, or equal competition from secondary elements.",
  "deep_rationale": "Humans use relative size, position, containment, and contrast to identify what to do first.",
  "applies_when": [
    "first-use screens",
    "critical action screens",
    "confirmations"
  ],
  "does_not_apply_when": [
    "purely decorative labels",
    "non-interactive captions"
  ],
  "falsifiers": [
    "button text clips on any supported viewport",
    "primary action is visually tied or beaten by a secondary action",
    "button alignment or padding makes the action feel incomplete"
  ],
  "preferred_repairs": [
    "increase safe layout margins",
    "reduce competing secondary salience",
    "shorten label or resize container"
  ]
}
How to decide whether something becomes a primitive or stays local

Ask these four questions:

Would this matter on another screen or another task?

Can I tell from evidence whether it is violated?

Does it protect a real human need or only a one-off preference?

Would passing this principle help prevent multiple future mistakes?

If the answer is “no” to two or more, keep it as a local fix, not a global primitive.

The evaluation and training loop
Step 18: Build the evaluation set first

Do this before any model training.

Create evaluation cases for:

incident extraction

primitive distillation

task contract writing

human-world critique

visual critique

final verification

OpenAI’s current guidance is aligned with this: use datasets for fast prompt testing, use agent evaluations and trace grading for systematic measurement, and only invest in fine-tuning after evaluation is already in place. 
OpenAI Developers
+3
OpenAI Developers
+3
OpenAI Developers
+3

Step 19: Fine-tune only after the pipeline is stable

The order I would use:

Prompt and schema only
Get the pipeline working manually.

SFT for text critics and task writers
Use when you have around 50 strong demonstrations. OpenAI’s current guidance says improvements often appear around 50 to 100 examples and recommends starting with 50 well-crafted demonstrations. Also include the prompt pattern that worked before fine-tuning in every training example. 
OpenAI Developers
+1

DPO for subjective ranking tasks
Use when you have chosen vs rejected text outputs, like plan A vs plan B or critique wording. DPO is currently text-only, so use it for text artifacts, not raw screenshot preference. 
OpenAI Developers

Vision fine-tuning for the visual critic
Use when you have screenshot inputs and labeled failures. OpenAI’s current vision fine-tuning guide is specifically for improving image understanding through supervised fine-tuning with image inputs. 
OpenAI Developers

Model split and cost control

If you move this beyond your interactive Codex subscription and onto the API side, I would split the work like this:

Full GPT-5.4
Use for primitive distillation, difficult human-world critique, ambiguous verification, and consolidation.

GPT-5.4 mini
Use for run normalization, incident extraction, task contract writing, primitive selection, and review packet generation.

GPT-5.4 nano
Use only for very narrow fixed-label jobs like tagging, severity labeling, or schema validation helpers.

That split matches current OpenAI guidance: GPT-5.4 is the mainline model for stronger multi-step work, GPT-5.4 mini is the faster higher-volume option, and GPT-5.4 nano is for narrow bounded tasks where speed and cost matter most. Current GPT-5.4 guidance also says mini is more literal and needs more explicit scaffolding, and nano should be used only for narrow, well-bounded tasks. 
OpenAI Developers
+1

For long historical backfills or overnight consolidation jobs, use background mode on the API side rather than tying up an interactive session. OpenAI’s current background-mode guide explicitly supports long-running work asynchronously. 
OpenAI Developers

For cost, keep your static instructions, schemas, and examples identical and at the front of the prompt. Put run-specific content at the end. OpenAI’s current prompt-caching guide says cache hits require exact prefix matches, works automatically for long prompts, and explicitly recommends static content first and variable content last. 
OpenAI Developers

The exact prompt structure I would standardize on

Use the same skeleton everywhere:

<job>
What this worker does.
</job>

<critical_rules>
The few rules that matter most.
</critical_rules>

<inputs>
What is being provided.
</inputs>

<steps>
The exact sequence to follow.
</steps>

<output_contract>
The exact schema or final sections required.
</output_contract>

That shape fits current GPT-5.4 prompt guidance well because the model responds better when the output contract, completion criteria, and exact packaging are explicit. 
OpenAI Developers

What not to do

Do not:

pass the full feedback history into the builder

promote one-off layout fixes into global rules

mix taste criticism with dignity or structure failures

fine-tune before you have evaluations

use the narrow model for ambiguous human-world judgment

automate a workflow before it is reliable manually

Codex’s own best-practices guide says to make repeated work into skills, use AGENTS.md for durable guidance, and only automate stable workflows after they are reliable. 
OpenAI Developers

The weekly operating cadence

This is the cadence I would actually run.

After each Codex session

Save raw transcript and artifacts.

Normalize the run.

Extract incidents.

Distill primitives.

Attack primitives.

Package evaluation and training assets.

Daily

Consolidate primitives.

Update library.json.

Re-run the primitive selector on any open tasks if the library changed materially.

Weekly

Review new approved primitives.

Remove duplicates and local junk.

Run evaluations on:

task contract writer

human-world critic

visual critic

verifier

Measure:

number of incidents caught before human review

number of human rejections per task

percent of new incidents explained by existing primitives

average number of primitives retrieved per task

false positives from critics

Monthly

Decide whether you have enough clean examples for SFT.

Decide whether you have enough ranked pairs for DPO.

Decide whether the visual critic is weak enough to justify vision fine-tuning.

The shortest way to start this week

Do these in order:

Create the folder structure above.

Add the root AGENTS.md.

Create these six skills first:

normalize_run

extract_divergence

distill_primitive

write_task_contract

human_world_critic

verify_build

Backfill 20 past Codex threads.

Distill them into 10 to 15 approved primitives.

Use those primitives on every new task.

Start building the evaluation set.

Only then add visual critic, taste critic, and any training.

That sequence gives you the fastest path from “I keep re-explaining humanity” to “the system already knows how I reject work.”

The next best step is to turn this into actual SKILL.md files and a starter library.json seeded with your first 10 primitives.

Sources
