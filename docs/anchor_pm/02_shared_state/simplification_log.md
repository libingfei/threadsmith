# Simplification Log

Layer: `2`

Status: compatibility mirror.

Authoritative source until promoted:

- `docs/anchor_pm/simplification.md`

Purpose: prevent rule and document growth.

## Placement Rules

Before adding a new coordination document, check whether the content belongs in:

- `AGENTS.md`, for minimal AI discovery and project-level safety rules;
- `docs/anchor_pm/00_framework_baseline/`, for Anchor PM framework/package
  baseline information;
- `docs/anchor_pm/01_thread_definitions/`, for thread scope, boundaries,
  acceptance, hard rules, handoff rules, and initialization prompts;
- `docs/anchor_pm/02_shared_state/`, for cross-thread shared state, workflow
  notices, repeated mechanism issues, and migration decisions;
- `docs/module_state/<thread>.md`, for canonical thread local state during the
  1.0 compatibility period;
- `docs/module_state/<thread>/`, for category-level local memory only when a
  thread state file becomes too large or volatile;
- an existing product reference document.

## Current Simplification Decisions

- `PRODUCT_PRINCIPLES.md`, `MVP_SPEC.md`, and
  `ANCHOR_PM_FRAMEWORK_DESIGN.md` remain reference documents, not higher-level
  rules.
- `docs/anchor_pm/contracts.md` remains the current authoritative Layer 1
  compatibility source until Coordination promotes
  `docs/anchor_pm/01_thread_definitions/*.md`.
- `docs/anchor_pm/thread_initialization.md` remains the current authoritative
  prompt source until Coordination promotes per-thread definition prompts.
- `docs/module_state/*.md` remain the canonical Layer 3 state files during the
  1.0 compatibility period.
- Category files under `docs/module_state/<thread>/` are split points, not
  authoritative replacements, until Coordination promotes them.
