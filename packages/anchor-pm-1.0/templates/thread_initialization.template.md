# Anchor PM Thread Initialization

This document tells users how many Codex threads to create and what prompt to paste into each one.

## Recommended Thread Count

Start with 3 threads unless the project clearly needs more:

- `Coordination`
- `Implementation`
- `Validation`

Use up to 5 threads when responsibilities are clearly separate.

Avoid creating speculative threads.

## Proposed Threads

{{THREAD_LIST}}

## Thread Prompts

The installer must replace this section with one complete prompt per proposed thread.

Each final prompt must be ready to copy and paste. Do not leave `<thread name>`, `<thread_file>`, `{{...}}`, or similar user-filled placeholders in the generated target document.

For each thread, include:

- the exact thread name;
- the exact module state file path;
- a one-sentence scope summary from `docs/anchor_pm/contracts.md`;
- the cross-thread handoff rule.

{{THREAD_PROMPTS}}
