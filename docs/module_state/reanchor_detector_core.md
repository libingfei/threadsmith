# Reanchor Detector Core Thread State

## Contract

Thread contract: `Reanchor Detector Core` in `docs/anchor_pm/contracts.md`

## Current State

- `Contract Version / Reanchor State Detector` is defined in `docs/anchor_pm/internal_function_spec.md`.
- The detector is considered a core code subsystem because its reliability, input/output contract, and efficiency directly affect Anchor PM feasibility.
- Initial detector code now exists as a Go core package under
  `pkg/reanchor/`, with module root `go.mod`.
- The package exposes `Evaluate(ctx, ReanchorRequest) (ReanchorResult, error)`
  and implements fixed structured request/result types, default Layer 0 through
  Layer 3 registry discovery, checkpoint comparison, conservative file status
  handling, periodic reanchor detection, required reads, checkpoint update
  proposals, minimal chat output, and closeout required-update planning.
- Future CLI and package workflows should call or mirror this detector instead of duplicating version-state logic.

## Open Issues

- Run authoritative validation in the local `rserver` container once WSL Docker
  integration or another rserver execution path is available.
- Run Windows local validation once a Windows Go toolchain is available through
  the active execution environment.
- Coordination still needs to decide when Layer 1 / Layer 2 / Layer 3 mirrors
  become authoritative instead of compatibility mirrors.
- CLI Core still needs to decide future command packaging without duplicating
  detector semantics.

## Runbook

Before substantial detector work:

1. Run Reanchor Start automatically.
2. Use the detector if available.
3. If unavailable, read `AGENTS.md`, `docs/anchor_pm/current_version.md`,
   `docs/anchor_pm/contracts.md`, and this file.
4. Read `docs/anchor_pm/internal_function_spec.md`, section
   `Contract Version / Reanchor State Detector`.
5. Show a short anchor-state line before continuing.

When designing detector behavior:

1. Specify inputs.
2. Specify outputs.
3. Specify failure behavior.
4. Specify test fixtures.
5. Check that unknown state forces refresh rather than stale continuation.

## History / Notes

- Created after reclassifying the detector from an internal module to a dedicated core subsystem thread.
- Implemented initial Go detector core and fixture-style tests covering missing
  checkpoint, unchanged state, Layer 0 / 1 / 2 / 3 changes, conservative error
  handling, periodic reanchor, closeout required updates, backslash paths, and
  symlink traversal.
