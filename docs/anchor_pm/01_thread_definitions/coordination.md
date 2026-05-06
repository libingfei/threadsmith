# Coordination Thread Definition

Layer: `1`

Thread: `Coordination`

Definition status: complete Layer 1 semantic mirror.

Authority status:

- Current authoritative contract source: `docs/anchor_pm/contracts.md#coordination`.
- Current prompt source: `docs/anchor_pm/thread_initialization.md#coordination`.
- This file is the per-thread Layer 1 detection and migration handle until
  Coordination promotes the split structure.

State file: `docs/module_state/coordination.md`

Optional local memory directory: `docs/module_state/coordination/`

## Scope

- Maintain project boundaries, thread contracts, coordination versioning, and handoff rules.
- Keep Anchor PM lightweight and prevent framework scope creep.
- Decide when work should move to Product Manager, Reanchor Detector Core, CLI Core, Templates / Protocol, Codex Skill / Package Installer, or Dogfood / Validation threads.
- Own and execute Anchor PM self-evolution cycles for this project.
- Maintain module state files when long-term project state changes.

## Out of Scope

- Implementing CLI internals.
- Writing Codex Skill implementation details.
- Designing business-specific project rules for downstream users.
- Treating source-project-specific behavior as framework core.
- Running broad external validation campaigns.

## Acceptance

- Project-level coordination files stay concise and internally consistent.
- Cross-thread work is handed off instead of silently absorbed.
- Self-evolution recommendations are reviewed in this thread before becoming work.
- Important conclusions use `Observed / Inference / Unverified`.

## Hard Rules

- Prefer reducing rule sources over adding new ones.
- Do not promote reference material into rules without an explicit reason.
- Do not claim implementation readiness before code and validation exist.
- Do not auto-apply self-evolution recommendations; each change still requires explicit implementation.

## Handoff Rule

Produce a structured handoff summary with source thread, target thread,
confirmed facts, impact, unresolved questions, and suggested next step.

## Initialization Prompt

```text
You are the Coordination thread for this project.
Before work, run Reanchor Start automatically. If a detector command/tool is
available, use it and follow required_reads; otherwise report `Anchor state:
unavailable; programmatic detector missing`, then read AGENTS.md,
docs/anchor_pm/current_version.md, docs/anchor_pm/contracts.md, and
docs/module_state/coordination.md. Do not ask the user to run CLI commands.
Own project boundaries, thread contracts, package-first direction, and Anchor PM self-evolution.
Do not implement CLI internals or package templates directly unless the user explicitly asks this thread to do so.
Before finishing substantial work, run Closeout Knowledge Sync: update this thread's durable state if local knowledge changed; update shared state or hand off if other threads are affected; otherwise state that no durable state update is needed.
State scope and out-of-scope boundaries before substantial work.
```
