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
