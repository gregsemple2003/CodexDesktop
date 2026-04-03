# Task 0001 Plan

## Planning Verdict

This plan is approved for execution under the standing user instruction to auto-approve plan gates that do not expand scope, spend money, delete data, or violate workflow rules.

## Implementation Strategy

Build the first working version as a Windows-first Python desktop utility with:

- a polling ingest loop over `C:\Users\gregs\.codex\sessions\`
- SQLite-backed cursor and event persistence
- a Tkinter overlay window for the chart and budget panel
- Win32 hotkey registration through `ctypes`
- a small config layer for weekly budget, polling interval, hotkey, and startup behavior

Keep the code organized so the ingest and aggregation concepts could later be ported to Tauri if the prototype succeeds.

## Pass Order

### PASS-0000 Repo Bootstrap And Ingest Core

Objective:

- create the repo-local baseline needed to build honestly in this repo
- implement the ingest engine, SQLite store, aggregation logic, and unit coverage

Expected work:

- add repo-root docs:
  - `AGENTS.md`
  - `TESTING.md`
  - `REGRESSION.md`
- add repo metadata:
  - `.gitignore`
  - `README.md`
  - Python project entrypoint
- implement:
  - config model
  - session file polling and cursor tracking
  - SQLite schema
  - token-event persistence
  - interval aggregation queries
  - weekly redline math
- add focused unit tests for ingest, dedupe, and aggregation behavior

Verification:

- run the unit test suite
- run a fixture-driven ingest smoke that proves event extraction and bucket math

Exit bar:

- repo-root docs exist and define the repo-specific truth
- ingest core can read real-shaped `token_count` events without double counting
- focused unit proof passes

### PASS-0001 Overlay UI And Desktop Integration

Objective:

- add the real desktop surface and background-friendly interaction model

Expected work:

- implement the overlay window and chart canvas
- implement visible interval switching for `1m`, `5m`, `15m`, `1h`, and `1d`
- implement global hotkey toggle
- implement weekly budget display and redline state
- implement a startup toggle using the Windows Startup folder
- add supporting tests for config and non-visual desktop integration pieces

Verification:

- run the unit test suite
- run a desktop smoke that launches the real app surface, forces an ingest refresh, toggles the overlay, and captures a task-owned artifact

Exit bar:

- the real desktop app can launch
- the hotkey path exists
- the overlay can present live aggregated data from fixture or local telemetry
- startup toggle and config persistence behave predictably

### PASS-0002 Regression, Handoff, And Closure

Objective:

- finish honest task-level proof and close the task if the implemented surface passes

Expected work:

- finalize any remaining documentation gaps
- execute the repo-root regression lane
- write `REGRESSION-RUN-0001.md`
- close any final watchouts or document blockers honestly

Verification:

- run the chosen regression lane from repo-root `REGRESSION.md`

Exit bar:

- the acceptance criteria are either satisfied or any remaining gap is explicit
- a task-owned regression artifact exists and names the actual lane run
- `HANDOFF.md` and `TASK-STATE.json` reflect the true closure state
