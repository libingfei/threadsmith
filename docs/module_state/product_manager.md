# Product Manager Thread State

## Contract

Thread contract: `Product Manager` in `docs/anchor_pm/contracts.md`

## Current State

- Anchor PM 1.0 uses a Codex package-first integration flow.
- The main user-facing entrypoints for target projects are repository-root language-specific copy targets such as `ANCHOR_PM_INSTALL_PROMPT.en.md` and `ANCHOR_PM_INSTALL_PROMPT.zh.md`.
- Target project installation should produce a proposal before writing files.
- The installation proposal should behave like a concise confirmation page, not a technical audit log.
- Thread initialization prompts must be complete and copy-paste-ready; users should not fill in placeholders Codex can generate.
- Installer replies should follow the user's usual language, while generated project docs may remain English unless requested otherwise.
- `Thread Management` should remain available after installation for future thread additions, removals, renames, or prompt regeneration.
- README should state the exact approval reply and explain that post-install long-lived threads are created from `docs/anchor_pm/thread_initialization.md`.
- README should frame the primary value as modular specialist threads: narrower business scope per thread, less unrelated context, better Codex answer quality, and reanchor/module-state synchronization when paths, data, commands, or methods change.
- Public positioning should explain Anchor PM as organization design for AI coding: stable general AI capability is directed into focused specialist threads, with lightweight communication channels for dependency information.
- Public examples should stay generic and anonymized so they can recruit testers without exposing private source projects.
- README should stay compact: one-sentence positioning, attractive value points, install path, one example, and a feedback-oriented close.
- README should help first-time readers self-qualify before installation by separating `Best Fit / 适合谁` from `What You Get / 你会得到什么`, including when Anchor PM is probably unnecessary.
- README should not imply public install is fully distribution-ready while install prompts still contain a local development package path.
- README should explicitly invite feedback about installation clarity, thread-boundary choice, and whether reanchor/module state reduces repeat explanations.
- Product work must continuously test whether Anchor PM improves real AI coding experience; if users struggle to understand/use it, or if evidence suggests it does not improve experience, this thread should report that clearly and recommend stopping or changing direction.
- Internal function behavior should be documented separately from README in `docs/anchor_pm/internal_function_spec.md`, with module-level inputs, outputs, owners, acceptance criteria, and failure signals.
- The system should include a Contract Version / Reanchor State Detector that checks whether shared anchors changed at thread start and decides whether common knowledge files must be reread.
- Contract state detector design should stay small and explicit: state model, fingerprints, decision rules, minimal chat output, checkpoint persistence, and failure signals.
- Contract state detector monitoring scope should distinguish framework baseline, thread definition, cross-thread shared state, thread local memory, Anchor PM development-only files, and business files that should not trigger reanchor by default.
- Contract state detector monitoring should be modeled as a horizontally extensible four-layer tree with task-specific refresh profiles instead of one flat monitored-file list.
- Contract state detector monitoring should use a four-layer product model: Layer 0 framework baseline, Layer 1 thread definition, Layer 2 cross-thread shared state, and Layer 3 thread local memory. Non-thread-management threads should confirm all four layer states at conversation start, then read only changed or task-relevant layers. At closeout, they should update Layer 3 for durable local knowledge and Layer 2 or handoff when other threads are affected; Layer 0 and Layer 1 changes belong to framework upgrade or Thread Management flows.
- Contract state detector file granularity should keep status checks cheap: Layer 1 has one thread-definition handle per thread; Layer 2 supports sparse directed dependency files such as `<source>__to_<target>` instead of eagerly creating every possible `N * (N - 1)` file; Layer 3 has category split points and should split local memory by category only when the state file becomes too large or volatile.
- Long-lived threads should run a periodic reanchor safety check every 10 conversation rounds: confirm all registered layer fingerprints, then read only changed, unknown, unreadable, or task-relevant layers.
- This repository now has an initialized numbered Layer 0 through Layer 3 structure: `00_framework_baseline`, `01_thread_definitions`, `02_shared_state`, and `03_thread_local_memory.md`, plus `docs/module_state/product_manager/` as the first category-level Layer 3 pilot. Current authoritative files remain `AGENTS.md`, `current_version.md`, `contracts.md`, and `docs/module_state/*.md` until Coordination promotes the split.
- Internal anchor directories should not use `README.md` by default. Reserve `README.md` for locations users are expected to open and understand; use `index.md`, named specs, or category files for internal detector/agent structures.
- Layer 1 semantic migration has started: all seven `docs/anchor_pm/01_thread_definitions/*.md` files now mirror full contract and initialization-prompt content. `contracts.md` and `thread_initialization.md` remain authoritative until Coordination promotes the split. A Layer 2 handoff to Coordination records the promotion decision needed.
- Layer 2 semantic migration has started: `current_version.md`, `interaction_guide.md`, `review_log.md`, and `simplification.md` now have mirrors under `docs/anchor_pm/02_shared_state/`. Old files remain authoritative until Coordination and owning threads promote the split or create compatibility redirects.
- Layer 3 semantic migration has started: all seven `docs/module_state/*.md` files now have category-level mirrors under `docs/module_state/<thread>/`. Old thread state files remain authoritative until Coordination promotes the split and each thread validates its own mirror.
- Reanchor Detector Core development should use `docs/anchor_pm/00_framework_baseline/reanchor_module_io_spec.md` as the fixed programmatic I/O contract. The module should compare registered file fingerprints, classify file status, output required reads and next actions, and never store chat history or business data.
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
- If a target project already has `AGENTS.md`, the installer should inspect it for conflicts. If no conflict is found, it should propose appending a short Anchor PM discovery section with the relevant `docs/anchor_pm/` and `docs/module_state/` paths.

