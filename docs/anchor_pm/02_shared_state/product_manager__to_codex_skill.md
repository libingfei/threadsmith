# Product Manager To Codex Skill

Layer: `2`

Source thread: `Product Manager`

Target thread: `Codex Skill / Package Installer`

Status: active shared dependency.

## Confirmed Product Requirements

- Install and thread-creation flows should set the expectation that Codex, not
  the user, triggers Reanchor Start before substantial work.
- A future Skill may describe how to run and interpret reanchor, but it must not
  be treated as a guaranteed session-start hook.
- Installer wording should avoid telling users to run CLI commands before each
  thread interaction.
- If a target project has a detector command/tool, generated thread prompts
  should tell Codex to use it automatically.
- If no detector command/tool exists, generated thread prompts should preserve a
  degraded fallback that Codex performs itself by reading the required anchors.
- Skill or installer wording must not present fallback rereading as completed
  programmatic anchoring.

## Target Next Step

Update install prompt behavior and any future Skill guidance so automatic
Reanchor Start is included in generated thread prompts without adding a user
manual step.
