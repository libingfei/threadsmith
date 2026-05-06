# CLI Core Thread State

## Contract

Thread contract: `CLI Core` in `docs/anchor_pm/contracts.md`

## Current State

- No CLI source exists yet.
- Intended command name: `anchorpm`.
- CLI is deferred until after Anchor PM 1.0 package-first release.
- Future CLI must reproduce package-first behavior rather than replace it.
- Contract-version and reanchor-state logic is owned by Reanchor Detector Core; CLI should package or call it rather than reimplement it.

## Open Issues

- Confirm Go toolchain availability in `rserver` only when CLI work begins.
- Freeze repository layout for CLI source after package-first validation.
- Define stable output schemas for audit, status, reanchor, and handoff.
- Define safe apply behavior and diff format.
- Define CLI integration points for Reanchor Detector Core.

## Runbook

Expected first steps after 1.0:

1. Check `rserver` toolchain availability.
2. Propose minimal Go module layout.
3. Implement read-only `status` and `reanchor` against package-generated anchors.
4. Add tests that compare CLI output to package workflow expectations.
5. Implement `init` only after status/reanchor behavior is stable.

## History / Notes

- Created as a planned implementation thread during self-dogfood adoption.
