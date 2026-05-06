# Reanchor Detector Core History

Layer: `3`

Thread: `Reanchor Detector Core`

Status: compatibility mirror of `docs/module_state/reanchor_detector_core.md`.

- Implemented initial Go detector core and fixture-style tests covering missing
  checkpoint, unchanged state, Layer 0 / 1 / 2 / 3 changes, conservative error
  handling, periodic reanchor, closeout required updates, backslash paths, and
  symlink traversal.

- Created after reclassifying the detector from an internal module to a
  dedicated core subsystem thread.
