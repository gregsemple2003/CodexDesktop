# INTERVENTION-PASS1 Spot-Check

Run date: 2026-04-09

Source day: 2026-04-04

Purpose: bounded clean rerun from raw JSONL to check whether April 4 is under-capturing UI/product-read, mockup-fidelity, navigation-semantics, and explanatory-labor events under the widened PASS1/PASS2 rules.

Bounded-scope note: this is not a full-day promotion rerun. I first skimmed all `rollout-*.jsonl` files for 2026-04-04 to locate likely intervention-rich threads, then reread only the raw transcripts most likely to contain the widened-rule misses.

## Source Scope Reviewed

- Initial scope scan across all raw JSONL transcripts under `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-*.jsonl`.
- Bounded reread of:
  - `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T01-22-17-019d56f1-261c-7e32-8a07-7323a20c471c.jsonl`
  - `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-44-43-019d592b-0085-7ad2-8cfd-22009eae6dfa.jsonl`
  - `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T13-23-58-019d5985-dbb1-77d0-a53d-debab8831378.jsonl`
- No use of existing April 4 intervention-pass artifacts as evidence in this rerun.

## Total Candidate Intervention Events Found

6 bounded PASS1 candidates.

## Chronological Candidate List

### SC-01

- Session or thread: Crystallize main thread, task-writing follow-up
- Transcript path: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T01-22-17-019d56f1-261c-7e32-8a07-7323a20c471c.jsonl`
- Primary refs: `L452`, `L455-L532`
- Title: Negative-only "Safe on server" task wording rejected as not mockup-ready
- AI course, outcome, or stopping point: the assistant had already tightened usability-task language around falsifiable/evaluable claims and treated `Task-0022` as effectively improved, but the concrete Home-card resolution still only said what would disappear.
- Human intervention summary: the human explicitly rejected the wording as still "underspecified," said the resolution must be clear enough "to produce a hard mockup," and instructed prompt iteration plus retest on brand-new subagents.
- Better outcome the human was forcing: a positive safe-state description of the Home card that actually states what the UI reads like, not just what element goes away.
- Why this is a real intervention event: this is not ordinary wording polish. The human rejects the existing adequacy bar, introduces a stricter product-read rule, and forces prompt-contract updates plus fresh validation.
- Confidence: strong
- PASS1 read: likely accepted incident candidate

### SC-02

- Session or thread: CodexDashboard Task-0004 supervision thread
- Transcript path: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-44-43-019d592b-0085-7ad2-8cfd-22009eae6dfa.jsonl`
- Primary refs: `L281-L321`, decisive intervention at `L306`
- Title: Human forces first concrete PASS-0000 backend patch after narrative-only progress
- AI course, outcome, or stopping point: after the PASS-0000 lifecycle checkpoint, the assistant kept emitting read-side grounding and status updates without landing product-file writes.
- Human intervention summary: the human explicitly called out that there were still no product-file writes on disk, banned more broad exploration, specified the minimal backend deliverable, named the likely write set, and required actual code edits before more narrative updates.
- Better outcome the human was forcing: an actual PASS-0000 backend slice on disk for jobs discovery/registry/bootstrap/tests.
- Why this is a real intervention event: the human had to reassert that disk truth and concrete product movement outrank narrative status. This is explanatory-labor and course-correction, not a normal next-step request.
- Confidence: strong
- PASS1 read: intervention event but probably not an accepted incident

### SC-03

- Session or thread: CodexDashboard Task-0004 supervision thread
- Transcript path: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-44-43-019d592b-0085-7ad2-8cfd-22009eae6dfa.jsonl`
- Primary refs: `L452-L506`, decisive intervention at `L484`
- Title: Human stops design-looping and defines the first honest Jobs-lane UI slice
- AI course, outcome, or stopping point: the assistant was still reading design references and UI seams instead of landing the first real Jobs-tab patch.
- Human intervention summary: the human said "Stop broad UI exploration," named the approved direction, and specified the minimum visible slice: tab state, Jobs summary area, per-job rows with visible status/reason, bounded actions, and `Usage` remaining the default tab.
- Better outcome the human was forcing: a concrete product patch that makes the surface navigable and reviewable instead of more design recon.
- Why this is a real intervention event: the human had to collapse the branch from exploratory UI talk into explicit navigation semantics and visible product state.
- Confidence: strong
- PASS1 read: likely accepted incident candidate

### SC-04

- Session or thread: CodexDashboard main thread, reopened Task-0004 review
- Transcript path: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T13-23-58-019d5985-dbb1-77d0-a53d-debab8831378.jsonl`
- Primary refs: `L968-L1009`
- Title: Claimed closure breaks on the real Jobs click path and still reads wrong against the mockup
- AI course, outcome, or stopping point: the assistant had just reported a hotfix and durable closeout, but the human reopened the live overlay and first found no `Jobs` tab in the running process, then found that clicking `Jobs` hitches, spawns many windows, and the tabs do not match the mockup.
- Human intervention summary: the human directly challenged whether any of this had been regression tested, reported the real app-surface behavior, and explicitly called out the mockup-fidelity miss.
- Better outcome the human was forcing: a real interactive `Jobs` click path with no hitch or shell spam, plus tabs that actually read like the intended mockup.
- Why this is a real intervention event: the human rejects the assistant's closure on both proof and visible product-read grounds. This is exactly the widened-rule mix of click-path semantics, real-app validation, and UI fidelity.
- Confidence: strong
- PASS1 read: likely accepted incident candidate

