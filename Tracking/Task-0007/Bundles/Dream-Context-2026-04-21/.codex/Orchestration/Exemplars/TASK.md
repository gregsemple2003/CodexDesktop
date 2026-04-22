# Task Exemplar

This file is an exemplar for the intended structure of `Tracking/Task-<id>/TASK.md`.

Canonical task-writeup rules now live in:

- `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md`

It uses Crystallize `task-0002` as source material, but it is written as if this format were already the normative shared shape.

Use `TASK-CREATE.md` for:

- choosing the correct task writeup type
- deciding when a task must be concrete implementation, consensus, or research
- the full specificity and falsifiability bar
- optional sections such as `Expected Resolution`, `What Does Not Count`, and `Proof Plan`

The minimum base sections for most `TASK.md` files are:

- `## Title`
- `## Summary` or `## Context`
- `## Goals`
- `## Non-Goals`
- `## Implementation Home`
- `## Acceptance Criteria`

This exemplar shows one concrete implementation task shape with an optional `## Constraints And Baseline` section.

# Task 0002

## Title

Home-server ingest and processing MVP for uploaded recorder clips.

## Context

`task-0001` already established the Android capture side:

- the app saves session-scoped WAV clips locally
- upload is deferred rather than live-streamed
- the Android worker uploads raw WAV bytes to `PUT /api/v1/clips/{clipId}`
- Android currently treats any HTTP `2xx` as final acceptance and retries any non-`2xx`

What is still missing is the real server-side target that:

- accepts those uploads durably
- validates and records them
- processes them outside the request path
- gives later Android work a stable backend contract to target

## Goals

- accept clip uploads from the current Android app without changing the client contract
- durably store the raw uploaded WAV before returning success
- validate upload integrity at ingest time
- persist clip metadata and processing state
- queue post-ingest work instead of doing heavy work inline
- produce a useful normalized downstream artifact
- keep the MVP small, local-first, and debuggable

## Non-Goals

- changing the Android upload contract before the first server version exists
- doing heavy processing inline in the upload request
- building a cloud-native or multitenant service
- shipping full diarization, transcription, or summarization in the MVP
- requiring Redis, Celery, Postgres, Docker, Whisper, or pyannote in the first pass

## Constraints And Baseline

- the existing Android upload path and headers are already in use and should be preserved for MVP
- duplicate `PUT` uploads are expected and must be handled safely
- raw audio must be durably stored before any `2xx` response
- the target environment is a private home PC, mini-server, or NAS rather than internet-scale infrastructure
- local filesystem storage and SQLite are acceptable for MVP

## Implementation Home

Keep task-owned planning, handoff, testing, and research artifacts under `Tracking/Task-0002/`.

Keep the implementation under `server/memory-server/`.

## Acceptance Criteria

- `PUT /api/v1/clips/{clipId}` accepts the current Android request shape without requiring client changes
- the server only returns `2xx` after the raw WAV is stored durably, metadata is persisted, and a processing job is queued
- the server validates checksum and basic WAV structure during ingest
- repeated uploads of the same `clipId` with the same checksum are handled idempotently without duplicating stored work
- conflicting uploads for the same `clipId` are rejected or quarantined explicitly for investigation
- clip metadata and processing/job state are persisted durably enough for later debugging and recovery
- async processing produces at least one normalized artifact, such as mono `16 kHz` WAV
- an operator can verify raw artifacts, normalized artifacts, and queue or processing state on disk and in SQLite
- `GET /healthz` exists and reports the service as runnable
