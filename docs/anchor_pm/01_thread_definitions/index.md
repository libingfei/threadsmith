# Thread Definition Index

Layer: `1`

Current thread count: `7`

Migration status: `semantic mirror initialized`

Canonical contract source: `docs/anchor_pm/contracts.md`

Thread prompt source: `docs/anchor_pm/thread_initialization.md`

Migration tracker: `docs/anchor_pm/01_thread_definitions/migration_status.md`

## Threads

| Thread | Definition File | State File |
| --- | --- | --- |
| Coordination | `docs/anchor_pm/01_thread_definitions/coordination.md` | `docs/module_state/coordination.md` |
| Product Manager | `docs/anchor_pm/01_thread_definitions/product_manager.md` | `docs/module_state/product_manager.md` |
| Reanchor Detector Core | `docs/anchor_pm/01_thread_definitions/reanchor_detector_core.md` | `docs/module_state/reanchor_detector_core.md` |
| CLI Core | `docs/anchor_pm/01_thread_definitions/cli_core.md` | `docs/module_state/cli_core.md` |
| Templates / Protocol | `docs/anchor_pm/01_thread_definitions/templates_protocol.md` | `docs/module_state/templates_protocol.md` |
| Codex Skill / Package Installer | `docs/anchor_pm/01_thread_definitions/codex_skill.md` | `docs/module_state/codex_skill.md` |
| Dogfood / Validation | `docs/anchor_pm/01_thread_definitions/dogfood_validation.md` | `docs/module_state/dogfood_validation.md` |

## Detector Use

Ordinary threads should confirm this layer at startup. If their own definition
changed, they should refresh that definition before work. If the thread topology
changed, they should hand off to Thread Management or Coordination.

## Compatibility

The per-thread files now contain complete Layer 1 contract and prompt content.
They are still compatibility mirrors until Coordination promotes this split as
authoritative.
