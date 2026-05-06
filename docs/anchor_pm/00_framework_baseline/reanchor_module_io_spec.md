# Reanchor Module I/O Spec

Layer: `0`

Status: engineering contract draft.

Owner:

- Reanchor Detector Core owns implementation behavior.
- Coordination owns version semantics and promotion decisions.
- Templates / Protocol owns prompt-facing wording.
- CLI Core may package this interface later without changing the core semantics.

Purpose: define a fixed, programmatic input/output contract for the reanchor
module using the Layer 0 through Layer 3 structure.

The module is based on file update state. It confirms whether registered anchor
files changed, plans the minimum required reads, and produces a stable result
that a prompt, package workflow, or future CLI can consume.

## Runtime Invocation Requirement

Programmatic anchoring requires a callable runtime entrypoint. The preferred
shape is:

```text
anchorpm reanchor --thread <thread_id> --operation ordinary_thread_start --json
```

An equivalent Codex tool call or project-local executable is acceptable if it
returns the same `ReanchorResult` schema. If no callable entrypoint exists,
Codex may use a documented fallback that reads required anchors directly, but
that state must be reported as degraded and must not be treated as completed
programmatic anchoring.

## Boundary

The reanchor module does:

- load or receive a registered anchor-file list;
- compare current fingerprints with the prior checkpoint;
- classify changed, unchanged, unknown, missing, and unreadable files;
- map file changes to Layer 0 through Layer 3 refresh decisions;
- decide which files must be read before work continues;
- produce checkpoint updates after a successful refresh.

The reanchor module does not:

- decide business task scope;
- read business source files by default;
- edit anchor documents;
- generate handoff prose;
- store chat history;
- silently treat missing or unreadable required files as unchanged.

## Layer Model

```text
Layer 0: Framework Baseline
  docs/anchor_pm/00_framework_baseline/baseline.md

Layer 1: Thread Definition
  docs/anchor_pm/01_thread_definitions/index.md
  docs/anchor_pm/01_thread_definitions/<thread>.md

Layer 2: Cross-Thread Shared State
  docs/anchor_pm/02_shared_state/index.md
  docs/anchor_pm/02_shared_state/<source>__to_<target>.md
  docs/anchor_pm/02_shared_state/project-level mirrors

Layer 3: Thread Local Memory
  docs/anchor_pm/03_thread_local_memory.md
  docs/module_state/<thread>.md
  docs/module_state/<thread>/<category>.md
```

Current 1.0 compatibility files may still be registered until Coordination
promotes the split:

- `AGENTS.md`
- `docs/anchor_pm/current_version.md`
- `docs/anchor_pm/contracts.md`
- `docs/anchor_pm/thread_initialization.md`
- `docs/module_state/*.md`

## Core Types

### Layer Id

```text
layer_0_framework_baseline
layer_1_thread_definition
layer_2_cross_thread_shared_state
layer_3_thread_local_memory
anchor_pm_development
```

Implementations may add extension layers, but must preserve these names.

### File Status

```text
unchanged
changed
unknown
missing_required
missing_optional
unreadable
invalid_path
```

Required missing, unreadable, invalid, or unknown files must force a conservative
refresh decision.

### Fingerprint

Use this priority order:

1. explicit version id, when the file provides one;
2. content hash;
3. mtime plus size as a weak fallback;
4. `unknown`.

Recommended hash algorithm: `sha256`.

## Input: ReanchorRequest

The module input should be a single structured request.

```json
{
  "schema_version": "anchorpm.reanchor.request.v1",
  "operation": "ordinary_thread_start",
  "project_root": "/absolute/project/root",
  "thread": {
    "id": "product_manager",
    "name": "Product Manager",
    "is_thread_management": false
  },
  "conversation": {
    "round_count": 7,
    "force_periodic": false,
    "user_task_hint": "short optional task label"
  },
  "checkpoint": {
    "path": "docs/anchor_pm/.state/product_manager.anchor_state.json",
    "state": null
  },
  "registry": {
    "mode": "default_layered",
    "files": []
  },
  "handoff": {
    "named_files": [],
    "source_thread": null,
    "target_thread": null
  },
  "options": {
    "hash_algorithm": "sha256",
    "read_file_contents": false,
    "include_compatibility_files": true,
    "max_required_reads": 20
  }
}
```

Required fields:

- `schema_version`
- `operation`
- `project_root`
- `thread.id`
- `checkpoint.path` or `checkpoint.state`
- `registry.mode`

Optional fields must have deterministic defaults.

## Operations

```text
ordinary_thread_start
periodic_reanchor
ordinary_thread_closeout
thread_management
framework_upgrade
cross_thread_handoff
detector_development
validation
```

Rules:

- `ordinary_thread_start` confirms Layer 0 through Layer 3 status, then reads
  only changed, unknown, unreadable, inbound, or task-relevant files.
- `periodic_reanchor` runs every 10 conversation rounds and confirms all
  registered layer fingerprints.
