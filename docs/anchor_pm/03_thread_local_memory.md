# Layer 3: Thread Local Memory

Purpose: describe where thread-local memory lives without placing user-facing
README files inside every internal state directory.

Current 1.0 canonical state files remain:

- `docs/module_state/coordination.md`
- `docs/module_state/product_manager.md`
- `docs/module_state/reanchor_detector_core.md`
- `docs/module_state/cli_core.md`
- `docs/module_state/templates_protocol.md`
- `docs/module_state/codex_skill.md`
- `docs/module_state/dogfood_validation.md`

Optional category split points live under:

```text
docs/module_state/<thread>/
```

Only create a thread-local memory directory when the thread actually needs
category-level files. Do not keep empty placeholder directories.

Suggested categories:

- `bugs.md`
- `style.md`
- `conventions.md`
- `aliases.md`
- `runbook.md`
- `key_memory.md`

Compatibility rule:

- The canonical `docs/module_state/*.md` files remain authoritative.
- Category files are detection and memory split points only until Coordination
  promotes them.
- Do not add `README.md` files inside internal memory directories unless that
  directory is intentionally user-facing.

## Migration Status

Status: `semantic mirrors initialized`

The following threads now have category-level Layer 3 mirrors:

- `docs/module_state/coordination/`
- `docs/module_state/product_manager/`
- `docs/module_state/reanchor_detector_core/`
- `docs/module_state/cli_core/`
- `docs/module_state/templates_protocol/`
- `docs/module_state/codex_skill/`
- `docs/module_state/dogfood_validation/`

Standard mirror categories:

- `current_state.md`
- `open_issues.md`
- `runbook.md`
- `history.md`

Product Manager also has pilot category files:

- `bugs.md`
- `style.md`
- `conventions.md`
- `aliases.md`
- `key_memory.md`

Still required before deleting `docs/module_state/*.md`:

- Coordination must decide that category-level Layer 3 files become
  authoritative.
- Each thread must verify no current state, open issue, runbook, or history
  entry was lost during the split.
- Thread prompts and reanchor guidance must point to the category-level files or
  a generated Layer 3 index.
- A compatibility redirect period must run before deletion.
