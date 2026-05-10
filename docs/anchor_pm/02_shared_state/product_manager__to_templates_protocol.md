# Product Manager To Templates Protocol

Layer: `2`

Source thread: `Product Manager`

Target thread: `Templates / Protocol`

Status: active shared dependency.

## Confirmed Product Requirements

- Reanchor must be presented as automatic Codex behavior, not a user-operated
  CLI step.
- Package workflow text should say Codex runs Reanchor Start before substantial
  long-lived thread work.
- If a detector command/tool is available, Codex should run it and follow
  `required_reads`, `blocked_by`, and `next_action`.
- If the detector is unavailable, Codex should report the degraded state, then
  read the required anchors itself as a compatibility fallback.
- Template wording must not present fallback rereading as the intended product
  behavior.
- Users should see only a short anchor-state line unless the thread is blocked.
- Templates should not imply that users must remember to request reanchor or run
  a CLI command manually.
- Existing-project install templates must generate module/subsystem-based
  project specialist threads, not default Coordination / Implementation /
  Validation responsibility buckets.
- `Coordination` is an Anchor PM internal/self-hosting role and should not be
  proposed as an ordinary target-project thread.
- The current install conversation is the Thread Management entrypoint; user
  prompts must not require the Codex conversation to be named before the first
  message.
- Each proposed target-project specialist should normally own code, tests, docs,
  and validation evidence for its module. Standalone `Validation` should be used
  only when the target project has a clear independent validation/release
  subsystem or the user asks for it.
- User-facing install proposals should hide verbose internal safety constraints
  and package execution details unless the user asks.
- User-visible thread names must follow the selected install prompt language.
  For the Chinese root prompt, thread names should be Chinese while technical
  terms such as Sans-IO, CLI, HTTP, and API may remain in English.
- "Install docs only; do not update AGENTS.md" should not appear as a default
  approval option. If AGENTS handling is blocked, the package should describe
  the consequence and ask for a specific merge/skip decision.
- `Adjust AGENTS.md` should not appear as a default reply option. Default
  options should be approve, adjust threads, and cancel; AGENTS-specific choices
  should appear only as concrete decisions when a real conflict or risk exists.
- Every generated long-lived thread prompt should include a lightweight Anchor
  Gate before work and Knowledge Sync Gate before final response. The gates
  should default to silent/no-write/minimal-read behavior.
- Anchor Gate should include same-turn user correction handling and Reanchor
  Start, but should not be exposed as multiple user-visible steps.
- Knowledge Sync Gate should update or hand off only durable local/shared
  knowledge; no durable change means no visible closeout note by default.
- Wording should enforce an anchor budget: gate handling must not crowd out the
  actual task response.
- Root install prompts should stay under 20 lines and only point Codex to the
  public package source, package directory, package install instructions, and
  package proposal template.
- Package templates should include `templates/install_proposal.template.md` as
  the stable first-response output contract so proposal sections do not drift
  back to verbose audit-style content.
- Package templates should include `templates/install_completion.template.md`
  as the stable post-approval completion contract. The completion page should
  be localized, start with a successful-install confirmation, briefly explain
  what the user now has, teach users how to create specialist conversations,
  show each complete copy-paste-ready thread prompt, and avoid turning the
  reply into a generated-file inventory.
- Thread creation prompts in `thread_initialization.md` are user-facing output
  and must match the install prompt language. Generated docs may otherwise be
  English unless requested, but copy-paste thread prompts should not switch
  language.
- Ordinary target-project installs should place all Anchor PM files under one
  `.threadsmith/` directory, including `.threadsmith/AGENTS.md`,
  `.threadsmith/thread_initialization.md`, and
  `.threadsmith/module_state/<thread>.md`.
- Root `AGENTS.md` integration should be opt-in only. Default installation must
  not create or modify root `AGENTS.md`, so unrelated Codex conversations are
  not automatically pulled into Anchor PM behavior.
- The completion page should explicitly tell users that `.threadsmith/` keeps
  Threadsmith isolated from ordinary Codex conversations unless those
  conversations start with generated thread prompts, and that the current
  install conversation remains Thread Management for creating, deleting,
  changing, regenerating, or querying threads.

## Target Next Step

Mirror this behavior into package workflow text and reusable protocol wording.
Known package locations still needing owner-thread update include
`packages/*/workflows/reanchor.md`,
`packages/*/templates/interaction_guide.template.md`, and
`packages/*/templates/module_state.template.md`.

Additional locations needing owner-thread review after the Flask install
dry-run include `packages/*/templates/thread_initialization.template.md`,
`packages/*/workflows/new_project_bootstrap.md`,
`packages/*/workflows/existing_project_adoption.md`, and package
`INSTALL_PROMPT.md` files. Anchor Gate and Knowledge Sync Gate wording needs
mirroring into `packages/*/templates/interaction_guide.template.md` and any
generated per-thread prompt text.

Product Manager has made a narrow package-surface update to unblock public
testing: shortened root prompts, added `templates/install_proposal.template.md`,
and updated standard install prompt/plan/workflow wording. Templates / Protocol
should review and normalize these changes across future package releases.

Product Manager also added `templates/install_completion.template.md` and a
thread-prompt language rule after a Chinese Flask install produced English
thread prompts and no clear next-thread creation teaching.

Product Manager then changed the standard target-project footprint to
`.threadsmith/` after testing showed files scattered across root, `docs/`, and
module state directories were harder to maintain and leaked the framework into
unrelated threads.

Product Manager then revised the install completion template from a file-report
shape into a success/next-action page that includes recommended thread prompts
directly and explains `.threadsmith/` isolation plus Thread Management
capabilities.

Do not change detector internals or CLI packaging here.
