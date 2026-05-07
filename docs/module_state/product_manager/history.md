# Product Manager History

Layer: `3`

Thread: `Product Manager`

Status: compatibility mirror of `docs/module_state/product_manager.md`.

- Created to own user operation flows and UX optimization for Anchor PM.
- Added product decision that README exposes multilingual install prompt files directly, avoiding an unclear "prompt index" as the primary user-facing concept, and installer chat output should show decisions while generated files hold long-form details.
- Added README decision that approval wording and post-install next steps should be visible before the user opens the install prompt.
- Revised README example around Anchor PM's clearest advantage: modular specialist context plus reanchor-based synchronization across changing modules.
- Reworked the README example into a generic promotional scenario for solo developers or small teams whose projects have outgrown one Codex chat.
- Added README tester qualification and feedback invitation sections.
- Added README current-status caveat that public testers must use a release package URL or their local cloned package path while prompts still point to a development path.
- Replaced local development package paths in public install prompts with the
  public Threadsmith GitHub repository and package directory.
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
- Added the Layer 0 Reanchor Module I/O Spec and linked it from detector,
  internal function, current version, and Product Manager shared handoff docs.
- Corrected Layer 2 directed dependency naming to
  `<source>__to_<target>.md` so programmatic discovery matches existing files.
- Promoted automatic Reanchor Start into project rules, current-version
  requirement, thread initialization prompts, Layer 1 thread definitions,
  interaction guidance, and detector product docs.
- Clarified that programmatic anchoring, not automatic anchor-file rereading, is
  the target behavior; missing callable detector entrypoint is a product gap for
  Reanchor Detector Core and CLI Core.
- Flask install dry-run exposed a product failure: the proposal used generic
  Coordination / Implementation / Validation threads, required a pre-named
  thread, showed too much internal safety explanation, and offered a confusing
  docs-only/no-AGENTS option.
- Added PM Review Gate after the Flask dry-run mistake: Product Manager must
  evaluate semantic usefulness and real user operation, not only checklist
  field presence, and must record durable behavior changes into thread state.
- Promoted Closeout Knowledge Sync as a core all-thread workflow: every
  long-lived thread must check local durable knowledge, cross-thread shared
  knowledge, thread-definition impact, and framework-level impact before final
  response.
- Defined Reanchor Start and Closeout Knowledge Sync as symmetric thread
  lifecycle hooks: read changed knowledge before work, write or hand off new
  durable knowledge before reply.
- Rewrote the MVP manual validation protocol around the complete lifecycle from
  new project adoption through issue solving, handoff, closeout knowledge sync,
  shared-state recovery, and restore.
- Added Pre-test GitHub Sync Gate so public GitHub-based tests use the latest
  prompts and validation docs rather than stale remote files.
- Flask install retest showed improved module-based thread splitting, but
  thread names were still English under the Chinese prompt and `Adjust
  AGENTS.md` remained as a default reply option; recorded both as product
  wording fixes.
- Added User Delta Triage as the pre-reanchor intake step: explicit durable user
  corrections should be classified before reanchor so stale anchors do not
  override the newest user intent; closeout remains necessary for knowledge
  learned during execution.
- Compressed anchor lifecycle into `Anchor Gate` and `Knowledge Sync Gate` so
  ordinary turns do not lose context space to process. Default gate behavior is
  silent, no-write, and minimal-read.
- Updated root English and Chinese install prompts to require generated thread
  prompts and interaction docs to use lightweight `Anchor Gate` and
  `Knowledge Sync Gate` semantics, even if package templates still contain older
  standalone Reanchor/Closeout wording.
- Simplified install proposal UX: removed default `Recommendation`, `Changes`,
  `Needs your decision`, and `Decision details` sections from the main proposal;
  optional rationale should move to a collapsed block or
  `docs/anchor_pm/install_decision_record.md`.
- Removed `AGENTS.md` handling from the default install proposal main view. It
  should appear only as a short risk/decision line when a real conflict or merge
  choice requires user input; otherwise record it after approval.
