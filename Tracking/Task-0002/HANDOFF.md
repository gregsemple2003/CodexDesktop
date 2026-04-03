# Task 0002 Handoff

## Current Status

`Task-0002` research is complete, the plan is approved under the standing auto-approve instruction, and `PASS-0000` is the active implementation checkpoint.

## Baseline

The current product baseline comes from `Task-0001`:

- a working Windows hotkey-first dashboard
- background ingest and overlay behavior already proven locally
- weekly burn and redline concepts already established in the product direction

The new task is not a product rewrite. It is a UI redesign pass that keeps the existing hotkey-overlay/operator-console intent and reshapes the surface around the Stitch mockup.

## Next Step

Implement `PASS-0000` in `app/codex_dashboard/ui.py`:

- compress the header into the Stitch utility strip
- restyle the metrics row as compact instrument cards
- restyle the single chart field around the monolithic control-room framing
- move budget editing, startup, advisory context, and footer actions into a tighter lower operational rail

## Watchouts

- Do not let the Stitch mockup quietly widen the task into a session explorer or multi-pane console.
- Do not drop required general-design intent such as budget editing, startup toggle, weekly burn, or advisory context just because the mockup is visually minimal.
- Treat the Stitch footer labels and metric wording as composition guidance, not as fixed product semantics if they conflict with the repo-root design anchor.
- Keep ingest, persistence, and background telemetry outside the scope of this task.

## References

- `Tracking/Task-0001/HANDOFF.md`
- `Tracking/Task-0001/TASK.md`
- `Tracking/Task-0002/RESEARCH-PLAN.md`
- `Tracking/Task-0002/RESEARCH-ANALYSIS.md`
- `Tracking/Task-0002/RESEARCH.md`
- `Tracking/Task-0002/PLAN.md`
- `Design/GENERAL-DESIGN.md`
- `Design/Mockups/stitch/DESIGN.md`
- `Design/Mockups/stitch/code.html`
- `Design/Mockups/stitch/screen.png`
