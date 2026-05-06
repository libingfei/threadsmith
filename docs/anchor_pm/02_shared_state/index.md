# Shared State Index

Layer: `2`

Purpose: list active cross-thread shared-state files.

## Project-Level Shared Mirrors

| File | Mirrors | Reason |
| --- | --- | --- |
| `docs/anchor_pm/02_shared_state/migration_status.md` | Layer 2 migration state | Shared migration tracking |
| `docs/anchor_pm/02_shared_state/project_version_notice.md` | `docs/anchor_pm/current_version.md` | Layer 0 / Layer 2 split preparation |
| `docs/anchor_pm/02_shared_state/interaction_workflow.md` | `docs/anchor_pm/interaction_guide.md` | Workflow notice migration |
| `docs/anchor_pm/02_shared_state/review_log.md` | `docs/anchor_pm/review_log.md` | Repeated mechanism issue migration |
| `docs/anchor_pm/02_shared_state/simplification_log.md` | `docs/anchor_pm/simplification.md` | Document-growth governance migration |

## Active Directed Channels

| Source | Target | File | Reason |
| --- | --- | --- | --- |
| Product Manager | Coordination | `docs/anchor_pm/02_shared_state/product_manager__to_coordination.md` | Layer 1 migration promotion decision |
| Product Manager | Reanchor Detector Core | `docs/anchor_pm/02_shared_state/product_manager__to_reanchor_detector_core.md` | Reanchor model product requirements |
| Product Manager | CLI Core | `docs/anchor_pm/02_shared_state/product_manager__to_cli_core.md` | Programmatic reanchor runtime entrypoint |
| Product Manager | Templates / Protocol | `docs/anchor_pm/02_shared_state/product_manager__to_templates_protocol.md` | Automatic Reanchor Start workflow wording |
| Product Manager | Codex Skill / Package Installer | `docs/anchor_pm/02_shared_state/product_manager__to_codex_skill.md` | Automatic Reanchor Start install guidance |

## Detector Use

Ordinary threads should confirm this layer at startup. They should read only
inbound files addressed to their thread or files explicitly named in a handoff.