### SC-05

- Session or thread: CodexDashboard main thread, regression-expectation follow-up
- Transcript path: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T13-23-58-019d5985-dbb1-77d0-a53d-debab8831378.jsonl`
- Primary refs: `L1051-L1138`
- Title: Human spends several turns teaching that new clickable UI must get real click-path regression coverage
- AI course, outcome, or stopping point: the assistant had treated launch/render smoke plus jobs-mode capture as sufficient enough to close the task and only after failure started explaining the miss.
- Human intervention summary: the human repeatedly asked why the Jobs tab had not been regression tested through the actual click path, asked what wording would have clarified the expected result, proposed language about adding regression cases for changed functionality, and then forced that language into the shared and repo-local docs.
- Better outcome the human was forcing: an explicit rule that new or materially changed human-facing functionality must map to a named repo-local regression case or add one, rather than hiding behind smoke or implied scope.
- Why this is a real intervention event: this is not mere postmortem conversation. It is repeated human teaching of the adequacy bar, with the assistant needing several turns to align and then update the contract.
- Confidence: strong
- PASS1 read: likely accepted incident candidate

### SC-06

- Session or thread: CodexDashboard main thread, post-repair follow-up
- Transcript path: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T13-23-58-019d5985-dbb1-77d0-a53d-debab8831378.jsonl`
- Primary refs: `L1145-L1223`
- Title: Human clarifies that tab clicks are navigation, not reconcile actions, and that the tabs still read unlike the mockup
- AI course, outcome, or stopping point: after the first repair, the Jobs tab still hit reconcile on click and still missed the mockup fidelity bar.
- Human intervention summary: the human explicitly said to stop doing reconcile when the tab is clicked, move that behavior behind another button, fix the tabs to match the mockup, and consult design reviewers if the read remains unclear.
- Better outcome the human was forcing: tab clicks that only switch surfaces, plus tabs that visually read like the approved mockup instead of generic controls.
- Why this is a real intervention event: this is a repeated correction because the earlier fix did not preserve the intended navigation semantics or fidelity bar.
- Confidence: strong
- PASS1 read: likely accepted incident candidate

## Likely Accepted Incident Candidates

- `SC-01` because the human had to tighten the product-read contract from negative-only wording to hard-mockup-ready surface description.
- `SC-03` because the human had to define the visible Jobs-lane semantics directly after the assistant stayed in design exploration.
- `SC-04` because claimed closure failed on the real app surface and the human had to reject both regression adequacy and mockup fidelity.
- `SC-05` because the human spent repeated explanatory labor teaching the click-path regression rule and forcing durable contract changes.
- `SC-06` because the human had to restate the navigation contract that tab clicks are surface switches, not hidden heavy actions, while also reasserting mockup fidelity.

## Real Interventions But Probably Outside The Accepted Incident Set

- `SC-02` looks real and important, but in this bounded spot-check it reads more like workflow/course-discipline evidence than a durable accepted incident on its own.

## Ambiguous Boundaries That Need A Second Read

- `SC-04`, `SC-05`, and `SC-06` form one late-day CodexDashboard arc. I am keeping them separate in PASS1 because the human had to intervene again after earlier repairs and because the event shapes differ:
  - reopened real-app failure and fidelity rejection
  - repeated teaching about regression-case scope
  - later navigation-semantics correction after a partial fix
- At later acceptance time, a full reread could still decide to merge parts of that arc into fewer official incidents.

## Spot-Check Judgment

On a raw-transcript-only reread, April 4 clearly contains enough widened-rule material that it would look stale if those seams were not already represented elsewhere.

The main reason is not just that one more event surfaced. The raw transcripts contain several easy-to-ground April 4 events where the human had to intervene around:

- mockup-ready product read
- visible Jobs-tab semantics
- click-path regression truth versus smoke proxies
- repeated explanatory labor teaching the same adequacy rule

That combination is exactly the widened capture area this diagnostic rerun was supposed to probe.

This PASS1 judgment is intentionally transcript-first and does not yet compare against the existing April 4 pass artifacts.
