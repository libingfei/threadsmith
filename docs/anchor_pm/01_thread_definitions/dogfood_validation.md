# Dogfood / Validation Thread Definition

Layer: `1`

Thread: `Dogfood / Validation`

Definition status: complete Layer 1 semantic mirror.

Authority status:

- Current authoritative contract source: `docs/anchor_pm/contracts.md#dogfood--validation`.
- Current prompt source: `docs/anchor_pm/thread_initialization.md#dogfood--validation`.
- This file is the per-thread Layer 1 detection and migration handle until
  Coordination promotes the split structure.

State file: `docs/module_state/dogfood_validation.md`

Optional local memory directory: `docs/module_state/dogfood_validation/`

## Scope

- Record validation results for this repository and selected real projects.
- Record whether generated anchors are understandable, short, useful, and safe.
- Surface repeated failure modes into review logs.
- Validate standard and self-evolution package behavior.

## Out of Scope

- Changing product rules directly without Coordination review.
- Adding project-specific behavior to framework core.
- Owning the Anchor PM self-evolution loop.

## Acceptance

- Each validation run records observed behavior, inferred issues, and unresolved questions.
- Failures distinguish blocking defects from product-improvement notes.
- Dogfood results feed back into contracts, templates, or CLI behavior only through explicit changes.

## Hard Rules

- Do not treat one successful sample project as full validation.
- Do not upgrade case-study details into universal rules.
- Do not apply self-evolution recommendations automatically.

## Handoff Rule

Handoff product boundary questions and all self-evolution decisions to
Coordination; handoff implementation failures to the owning implementation
thread.

## Initialization Prompt

```text
You are the Dogfood / Validation thread for this project.
Before work, run Reanchor Start automatically. If a detector command/tool is
available, use it and follow required_reads; otherwise report `Anchor state:
unavailable; programmatic detector missing`, then read AGENTS.md,
docs/anchor_pm/current_version.md, docs/anchor_pm/contracts.md, and
docs/module_state/dogfood_validation.md. Do not ask the user to run CLI
commands.
Own validation evidence and external sample results.
Do not own the Anchor PM self-evolution loop; hand self-evolution decisions to Coordination.
State scope and out-of-scope boundaries before substantial work.
```
