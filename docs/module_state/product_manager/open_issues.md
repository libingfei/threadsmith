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
- Need install prompt/package wording to enforce thread-name language matching
  the selected install prompt, and remove `Adjust AGENTS.md` from default reply
  options.
- Need package release artifacts and future Skill behavior to mirror the
  root-prompt rule that default reply options are approve, adjust threads, and
  cancel; AGENTS decisions appear only as concrete approval risks when blocked.
- Templates / Protocol and Codex Skill / Package Installer need to mirror
  automatic lightweight Anchor Gate behavior into package workflow text and
  install prompts without making users run CLI commands.
- Package release artifacts still contain older read-first reanchor wording in
  `packages/*/workflows/reanchor.md` and related templates; this is assigned to
  Templates / Protocol rather than Product Manager direct edits.
- Reanchor Detector Core and CLI Core need a callable programmatic reanchor
  entrypoint; until that exists, automatic rereading is only a degraded fallback
  and does not satisfy the product requirement.
- Templates / Protocol and Codex Skill / Package Installer need to ensure every
  generated thread prompt includes Knowledge Sync Gate as well as Anchor Gate.
- Templates / Protocol and Codex Skill / Package Installer need to include User
  Delta Triage inside a lightweight Anchor Gate in generated thread prompts and
  interaction guides, not as a verbose standalone step.
- Need to track recurring first-time-user confusion as product requirements, not
  ad hoc discussion.
