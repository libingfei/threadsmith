# Anchor PM Internal Function Spec

Audience: developers, prompt engineers, package maintainers, and validation
threads.

Purpose: define Anchor PM functional modules, their inputs, outputs, and
handoff points. README remains the user-facing introduction; this document is
the internal build reference.

## Product Boundary

Anchor PM 1.0 is a Codex package-first Markdown coordination package.

It does not implement a CLI-first workflow, RAG, background scheduler, autonomous
agent router, deployment system, or business-code modifier.

Core user value:

- turn broad AI coding chats into focused specialist threads;
- keep each thread's context narrow;
- provide lightweight reanchor and handoff paths for cross-module dependency
  information;
- keep conclusions conservative and auditable.

## Module Map

```text
User README / Prompt Entry
  -> Thread Management Installer
  -> Target Project Audit
  -> Thread Plan Generator
  -> Anchor File Generator
  -> AGENTS.md Integration
  -> Post-Install Thread Creation Guide
  -> Contract Version / Reanchor State Detector
  -> Reanchor / Handoff / Status Workflows
  -> Validation / Feedback Loop
```

## 1. User README / Prompt Entry

Owner: Product Manager for flow; Templates / Protocol for wording; Codex Skill /
Package Installer for release prompt behavior.

Inputs:

- Repository-root `README.md`.
- Language-specific install prompts:
  - `ANCHOR_PM_INSTALL_PROMPT.en.md`
  - `ANCHOR_PM_INSTALL_PROMPT.zh.md`
- Package path or release package URL.
- User's preferred conversation language.

Outputs:

- User understands whether Anchor PM is a fit.
- User opens target project in Codex.
- User creates `Thread Management` / `线程管理` thread.
- User pastes one language-specific install prompt.

Acceptance:

- The entry path is copy-paste-ready.
- No user-filled placeholders are required for the package path in released
  prompts.
- README states current development status when prompts still use local paths.
- User can tell when Anchor PM is probably unnecessary.

Failure Signals:

- User asks what file to copy.
- User copies the wrong prompt language.
- User cannot tell whether to use Anchor PM for their project.
- Local package path causes public install failure without warning.

## 2. Thread Management Installer

Owner: Codex Skill / Package Installer.

Inputs:

- User-pasted install prompt.
- Target project root.
- Package root.
- `PACKAGE_MANIFEST.md`.
- `ACTIVE_INSTALL_PLAN.md`.
- Referenced install plan, workflows, checklists, and templates.

Outputs Before Approval:

- Concise installation proposal, not a package execution log.
- Target project path.
- Detected project type: `existing project` or `new project`.
- Recommended thread count and thread names.
- One-sentence responsibility per thread.
- File create/update summary.
- `AGENTS.md` handling recommendation.
- Safety promises.
- 1-3 approval-relevant risks or decisions.
- Localized reply options such as:
  - `Approve install`
  - `Adjust threads: ...`
  - `Install docs only; do not update AGENTS.md`
  - `Cancel`

Outputs After Approval:

- Approved files created.
- Approved files updated.
- Files intentionally left untouched.
- Links to generated thread prompts and usage docs.
- Clear next step: create long-lived threads from
  `docs/anchor_pm/thread_initialization.md`.
- Reminder that `Thread Management` remains available for future thread changes.

Acceptance:

- Does not write before explicit approval.
- Does not expose package internals unless user asks.
- Does not continue optimizing the business project after standard install.
- Follows the user's conversation language.

Failure Signals:

- User must infer what to approve.
- Proposal is longer than needed for a decision.
- Installer shows manifest/workflow logs as primary output.
- Installer starts implementation work after installing anchors.

## 3. Target Project Audit

Owner: Codex Skill / Package Installer; Templates / Protocol for audit wording.

Inputs:

- Target root filesystem.
- Existing project files, especially:
  - `README*`
  - `AGENTS.md`
  - existing docs
  - scripts/config/tests
  - CI/deploy files
- User-provided module or thread hints, if any.

Outputs:

- `Observed`: direct file and directory facts.
- `Inference`: likely modules, thread boundaries, and install mode.
- `Needs Confirmation`: assumptions that must not become formal contracts yet.
- Existing rule files and conflict risks.

Acceptance:

- Audit is read-only.
- Existing project rules are not overwritten silently.
- Business facts inferred from filenames are kept as inference until confirmed.

Failure Signals:

- Installer promotes guessed module boundaries into contracts.
- Existing `AGENTS.md` is overwritten.
- Audit tries to run build, deploy, migration, or business workflows.

## 4. Thread Plan Generator

Owner: Templates / Protocol for defaults; Product Manager for UX; Coordination
for boundary decisions.

Inputs:

