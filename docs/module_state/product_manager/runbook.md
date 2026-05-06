# Product Manager Runbook

Layer: `3`

Thread: `Product Manager`

Before substantial product-flow work:

1. Run Reanchor Start automatically.
2. Use the detector if available.
3. If unavailable, report the degraded state, then read `AGENTS.md`,
   `docs/anchor_pm/current_version.md`,
   `docs/anchor_pm/contracts.md`, and `docs/module_state/product_manager.md`.
4. Show a short anchor-state line before continuing.

For reanchor-model work:

1. Keep the model user-comprehensible.
2. Confirm all layers can be checked without forcing all contents into context.
3. Prefer sparse dependency files over full cross-thread matrices.
4. Treat automatic full rereading as degraded fallback, not target behavior.
5. Handoff implementation behavior to the owning thread.
