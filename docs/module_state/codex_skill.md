# Codex Skill / Package Installer Thread State

## Contract

Thread contract: `Codex Skill / Package Installer` in `docs/anchor_pm/contracts.md`

## Current State

- No formal Codex Skill file exists yet.
- Anchor PM 1.0 uses `INSTALL_PROMPT.md` as the Codex-first entrypoint.
- The installer should read package docs and target Markdown anchors, not store project-specific rules.

## Open Issues

- Test `INSTALL_PROMPT.md` in Codex App on a real target project.
- Decide whether a future Skill should wrap the package or remain unnecessary.
- Define how a future Skill should present reanchor and handoff outputs.
- Verify installer replies follow the user's usual conversation language while generated project docs may remain English.

## Runbook

Expected next steps:

1. Use `packages/anchor-pm-1.0-standard/INSTALL_PROMPT.md` in a sample target project.
2. Use `packages/anchor-pm-1.0-self-evolution/INSTALL_PROMPT.md` on this repository.
3. Record friction before deciding whether to add a formal Skill.

## History / Notes

- Created as a planned Codex adapter thread during self-dogfood adoption.