- Target audit output.
- User-provided thread preferences.
- Default bootstrap rule:
  - small/new project: `Coordination`, `Implementation`, `Validation`
  - larger project: only add extra threads when responsibilities are clearly
    separate
- Existing project structure and docs.

Outputs:

- Proposed thread list.
- Per-thread:
  - name;
  - scope;
  - out-of-scope boundaries;
  - state file path;
  - one-sentence user-facing responsibility;
  - final copy-paste-ready initialization prompt.

Acceptance:

- Thread count is minimal.
- User does not fill `<thread name>`, `<thread_file>`, or similar placeholders.
- Thread boundaries are easy to read.
- Cross-boundary work has a handoff path.

Failure Signals:

- Threads are added for hypothetical future work.
- User cannot distinguish two thread responsibilities.
- The generated prompts contain unresolved placeholders.

## 5. Anchor File Generator

Owner: Templates / Protocol.

Inputs:

- Approved installation proposal.
- Package templates:
  - `AGENTS.template.md`
  - `current_version.template.md`
  - `contracts.template.md`
  - `thread_initialization.template.md`
  - `interaction_guide.template.md`
  - `module_state.template.md`
  - `review_log.template.md`
  - `simplification.template.md`
- Confirmed project facts.
- User-approved inferred facts.

Outputs:

- `AGENTS.md` create or approved update.
- `docs/anchor_pm/current_version.md`.
- `docs/anchor_pm/contracts.md`.
- `docs/anchor_pm/thread_initialization.md`.
- `docs/anchor_pm/interaction_guide.md`.
- Optional support docs:
  - `docs/anchor_pm/review_log.md`
  - `docs/anchor_pm/simplification.md`
- `docs/module_state/<thread>.md` files.

Acceptance:

- Generated files are concise.
- Generated docs are project-neutral unless confirmed project facts are provided.
- Module state files are sparse state, not chat transcripts.
- Strong conclusion protocol is present.

Failure Signals:

- Generated docs are longer than the project can maintain.
- Multiple files repeat the same rules.
- Inference is written as confirmed project fact.

## 6. AGENTS.md Integration

Owner: Codex Skill / Package Installer for behavior; Templates / Protocol for
wording; Coordination for rule conflicts.

Inputs:

- Existing `AGENTS.md`, if present.
- Proposed Anchor PM paths.
- Existing project rule hierarchy.
- User approval.

Outputs:

- If missing: proposed new `AGENTS.md`.
- If present and no conflict: proposed short Anchor PM discovery section pointing
  to:
  - `docs/anchor_pm/current_version.md`
  - `docs/anchor_pm/contracts.md`
  - `docs/anchor_pm/thread_initialization.md`
  - `docs/anchor_pm/interaction_guide.md`
  - `docs/module_state/`
- If conflict or uncertainty: conflict summary and explicit merge decision.

Acceptance:

- Existing rules remain authoritative unless user approves a merge.
- Anchor PM section helps AI agents discover anchors.
- No business rules are invented.

Failure Signals:

- Installer replaces existing project policy.
- Anchor PM discovery section becomes a second project constitution.
- Conflict is hidden inside a large proposal.

## 7. Post-Install Thread Creation Guide

Owner: Product Manager for flow; Templates / Protocol for generated document.

Inputs:

- Final thread plan.
- Generated contracts.
- Generated module state filenames.

Outputs:

- `docs/anchor_pm/thread_initialization.md` with one complete prompt per thread.
- Installer completion message linking to:
  - `docs/anchor_pm/thread_initialization.md`
  - `docs/anchor_pm/contracts.md`
  - `docs/anchor_pm/interaction_guide.md`
  - `docs/anchor_pm/current_version.md`

Acceptance:

- User can create long-lived threads without inventing prompts.
- `Thread Management` is preserved for future thread maintenance.
- Standard install stops after this step.

Failure Signals:

- User must manually assemble thread prompts.
- Completion message dumps full prompts into chat instead of linking generated
  files.
- Installer continues into ordinary project work.

## 8. Contract Version / Reanchor State Detector

Owner: Reanchor Detector Core for code subsystem behavior; Templates / Protocol
for prompt behavior; CLI Core for command packaging; Coordination for version
semantics.

Detailed design: `docs/anchor_pm/contract_state_detector.md`.

Engineering I/O contract:
`docs/anchor_pm/00_framework_baseline/reanchor_module_io_spec.md`.

Inputs:

- `ReanchorRequest` with schema version, operation, project root, thread id,
  checkpoint, registry mode, optional handoff files, and options.
- Registered Layer 0 through Layer 3 anchor files.
- Previous checkpoint fingerprints, if available.
- Current conversation round count for periodic reanchor.
- Optional closeout events for local memory or cross-thread dependency updates.

Outputs:

