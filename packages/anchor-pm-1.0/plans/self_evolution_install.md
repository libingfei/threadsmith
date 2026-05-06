# Self-Evolution Install Plan

Use this plan only when the target project is Anchor PM itself.

## Goal

Install or refresh Anchor PM coordination anchors, then generate one self-optimization report for the next source version.

This plan stops at recommendations. It does not automatically modify source, product specs, templates, or package files after the report.

## Inputs

- Anchor PM source root
- Anchor PM 1.0 package root
- Current coordination anchors, if present
- Current product reference docs

## Phase 1: Read Package

1. Read `PACKAGE_MANIFEST.md`.
2. Read this active plan.
3. Read:
   - `workflows/existing_project_adoption.md`
   - `workflows/self_optimization.md`
   - `workflows/status_check.md`
   - `checklists/safety_check.md`
   - `checklists/conclusion_check.md`
   - `checklists/drift_check.md`

## Phase 2: Verify Target

Confirm the target looks like Anchor PM.

Required evidence should include at least one of:

- Product docs mentioning Anchor PM.
- `docs/anchor_pm/` coordination files.
- `packages/anchor-pm-1.0/` package files.

If the target is not clearly Anchor PM, stop and ask the user whether standard mode was intended.

## Phase 3: Deploy Or Refresh Anchors

Follow existing-project adoption rules:

- Read current anchors.
- Identify missing or stale coordination files.
- Propose creates or updates.
- Do not write until user approval.

## Phase 4: Generate Self-Optimization Report

After deployment or if anchors already exist, run `workflows/self_optimization.md`.

The report must include:

- `Observed`
- `Inference`
- `Unverified`
- Candidate improvements for `Sn -> Sn+1`
- Blocking issues
- Non-blocking risks
- Suggested handoffs by thread

## Phase 5: Stop

Do not apply the optimization report automatically.

The next source version requires human review and explicit implementation work.
