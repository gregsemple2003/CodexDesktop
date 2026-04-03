# Pass 0000 Audit

## Scope

`PASS-0000` reshaped the existing Tk overlay around the Stitch-inspired composition while keeping the repo-root product intent intact.

Implemented in this pass:

- compressed the overlay shell into a tighter header, metrics strip, chart field, and footer rail
- moved interval controls into the header utility strip
- restyled the summary metrics as compact instrument cards
- restyled the chart field with a visible redline threshold and denser control-room framing
- retained budget editing, startup control, advisory context, and dismiss/quit actions
- added `interval_redline_tokens(...)` to make the chart threshold explicit
- added a smoke-mode hotkey-registration fallback so real UI smoke can still run when the hotkey chord is already claimed on the host

## Verification

Executed from repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0002/Testing/PASS-0000-UI-SMOKE-0001
```

Observed result:

- `13` unit tests passed
- the ingest smoke completed successfully against the real local telemetry tree
- the real Tk app launched in smoke mode and exported:
  - `Tracking/Task-0002/Testing/PASS-0000-UI-SMOKE-0001/overlay-chart.ps`
  - `Tracking/Task-0002/Testing/PASS-0000-UI-SMOKE-0001/overlay-summary.txt`
- the smoke summary captured:
  - last ingest status
  - selected interval
  - weekly budget
  - startup state
  - seven-day total
  - projected burn
  - current bucket value
  - advisory context

## UI Fidelity Note

The new overlay structure is grounded in:

- `Design/Mockups/stitch/code.html`
- `Design/Mockups/stitch/screen.png`
- `Design/GENERAL-DESIGN.md`

The implementation now matches the intended structural shape more closely than the previous baseline:

- compressed utility header
- compact four-card metric strip
- one focal chart field
- lower operational rail for budget/startup/actions

Pass-local proof is still slightly weaker than an ideal full-window screenshot comparison because the durable runtime artifacts are a chart postscript plus overlay summary rather than a captured full overlay image.

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| Follow the Stitch overlay composition without widening into a console browser | `app/codex_dashboard/ui.py`; `Tracking/Task-0002/RESEARCH-ANALYSIS.md`; smoke artifact | Passed |
| Keep the overlay hotkey-first and dismissible | `app/codex_dashboard/ui.py`; smoke run | Passed |
| Keep weekly burn/redline, budget editing, startup toggle, and advisory context visible | `app/codex_dashboard/ui.py`; `overlay-summary.txt` | Passed |
| Keep one token-velocity chart rather than multiple panes | `app/codex_dashboard/ui.py`; `overlay-chart.ps` | Passed |
| Keep pass-local proof separate from task-level regression | unit suite, ingest smoke, and smoke artifact only | Passed |

## Notes

- normal-mode hotkey behavior is unchanged; the hotkey fallback only applies when smoke mode is active and the host already has the chord claimed
- upstream push is currently blocked because this repo has no configured `upstream` remote
- task-level regression remains for `PASS-0001`

## Verdict

`ready_with_caveats`