## Open Issues

- Need a real user-facing walkthrough from new target project to installed Anchor PM anchors using the language-specific root install prompt files.
- Need Codex App testing of `Thread Management` install prompt.
- Need Codex Skill / Package Installer and Templates / Protocol handoff to mirror the new confirmation-page output shape into package release prompts.
- Templates / Protocol and Codex Skill / Package Installer need to mirror the
  automatic Reanchor Start behavior into package workflow text and install
  prompts without making users run CLI commands.
- Package release artifacts still contain older read-first reanchor wording in
  `packages/*/workflows/reanchor.md` and related templates; this is assigned to
  Templates / Protocol rather than Product Manager direct edits.
- Reanchor Detector Core and CLI Core need a callable programmatic reanchor
  entrypoint; until that exists, automatic rereading is only a degraded fallback
  and does not satisfy the product requirement.
- Need to track recurring first-time-user confusion as product requirements, not ad hoc discussion.

## Runbook

Before substantial product-flow work:

1. Run Reanchor Start automatically.
2. Use the detector if available.
3. If unavailable, report the degraded state, then read `AGENTS.md`,
   `docs/anchor_pm/current_version.md`,
   `docs/anchor_pm/contracts.md`, and this file.
4. Show a short anchor-state line before continuing.

When evaluating a user flow:

1. Identify the user goal.
2. Identify the first prompt or document the user sees.
3. Mark any place where the user must guess a value.
4. Propose the smallest change that removes the guess.
5. Handoff implementation details to the owning thread.

## History / Notes

- Created to own user operation flows and UX optimization for Anchor PM.
- Added product decision that README exposes multilingual install prompt files directly, avoiding an unclear "prompt index" as the primary user-facing concept, and installer chat output should show decisions while generated files hold long-form details.
- Added README decision that approval wording and post-install next steps should be visible before the user opens the install prompt.
- Revised README example around Anchor PM's clearest advantage: modular specialist context plus reanchor-based synchronization across changing modules.
- Reworked the README example into a generic promotional scenario for solo developers or small teams whose projects have outgrown one Codex chat.
- Added README tester qualification and feedback invitation sections.
- Added README current-status caveat that public testers must use a release package URL or their local cloned package path while prompts still point to a development path.
- Added README positioning that Anchor PM borrows from real project organization by turning general chats into specialists and giving them lightweight communication channels.
- Rewrote README into a compact public entry page with one-sentence positioning, value points, install instructions, one anonymized example, and tester feedback close.
- Revised README from a PM/UX review perspective: lead with fit, benefits, install, example, and tester feedback while keeping bilingual content compact.
- Recorded core product principle: prioritize real user AI coding improvement, surface UX/value failures early, and stop meaningless work if the approach does not improve experience.
- Added internal function spec for developers and prompt engineers, separating system behavior from user-facing README.
- Added internal Contract Version / Reanchor State Detector module to support version-change detection and shared-knowledge refresh decisions.
- Added focused `contract_state_detector.md` product design for the core anchor refresh module.
- Added monitored-document scope for the contract state detector to prevent it from becoming a whole-repository scanner.
- Revised contract detector monitoring from a flat list into a layered reanchor tree with task-specific refresh profiles.
- Corrected ordinary specialist thread behavior: conversation start must confirm Layer 0 through Layer 3 status, but should selectively read only changed or task-relevant layers; closeout updates Layer 3 for local durable knowledge and Layer 2 or handoff when other threads are affected.
- Added detector file-granularity direction and periodic 10-round reanchor safety check to balance stale-context risk against context cost.
- Initialized this repository's Layer 0 through Layer 3 directory skeleton without deleting or migrating current authoritative coordination files.
- Renamed internal Layer directories with numeric prefixes and removed internal `README.md` files so only user-facing locations use README-style entrypoints.
- Mirrored Layer 1 contract and initialization prompt semantics into per-thread definition files and created a Coordination handoff for promotion/deprecation decisions.
- Mirrored Layer 2 version/workflow/review/simplification semantics into `02_shared_state/` and recorded the remaining no-loss deletion gates.
- Mirrored Layer 3 thread-state semantics into category-level files under `docs/module_state/<thread>/`.
- Added the Reanchor Module I/O Spec under Layer 0 and linked it from detector/product handoff docs.
- Promoted automatic Reanchor Start into project rules, current-version
  requirement, thread initialization prompts, Layer 1 thread definitions,
  interaction guidance, and detector product docs.
- Clarified that programmatic anchoring, not automatic anchor-file rereading, is
  the target behavior; missing callable detector entrypoint is a product gap for
  Reanchor Detector Core and CLI Core.
