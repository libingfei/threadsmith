# Template: Install Completion Main View

Use this template for the first reply after approved installation writes are complete.

Localize headings, instructions, and thread-creation guidance to the user's
install-prompt language. Do not use English headings such as `Files Created` in
a Chinese install flow.

## English Shape

```markdown
# Anchor PM Installed

Next: create project specialist threads
1. Open [thread_initialization.md](<absolute path>/docs/anchor_pm/thread_initialization.md).
2. Pick one specialist thread.
3. Start a new Codex conversation in the same target project.
4. Copy the complete prompt under that thread and send it as the first message.
5. Keep this Thread Management conversation for future thread changes.

Key links:
- [thread_initialization.md](<absolute path>/docs/anchor_pm/thread_initialization.md)
- [contracts.md](<absolute path>/docs/anchor_pm/contracts.md)
- [interaction_guide.md](<absolute path>/docs/anchor_pm/interaction_guide.md)
- [current_version.md](<absolute path>/docs/anchor_pm/current_version.md)
- [install_decision_record.md](<absolute path>/docs/anchor_pm/install_decision_record.md)

Change summary:
- Created: <short file/directory summary>
- Updated: <short file/directory summary or none>
- Left untouched: <short business-file summary>
```

## Chinese Shape

```markdown
# Anchor PM 已安装

下一步：创建项目专家线程
1. 打开 [thread_initialization.md](<目标项目绝对路径>/docs/anchor_pm/thread_initialization.md)。
2. 选择一个项目专家线程。
3. 在同一个目标项目中开启一个新的 Codex 对话。
4. 复制该线程下的完整提示词，作为第一条消息发送。
5. 保留当前线程管理对话；以后新增、删除、重命名或重新生成线程时回到这里。

常用链接：
- [thread_initialization.md](<目标项目绝对路径>/docs/anchor_pm/thread_initialization.md)
- [contracts.md](<目标项目绝对路径>/docs/anchor_pm/contracts.md)
- [interaction_guide.md](<目标项目绝对路径>/docs/anchor_pm/interaction_guide.md)
- [current_version.md](<目标项目绝对路径>/docs/anchor_pm/current_version.md)
- [install_decision_record.md](<目标项目绝对路径>/docs/anchor_pm/install_decision_record.md)

变更摘要：
- 已创建：<简短文件/目录摘要>
- 已更新：<简短文件/目录摘要，或“无”>
- 保持不变：<简短业务文件摘要>
```

## Completion Rules

- Do not paste every thread prompt into the chat by default. Link to
  `docs/anchor_pm/thread_initialization.md` and teach the user how to use it.
- If showing an example prompt, show at most one and only if it matches the
  install-prompt language.
- The generated `thread_initialization.md` must contain full copy-paste-ready
  prompts in the install-prompt language.
- In a Chinese install flow, generated thread prompts should start in Chinese,
  for example `你是 <线程名> 线程...`, while conventional technical terms such
  as project names, Sans-IO, CLI, HTTP, API, JSON may remain in English.
- In an English install flow, generated thread prompts should be English.
- Keep completion concise. Detailed rationale belongs in
  `docs/anchor_pm/install_decision_record.md`.
- Replace every placeholder in this template before user-visible output.
