# Pass 0000 Audit

## Scope

`PASS-0000` established the repo bootstrap and ingest core for CodexDashboard.

Implemented in this pass:

- initialized the local repo baseline for a new Windows-first Python desktop utility
- added repo-root `AGENTS.md`, `TESTING.md`, and `REGRESSION.md`
- added `.gitignore` and `README.md`
- implemented the Python package under `app/codex_dashboard/`
- implemented:
  - config loading and saving
  - runtime paths
  - SQLite schema and persistence
  - session file polling
  - cursor tracking
  - `token_count` parsing
  - interval aggregation
  - weekly redline projection
- added focused unit coverage for:
  - partial trailing lines
  - dedupe
  - null-info token events
  - bucket aggregation
  - weekly redline math

## Verification

Executed from repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
```

Observed result:

- `5` unit tests passed
- the real-session ingest smoke scanned `209` files
- the ingest smoke persisted `16,324` token events
- the live summary reported:
  - `last_7d_tokens=1728840441`
  - `interval=1h`
  - `current_bucket_tokens=33209201`
  - `projected_weekly_burn=5579145768`
  - `weekly_budget_tokens=8000000`
  - `over_redline=True`

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| Parse real `token_count` events rather than file counts | `app/codex_dashboard/scanner.py`; `tests/test_ingest_core.py`; scan-once smoke | Passed |
| Persist enough state to avoid double counting | `app/codex_dashboard/storage.py`; `tests/test_ingest_core.py` | Passed |
| Handle incomplete trailing lines safely | `app/codex_dashboard/scanner.py`; `tests/test_ingest_core.py` | Passed |
| Aggregate recent token usage into interval buckets | `app/codex_dashboard/aggregation.py`; `tests/test_ingest_core.py` | Passed |
| Compute weekly redline math against a configured budget | `app/codex_dashboard/aggregation.py`; scan-once smoke | Passed |
| Define repo-local truth for testing and regression | repo-root `AGENTS.md`; `TESTING.md`; `REGRESSION.md` | Passed |

## Notes

- upstream push was intentionally deferred by explicit user instruction while the prototype is still proving itself locally
- the overlay UI, hotkey toggle, and startup integration remain for `PASS-0001`

## Verdict

`ready`
