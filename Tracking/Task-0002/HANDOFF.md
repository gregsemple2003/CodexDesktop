# Task 0002 Handoff

## Current Status

`Task-0002` is created and ready for research and planning. No implementation work has started.

## Baseline

The current product baseline comes from `Task-0001`:

- a working Windows hotkey-first dashboard
- background ingest and overlay behavior already proven locally
- weekly burn and redline concepts already established in the product direction

The new task is not a product rewrite. It is a UI redesign pass that keeps the existing hotkey-overlay/operator-console intent and reshapes the surface around the Stitch mockup.

## Next Step

Ground the redesign in the existing `app/` surface and map the Stitch composition to the current overlay implementation before any code changes.

That next step should answer:

- which current UI pieces can be retained as-is
- which pieces need to be restyled or regrouped to match the Stitch structure
- where `Design/GENERAL-DESIGN.md` intentionally overrides the Stitch mockup

## Watchouts

- Do not let the Stitch mockup quietly widen the task into a session explorer or multi-pane console.
- Do not drop required general-design intent such as budget editing, startup toggle, weekly burn, or advisory context just because the mockup is visually minimal.
- Treat the Stitch footer labels and metric wording as composition guidance, not as fixed product semantics if they conflict with the repo-root design anchor.
- Keep ingest, persistence, and background telemetry outside the scope of this task.

## References

- `Tracking/Task-0001/HANDOFF.md`
- `Tracking/Task-0001/TASK.md`
- `Design/GENERAL-DESIGN.md`
- `Design/Mockups/stitch/DESIGN.md`
- `Design/Mockups/stitch/code.html`
- `Design/Mockups/stitch/screen.png`
