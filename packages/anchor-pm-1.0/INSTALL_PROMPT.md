# Anchor PM 1.0 Standard Install Instructions

This file is package-internal guidance for Codex after a root install prompt points here.
Do not ask the user to paste this full file unless debugging the package.

## Required Behavior

1. Reply in the user's install-prompt language.
2. Treat the current conversation as Thread Management for installation and future thread changes.
3. Inspect the target project before writing files.
4. Do not write until the user explicitly approves installation.
5. Use `templates/install_proposal.template.md` for the first installation proposal.
6. After approval, write detailed rationale and file changes to `.threadsmith/install_decision_record.md`.
7. Use `templates/install_completion.template.md` for the first completion reply.
8. Stop after installation unless the user asks for follow-up work.

If package source is not already available, clone or fetch Threadsmith outside
the target project. Do not copy the Threadsmith repository into the target
project.

Default install root: `.threadsmith/`. Keep all Anchor PM files under that one
directory unless the user explicitly asks for root-level discovery.

Do not create or modify the target project's root `AGENTS.md` by default. This
keeps Anchor PM isolated: ordinary Codex conversations that do not use the
generated thread prompts should not automatically see the thread-management
anchors.

## Project Classification

- `existing`: the target has source, docs, tests, rules, scripts, or project history.
- `new`: the target is empty or nearly empty.

For existing projects, proposed threads must come from real modules, subsystems,
runtime surfaces, docs/support surfaces, or durable maintenance boundaries.
Do not use generic `Coordination / Implementation / Validation` threads for
existing projects.

For new projects, propose starter threads only as adjustable placeholders.
Avoid `Coordination` because Thread Management is already handled by the
installation conversation.

## Proposal Main View

The user-visible proposal should show only:

- project path and detected type;
- project specialist thread list with one-sentence responsibilities;
- reply options: approve install, adjust threads, cancel.

Do not show by default:

- file create/update counts;
- root `AGENTS.md` integration;
- `Observed / Inference / Needs Confirmation`;
- package execution logs;
- active-plan or workflow internals;
- internal safety explanations;
- `Adjust AGENTS.md` or root `AGENTS.md` options;
- docs-only/no-AGENTS partial-install options.

Only discuss root `AGENTS.md` when the user explicitly asks for global
discovery or when a requested root integration creates an approval-blocking
merge choice.

## Generated Files

Generated `.threadsmith/thread_initialization.md` must contain complete,
copy-paste-ready prompts for every proposed long-lived thread. Do not leave
`<thread name>`, `<thread_file>`, `{{...}}`, or other user-filled placeholders.

User-facing thread names and all thread creation prompts in
`thread_initialization.md` must match the install-prompt language. Generated
project docs may otherwise remain English, but the prompts users paste into new
Codex conversations are user-facing and must not silently switch language.

For a Chinese install flow, write thread prompts in Chinese, for example
`你是 <线程名> 线程，负责 <目标项目路径> 中的 <scope>...`. Project names and
conventional technical terms such as Sans-IO, CLI, HTTP, API, JSON can stay in
English.

For an English install flow, write thread prompts in English.

Every generated long-lived thread prompt must include:

- scope and out-of-scope boundaries;
- lightweight Anchor Gate before work;
- lightweight Knowledge Sync Gate before final response;
- cross-thread handoff behavior.

Anchor Gate should stay silent by default and only expand for changed,
blocked, unknown, conflicting, degraded, or durable-correction cases.
Knowledge Sync Gate should write or hand off only durable local/shared
knowledge; when no durable knowledge changed, keep it silent.

## Completion Main View

The completion reply must be localized and feel like a successful setup page,
not an internal file report. It should:

- start by telling the user the installation succeeded;
- briefly explain what they now have: a `.threadsmith/` workspace, project
  specialist threads, shared contracts, and lightweight reanchor/knowledge-sync
  instructions;
- explain that each recommended specialist thread should be created as a new
  Codex conversation by copying that thread's full prompt;
- show the complete copy-paste-ready prompt for each recommended thread in the
  completion reply;
- remind the user to keep the current Thread Management conversation for future
  thread changes;
- explain that all Threadsmith files live under `.threadsmith/` and ordinary
  Codex conversations that do not use the generated prompts are unaffected;
- tell the user they can ask this Thread Management conversation to create new
  threads, change thread boundaries, delete threads, regenerate prompts, or
  query thread information.

Do not show a generated-file inventory by default. Detailed file changes and
rationale belong in `.threadsmith/install_decision_record.md`.
