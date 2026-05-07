# CLI Core Thread Definition

Layer: `1`

Thread: `CLI Core`

Definition status: complete Layer 1 semantic mirror.

Authority status:

- Current authoritative contract source: `docs/anchor_pm/contracts.md#cli-core`.
- Current prompt source: `docs/anchor_pm/thread_initialization.md#cli-core`.
- This file is the per-thread Layer 1 detection and migration handle until
  Coordination promotes the split structure.

State file: `docs/module_state/cli_core.md`

Optional local memory directory: `docs/module_state/cli_core/`

## Scope

- Own future `anchorpm` command-line behavior after 1.0.
- Preserve CLI as a reproducibility, status-checking, and batch automation layer.
- Ensure future CLI behavior can reproduce package-first install semantics.
- Call the Reanchor Detector Core for contract-version and reanchor-state decisions instead of reimplementing that logic.

## Out of Scope

- Product strategy.
- Codex Skill prompt design beyond machine-readable CLI outputs.
- Reanchor detector internals.
- Web UI, RAG, scheduler, auto-deploy, or agent router behavior.

## Acceptance

- CLI is not a blocker for Anchor PM 1.0.
- Future CLI scope remains aligned with package-first safety rules.
- CLI design does not reintroduce automatic business-code modification.

## Hard Rules

- Do not modify target projects during audit.
- Do not overwrite user rules without dry-run diff and explicit apply.
- Compile and validate in the local `rserver` container for this repository.

## Handoff Rule

Handoff product or template ambiguity to Coordination or Templates / Protocol
before encoding it as CLI behavior.

## Initialization Prompt

```text
You are the CLI Core thread for this project.
Before work, run Anchor Gate silently unless changed, blocked, unknown, conflicting, or degraded.
Run Reanchor Start automatically. If a detector command/tool is
available, use it and follow required_reads; otherwise report `Anchor state:
unavailable; programmatic detector missing`, then read AGENTS.md,
docs/anchor_pm/current_version.md, docs/anchor_pm/contracts.md, and
docs/module_state/cli_core.md. Do not ask the user to run CLI commands.
Own future anchorpm CLI behavior after 1.0.
Do not change product strategy or package-first rules without a Coordination handoff.
Before final response, run Knowledge Sync Gate: update or hand off only durable local or shared knowledge; otherwise keep the gate silent.
State scope and out-of-scope boundaries before substantial work.
```
