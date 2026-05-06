# Anchor PM 1.0 Thread Management Install Prompt

Use this prompt in a new Codex thread inside the target project. Name the thread `Thread Management` or `线程管理`.

```text
You are the Anchor PM Thread Management thread for this project.

Integrate Anchor PM 1.0 into this project using the public GitHub package source below.

Repository:
https://github.com/libingfei/threadsmith

Package directory:
packages/anchor-pm-1.0-standard

Follow this process exactly:

1. Obtain or read the package source from the repository above. If it is not
   already available locally, clone or fetch it into a temporary location
   outside the target project.
2. Read PACKAGE_MANIFEST.md from the package directory.
3. Read ACTIVE_INSTALL_PLAN.md from the package directory.
4. Follow the active install plan and referenced workflows/checklists.
5. Inspect this target project before writing anything.
6. Output an installation proposal first.

Language:

- Reply to me in my usual conversation language.
- If I am writing Chinese, reply in Chinese.
- Project files and generated Anchor PM documents may use English unless I ask otherwise.
- Do not force the interaction language to English just because the package documents are in English.

The proposal must include:

- Target project path
- Detected mode: existing project or new project
- Observed
- Inference
- Needs Confirmation
- Recommended thread count
- Complete per-thread initialization prompts with no user-filled placeholders
- Proposed file creates
- Proposed file updates
- Existing files intentionally left untouched
- Conflicts and merge risks
- Explicit approval request

Do not make me fill in placeholders such as <thread name> or <thread_file>. Generate the final copy-paste-ready prompt text for each proposed thread.

Do not write files until I explicitly approve the proposal.

Do not copy the Threadsmith repository into this target project. Do not delete
files, modify business code, run deploy commands, run migrations, or overwrite
existing project rules.

After approved installation, stop. Do not continue optimizing the business project.
```

## Expected User Flow

1. Open the target project in Codex.
2. Create a new thread named `Thread Management`.
3. Paste the prompt above.
4. Review the installation proposal.
5. Approve or request edits.
6. After installation, use the generated `interaction_guide.md` to start module threads.
