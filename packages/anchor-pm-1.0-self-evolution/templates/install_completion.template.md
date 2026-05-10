# Template: Install Completion Main View

Use this template for the first reply after approved installation writes are
complete.

The completion page is a user-success page, not an internal file report.
Localize headings, instructions, and thread prompts to the user's install-prompt
language. Do not use English headings such as `Files Created` in a Chinese
install flow.

## English Shape

````markdown
# Anchor PM Installed Successfully

Your project now has a Threadsmith workspace. It gives you focused specialist
threads, shared contracts, and lightweight reanchor/knowledge-sync instructions
under one `.threadsmith/` folder.

Create the recommended threads:
1. Keep this conversation as Thread Management.
2. For each specialist below, start a new Codex conversation in the same project.
3. Copy that specialist's full prompt as the first message.

Recommended thread prompts:

### <thread name>
```text
<complete copy-paste-ready prompt in the install-prompt language>
```

### <thread name>
```text
<complete copy-paste-ready prompt in the install-prompt language>
```

All Threadsmith files are in [`.threadsmith/`](<absolute path>/.threadsmith/).
Ordinary Codex conversations that are not started with these prompts will not
automatically use Threadsmith.

This conversation remains your Thread Management thread. You can ask it to
create a new thread, change thread boundaries, remove a thread, regenerate
prompts, or query current thread information.
````

## Chinese Shape

````markdown
# Anchor PM 安装成功

现在你的项目已经拥有一个 Threadsmith 工作区：它会提供聚焦的项目专家线程、共享契约，以及轻量的重锚/知识同步规则，全部集中在 `.threadsmith/` 目录中。

创建推荐线程：
1. 保留当前对话作为线程管理入口。
2. 对下面每个项目专家，在同一个项目里新建一个 Codex 对话。
3. 复制该专家的完整提示词，作为新对话第一条消息发送。

推荐线程提示词：

### <线程名>
```text
<与安装语言一致的完整可复制提示词>
```

### <线程名>
```text
<与安装语言一致的完整可复制提示词>
```

所有 Threadsmith 文件都在 [`.threadsmith/`](<目标项目绝对路径>/.threadsmith/) 中。没有使用这些提示词创建的普通 Codex 对话，不会自动进入 Threadsmith 的线程管理体系。

当前对话就是线程管理线程。你可以在这里要求：创建新线程、修改线程边界、删除线程、重新生成提示词，或查询当前线程信息。
````

## Completion Rules

- Show the generated recommended thread prompts directly in the completion page.
  This is the user's next action and should not be hidden behind a file list.
- Also write the same complete prompts to `.threadsmith/thread_initialization.md`
  as a durable copy.
- The generated `thread_initialization.md` must contain full copy-paste-ready
  prompts in the install-prompt language.
- In a Chinese install flow, generated thread prompts should start in Chinese,
  for example `你是 <线程名> 线程...`, while conventional technical terms such
  as project names, Sans-IO, CLI, HTTP, API, JSON may remain in English.
- In an English install flow, generated thread prompts should be English.
- Do not show a file inventory by default. Detailed created/updated/untouched
  file lists and rationale belong in
  `.threadsmith/install_decision_record.md`.
- If file links are useful, show only the `.threadsmith/` folder and optionally
  `.threadsmith/thread_initialization.md`; do not dump every generated file.
- The final paragraph must explain isolation: Threadsmith files stay in
  `.threadsmith/`, and ordinary Codex conversations are unaffected unless they
  start with the generated prompts.
- The final paragraph must explain that the current conversation remains Thread
  Management and can create, remove, rename, adjust, regenerate, or query
  threads.
- Replace every placeholder in this template before user-visible output.
