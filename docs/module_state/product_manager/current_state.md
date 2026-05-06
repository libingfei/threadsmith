# Product Manager Current State

Layer: `3`

Thread: `Product Manager`

Status: compatibility mirror of `docs/module_state/product_manager.md`.

- Anchor PM 1.0 uses a Codex package-first integration flow.
- The main user-facing entrypoints for target projects are repository-root
  language-specific copy targets such as `ANCHOR_PM_INSTALL_PROMPT.en.md` and
  `ANCHOR_PM_INSTALL_PROMPT.zh.md`.
- Target project installation should produce a proposal before writing files.
- The installation proposal should behave like a concise confirmation page, not
  a technical audit log.
- Thread initialization prompts must be complete and copy-paste-ready; users
  should not fill in placeholders Codex can generate.
- Installer replies should follow the user's usual language, while generated
  project docs may remain English unless requested otherwise.
- `Thread Management` should remain available after installation for future
  thread additions, removals, renames, or prompt regeneration.
- Install prompts must not depend on the user naming the Codex conversation
  before the first message; thread naming can be optional when the client
  supports it.
- The current install conversation is the Thread Management entrypoint, not a
  target-project business specialist thread.
- Existing-project thread generation should be based on the target project's
  real modules, subsystems, and durable maintenance boundaries, not generic
  responsibility buckets.
- Ordinary target projects should not receive an Anchor PM internal
  `Coordination` thread. Generic `Implementation` or standalone `Validation`
  threads are fallback-only and require clear justification.
- Installation proposals should stay user-facing and decision-focused; verbose
  internal safety explanations and package execution details should remain
  hidden unless the user asks.
- User-visible thread names should use the same language as the selected
  install prompt. Technical terms may remain in their conventional language.
- Default install reply options should be limited to approve, adjust threads,
  and cancel. `Adjust AGENTS.md` should not appear as a default option; AGENTS
  conflicts should be handled as specific decisions only when needed.
- "Install docs only; do not update AGENTS.md" should not be a default approval
  option because it creates a confusing partial integration. If AGENTS handling
  is blocked, the installer should explain the consequence and ask for a
  specific merge/skip decision.
- README should state the exact approval reply and explain that post-install
  long-lived threads are created from `docs/anchor_pm/thread_initialization.md`.
- README should frame the primary value as modular specialist threads: narrower
  business scope per thread, less unrelated context, better Codex answer
  quality, and reanchor/module-state synchronization when paths, data,
  commands, or methods change.
- Public positioning should explain Anchor PM as organization design for AI
  coding: stable general AI capability is directed into focused specialist
  threads, with lightweight communication channels for dependency information.
- Public examples should stay generic and anonymized so they can recruit
  testers without exposing private source projects.
- README should stay compact: one-sentence positioning, attractive value points,
  install path, one example, and a feedback-oriented close.
- README should help first-time readers self-qualify before installation by
  separating `Best Fit / 适合谁` from `What You Get / 你会得到什么`, including
  when Anchor PM is probably unnecessary.
- README should state the public GitHub package source used by install prompts:
  `https://github.com/libingfei/threadsmith` with package directory
  `packages/anchor-pm-1.0-standard`.
- README should explicitly invite feedback about installation clarity,
  thread-boundary choice, and whether reanchor/module state reduces repeat
  explanations.
- Product work must continuously test whether Anchor PM improves real AI coding
  experience; if users struggle to understand/use it, or if evidence suggests it
  does not improve experience, this thread should report that clearly and
  recommend stopping or changing direction.
- Product Manager must not treat checklist completeness as product success.
  Install and workflow evaluations require a PM Review Gate that checks real
  Codex operation, user-visible clarity, project-specific thread semantics, and
  whether each option has meaningful user value.
- Internal function behavior should be documented separately from README in
  `docs/anchor_pm/internal_function_spec.md`, with module-level inputs, outputs,
  owners, acceptance criteria, and failure signals.
- The system should include a Contract Version / Reanchor State Detector that
  checks whether shared anchors changed at thread start and decides whether
  common knowledge files must be reread.
- Contract state detector design should stay small and explicit: state model,
  fingerprints, decision rules, minimal chat output, checkpoint persistence, and
  failure signals.
- Contract state detector monitoring scope should distinguish framework
  baseline, thread definition, cross-thread shared state, thread local memory,
  Anchor PM development-only files, and business files that should not trigger
  reanchor by default.
- Contract state detector monitoring should be modeled as a horizontally
  extensible four-layer tree with task-specific refresh profiles instead of one
  flat monitored-file list.
- Contract state detector monitoring should use a four-layer product model:
  Layer 0 framework baseline, Layer 1 thread definition, Layer 2 cross-thread
  shared state, and Layer 3 thread local memory.
- Closeout Knowledge Sync is a core Anchor PM workflow for every long-lived
  thread: before final response after substantial work, the thread must decide
  whether new or changed knowledge belongs in its own Layer 3 state, a Layer 2
  shared-state file/handoff, a Thread Management Layer 1 update request, or a
  framework-owner handoff. If nothing durable changed, it should state that no
  durable state update is needed.
- Reanchor Start and Closeout Knowledge Sync are symmetric lifecycle hooks:
  Reanchor Start is the pre-work read/refresh/boundary-confirmation side;
  Closeout Knowledge Sync is the pre-response write/propagation side. Anchor PM
  only evolves across conversation rounds when both hooks run consistently.
- MVP manual validation must test the full lifecycle, not only installation:
  project adoption, module/subsystem thread generation, thread prompt creation,
  Reanchor Start, scoped issue solving, handoff, Closeout Knowledge Sync,
  shared-state recovery, and repository restore.
- Before the user starts any public GitHub-based test, Product Manager should
  run a Pre-test GitHub Sync Gate: check for local changes, commit the intended
  changes, push to `origin/main`, verify remote HEAD, and tell the user which
  commit and GitHub links to test.
- Contract state detector file granularity should keep status checks cheap.
- Long-lived threads should run a periodic reanchor safety check every 10
  conversation rounds.
- Reanchor Detector Core development should use
  `docs/anchor_pm/00_framework_baseline/reanchor_module_io_spec.md` as the fixed
  programmatic I/O contract: request in, file fingerprint comparison, required
  reads, blockers, next action, and checkpoint update proposal out.
- Reanchor output should not dump file contents, chat history, or business data;
  it should return machine-readable refresh decisions and a short chat-facing
  summary.
- Reanchor should be productized as automatic Codex behavior, not a user action:
  before substantial work, Codex runs Reanchor Start, uses the detector if
  available, falls back to reading required anchors only as a degraded
  compatibility path if unavailable, and shows only a short anchor-state line.
- Programmatic reanchor is the target behavior: Codex should receive a
  machine-readable detector result and read only `required_reads`; automatic
  full anchor rereading is not considered a completed implementation.
- Current self-dogfood has a product gap: `anchorpm` is not callable in this
  environment, so Reanchor Start currently degrades to file rereads.
- Package release artifacts still contain older read-first reanchor wording; the
  owning Templates / Protocol thread must update package workflows and templates
  to match the programmatic/degraded-fallback distinction.
- This repository now has an initialized numbered Layer 0 through Layer 3
  structure.
- Internal anchor directories should not use `README.md` by default.
- Layer 1 semantic migration has started.
- Layer 2 semantic migration has started.
