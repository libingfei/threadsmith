# Product Manager To Reanchor Detector Core

Layer: `2`

Source thread: `Product Manager`

Target thread: `Reanchor Detector Core`

Status: active shared dependency.

## Confirmed Product Requirements

- Reanchor state uses a four-layer model:
  - Layer 0: framework baseline.
  - Layer 1: thread definition.
  - Layer 2: cross-thread shared state.
  - Layer 3: thread local memory.
- Ordinary specialist threads confirm all registered layer states at startup.
- Confirming status is not the same as reading full file contents.
- Threads should read only changed, unknown, unreadable, inbound, or
  task-relevant files.
- Layer 1 should support one thread-definition detection handle per thread.
- Layer 2 should support sparse directed dependency files rather than eagerly
  generating every possible `N * (N - 1)` channel.
- Layer 3 may split local memory by category when the single state file becomes
  too large or volatile.
- Long-lived threads should run a periodic reanchor check every 10 conversation
  rounds.
- Closeout Knowledge Sync should be supported as a first-class operation:
  every long-lived thread needs to classify closeout changes into Layer 3 local
  memory, Layer 2 shared state/handoff, Thread Management Layer 1 update
  request, framework-owner handoff, or no durable update.
- Reanchor Detector Core should model `ordinary_thread_start` and
  `ordinary_thread_closeout` as symmetric lifecycle operations: start imports
  changed anchor knowledge before work; closeout exports newly durable knowledge
  before reply.
- Reanchor should be triggered by Codex before substantial work; users should
  not be asked to run a detector or CLI command manually.
- The target behavior is programmatic anchoring, not automatic full rereading:
  a detector command/tool should return machine-readable `required_reads`,
  `blocked_by`, `next_action`, and checkpoint update data.
- If no detector command/tool is callable, Reanchor Start should be reported as
  degraded/unavailable before using the fallback read path.
- Reanchor Detector Core development should use the fixed input/output contract
  in `docs/anchor_pm/00_framework_baseline/reanchor_module_io_spec.md`.
- The detector should output required reads and checkpoint updates; it should
  not dump file contents or chat history.
- Closeout planning should be machine-readable and should never store chat
  history in checkpoint state.

## Target Next Step

Use this as product input when designing detector interfaces, fixture cases, and
machine-readable outputs. Use
`docs/anchor_pm/00_framework_baseline/reanchor_module_io_spec.md` as the
engineering I/O contract. The next product-relevant gap is a callable runtime
entrypoint that lets Codex receive `ReanchorResult` instead of rereading anchor
files itself. Implementation details still belong to the Reanchor Detector Core
thread.
