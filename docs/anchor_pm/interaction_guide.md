# Anchor PM Interaction Guide

## Thread Setup

Use `docs/anchor_pm/thread_initialization.md` to decide which threads to create and what prompt to paste into each new Codex thread.

Use the repository-root language-specific install prompt files when adding
Anchor PM to a new target project:

- `ANCHOR_PM_INSTALL_PROMPT.en.md`
- `ANCHOR_PM_INSTALL_PROMPT.zh.md`

Choose the file that matches the user's normal Codex conversation. `Thread
Management` should remain available after installation for future thread
additions or prompt regeneration. If the Codex client cannot name a thread
before the first message, the user can paste the install prompt first and rename
the conversation later if supported.

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

## Thread Lifecycle Symmetry

Every substantial long-lived thread turn has two symmetric anchor hooks:

1. `Reanchor Start`, before work: read or refresh changed knowledge and confirm
   the current thread boundary.
2. `Closeout Knowledge Sync`, before final response: write back or hand off new
   durable knowledge created by the turn.

Reanchor Start protects the thread from stale context. Closeout Knowledge Sync
keeps the anchors evolving as conversations accumulate. If either side is
missing, the anchor system degrades: threads may start from stale knowledge, or
new corrections remain trapped in chat history.

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

Before finishing substantial work, every long-lived thread must run Closeout
Knowledge Sync.

Decision:

1. If this conversation produced durable knowledge for the current thread, update
   that thread's Layer 3 module state or category file.
2. If this conversation changed information other threads depend on, update the
   relevant Layer 2 shared-state file or produce a handoff naming the affected
   thread.
3. If this conversation implies a thread-definition, scope, or ownership change,
   hand off to Thread Management instead of silently changing the current
   thread's Layer 1 definition.
4. If this conversation implies framework-level behavior changes, hand off to
   the owning framework thread.
5. If no durable or shared knowledge changed, say so briefly.

If not, say:

```text
Closeout Knowledge Sync: no durable state update needed.
```
