# Anchor PM Thread Contracts

## Coordination

Scope:

- Maintain project boundaries, thread contracts, coordination versioning, and handoff rules.
- Keep Anchor PM lightweight and prevent framework scope creep.
- Decide when work should move to Product Manager, Reanchor Detector Core, CLI Core, Templates / Protocol, Codex Skill / Package Installer, or Dogfood / Validation threads.
- Own and execute Anchor PM self-evolution cycles for this project.
- Maintain module state files when long-term project state changes.

Out of Scope:

- Implementing CLI internals.
- Writing Codex Skill implementation details.
- Designing business-specific project rules for downstream users.
- Treating source-project-specific behavior as framework core.
- Running broad external validation campaigns.

Acceptance:

- Project-level coordination files stay concise and internally consistent.
- Cross-thread work is handed off instead of silently absorbed.
- Self-evolution recommendations are reviewed in this thread before becoming work.
- Important conclusions use `Observed / Inference / Unverified`.

Hard Rules:

- Prefer reducing rule sources over adding new ones.
- Do not promote reference material into rules without an explicit reason.
- Do not claim implementation readiness before code and validation exist.
- Do not auto-apply self-evolution recommendations; each change still requires explicit implementation.

State File:

- `docs/module_state/coordination.md`

Handoff Rule:

- Produce a structured handoff summary with source thread, target thread, confirmed facts, impact, unresolved questions, and suggested next step.

## Product Manager

Scope:

- Own user operation flows, onboarding paths, and experience optimization for Anchor PM.
- Maintain clarity of install prompts, thread creation guidance, and user-facing workflow expectations.
- Translate user friction into concrete product requirements for package, template, or installer threads.
- Ensure Anchor PM feels comfortable for first-time users and does not require users to guess placeholders or hidden steps.

Out of Scope:

- Implementing package templates or CLI internals directly.
- Owning thread contracts or self-evolution execution.
- Running validation campaigns.
- Making business-specific rules for downstream projects.

Acceptance:

- User-facing flows have clear entrypoints and next steps.
- Install and thread initialization prompts are copy-paste-ready where possible.
- Interaction language follows the user's usual language while project docs may remain English.
- Product requirements distinguish user pain, proposed behavior, and implementation owner.

Hard Rules:

- Do not add new concepts when a clearer prompt or shorter workflow solves the problem.
- Do not require users to fill in values Codex can infer or generate.
- Do not weaken safety gates for smoother UX.

State File:

- `docs/module_state/product_manager.md`

Handoff Rule:

- Handoff package wording to Templates / Protocol, install-flow behavior to Codex Skill / Package Installer, validation evidence to Dogfood / Validation, and boundary decisions to Coordination.

## Reanchor Detector Core

Scope:

- Own the Contract Version / Reanchor State Detector as a core code subsystem.
- Define and maintain its input/output contract, file-reading behavior, error handling, and efficiency requirements.
- Detect whether shared anchors changed and whether a thread must reread `AGENTS.md`, `contracts.md`, and/or its `module_state`.
- Treat unknown, missing, or unreadable version state conservatively.
- Provide a stable internal interface for future CLI, package, and prompt workflows.

Out of Scope:

- Deciding business task scope.
- Owning general CLI command UX.
- Writing user-facing product flows.
- Defining template prose except where required by the detector interface.
- Running broad validation campaigns.

Acceptance:

- Detector distinguishes `unchanged`, `changed`, and `unknown` version states.
- Detector outputs a concise refresh decision and reason.
- Detector avoids unnecessary file reads while remaining conservative on unknown state.
- Detector behavior is testable with fixture projects and missing/corrupt anchor files.
- Detector does not expand into business validation or project management.

Hard Rules:

- Never silently treat missing or unreadable version files as unchanged.
- Never require reading unrelated project files to answer reanchor state.
- Keep machine-readable output stable once introduced.
- Prefer deterministic file parsing over prompt-only inference for code paths.

State File:

- `docs/module_state/reanchor_detector_core.md`

Handoff Rule:

- Handoff version semantics to Coordination, prompt wording to Templates / Protocol, CLI command packaging to CLI Core, and user-facing flow concerns to Product Manager.

## CLI Core

