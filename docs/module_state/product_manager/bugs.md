# Product Manager Bugs

Layer: `3`

Thread: `Product Manager`

Current entries:

- Flask install dry-run exposed a product failure in the public install prompt:
  it produced generic Coordination / Implementation / Validation threads for an
  existing project, required a pre-named conversation, showed too much internal
  safety explanation, and offered a confusing docs-only/no-AGENTS option.
- Flask install retest still had a language/option defect: Chinese install
  prompt output used English thread names and exposed `Adjust AGENTS.md` as a
  default reply option.
