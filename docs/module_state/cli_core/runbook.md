# CLI Core Runbook

Layer: `3`

Thread: `CLI Core`

Status: compatibility mirror of `docs/module_state/cli_core.md`.

Expected first steps after 1.0:

1. Check `rserver` toolchain availability.
2. Propose minimal Go module layout.
3. Implement read-only `status` and `reanchor` against package-generated anchors.
4. Add tests that compare CLI output to package workflow expectations.
5. Implement `init` only after status/reanchor behavior is stable.
