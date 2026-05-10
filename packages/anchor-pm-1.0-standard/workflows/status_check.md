# Workflow: Status Check

Use this workflow to inspect Anchor PM anchor health.

## PASS

All required anchor files exist, contracts reference existing module state files, and state files contain required sections.

## WARN

Use `WARN` when:

- reference material appears to be treated as a higher-level rule;
- duplicate rules appear in multiple places;
- a state file has stale or ambiguous current state;
- there are unresolved `Needs Confirmation` items.

## FAIL

Use `FAIL` when:

- `.threadsmith/current_version.md` is missing;
- `.threadsmith/contracts.md` is missing;
- a contract references a missing module state file;
- a module state file lacks required sections;
- an existing project rule would be overwritten without approval.

## Required State Sections

- `Contract`
- `Current State`
- `Open Issues`
- `Runbook`
- `History / Notes`
