# General Design

## Product Intent

CodexDashboard should feel like a private cockpit, not a normal productivity app window.

The primary use case is:

- notice a spike in token burn
- tap one global hotkey
- read the answer immediately
- dismiss the surface and keep working

## Interaction Model

- the app runs quietly in the background
- the primary entry path is a global hotkey toggle
- the visible surface is a compact overlay rather than a maximized window
- the overlay should be dismissible immediately with either the hotkey or `Escape`

## Visual Direction

- dark control-room palette rather than light spreadsheet chrome
- strong contrast between:
  - chart field
  - metric field
  - control row
- recent or active risk should use a warm red-orange accent
- normal bars should use a cool cyan accent
- typography should feel compact and operational rather than decorative

## Overlay Structure

The overlay should contain:

1. a clear title and an immediate hide path
2. a one-line status summary for the last ingest
3. interval buttons for `1m`, `5m`, `15m`, `1h`, and `1d`
4. a single bar chart focused on token velocity
5. a local seven-day total
6. a projected weekly burn line
7. a visible redline state
8. a budget editor
9. a startup toggle
10. advisory Codex weekly-window context when it is available

## Non-Goals

- multi-pane browsing
- transcript inspection
- a general session explorer
- decorative animations
- hidden or ambiguous controls
