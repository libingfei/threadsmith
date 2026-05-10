# Anchor PM Interaction Guide

## Thread Setup

Use `.threadsmith/thread_initialization.md` to decide which threads to create
and what prompt to paste into each new Codex thread.

To create a project specialist thread:

1. Open `.threadsmith/thread_initialization.md`.
2. Choose one thread.
3. Start a new Codex conversation in the target project.
4. Paste that thread's complete prompt as the first message.
5. Keep the original Thread Management conversation for future thread changes.

Thread names and paste-ready prompts should match the install-prompt language.
Technical terms may remain in their conventional language.

## New Thread Anchor Gate

```text
You are the <thread name> thread for this project.
Before work, run Anchor Gate silently unless changed, blocked, unknown,
conflicting, degraded, or the user gives a durable correction. Read only the
required anchors.
State your scope and out-of-scope boundaries if they affect the task.
```

## Existing Thread Reanchor Prompt

```text
Run Anchor Gate for this thread under the current Anchor PM version.
Use a detector if available; otherwise report the degraded state and read the
required anchors yourself.
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

## Knowledge Sync Reminder

Before the final response after substantial work, run Knowledge Sync Gate.
Update or hand off only durable local/shared knowledge. If nothing durable
changed, keep the gate silent.
