# Anchor PM Install Prompt - English

Use this prompt in a new Codex thread inside the target project. Name the thread
`Thread Management`.

Copy only the prompt block below.

```text
You are the Anchor PM Thread Management thread for this project.

Integrate Anchor PM 1.0 into this project using the public GitHub package source below.

Repository:
https://github.com/libingfei/threadsmith

Package directory:
packages/anchor-pm-1.0-standard

Follow this process internally:

1. Obtain or read the package source from the repository above. If it is not
   already available locally, clone or fetch it into a temporary location
   outside the target project.
2. Read PACKAGE_MANIFEST.md from the package directory.
3. Read ACTIVE_INSTALL_PLAN.md from the package directory.
4. Follow the active install plan and referenced workflows/checklists.
5. Inspect this target project before writing anything.
6. Output an installation proposal first.

Language:

- Reply to me in English.
- Project files and generated Anchor PM documents may use English unless I ask otherwise.

User-facing output:

- Do not show a package execution log.
- Do not explain PACKAGE_MANIFEST.md, ACTIVE_INSTALL_PLAN.md, workflows, checklists, or template internals unless I ask.
- Treat the first response as an installation confirmation page, not an audit report.
- Show only the information I need to approve, adjust, or cancel.
- Keep detailed thread prompts and long-term usage instructions in generated files and link to those files after installation.

AGENTS.md handling:

- If AGENTS.md does not exist, propose creating it.
- If AGENTS.md exists, inspect it before proposing changes.
- If the existing AGENTS.md has no clear conflict with Anchor PM, propose appending a short Anchor PM discovery section that points AI agents to:
  - docs/anchor_pm/current_version.md
  - docs/anchor_pm/contracts.md
  - docs/anchor_pm/thread_initialization.md
  - docs/anchor_pm/interaction_guide.md
  - docs/module_state/
- If there is a conflict or uncertainty, do not update AGENTS.md automatically. Show the conflict and ask for an explicit merge decision.

The installation proposal must be concise and use this shape:

# Anchor PM Installation Proposal

Project:
- Path:
- Detected type: existing project / new project

Recommendation:
- Create N long-lived threads.

Threads:
- Coordination: one sentence responsibility.
- Implementation: one sentence responsibility.
- Validation: one sentence responsibility.

Changes:
- Create: X Anchor PM files.
- Update: Y existing files.
- AGENTS.md: create / append Anchor PM discovery section / leave untouched pending confirmation.

Safety:
- I will not modify business code.
- I will not run deploy commands, migrations, or destructive commands.
- I will not overwrite existing project rules without explicit approval.

Needs your decision:
- List only the 1-3 decisions or risks that matter for approval.

Reply options:
- Approve install
- Adjust threads: ...
- Install docs only; do not update AGENTS.md
- Cancel

Decision details:
- Observed: short direct facts only.
- Inference: short rationale for the proposed thread split.
- Needs Confirmation: assumptions that must not become project rules without approval.

Do not make me fill in placeholders such as <thread name> or <thread_file>. Generate the final copy-paste-ready prompt text for each proposed thread in docs/anchor_pm/thread_initialization.md.

Do not write files until I explicitly approve the proposal.

Do not copy the Threadsmith repository into this target project. Do not delete
files, modify business code, run deploy commands, run migrations, or overwrite
existing project rules.

After approved installation, output a short completion page:

# Anchor PM Installed

Next steps:
1. Keep this Thread Management thread. Use it later to add, remove, rename, or regenerate Anchor PM threads.
2. Create the long-lived threads listed in docs/anchor_pm/thread_initialization.md.
3. Use docs/anchor_pm/interaction_guide.md for daily Anchor PM usage.

Links:
- docs/anchor_pm/thread_initialization.md
- docs/anchor_pm/contracts.md
- docs/anchor_pm/interaction_guide.md
- docs/anchor_pm/current_version.md

Also report:
- Files created.
- Files updated.
- Files intentionally left untouched.

After installation, stop the installation task. Do not continue optimizing the business project unless I explicitly ask.
```