Scope:

- Own future `anchorpm` command-line behavior after 1.0.
- Preserve CLI as a reproducibility, status-checking, and batch automation layer.
- Ensure future CLI behavior can reproduce package-first install semantics.
- Call the Reanchor Detector Core for contract-version and reanchor-state decisions instead of reimplementing that logic.

Out of Scope:

- Product strategy.
- Codex Skill prompt design beyond machine-readable CLI outputs.
- Reanchor detector internals.
- Web UI, RAG, scheduler, auto-deploy, or agent router behavior.

Acceptance:

- CLI is not a blocker for Anchor PM 1.0.
- Future CLI scope remains aligned with package-first safety rules.
- CLI design does not reintroduce automatic business-code modification.

Hard Rules:

- Do not modify target projects during audit.
- Do not overwrite user rules without dry-run diff and explicit apply.
- Compile and validate in the local `rserver` container for this repository.

State File:

- `docs/module_state/cli_core.md`

Handoff Rule:

- Handoff product or template ambiguity to Coordination or Templates / Protocol before encoding it as CLI behavior.

## Templates / Protocol

Scope:

- Own default Markdown templates, package workflow text, and protocol text.
- Keep templates short, auditable, and project-neutral.
- Define required sections for `AGENTS.md`, `contracts.md`, `current_version.md`, `interaction_guide.md`, and `module_state/<thread>.md`.
- Maintain package-first install plans, checklists, and workflow documents.

Out of Scope:

- CLI parsing and filesystem behavior.
- Codex App behavior outside the package instructions.
- Case-study-specific logic.

Acceptance:

- Templates and workflows are small enough for users to review quickly.
- Templates distinguish direct observation from inference.
- Templates do not assume a specific business domain.
- Standard and self-evolution package structures remain identical except `ACTIVE_INSTALL_PLAN.md` and `INSTALL_PROMPT.md`.

Hard Rules:

- Do not create a second project brain.
- Do not copy source-project-specific terms into framework core.
- Prefer one reusable section over multiple near-duplicate documents.

State File:

- `docs/module_state/templates_protocol.md`

Handoff Rule:

- Handoff implementation concerns to CLI Core; handoff interaction wording concerns to Codex Skill.

## Codex Skill / Package Installer

Scope:

- Provide Codex-first usage of Anchor PM through package instructions.
- Own `INSTALL_PROMPT.md` behavior and target-thread installation flow.
- Ensure Codex reads `PACKAGE_MANIFEST.md` and `ACTIVE_INSTALL_PLAN.md` before acting.

Out of Scope:

- Implementing a CLI.
- Storing project-specific rules inside the Skill.
- Acting as a general project manager.

Acceptance:

- Installation behavior remains a thin adapter over package docs and target Markdown anchors.
- Standard mode stops after deployment.
- Self-evolution mode stops after producing recommendations.
- The installer does not invent project facts.

Hard Rules:

- Project rules live in the target project, not inside the Skill.
- The Skill must preserve `Observed / Inference / Unverified` distinctions.
- User confirmation is required before writing files.

State File:

- `docs/module_state/codex_skill.md`

Handoff Rule:

- Handoff future CLI features to CLI Core; handoff protocol ambiguity to Coordination.

## Dogfood / Validation

Scope:

- Record validation results for this repository and selected real projects.
- Record whether generated anchors are understandable, short, useful, and safe.
- Surface repeated failure modes into review logs.
- Validate standard and self-evolution package behavior.

Out of Scope:

- Changing product rules directly without Coordination review.
- Adding project-specific behavior to framework core.
- Owning the Anchor PM self-evolution loop.

Acceptance:

- Each validation run records observed behavior, inferred issues, and unresolved questions.
- Failures distinguish blocking defects from product-improvement notes.
- Dogfood results feed back into contracts, templates, or CLI behavior only through explicit changes.

Hard Rules:

- Do not treat one successful sample project as full validation.
- Do not upgrade case-study details into universal rules.
- Do not apply self-evolution recommendations automatically.

State File:

- `docs/module_state/dogfood_validation.md`

Handoff Rule:

- Handoff product boundary questions and all self-evolution decisions to Coordination; handoff implementation failures to the owning implementation thread.
