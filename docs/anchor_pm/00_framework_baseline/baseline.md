# Anchor PM Framework Baseline

Layer: `0`

Baseline id: `package-first-v1.0`

Package model: Markdown package-first.

Current coordination version source:

- `docs/anchor_pm/current_version.md`

Layer 2 shared version mirror:

- `docs/anchor_pm/02_shared_state/project_version_notice.md`

Primary package roots:

- `packages/anchor-pm-1.0/`
- `packages/anchor-pm-1.0-standard/`
- `packages/anchor-pm-1.0-self-evolution/`

Protocol references:

- `docs/anchor_pm/00_framework_baseline/reanchor_module_io_spec.md`
- `docs/anchor_pm/contract_state_detector.md`
- `docs/anchor_pm/internal_function_spec.md`

Refresh trigger:

- Anchor PM package, template, protocol, or detector behavior changes.

Ordinary thread behavior:

- Trigger Reanchor Start automatically before substantial work.
- Confirm this baseline id through the detector when available, or through the
  documented degraded fallback when unavailable.
- If it changed, stop ordinary work and hand off to Thread Management or the
  framework-upgrade flow.

Compatibility status:

- This file mirrors the Layer 0 meaning of `docs/anchor_pm/current_version.md`.
- `docs/anchor_pm/current_version.md` remains authoritative until Coordination
  promotes the Layer 0 / Layer 2 split.
