# Dogfood / Validation Thread State

## Contract

Thread contract: `Dogfood / Validation` in `docs/anchor_pm/contracts.md`

## Current State

- This repository is the first self-dogfood target.
- Documentation-level adoption has been performed.
- Anchor PM 1.0 package-first release directories now exist.
- CLI-generated adoption is intentionally deferred.
- Self-evolution execution is now owned by the Coordination thread.
- Dogfood / Validation records validation evidence and external sample results.

## Open Issues

- Run standard package dry-run on at least one real ordinary project.
- Compare future CLI output with package-generated anchors after CLI exists.
- Define what evidence is required before declaring package behavior broadly validated.

## Runbook

For each validation run, record:

- Target project
- Mode: existing-project adoption or new-project bootstrap
- Observed behavior
- Inferred issues
- Needs confirmation
- Blocking failures
- Non-blocking product notes

For self-evolution records, link to the Coordination-owned report instead of rerunning the loop independently.

## History / Notes

- Created during self-dogfood adoption.
- Self-evolution ownership moved to Coordination after deciding the project is small enough for one management thread.
