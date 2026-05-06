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
