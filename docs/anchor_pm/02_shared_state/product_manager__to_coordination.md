# Product Manager To Coordination

Layer: `2`

Source thread: `Product Manager`

Target thread: `Coordination`

Status: active shared dependency.

## Confirmed Facts

- Layer 0 through Layer 3 structure has been initialized with numbered internal
  paths.
- Internal Layer directories no longer use `README.md` files.
- `docs/anchor_pm/01_thread_definitions/*.md` now contain complete Layer 1
  semantic mirrors for all seven current project threads.
- The mirrored content includes scope, out-of-scope boundaries, acceptance
  criteria, hard rules, state file path, handoff rule, and initialization
  prompt.
- `docs/anchor_pm/contracts.md` and `docs/anchor_pm/thread_initialization.md`
  remain current authoritative compatibility sources.

## Product Recommendation

Do not delete old Layer 1 source files yet.

Recommended next Coordination decision:

- Decide whether `docs/anchor_pm/01_thread_definitions/*.md` should become the
  authoritative Layer 1 source in the next coordination version.
- If yes, convert `contracts.md` and `thread_initialization.md` into short
  compatibility redirects after validation.
- If no, keep the split files as detector handles only and document the reason.

## Suggested Validation Before Promotion

- Verify every thread can reanchor from its own `01_thread_definitions/<thread>.md`.
- Verify no scope, boundary, acceptance, hard rule, state file, handoff, or
  prompt information was lost during the mirror.
- Verify package templates and install prompts can either generate from or point
  to the new Layer 1 files.

## Requested Owner Action

Coordination should decide whether to promote Layer 1 split files in the next
version, or explicitly keep the old Layer 1 source files authoritative.
