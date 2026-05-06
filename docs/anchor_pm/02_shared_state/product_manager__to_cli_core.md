# Product Manager To CLI Core

Layer: `2`

Source thread: `Product Manager`

Target thread: `CLI Core`

Status: active shared dependency.

## Confirmed Product Requirements

- Programmatic reanchor is required for the intended user experience. Automatic
  rereading of anchor files is only a degraded compatibility path.
- Codex should be able to invoke a stable project-local command or equivalent
  tool before substantial work and receive machine-readable `ReanchorResult`
  JSON.
- The minimum user-facing behavior is a short anchor-state line; users should
  not see fingerprint tables, registry dumps, or CLI operation details.
- Missing runtime entrypoint should be visible as
  `Anchor state: unavailable; programmatic detector missing`, then Codex may
  use the fallback read path.
- This does not require turning Anchor PM into a broad CLI-first product. The
  immediate need is a narrow callable reanchor entrypoint that packages or calls
  Reanchor Detector Core.

## Suggested Runtime Shape

```text
anchorpm reanchor --thread <thread_id> --operation ordinary_thread_start --json
```

Equivalent local executable or Codex tool integration is acceptable if it
returns the same schema documented in
`docs/anchor_pm/00_framework_baseline/reanchor_module_io_spec.md`.

## Target Next Step

Design the smallest CLI/tool wrapper that exposes Reanchor Detector Core without
duplicating detector semantics. Keep audit/init/apply CLI features separate
from this minimum runtime entrypoint.
