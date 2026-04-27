# BUG-0007 Validation

Status: passing

Date: 2026-04-27

## Scope

Validation for:

- [BUG-0007](../BUG-0007.md): duplicate `Open Task` buttons have no observable
  effect
- [BUG-0010](../BUG-0010.md): `Open Task` buttons do not raise docs in
  VSCodium

## Commands

```powershell
python -m unittest tests.test_tasks_backend tests.test_desktop_support -v
```

Result: passing, 52 tests.

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
```

Result: passing, 104 tests.

## App-Surface Smoke

The dashboard was launched from the working tree with isolated smoke config and
synthetic Codex telemetry, pointed at the service backend for task readback.

Evidence:

- [BUG-0007-20260427-121303/overlay-summary.txt](./BUG-0007-20260427-121303/overlay-summary.txt)
- [BUG-0007-20260427-121303/overlay.png](./BUG-0007-20260427-121303/overlay.png)

Observed selected actions:

```text
tasks_selected_actions=Dispatch,Open Task,Open Handoff,Open Plan
```

This proves the previous visible duplicate-label issue is absent on the real
dashboard surface for the planning-task action shape.

## Unit Coverage Added

- distinct task-artifact launch labels for task folder, plan, and handoff
- fallback from unavailable backend `code` command to installed VSCodium command
- normalized `file:///C:/...` fallback launch path
