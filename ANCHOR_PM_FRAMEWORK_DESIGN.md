# Anchor PM Framework 设计草案

## 1. 定位

`Anchor PM` 是面向 AI coding agents 的轻量项目协调层。

它不是业务开发代理，也不是全知项目大脑，而是一个“项目经理型协调代理”：

- 维护项目边界、线程职责、共享知识与交接机制。
- 帮助模块线程在短上下文内高质量工作。
- 帮助人类把长期项目知识从聊天记忆中外置到稳定锚点。
- 不替代人类的业务裁决，也不替代模块线程的技术细节判断。

核心目标：

- 降低长线程上下文遗忘。
- 降低多线程之间的信息割裂。
- 降低跨模块串线。
- 降低规则、状态、文档膨胀。
- 让任意线程都能通过少量锚点恢复当前上下文。

## 2. 从源项目抽象出的核心经验

Anchor PM 来自一个真实中大型项目的协作经验，但源项目不是通用模板本身。

可抽象的部分：

- 一个项目需要有最高规则锚点。
- 每个长期线程需要有明确职责、排除范围、验收标准和状态文件。
- 线程开始前需要检查契约版本，必要时重锚。
- 跨线程任务不应静默扩展，应生成交接摘要。
- 重要结论需要区分观察、推断和未验证。
- 状态变化不能只留在聊天里，结束前应判断是否回写。
- 复发性错误应进入版本 bug log 或复盘记录。
- 文档治理应持续减少规则源，而不是不断新增规则文档。

不可直接复用的源项目业务细节：

- 源项目的目录名、阶段名和数据层命名。
- 源项目的业务术语和业务流程含义。
- 源项目的运行环境、语言偏好和部署约束。
- 源项目的数据审计、训练、验证和生产规则。

抽象原则：

- 保留结构，不搬业务。
- 保留协议，不搬术语。
- 保留协作机制，不搬具体实现。

## 3. 角色模型

### Human / Domain Expert

职责：

- 提供需求。
- 提供行业判断。
- 做跨模块裁决。
- 决定上线、替代、发布和优先级。

特点：

- 可以直接向任意模块线程下发需求。
- 不必须通过 Anchor PM 转发所有任务。

### Anchor PM / Coordination Layer

职责：

- 维护项目宪法。
- 维护线程契约。
- 维护模块状态结构。
- 管理版本重锚。
- 维护跨线程交接格式。
- 审查文档漂移、重复和规则膨胀。
- 记录复发性机制问题。

不做：

- 不替模块线程做业务实现。
- 不替人类做最终业务裁决。
- 不强制所有任务经过自己路由。
- 不把事实层、日志层、案例层提升成上位规则。

### Module Threads / Specialist Agents

职责：

- 负责一个高内聚模块。
- 读取本模块契约和状态。
- 在本模块内执行代码、验证和局部判断。
- 越界时生成交接摘要。
- 收尾时判断是否更新状态文件。

## 4. 标准锚点层级

Anchor PM 的通用文件层级建议如下。

```text
<project>/
  AGENTS.md
  docs/
    README.md
    anchor_pm/
      current_version.md
      contracts.md
      interaction_guide.md
      review_log.md
      simplification.md
    module_state/
      <thread>.md
```

### `AGENTS.md`

项目最高执行规则。

只放：

- 环境边界。
- 权限边界。
- 失败策略。
- 结论协议。
- 安全约束。
- 项目级不可违背原则。

不放：

- 具体线程任务。
- 历史状态。
- 一次性排障结论。

### `docs/anchor_pm/current_version.md`

线程契约版本入口。

用途：

- 明确当前契约版本。
- 说明版本变化点。
- 指示是否需要重锚。

### `docs/anchor_pm/contracts.md`

线程契约总表。

每个线程固定字段：

- `Scope`
- `Out of Scope`
- `Acceptance`
- `Hard Rules`
- `State File`
- `Handoff Rule`

### `docs/module_state/<thread>.md`

每个线程一个状态文件。

固定章节：

- `Contract`
- `Current State`
- `Open Issues`
- `Runbook`
- `History / Notes`

### `docs/anchor_pm/interaction_guide.md`

人类与线程交互的最小话术。

只保留：

- 新线程绑定模板。
- 老线程重锚模板。
- 跨线程交接模板。
- 收尾状态回写提醒。