- `ordinary_thread_closeout` plans Layer 3 updates for local durable memory and
  Layer 2 updates or handoffs for cross-thread impact.
- `thread_management` may inspect all Layer 1 thread definitions.
- `framework_upgrade` may inspect Layer 0 and approved Layer 1 regeneration
  outputs.
- `cross_thread_handoff` inspects named Layer 2 dependency files and named Layer
  3 source/target files.

## Lifecycle Symmetry

`ordinary_thread_start` and `ordinary_thread_closeout` are paired lifecycle
operations.

- `ordinary_thread_start` is the read-side hook before work: compare
  fingerprints, return changed/unknown anchors, and tell Codex which files must
  be read before substantial reasoning.
- `ordinary_thread_closeout` is the write-side hook before final response:
  classify durable changes into Layer 3 local memory, Layer 2 shared state,
  Thread Management Layer 1 requests, framework-owner handoffs, or no durable
  update.

The system only improves over long conversations if both hooks run. Start
without closeout refreshes stale context but never captures new corrections.
Closeout without start risks preserving conclusions made from stale anchors.

## Registry File Entry

Each registered anchor file should use this shape.

```json
{
  "id": "thread_definition.product_manager",
  "layer": "layer_1_thread_definition",
  "path": "docs/anchor_pm/01_thread_definitions/product_manager.md",
  "required": true,
  "owner_thread": "product_manager",
  "source_thread": null,
  "target_thread": null,
  "category": "definition",
  "check_mode": "hash",
  "read_policy": "content_if_changed",
  "compatibility_source": "docs/anchor_pm/contracts.md"
}
```

Fields:

- `id`: stable machine id.
- `layer`: one of the core layer ids or an extension layer.
- `path`: repository-relative path under `project_root`.
- `required`: whether missing/unreadable state blocks normal continuation.
- `owner_thread`: owning thread id, if any.
- `source_thread` and `target_thread`: used for Layer 2 directed channels.
- `category`: category such as `baseline`, `definition`, `shared_notice`,
  `current_state`, `bugs`, `runbook`, or `history`.
- `check_mode`: `version`, `hash`, `mtime_size`, or `exists`.
- `read_policy`: `never`, `content_if_changed`, `content_if_relevant`, or
  `always_for_owner`.
- `compatibility_source`: old authoritative path, if this file is a mirror.

## Default Registry Discovery

For `ordinary_thread_start`, the default registry should include:

- Layer 0:
  - `docs/anchor_pm/00_framework_baseline/baseline.md`
  - compatibility: `docs/anchor_pm/current_version.md`
- Layer 1:
  - `docs/anchor_pm/01_thread_definitions/index.md`
  - `docs/anchor_pm/01_thread_definitions/<thread>.md`
  - compatibility: `docs/anchor_pm/contracts.md`
- Layer 2:
  - `docs/anchor_pm/02_shared_state/index.md`
  - project-level shared mirrors in `docs/anchor_pm/02_shared_state/`
  - inbound directed files matching `*__to_<thread>.md`
  - handoff-named files
  - compatibility: `docs/anchor_pm/current_version.md`,
    `docs/anchor_pm/interaction_guide.md`, `docs/anchor_pm/review_log.md`,
    `docs/anchor_pm/simplification.md`
- Layer 3:
  - `docs/anchor_pm/03_thread_local_memory.md`
  - `docs/module_state/<thread>.md`
  - files under `docs/module_state/<thread>/`, when present

For `thread_management`, include all files under
`docs/anchor_pm/01_thread_definitions/`.

For `validation`, include validation-specific Layer 2 reports only when the
operation or task asks for validation.

## Output: ReanchorResult

The module output must be stable and machine-readable.

```json
{
  "schema_version": "anchorpm.reanchor.result.v1",
  "operation": "ordinary_thread_start",
  "thread_id": "product_manager",
  "checked_at": "2026-05-06T00:00:00Z",
  "anchor_state": "changed",
  "periodic_due": false,
  "changed_layers": [
    "layer_2_cross_thread_shared_state"
  ],
  "file_statuses": [
    {
      "id": "shared.product_manager_to_reanchor_detector_core",
      "layer": "layer_2_cross_thread_shared_state",
      "path": "docs/anchor_pm/02_shared_state/product_manager__to_reanchor_detector_core.md",
      "previous_fingerprint": "sha256:old",
      "current_fingerprint": "sha256:new",
      "status": "changed",
      "error": null
    }
  ],
  "required_reads": [
    {
      "path": "docs/anchor_pm/02_shared_state/product_manager__to_reanchor_detector_core.md",
      "layer": "layer_2_cross_thread_shared_state",
      "reason": "inbound shared dependency changed",
      "read_mode": "full",
      "priority": "high"
    }
  ],
  "blocked_by": [],
  "next_action": "read_required_files_then_continue",
  "checkpoint_update": {
    "write_allowed": false,
    "path": "docs/anchor_pm/.state/product_manager.anchor_state.json",
    "state": {
      "schema_version": "anchorpm.reanchor.checkpoint.v1",
      "thread_id": "product_manager",
      "round_count": 8,
      "anchors": []
    }
  },
  "minimal_chat_output": {
    "anchor_state": "changed",
    "refresh": "layer_2",
    "next": "read required Layer 2 shared dependency before work"
  }
}
```

