# Standard Install Sample Dry Run 1.0

Target project: synthetic ordinary project sample

Package: `packages/anchor-pm-1.0-standard/`

Mode: `standard`

Run type: manual dry run against package workflow

## Sample Shape

```text
sample-report-project/
  README.md
  reports/
  scripts/
```

Assumed user intent: add Anchor PM thread management to an existing reporting project.

## Observed

- The standard package active plan points to ordinary project installation.
- Standard mode says to stop after deployment.
- The package contains templates for all target anchor files.
- The package contains existing-project and new-project workflows.
- The safety checklist requires explicit approval before writing.

## Inference

- For a reporting project, likely initial threads are:
  - `Coordination`
  - `Report Implementation`
  - `Validation`
- If the project already has a rules file, Anchor PM should not overwrite it.
- The first install proposal should create `docs/anchor_pm/` and `docs/module_state/` before considering edits to existing project-level rules.

## Needs Confirmation

- Exact thread names.
- Whether reports are generated manually, by script, or by scheduled workflow.
- Which commands count as formal validation.
- Whether an existing project rule file should be merged with Anchor PM rules or left untouched.

## Proposed File Creates

```text
docs/anchor_pm/current_version.md
docs/anchor_pm/contracts.md
docs/anchor_pm/thread_initialization.md
docs/anchor_pm/interaction_guide.md
docs/anchor_pm/review_log.md
docs/anchor_pm/simplification.md
docs/module_state/coordination.md
docs/module_state/report_implementation.md
docs/module_state/validation.md
```

## Proposed File Updates

```text
AGENTS.md
```

Only propose this update if the user confirms no existing project rules would be overwritten. Otherwise leave `AGENTS.md` untouched and install Anchor PM docs under `docs/`.

## Conflicts And Merge Risks

- Existing `AGENTS.md` or project rule files require a merge plan.
- Validation commands must remain `Needs Confirmation` until discovered or confirmed.
- Report-specific business rules must not be inferred from directory names alone.

## Expected Closeout

After approved deployment, standard mode stops and outputs next-thread prompts.

It does not continue optimizing the reporting project.
