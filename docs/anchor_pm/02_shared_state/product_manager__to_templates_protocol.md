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

## Target Next Step

Mirror this behavior into package workflow text and reusable protocol wording.
Known package locations still needing owner-thread update include
`packages/*/workflows/reanchor.md`,
`packages/*/templates/interaction_guide.template.md`, and
`packages/*/templates/module_state.template.md`.

Do not change detector internals or CLI packaging here.
