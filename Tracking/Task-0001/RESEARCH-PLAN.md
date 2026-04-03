# Task 0001 Research Plan

## Purpose

Determine the smallest honest implementation path for the first working version of the Codex token-velocity dashboard.

## Questions To Answer

1. Which shell is the best repo-fit for this machine right now:
   - Tauri
   - Electron
   - Python/Tkinter
2. How should the app ingest `.jsonl` session telemetry safely:
   - filesystem watcher
   - periodic polling
   - hybrid polling plus cursor persistence
3. What local persistence model is sufficient for:
   - dedupe
   - restart recovery
   - interval aggregation
4. What hotkey and overlay model is viable without extra platform dependencies?
5. What regression lane should this repo use for a desktop utility where the primary proof is the real overlay surface?

## Research Inputs

- `Tracking/Task-0001/TASK.md`
- `Tracking/Task-0001/HANDOFF.md`
- local `C:\Users\gregs\.codex\sessions\` telemetry shape
- local toolchain availability on this Windows host

## Exit Bar

- choose one concrete implementation stack
- choose one concrete ingest model
- choose one concrete local persistence model
- define the first honest regression lane for this repo
- write a bounded recommendation that is planning-ready
