# Anchor PM 1.0 Package Manifest

Package name: `anchor-pm`

Package version: `1.0`

Package protocol: `codex-package-first`

Active mode source: `ACTIVE_INSTALL_PLAN.md`

## Purpose

This package installs Anchor PM coordination anchors into a Codex project.

It does not install a CLI, run deployment commands, modify business code, or act as an autonomous project manager.

## Modes

The package supports two release modes with the same directory structure:

- `standard`: install Anchor PM anchors into an ordinary project, then stop.
- `self-evolution`: install or refresh Anchor PM anchors in the Anchor PM project, then generate one optimization report.

The only release-level differences are:

- `ACTIVE_INSTALL_PLAN.md`
- `INSTALL_PROMPT.md`

`INSTALL_PROMPT.md` may differ because development builds include a concrete local package path.

## Required Package Files

- `INSTALL_PROMPT.md`
- `PACKAGE_MANIFEST.md`
- `ACTIVE_INSTALL_PLAN.md`
- `plans/standard_project_install.md`
- `plans/self_evolution_install.md`
- `templates/thread_initialization.template.md`
- `templates/*.template.md`
- `workflows/*.md`
- `checklists/*.md`

## Target Outputs

The installer may propose these target-project files:

- `AGENTS.md`
- `docs/anchor_pm/current_version.md`
- `docs/anchor_pm/contracts.md`
- `docs/anchor_pm/thread_initialization.md`
- `docs/anchor_pm/interaction_guide.md`
- `docs/anchor_pm/review_log.md`
- `docs/anchor_pm/simplification.md`
- `docs/module_state/<thread>.md`

The installer must not write them until the user confirms the installation plan.

## Installer Output Contract

The installer should reply in the user's usual conversation language. Package files may be English, and generated project documents may be English unless the user asks otherwise.

Before writing, the installer must output:

- Target project path
- Detected mode: existing project or new project
- `Observed`
- `Inference`
- `Needs Confirmation`
- Recommended thread count
- Complete per-thread initialization prompts with no user-filled placeholders
- Proposed file creates
- Proposed file updates
- Conflicts and merge risks
- Explicit approval request

After writing, the installer must output:

- Files created
- Files updated
- Files intentionally left untouched
- Next-thread usage instructions with copy-paste-ready prompts

## Safety Contract

- Do not delete files.
- Do not modify business code.
- Do not run build, test, deploy, or migration commands unless the user explicitly asks.
- Do not overwrite existing project rules without showing a merge plan.
- Do not turn inference into formal contracts.
- Strong conclusions must use `Observed / Inference / Unverified`.
