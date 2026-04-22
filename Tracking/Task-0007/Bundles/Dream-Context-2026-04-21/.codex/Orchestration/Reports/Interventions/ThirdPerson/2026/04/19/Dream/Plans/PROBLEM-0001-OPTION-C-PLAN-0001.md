# Problem 0001 Option C Plan 0001

## Intent

Require a human-review proof packet that labels the lane used, the limits of the evidence, and why the proof does or does not count.

## What Changes

- Add a small proof summary bundle for every approval or closeout ask.
- Force explicit statements about lane, runtime surface, and whether the evidence is supporting-only.
- Keep the proof narrative readable even when multiple evidence types exist.

## Files Or Artifact Types That Move

- Approval-facing summaries in task `HANDOFF.md`, pass closeout notes, or packet-local review briefs.
- Link bundles for screenshots, logs, and run manifests.
- Shared templates for proof summaries.

## Rollout

1. Define a standard proof summary layout with lane, proof claim, proof limits, and exact links.
2. Add the layout to approval and closeout surfaces.
3. Test it on a mixed-evidence `ThirdPerson` pass where runtime and non-runtime artifacts coexist.
4. Refine the wording until a reviewer can tell within one screen what counted and what did not.

## Success Check

- Approval packets always say what lane was used.
- Supporting-only proof is labeled as such.
- Review friction drops even when the underlying evidence mix is complex.

## Burden Reduction Under Directional Context

`Truth`: the proof story is harder to blur.

`Compassion`: the human spends less time hunting through unlabeled links.

`Tolerance`: mixed evidence can still be communicated clearly without forcing the human to reconstruct the lane logic alone.
