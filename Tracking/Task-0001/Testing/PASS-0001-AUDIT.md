# Pass 0001 Audit

## Scope

`PASS-0001` added the first real desktop surface and the Windows integration needed to use CodexDashboard as a background utility.

Implemented in this pass:

- added the repo-root design anchor at `Design/GENERAL-DESIGN.md`
- implemented the Tk overlay window
- implemented the chart canvas and visible interval switching
- implemented the global hotkey toggle path
- implemented budget editing and config persistence
- implemented Startup-folder integration
- added a smoke mode that launches the real overlay surface, exports chart artifacts, and exits
- added desktop-support unit coverage for hotkey parsing and startup launcher generation

## Verification

Executed from repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0001/Testing/UI-SMOKE-0001
```

Observed result:

- `7` unit tests passed
- the real desktop app launched in smoke mode, exported:
  - `Tracking/Task-0001/Testing/UI-SMOKE-0001/overlay-chart.ps`
  - `Tracking/Task-0001/Testing/UI-SMOKE-0001/overlay-summary.txt`
- the smoke summary captured:
  - last ingest status
  - local 7-day token total
  - projected weekly burn
  - Codex advisory weekly-window percent

## UI Fidelity Note

This repo did not previously have a mockup set, so the durable comparison anchor for this first surface is the new repo-root `Design/GENERAL-DESIGN.md`.

The smoke artifact matches that design anchor on the intended structure:

- compact overlay
- one chart field
- interval controls
- seven-day total
- projected weekly burn
- redline state
- budget editing
- startup toggle

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| Provide a hotkey-first overlay instead of a conventional main window | `app/codex_dashboard/hotkey.py`; `app/codex_dashboard/ui.py`; smoke run | Passed |
| Show interval bars for `1m`, `5m`, `15m`, `1h`, and `1d` | `app/codex_dashboard/ui.py`; `app/codex_dashboard/aggregation.py` | Passed |
| Show a visible weekly redline state | `app/codex_dashboard/ui.py`; `app/codex_dashboard/aggregation.py`; `overlay-summary.txt` | Passed |
| Support a user-configured weekly budget | `app/codex_dashboard/config.py`; `app/codex_dashboard/ui.py`; `%LOCALAPPDATA%\\CodexDashboard\\config.json` | Passed |
| Support Windows startup integration | `app/codex_dashboard/startup.py`; desktop-support unit test | Passed |

## Notes

- upstream push remains intentionally deferred by explicit user instruction while the prototype is still local-only
- task-level regression remains for `PASS-0002`

## Verdict

`ready`
