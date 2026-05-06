# Workflow: Self-Optimization

Use this workflow only for Anchor PM self-evolution mode.

## Goal

Generate one report that identifies improvements for the next Anchor PM source version.

## Inspect

Read:

- `AGENTS.md`
- `docs/anchor_pm/current_version.md`
- `docs/anchor_pm/contracts.md`
- `docs/module_state/*.md`
- product reference docs
- package files under `packages/`

## Evaluate

Check:

- package-first rules are reflected in specs;
- standard and self-evolution modes have identical structure;
- standard mode stops after install;
- self-evolution mode stops after recommendations;
- templates are short and project-neutral;
- status, reanchor, handoff, and safety workflows are present;
- module state files reflect current work.

## Report Format

```text
# Self-Optimization Report

Observed:

Inference:

Unverified:

Candidate Sn -> Sn+1 Improvements:

Blocking Issues:

Non-Blocking Risks:

Suggested Handoffs:
```

Do not apply the report automatically.
