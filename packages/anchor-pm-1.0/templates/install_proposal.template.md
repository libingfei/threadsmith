# Template: Install Proposal Main View

Use this template for the first installer reply before any file writes.

Localize headings, thread names, and reply options to the user's install-prompt language.

```markdown
# Anchor PM Installation Proposal

Project: <absolute target path> (<existing project | new project>)

Threads:
- <project module or subsystem thread name>: <one-sentence responsibility>
- <project module or subsystem thread name>: <one-sentence responsibility>
- <project module or subsystem thread name>: <one-sentence responsibility>

Reply:
- <approve install phrase>
- <adjust threads phrase>: ...
- <cancel phrase>
```

Chinese shape:

```markdown
# Anchor PM 安装提案

项目：<目标项目绝对路径>（已有项目 / 新项目）

线程：
- <项目模块或子系统线程名>：<一句话职责>
- <项目模块或子系统线程名>：<一句话职责>
- <项目模块或子系统线程名>：<一句话职责>

回复：
- 批准安装
- 调整线程：...
- 取消
```

## Main-View Rules

- Show only project path/type, specialist threads, and reply options.
- Do not show file create/update counts in the main view.
- Do not show root `AGENTS.md` integration in the main view unless the user explicitly requested it and a real approval-blocking merge choice exists.
- Do not show `Observed / Inference / Needs Confirmation` in the main view by default.
- Do not show package execution logs, active-plan details, workflow internals, or internal safety explanations.
- Do not expose `Adjust AGENTS.md` or docs-only/no-AGENTS as default reply options.
- Do not use generic `Coordination / Implementation / Validation` threads for existing projects.
- Do not leave placeholders in generated files; placeholders in this template must be replaced before user-visible output.

## Optional Risk Line

Only when a real approval-blocking risk exists, add one short line after the thread list:

```markdown
Note: <one sentence risk that affects approval>
```

Chinese:

```markdown
注意：<一句话说明真正影响批准的风险>
```

## Optional Details

Do not include a details block by default. Use this only when a real conflict/risk exists or the user asks for rationale:

```markdown
<details>
<summary>View rationale and file changes</summary>

- Files to create/update.
- Root AGENTS.md integration: not changed by default / requested merge pending confirmation.
- Observed: short direct facts.
- Inference: short rationale for the proposed split.
- Needs Confirmation: assumptions that must not become project rules without approval.

</details>
```

After approval, write the full rationale and file changes to `.threadsmith/install_decision_record.md`.
