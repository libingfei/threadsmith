# Product Manager To Templates Protocol

Layer: `2`

Source thread: `Product Manager`

Target thread: `Templates / Protocol`

Status: active shared dependency.

## Confirmed Product Requirements

- Reanchor must be presented as automatic Codex behavior, not a user-operated
  CLI step.
- Package workflow text should say Codex runs Reanchor Start before substantial
  long-lived thread work.
- If a detector command/tool is available, Codex should run it and follow
  `required_reads`, `blocked_by`, and `next_action`.
- If the detector is unavailable, Codex should report the degraded state, then
  read the required anchors itself as a compatibility fallback.
- Template wording must not present fallback rereading as the intended product
  behavior.
- Users should see only a short anchor-state line unless the thread is blocked.
- Templates should not imply that users must remember to request reanchor or run
  a CLI command manually.
- Existing-project install templates must generate module/subsystem-based
  project specialist threads, not default Coordination / Implementation /
  Validation responsibility buckets.
- `Coordination` is an Anchor PM internal/self-hosting role and should not be
  proposed as an ordinary target-project thread.
- The current install conversation is the Thread Management entrypoint; user
  prompts must not require the Codex conversation to be named before the first
  message.
- Each proposed target-project specialist should normally own code, tests, docs,
  and validation evidence for its module. Standalone `Validation` should be used
  only when the target project has a clear independent validation/release
  subsystem or the user asks for it.
- User-facing install proposals should hide verbose internal safety constraints
  and package execution details unless the user asks.
- "Install docs only; do not update AGENTS.md" should not appear as a default
  approval option. If AGENTS handling is blocked, the package should describe
  the consequence and ask for a specific merge/skip decision.
- Every generated long-lived thread prompt must include Closeout Knowledge Sync:
  before final response after substantial work, the thread checks whether new or
  changed knowledge belongs in its own Layer 3 state, a Layer 2 shared-state
  file or handoff, a Thread Management Layer 1 update request, or a
  framework-owner handoff.
- The generated interaction guide should treat Closeout Knowledge Sync as
  mandatory all-thread behavior, parallel to Reanchor Start.
- Wording should present Reanchor Start and Closeout Knowledge Sync as symmetric
  lifecycle hooks: start reads changed knowledge and confirms boundaries before
  work; closeout writes or hands off new durable knowledge before the final
  reply.

## Target Next Step

Mirror this behavior into package workflow text and reusable protocol wording.
Known package locations still needing owner-thread update include
`packages/*/workflows/reanchor.md`,
`packages/*/templates/interaction_guide.template.md`, and
`packages/*/templates/module_state.template.md`.

Additional locations needing owner-thread review after the Flask install
dry-run include `packages/*/templates/thread_initialization.template.md`,
`packages/*/workflows/new_project_bootstrap.md`,
`packages/*/workflows/existing_project_adoption.md`, and package
`INSTALL_PROMPT.md` files. Closeout Knowledge Sync also needs mirroring into
`packages/*/templates/interaction_guide.template.md` and any generated
per-thread prompt text.

Do not change detector internals or CLI packaging here.
