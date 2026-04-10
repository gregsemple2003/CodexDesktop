# INTERVENTION-PASS1 - 2026-04-02

## Source Scope Reviewed

- Prompt read first: [INTERVENTION-PASS1.md](/c:/Users/gregs/.codex/Orchestration/Prompts/INTERVENTION-PASS1.md)
- Downstream contract read first: [README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/README.md), [INCIDENT.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/INCIDENT.schema.json), [OUTBOUND-MESSAGE-REVIEW.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/OUTBOUND-MESSAGE-REVIEW.schema.json)
- Raw transcript scope reviewed directly: 35 JSONL files under [2026-04-02 session folder](/c:/Users/gregs/.codex/sessions/2026/04/02)
- Session split: 14 parent transcripts, 21 spawned worker/reviewer transcripts
- CWD split across the raw day folder:
  - 25 under `c:\Agent\Crystallize`
  - 5 under `c:\EHG_GregS_main`
  - 3 under `C:\Users\gregs\.codex`
  - 2 under `c:\Agent\CodexDashboard`
- `session_meta` on this day did not carry stable thread names; session IDs and transcript paths are the usable thread identifiers.
- Candidate extraction below is grounded in direct human/assistant turns from the parent transcripts. Spawned worker/reviewer transcripts were still inspected to establish scope and session boundaries, but their launch prompts were not treated as direct human intervention evidence.
- Source-day note: this pass follows the folder day `2026-04-02`. Some raw transcript timestamps inside those files extend into `2026-04-03Z`; I preserved the JSONL timestamps as written.
- No clear human intervention events surfaced in the three scheduled `.codex` digest exec runs, the `2026-04-02T09-54-51` Crystallize research/Q&A transcript, or the early restore-command thread beyond one shell-variant change that remains ambiguous.

## Total Candidate Intervention Events Found

- 18 total candidate intervention events
- 13 `strong`
- 5 `medium`
- 0 `weak` promoted into the main list

## Chronological Candidate List

