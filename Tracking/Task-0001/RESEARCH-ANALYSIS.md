# Task 0001 Research Analysis

## Problem 0001: Shell And Toolchain Fit

### Options Considered

- `Tauri 2`
- `Electron`
- `Python/Tkinter`

### Findings

- the local machine does not currently expose `node`, `npm`, `cargo`, or `rustc` on `PATH`
- the local machine does expose `python` / `py` at `3.13.12`
- standard-library `tkinter` is available and usable on this host
- the Windows Startup folder exists and can be targeted from Python without extra dependencies

### Decision

Choose `Python/Tkinter` for the first working version.

### Why

- it is the only viable stack already runnable on this host without bootstrapping a larger toolchain first
- it can satisfy the actual task requirements:
  - background polling
  - local SQLite persistence
  - Win32 global hotkey registration through `ctypes`
  - a lightweight always-on-top overlay through `tkinter`
- it keeps the prototype Windows-first, which matches the task scope

### Consequence

- the task should preserve the Tauri-style product intent, but the actual first implementation should be a Python prototype instead of blocking on Node/Rust setup

## Problem 0002: Ingest And Persistence Model

### Options Considered

- pure filesystem watcher
- pure periodic full rescan
- periodic scan with per-file byte-offset cursors in SQLite

### Findings

- the source files are append-heavy and may contain incomplete trailing lines while Codex is still writing
- dedupe must survive restarts
- file counts alone are not sufficient; the app needs event-level token deltas
- a pure watcher is unnecessary for the first version and increases moving parts

### Decision

Use periodic polling plus SQLite cursor persistence.

### Why

- polling every few seconds is enough for a personal velocity dashboard
- storing `(path, last_size, last_offset)` in SQLite makes recovery simple and avoids double counting
- the reader can discard incomplete trailing lines and retry them on the next poll

## Problem 0003: Hotkey, Overlay, And Regression

### Hotkey Decision

Start with global hotkey toggle, not hold-to-peek.

### Why

- toggle semantics are simpler and more reliable for the first Windows prototype
- the same architecture can later be extended to press-and-hold behavior if needed

### Overlay Decision

Use one hidden top-level Tk window that toggles a borderless always-on-top overlay.

### Regression Decision

Define the repo's default regression lane as:

- start the real desktop app
- point it at a real `.codex` session tree or a fixture tree with real `token_count` events
- wait for a real ingest cycle
- trigger the real overlay surface
- verify the visible interval chart and budget panel through app-generated artifact capture

### Why

- the shared workflow still requires a real app-surface lane
- for this repo, a desktop utility's honest regression proof can be an automated app-surface smoke instead of a mobile-device or browser lane
