# Product Manager Runbook

Layer: `3`

Thread: `Product Manager`

Before substantial product-flow work:

1. Run Anchor Gate silently unless changed, blocked, unknown, conflicting, or
   degraded.
2. Inside the gate, classify explicit durable user corrections or update
   requests before Reanchor Start; write/stage only safe owned changes.
3. Use the detector if available.
4. If unavailable, report the degraded state, then read `AGENTS.md`,
   `docs/anchor_pm/current_version.md`,
   `docs/anchor_pm/contracts.md`, and `docs/module_state/product_manager.md`.
5. Keep anchor handling below the task budget; do not explain unchanged gates.

For reanchor-model work:

1. Keep the model user-comprehensible.
2. Confirm all layers can be checked without forcing all contents into context.
3. Prefer sparse dependency files over full cross-thread matrices.
4. Treat automatic full rereading as degraded fallback, not target behavior.
5. Handoff implementation behavior to the owning thread.

PM Review Gate before declaring an install or workflow acceptable:

1. Confirm the flow matches real Codex client behavior; do not depend on steps
   the user cannot perform before the first message.
2. Judge whether the output helps the target-project user decide the next step,
   not whether it exposes enough Anchor PM internals.
3. Verify proposed threads come from the target project's modules/subsystems or
   durable maintenance boundaries, not from generic role buckets.
4. Check whether each visible option has a clear, useful outcome. Remove or flag
   options that create confusing partial integration.
5. Try the target maintainer perspective: if the proposed threads would not
   reduce context scope or repeated explanation, mark the result as retry or
   blocked even if every checklist field exists.
6. If this review changes product behavior, update thread state before final
   response; otherwise keep Knowledge Sync Gate silent unless status was
   requested.

Pre-test GitHub Sync Gate:

1. When the user says they are about to test public GitHub behavior, check
   `git status --short --branch`.
2. If local intended changes exist, commit and push them before the user starts
   the test.
3. Verify `git ls-remote origin refs/heads/main` matches local `HEAD`.
4. Give the user the commit SHA and direct GitHub links for the files they will
   paste or follow.
5. If changes should not be pushed, clearly tell the user to test with local
   files instead of GitHub links.
