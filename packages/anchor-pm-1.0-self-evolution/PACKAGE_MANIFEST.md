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

`INSTALL_PROMPT.md` may differ by mode because it points to the relevant public
package directory.

## Required Package Files

- `INSTALL_PROMPT.md`
- `PACKAGE_MANIFEST.md`
- `ACTIVE_INSTALL_PLAN.md`
- `plans/standard_project_install.md`
- `plans/self_evolution_install.md`
- `templates/install_proposal.template.md`
- `templates/install_completion.template.md`
- `templates/thread_initialization.template.md`
- `templates/*.template.md`
- `workflows/*.md`
- `checklists/*.md`

## Target Outputs

The installer may propose these target-project files:

- `.threadsmith/AGENTS.md`
- `.threadsmith/current_version.md`
- `.threadsmith/contracts.md`
- `.threadsmith/thread_initialization.md`
- `.threadsmith/interaction_guide.md`
- `.threadsmith/install_decision_record.md`
- `.threadsmith/review_log.md`
- `.threadsmith/simplification.md`
- `.threadsmith/module_state/<thread>.md`

The installer must not write them until the user confirms the installation plan.
By default, all Anchor PM files live under `.threadsmith/` so ordinary project
files and unrelated Codex conversations are not affected.

## Installer Output Contract

The installer should reply in the user's usual conversation language. Package files may be English, and generated project documents may be English unless the user asks otherwise.

Exception: user-facing thread names, thread creation prompts in
`.threadsmith/thread_initialization.md`, and the installation completion
message must match the install-prompt language. These are instructions users
copy into Codex, so they must not silently switch language.

Before writing, the installer must use
`templates/install_proposal.template.md` for the user-visible proposal.

The main proposal view should show only:

- target project path and detected type;
- proposed project specialist threads with one-sentence responsibilities;
- reply options to approve install, adjust threads, or cancel.

Do not show file counts, root `AGENTS.md` integration, `Observed / Inference /
Needs Confirmation`, package execution details, or internal safety explanations
in the main view by default. Show a one-line risk only when it affects approval.
Detailed rationale belongs in an optional collapsed block before approval, or in
`.threadsmith/install_decision_record.md` after approval.

After writing, the installer must output:

- a localized completion page using `templates/install_completion.template.md`;
- next-thread creation instructions that teach the user to open
  `.threadsmith/thread_initialization.md`, create a new Codex conversation,
  and paste the chosen thread's full prompt;
- links to key generated files;
- short summaries of files created, files updated, and files intentionally left
  untouched.

Do not paste every thread prompt into the completion chat by default.

## Safety Contract

- Do not delete files.
- Do not copy the Threadsmith repository into the target project.
- Do not modify business code.
- Do not run build, test, deploy, or migration commands unless the user explicitly asks.
- Do not create or modify root `AGENTS.md` by default.
- Do not overwrite existing project rules without showing a merge plan.
- Do not turn inference into formal contracts.
- Strong conclusions must use `Observed / Inference / Unverified`.
