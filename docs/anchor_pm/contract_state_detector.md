# Contract State Detector

Purpose: decide whether a thread can continue with its current shared context,
or must refresh project anchors before work.

This is a core Anchor PM module. It protects the main product promise: focused
specialist threads with fresh enough shared knowledge and low context cost.

Engineering I/O contract:

- `docs/anchor_pm/00_framework_baseline/reanchor_module_io_spec.md`

## Principle

Do not reread everything by default.

Do not trust stale thread memory by default.

Read only the anchors whose state changed, and treat unknown state as changed.

## Trigger Model

Reanchor should feel automatic to users.

- Codex runs a lightweight `Anchor Gate` before substantial work in a long-lived
  Anchor PM thread.
- The gate checks for explicit durable user corrections, rule updates,
  preferences, or cross-thread facts, then runs Reanchor Start only as far as
  needed for the task.
- The user should not be asked to run a CLI command manually.
- If a detector command/tool is available, Codex runs it and follows
  `required_reads`, `blocked_by`, and `next_action`.
- If no detector command/tool is available, Codex reports a degraded/unavailable
  anchor state, then performs the compatibility fallback: read the project
  version, thread contracts, and current thread state file.
- Chat output should stay silent for unchanged anchor checks and show only a
  minimal status when changed, unknown, blocked, conflicting, or degraded.

## Lifecycle Symmetry

Anchor PM has two lightweight gates:

- `Anchor Gate` runs before substantial work. It handles same-turn user deltas,
  reads or refreshes changed knowledge, and confirms the current boundary.
- `Knowledge Sync Gate` runs before the final response. It writes back local
  durable knowledge, updates shared state, or produces handoffs only when needed.

The detector exists to make both gates cheap and deterministic. Over many
conversation rounds, the first gate protects newest user intent and imports
relevant anchor changes, while the second exports durable new knowledge. If gate
handling crowds out the actual task, the product has failed its UX budget.

## State Model

The detector compares the thread's last known anchor state with the current
project anchor state.

The state model is a logical tree, not a whole-repository scan. Each layer has a
different change trigger, update frequency, and blast radius.

```text
Project Anchor Root
  Layer 0: Framework Baseline
    docs/anchor_pm/00_framework_baseline/baseline.md
    package templates and default detector protocol
  Layer 1: Thread Definition
    AGENTS.md Anchor PM discovery section
    docs/anchor_pm/contracts.md
    docs/anchor_pm/thread_initialization.md
    docs/anchor_pm/01_thread_definitions/<thread>.md
  Layer 2: Cross-Thread Shared State
    docs/anchor_pm/current_version.md
    docs/anchor_pm/interaction_guide.md, when workflow behavior changed
    docs/anchor_pm/review_log.md
    handoff-named docs/module_state/*.md
    docs/anchor_pm/02_shared_state/<source>__to_<target>.md
  Layer 3: Thread Local Memory
    docs/anchor_pm/03_thread_local_memory.md
    docs/module_state/<current-thread>.md
    docs/module_state/<thread>/<category>.md
```

Default status checks for a normal non-thread-management long-lived thread:

- Layer 0 framework baseline version.
- Layer 1 thread definition version.
- Layer 2 files relevant to the current task.
- Layer 3 current `docs/module_state/<thread>.md`.

Confirming a layer changed is not the same as reading that full layer into the
conversation. Ordinary threads should cheaply confirm all layer states, then
read only the layers required by the refresh decision.

Each tracked anchor should have a lightweight fingerprint:

- version id when available;
- content hash when automated;
- mtime only as a weak fallback;
- `unknown` when no reliable prior state exists.

## Monitoring Scope

The detector must not become a whole-repository scanner. It monitors anchor
documents, not business files.

### Layer 0: Framework Baseline

This layer is common to projects that installed the same Anchor PM package or
protocol version.

Typical contents:

- `docs/anchor_pm/00_framework_baseline/baseline.md`;
- Anchor PM package version;
- template and prompt protocol version;
- detector interface version;
- default reanchor and handoff rules shipped by Anchor PM.

Change trigger:

- Anchor PM itself is upgraded or reinstalled in the target project.

Refresh behavior:

- Ordinary project threads confirm this layer's status at conversation start.
- `Thread Management`, installer, or framework-upgrade flow checks this layer.
- If this layer changes, affected projects should rerun the approved upgrade or
  regeneration path instead of silently changing thread behavior.

### Layer 1: Thread Definition

