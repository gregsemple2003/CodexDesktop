# PASS-0000 Product Surface Contract

## Purpose

This contract freezes the first implementation target for the CodexDashboard `Tasks` tab.

It exists so implementation can proceed without rediscovering the product surface from the long design brief, generated mockup, backend contracts, and planning-gate corrections.

## Source Inputs

- [../TASK.md](../TASK.md)
- [../PLAN.md](../PLAN.md)
- [./HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md](./HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md)
- [./STITCH-PROMPT.md](./STITCH-PROMPT.md)
- [../Mockup/stitch_task_tab/screen.png](../Mockup/stitch_task_tab/screen.png)
- [../Mockup/stitch_task_tab/code.html](../Mockup/stitch_task_tab/code.html)
- [../../Task-0008/HANDOFF.md](../../Task-0008/HANDOFF.md)
- [../../Task-0011/TASK.md](../../Task-0011/TASK.md)

## Product Role

`Tasks` is the committed-work cockpit.

It should help the human answer quickly:

- what needs attention
- what is running
- what is paused, waiting, sleeping, or blocked
- what is ready to dispatch
- where to open the live working context
- what provenance makes this a real committed task

It is not:

- the intake queue
- a Dream candidate browser
- a raw backend inspector
- a transcript explorer
- a spreadsheet of task ids

## First-Release Information Architecture

The first implemented `Tasks` surface should use:

- top application navigation consistent with the existing dashboard shell
- page header with `Tasks`, freshness copy, and refresh control
- summary strip with committed-work counts:
  - `Needs you`
  - `Sleeping`
  - `Running`
  - `Blocked`
  - `Ready`
- split body:
  - left: grouped task stream
  - right: persistent selected-task detail pane

The generated mockup's fifth `Candidates` summary card is semantic drift. Replace it with `Blocked` or another committed-work state required by the current backend readback. Do not show unpromoted candidates on the `Tasks` surface.

## Grouped Task Stream

Use groups by human meaning, not raw backend taxonomy:

- `Needs Attention`
- `Running`
- `Ready to Dispatch`
- `Waiting or Blocked`
- `Sleeping / Stalled` when separate display is useful

Each row should show:

- human-readable title
- one burden or meaning summary
- visible state chip
- source provenance
- freshness text
- reason line

Rows should not show AI-run progress bars. Prefer:

- elapsed time
- last meaningful progress
- current state
- next expected event
- concise recent-change copy

## Detail Pane

The selected-task detail pane should absorb the first layer of comprehension before the human opens markdown.

Required sections:

- `Summary`
- `Why this task exists`
- `Current state`
- `What changed recently`
- `Next expected step`
- `Artifacts`
- `Actions`

The pane may show raw ids or backend details only as secondary metadata.

## Provenance Rules

`Tasks` shows committed work only.

Allowed provenance examples:

- `Authored`
- `Promoted from Dream`
- `Promoted from Review`
- `System-authored` only when the backend has durable source evidence

Forbidden committed-work labels:

- `Candidate`
- `Prov: Candidate`
- any label implying unpromoted review/intake work is a normal task

Unpromoted candidate/intake items belong to the `Review` surface owned by Task-0011 until promotion creates a committed task.

## Action Model

Human-facing action labels:

- `Dispatch`
- `Pause`
- `Poke`
- `Open Task`
- `Open Working Context`
- `Open Live Thread`
- `Open Thread`
- `Open Transcript`
- `Resume` or `Continue` only when backend readback says the existing run can continue

Task-0009 may call Task-0008 backend contracts that still use internal names such as `interrupt`, but visible copy should say `Pause` when the human is choosing to stop or hold a run.

The UI must not kill arbitrary local processes. Pause/run-control actions must be backed by backend readback and reflected durable state.

## Human Instruction Path

For paused or waiting-for-human runs, the preferred instruction path is:

1. show why the run needs the human
2. expose `Open Live Thread` or the best available thread/working-context launch target
3. let the human type directly in the live Codex/VSCodium thread

Do not make a dashboard instruction text box the main path.

If no live thread target exists, fall back to a precise task or working-context target and say direct thread continuation is unavailable.

If Task-0008 needs a Temporal resume signal after the live-thread response, Task-0008 owns bridging that response back into workflow state. Task-0009 owns the launch target and visible state/copy.

## State And Copy Rules

Use plain human-facing copy:

- `Waiting on you`
- `Paused`
- `Sleeping`
- `Blocked`
- `Ready`
- `Running`

Avoid backend-first default copy:

- `workflow`
- `activity`
- `reconcile`
- `desired state`
- `execution artifact`
- backend-internal `interrupt` as the primary action label

Every non-terminal state should include a reason line and a next expected event when backend readback provides one.

## Visual Standard

Use the active generated mockup and Stitch prompt as style references, not exhaustive pixel specs.

Preserve:

- dark control-room surface
- hard-edged, dense, operational layout
- high contrast
- crafted rows and panels
- cyan running accents
- warm urgent/needs-you accents
- serious sleeping/failed accents
- compact metadata and strong task titles

Override:

- `Candidates` summary card
- `Prov: Candidate`
- AI-run progress bars
- any visual that implies unpromoted suggestions are committed work

## Empty, Loading, Stale, And Unavailable States

The first implementation must represent:

- loading
- populated
- stale
- backend unavailable
- empty-but-healthy

These states should preserve the shell and explain what is trustworthy. Do not show an empty task list as if the system has no tasks when backend readback is unavailable or stale.

## Validation And Regression Lane

Implementation, validation, screenshots, and regression must use isolated Task-0009 lanes:

- agent-owned data
- agent-owned ports
- no human dashboard service lane
- no human dashboard app data

If isolated proof cannot run, record the blocker in task-owned testing artifacts.

## PASS-0001 Ready Baseline

`PASS-0001` can start from this contract.

The first product-code slice should build the read-only surface shell, grouped stream, selected detail pane, state rendering, provenance rendering, and non-destructive states before wiring live dispatch controls.
