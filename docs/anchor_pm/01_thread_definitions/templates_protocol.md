# Templates / Protocol Thread Definition

Layer: `1`

Thread: `Templates / Protocol`

Definition status: complete Layer 1 semantic mirror.

Authority status:

- Current authoritative contract source: `docs/anchor_pm/contracts.md#templates--protocol`.
- Current prompt source: `docs/anchor_pm/thread_initialization.md#templates--protocol`.
- This file is the per-thread Layer 1 detection and migration handle until
  Coordination promotes the split structure.

State file: `docs/module_state/templates_protocol.md`

Optional local memory directory: `docs/module_state/templates_protocol/`

## Scope

- Own default Markdown templates, package workflow text, and protocol text.
- Keep templates short, auditable, and project-neutral.
- Define required sections for `AGENTS.md`, `contracts.md`, `current_version.md`, `interaction_guide.md`, and `module_state/<thread>.md`.
- Maintain package-first install plans, checklists, and workflow documents.

## Out of Scope

- CLI parsing and filesystem behavior.
- Codex App behavior outside the package instructions.
- Case-study-specific logic.

## Acceptance

- Templates and workflows are small enough for users to review quickly.
- Templates distinguish direct observation from inference.
- Templates do not assume a specific business domain.
- Standard and self-evolution package structures remain identical except `ACTIVE_INSTALL_PLAN.md` and `INSTALL_PROMPT.md`.

## Hard Rules

- Do not create a second project brain.
- Do not copy source-project-specific terms into framework core.
- Prefer one reusable section over multiple near-duplicate documents.

## Handoff Rule

Handoff implementation concerns to CLI Core; handoff interaction wording
concerns to Codex Skill.

## Initialization Prompt

```text
You are the Templates / Protocol thread for this project.
Before work, run Anchor Gate silently unless changed, blocked, unknown, conflicting, or degraded.
Run Reanchor Start automatically. If a detector command/tool is
available, use it and follow required_reads; otherwise report `Anchor state:
unavailable; programmatic detector missing`, then read AGENTS.md,
docs/anchor_pm/current_version.md, docs/anchor_pm/contracts.md, and
docs/module_state/templates_protocol.md. Do not ask the user to run CLI
commands.
Own package templates, workflow text, checklists, and protocol wording.
Keep documents short, project-neutral, and auditable.
Before final response, run Knowledge Sync Gate: update or hand off only durable local or shared knowledge; otherwise keep the gate silent.
State scope and out-of-scope boundaries before substantial work.
```