This layer defines the project's thread topology: which threads exist, what each
thread owns, and how users create those threads.

Typical contents:

- `AGENTS.md`
  - only the Anchor PM discovery section and project-level AI rules relevant to
    finding the anchors;
- `docs/anchor_pm/contracts.md`
  - thread scope, out-of-scope boundaries, handoff rules;
- `docs/anchor_pm/thread_initialization.md`
  - copy-paste-ready prompts for creating or restoring long-lived threads.
- `docs/anchor_pm/01_thread_definitions/<thread>.md`
  - per-thread detection handle for status checks.

Granularity:

- one thread-definition file per thread, for example:
  `docs/anchor_pm/01_thread_definitions/<thread>.md`;
- a small index can list thread names and definition fingerprints;
- a thread should refresh only its own definition unless thread topology
  changed.

Change trigger:

- new thread;
- removed thread;
- renamed thread;
- responsibility or boundary change;
- approved change to how threads should be created.

Refresh behavior:

- `Thread Management` owns this layer.
- Ordinary threads confirm this layer's status at conversation start.
- If this layer changed, the ordinary thread should refresh its relevant
  definition before continuing or ask `Thread Management` to resolve the
  change.
- If an ordinary thread finds that its task no longer fits its definition, it
  should produce a handoff or ask `Thread Management` to update the definition.

### Layer 2: Cross-Thread Shared State

This layer exists only for information that affects more than one thread.

Typical contents:

- `docs/anchor_pm/current_version.md`
  - project-level shared-anchor version or coordination change notice;
- `docs/anchor_pm/interaction_guide.md`
  - only when workflow behavior changed;
- `docs/anchor_pm/review_log.md`
  - repeated failures, validation findings, or version bugs with cross-thread
    impact;
- handoff-named `docs/module_state/*.md`
  - only when another thread's state contains dependency information needed by
    the current task.
- `docs/anchor_pm/02_shared_state/<source>__to_<target>.md`
  - sparse directed dependency channel.

Granularity:

- use directed shared-state files for cross-thread dependencies, for example:
  `docs/anchor_pm/02_shared_state/<source>__to_<target>.md`;
- conceptually, a project with `N` threads can have up to `N * (N - 1)`
  directed dependency channels;
- in practice, create sparse files only for dependencies that actually exist;
- a target thread should inspect only inbound files addressed to it, plus
  handoff-named files.

Change trigger:

- a path, command, interface, data dependency, workflow, decision, or constraint
  changed in a way that another thread must know.

Refresh behavior:

- Non-thread-management threads check this layer at conversation start.
- The check should be scoped to the current task and named handoffs.
- A Layer 2 change can require reading a specific other thread's state, but
  should not force all module states into context.

### Layer 3: Thread Local Memory

This layer belongs to one thread.

Typical contents:

- current `docs/module_state/<thread>.md`
  - current state;
  - open issues;
  - runbook;
  - known thread-local bugs;
  - key memory points;
  - history/notes.
- optional category files under `docs/module_state/<thread>/`.
- `docs/anchor_pm/03_thread_local_memory.md`
  - Layer 3 directory and compatibility index.

Granularity:

- split local memory by logical category only when the single state file becomes
  too large or too volatile;
- possible categories:
  - `bugs.md`
  - `style.md`
  - `conventions.md`
  - `aliases.md`
  - `runbook.md`
  - `key_memory.md`
- the detector should refresh only changed local-memory categories.

Change trigger:

- this thread learned a durable fact;
- this thread found a recurring error;
- this thread changed its local runbook;
- this thread needs to resume after context loss.

Refresh behavior:

- Non-thread-management threads check this layer at conversation start.
- At closeout, the thread updates this layer when the conversation changed its
  durable local state.

## Non-Thread-Management Thread Flow

For ordinary specialist threads, conversation start should stay cheap:

```text
Start:
  1. Identify thread identity.
  2. Confirm Layer 0, Layer 1, Layer 2, and Layer 3 status.
  3. If Layer 0 changed, stop ordinary work and use framework-upgrade or
     Thread Management flow.
  4. If Layer 1 changed, refresh this thread's definition before work.
  5. If Layer 2 changed, read the relevant shared state or named handoff.
  6. If Layer 3 changed, read this thread's local memory.
  7. Continue within the confirmed Layer 1 definition.
```

Ordinary threads should confirm every layer at startup, but they should not
apply Layer 0 or Layer 1 changes themselves unless the current thread is the
owning management thread. Confirmation protects correctness; selective reading
protects context size.

