# Anchor PM 1.0 MVP 规格

## 1. 目标

Anchor PM 1.0 采用 Codex package-first 路线。

1.0 的目标不是先交付跨平台 CLI，而是交付一个 Codex 可读取、可执行安装计划的项目包。用户在目标项目中新建 `Thread Management` / `线程管理` 线程，粘贴带有 Anchor PM 包路径或 URL 的安装提示词，Codex 按包内计划完成审计、安装计划生成和经确认后的锚点部署。

1.0 只解决 5 个问题：

- 项目需要哪些长期线程
- 每个线程负责什么
- 每个线程启动时该读什么
- 跨线程如何交接
- 状态和文档是否发生漂移

1.0 不解决：

- Go CLI 或跨平台二进制分发
- RAG / 向量检索
- Web UI
- 自动调度 agent
- 自动修改业务代码
- 自动跑 CI / deploy
- 自动替用户裁决上线

CLI 是后续复现、检查和批量化工具，不是 1.0 主路径。

## 2. 交付物

1.0 提供一个统一包源目录和两个 release 目录：

```text
packages/anchor-pm-1.0/
packages/anchor-pm-1.0-standard/
packages/anchor-pm-1.0-self-evolution/
```

两个 release 目录结构必须一致，只允许 `ACTIVE_INSTALL_PLAN.md` 和 `INSTALL_PROMPT.md` 不同。

standard 模式：

- 面向普通项目
- 完成 Anchor PM 集成后停止
- 不继续优化业务项目

self-evolution 模式：

- 仅面向 Anchor PM 自身
- 部署或刷新锚点后继续生成一份 self-optimization report
- 不自动修改产品文档、源码或包文件
- 由 Anchor PM 的 Coordination 线程统一管理，不单独拆出执行线程

## 3. 包结构

1.0 包结构固定为：

```text
PACKAGE_MANIFEST.md
INSTALL_PROMPT.md
ACTIVE_INSTALL_PLAN.md
plans/
  standard_project_install.md
  self_evolution_install.md
templates/
  AGENTS.template.md
  current_version.template.md
  contracts.template.md
  thread_initialization.template.md
  interaction_guide.template.md
  module_state.template.md
  review_log.template.md
  simplification.template.md
workflows/
  existing_project_adoption.md
  new_project_bootstrap.md
  self_optimization.md
  handoff.md
  reanchor.md
  status_check.md
checklists/
  safety_check.md
  conclusion_check.md
  drift_check.md
```

## 4. 安装流程

用户路径：

1. 用户在目标项目中新建 `Thread Management` / `线程管理` 线程。
2. 用户粘贴 `INSTALL_PROMPT.md` 或 `docs/anchor_pm/thread_management_install_prompt.md` 中的提示词。
3. 提示词必须包含具体 Anchor PM package path 或 URL；开发阶段使用本地路径，发布阶段使用 release URL。
4. Codex 读取 `PACKAGE_MANIFEST.md`。
5. Codex 读取并执行 `ACTIVE_INSTALL_PLAN.md`。
6. Codex 审计目标项目。
7. Codex 输出安装计划。
8. 用户确认后 Codex 才写入文件。

写入前必须输出：

- Target project path
- Detected mode: existing project or new project
- `Observed`
- `Inference`
- `Needs Confirmation`
- Recommended thread count
- Thread initialization prompts
- Proposed file creates
- Proposed file updates
- Conflicts and merge risks
- Explicit approval request

## 5. 安全规则

1.0 默认不污染项目。

安装线程不得：

- 删除文件
- 修改业务代码
- 运行部署命令
- 运行迁移命令
- 覆盖已有规则文件
- 把未确认推断写成正式契约

如果目标项目已有 `AGENTS.md` 或其他规则源，必须先输出冲突和合并计划。

即使是新项目，也必须先展示文件列表和线程列表，再写入。

## 5.1 用户不填占位符原则

Anchor PM 1.0 必须让用户直接复制可用提示词。

安装线程应根据候选线程名、状态文件名和契约内容，生成每个线程的完整初始化提示词。