### `docs/anchor_pm/review_log.md`

复盘与版本 bug log。

只记录：

- 已观察到。
- 会重复。
- 值得下一版本优化。

不记录：

- 一次性猜测。
- 临时情绪。
- 低价值操作细节。

### `docs/anchor_pm/simplification.md`

简化治理规则。

只回答：

- 哪些是真共用。
- 哪些只是表面相似。
- 哪些是重复残留。
- 新规则能否并入现有主文档。

## 5. 核心协议

### 5.1 重锚协议

每个长期线程开始前执行：

1. 读取 `current_version.md`。
2. 若版本变化，重读：
   - `AGENTS.md`
   - `contracts.md`
   - 当前线程的 `module_state/<thread>.md`
3. 回复：
   - 当前线程身份
   - 当前版本
   - 当前范围
   - 排除范围

### 5.2 越界协议

当任务超出当前线程范围：

1. 当前线程不静默扩展执行。
2. 只记录对当前线程的影响。
3. 生成交接摘要。

交接摘要格式：

```text
来源线程：
目标线程：
当前结论：
已确认事实：
影响范围：
不需要重复讨论的内容：
需要目标线程处理的问题：
建议下一步：
```

### 5.3 状态回写协议

任务结束前判断：

- 是否改变长期状态。
- 是否产生新风险。
- 是否产生未决问题。
- 是否改变默认入口、命令或 runbook。

若是，更新对应 `module_state/<thread>.md`。

若否，明确回复：

```text
本次无需更新模块状态文件
```

### 5.4 高保守结论协议

重要结论默认拆成：

- `Observed`
- `Inference`
- `Unverified`

只有在正式验证、报告或产物已检查后，才允许说：

- `pass`
- `consistent`
- `ready`
- `safe`
- `can ship`

### 5.5 复盘协议

进入 review log 的问题必须满足：

- 已观察到，不只是猜测。
- 具有复发可能。
- 值得下一版本优化。
- 能提炼成机制问题，而不是单次操作噪音。

## 6. 工具化设计

Anchor PM 1.0 应是 Codex package-first，而不是复杂平台或 CLI-first 工具。

1.0 的核心交付物是一个普通目录包：

```text
packages/anchor-pm-1.0/
  PACKAGE_MANIFEST.md
  INSTALL_PROMPT.md
  ACTIVE_INSTALL_PLAN.md
  plans/
  templates/
  workflows/
  checklists/
```

Codex 读取包内 manifest、active install plan、workflow 和模板后，在目标项目中生成安装计划。用户确认后才写入锚点文件。

CLI 仍是合理的后续工具，但在 1.0 中降级为未来复现、检查和批量化能力。

### 6.1 包模式

1.0 有两种 release：

- `standard`：普通项目安装，部署后停止。
- `self-evolution`：Anchor PM 自身安装，部署后生成一份下一版本优化建议。

两种 release 的目录结构完全一致，只允许 `ACTIVE_INSTALL_PLAN.md` 和 `INSTALL_PROMPT.md` 不同。

`INSTALL_PROMPT.md` 允许按 release 不同，因为开发阶段必须写入具体本地包路径，standard 和 self-evolution 的路径不同。

Anchor PM 项目自身规模较小时，self-evolution 执行和 review 由 Coordination 线程统一管理，不额外拆分执行线程。

### 6.2 安装计划

安装计划必须规定：

- 先审计目标项目。
- 输出 `Observed / Inference / Needs Confirmation`。
- 输出拟创建文件、拟更新文件和冲突。
- 用户确认前不写入。
- 不删除文件、不修改业务代码、不运行部署命令。

### 6.3 模板与 workflow

1.0 通过模板生成：

- `AGENTS.md`
- `current_version.md`
- `contracts.md`
- `thread_initialization.md`
- `interaction_guide.md`
- `module_state.md`
- `review_log.md`
- `simplification.md`

workflow 至少覆盖：

- 已有项目接入。
- 新项目初始化。
- 线程数量建议和初始化提示词。
- self-evolution 优化建议。
- 重锚。
- 交接。
- 状态检查。

## 7. Codex Skill 设计

1.0 不要求先提供正式 Codex Skill。

1.0 的 Codex-first 入口是 `INSTALL_PROMPT.md`。用户在目标项目中新建 `Thread Management` / `线程管理` 线程后，粘贴已经包含包路径或 URL 的安装提示词。

