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

File counts, `AGENTS.md` handling, `Observed / Inference / Needs Confirmation`,
and detailed rationale must stay out of the main view by default. Use a one-line
risk only for approval-blocking conflicts. Put details in a collapsed block only
when needed, then write the full decision record after approval.

Default target files:

```text
AGENTS.md
docs/anchor_pm/current_version.md
docs/anchor_pm/contracts.md
docs/anchor_pm/thread_initialization.md
docs/anchor_pm/interaction_guide.md
docs/anchor_pm/install_decision_record.md
docs/anchor_pm/review_log.md
docs/anchor_pm/simplification.md
docs/module_state/<thread>.md
```

## Phase 4: Write After Approval

Only after explicit user approval:

1. Create missing directories.
2. Create approved new files from templates.
3. Apply approved updates to existing files.
4. Preserve existing user rules unless a specific merge was approved.

## Phase 5: Close

Output:

- Files created.
- Files updated.
- Files left untouched.
- Localized next-thread creation instructions using
  `templates/install_completion.template.md`.
- A link to `docs/anchor_pm/thread_initialization.md`, where the complete
  copy-paste-ready prompts live.

Stop here. Do not continue optimizing the target project.