最终写入目标项目的 `docs/anchor_pm/thread_initialization.md` 不应要求用户填写 `<thread name>`、`<thread_file>` 或类似占位符。占位符只允许存在于包模板中，作为 Codex 安装线程生成最终文档的输入。

同样，提供给 `Thread Management` / `线程管理` 线程的安装提示词必须指定包位置。开发阶段使用本地路径；发布阶段可以替换成下载链接或连接器地址。

## 5.2 交互语言原则

Anchor PM 的项目文档和包文档可以使用英语，但 Codex 线程与用户的交互语言应跟随用户惯用语言。

安装提示词必须明确：

- 如果用户用中文交流，安装线程用中文回复。
- 如果用户用其他语言交流，安装线程跟随该语言。
- 不因为 package 文档是英文就强制用英文交互。
- 生成到项目里的 Anchor PM 文档默认可用英文，除非用户要求本地化。

## 6. 目标锚点

安装后目标项目可以包含：

```text
AGENTS.md
docs/anchor_pm/current_version.md
docs/anchor_pm/contracts.md
docs/anchor_pm/thread_initialization.md
docs/anchor_pm/interaction_guide.md
docs/anchor_pm/review_log.md
docs/anchor_pm/simplification.md
docs/module_state/<thread>.md
```

必需模板：

- `AGENTS.template.md`
- `current_version.template.md`
- `contracts.template.md`
- `thread_initialization.template.md`
- `interaction_guide.template.md`
- `module_state.template.md`
- `review_log.template.md`
- `simplification.template.md`

## 7. Existing Project Adoption

已有项目接入必须先审计，再计划，再确认写入。

审计对象：

- README
- AGENTS 或其他规则文件
- docs
- scripts
- config
- tests
- CI / deploy 文件

审计输出必须分层：

- `Observed`
- `Inference`
- `Needs Confirmation`

已有项目默认保留已有规则源，不直接覆盖。

## 8. New Project Bootstrap

新项目初始化可以使用默认短线程集合：

- `Coordination`
- `Implementation`
- `Validation`

如果用户提供线程列表，以用户列表为准。

新项目也必须先展示计划，再写入。

## 9. Self-Evolution

Anchor PM 自身使用 self-evolution release。

因为 Anchor PM 1.0 本身不复杂，自迭代过程由当前 Coordination 线程统一管理。Dogfood / Validation 只记录验证证据和外部样本结果。

流程：

1. 用 `anchor-pm-1.0-self-evolution` 读取 Anchor PM 当前项目。
2. 检查或刷新 Anchor PM 锚点。
3. 运行 `workflows/self_optimization.md`。
4. 生成 self-optimization report。
5. 停止。

报告必须包含：

- `Observed`
- `Inference`
- `Unverified`
- Candidate `Sn -> Sn+1` improvements
- Blocking issues
- Non-blocking risks
- Suggested handoffs

self-evolution 模式一次只跑一轮，不自动进入下一轮。

## 10. 验收标准

功能验收：

- standard package 结构完整
- self-evolution package 结构完整
- 两种 release 目录结构一致
- 两种 release 只在 `ACTIVE_INSTALL_PLAN.md` 和 `INSTALL_PROMPT.md` 分流
- 安装提示词可指导 Codex 读取包
- 模板覆盖目标锚点
- workflow 覆盖已有项目、新项目、重锚、交接、状态检查、自优化

质量验收：

- 1.0 不要求用户安装 CLI 或运行时
- 写入前必须生成安装计划
- 已有规则不被静默覆盖
- 未确认推断不进入正式契约
- standard 模式部署后停止
- self-evolution 模式只生成优化建议

失败条件：

- 工具包要求所有任务都经过 Anchor PM
- 安装流程默认修改业务代码
- self-evolution 自动修改产品文档
- 模板制造大量低价值文档
- 用户无法在 10 分钟内理解线程结构

## 11. 后续版本

1.1 之后可考虑：

- Go CLI 复现 package 行为
- `anchorpm status`
- `anchorpm reanchor`
- `anchorpm init --mode existing`
- release 打包脚本
- Codex Skill / Plugin 包装

这些不是 1.0 的交付前置条件。
