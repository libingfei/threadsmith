# Layer 2 Migration Status

Purpose: track whether shared workflow and governance files can eventually be
reduced to compatibility redirects.

## Current Status

Status: `semantic mirrors initialized`

Layer 2 now has mirrors for:

- `docs/anchor_pm/current_version.md`
  - mirror: `docs/anchor_pm/02_shared_state/project_version_notice.md`
- `docs/anchor_pm/interaction_guide.md`
  - mirror: `docs/anchor_pm/02_shared_state/interaction_workflow.md`
- `docs/anchor_pm/review_log.md`
  - mirror: `docs/anchor_pm/02_shared_state/review_log.md`
- `docs/anchor_pm/simplification.md`
  - mirror: `docs/anchor_pm/02_shared_state/simplification_log.md`

Layer 2 also has directed dependency handoff files for cross-thread changes:

- `docs/anchor_pm/02_shared_state/product_manager__to_coordination.md`
- `docs/anchor_pm/02_shared_state/product_manager__to_reanchor_detector_core.md`
- `docs/anchor_pm/02_shared_state/product_manager__to_cli_core.md`
- `docs/anchor_pm/02_shared_state/product_manager__to_templates_protocol.md`
- `docs/anchor_pm/02_shared_state/product_manager__to_codex_skill.md`

## Still Required Before Deleting Old Files

- Coordination must decide which Layer 2 mirrors become authoritative.
- Templates / Protocol must confirm package and workflow wording can point to
  the Layer 2 files.
- `docs/anchor_pm/current_version.md` must be split safely between Layer 0
  baseline and Layer 2 shared notice responsibilities.
- Existing install prompts and generated user docs must stop depending on old
  paths or the old paths must remain compatibility redirects.
- A validation pass must show that existing threads can reanchor without losing
  workflow, version, review, or simplification information.

## Do Not Delete Yet

- `docs/anchor_pm/current_version.md`
- `docs/anchor_pm/interaction_guide.md`
- `docs/anchor_pm/review_log.md`
- `docs/anchor_pm/simplification.md`
