# Coordination Runbook

Layer: `3`

Thread: `Coordination`

Status: compatibility mirror of `docs/module_state/coordination.md`.

Before substantial coordination work:

1. Run Reanchor Start automatically.
2. Use the detector if available.
3. If unavailable, read `AGENTS.md`, `docs/anchor_pm/current_version.md`,
   `docs/anchor_pm/contracts.md`, and `docs/module_state/coordination.md`.
4. Show a short anchor-state line before continuing.

When a request belongs to another thread:

1. State the boundary.
2. Produce a handoff summary.
3. Update this file only if long-term coordination state changed.

When running self-evolution:

1. Use `packages/anchor-pm-1.0-self-evolution/`.
2. Generate one self-optimization report.
3. Separate `Observed`, `Inference`, and `Unverified`.
4. Convert accepted recommendations into explicit follow-up work.
5. Do not auto-apply recommendations within the same step unless the user
   separately asks for implementation.

When making strong product conclusions:

1. Separate `Observed`, `Inference`, and `Unverified`.
2. Avoid closure wording unless a formal check or validation run exists.
