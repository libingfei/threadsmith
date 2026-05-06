# Layer 1 Migration Status

Purpose: track whether `docs/anchor_pm/contracts.md` and
`docs/anchor_pm/thread_initialization.md` can eventually be reduced to
compatibility redirects.

## Current Status

Status: `semantic mirror initialized`

The per-thread files under `docs/anchor_pm/01_thread_definitions/` now contain:

- full scope;
- full out-of-scope boundaries;
- acceptance criteria;
- hard rules;
- state file path;
- handoff rule;
- initialization prompt.

## Still Required Before Deleting Old Files

- Coordination must decide that Layer 1 split files become authoritative.
- `AGENTS.md` must point to the Layer 1 index and current thread definition.
- `docs/anchor_pm/thread_initialization.md` must be replaced by a compatibility
  redirect or generated from `01_thread_definitions/*.md`.
- `docs/anchor_pm/contracts.md` must be replaced by a compatibility redirect or
  generated from `01_thread_definitions/*.md`.
- Package templates and install prompts must stop depending on the old Layer 1
  source files.
- A validation pass must show that every thread can reanchor from the new Layer
  1 files without losing scope, boundary, handoff, or prompt information.

## Do Not Delete Yet

- `docs/anchor_pm/contracts.md`
- `docs/anchor_pm/thread_initialization.md`
