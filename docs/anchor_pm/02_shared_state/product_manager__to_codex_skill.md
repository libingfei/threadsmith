# Product Manager To Codex Skill

Layer: `2`

Source thread: `Product Manager`

Target thread: `Codex Skill / Package Installer`

Status: active shared dependency.

## Confirmed Product Requirements

- Install and thread-creation flows should set the expectation that Codex, not
  the user, triggers Reanchor Start before substantial work.
- A future Skill may describe how to run and interpret reanchor, but it must not
  be treated as a guaranteed session-start hook.
- Installer wording should avoid telling users to run CLI commands before each
  thread interaction.
- If a target project has a detector command/tool, generated thread prompts
  should tell Codex to use it automatically.
- If no detector command/tool exists, generated thread prompts should preserve a
  degraded fallback that Codex performs itself by reading the required anchors.
- Skill or installer wording must not present fallback rereading as completed
  programmatic anchoring.
- Existing-project installation should derive project specialist threads from
  real target-project modules/subsystems instead of using generic Coordination /
  Implementation / Validation defaults.
- The current install conversation is the Thread Management entrypoint. The
  installer must not depend on the user being able to name the Codex
  conversation before the first message.
- `Coordination` should not be proposed as an ordinary target-project thread.
- User-facing install proposals should be concise confirmation pages. Safety
  constraints and package execution internals should stay internal unless the
  user asks.
- Do not expose "install docs only; do not update AGENTS.md" as a normal
  approval option; if AGENTS handling is blocked, ask for a specific merge/skip
  decision and mark the consequence.
- Generated thread prompts must include Closeout Knowledge Sync, not only
  Reanchor Start. At closeout, Codex should decide whether to update local
  Layer 3 state, Layer 2 shared state/handoff, request Thread Management for
  Layer 1 changes, or state that no durable update is needed.
- Skill/installer guidance should frame Reanchor Start and Closeout Knowledge
  Sync as symmetric lifecycle hooks: read before work, write or hand off before
  reply.

## Target Next Step

Update install prompt behavior and any future Skill guidance so automatic
Reanchor Start is included in generated thread prompts without adding a user
manual step, and so existing-project thread proposals are module/subsystem based
instead of generic role based. Include Closeout Knowledge Sync in the generated
thread behavior so knowledge updates become durable instead of remaining only in
chat history.
