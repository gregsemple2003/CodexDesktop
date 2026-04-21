Cognition | Devin can now Schedule Devins

devin

- Overview

- Enterprise

- Pricing

- Customers

windsurf

- Overview

- Install

- Enterprise

- Pricing

- blog

- contact us

- government

- careers

Try devin Menu Devin Overview Pricing Enterprise Customers Windsurf Overview Install Enterprise Pricing Blog Contact us Government Careers Try devin Blog / Product Updates March 20, 2026

# Devin can now Schedule Devins

by The Cognition Team In this article:

Every engineering team has a quiet backlog of things that should happen on schedule, but often slip – like weekly release notes, feature flag cleanups, and onboarding QA sweeps. None of these tasks are difficult, but they slip through the cracks precisely because no single instance feels urgent enough to prioritize.

Starting today, Devin can schedule its own sessions to take care of work like this. You simply describe what should happen on a recurring basis, and Devin figures out the cadence, sets up the schedule, and runs it automatically going forward. There's no cron job to configure and no workflow builder to learn; instead, you're just telling Devin what you need done, the same way you would for any other task, and Devin takes ownership of making sure it keeps happening.

## It starts with one good session

The most natural way into this feature is to let Devin do something once and then tell it to keep doing it.

For example: you might ask Devin to audit your feature flags, and find anything that's been at full rollout for more than two weeks. Devin opens a PR to remove the flag and pings the owner. The results are exactly what you wanted, and it occurs to you that this is the kind of thing that should just happen every Monday morning.

Now, you can simply tell Devin: "Schedule this for every Monday at 9am." Devin will set up a schedule to wake up each Monday at 9am and run the same workflow, so you can stop thinking about it. One good session can save you time every day, week, or month on recurring tasks to support your work.

## Devin remembers what it did last time

One feature that makes this meaningfully different from a scheduled script is that Devin carries state between runs. It reads and writes its own notes across sessions, which means each run builds on the context of the one before it rather than starting from scratch.

This matters most for tasks that accumulate over time. If you set up a Devin to compile release notes every Friday, it won't re-summarize the PRs it already covered last week, because Devin knows where it left off, picks up from there, and gives you a clean summary of just what's new. If you have a Devin watching your #feature-requests channel every morning, it tracks which messages it's already processed, notices when themes recur, and surfaces only what's changed since yesterday. The longer a scheduled Devin runs, the more useful its output becomes, because it's building on a growing body of its own context rather than treating every run as an isolated event.

## Works with Managed Devins

Earlier this week we launched Managed Devins , which let Devin break large tasks into pieces and delegate them to parallel agents, each running in their own isolated VM. Scheduled Devins and Managed Devins compose naturally together.

That means you can set up a weekly QA pass where Devin spins up a managed Devin for each page of your application, tests them all in parallel, compiles the results into a single report, and posts it to your team's Slack channel — automatically, every Friday afternoon, without anyone needing to kick it off. The combination of parallel execution and recurring scheduling is where this starts to feel less like a developer tool and more like a member of the team that just handles things.

## Try it now

Here are a few examples of the kinds of recurring work that Scheduled Devins handle well. Click any prompt to open it directly in Devin.

### Clean up stale feature flags every Monday

Check our feature flags. If any flag has been at 100% rollout for more than 14 days, open a PR to remove it and ping the owner. Then schedule this to run every Monday at 9am.

### Compile release notes every Friday

Every Friday at 5pm, pull all PRs merged this week, categorize the changes into bug fixes, new features, and infrastructure, write release notes, and post them to #engineering in Slack.

### Run QA against staging every morning

Every morning at 8am, run QA against our staging environment. Navigate through the core user flows, take screenshots, flag any regressions or visual issues, and post a report to #engineering. Use managed Devins to test flows in parallel.

Tags:

- Product Updates

Follow us on:

- Linkedin

- Twitter [ x ]

## Related posts

Previous Next

- April 15, 2026 The Cognition Team

### Devin in Windsurf

- March 19, 2026 The Cognition Team

### Devin can now Manage Devins

- February 24, 2026 The Cognition Team

### Introducing Devin 2.2

- January 21, 2026 The Cognition Team

### Devin Review: AI to Stop Slop

- November 4, 2025 The Cognition Team

### Windsurf Codemaps: Understand Code, Before You Vibe It

- September 29, 2025 Theodor Marcu

### Announcing Devin Agent Preview with Sonnet 4.5

- August 28, 2025 The Cognition Team

### Build Your Own AI Data Analyst

- May 22, 2025 The Cognition Team

### The DeepWiki MCP Server

- May 15, 2025 The Cognition Team

### Devin 2.1

- May 5, 2025 The Cognition Team

### DeepWiki: AI docs for any repo

- April 3, 2025 The Cognition Team

### Devin 2.0

- February 26, 2025 The Cognition Team

### Devin February '25 Product Update

- January 16, 2025 The Cognition Team

### Devin January '25 Product Update

- December 23, 2024 The Cognition Team

### Devin December '24 Product Update (Part 2)

- December 3, 2024 The Cognition Team

### Devin December '24 Product Update

- September 5, 2024 The Cognition Team

### Devin September '24 Product Update

- June 5, 2024 The Cognition Team

### Devin June '24 Product Update

In this article:

### Hire [ devin The AI software engineer ] devin [ The AI software engineer ]

get started with devin learn about devin Follow us on:

- Linkedin

- Twitter | X

- Website Terms of Use

- Enterprise Terms of Service

- Platform Terms of Service

- Security

- Privacy policy

- Acceptable Use Policy

- Data Processing Addendum

- Brand