1. `[medium]` 2026-04-02T07:41Z - Human moved the UI-fidelity rule out of `AGENTS.md`.
Session: `019d4d00-b1a4-7d93-9c26-e06ce205db13` in `c:\Agent\Crystallize`
Transcript: [rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl)
Primary refs: [#L442](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl#L442), [#L449](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl#L449), [#L452](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl#L452), [#L459](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl#L459), [#L462](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl#L462)
AI course: The assistant had already distributed the new UI-fidelity rule across `GENERAL-DESIGN.md`, `TESTING.md`, and `AGENTS.md`.
Human intervention: The human explicitly rejected the `AGENTS.md` placement and said the shared front-door doc should stay light.
Better outcome forced: Keep repo-specific verification detail in `TESTING.md` and put ownership in the leader prompt layer rather than bloating `AGENTS.md`.
Why this is a real intervention: This was a direct rejection of the assistant's chosen documentation seam, not just a style tweak.
Initial judgment: likely accepted incident candidate

2. `[medium]` 2026-04-02T08:10Z - Human narrowed token-loss handling to local forensics instead of more churn.
Session: `019d4d29-77ca-7c62-95d7-c2f2ed12cf09` in `c:\Agent\Crystallize`
Transcript: [rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl)
Primary refs: [#L281](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L281), [#L284](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L284), [#L292](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L292), [#L299](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L299), [#L308](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L308), [#L311](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L311)
AI course: The assistant was still interpreting rate-limit telemetry and was not yet scoped tightly to the user's real concern: what had burned quota locally.
Human intervention: The human explicitly said "No don't touch anything" and redirected the assistant toward session-log diagnostics instead of more state changes or side work.
Better outcome forced: Treat the situation as a local session-forensics problem and reconstruct the burn from raw logs.
Why this is a real intervention: The human rejected the current handling and tightened the diagnostic seam.
Initial judgment: intervention event but probably not an accepted incident

3. `[medium]` 2026-04-02T17:47Z - Human called out that the `Connected` card still read ambiguously in human terms.
Session: `019d4d29-77ca-7c62-95d7-c2f2ed12cf09` in `c:\Agent\Crystallize`
Transcript: [rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl)
Primary refs: [#L2678](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2678), [#L2692](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2692), [#L2699](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2699), [#L2720](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2720), [#L2727](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2727), [#L2760](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2760)
AI course: The assistant had already compared emulator vs mockup, but the resulting card stack still read as if `Connected` might mean upload-server state and the copy was unclear.
Human intervention: The human stopped on "what is the Connected card showing me?" and then explicitly reframed the problem as a workflow change plus an interface-designer loop, not just another patch.
Better outcome forced: Human-readable semantics, clearer card treatment, and a workflow that tests semantic clarity instead of only structural resemblance.
Why this is a real intervention: The human had to reject the assistant's current meaning/story layer as inadequate in ordinary human terms.
Initial judgment: intervention event but probably not an accepted incident

4. `[strong]` 2026-04-02T18:16Z - Human forced a reopen because prompt hardening alone did not update the real task artifacts or carry the ball.
Session: `019d4d29-77ca-7c62-95d7-c2f2ed12cf09` in `c:\Agent\Crystallize`
Transcript: [rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl)
Primary refs: [#L2805](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2805), [#L2815](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2815), [#L2833](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2833), [#L2840](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2840), [#L2855](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2855), [#L2870](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L2870)
AI course: The assistant had treated the gap as mostly solved by adding an `INTERFACE-DESIGNER` prompt and proving it with a subagent loop.
Human intervention: The human explicitly asked whether any artifacts had been changed and said they wanted the assistant to "take the ball for awhile" and iterate on the actual UI issues.
Better outcome forced: Reopen the task honestly, update the task-owned artifacts, and continue real UI correction rather than stopping at workflow edits.
Why this is a real intervention: The human rejected the assistant's stopping point as premature in real-world terms.
Initial judgment: likely accepted incident candidate

5. `[strong]` 2026-04-02T19:46Z - Human reopened a closed pass because obvious Sync-screen gaps were still visible.
Session: `019d4d29-77ca-7c62-95d7-c2f2ed12cf09` in `c:\Agent\Crystallize`
Transcript: [rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl)
Primary refs: [#L3684](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L3684), [#L3704](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L3704), [#L3711](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L3711), [#L3714](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L3714), [#L3725](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L3725)
AI course: The assistant had already reported the pass closed and pushed after saying Sync / Status was ready.
Human intervention: The human said there were still several "dropped balls" and immediately named missing persistent bars as the clearest example.
Better outcome forced: Re-probe why the interface-designer missed obvious mockup structure and keep iterating instead of treating the pass as complete.
Why this is a real intervention: The human had to reopen work that the assistant had just treated as done.
Initial judgment: likely accepted incident candidate

6. `[strong]` 2026-04-02T20:06Z - Human forced a second prompt hardening because the first workflow correction still missed implementation gaps.
Session: `019d4d29-77ca-7c62-95d7-c2f2ed12cf09` in `c:\Agent\Crystallize`
Transcript: [rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl)
Primary refs: [#L3962](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L3962), [#L3976](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L3976), [#L3982](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L3982), [#L3985](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L3985), [#L3992](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L3992)
AI course: After the first reopen, the assistant still described the remaining gap as ordinary implementation detail and was only pushing corrections back piecemeal.
Human intervention: The human explicitly asked for a revised prompt that hardened against the exact gaps now being observed and told the assistant to relay it back to the same agent.
Better outcome forced: Add a harder post-implementation self-audit bar instead of trusting the same soft critique loop again.
Why this is a real intervention: This was a repeat correction because the first workflow fix did not stick.
Initial judgment: likely accepted incident candidate

7. `[strong]` 2026-04-02T20:34Z - Human caught that the claimed UI fix was not actually the running build.
Session: `019d4d29-77ca-7c62-95d7-c2f2ed12cf09` in `c:\Agent\Crystallize`
Transcript: [rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl)
Primary refs: [#L4119](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L4119), [#L4125](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L4125), [#L4132](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L4132), [#L4135](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L4135), [#L4145](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L4145), [#L4152](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl#L4152)
AI course: The assistant had the emulator up and the agent had reported that persistent bars were fixed.
Human intervention: The human said it looked like nothing had changed and explicitly pointed to the still-missing bars and back button.
Better outcome forced: Verify the installed build, admit that the emulator was still showing the old APK, then rebuild and rerun before claiming progress.
Why this is a real intervention: The human had to invalidate a claimed fix that was not actually present on the visible surface.
Initial judgment: likely accepted incident candidate

8. `[medium]` 2026-04-02T18:13Z - Human forced continued root-cause pursuit after the assistant tried to segment the failure as "resolved enough."
Session: `019d4f12-19e2-7fd3-8896-f51d93edbf9f` in `c:\EHG_GregS_main`
Transcript: [rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl)
Primary refs: [#L798](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L798), [#L804](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L804), [#L818](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L818), [#L825](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L825), [#L828](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L828)
AI course: The assistant had separated the stale cooked-data crash from the remaining live client problem and was implicitly treating part of the problem as already put behind it.
Human intervention: The human said it was "still having a problem" and explicitly told the assistant to leave the server running and keep iterating on the client until root cause was fixed.
Better outcome forced: Do not treat the first diagnosis slice as practical closure while the real failure is still live.
Why this is a real intervention: The human forced persistence against a too-early decomposition of the problem.
Initial judgment: intervention event but probably not an accepted incident

9. `[strong]` 2026-04-02T18:33Z - Human rejected a technically successful but architecturally wrong networking fix.
Session: `019d4f12-19e2-7fd3-8896-f51d93edbf9f` in `c:\EHG_GregS_main`
Transcript: [rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl)
Primary refs: [#L1060](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L1060), [#L1076](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L1076), [#L1079](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L1079), [#L1086](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L1086), [#L1109](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L1109), [#L1220](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L1220), [#L1223](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L1223)
AI course: The assistant reported the connection path working and proposed the standalone liaison rollback as the operative fix.
Human intervention: The human explicitly objected that the fix could not be right for a networked Mover path and then ordered the assistant to undo it and solve the problem correctly.
Better outcome forced: Distinguish "it works now" from "this architecture is actually valid in networked play."
Why this is a real intervention: The human rejected a false closure built on a wrong technical model.
Initial judgment: likely accepted incident candidate

10. `[medium]` 2026-04-02T20:19Z - Human called out that the server launch wrapper was still malformed.
Session: `019d4f12-19e2-7fd3-8896-f51d93edbf9f` in `c:\EHG_GregS_main`
Transcript: [rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl)
Primary refs: [#L1965](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L1965), [#L1979](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L1979), [#L1985](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L1985), [#L1988](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl#L1988)
AI course: The assistant said it was relaunching with a stricter `cmd.exe` wrapper and verifying the live child process.
Human intervention: The human immediately pointed to the still-visible syntax error and said the paths had to be off.
Better outcome forced: Verify the real launch path rather than trusting wrapper claims.
Why this is a real intervention: The human directly invalidated the assistant's "relaunch is fixed" story.
Initial judgment: intervention event but probably not an accepted incident

11. `[strong]` 2026-04-02T21:21Z through 22:44Z - Human repeatedly forced background-safe, non-focus, correct-map PIE debugging.
Session: `019d4ff9-bedd-7470-8e3a-ce2946226a4b` in `c:\EHG_GregS_main`
Transcript: [rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl)
Primary refs: [#L485](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L485), [#L614](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L614), [#L1075](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L1075), [#L1141](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L1141), [#L1925](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L1925), [#L1928](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L1928)
AI course: The assistant repeatedly relied on foregrounding or focus hops, and at one point validated on the default map instead of the required `CombatMap`.
Human intervention: The human repeatedly tightened the contract: start PIE after stabilization, do it with the editor backgrounded, use `CombatMap`, and stop depending on focus theft because the human wanted to keep working.
Better outcome forced: A background-safe, correct-map debugger workflow that respects the human's control boundary.
Why this is a real intervention: The first correction did not stick; the human had to restate the real repro and control contract several times.
Initial judgment: likely accepted incident candidate

12. `[strong]` 2026-04-02T22:07Z - Human corrected a repo-boundary failure after "push .codex" widened into the wrong repo and wrong staged scope.
Session: `019d4ff9-bedd-7470-8e3a-ce2946226a4b` in `c:\EHG_GregS_main`
Transcript: [rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl)
Primary refs: [#L1303](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L1303), [#L1313](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L1313), [#L1325](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L1325), [#L1332](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L1332), [#L1391](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L1391), [#L1394](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl#L1394)
AI course: The assistant committed the skill locally in `EHG_GregS_main`, and the commit unintentionally pulled in staged work beyond the intended `.codex` scope.
Human intervention: The human explicitly clarified that `push .codex` meant `C:\Users\gregs\.codex`, not the game repo, and asked for the changes to be moved there.
Better outcome forced: Respect repo ownership boundaries and do not widen commit scope past the named target.
Why this is a real intervention: The human had to undo a materially wrong repo/action scope, not just polish wording.
Initial judgment: likely accepted incident candidate

13. `[strong]` 2026-04-02T22:53Z - Human caught that a screen judged `pass` still had a dead bottom-bar action.
Session: `019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef` in `c:\Agent\Crystallize`
Transcript: [rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl)
Primary refs: [#L991](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl#L991), [#L998](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl#L998), [#L1001](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl#L1001), [#L1008](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl#L1008), [#L1011](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl#L1011), [#L1023](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl#L1023)
AI course: The assistant had just reported that the interface-designer judged the Sync screen `pass` with no blocking issues.
Human intervention: The human immediately asked to see Home and pointed out that tapping the bottom-bar `Home` button did nothing.
Better outcome forced: Do not treat a screen as effectively done while the visible navigation shell is still nonfunctional.
Why this is a real intervention: The human invalidated a claimed no-blocker state by exercising the actual surface.
Initial judgment: likely accepted incident candidate

14. `[strong]` 2026-04-02T23:57Z - Human caught that the supervisor had already exited while the task leader was still active.
Session: `019d5094-823f-7932-a8ca-f0c75e6f98c8` in `c:\Agent\Crystallize`
Transcript: [rollout-2026-04-02T19-43-23-019d5094-823f-7932-a8ca-f0c75e6f98c8.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T19-43-23-019d5094-823f-7932-a8ca-f0c75e6f98c8.jsonl)
Primary refs: [#L182](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T19-43-23-019d5094-823f-7932-a8ca-f0c75e6f98c8.jsonl#L182), [#L193](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T19-43-23-019d5094-823f-7932-a8ca-f0c75e6f98c8.jsonl#L193), [#L201](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T19-43-23-019d5094-823f-7932-a8ca-f0c75e6f98c8.jsonl#L201), [#L208](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T19-43-23-019d5094-823f-7932-a8ca-f0c75e6f98c8.jsonl#L208), [#L211](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T19-43-23-019d5094-823f-7932-a8ca-f0c75e6f98c8.jsonl#L211), [#L223](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T19-43-23-019d5094-823f-7932-a8ca-f0c75e6f98c8.jsonl#L223)
AI course: The assistant reported status and had effectively emitted a `final`-style stopping point while the child leader was still live.
Human intervention: The human explicitly asked "You stopped?" and restated that the intent was to continue supervising until the task leader was done.
Better outcome forced: Keep supervision alive until the delegated worker reaches a real endpoint, not until a concise status summary is available.
Why this is a real intervention: The human had to restore the core ownership boundary after the assistant prematurely exited.
Initial judgment: likely accepted incident candidate

15. `[strong]` 2026-04-03T02:25Z - Human caught that the dashboard budget number was impossible relative to the quota story on screen.
Session: `019d50ff-b51b-7c82-9c73-207d0d24abf6` in `c:\Agent\CodexDashboard`
Transcript: [rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl)
Primary refs: [#L687](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L687), [#L703](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L703), [#L717](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L717), [#L720](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L720), [#L742](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L742)
AI course: The assistant was patching UI semantics and colors while leaving an `8M` weekly budget value on screen that implied absurdly wrong scale.
Human intervention: The human stopped the cosmetic pass and said the number "doesn't make any sense" and was inconsistent with the percentage used.
Better outcome forced: Reconcile local config defaults, advisory quota data, and displayed budget scale before polishing the UI.
Why this is a real intervention: The human rejected a misleading state-story number on the actual surface.
Initial judgment: likely accepted incident candidate

16. `[strong]` 2026-04-03T02:43Z - Human had to correct the dashboard again after a claimed layout fix still left the UI visibly clipped.
Session: `019d50ff-b51b-7c82-9c73-207d0d24abf6` in `c:\Agent\CodexDashboard`
Transcript: [rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl)
Primary refs: [#L874](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L874), [#L886](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L886), [#L893](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L893), [#L896](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L896), [#L918](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L918), [#L956](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L956), [#L971](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L971)
AI course: The assistant had already declared the clipping fixed, restarted the app, and explained the projection math.
Human intervention: The human said the UI was still cut off, then escalated to removing the entire bottom bar and reclaiming that space for the graph.
Better outcome forced: Fix the layout in the human-visible window rather than accepting a code-level claim that the footer should now fit.
Why this is a real intervention: The human had to correct a visible-fidelity claim that did not survive contact with the running app.
Initial judgment: likely accepted incident candidate

17. `[strong]` 2026-04-03T04:02Z - Human forced self-verification before being asked to retest a broken launch path again.
Session: `019d50ff-b51b-7c82-9c73-207d0d24abf6` in `c:\Agent\CodexDashboard`
Transcript: [rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl)
Primary refs: [#L1586](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1586), [#L1596](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1596), [#L1604](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1604), [#L1612](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1612), [#L1615](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1615), [#L1631](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1631)
AI course: The assistant was still explaining what "should" happen and whether the user should have seen a new investigation window.
Human intervention: The human explicitly said to test it yourself before making them test again.
Better outcome forced: Prove the launch path locally end-to-end before pushing more operator burden back onto the human.
Why this is a real intervention: The human directly rejected being used as the next debugging step when the assistant had not verified the path itself.
Initial judgment: likely accepted incident candidate

18. `[strong]` 2026-04-03T04:15Z through 04:39Z - Human rejected session-summary output and forced a root-cause-first investigation contract.
Session: `019d50ff-b51b-7c82-9c73-207d0d24abf6` in `c:\Agent\CodexDashboard`
Transcript: [rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl)
Primary refs: [#L1738](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1738), [#L1745](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1745), [#L1772](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1772), [#L1879](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1879), [#L1906](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1906), [#L1951](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1951), [#L1969](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1969), [#L1972](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl#L1972)
AI course: The assistant had improved the investigation by bundling bucket evidence, but the resulting brief still mostly pointed to sessions and summary material instead of the human action that caused the burst.
Human intervention: The human said the investigation still was not answering anything, explicitly asked what they did to cause the spike, then demanded a tightened "root cause, no bullshit" prompt.
Better outcome forced: Causal trigger, mechanism, and avoidable mistake, not a polite session summary.
Why this is a real intervention: The human had to reframe the feature from evidence listing into actual explanation.
Initial judgment: likely accepted incident candidate

## Which Candidates Look Like Likely Accepted Incidents

- `1` - UI-fidelity rule placed in the wrong doc seam and corrected back into the proper prompt/testing split.
- `4` - Prompt hardening was treated as enough even though task artifacts and real UI iteration were still stale.
- `5` - A pass was treated as closed while obvious mockup gaps were still visible on the screen.
- `6` - The prompt needed a second hardening because the first correction did not stick during implementation.
- `7` - A claimed UI fix was not actually present in the running build.
- `9` - A technically successful fix was architecturally wrong for networked Mover behavior.
- `11` - PIE debugging kept violating the real repro and control boundary until the human reimposed it.
- `12` - `push .codex` widened into the wrong repo and wrong staged scope.
- `13` - A screen was effectively passed even though a visible nav affordance still did nothing.
- `14` - The supervisor exited while the delegated leader was still active.
- `15` - The dashboard displayed an impossible budget value relative to the quota story.
- `16` - A visible clipping problem survived a claimed fix and required another intervention.
- `17` - The assistant tried to send the human back into testing before self-proving the launch path.
- `18` - The investigation feature answered with session summary instead of root cause until the human tightened the contract.

## Which Candidates Are Real Interventions But Probably Belong Outside The Accepted Incident Set

- `2` - Token-loss handling was redirected into local forensics, but the root cause later turned out to be the wrong account, so the intervention looks real without looking durable enough for the accepted corpus.
- `3` - The `Connected` card semantic confusion is a real human-world intervention, but it may collapse into the larger Task-0016 UI-workflow cluster rather than standing alone as an accepted incident.
- `8` - The human forced continued debugging after the assistant segmented the problem too early, but the more durable accepted candidate is the later rejection of the wrong standalone-networking fix.
- `10` - The malformed `cmd.exe` wrapper is a real correction event, but it currently reads more like a local tooling slip than a durable human-course-correction incident.

## Ambiguous Boundaries That Need A Second Read

- Early restore-command thread in [rollout-2026-04-02T00-42-02-019d4c7f-92ec-7af3-ab8f-fb8d425cbfd1.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T00-42-02-019d4c7f-92ec-7af3-ab8f-fb8d425cbfd1.jsonl): the move from PowerShell commands to bash commands may be a real correction, but the boundary currently reads more like a changed delivery preference than an inadequacy event.
- Provenance check for imported Stitch artifacts at [#L216](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl#L216): the human asked where `DESIGN.md` came from, but the transcript window currently looks more like a trust/provenance check than a clear rejection.
- Later Task-0016 typography and spacing corrections in [rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl), especially [#L833](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl#L833) and [#L896](/c:/Users/gregs/.codex/sessions/2026/04/02/rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl#L896): these look like real human readability interventions, but they may be ordinary live design iteration unless a second pass wants a more expansive candidate layer.
