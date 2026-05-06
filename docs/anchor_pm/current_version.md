# Anchor PM Current Version

Current coordination version: `package-first-v1.0`

## Status

Anchor PM is now self-hosting at the coordination-document and package level.

This means this repository uses its own minimal Anchor PM structure and now contains an Anchor PM 1.0 package-first release.

## Version Notes

- Added project-level `AGENTS.md`.
- Added thread contracts in `docs/anchor_pm/contracts.md`.
- Added initial module state files under `docs/module_state/`.
- Established self-dogfood as a product constraint.
- Added `packages/anchor-pm-1.0/` canonical package source.
- Added `packages/anchor-pm-1.0-standard/` and `packages/anchor-pm-1.0-self-evolution/` release directories.
- Repositioned CLI as a post-1.0 reproducibility/checking tool.
- Added `docs/anchor_pm/thread_initialization.md` so users can create threads without inventing prompts.
- Added `docs/anchor_pm/thread_management_install_prompt.md` as the copy-paste prompt for integrating Anchor PM into a target project.
- Added `Product Manager` thread for user operation flows and experience optimization.
- Added `docs/anchor_pm/internal_function_spec.md` as the internal module input/output reference for development and prompt engineering.
- Added `Reanchor Detector Core` thread for Contract Version / Reanchor State Detector reliability and implementation.
- Added `docs/anchor_pm/contract_state_detector.md` as the focused design note for anchor version/change detection.
- Initialized Layer 0 through Layer 3 structure:
  - `docs/anchor_pm/00_framework_baseline/`
  - `docs/anchor_pm/01_thread_definitions/`
  - `docs/anchor_pm/02_shared_state/`
  - `docs/anchor_pm/03_thread_local_memory.md`
  - `docs/module_state/product_manager/` as the first category-level Layer 3
    pilot.
- Internal Layer directories use numbered prefixes for scanability and avoid
  `README.md` files unless the location is intentionally user-facing.
- Mirrored `contracts.md` and `thread_initialization.md` into complete
  per-thread Layer 1 definition files under `docs/anchor_pm/01_thread_definitions/`.
- Added a Product Manager to Coordination shared-state handoff for deciding
  whether the Layer 1 split should become authoritative.
- Mirrored Layer 2 shared files into `docs/anchor_pm/02_shared_state/`:
  project version notice, interaction workflow, review log, simplification log,
  and Layer 2 migration status.
- Mirrored Layer 3 `docs/module_state/*.md` files into category-level files
  under `docs/module_state/<thread>/`.
- Added `docs/anchor_pm/00_framework_baseline/reanchor_module_io_spec.md` as
  the fixed programmatic input/output contract for Reanchor Detector Core
  development.
- Added automatic `Reanchor Start` protocol: Codex should trigger reanchor
  before substantial thread work, use a programmatic detector when available,
  fall back to reading required anchors only as a degraded compatibility path,
  and avoid asking users to run CLI commands manually.
- Updated public install prompts to use
  `https://github.com/libingfei/threadsmith` and package directory
  `packages/anchor-pm-1.0-standard` instead of a local development package path.
- Promoted Closeout Knowledge Sync as a core all-thread workflow in product
  state, interaction guidance, internal function spec, detector product docs,
  and current project thread initialization prompts.
- Defined Reanchor Start and Closeout Knowledge Sync as symmetric lifecycle
  hooks: read/refresh/confirm before work, then write/handoff durable knowledge
  before reply.
- Added `docs/anchor_pm/mvp_manual_test_protocol.md` so the full MVP manual
  validation flow is available from the public repository.
- Tightened install proposal UX: user-visible thread names follow the selected
  install prompt language, and `Adjust AGENTS.md` is no longer a default reply
  option.

## Reanchor Requirement

For substantial work, run `Reanchor Start` automatically. The intended behavior
is programmatic anchoring: Codex calls a detector command/tool, receives a
machine-readable refresh decision, and reads only the files named in
`required_reads`.

1. Use the Reanchor Detector command/tool when available.
2. Follow its `required_reads`, `blocked_by`, and `next_action`.
3. If the detector is unavailable, report the degraded state, then read:
   - `AGENTS.md`
   - `docs/anchor_pm/current_version.md`
   - `docs/anchor_pm/contracts.md`
   - the relevant `docs/module_state/<thread>.md`
4. Do not treat this fallback as the target product behavior or as proof that
   programmatic anchoring is complete.
5. Show a minimal anchor-state line to the user and continue only within the
   current thread scope.

If this version changes, reread all four before continuing.
