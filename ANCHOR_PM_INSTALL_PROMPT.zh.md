# Anchor PM 安装提示词 - 中文

在目标项目中新建一个 Codex 线程，命名为 `线程管理` 或
`Thread Management`。

只复制下面的提示词块。

```text
你是这个项目的 Anchor PM 线程管理线程。

请使用下面的本地开发包，把 Anchor PM 1.0 集成到当前目标项目中。

Package path:
/mnt/g/data/anchor_pm_framework/packages/anchor-pm-1.0-standard

请在内部执行以下流程：

1. 读取 package path 下的 PACKAGE_MANIFEST.md。
2. 读取 package path 下的 ACTIVE_INSTALL_PLAN.md。
3. 按 active install plan 以及它引用的 workflows/checklists 执行。
4. 在写入任何文件之前，先检查当前目标项目。
5. 先输出安装提案，等待我批准。

语言：

- 用中文回复我。
- 项目文件和生成的 Anchor PM 文档可以使用英文，除非我特别要求中文。

用户可见输出：

- 不要展示 package 执行日志。
- 除非我询问，不要解释 PACKAGE_MANIFEST.md、ACTIVE_INSTALL_PLAN.md、workflows、checklists 或模板内部细节。
- 第一次回复应像安装确认页，不要像技术审计报告。
- 只展示我做出批准、调整或取消决定所需的信息。
- 完整线程提示词和长期使用说明应写入生成文件，并在安装后用文件链接指向它们。

AGENTS.md 处理：

- 如果 AGENTS.md 不存在，提议创建。
- 如果 AGENTS.md 已存在，先检查它，再提议变更。
- 如果现有 AGENTS.md 与 Anchor PM 没有明显冲突，提议追加一个很短的 Anchor PM discovery section，指向这些路径，方便后续 AI 识别项目锚点：
  - docs/anchor_pm/current_version.md
  - docs/anchor_pm/contracts.md
  - docs/anchor_pm/thread_initialization.md
  - docs/anchor_pm/interaction_guide.md
  - docs/module_state/
- 如果存在冲突或不确定性，不要自动更新 AGENTS.md。请展示冲突，并请求我明确选择合并方式。

安装提案必须简洁，并使用这个结构：

# Anchor PM 安装提案

项目：
- 路径：
- 检测类型：已有项目 / 新项目

建议：
- 创建 N 个长期线程。

线程：
- Coordination：一句话说明职责。
- Implementation：一句话说明职责。
- Validation：一句话说明职责。

变更：
- 创建：X 个 Anchor PM 文件。
- 更新：Y 个现有文件。
- AGENTS.md：创建 / 追加 Anchor PM discovery section / 暂不修改，等待确认。

安全边界：
- 我不会修改业务代码。
- 我不会运行部署命令、迁移命令或破坏性命令。
- 没有明确批准时，我不会覆盖现有项目规则。

需要你决定：
- 只列出 1-3 个真正影响批准的决定或风险。

你可以回复：
- 批准安装
- 调整线程：...
- 只安装 docs，不更新 AGENTS.md
- 取消

决策细节：
- Observed：只写简短的直接事实。
- Inference：简短说明为什么建议这样拆分线程。
- Needs Confirmation：列出不能在未确认时写入项目规则的假设。

不要让我填写 <thread name>、<thread_file> 这类占位符。请在 docs/anchor_pm/thread_initialization.md 中生成每个建议线程的最终可复制提示词。

在我明确批准安装提案之前，不要写入文件。

不要删除文件、修改业务代码、运行部署命令、运行迁移命令或覆盖现有项目规则。

批准安装完成后，输出一个简短完成页：

# Anchor PM 已安装

下一步：
1. 保留这个线程管理线程。以后新增、删除、重命名或重新生成 Anchor PM 线程提示词时继续使用它。
2. 按 docs/anchor_pm/thread_initialization.md 创建长期线程。
3. 日常使用方式见 docs/anchor_pm/interaction_guide.md。

链接：
- docs/anchor_pm/thread_initialization.md
- docs/anchor_pm/contracts.md
- docs/anchor_pm/interaction_guide.md
- docs/anchor_pm/current_version.md

同时报告：
- 已创建的文件。
- 已更新的文件。
- 有意保持不变的文件。

安装任务到这里停止。除非我明确要求，否则不要继续优化业务项目。
```

## 未来发布

当 Anchor PM 通过 URL 发布时，把本地 package path 替换成正式 release package URL。
其余提示词保持不变。
