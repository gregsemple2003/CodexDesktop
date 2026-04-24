# Review Surface Survey

## Why This Research Exists

`Task-0011` grows out of a valid earlier instinct from [Task-0009](../../Task-0009/TASK.md):

- make `Tasks` the high-level human surface for dispatch and monitoring

But a new pressure is now obvious.

The dashboard is not only going to manage already-committed work.
It is also going to receive batches of incoming asks from multiple pipelines:

- Dream runs that propose option tasks
- interface-design reviews that surface UI fixes
- QA or adversarial bug sweeps that surface defects
- general-design or first-principles passes that surface unmet needs
- approval flows that need an explicit human decision
- runtime exceptions that need acknowledgement, poke, reroute, or interruption

That is not one narrow `candidate task email` problem.

It is a broader product need:

- one trustworthy place where a human can review incoming asks, understand provenance, and take the right next action without confusing proposed work with committed work

This note surveys related patterns in frontier task-management and automation tools, then turns those patterns into a concrete direction for CodexDashboard.

## Working Question

Should email remain the canonical home for candidate tasks and other incoming asks?

Current answer:

- no

Email still has value, but only as:

- notification
- digest
- reminder
- deep-link transport into the dashboard

It should not be the canonical review surface.

## Survey

### Linear: Separate Personal Inbox From Team Triage

Official docs:

