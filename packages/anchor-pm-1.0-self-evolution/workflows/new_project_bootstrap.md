# Workflow: New Project Bootstrap

Use this workflow when the target project is empty or nearly empty.

## Goal

Create the smallest useful Anchor PM anchor set before the project grows.

## Default Threads

If the user does not provide thread names, propose:

- `Coordination`
- `Implementation`
- `Validation`

Keep the list short. More threads can be added later.

## Proposed Files

```text
AGENTS.md
docs/anchor_pm/current_version.md
docs/anchor_pm/contracts.md
docs/anchor_pm/thread_initialization.md
docs/anchor_pm/interaction_guide.md
docs/anchor_pm/review_log.md
docs/anchor_pm/simplification.md
docs/module_state/coordination.md
docs/module_state/implementation.md
docs/module_state/validation.md
```

## Confirmation

Even for new projects, show the proposed file list, thread list, and complete per-thread initialization prompts before writing.

Do not ask the user to fill in thread-name or state-file placeholders.
