# Anchor PM Thread Initialization

This document tells users how many Codex threads to create and what prompt to paste into each one.

The installer must localize user-facing thread names, responsibilities, and
copy-paste prompts to the install-prompt language. Package docs may be English,
but the prompts users paste into Codex must use the user's install language.

## Recommended Thread Count

For existing projects, derive threads from real project modules, subsystems, or
durable maintenance boundaries.

For new or nearly empty projects, start with a small adjustable set:

- `Project Direction`
- `Core Implementation`
- `Quality and Release`

Use up to 5 threads when responsibilities are clearly separate.
Existing projects may need more than 5 when module boundaries are strong, but
avoid speculative threads.

Avoid creating speculative threads.
Thread Management is handled by the installation conversation; do not include a
default `Coordination` thread for ordinary target projects.

## Proposed Threads

{{THREAD_LIST}}

## Thread Prompts

The installer must replace this section with one complete prompt per proposed thread.

Each final prompt must be ready to copy and paste. Do not leave `<thread name>`, `<thread_file>`, `{{...}}`, or similar user-filled placeholders in the generated target document.

For a Chinese install flow, each prompt should be written in Chinese and may
retain project names and conventional technical terms such as Sans-IO, CLI,
HTTP, API, or JSON. For an English install flow, prompts should be English.

For each thread, include:

- the exact thread name;
- the exact module state file path;
- a one-sentence scope summary from `docs/anchor_pm/contracts.md`;
- lightweight Anchor Gate behavior before work;
- lightweight Knowledge Sync Gate behavior before final response;
- the cross-thread handoff rule.

## Prompt Language Patterns

Chinese install flow:

```text
你是 <线程名> 线程，负责 <目标项目路径> 中的 <一句话 scope>。
开工前运行 Anchor Gate：默认静默；如锚点 changed / blocked / unknown / conflicting / degraded，或用户给出持久纠偏，再读取必要锚点。
需要读取的锚点：AGENTS.md、docs/anchor_pm/current_version.md、docs/anchor_pm/contracts.md、docs/module_state/<thread>.md。
Scope：<本线程负责什么>。
Out-of-scope：<不属于本线程的事项>。
回复前运行 Knowledge Sync Gate：只在产生持久本地知识、共享知识或 handoff 时更新/说明；没有持久变化时保持静默。
跨线程事项请生成 handoff，不要直接扩展职责。
```

English install flow:

```text
You are the <thread name> thread for <target project path>.
Before work, run Anchor Gate silently unless changed, blocked, unknown, conflicting, degraded, or the user gives a durable correction.
Read required anchors: AGENTS.md, docs/anchor_pm/current_version.md, docs/anchor_pm/contracts.md, and docs/module_state/<thread>.md.
Scope: <what this thread owns>.
Out-of-scope: <what this thread must hand off>.
Before final response, run Knowledge Sync Gate: update or hand off only durable local/shared knowledge; otherwise keep it silent.
For cross-thread work, produce a handoff instead of expanding scope.
```

{{THREAD_PROMPTS}}
