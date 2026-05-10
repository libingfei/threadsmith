# Thread Management Install Prompt

Canonical user-facing entrypoints:

- `ANCHOR_PM_INSTALL_PROMPT.en.md` for English users.
- `ANCHOR_PM_INSTALL_PROMPT.zh.md` for Chinese users.

README should link directly to language-specific copy targets for GitHub users.
`ANCHOR_PM_INSTALL_PROMPT.md` may exist as a chooser page, but should not be the
primary call to action when direct language links are available.

## Root Prompt Rule

Root install prompts are launchers, not full instruction manuals. Keep each
language-specific root prompt under 20 lines and point Codex to the package
source:

- repository: `https://github.com/libingfei/threadsmith`
- package directory: `packages/anchor-pm-1.0-standard`
- required package files: `PACKAGE_MANIFEST.md`, `ACTIVE_INSTALL_PLAN.md`,
  `INSTALL_PROMPT.md`
- fixed proposal template: `templates/install_proposal.template.md`
- fixed completion template: `templates/install_completion.template.md`

Detailed behavior belongs inside the package so public GitHub users copy a
small stable prompt while package updates can improve the installer without
expanding the root prompt.

## Product Output Rule

The `Thread Management` installer should treat its first response as an
installation confirmation page, not as a technical audit log.

Show in chat:

- target project path and detected type;
- thread names with one-sentence responsibilities;
- explicit reply options in the user's conversation language.
- at most one short approval-relevant risk line, only when needed.

Do not show by default in the main proposal:

- create/update counts;
- `AGENTS.md` handling details;
- `Observed / Inference / Needs Confirmation`;
- internal safety constraints;
- package execution details.

Omit optional rationale before approval unless there is a real approval-blocking
risk or the user asks for it. If needed, use a collapsed block. After approval,
write details to `docs/anchor_pm/install_decision_record.md`.

Thread names shown to users should follow the selected install prompt language.
For Chinese prompts, use Chinese thread names while preserving technical terms
such as Sans-IO, CLI, HTTP, and API when useful.

Thread creation prompts are also user-facing output. The complete prompts
written to `docs/anchor_pm/thread_initialization.md` must match the selected
install prompt language. For Chinese installs, prompts should be Chinese, for
example `你是 ... 线程`, while conventional technical terms may remain English.

Default reply options should be limited to approve, adjust threads, and cancel.
Do not expose `Adjust AGENTS.md` as a default option; AGENTS-specific decisions
belong in the approval-risk area only when the inspected project requires one.

The package must provide and follow `templates/install_proposal.template.md` so
the proposal shape does not drift between tests.

After approval, the completion reply should follow
`templates/install_completion.template.md`: teach the user to open
`docs/anchor_pm/thread_initialization.md`, choose one specialist, create a new
Codex conversation, paste that thread's full prompt, and keep the current
Thread Management conversation for future thread changes.

Do not paste every thread prompt into the completion chat by default. Link to
`thread_initialization.md` instead.

Link or write to generated files:

- complete per-thread prompts in `docs/anchor_pm/thread_initialization.md`;
- thread boundaries in `docs/anchor_pm/contracts.md`;
- daily usage guidance in `docs/anchor_pm/interaction_guide.md`;
- version/install state in `docs/anchor_pm/current_version.md`;
- detailed install rationale in `docs/anchor_pm/install_decision_record.md`;
- long-term thread state in `docs/module_state/*.md`.

Do not show by default:

- package file read logs;
- active plan internals;
- workflow/checklist execution detail;
- full scan file lists;
- template internals.

## Thread Management Lifecycle

After approved installation, `Thread Management` should stop the installation
task but remain available for future thread management operations, such as
adding, removing, renaming, or regenerating thread prompts.

It should not continue into ordinary business implementation work.
