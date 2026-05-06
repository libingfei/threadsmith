# Anchor PM Interaction Guide

## Thread Setup

Use `docs/anchor_pm/thread_initialization.md` to decide which threads to create and what prompt to paste into each new Codex thread.

Use the repository-root language-specific install prompt files when adding
Anchor PM to a new target project:

- `ANCHOR_PM_INSTALL_PROMPT.en.md`
- `ANCHOR_PM_INSTALL_PROMPT.zh.md`

Choose the file that matches the user's normal Codex conversation. `Thread
Management` should remain available after installation for future thread
additions or prompt regeneration.

## Automatic Reanchor Behavior

Codex should trigger reanchor automatically before substantial work in a
long-lived Anchor PM thread. Users should not need to ask for reanchor or run a
CLI command manually.

Preferred behavior:

1. Use the detector command/tool when available.
2. Follow its required reads and blockers.
3. If unavailable, report the degraded state, then read the required anchors as
   a compatibility fallback.
4. Show only a short `Anchor state` line unless blocked.

## New Thread Reanchor Prompt

For user-facing thread creation, prefer the complete prompts in `docs/anchor_pm/thread_initialization.md`.

This generic pattern is for maintainers and package generation only:

```text
You are the <thread name> thread for this project.
Before work, run Reanchor Start automatically. If a detector command/tool is
available, use it and follow required_reads; otherwise report `Anchor state:
unavailable; programmatic detector missing`, then read AGENTS.md,
docs/anchor_pm/current_version.md, docs/anchor_pm/contracts.md, and
docs/module_state/<thread>.md. Do not ask the user to run CLI commands.
State your scope and out-of-scope boundaries if they affect the task.
```

## Existing Thread Reanchor Prompt

```text
Run Reanchor Start for this thread under the current Anchor PM version.
Use the detector if available; otherwise report the degraded state, then check
docs/anchor_pm/current_version.md and the relevant module state file yourself.
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