At closeout:

```text
Knowledge Sync Gate:
  1. Run before the final response after substantial work.
  2. If this thread learned durable local information, update Layer 3.
  3. If this thread changed something other threads depend on, update Layer 2 or
     produce a handoff that names the affected thread state.
  4. If the work implies thread-definition changes, hand off to
     Thread Management instead of editing Layer 1 directly.
  5. If the work implies framework-level behavior changes, hand off to
     Coordination or the owning framework thread.
  6. If no durable or shared knowledge changed, keep the gate silent unless the
     user asked for status.
```

This preserves the main product promise: ordinary specialist threads keep a
small context window while still receiving cross-thread dependency updates when
they matter.

## Periodic Reanchor

File fingerprints reduce routine context cost, but long-lived threads still
need a periodic safety check.

Rule:

- Every 10 conversation rounds in a long-lived thread, run a periodic reanchor
  check across all registered layers.
- A conversation round means one user request plus the assistant's completed
  response.
- The periodic check confirms Layer 0 through Layer 3 status and any registered
  Anchor PM development layer for this repository.
- It should refresh fingerprints for all registered anchor files.
- It should read file contents only for changed, unknown, unreadable, or
  task-relevant layers.

Purpose:

- catch missed dependency updates;
- catch stale thread definitions;
- prevent local memory from drifting too far from project anchors;
- keep the cost bounded by file granularity instead of full-document rereads.

### Anchor PM Development Files

These are monitored by Anchor PM's own development threads, not by ordinary
target projects unless the target project is Anchor PM itself:

- `README.md`
- `ANCHOR_PM_INSTALL_PROMPT*.md`
- `PRODUCT_PRINCIPLES.md`
- `MVP_SPEC.md`
- `ANCHOR_PM_FRAMEWORK_DESIGN.md`
- `docs/anchor_pm/internal_function_spec.md`
- `docs/anchor_pm/contract_state_detector.md`
- package files under `packages/anchor-pm-1.0*/`

This section exists so Anchor PM can dogfood its own development. It should not
be installed as a default burden for downstream target projects.

### Non-Anchor Project Files

These may matter to a task, but they are not shared-anchor triggers:

- business source files;
- generated data;
- logs;
- build outputs;
- model artifacts;
- temporary files;
- unrelated docs not referenced by Anchor PM anchors.

Business files should be read because the task needs them, not because the
reanchor detector changed state.

## Refresh Profiles

The detector should choose a refresh profile from the tree instead of returning
only a flat file list.

```text
ordinary_thread_start:
  Confirm Layer 0, Layer 1, Layer 2, and Layer 3 status
  Read Layer 1 only if thread definition changed
  Read Layer 2 only if relevant shared state changed
  Read Layer 3 only if local memory changed

periodic_reanchor:
  Every 10 conversation rounds
  Confirm all registered layer fingerprints
  Read only changed, unknown, unreadable, or task-relevant layers

ordinary_thread_closeout:
  Layer 3: docs/module_state/<current-thread>.md, if local memory changed
  Layer 2: shared notice or handoff, if other threads are affected

thread_management:
  Layer 1: AGENTS.md Anchor PM discovery section
  Layer 1: docs/anchor_pm/contracts.md
  Layer 1: docs/anchor_pm/thread_initialization.md
  Layer 1: docs/anchor_pm/01_thread_definitions/*.md
  Layer 2: affected shared state

framework_upgrade:
  Layer 0: docs/anchor_pm/00_framework_baseline/baseline.md
  Layer 1: regenerated thread definitions, if approved
  Layer 2: shared upgrade notice

cross_thread_handoff:
  Layer 2: named shared dependency
  Layer 3: source or target thread state named by the handoff

detector_development:
  Anchor PM development files
  docs/module_state/<current-thread>.md

product_or_install_flow:
  Anchor PM development files
  README.md
  ANCHOR_PM_INSTALL_PROMPT*.md
  docs/anchor_pm/thread_management_install_prompt.md
  docs/anchor_pm/interaction_guide.md

validation:
  Layer 2: docs/anchor_pm/review_log.md
  Layer 2: docs/anchor_pm/reports/*
  Layer 3: docs/module_state/dogfood_validation.md
```

Profiles are additive. A task can activate more than one profile, but the
detector should still explain why each non-core layer was included.

## Not Monitored By The Detector

The detector should not monitor:

