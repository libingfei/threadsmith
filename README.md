# Anchor PM

## English

Anchor PM turns scattered AI coding chats into anchored specialist threads.

Most AI coding tools optimize a single conversation. Anchor PM is for projects
that have outgrown one chat: it gives each long-lived Codex thread a clear
specialty, a small Markdown state file, and a handoff path to other threads.

### Best Fit

Use Anchor PM when your project has multiple active modules, recurring context
resets, changing file paths or commands, and more than one persistent Codex
conversation.

Skip it for one-off fixes, small scripts, or projects that fit comfortably in
one short chat.

### What You Get

- Focused specialist threads instead of one overloaded conversation.
- Reanchor prompts that tell each thread what to reread before work.
- Module state files for current paths, commands, decisions, risks, and open
  issues.
- Handoff summaries when one thread needs another specialist's context.
- Conservative conclusion rules that separate observed facts, inference, and
  unverified risks.

### Install

Open the target project in Codex and start a new installation/thread-management
conversation. Rename it `Thread Management` if your client supports that, then
paste one install prompt:

- [English install prompt](./ANCHOR_PM_INSTALL_PROMPT.en.md)
- [中文安装提示词](./ANCHOR_PM_INSTALL_PROMPT.zh.md)

The installer returns a short proposal before writing files. If it looks right,
reply:

- `Approve install`

The install prompts use the public GitHub package source:
`https://github.com/libingfei/threadsmith`, package directory
`packages/anchor-pm-1.0-standard`.

By default, generated Anchor PM files stay inside one `.threadsmith/` folder in
the target project.

### Example

A project starts as one feature, then grows into data import, backend rules,
model/reporting code, UI or API workflows, and validation. In one long chat,
Codex starts reusing stale paths and mixing module-specific assumptions.

With Anchor PM, those areas become specialist threads. Each thread rereads the
same project anchors, owns a smaller scope, updates its module state when facts
change, and hands off cross-module dependencies explicitly.

### Help Test It

Try Anchor PM on a real multi-thread Codex project. The most useful feedback is
whether installation is clear, thread boundaries are easy to choose, and
reanchor/module state reduces repeated explanations.

## 中文

Anchor PM 把分散的 AI 编程对话，收束成一组可重锚、可交接、可维护的专家线程。

多数 AI coding 工具在优化单个对话。Anchor PM 面向已经超过单个对话承载范围
的项目：它为每个长期 Codex 线程提供清晰专长、小型 Markdown 状态文件，以及
通向其他线程的 handoff 路径。

### 适合谁

如果你的项目有多个活跃模块、经常发生上下文重置、文件路径或命令会变化，并
且你会保留多个长期 Codex 对话，那么 Anchor PM 更可能有用。

如果只是一次性修复、小脚本，或一个短对话就能装下的项目，可以先不用。

### 你会得到什么

- 聚焦的专家线程，而不是一个过载的长对话。
- reanchor 提示词，让每个线程开工前知道该重读什么。
- 模块状态文件，用来记录当前路径、命令、决策、风险和未决问题。
- handoff 摘要，让一个线程能把依赖信息交给另一个专家线程。
- 保守结论规则，区分已观察事实、推断和未验证风险。

### 安装

在 Codex 中打开目标项目，开启一个安装/线程管理对话。如果客户端支持重命名，
可以命名为 `线程管理` 或 `Thread Management`，然后粘贴一个安装提示词：

- [English install prompt](./ANCHOR_PM_INSTALL_PROMPT.en.md)
- [中文安装提示词](./ANCHOR_PM_INSTALL_PROMPT.zh.md)

安装器会在写入文件前返回简短提案。如果看起来合适，回复：

- `批准安装`

安装提示词使用公开 GitHub package source：
`https://github.com/libingfei/threadsmith`，package directory 为
`packages/anchor-pm-1.0-standard`。

默认情况下，生成的 Anchor PM 文件会集中放在目标项目的 `.threadsmith/` 目录中。

### 案例

一个项目最初只有一个功能，后来增长出数据导入、后端规则、模型/报表代码、
UI 或 API 流程、验证检查。放在一个长对话里，Codex 开始复用过期路径，并混
用不同模块的假设。

使用 Anchor PM 后，这些区域会变成专家线程。每个线程重读同一套项目锚点，
负责更小范围，在事实变化时更新自己的模块状态，并显式 handoff 跨模块依赖。

### 参与测试

请在真实的多线程 Codex 项目中试用 Anchor PM。最有价值的反馈是：安装是否
清楚、线程边界是否好选，以及 reanchor/module state 是否减少了重复解释。
