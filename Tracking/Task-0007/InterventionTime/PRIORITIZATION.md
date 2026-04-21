# Intervention Time Prioritization

`InterventionTime` is the default rough proxy for human cost.

It should be used as:

- a first approximation to the utility cost imposed on the human
- an approximate sizing and ranking signal for where the system should focus effort
- an ordinal tool first, so `much larger`, `somewhat larger`, and `roughly similar` matter more than false numeric precision

It should not be treated as:

- a complete or absolute measure of utility
- a final moral or product judgment
- proof that all relevant human costs have already been captured in one scalar

Use rule:

- prioritize attention primarily by estimated `InterventionTime`
- remain open to situational refinement when truth-seeking, compassion, and human-sensitive judgment indicate that the rough proxy is incomplete
- make that refinement explicit in context rather than pretending the base metric already captured it

This keeps the prioritization simple enough to act on while staying honest about uncertainty and human context.
