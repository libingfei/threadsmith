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

Report:

- `Observed`: direct file facts.
- `Inference`: likely thread boundaries or project rules.
- `Needs Confirmation`: anything that would affect formal contracts.

## Proposed Integration

Prefer creating Anchor PM docs without replacing existing project rules.

If `AGENTS.md` already exists, propose one of:

- keep existing `AGENTS.md` and add a short Anchor PM section;
- create a draft merge plan;
- leave `AGENTS.md` untouched and install only `docs/anchor_pm/` files.

## Stop Conditions

Stop and ask for approval when:

- an existing rule file would be changed;
- a thread boundary is unclear;
- inferred project facts would enter contracts;
- the target project appears outside the requested root.
