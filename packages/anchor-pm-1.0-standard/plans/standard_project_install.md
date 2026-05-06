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
4. Read templates only when preparing the proposal.

## Phase 2: Audit Target

Classify the target as:

- `existing`: non-empty project, or project with existing rules/docs/code.
- `new`: empty or nearly empty project.

Output findings as:

- `Observed`: direct file and directory facts.
- `Inference`: likely thread/module boundaries.
- `Needs Confirmation`: assumptions that must not become formal contracts yet.

## Phase 3: Propose Installation

Before writing, output:

- Proposed mode: existing or new.
- Proposed thread list.
- Files to create.
- Files to update.
- Existing rule files that will not be overwritten.
- Conflicts requiring user choice.
- Recommended thread count and complete per-thread initialization prompts.
- Exact approval request.

Default target files:

```text
AGENTS.md
docs/anchor_pm/current_version.md
docs/anchor_pm/contracts.md
docs/anchor_pm/thread_initialization.md
docs/anchor_pm/interaction_guide.md
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
- Recommended next thread prompts that are ready to copy and paste.

Stop here. Do not continue optimizing the target project.
