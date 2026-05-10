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

Lifecycle model:

- `Anchor Gate` runs before substantial work. It combines user-delta triage and
  Reanchor Start, but defaults to no write, no full reread, and no visible
  process explanation.
- `Knowledge Sync Gate` runs before the final answer. It updates local state,
  shared state, or handoff only when durable knowledge changed.
- Gate handling should stay smaller than task handling. It escalates only for
  changed, unknown, blocked, conflicting, or durable user-correction cases.

## Module Map

```text
User README / Prompt Entry
  -> Thread Management Installer
  -> Target Project Audit
  -> Thread Plan Generator
  -> Anchor File Generator
  -> Optional Root AGENTS.md Integration
  -> Post-Install Thread Creation Guide
  -> Anchor Gate
  -> Contract Version / Reanchor State Detector
  -> Reanchor / Handoff / Status Workflows
  -> Knowledge Sync Gate
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
- Public GitHub repository URL and package directory.
- User's preferred conversation language.

Outputs:

- User understands whether Anchor PM is a fit.
- User opens target project in Codex.
- User creates `Thread Management` / `线程管理` thread.
- User pastes one language-specific install prompt.
- The root prompt acts as a package index and points Codex to package-internal
  instructions and templates.

Acceptance:

- The entry path is copy-paste-ready.
- No user-filled placeholders are required for the package path in released
  prompts.
- README states the public package source used by the install prompts.
- Each root install prompt stays under 20 lines.
- Detailed install behavior lives in package files rather than in the root
  prompt.
- User can tell when Anchor PM is probably unnecessary.

Failure Signals:

- User asks what file to copy.
- User copies the wrong prompt language.
- User cannot tell whether to use Anchor PM for their project.
- Public install prompt points to a stale local package path.

## 2. Thread Management Installer

Owner: Codex Skill / Package Installer.

Inputs:

- User-pasted install prompt.
- Target project root.
- Package root.
- `PACKAGE_MANIFEST.md`.
- `ACTIVE_INSTALL_PLAN.md`.
- `INSTALL_PROMPT.md`.
- `templates/install_proposal.template.md`.
- `templates/install_completion.template.md`.
- Referenced install plan, workflows, checklists, and templates.

Outputs Before Approval:

- Extremely concise installation proposal, not a package execution log or audit
  report.
- Main visible proposal contains only:
  - target project path and detected type;
  - proposed project specialist thread names with one-sentence responsibilities;
  - localized reply options such as:
  - `Approve install`
  - `Adjust threads: ...`
  - `Cancel`
- Thread-name language follows the selected install prompt language.
- Generated thread creation prompts in `.threadsmith/thread_initialization.md`
  follow the selected install prompt language.
- File create/update counts, root `AGENTS.md` integration, `Observed`,
  `Inference`, and `Needs Confirmation` should not appear in the main visible
  proposal by default.
- Approval-relevant conflicts or risks may appear as one short `Note` line.
- Optional rationale should be omitted before approval unless there is a real
  approval-blocking risk or the user asks for it. If needed, use a collapsed
  details block; after approval, write details to
  `.threadsmith/install_decision_record.md`.
- Internal safety constraints should not be expanded as a visible proposal
  section unless the user asks.

Outputs After Approval:

- Localized installation success message.
- Brief feature summary: the target project now has a `.threadsmith/`
  workspace, recommended project specialist threads, shared contracts, and
  lightweight reanchor/knowledge-sync instructions.
- Clear next step: keep the current conversation as Thread Management, create a
  new Codex conversation for each recommended specialist, and paste that
  specialist's full prompt as the first message.
- Complete copy-paste-ready prompts for each recommended specialist thread.
- Isolation note: all Threadsmith files are under `.threadsmith/`; ordinary
  Codex conversations that are not started with the generated prompts are not
  pulled into Threadsmith behavior.
- Thread Management capability note: the current conversation can create,
  remove, rename, adjust, regenerate, or query threads.
- `.threadsmith/install_decision_record.md` containing detailed rationale and
  file change information.
- Generated-file inventory omitted from the main completion view by default.

Acceptance:

- Does not write before explicit approval.
- Does not expose package internals unless user asks.
- Does not continue optimizing the business project after standard install.
- Follows the user's conversation language.
- Completion page and thread creation prompts match the install prompt language.

Failure Signals:

- User must infer what to approve.
- Proposal is longer than needed for a decision.
- Installer shows manifest/workflow logs as primary output.
- Installer starts implementation work after installing anchors.
- Chinese install produces English thread creation prompts or English
  completion headings.
- Completion page links to files but fails to teach the user how to create the
  next Codex thread.
- Completion page forces the user to interpret a long generated-file list
  instead of showing the recommended next actions and thread prompts.

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
- Thread-management context:
  - the current install conversation is the Thread Management entrypoint;
  - Thread Management is not counted as a target-project business specialist;
  - Anchor PM internal roles such as `Coordination` must not be proposed as
    ordinary target-project specialist threads.
- Existing-project split rule:
  - derive threads from the target project's real modules, subsystems, or
    durable maintenance boundaries;
  - inspect source packages, runtime surfaces, APIs/plugins, CLI, docs, tests,
    config, and CI signals before proposing boundaries;
  - avoid broad function buckets such as `Implementation` or default standalone
    `Validation` unless the project structure genuinely supports them;
  - each module specialist should normally own code, tests, docs, and validation
    evidence for that module.
- New/empty project fallback:
  - starter threads are allowed only as provisional defaults;
  - they must be marked adjustable and should not be presented as discovered
    project modules.
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
- User-visible thread names use the same language as the selected install
  prompt, while technical terms may remain in their conventional language.
- Thread boundaries are easy to read.
- Cross-boundary work has a handoff path.
- Existing projects receive module/subsystem-based specialist threads, not
  Anchor PM internal threads.
- The first Codex message does not require the user to have renamed the
  conversation in advance.

Failure Signals:

- Installer proposes `Coordination` for an ordinary target project.
- Installer collapses most business code into one broad `Implementation` thread
  when observable module boundaries exist.
- Installer presents generic starter threads as discovered project structure.
- Installer generates thread names in the wrong language for the selected
  install prompt.
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

- `.threadsmith/AGENTS.md`.
- `.threadsmith/current_version.md`.
- `.threadsmith/contracts.md`.
- `.threadsmith/thread_initialization.md`.
- `.threadsmith/interaction_guide.md`.
- Optional support docs:
  - `.threadsmith/review_log.md`
  - `.threadsmith/simplification.md`
- `.threadsmith/module_state/<thread>.md` files.

Acceptance:

- Generated files are concise.
- Generated docs are project-neutral unless confirmed project facts are provided.
- Module state files are sparse state, not chat transcripts.
- Strong conclusion protocol is present.

Failure Signals:

- Generated docs are longer than the project can maintain.
- Multiple files repeat the same rules.
- Inference is written as confirmed project fact.

## 6. Optional Root AGENTS.md Integration

Owner: Codex Skill / Package Installer for behavior; Templates / Protocol for
wording; Coordination for rule conflicts.

Inputs:

- Existing `AGENTS.md`, if present.
- Proposed Anchor PM paths.
- Existing project rule hierarchy.
- User approval.

Outputs:

- Default: no root `AGENTS.md` create/update.
- If the user explicitly asks for global discovery: proposed short Anchor PM
  discovery section pointing to `.threadsmith/`.
- If conflict or uncertainty: conflict summary and explicit merge decision.

Acceptance:

- Existing rules remain authoritative unless user approves a merge.
- Default install is isolated: unrelated Codex conversations do not
  automatically enter Anchor PM behavior.
- Optional Anchor PM section helps AI agents discover anchors only when the user
  wants project-wide discovery.
- No business rules are invented.

Failure Signals:

- Installer replaces existing project policy.
- Anchor PM discovery section becomes a second project constitution.
- Conflict is hidden inside a large proposal.
- Root `AGENTS.md` is created or modified by default.

## 7. Post-Install Thread Creation Guide

Owner: Product Manager for flow; Templates / Protocol for generated document.

Inputs:

- Final thread plan.
- Generated contracts.
- Generated module state filenames.

Outputs:

- `.threadsmith/thread_initialization.md` with one complete prompt per thread.
- Installer completion message linking to:
  - `.threadsmith/thread_initialization.md`
  - `.threadsmith/contracts.md`
  - `.threadsmith/interaction_guide.md`
  - `.threadsmith/current_version.md`

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

## 9. Anchor Gate Workflow

Owner: Product Manager for UX and priority semantics; Templates / Protocol for
prompt wording; Reanchor Detector Core for future machine-readable pending
delta support.

Purpose: keep pre-work anchor handling correct but cheap. The gate prevents
stale anchors from overriding a fresh user correction, then runs Reanchor Start
only as far as needed for the task.

Inputs:

- Latest user message.
- Current thread identity and ownership boundaries.
- Known Layer 0 through Layer 3 ownership rules.
- Existing anchor state, when needed to apply a safe update.

Outputs:

- One of:
  - no visible anchor action needed;
  - safe local Layer 3 update or staged patch before reanchor;
  - Layer 2 shared-state update or handoff before reanchor;
  - pending high-priority user delta to carry into Reanchor Start;
  - minimal required-read list from the detector or fallback;
  - confirmation or owner-thread handoff for scope/framework changes.

Decision Rules:

- Default to no anchor write, no full reread, and no visible explanation.
- Apply this only to explicit durable user corrections, rule changes,
  preferences, or cross-thread facts. Ordinary task instructions do not become
  durable anchors by default.
- The latest user correction is a high-priority pending delta for the current
  turn. Do not let older anchor text silently override it.
- If the current thread owns the relevant state and the change is low-risk,
  update or stage the anchor before Reanchor Start.
- If the change affects Layer 0, Layer 1, another thread's Layer 3, or unclear
  shared state, create a handoff or ask for confirmation instead of silently
  editing.
- If the user correction conflicts with current anchors, surface the conflict
  briefly and treat the user correction as pending until reconciled.
- Show an anchor status line only when changed, blocked, unknown, degraded, or
  conflict-relevant.
- Anchor Gate cost should stay below task cost; if it expands, collapse to a
  handoff or confirmation instead of spending the turn on process.

Acceptance:

- User-discovered product or workflow corrections are not lost behind old
  reanchor output.
- Pre-work writes remain sparse and ownership-safe.
- The thread can distinguish durable corrections from transient task
  instructions.
- Normal turns proceed directly to the actual task without visible ritual.

Failure Signals:

- A thread reanchors into old rules and then ignores the user's explicit
  correction from the same message.
- Every user comment creates state churn.
- The gate output consumes more visible chat than the task itself.
- A thread rewrites scope, framework rules, or another thread's state without
  ownership or confirmation.

## 10. Reanchor Workflow

Owner: Templates / Protocol; Coordination for version semantics.

Inputs:

- Optional detector result, when a detector command/tool is available.
- `.threadsmith/AGENTS.md`.
- `.threadsmith/current_version.md`.
- `.threadsmith/contracts.md`.
- Current thread's `.threadsmith/module_state/<thread>.md`.
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

## 11. Knowledge Sync Gate Workflow

Owner: Templates / Protocol for default workflow text; Product Manager for UX;
Reanchor Detector Core for future programmatic closeout planning; Coordination
for cross-thread policy.

Purpose: ensure every long-lived thread checks whether the conversation created
new durable knowledge or changed existing knowledge before the final response,
without turning every response into process output.

This is the write-side counterpart to Anchor Gate. Anchor Gate keeps the thread
fresh before work; Knowledge Sync Gate preserves only durable new knowledge
before reply.

Inputs:

- Current thread identity and scope.
- Current task outcome.
- Files changed or decisions made in this conversation.
- User corrections, critiques, or confirmed product/business facts.
- Known Layer 2 shared-state dependencies and current Layer 3 module state.

Outputs:

- One of:
  - no durable state update needed, with no visible note by default;
  - Layer 3 local-memory update for the current thread;
  - Layer 2 shared-state update or directed handoff for affected threads;
  - handoff to Thread Management for Layer 1 thread-definition changes;
  - handoff to Coordination or framework owner for Layer 0/framework changes.
- Final response should mention the state update when it matters, without
  dumping internal logs.

Decision Rules:

- Local durable knowledge goes to the current thread's Layer 3 module state or
  category file.
- Knowledge other threads need goes to Layer 2 shared state or a structured
  handoff.
- Scope, ownership, thread list, or thread prompt changes go through Thread
  Management; ordinary specialist threads should not silently rewrite Layer 1.
- Framework baseline or global protocol changes go through Coordination or the
  owning framework thread.
- Temporary observations, unverified guesses, and one-off task details should
  not be promoted into durable state unless they affect future work.
- If no durable or shared knowledge changed, do not spend visible chat space on
  closeout status unless the user asked for it.
- Knowledge Sync Gate cost should stay below task-result explanation cost.

Acceptance:

- Every substantial long-lived thread closeout performs this check.
- Important user corrections become durable state instead of remaining only in
  chat history.
- Shared knowledge reaches affected threads without forcing them to reread all
  thread-local files.
- State files stay sparse; they do not become full chat transcripts.
- Normal turns with no durable change finish without visible closeout ritual.

Failure Signals:

- A thread repeats a previously corrected mistake because the correction was not
  written to its state.
- A cross-thread dependency changes but no Layer 2 update or handoff is made.
- A thread silently changes its own scope or another thread's scope.
- State files accumulate transient discussion instead of durable facts.
- Closeout narration crowds out the actual answer.

## 12. Handoff Workflow

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

## 13. Status / Drift / Review Workflows

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

## 14. Validation Criteria

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
