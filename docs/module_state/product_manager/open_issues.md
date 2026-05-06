# Product Manager Open Issues

Layer: `3`

Thread: `Product Manager`

Status: compatibility mirror of `docs/module_state/product_manager.md`.

- Need a real user-facing walkthrough from new target project to installed
  Anchor PM anchors using the language-specific root install prompt files.
- Need Codex App testing of `Thread Management` install prompt.
- Need Codex Skill / Package Installer and Templates / Protocol handoff to
  mirror the new confirmation-page output shape into package release prompts.
- Need Templates / Protocol and Codex Skill / Package Installer to replace
  generic Coordination / Implementation / Validation defaults for existing
  projects with module/subsystem-based thread generation.
- Need package release artifacts to stop exposing docs-only/no-AGENTS as a
  default approval path; partial integration should be exceptional and clearly
  labeled.
- Templates / Protocol and Codex Skill / Package Installer need to mirror
  automatic Reanchor Start behavior into package workflow text and install
  prompts without making users run CLI commands.
- Package release artifacts still contain older read-first reanchor wording in
  `packages/*/workflows/reanchor.md` and related templates; this is assigned to
  Templates / Protocol rather than Product Manager direct edits.
- Reanchor Detector Core and CLI Core need a callable programmatic reanchor
  entrypoint; until that exists, automatic rereading is only a degraded fallback
  and does not satisfy the product requirement.
- Templates / Protocol and Codex Skill / Package Installer need to ensure every
  generated thread prompt includes Closeout Knowledge Sync, not only Reanchor
  Start.
- Need to track recurring first-time-user confusion as product requirements, not
  ad hoc discussion.