Required result fields:

- `schema_version`
- `operation`
- `thread_id`
- `checked_at`
- `anchor_state`
- `changed_layers`
- `file_statuses`
- `required_reads`
- `blocked_by`
- `next_action`
- `checkpoint_update`
- `minimal_chat_output`

## Result Enums

### anchor_state

```text
unchanged
changed
unknown
blocked
```

### next_action

```text
continue
read_required_files_then_continue
run_framework_upgrade_flow
refresh_thread_definition
handoff_to_thread_management
perform_cross_thread_handoff
fail_safe_full_reanchor
stop_for_read_error
```

### read_mode

```text
metadata_only
full
section
generated_summary
```

`generated_summary` is allowed only when a deterministic summarizer exists.
Prompt-only summarization must not be treated as a stable machine output.

## Checkpoint

Preferred checkpoint path:

```text
docs/anchor_pm/.state/<thread>.anchor_state.json
```

Checkpoint shape:

```json
{
  "schema_version": "anchorpm.reanchor.checkpoint.v1",
  "project_id": "anchor_pm_framework",
  "thread_id": "product_manager",
  "updated_at": "2026-05-06T00:00:00Z",
  "round_count": 8,
  "last_operation": "ordinary_thread_start",
  "anchors": [
    {
      "id": "layer0.baseline",
      "layer": "layer_0_framework_baseline",
      "path": "docs/anchor_pm/00_framework_baseline/baseline.md",
      "fingerprint": "sha256:...",
      "fingerprint_type": "sha256"
    }
  ]
}
```

Checkpoint rules:

- Store fingerprints, not chat history.
- Store round count for periodic reanchor.
- Do not store business data.
- Do not update checkpoint until required reads have been completed or the user
  explicitly accepts the risk.

## Closeout Input

Every substantial long-lived thread should run Closeout Knowledge Sync before
the final response. Closeout uses the same request envelope plus
durable-change events. If no durable or shared knowledge changed, the caller may
record no events and report that no closeout state update is needed.

```json
{
  "schema_version": "anchorpm.reanchor.request.v1",
  "operation": "ordinary_thread_closeout",
  "thread": {
    "id": "product_manager"
  },
  "closeout_events": [
    {
      "type": "local_memory_changed",
      "layer": "layer_3_thread_local_memory",
      "target_path": "docs/module_state/product_manager/key_memory.md",
      "reason": "durable product principle changed"
    },
    {
      "type": "cross_thread_dependency_changed",
      "layer": "layer_2_cross_thread_shared_state",
      "source_thread": "product_manager",
      "target_thread": "reanchor_detector_core",
      "target_path": "docs/anchor_pm/02_shared_state/product_manager__to_reanchor_detector_core.md",
      "reason": "detector I/O contract requirements changed"
    }
  ]
}
```

Closeout result adds:

```json
{
  "required_updates": [
    {
      "path": "docs/module_state/product_manager/key_memory.md",
      "layer": "layer_3_thread_local_memory",
      "update_type": "local_memory",
      "owner_thread": "product_manager"
    },
    {
      "path": "docs/anchor_pm/02_shared_state/product_manager__to_reanchor_detector_core.md",
      "layer": "layer_2_cross_thread_shared_state",
      "update_type": "cross_thread_dependency",
      "owner_thread": "product_manager",
      "target_thread": "reanchor_detector_core"
    }
  ]
}
```

The module may plan required updates, but writing prose remains the owning
thread's responsibility unless a separate writer tool is explicitly introduced.

## Error Strategy

- Missing required file: `anchor_state=blocked`, add `blocked_by`, recommend
  fail-safe refresh or Thread Management.
- Missing optional file: mark `missing_optional`, continue unless named by
  handoff.
- Unreadable file: mark `unreadable`, require refresh or stop for read error.
- Invalid path outside `project_root`: mark `invalid_path`, block.
- Missing checkpoint: mark prior state `unknown`; require first-run refresh,
  then create checkpoint after required reads.
- Unknown fingerprint: treat as changed.

## Extensibility

The interface is extensible through data, not new ad hoc rules:

- add new file categories through registry entries;
- add new directed dependencies through Layer 2 files;
- add new threads through Layer 1 definitions;
- add extension layers only when the four core layers cannot express the need;
- preserve `schema_version` and add fields instead of changing field meanings.

## Minimal User-Facing Output

The programmatic output may be detailed. Chat output should stay short:

```text
Anchor state: changed
Refresh: layer_2
Reason: inbound shared dependency changed.
Next: read docs/anchor_pm/02_shared_state/product_manager__to_reanchor_detector_core.md before work.
```

Do not expose fingerprint tables or registry dumps unless the user asks for
debug details.
