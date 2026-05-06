# Anchor PM Simplification Log

This file is used to prevent rule and document growth.

Before adding a new coordination document, check whether the content belongs in:

- `AGENTS.md`
- `docs/anchor_pm/00_framework_baseline/`
- `docs/anchor_pm/01_thread_definitions/`
- `docs/anchor_pm/02_shared_state/`
- `docs/module_state/<thread>.md`
- `docs/module_state/<thread>/`, only when category-level local memory is needed
- an existing product reference document

## Current Simplification Decisions

- `PRODUCT_PRINCIPLES.md`, `MVP_SPEC.md`, and `ANCHOR_PM_FRAMEWORK_DESIGN.md` remain reference documents, not higher-level rules.
- Thread contracts currently remain authoritative in `docs/anchor_pm/contracts.md`.
- Complete Layer 1 semantic mirrors now exist under `docs/anchor_pm/01_thread_definitions/`.
- `docs/anchor_pm/thread_initialization.md` remains the current prompt source until Coordination promotes the Layer 1 split.
- Long-term thread state currently remains authoritative in `docs/module_state/*.md`.
- Category files under `docs/module_state/<thread>/` are split points, not authoritative replacements, until Coordination promotes them.
