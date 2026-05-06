# Anchor PM Interaction Guide

## Thread Setup

Use `docs/anchor_pm/thread_initialization.md` to decide which threads to create and what prompt to paste into each new Codex thread.

## New Thread Reanchor Prompt

```text
You are the <thread name> thread for this project.
Before work, read AGENTS.md, docs/anchor_pm/current_version.md,
docs/anchor_pm/contracts.md, and docs/module_state/<thread>.md.
State your scope and out-of-scope boundaries if they affect the task.
```

## Existing Thread Reanchor Prompt

```text
Reanchor this thread under the current Anchor PM version.
Check docs/anchor_pm/current_version.md and the relevant module state file.
Continue only within this thread's contract.
```

## Handoff Template

```text
Source thread:
Target thread:
Current conclusion:
Confirmed facts:
Impact:
Do not repeat:
Questions for target thread:
Suggested next step:
```

## Closeout Reminder

Before finishing substantial work, decide whether any module state file needs updates.

If not, say:

```text
No module state update needed for this task.
```
