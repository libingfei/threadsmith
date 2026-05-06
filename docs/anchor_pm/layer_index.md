# Anchor PM Layer Index

Purpose: show the initialized Layer 0 through Layer 3 structure for this
repository.

This index does not replace the current authoritative files listed in
`AGENTS.md`. It gives the reanchor detector a clear directory model to evolve
toward while Anchor PM keeps its 1.0 package-first documents stable.

## Layer Map

```text
Layer 0: Framework Baseline
  docs/anchor_pm/00_framework_baseline/

Layer 1: Thread Definition
  docs/anchor_pm/01_thread_definitions/
  status: semantic mirror initialized, not yet authoritative

Layer 2: Cross-Thread Shared State
  docs/anchor_pm/02_shared_state/
  status: semantic mirrors initialized, not yet authoritative

Layer 3: Thread Local Memory
  docs/anchor_pm/03_thread_local_memory.md
  docs/module_state/
  docs/module_state/<thread>/, created only when category split is needed
  status: semantic mirrors initialized, not yet authoritative
```

## Current Compatibility Rule

- `AGENTS.md`, `docs/anchor_pm/current_version.md`,
  `docs/anchor_pm/contracts.md`, and `docs/module_state/*.md` remain the current
  authoritative coordination files.
- The new layer directories provide smaller detection handles, indexes, and
  split points.
- Do not migrate authority from the current files into the layer directories
  without a Coordination decision.
- `README.md` files should be reserved for user-facing entrypoints. Internal
  anchor directories should use `index.md`, named spec files, or category files
  instead.
