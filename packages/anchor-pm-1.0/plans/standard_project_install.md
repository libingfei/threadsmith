# Standard Project Install Plan

Use this plan when installing Anchor PM into a non-Anchor PM project.

## Goal

Deploy the smallest useful Anchor PM coordination structure into the target project, then stop.

## Inputs

- Target project root
- Package root
- Optional user-provided thread names
- Existing project documents and rules

## Phase 1: Read Package

1. Read `PACKAGE_MANIFEST.md`.
2. Read this active plan.
3. Read:
   - `workflows/existing_project_adoption.md`
   - `workflows/new_project_bootstrap.md`
   - `checklists/safety_check.md`
   - `checklists/conclusion_check.md`
4. Read `templates/install_proposal.template.md` before composing the first
   user-visible proposal.
5. Read `templates/install_completion.template.md` before composing the
   post-approval completion reply.
6. Read other templates only when preparing approved file outputs.

## Phase 2: Audit Target

Classify the target as:

- `existing`: non-empty project, or project with existing rules/docs/code.
- `new`: empty or nearly empty project.

Output findings as:

- `Observed`: direct file and directory facts.
- `Inference`: likely thread/module boundaries.
- `Needs Confirmation`: assumptions that must not become formal contracts yet.

## Phase 3: Propose Installation

Before writing, output the concise proposal defined in
`templates/install_proposal.template.md`.

Main-view content is limited to:

- target project path and detected type;
- proposed project specialist threads with one-sentence responsibilities;
- approval / adjust threads / cancel reply options.

File counts, root `AGENTS.md` integration, `Observed / Inference / Needs
Confirmation`, and detailed rationale must stay out of the main view by
default. Use a one-line risk only for approval-blocking conflicts. Put details
in a collapsed block only when needed, then write the full decision record after
approval.

Default target files:

```text
.threadsmith/AGENTS.md
.threadsmith/current_version.md
.threadsmith/contracts.md
.threadsmith/thread_initialization.md
.threadsmith/interaction_guide.md
.threadsmith/install_decision_record.md
.threadsmith/review_log.md
.threadsmith/simplification.md
.threadsmith/module_state/<thread>.md
```

## Phase 4: Write After Approval

Only after explicit user approval:

1. Create missing directories.
2. Create approved new files from templates.
3. Apply approved updates to existing files.
4. Preserve existing user rules unless a specific root-level merge was
   explicitly requested and approved.

## Phase 5: Close

Output:

- A localized installation-success message using
  `templates/install_completion.template.md`.
- A short user-facing explanation of what Threadsmith added.
- Localized next-thread creation instructions.
- Complete copy-paste-ready prompts for each recommended specialist thread.
- A short isolation note: all Threadsmith files are under `.threadsmith/`, and
  ordinary Codex conversations are unaffected unless started with the generated
  prompts.
- A short Thread Management note: this current conversation can create, remove,
  rename, adjust, regenerate, or query threads.

Do not show the generated-file inventory by default. Store detailed file
changes and rationale in `.threadsmith/install_decision_record.md`.

Stop here. Do not continue optimizing the target project.
