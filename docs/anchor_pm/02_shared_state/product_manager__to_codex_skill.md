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
- User-visible thread names must follow the selected install prompt language.
  For the Chinese install prompt, use Chinese thread names while preserving
  conventional technical terms when useful.
- Do not expose "install docs only; do not update AGENTS.md" as a normal
  approval option; if AGENTS handling is blocked, ask for a specific merge/skip
  decision and mark the consequence.
- Do not expose `Adjust AGENTS.md` as a default reply option. The default
  install choices should be approve, adjust threads, and cancel; AGENTS-specific
  decisions should appear only when the inspected project requires one.
- Generated thread prompts should include a lightweight Anchor Gate before work
  and Knowledge Sync Gate before final response. The gates should default to
  silent/no-write/minimal-read behavior.
- Anchor Gate includes same-turn user correction handling and Reanchor Start,
  but should not be exposed as multiple user-visible steps.
- Knowledge Sync Gate updates or hands off only durable local/shared knowledge;
  no durable change means no visible closeout note by default.
- Skill/installer guidance should enforce an anchor budget: gate handling must
  not crowd out the actual task response.
- Public root install prompts should stay under 20 lines and function as
  package indexes. The detailed install behavior and fixed proposal shape live
  in the package, especially `templates/install_proposal.template.md`.

## Target Next Step

Update install prompt behavior and any future Skill guidance so automatic
Anchor Gate is included in generated thread prompts without adding a user manual
step, and so existing-project thread proposals are module/subsystem based
instead of generic role based. Include Knowledge Sync Gate in the generated
thread behavior so durable updates do not remain only in chat history, while
normal turns stay silent when no state change is needed.
