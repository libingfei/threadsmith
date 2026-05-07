# Thread Management Install Prompt

Canonical user-facing entrypoints:

- `ANCHOR_PM_INSTALL_PROMPT.en.md` for English users.
- `ANCHOR_PM_INSTALL_PROMPT.zh.md` for Chinese users.

README should link directly to language-specific copy targets for GitHub users.
`ANCHOR_PM_INSTALL_PROMPT.md` may exist as a chooser page, but should not be the
primary call to action when direct language links are available.

## Product Output Rule

The `Thread Management` installer should treat its first response as an
installation confirmation page, not as a technical audit log.

Show in chat:

- target project path and detected type;
- thread names with one-sentence responsibilities;
- `AGENTS.md` handling choice;
- explicit reply options in the user's conversation language.
- at most one short approval-relevant risk line, only when needed.

Do not show by default in the main proposal:

- create/update counts;
- `Observed / Inference / Needs Confirmation`;
- internal safety constraints;
- package execution details.

Omit optional rationale before approval unless there is a real approval-blocking
risk or the user asks for it. If needed, use a collapsed block. After approval,
write details to `docs/anchor_pm/install_decision_record.md`.

Thread names shown to users should follow the selected install prompt language.
For Chinese prompts, use Chinese thread names while preserving technical terms
such as Sans-IO, CLI, HTTP, and API when useful.

Default reply options should be limited to approve, adjust threads, and cancel.
Do not expose `Adjust AGENTS.md` as a default option; AGENTS-specific decisions
belong in the approval-risk area only when the inspected project requires one.

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