- whole directories by default;
- package internals in downstream target projects;
- user business files;
- generated artifacts;
- chat transcripts or long summaries.

## Inputs

- Thread identity.
- Last known anchor fingerprints for this thread, if available.
- Current anchor fingerprints.
- Current user task.
- Optional explicit user request to reanchor fully.

## Outputs

```text
refresh_required: yes | no | unknown
changed_nodes:
  - framework_baseline
  - thread_definition
  - cross_thread_shared_state
  - thread_local_memory
  - anchor_pm_development_files
refresh_profile:
  - ordinary_thread_start
  - periodic_reanchor
  - ordinary_thread_closeout
  - thread_management
  - framework_upgrade
  - cross_thread_handoff
  - detector_development
  - product_or_install_flow
  - validation
refresh_scope:
  - layer_0
  - layer_1
  - layer_2
  - layer_3
  - anchor_pm_development_files
  - full
  - none
required_reads:
reason:
risk_if_skipped:
next_action:
```

## Decision Rules

1. If prior state is missing, set `refresh_required=unknown` and
   `refresh_scope=full` for initial setup. After setup, ordinary starts should
   confirm all layer states and read only changed or relevant layers.
2. If Layer 0 changed, activate `framework_upgrade`; ordinary specialist
   threads should not apply that change directly.
3. If Layer 1 changed, activate `thread_management`; ordinary specialist
   threads should continue only after the relevant definition is refreshed or
   confirmed.
4. If Layer 2 changed, read only the shared state or handoff-named module state
   relevant to the current task.
5. If Layer 3 changed, refresh the current thread's local state.
6. If another thread's Layer 3 changed without Layer 2 notice or handoff, do not
   read it by default.
7. If a non-thread-management thread changed only local knowledge, update Layer
   3 at closeout.
8. If a non-thread-management thread changed information other threads depend
   on, update Layer 2 or produce a handoff at closeout.
9. If the user task crosses the current thread boundary, run reanchor and then
   handoff instead of silently expanding scope.
10. If a tracked anchor is unreadable, treat it as changed and report the read
    failure.
11. Every 10 conversation rounds, run `periodic_reanchor`: refresh fingerprints
    across all registered layers, then read only changed, unknown, unreadable,
    or task-relevant layers.

## Minimal Chat Output

Only show the result needed for the current thread.

Changed:

```text
Anchor state: changed
Profile: ordinary_thread_start
Refresh: layer_1, layer_2, layer_3
Reason: thread definition changed, shared dependency notice changed, and this
thread's local state changed.
Next: refresh this thread's definition, then read the named shared notice and
docs/module_state/<thread>.md.
```

Unchanged:

```text
Anchor state: unchanged
Refresh: none
Next: continue within current Layer 1 thread definition.
```

## Persistence

The detector needs a small per-thread checkpoint.

Preferred future location:

```text
docs/anchor_pm/.state/<thread>.anchor_state.json
```

Minimum fields:

```json
{
  "thread": "Implementation",
  "checked_at": "2026-05-06T00:00:00Z",
  "anchors": {
    "AGENTS.md": "sha256:...",
    "docs/anchor_pm/current_version.md": "package-first-v1.0",
    "docs/anchor_pm/contracts.md": "sha256:...",
    "docs/module_state/implementation.md": "sha256:..."
  }
}
```

Do not store chat history in this checkpoint.

## Implementation Forms

Detector command/tool form:

- `anchorpm reanchor --thread <name> --operation ordinary_thread_start --json`
  or equivalent tool call;
- computes fingerprints;
- compares against checkpoint;
- prints machine-readable refresh decision;
- optionally updates checkpoint after successful reanchor.

This is the intended product behavior. It avoids reloading full anchor contents
unless the detector returns those files in `required_reads`.

Protocol-triggered degraded fallback form:

- `AGENTS.md` and thread prompts require Codex to run `Reanchor Start`
  automatically.
- If no detector command/tool is available, Codex reads `current_version.md`,
  `contracts.md`, and the current thread state file itself.
- This fallback must be automatic; users should not be asked to run CLI
  commands.
- This fallback is not programmatic anchoring. It should be reported as
  degraded/unavailable and treated as a temporary compatibility path.

## Failure Signals

- Thread continues after shared anchors changed.
- Every thread rereads every anchor on every message.
- Detector stores large summaries or chat history.
- Detector hides unreadable files.
- Detector decides business scope instead of refresh scope.
- Programmatic detector entrypoint is missing, so every session still relies on
  fallback rereads.
