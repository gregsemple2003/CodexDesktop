# Stitch Prompt

Design a Windows-first desktop app screen for a product called `CodexDashboard`.

This screen is the new `Tasks` tab and it should feel like the beating heart of the committed-work side of the product:

- a private cockpit for supervising real AI task work
- not a generic SaaS admin panel
- not a kanban board
- not a raw database table
- not the intake queue for new asks

The human need is:

- let one person see what matters, what needs them, what is running, what fell asleep, and what committed work can be dispatched next without carrying the whole system in their head

Visual direction:

- dark control-room atmosphere
- charcoal and deep slate base
- warm amber for `Needs you`
- cool cyan for `Running`
- muted olive for `Ready`
- coral or red-orange for `Sleeping` and `Failed`
- subtle texture or gradients, not flat white or generic purple UI
- typography should feel intentional and operational
- use an expressive sans for headings and a restrained monospace only for ids and machine-like metadata

Layout:

- desktop window, roughly 1440x960
- top header with:
  - `Tasks`
  - a one-line freshness/status sentence
  - a small refresh control
- second row with 5 summary cards:
  - `Needs you`
  - `Sleeping`
  - `Running`
  - `Blocked`
  - `Ready`
- main split body:
  - left side is a grouped task stream
  - right side is a persistent detail pane for the selected task

Task stream groups on the left:

- `Needs Attention`
- `Running`
- `Ready to Dispatch`
- `Waiting or Blocked`

Each task row should show:

- a strong human-readable task title
- one short burden or meaning summary
- a visible state pill
- a provenance label such as `Authored` or `Promoted`
- a freshness indicator like `Updated 12m ago`
- one reason line such as:
  - `Waiting on your approval of the plan`
  - `No progress since 2:14 PM and no wait reason recorded`
  - `Promoted from Review after Dream digest approval`
  - `Ready to dispatch after plan approval`

Selected detail pane on the right:

- `Summary`
- `Why this task exists`
- `Current state`
- `What changed recently`
- `Next expected step`
- `Artifacts`
- `Actions`

The detail pane should include clear buttons such as:

- `Dispatch`
- `Open Task`
- `Open Thread`
- `Poke`
- `Interrupt`

Important product semantics:

- `Sleeping` is a first-class state and should look serious
- `Waiting on you` should be very visible and understandable
- `Tasks` only shows committed work; incoming asks belong on a separate `Review` tab
- `Promoted` tasks must be clearly distinct from `Authored` tasks, but both are real work
- this screen should feel trustworthy and high-signal, not noisy
- action labels should be explicit and concrete

Interface style:

- bold, intentional hierarchy
- plenty of contrast
- clear grouped blocks
- cards and rows should feel crafted, not default widget styling
- no overstuffed toolbar
- no spreadsheet aesthetic
- no generic Trello columns

Make it look like a serious local supervision tool for long-running AI work on Windows.
