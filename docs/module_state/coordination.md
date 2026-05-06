# Coordination Thread State

## Contract

Thread contract: `Coordination` in `docs/anchor_pm/contracts.md`

Current coordination version: `package-first-v1.0`

## Current State

- Anchor PM is a product/framework project for Codex-first AI coding thread coordination.
- The project is self-hosting at the documentation, thread-contract, and package-first release level.
- Current source material exists in `PRODUCT_PRINCIPLES.md`, `MVP_SPEC.md`, and `ANCHOR_PM_FRAMEWORK_DESIGN.md`.
- Anchor PM 1.0 uses Codex package-first installation, not CLI-first implementation.
- Future CLI implementation remains a reproducibility and checking layer, not a 1.0 blocker.
- The product must support both existing-project adoption and new-project bootstrap.
- The product must support standard and self-evolution package modes with identical structure and different active plans.
- Because this project is small, the Coordination thread owns Anchor PM self-evolution execution and review.

## Open Issues

- CLI implementation is deferred until after 1.0.
- Codex Skill / Plugin packaging remains future work.
- Standard package dry-run has only been performed against a local sample, not a real external user project.
- Self-evolution dry-run has produced recommendations but those recommendations have not been applied.
- Dogfood / Validation remains useful for recording external validation evidence, but it no longer owns the self-evolution loop.
- Self-evolution round 1 has produced `docs/anchor_pm/reports/self_evolution_round_1.md`; its recommendations have not been applied.
- Thread initialization prompts now live in `docs/anchor_pm/thread_initialization.md` so users do not need to invent thread setup prompts.
- Final user-facing thread prompts must be complete; users should not have to fill in placeholders that Codex can generate.
- Target-project integration now has a copy-paste `Thread Management` prompt at `docs/anchor_pm/thread_management_install_prompt.md` with the development package path.
- Thread Management installer replies should follow the user's usual language; project docs may remain English unless requested otherwise.
- Product Manager thread now owns user operation flows and experience optimization.
- Reanchor Detector Core now owns the Contract Version / Reanchor State Detector as a core code subsystem.

## Runbook

Before substantial coordination work:

1. Run Reanchor Start automatically.
2. Use the detector if available.
3. If unavailable, read `AGENTS.md`, `docs/anchor_pm/current_version.md`,
   `docs/anchor_pm/contracts.md`, and this file.
4. Show a short anchor-state line before continuing.

When a request belongs to another thread:

1. State the boundary.
2. Produce a handoff summary.
3. Update this file only if long-term coordination state changed.

When running self-evolution:

1. Use `packages/anchor-pm-1.0-self-evolution/`.
2. Generate one self-optimization report.
3. Separate `Observed`, `Inference`, and `Unverified`.
4. Convert accepted recommendations into explicit follow-up work.
5. Do not auto-apply recommendations within the same step unless the user separately asks for implementation.

When making strong product conclusions:

1. Separate `Observed`, `Inference`, and `Unverified`.
2. Avoid closure wording unless a formal check or validation run exists.

## History / Notes

- `selfdogfood-v0.1`: Established this repository as its own first Anchor PM adoption target.
- `package-first-v1.0`: Added Anchor PM 1.0 package-first structure and release-mode split.
- `coordination-owned-self-evolution`: Moved self-evolution execution under the Coordination thread because the project is small.
- `self-evolution-round-1`: Ran the self-evolution package manually and recorded recommendations without applying them.
- `thread-initialization-prompts`: Added a dedicated document for recommended thread count and per-thread startup prompts.
- `no-user-placeholder-prompts`: Clarified that generated user-facing thread prompts must be copy-paste-ready.
- `thread-management-install-prompt`: Added a concrete local-path install prompt for target project integration.
- `user-language-install-interaction`: Added language-following rules for install prompts and package manifests.
- `product-manager-thread`: Added Product Manager as a long-lived thread for user flows and UX logic.
- `reanchor-detector-core-thread`: Added Reanchor Detector Core as a long-lived thread because detector reliability and I/O efficiency are central to system feasibility.
