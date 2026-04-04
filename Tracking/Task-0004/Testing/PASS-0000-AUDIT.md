# Pass 0000 Audit

## Scope

`PASS-0000` introduced the backend jobs model that Task 0004 needs before any Jobs tab UI can be honest.

Implemented in this pass:

- added `app/codex_dashboard/jobs.py` for:
  - file-backed registry bootstrap
  - startup-launcher and Scheduled Task discovery
  - desired-vs-observed reconciliation
  - bounded apply behavior for supported job kinds
- added the registry path helper in `app/codex_dashboard/paths.py`
- wired bootstrap into the existing CLI and app startup path through `app/codex_dashboard/__main__.py`
- fixed startup launcher writes so CRLF command text remains stable and idempotent on Windows
- added focused unit coverage in `tests/test_jobs.py`

## Verification

Executed from repo root:

```powershell
python -m unittest tests.test_jobs -v
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
```

Observed result:

- focused jobs tests passed:
  - registry path under `.codex\Orchestration`
  - bootstrap behavior
  - startup drift classification
  - scheduled-task disabled classification
  - startup apply idempotence
  - scheduled-task apply command path
- full unit coverage passed: `40` tests
- the real scan-once path completed successfully against the local telemetry tree
- bootstrap created the live managed registry at:
  - `C:\Users\gregs\.codex\Orchestration\codex-jobs-registry.json`
- the live registry now contains:
  - the CodexDashboard startup launcher
  - `Codex Daily Agentic SWE Digest`
  - `Codex Daily Physical Agents Digest`
  - `Codex Daily UE Determinism Digest`

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| File-backed managed jobs registry under `C:\Users\gregs\.codex\Orchestration` | `app/codex_dashboard/jobs.py`; live `codex-jobs-registry.json` | Passed |
| Support startup launchers and Scheduled Tasks | `app/codex_dashboard/jobs.py`; `tests/test_jobs.py` | Passed |
| Honest reconciliation states for missing, drifted, disabled, blocked, and in-sync | `app/codex_dashboard/jobs.py`; `tests/test_jobs.py` | Passed |
| Bootstrap the existing startup launcher and digest tasks without manual recreation | live registry contents from `--scan-once --print-summary` | Passed |
| Keep bootstrap on the existing app/CLI startup path | `app/codex_dashboard/__main__.py` | Passed |

## Notes

- `.codex` backup hygiene was reviewed in:
  - `C:\Users\gregs\.codex\.gitignore`
  - `C:\Users\gregs\.codex\SETUP.md`
- no `.codex` doc changes were required for this pass because the new registry is a tracked user-authored artifact under `Orchestration`, not an ignored generated path
- the pass intentionally stops short of UI work; the Jobs tab itself remains `PASS-0001`

## Verdict

`ready`
