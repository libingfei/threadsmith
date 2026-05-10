# Interaction Workflow

Layer: `2`

Status: compatibility mirror.

Authoritative source until promoted:

- `docs/anchor_pm/interaction_guide.md`

## Thread Setup

Use `docs/anchor_pm/thread_initialization.md` to decide which threads to create
and what prompt to paste into each new Codex thread.

Thread creation flow:

1. Open `docs/anchor_pm/thread_initialization.md`.
2. Choose one specialist thread.
3. Start a new Codex conversation in the same target project.
4. Paste that thread's complete prompt as the first message.
5. Keep the original Thread Management conversation for future thread changes.

Thread creation prompts are user-facing and should match the install prompt
language. Technical terms may remain in their conventional language.

Use the repository-root language-specific install prompt files when adding
Anchor PM to a new target project:

- `ANCHOR_PM_INSTALL_PROMPT.en.md`
- `ANCHOR_PM_INSTALL_PROMPT.zh.md`

Choose the file that matches the user's normal Codex conversation.
These root prompts should stay under 20 lines and act as launchers into the
package. Detailed install behavior and the fixed proposal shape live in the
package, especially `packages/anchor-pm-1.0-standard/INSTALL_PROMPT.md` and
`packages/anchor-pm-1.0-standard/templates/install_proposal.template.md`.

`Thread Management` should remain available after installation for future
thread additions or prompt regeneration. If the Codex client cannot name a
thread before the first message, the user can paste the install prompt first and
rename the conversation later if supported.

## Lightweight Anchor Gates

Codex should run anchor handling as lightweight gates, not as visible process.
Users should not need to ask for reanchor or run a CLI command manually.

### Anchor Gate

Run before substantial work. It combines user-delta triage and Reanchor Start.

Default path:

- Do not write anchors.
- Do not reread full anchor files.
- Do not explain the gate in chat.
- Use the detector when available and read only changed, unknown, inbound, or
  task-relevant anchors.

Escalate only when:

- The user gives an explicit durable correction, rule change, preference, or
  cross-thread fact.
- An anchor is changed, unknown, unreadable, or conflicts with the user message.
- The task crosses thread scope or requires a handoff.
- No detector is available and a fallback read is required.

When escalated, keep chat output to one short status line unless blocked.
Stale anchors must not silently override a same-turn user correction.

### Knowledge Sync Gate

Run before the final response after substantial work.

Default path:

- Do not write anchors.
- Do not mention closeout.

Escalate only when the turn created durable local knowledge, changed shared
knowledge, needs a thread-definition update, or needs a framework/owner handoff.

The anchor budget should stay below the task budget. If gate handling starts to
consume visible conversation space, collapse it back to status-only output.

## Reanchor Prompt Pattern

For user-facing thread creation, prefer the complete prompts in
`docs/anchor_pm/thread_initialization.md` or, after Layer 1 promotion, the
per-thread prompt in `docs/anchor_pm/01_thread_definitions/<thread>.md`.

Generic maintainer pattern:

```text
You are the <thread name> thread for this project.
Before work, run Anchor Gate silently unless changed, blocked, unknown,
conflicting, or degraded: classify explicit durable user corrections, run
Reanchor Start, use a detector if available, and read only required anchors. If
unavailable, report the degraded state and read AGENTS.md,
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

## Knowledge Sync Reminder

Before finishing substantial work, every long-lived thread must run the
Knowledge Sync Gate.

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
5. If no durable or shared knowledge changed, do not spend visible chat space on
   the gate unless the user asked for status.
