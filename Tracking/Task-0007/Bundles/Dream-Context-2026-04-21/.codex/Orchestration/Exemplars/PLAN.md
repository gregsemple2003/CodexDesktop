# PLAN.md Exemplar

This file is an exemplar for the intended structure of `Tracking/Task-<id>/PLAN.md`.

It is distilled from the planning style used for Crystallize `task-0002`, but it is written as if this format were already the shared default.

# Task 0002 Implementation Plan

## Planning Intent

This file turns `TASK.md` into a bounded implementation sequence.

It should describe the intended route to done, not pass history. Executed evidence belongs in pass audits, bug notes, and regression-run artifacts instead.

## Summary

Build a small home-server backend that:

- accepts the existing Android clip upload contract unchanged
- stores raw WAV uploads durably before any `2xx`
- persists clip and job state
- runs heavier processing in a separate worker
- produces a normalized downstream artifact that later pipeline work can build on

## Fixed Defaults

- public API: `GET /healthz` and `PUT /api/v1/clips/{clipId}`
- runtime shape: one API entrypoint plus one separate worker entrypoint
- metadata store: SQLite
- artifact roots:
  - `data/raw/`
  - `data/derived/`
  - `data/quarantine/`
  - `data/runtime/`
- normalization dependency: host `ffmpeg`
- acceptance stance: durable store first, then classify as accepted or quarantined

## Pass Plan

### Pass 0000 - Scaffold And Bootstrap

Goal:

- create the package skeleton, config loading, DB bootstrap, shared fixtures, and `GET /healthz`

Build:

- project layout under the implementation home
- configuration and logging helpers
- DB bootstrap and runtime directory creation
- pytest layout and tiny generated fixture helpers

Unit Proof:

- bootstrap and helper tests stay green
- `healthz` coverage proves the service can boot against a throwaway runtime root

Exit Bar:

- the project starts locally
- health and bootstrap tests pass
- the task folder contains `TASK.md`, `PLAN.md`, `HANDOFF.md`, and a writable `Testing/` home

### Pass 0001 - Durable Ingest Primitives

Goal:

- isolate the risky ingest logic before wiring the route

Build:

- checksum helper
- light WAV sanity parser
- storage path builder
- temp-write plus fsync plus atomic-promote helper
- ingest decision logic for new, duplicate, conflicting, and malformed uploads

Unit Proof:

- parser, storage, and decision-table tests pass without needing a live server

Exit Bar:

- route code can stay thin because the core ingest behavior already exists behind tested helpers

### Pass 0002 - Happy-Path Ingest

Goal:

- accept valid uploads end to end and queue downstream work once the bytes are durable

Build:

- streaming `PUT /api/v1/clips/{clipId}` route
- clip-row insert or upsert
- canonical raw-file persistence
- job-row creation for valid accepted clips
- idempotent handling for same-checksum replays

Unit Proof:

- touched unit coverage remains green

Exit Bar:

- valid upload and idempotent duplicate integration slices pass
- one manual operator-shaped upload lands on disk and creates a queued job

### Pass 0003 - Quarantine And Recovery

Goal:

- classify bad or conflicting uploads honestly without losing the bytes

Build:

- checksum mismatch handling
- malformed WAV handling
- conflicting duplicate handling
- quarantine storage
- error recording
- reconcile logic for expired leases or interrupted ingest state

Unit Proof:

- decision logic and helper coverage prove the quarantine branches and recovery math

Exit Bar:

- quarantined uploads do not overwrite canonical data
- duplicate retries cannot create unbounded downstream work

### Pass 0004 - Worker And Derived Artifacts

Goal:

- move post-ingest processing out of the request path

Build:

- job claim and lease logic
- deterministic `run_one_job()` worker core
- `ffmpeg` normalization to mono `16 kHz` WAV
- artifact-row creation
- clip and job state transitions

Unit Proof:

- queue math and command-builder coverage stay fast and deterministic

Exit Bar:

- worker success and retry slices pass
- an upload-to-artifact smoke proves the first useful downstream artifact exists on disk

### Pass 0005 - Operator Readiness

Goal:

- make the service runnable and debuggable by someone who did not implement it

Build:

- repo runbook
- host setup notes
- operator commands for API, worker, and reconcile entrypoints
- regression matrix updates

Unit Proof:

- full unit coverage remains green after the release-bar polish

Exit Bar:

- runbook and regression docs match the current artifact layout
- the task is ready for task-level regression

## Testing Strategy

- use unit tests during each implementation pass to prove pass-local correctness
- keep integration slices focused on real DB, filesystem, and entrypoint behavior
- do not treat task-level regression as complete until the planned passes are implemented and the real app or operator lane has been exercised honestly

## Deferred Work

Keep these out of the MVP unless the task expands intentionally:

- Android client changes
- transcript generation
- sessionization
- cloud orchestration
- auth-heavy internet deployment
- UI beyond the operator runbook and service surfaces
