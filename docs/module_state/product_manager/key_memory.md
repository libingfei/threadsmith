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
- Durable PM behavior changes must be written into thread state before closeout;
  otherwise reanchor cannot carry the correction into future work.
- Closeout Knowledge Sync is not PM-specific. Every long-lived thread must
  check before final response whether new or changed knowledge belongs in local
  Layer 3 state, Layer 2 shared state/handoff, a Thread Management Layer 1
  update request, or a framework-owner handoff.
- Reanchor Start and Closeout Knowledge Sync are a symmetric pair: read changed
  knowledge and confirm boundaries before work; write or hand off new durable
  knowledge before reply. Without both, anchors do not continuously improve.
- MVP validation should follow the same lifecycle: install anchors, create
  specialist thread, run Reanchor Start, solve a scoped problem, hand off
  cross-boundary work, run Closeout Knowledge Sync, verify shared-state recovery,
  then restore the repo.
- Before the user tests from GitHub, sync the repository first. Stale remote
  prompts invalidate install-flow test results.
