# Checklist: Safety

Before writing files, confirm:

- Target root is explicit.
- Installation mode is clear.
- Proposed file creates are tracked in optional details or the install decision
  record.
- Proposed file updates are tracked in optional details or the install decision
  record.
- Existing rule files are identified.
- Conflicts are listed.
- User approval has been requested.

Never do these by default:

- delete files;
- copy the Threadsmith repository into the target project;
- modify business code;
- run deploy commands;
- run migrations;
- create or overwrite root `AGENTS.md`;
- promote unconfirmed inference into contracts.
