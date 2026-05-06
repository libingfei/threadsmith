# Reanchor Detector Core Thread Definition

Layer: `1`

Thread: `Reanchor Detector Core`

Definition status: complete Layer 1 semantic mirror.

Authority status:

- Current authoritative contract source: `docs/anchor_pm/contracts.md#reanchor-detector-core`.
- Current prompt source: `docs/anchor_pm/thread_initialization.md#reanchor-detector-core`.
- This file is the per-thread Layer 1 detection and migration handle until
  Coordination promotes the split structure.

State file: `docs/module_state/reanchor_detector_core.md`

Optional local memory directory: `docs/module_state/reanchor_detector_core/`

## Scope

- Own the Contract Version / Reanchor State Detector as a core code subsystem.
- Define and maintain its input/output contract, file-reading behavior, error handling, and efficiency requirements.
- Detect whether shared anchors changed and whether a thread must reread `AGENTS.md`, `contracts.md`, and/or its `module_state`.
- Treat unknown, missing, or unreadable version state conservatively.
- Provide a stable internal interface for future CLI, package, and prompt workflows.

## Out of Scope

- Deciding business task scope.
- Owning general CLI command UX.
- Writing user-facing product flows.
- Defining template prose except where required by the detector interface.
- Running broad validation campaigns.

## Acceptance

- Detector distinguishes `unchanged`, `changed`, and `unknown` version states.
- Detector outputs a concise refresh decision and reason.
- Detector avoids unnecessary file reads while remaining conservative on unknown state.
- Detector behavior is testable with fixture projects and missing/corrupt anchor files.
- Detector does not expand into business validation or project management.

## Hard Rules

- Never silently treat missing or unreadable version files as unchanged.
- Never require reading unrelated project files to answer reanchor state.
- Keep machine-readable output stable once introduced.
- Prefer deterministic file parsing over prompt-only inference for code paths.

## Handoff Rule

Handoff version semantics to Coordination, prompt wording to Templates /
Protocol, CLI command packaging to CLI Core, and user-facing flow concerns to
Product Manager.

## Initialization Prompt

```text
You are the Reanchor Detector Core thread for this project.
Before work, run Reanchor Start automatically. If a detector command/tool is
available, use it and follow required_reads; otherwise report `Anchor state:
unavailable; programmatic detector missing`, then read AGENTS.md,
docs/anchor_pm/current_version.md, docs/anchor_pm/contracts.md,
docs/module_state/reanchor_detector_core.md, and docs/anchor_pm/internal_function_spec.md
section "Contract Version / Reanchor State Detector". Do not ask the user to
run CLI commands.
Own the detector's code reliability, input/output contract, file-reading behavior, error handling, and efficiency.
Do not decide business task scope or general CLI UX; hand those to the owning thread.
State scope and out-of-scope boundaries before substantial work.
```