- [Linear Inbox](https://linear.app/docs/inbox)
- [Linear Triage](https://linear.app/docs/triage)
- [Linear Triage Intelligence](https://linear.app/docs/triage-intelligence)

Observed pattern:

- Linear keeps `Inbox` and `Triage` separate.
- `Inbox` is the personal notification center.
- `Triage` is the team review surface for incoming issues before they enter the normal workflow.
- Triage supports concrete dispositions:
  - accept
  - decline
  - duplicate
  - snooze
- Triage also supports automation and AI-assisted routing, but that AI does not replace the review surface.

Why this matters:

- Incoming work should not dump directly into the normal work queue.
- A review layer exists before backlog or committed workflow.
- AI suggestions are useful, but they are not the same thing as accepted work.

CodexDashboard implication:

- Dream and other pipeline outputs should land in a review surface first.
- `Tasks` should remain the home for committed, dispatchable, or active work.

### Asana And ClickUp: Inbox Is A Notification Surface, Not The Work Model

Official docs:

- [Asana Inbox](https://help.asana.com/s/article/inbox?language=en_US)
- [ClickUp Inbox](https://help.clickup.com/hc/en-us/articles/33947959867543-What-is-the-Inbox)

Observed pattern:

- Both products use an Inbox for notification and attention management.
- Both allow lightweight actions like clear, save for later, snooze, or open-in-context.
- ClickUp explicitly separates notification tabs like:
  - Primary
  - Other
  - Later
  - Cleared
- ClickUp also keeps the related item context visible beside the notification.

Why this matters:

- A person often wants to review asks without fully entering the destination object first.
- Notification surfaces are good for:
  - freshness
  - reminders
  - temporary attention management
- They are not the same thing as the canonical lifecycle of tasks, approvals, or bug routing.

CodexDashboard implication:

- Email and digest delivery should behave like Inbox inputs, not as the durable place where state lives.
- The dashboard should own the durable state for:
  - reviewed
  - deferred
  - promoted
  - routed
  - dismissed

### Jira Service Management: Queues Are Canonical Triage

Official docs:

- [What are queues?](https://support.atlassian.com/jira-service-management-cloud/docs/what-are-queues/)
- [Best practices for managing queues at scale](https://support.atlassian.com/jira-service-management-cloud/docs/best-practices-for-managing-queues-at-scale/)
- [What are approvals?](https://support.atlassian.com/jira-service-management-cloud/docs/what-are-approvals/)

Observed pattern:

- Jira Service Management treats queues as the place where incoming work is organized, triaged, assigned, and acted on.
- Atlassian explicitly frames queues as a focused view of work.
- Atlassian also calls out mental load directly when discussing queue design at scale.
- The docs recommend prioritizing only the queues that matter and avoiding giant catch-all views.
- Approvals are modeled as explicit workflow steps, not loose side-band messages.

Why this matters:

- Once incoming asks grow, raw chronological feeds become weak.
- Grouping, prioritization, and focused queue design matter.
- Approval is a first-class state transition, not just a note in an email.

CodexDashboard implication:

- The review surface must support explicit ask states and dispositions.
- The dashboard should not collapse everything into one flat global inbox.
- Review should be grouped by meaning and importance, not just timestamp.

### n8n And Zapier: Human-In-The-Loop Means Paused State Plus Review Channel

Official docs:

- [n8n human-in-the-loop for AI tool calls](https://docs.n8n.io/advanced-ai/human-in-the-loop-tools/)
- [Zapier Request Approval](https://help.zapier.com/hc/en-us/articles/38731463206029-Request-approval-to-keep-your-workflow-running-with-Human-in-the-Loop)

Observed pattern:

- Both tools pause an automation or workflow at an explicit review checkpoint.
- The review request can be delivered through another channel:
  - email
  - Slack
  - other workflow
- But the canonical truth still lives in the workflow state.
- The human is reviewing a paused step with defined next actions:
  - approve
  - deny
  - sometimes edit or supply data

Why this matters:

- Delivery channel and canonical state should be separated.
- Email can notify.
- Slack can notify.
- Another workflow can notify.
- None of those should replace the durable review state.

CodexDashboard implication:

- Dream emails can remain useful.
- But the actual state of a review item must live in the dashboard or its backing store.
- Runtime exceptions, approvals, and promote-to-task actions should all be modeled as stateful asks, not just as messages.

### Frontier Pattern Summary

Across these tools, the strongest repeated pattern is:

- canonical review or triage happens in-product
- external channels deliver notifications or shortcuts back into that product

The next strongest pattern is:

- review surfaces are not the same as active-work surfaces

The third strong pattern is:

- scale requires grouped, stateful queues rather than one flat chronological stream

## What The Survey Says To Borrow

### Borrow: Separate Review From Committed Work

Keep a distinct surface for incoming asks before they become real tasks or active runs.

### Borrow: Model Dispositions Explicitly

Do not force all asks into `promote` or `ignore`.

The surface needs states and actions like:

- promote to task
- approve
- route to bug
- route to design
- merge with existing task
- defer
- dismiss
- mark duplicate

### Borrow: Preserve Provenance

A human should always be able to tell:

- what pipeline produced this item
- what batch it came from
- what artifacts support it
- what the next action will do

### Borrow: Use Email As Transport, Not Canonical State

Email is strong for:

- awareness
- reminders
- periodic summaries
- links

Email is weak for:

- dedupe
- explicit item state
- provenance collapse
- batch handling
- multi-action review

### Borrow: Design For Mental Load

Atlassian is right to call out mental load directly.

If this surface becomes a giant unstructured intake bucket, it will fail even if the data model is technically powerful.

The screen must help the human focus on:

- what needs review now
- what can wait
- what is duplicated
- what should become real work

## What To Avoid

### Avoid: Making Email The Home

This creates several failures:

- review state gets fragmented across inbox history and app state
- items are easy to lose or re-review
- batch comparisons are clumsy
- non-email asks become second-class citizens
- provenance becomes harder to preserve

### Avoid: Flattening Everything Into `Tasks`

This would confuse:

- candidate work
- approvals
- design questions
- runtime exceptions
- committed work

`Tasks` should stay clean enough that a human can trust it as the surface for work that is real, active, blocked, sleeping, or ready to dispatch.

### Avoid: A Generic Catch-All Inbox

If the product ships one giant list of miscellaneous asks, it will inherit the worst properties of both email and operator consoles.

The surface needs:

- grouped review
- explicit ask types
- meaningful defaults
- preserved provenance

## Concrete Direction

### Product Decision

CodexDashboard should add a canonical `Review` surface.

This should be a top-level tab or an equivalently first-class surface.

It should not be hidden as a sub-mode of `Tasks`, because the human job is materially different:

- `Review` = decide what incoming asks mean and what to do with them
- `Tasks` = supervise work that is already committed

### Product Identity

Use `Review` as the human-facing name.

Use `intake` as the internal plumbing concept.

Reason:

- `Review` matches the human job
- `Intake` sounds like backend plumbing
- the surface must handle more than just new work

It also handles:

- approvals
- reroutes
- exceptions
- stale asks
- candidate promotion

### Core Data Model

The product should normalize incoming information into:

1. `batch`
2. `review item`
3. `disposition`

Suggested shape:

- `batch`
  - source pipeline
  - repo
  - created at
  - count summary
  - source artifacts
  - freshness
- `review item`
  - ask type
  - title
  - one-paragraph summary
  - rationale
  - provenance
  - severity or importance
  - recommended action
  - available actions
- `disposition`
  - new
  - needs review
  - deferred
  - promoted
  - approved
  - routed
  - duplicate
  - dismissed

### Ask Types

The minimum viable ask taxonomy should likely include:

- `candidate_task`
- `approval_plan`
- `approval_task`
- `approval_closeout`
- `bug_or_defect`
- `design_gap`
- `runtime_exception`
- `follow_up_research`

### UI Shape

The best likely shape is:

- top summary strip
- left batch rail or source rail
- center grouped review queue
- right detail pane

That supports both:

- batch-level orientation
- item-level decision-making

This is better than forcing the human to read seven separate Dream candidates across multiple repos as one undifferentiated stream.

### Email Role

Email should remain:

- a digest
- a notification
- a reminder channel
- a deep-link launcher into a specific batch or review item

Email should stop being:

- the durable home of candidate tasks
- the place where review state is tracked

## Bottom Line

The survey does not support email as the canonical home for candidate tasks or other incoming asks.

It supports a different pattern:

- one canonical in-product review surface
- multiple delivery channels into it
- clear separation between review and active work

For CodexDashboard, the strongest next product move is:

- add `Review` as the canonical surface for incoming asks
- keep `Tasks` as the canonical surface for committed work
- let email become a useful but secondary delivery layer

## Sources

- [Linear Inbox](https://linear.app/docs/inbox)
- [Linear Triage](https://linear.app/docs/triage)
- [Linear Triage Intelligence](https://linear.app/docs/triage-intelligence)
- [Asana Inbox](https://help.asana.com/s/article/inbox?language=en_US)
- [ClickUp Inbox](https://help.clickup.com/hc/en-us/articles/33947959867543-What-is-the-Inbox)
- [Jira Service Management Queues](https://support.atlassian.com/jira-service-management-cloud/docs/what-are-queues/)
- [Jira Service Management queue best practices](https://support.atlassian.com/jira-service-management-cloud/docs/best-practices-for-managing-queues-at-scale/)
- [Jira Service Management approvals](https://support.atlassian.com/jira-service-management-cloud/docs/what-are-approvals/)
- [n8n human-in-the-loop for AI tool calls](https://docs.n8n.io/advanced-ai/human-in-the-loop-tools/)
- [Zapier Human in the Loop approval](https://help.zapier.com/hc/en-us/articles/38731463206029-Request-approval-to-keep-your-workflow-running-with-Human-in-the-Loop)