- `ReanchorResult` with schema version, operation, thread id, checked time, and
  anchor state:
  - `unchanged`
  - `changed`
  - `unknown`
  - `blocked`
- Per-file statuses:
  - `unchanged`
  - `changed`
  - `unknown`
  - `missing_required`
  - `missing_optional`
  - `unreadable`
  - `invalid_path`
- Changed layers.
- Required reads.
- Blockers.
- Next action.
- Checkpoint update proposal.
- Minimal chat output.

Acceptance:

- Every substantial long-lived thread start can confirm Layer 0 through Layer 3
  status.
- Reanchor is triggered by Codex protocol before substantial work; users are
  not asked to run detector or CLI commands manually.
- A callable detector command/tool can return `ReanchorResult` so Codex can
  avoid full anchor rereads during normal operation.
- File changes produce deterministic required reads instead of full-context
  rereads.
- Unknown, missing required, unreadable, or invalid required files are treated
  conservatively.
- Periodic reanchor can run every 10 conversation rounds using checkpoint round
  count and fingerprints.
- The detector does not decide business task scope; it only decides whether
  shared context must be refreshed or whether a management/handoff flow is
  required.

Failure Signals:

- Thread continues from stale contract after version changed.
- User must remember to request reanchor or manually run a CLI command before
  ordinary thread work.
- Detector only checks file existence, not version state.
- Detector silently treats unreadable or missing version files as unchanged.
- Detector expands into business validation or project management.
- Detector dumps file contents into chat instead of returning required reads.
- Detector stores chat history or business data in checkpoint state.
- No callable detector command/tool exists, so sessions rely on fallback anchor
  rereads rather than programmatic anchoring.

## 9. Reanchor Workflow

Owner: Templates / Protocol; Coordination for version semantics.

Inputs:

- Optional detector result, when a detector command/tool is available.
- `AGENTS.md`.
- `docs/anchor_pm/current_version.md`.
- `docs/anchor_pm/contracts.md`.
- Current thread's `docs/module_state/<thread>.md`.
- User task.

Outputs:

- Minimal anchor-state line.
- Thread identity.
- Current version.
- Scope and out-of-scope boundaries when relevant.
- Recovered current state.
- Decision whether task is in scope or requires handoff.

Acceptance:

- Codex automatically runs Reanchor Start before substantial work.
- Thread starts from current anchors, not stale chat memory.
- Reanchor cost stays small.
- Long-lived thread can resume after context reset.
- Users do not need to understand or manually operate the detector.

Failure Signals:

- Thread ignores updated module state.
- Thread waits for the user to ask for reanchor.
- Thread works outside contract without handoff.
- Reanchor requires reading a document pile.

## 10. Handoff Workflow

Owner: Templates / Protocol.

Inputs:

- Source thread identity.
- Target thread identity.
- Confirmed facts.
- Inference.
- Impact on current task.
- Unresolved questions.
- Suggested next step.

Outputs:

- Structured handoff summary.
- No target-thread decision made by source thread.
- Clear boundary of what not to repeat.

Acceptance:

- Cross-module dependency information moves without broad context dumping.
- Source thread preserves its own boundary.
- Target thread gets enough information to reanchor and continue.

Failure Signals:

- Source thread silently expands scope.
- Handoff includes unconfirmed inference as fact.
- Handoff is too vague for target thread to act.

## 11. Status / Drift / Review Workflows

Owner: Dogfood / Validation for evidence; Coordination for rule changes;
Templates / Protocol for workflow text.

Inputs:

- Current anchors.
- Generated package files.
- Module state files.
- Validation reports.
- User or tester feedback.

Outputs:

- Status check result.
- Drift findings.
- Review log entries for repeated failures.
- Suggested handoffs to owning thread.

Acceptance:

- Validation distinguishes blocking defects from product improvements.
- Repeated UX failures become product requirements.
- Rules are reduced or clarified rather than duplicated.

Failure Signals:

- One successful sample is treated as full validation.
- Review findings directly mutate contracts without Coordination review.
- UX confusion is dismissed as user error.

## 12. Validation Criteria

Anchor PM should continue development only if evidence supports real AI coding
experience improvement.

Primary validation questions:

- Can a first-time user understand whether Anchor PM fits their project?
- Can the user complete install without hidden placeholder work?
- Are generated thread boundaries easy to choose and understand?
- Does reanchor reduce repeated background explanation?
- Does module state reduce stale paths, stale commands, or stale assumptions?
- Does handoff reduce cross-thread confusion?
- Are generated docs short enough to maintain?

Stop or change direction if:

- Users repeatedly fail during install or thread creation.
- Users cannot understand why threads are split.
- State files become a second project brain.
- Anchor PM adds process without reducing repeated explanation, context
  correction, or synchronization cost.
- The product's value depends on users reading long manuals.
