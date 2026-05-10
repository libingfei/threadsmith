# Product Manager Key Memory

Layer: `3`

Thread: `Product Manager`

Key memory:

- Anchor PM's product goal is to improve real AI coding experience, not to
  create process for its own sake.
- The strongest value hypothesis is specialist threads with narrower context,
  plus lightweight reanchor and handoff for changing shared information.
- Reanchor should confirm all registered layer states, but only read changed,
  unknown, unreadable, inbound, or task-relevant content.
- Periodic reanchor safety check target: every 10 conversation rounds.
- Reanchor module development should be based on a fixed I/O contract:
  `ReanchorRequest` plus registered file fingerprints in, `ReanchorResult` with
  required reads, blockers, next action, and checkpoint update proposal out.
- Reanchor is a Codex responsibility, not a user step: Codex should trigger it
  automatically before substantial work and should not ask users to run CLI
  commands.
- Programmatic anchoring means Codex receives a detector result and reads only
  required files; automatically rereading anchors is only a degraded fallback.
- PM validation must evaluate semantic product value, not checklist field
  presence. A proposal that is structurally complete but gives the wrong thread
  model, impossible user steps, or confusing options is a product failure.
- Durable PM behavior changes must be written into thread state, but no-change
  closeout should stay silent unless status was requested.
- Knowledge Sync Gate is not PM-specific. Every long-lived thread must check
  before final response whether new or changed durable knowledge belongs in
  local Layer 3 state, Layer 2 shared state/handoff, a Thread Management Layer 1
  update request, or a framework-owner handoff.
- Anchor Gate and Knowledge Sync Gate are a pair: read only necessary changed
  knowledge before work; write or hand off only durable new knowledge before
  reply. Without both, anchors do not continuously improve.
- MVP validation should follow the same lifecycle: install anchors, create
  specialist thread, run Anchor Gate, solve a scoped problem, hand off
  cross-boundary work, run Knowledge Sync Gate, verify shared-state recovery,
  then restore the repo.
- Before the user tests from GitHub, sync the repository first. Stale remote
  prompts invalidate install-flow test results.
- User-visible thread names must follow the selected install prompt language.
  AGENTS.md handling belongs in the proposal risk/decision area only when it is
  approval-relevant; do not expose `Adjust AGENTS.md` as a default reply option.
- Use lightweight gates: Anchor Gate before work and Knowledge Sync Gate before
  final response. Default is silent/no-write/minimal-read; escalate only when
  changed, unknown, blocked, conflicting, degraded, or durability-relevant.
- Root install prompts should stay under 20 lines and work as package-index
  launchers. Fixed proposal shape belongs in
  `templates/install_proposal.template.md`, not in the root prompt.
- Thread creation prompts are user-facing output and must match the install
  prompt language. Completion pages should teach users to open
  `thread_initialization.md`, create a new Codex conversation, and paste one
  specialist prompt, instead of dumping every prompt into chat.
