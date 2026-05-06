# Reanchor Detector Core Runbook

Layer: `3`

Thread: `Reanchor Detector Core`

Status: compatibility mirror of `docs/module_state/reanchor_detector_core.md`.

Before substantial detector work:

1. Run Reanchor Start automatically.
2. Use the detector if available.
3. If unavailable, read `AGENTS.md`, `docs/anchor_pm/current_version.md`,
   `docs/anchor_pm/contracts.md`, and `docs/module_state/reanchor_detector_core.md`.
4. Read `docs/anchor_pm/internal_function_spec.md`, section
   `Contract Version / Reanchor State Detector`.
5. Show a short anchor-state line before continuing.

When designing detector behavior:

1. Specify inputs.
2. Specify outputs.
3. Specify failure behavior.
4. Specify test fixtures.
5. Check that unknown state forces refresh rather than stale continuation.
