# Workflow: Existing Project Adoption

Use this workflow when the target project already has files, docs, rules, source code, or project history.

## Audit

Inspect likely sources of truth:

- `README*`
- `AGENTS.md`
- existing docs
- scripts
- config files
- tests
- CI or deploy files

Do not modify files during audit.

## Classify Findings

Keep findings internally while preparing the proposal:

- `Observed`: direct file facts.
- `Inference`: likely module/subsystem boundaries or project rules.
- `Needs Confirmation`: anything that would affect formal contracts.

Do not show these as default proposal sections. Put them in an optional details
block only when they affect approval, or write them after approval to
`docs/anchor_pm/install_decision_record.md`.

## Proposed Integration

Prefer creating Anchor PM docs without replacing existing project rules.

For existing projects, propose project specialist threads from real modules,
subsystems, runtime surfaces, documentation/support surfaces, and durable
maintenance boundaries. Do not use generic `Coordination / Implementation /
Validation` as the default thread set.

If `AGENTS.md` already exists, inspect it for conflicts.

- If there is no clear conflict, plan a short Anchor PM discovery section.
- If there is a conflict or uncertainty, pause for a specific merge decision.
- Do not expose docs-only/no-AGENTS as a default reply option.
- Do not show `AGENTS.md` handling in the main proposal unless it creates an
  approval-blocking risk.

## Stop Conditions

Stop and ask for approval when:

- an existing rule file would be changed;
- a thread boundary is unclear;
- inferred project facts would enter contracts;
- the target project appears outside the requested root.
