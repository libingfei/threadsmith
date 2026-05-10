# Anchor PM Thread Initialization

This document tells users how many Codex threads to create and what prompt to paste into each one.

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

For each thread, include:

- the exact thread name;
- the exact module state file path;
- a one-sentence scope summary from `docs/anchor_pm/contracts.md`;
- lightweight Anchor Gate behavior before work;
- lightweight Knowledge Sync Gate behavior before final response;
- the cross-thread handoff rule.

{{THREAD_PROMPTS}}
