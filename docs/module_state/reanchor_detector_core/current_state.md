# Reanchor Detector Core Current State

Layer: `3`

Thread: `Reanchor Detector Core`

Status: compatibility mirror of `docs/module_state/reanchor_detector_core.md`.

- `Contract Version / Reanchor State Detector` is defined in
  `docs/anchor_pm/internal_function_spec.md`.
- The detector is considered a core code subsystem because its reliability,
  input/output contract, and efficiency directly affect Anchor PM feasibility.
- Initial detector code now exists as a Go core package under
  `pkg/reanchor/`, with module root `go.mod`.
- The package exposes `Evaluate(ctx, ReanchorRequest) (ReanchorResult, error)`
  and implements fixed structured request/result types, default Layer 0 through
  Layer 3 registry discovery, checkpoint comparison, conservative file status
  handling, periodic reanchor detection, required reads, checkpoint update
  proposals, minimal chat output, and closeout required-update planning.
- Future CLI and package workflows should call or mirror this detector instead
  of duplicating version-state logic.
