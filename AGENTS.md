# Anchor PM Framework Project Rules

## Project Boundary

This repository develops Anchor PM: a lightweight coordination framework for AI coding projects.

Anchor PM turns scattered AI coding chats into anchored specialist threads through Markdown-based contracts, module state files, reanchor prompts, handoff templates, and conservative conclusion rules.

Anchor PM 1.0 uses a Codex package-first delivery model. The 1.0 package is a Markdown-based installable coordination package, not a CLI-first implementation.

## Source of Truth

Authoritative project coordination files:

- `AGENTS.md`
- `docs/anchor_pm/current_version.md`
- `docs/anchor_pm/contracts.md`
- `docs/module_state/*.md`

Product design references:

- `PRODUCT_PRINCIPLES.md`
- `MVP_SPEC.md`
- `ANCHOR_PM_FRAMEWORK_DESIGN.md`

Reference documents are not higher-level rules unless promoted into the coordination files above.

Package artifacts:

- `packages/anchor-pm-1.0/`
- `packages/anchor-pm-1.0-standard/`
- `packages/anchor-pm-1.0-self-evolution/`

## Thread Protocol

Long-lived threads must run **Reanchor Start** automatically before substantial
work. The product target is programmatic anchoring: Codex should call a
Reanchor Detector command/tool that returns a machine-readable refresh decision,
not reread every anchor file by default. The user should not need to remember,
request, or manually run a reanchor command.

Reanchor Start:

1. If a Reanchor Detector command/tool is available, Codex should run it
   automatically for the current thread and follow its `required_reads`,
   `blocked_by`, and `next_action`.
2. If the detector is unavailable, Codex must report the degraded state, for
   example `Anchor state: unavailable; programmatic detector missing`, then fall
   back to reading:
   - `docs/anchor_pm/current_version.md`
   - `docs/anchor_pm/contracts.md`
   - the current thread state file under `docs/module_state/`
3. The fallback is a compatibility/degraded path, not the intended product
   behavior and not evidence that programmatic anchoring is complete.
4. Codex should show only a short anchor result to the user, for example
   `Anchor state: unchanged` or
   `Anchor state: changed; reading required Layer 2 dependency`.
5. State the thread identity, version, scope, and out-of-scope boundaries when
   needed.

Threads must not silently expand across module boundaries. If work crosses a boundary, produce a handoff summary.

## Safety Rules

- Do not default to modifying external user projects.
- Existing-project adoption must begin with read-only audit and draft output.
- New-project bootstrap may write initial anchors only when the target is empty or explicitly confirmed.
- Do not delete business files.
- Do not overwrite existing project rules without a displayed diff and explicit confirmation.
- Do not turn unconfirmed inference into formal contracts.
- Standard package mode must stop after installation.
- Self-evolution package mode must stop after generating recommendations.

## Conclusion Protocol

Important conclusions must separate:

- `Observed`
- `Inference`
- `Unverified`

Only use strong closure words such as `pass`, `ready`, `safe`, `consistent`, `no issue`, or `can ship` after a formal validation flow has run and its output has been checked.

## Development Environment

Project source compilation and validation should use the local `rserver` container unless a task explicitly says otherwise.

The product should still target ordinary cross-platform users. Local development constraints are not user-facing runtime requirements.
