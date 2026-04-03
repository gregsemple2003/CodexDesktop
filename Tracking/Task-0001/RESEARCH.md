# Task 0001 Research Summary

## Research Verdict

Task research is planning-ready.

## Key Decisions

- implement the first working version as a Windows-first `Python/Tkinter` app instead of blocking on `Tauri`
- use `sqlite3` for file cursors, token events, and time-bucket queries
- ingest by periodic polling with resume-safe byte offsets, not by file count heuristics
- start with a global hotkey toggle overlay
- define regression around the real desktop app surface, not around a server-only or parser-only smoke

## Why This Is The Honest Repo Fit

- the host already has Python 3.13 and `tkinter`
- the host does not currently expose the Node/Rust toolchain that a Tauri implementation needs
- the standard library is sufficient for:
  - background scanning
  - SQLite persistence
  - Win32 hotkey registration
  - Windows startup integration
  - a lightweight overlay UI

## Carry-Forward Constraints

- keep the product direction aligned with the original task intent:
  - hotkey-first
  - background-friendly
  - visually compact
  - token-velocity focused
- do not widen into a general session browser
- treat `rate_limits.secondary` as advisory context until exact weekly budget math is provable
- keep the code organized so a future Tauri rewrite can reuse the ingest and aggregation design if desired

## Planning Recommendation

Plan the work in three passes:

- `PASS-0000`: repo bootstrap, ingest engine, persistence, and supporting tests
- `PASS-0001`: desktop overlay UI, global hotkey toggle, budget config, and startup toggle
- `PASS-0002`: repo-root docs, regression harness, executed regression artifact, and closure

## Remaining Open Questions

- what exact global hotkey chord should be the default
- whether startup should use a Startup-folder launcher or a registry Run entry

Neither question blocks planning or implementation because both have safe defaults.
