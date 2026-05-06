# {{PROJECT_NAME}} Project Rules

## Project Boundary

{{PROJECT_BOUNDARY}}

## Source of Truth

Authoritative coordination files:

- `AGENTS.md`
- `docs/anchor_pm/current_version.md`
- `docs/anchor_pm/contracts.md`
- `docs/module_state/*.md`

Project reference documents remain reference material unless promoted into the coordination files above.

## Thread Protocol

Before substantial work, each long-lived thread should read:

1. `AGENTS.md`
2. `docs/anchor_pm/current_version.md`
3. `docs/anchor_pm/contracts.md`
4. Its own `docs/module_state/<thread>.md`

Threads must stay inside their contract. Cross-thread work requires a handoff summary.

## Safety Rules

- Do not delete project files unless the user explicitly asks.
- Do not overwrite existing rules without showing the planned change.
- Do not run deploy or destructive commands without explicit approval.
- Do not turn unconfirmed inference into project rules.

## Conclusion Protocol

Important conclusions must separate:

- `Observed`
- `Inference`
- `Unverified`

Use strong closure words only after a formal validation flow has run and its output has been checked.
