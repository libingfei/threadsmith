# Project Version Notice

Layer: `2`

Status: compatibility mirror.

Current coordination version: `package-first-v1.0`

Authoritative source until promoted:

- `docs/anchor_pm/current_version.md`

## Current Status

Anchor PM is self-hosting at the coordination-document and package level.

This repository uses its own minimal Anchor PM structure and contains an Anchor
PM 1.0 package-first release.

## Shared Change Notes

- Layer 0 through Layer 3 structure has been initialized.
- Internal Layer directories use numbered prefixes.
- Internal Layer directories avoid `README.md` files unless intentionally
  user-facing.
- Layer 1 contracts and initialization prompts have been mirrored into complete
  per-thread definition files under `docs/anchor_pm/01_thread_definitions/`.
- Product Manager created a Layer 2 handoff to Coordination for deciding whether
  the Layer 1 split should become authoritative.
- Automatic Reanchor Start is now part of the project protocol: Codex triggers
  reanchor before substantial work, uses the detector when available, falls back
  to reading required anchors only as a degraded compatibility path when
  unavailable, and does not ask users to run CLI commands manually.
- Product Manager created Layer 2 handoffs to Reanchor Detector Core, CLI Core,
  Templates / Protocol, and Codex Skill / Package Installer for turning
  Reanchor Start into programmatic anchoring with degraded fallback only when no
  detector entrypoint is callable.

## Reanchor Meaning

If this notice changes, ordinary threads should confirm whether the change
affects their Layer 1 thread definition, inbound Layer 2 dependencies, or Layer
3 local memory.
