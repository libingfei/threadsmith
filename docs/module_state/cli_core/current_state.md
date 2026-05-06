# CLI Core Current State

Layer: `3`

Thread: `CLI Core`

Status: compatibility mirror of `docs/module_state/cli_core.md`.

- No CLI source exists yet.
- Intended command name: `anchorpm`.
- CLI is deferred until after Anchor PM 1.0 package-first release.
- Future CLI must reproduce package-first behavior rather than replace it.
- Contract-version and reanchor-state logic is owned by Reanchor Detector Core;
  CLI should package or call it rather than reimplement it.
