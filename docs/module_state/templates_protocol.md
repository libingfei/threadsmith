# Templates / Protocol Thread State

## Contract

Thread contract: `Templates / Protocol` in `docs/anchor_pm/contracts.md`

## Current State

- Template requirements are described in `MVP_SPEC.md`.
- Concrete 1.0 template files exist under `packages/anchor-pm-1.0/templates/`.
- Required template families are `AGENTS`, `current_version`, `contracts`, `thread_initialization`, `interaction_guide`, `module_state`, `review_log`, and `simplification`.
- Package workflows and checklists exist under `packages/anchor-pm-1.0/workflows/` and `packages/anchor-pm-1.0/checklists/`.
- Reanchor detector prompt behavior should align with Reanchor Detector Core's code-level input/output contract.

## Open Issues

- Keep default templates project-neutral.
- Keep standard and self-evolution release directories structurally identical except `ACTIVE_INSTALL_PLAN.md` and `INSTALL_PROMPT.md`.
- Ensure existing-project adoption templates distinguish confirmed facts from inferred suggestions.
- Ensure generated thread initialization prompts are complete and do not leave user-filled placeholders.
- Decide whether future versions should generate zip artifacts or keep directory packages only.
- Keep prompt/workflow detector wording synchronized with Reanchor Detector Core behavior.

## Runbook

Expected next steps:

1. Review package templates against product principles.
2. Shorten any template that creates unnecessary rule surface.
3. Feed dry-run findings into package revisions only after Coordination review.
4. Hand future CLI reproducibility concerns to CLI Core.

## History / Notes

- Created as a planned protocol thread during self-dogfood adoption.