后续可以提供一个通用 skill：

```text
anchor-pm
```

触发场景：

- 用户要求重锚当前线程。
- 用户要求新增线程。
- 用户要求检查文档漂移。
- 用户要求跨线程交接。
- 用户要求记录当前版本 bug。
- 用户要求初始化项目协作体系。

Skill 行为：

- 优先读取 Anchor PM package。
- 若未来 CLI 可用，可调用 CLI 复现 package 行为。
- 不新增平行规则源。
- 不把工具说明当项目规则。

## 8. 隔离策略

为避免污染既有业务项目，Anchor PM 必须支持三种隔离模式。

### 8.1 外部实验目录

默认模式。

生成到：

```text
../<project>_anchor_pm_draft/
```

特点：

- 不修改原项目。
- 适合初次审计。
- 适合用户确认前的草案。

### 8.2 项目内草案目录

可选模式。

生成到：

```text
docs/.anchor_pm_draft/
```

特点：

- 仍不作为正式规则源。
- 适合需要在同仓库 review 的团队。

### 8.3 正式接入模式

用户确认后才执行。

生成或更新：

```text
AGENTS.md
docs/anchor_pm/
docs/module_state/
```

必须满足：

- 有 diff。
- 有审计报告。
- 有可回滚备份。
- 用户显式确认。

## 9. 安全边界

Anchor PM 不应默认执行：

- 删除文件。
- 修改生产配置。
- 运行部署脚本。
- 改业务代码。
- 生成不可审计的大量文档。

默认只读审计。

写入时遵循：

- 先生成草案。
- 再展示 diff。
- 最后由用户确认 apply。

## 10. 最小可行版本

1.0 只实现：

- package manifest
- install prompt
- standard install plan
- self-evolution install plan
- templates
- workflows
- checklists
- dry-run validation reports

不实现：

- CLI
- 自动调度 agent。
- 自动修改业务代码。
- 自动发布。
- 自动裁决跨模块争议。

### 10.1 自迭代要求

Anchor PM 本仓库必须作为第一个接入对象。

在 CLI 实现前，允许手工维护最小锚点：

```text
AGENTS.md
docs/anchor_pm/
docs/module_state/
```

在 1.0 package 实现后，self-evolution 包必须能审计这些锚点并生成下一版本优化建议。

在后续 CLI 实现后，工具必须能检查这些锚点，并能生成可对比的 draft。

自迭代不是特殊业务逻辑。它只验证通用协议是否足够轻、足够清楚、足够可恢复。

当前项目的自迭代过程由 Coordination 线程统一管理。Dogfood / Validation 只保留验证记录和外部样本证据职责。

## 11. 成功标准

短期成功标准：

- 新项目能在 10 分钟内生成一套可审阅的锚点草案。
- 用户能看懂每个线程负责什么、不负责什么。
- 任意线程能通过 `reanchor` 找到应读文件。
- 不污染原项目。
- Anchor PM 本仓库能用自己的锚点完成后续协作。

中期成功标准：

- 长线程遗忘问题下降。
- 跨线程信息断裂下降。
- 规则重复文档减少。
- 状态回写更稳定。

失败信号：

- 文档数量快速膨胀。
- 工具生成了用户不信任的规则。
- 项目经理层开始替业务线程做业务判断。
- 所有任务都被迫经过 Anchor PM，形成瓶颈。

## 12. 与源项目的关系

源项目只作为参考案例，不作为公开框架的一部分。

不得直接依赖：

- 源项目的目录。
- 源项目的业务术语。
- 源项目的运行环境。
- 源项目的业务工作流。

可以提炼：

- 五线程结构的经验。
- 重锚协议。
- 状态文件结构。
- 交接模板。
- 复盘机制。
- 简化治理原则。

未来如需把源项目经验迁入 Anchor PM，应采用匿名案例附件形式，而不是把案例逻辑写入框架核心。

## 13. 下一步建议

建议下一轮只做两件事：

1. 用 `anchor-pm-1.0-self-evolution` 对本仓库生成 self-optimization report。
2. 用 `anchor-pm-1.0-standard` 对一个普通项目样本生成 standard install proposal。

暂不在源项目内接入 Anchor PM，避免污染当前业务项目。
