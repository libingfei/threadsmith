# Threadsmith MVP 手动验收手册

本手册用于手动验证 Threadsmith MVP 是否真的形成一个可用的 AI coding 协作闭环，而不只是完成安装。

MVP 核心链路：

```text
接入项目
-> 生成项目专家线程
-> 提出一个问题
-> Anchor Gate：工作前轻量处理用户纠偏和必要锚点刷新
-> 在线程边界内解决问题
-> 必要时 handoff
-> Knowledge Sync Gate：回复前只在必要时写回或传播新知识
-> 下一轮从更新后的锚点继续
```

## 测试范围

本轮验证 Threadsmith MVP 的 package-first 体验：

- 从 public GitHub package source 接入真实开源项目。
- 安装前必须有确认提案，未批准前不得写入。
- 线程必须按目标项目模块/子系统/长期维护边界生成。
- 生成的线程提示词必须可直接复制使用。
- 每个长期线程必须包含轻量 `Anchor Gate` 和 `Knowledge Sync Gate`。
- 能完成一个小问题的 read-only 解决链路，并把必要知识同步到合适锚点。

本轮不把程序化 `anchorpm reanchor` 作为 blocker。若 detector/CLI 不可用，线程可以走降级读取路径，但必须明确表现为 degraded fallback，不能声称程序化锚定已完成。

## 测试前 GitHub 同步门禁

如果本轮要测试 public GitHub 链接，开始前必须确认远端已经是最新版本：

1. 在 Threadsmith 仓库运行 `git status --short --branch`。
2. 如果有本轮测试相关的本地修改，先 commit 并 push 到 `origin/main`。
3. 运行 `git ls-remote origin refs/heads/main`，确认远端提交等于本地 `HEAD`。
4. 记录本轮测试使用的 commit SHA。
5. 如果无法推送或不应推送，本轮必须改用本地提示词链接，不要用 GitHub 链接测试。

## 测试项目

| 测试档位 | 项目 | 本地目录 |
| --- | --- | --- |
| 小型 Python 项目 | pallets/flask | [/mnt/g/data/threadsmith_test_flask](/mnt/g/data/threadsmith_test_flask) |
| 中型 TS Monorepo | vitejs/vite | [/mnt/g/data/threadsmith_test_vite](/mnt/g/data/threadsmith_test_vite) |
| 可管理大型项目 | django/django | [/mnt/g/data/threadsmith_test_django](/mnt/g/data/threadsmith_test_django) |

Threadsmith 公共安装包来源：

```text
https://github.com/libingfei/threadsmith
packages/anchor-pm-1.0-standard
```

快速入口：

- [Threadsmith GitHub 仓库](https://github.com/libingfei/threadsmith)
- [标准安装包目录](https://github.com/libingfei/threadsmith/tree/main/packages/anchor-pm-1.0-standard)
- [中文安装提示词](https://github.com/libingfei/threadsmith/blob/main/ANCHOR_PM_INSTALL_PROMPT.zh.md)
- 本地中文安装提示词：[ANCHOR_PM_INSTALL_PROMPT.zh.md](/mnt/g/data/anchor_pm_framework/ANCHOR_PM_INSTALL_PROMPT.zh.md)
- 还原脚本：[threadsmith_restore_tests.sh](/mnt/g/data/threadsmith_restore_tests.sh)

如果测试 public GitHub 流程，先确认 GitHub 上的中文安装提示词已经包含最新规则：不要求首轮先命名线程、不默认 `Coordination / Implementation / Validation`、中文提示词下使用中文线程名、不默认提供 `调整 AGENTS.md`、包含轻量 `Anchor Gate` 和 `Knowledge Sync Gate`。如果还没有 push 最新版本，先用本地中文安装提示词验证产品逻辑。

## 安装后锚点文件快速入口

这些文件需要在完成对应项目安装后才会存在。测试时按当前项目点击对应列。

| 用途 | Flask | Vite | Django |
| --- | --- | --- | --- |
| 项目规则入口 | [AGENTS.md](/mnt/g/data/threadsmith_test_flask/AGENTS.md) | [AGENTS.md](/mnt/g/data/threadsmith_test_vite/AGENTS.md) | [AGENTS.md](/mnt/g/data/threadsmith_test_django/AGENTS.md) |
| 线程提示词 | [thread_initialization.md](/mnt/g/data/threadsmith_test_flask/docs/anchor_pm/thread_initialization.md) | [thread_initialization.md](/mnt/g/data/threadsmith_test_vite/docs/anchor_pm/thread_initialization.md) | [thread_initialization.md](/mnt/g/data/threadsmith_test_django/docs/anchor_pm/thread_initialization.md) |
| 线程契约 | [contracts.md](/mnt/g/data/threadsmith_test_flask/docs/anchor_pm/contracts.md) | [contracts.md](/mnt/g/data/threadsmith_test_vite/docs/anchor_pm/contracts.md) | [contracts.md](/mnt/g/data/threadsmith_test_django/docs/anchor_pm/contracts.md) |
| 交互指南 | [interaction_guide.md](/mnt/g/data/threadsmith_test_flask/docs/anchor_pm/interaction_guide.md) | [interaction_guide.md](/mnt/g/data/threadsmith_test_vite/docs/anchor_pm/interaction_guide.md) | [interaction_guide.md](/mnt/g/data/threadsmith_test_django/docs/anchor_pm/interaction_guide.md) |
| 当前版本 | [current_version.md](/mnt/g/data/threadsmith_test_flask/docs/anchor_pm/current_version.md) | [current_version.md](/mnt/g/data/threadsmith_test_vite/docs/anchor_pm/current_version.md) | [current_version.md](/mnt/g/data/threadsmith_test_django/docs/anchor_pm/current_version.md) |
| 线程状态目录 | [module_state/](/mnt/g/data/threadsmith_test_flask/docs/module_state) | [module_state/](/mnt/g/data/threadsmith_test_vite/docs/module_state) | [module_state/](/mnt/g/data/threadsmith_test_django/docs/module_state) |

## 每个项目的完整测试链路

每个测试项目都完整执行一次下面流程。

### 0. 测试前还原

在开始某个项目测试前，先还原三个测试项目，避免上一次测试残留影响结果：

```bash
/mnt/g/data/threadsmith_restore_tests.sh
```

确认目标项目 `git status --short` 为空。

### 1. 启动安装/线程管理对话

1. 在 Codex 中打开目标项目目录。
2. 开启一个安装/线程管理对话。如果 Codex 客户端支持重命名，可以命名为 `线程管理`；如果不支持，直接继续。
3. 粘贴 [中文安装提示词](https://github.com/libingfei/threadsmith/blob/main/ANCHOR_PM_INSTALL_PROMPT.zh.md) 的完整提示词。
4. 第一轮只观察安装提案，不批准。

提案必须包含：

- 当前项目路径。
- 已有项目 / 新项目判断。
- 每个专家线程的一句话职责。
- 线程名称语言与安装提示词语言一致。中文提示词下应使用中文线程名称，技术名词可保留英文。
- [AGENTS.md](#安装后锚点文件快速入口) 处理策略。
- 回复选项：`批准安装`、`调整线程：...`、`取消`。

提案不应包含：

- `建议`、`变更`、`需要你决定`、`决策细节` 作为默认大段落。
- 文件创建/更新数量作为主视图内容。
- `Observed / Inference / Needs Confirmation` 作为主视图内容；如需要，应折叠或写入安装后的细节文档。
- `Coordination / Implementation / Validation` 作为已有项目默认线程。
- 大段 package 执行日志。
- 大段内部安全声明。
- `调整 AGENTS.md` 作为默认回复选项。
- `只安装 docs，不更新 AGENTS.md` 这类默认半集成选项。
- 要求用户手动填写 `<thread name>`、`{{...}}` 等占位符。

### 2. 批准安装

如果提案可接受，回复：

```text
批准安装
```

安装完成页必须包含：

- 已创建文件。
- 已更新文件。
- 有意保持不变的文件。
- [docs/anchor_pm/thread_initialization.md](#安装后锚点文件快速入口)。
- [docs/anchor_pm/contracts.md](#安装后锚点文件快速入口)。
- [docs/anchor_pm/interaction_guide.md](#安装后锚点文件快速入口)。
- [docs/anchor_pm/current_version.md](#安装后锚点文件快速入口)。
- `docs/anchor_pm/install_decision_record.md`，用于保存判断依据和文件变更细节。

安装后检查：

- 没有修改业务代码、测试、配置或原有项目文档正文。
- [AGENTS.md](#安装后锚点文件快速入口) 不覆盖已有规则；若原来不存在，可以创建。
- [thread_initialization.md](#安装后锚点文件快速入口) 中每个线程提示词可直接复制。
- 每个线程提示词同时包含 `Anchor Gate` 和 `Knowledge Sync Gate`。

### 3. 创建项目专家线程

1. 打开 [thread_initialization.md](#安装后锚点文件快速入口)。
2. 选择一个安装器推荐的项目专家线程。
3. 新建 Codex 对话，粘贴该线程完整提示词。

第一轮输入：

```text
请先说明你的 scope 和 out-of-scope。然后简短说明你在工作前如何执行 Anchor Gate，以及回复前如何执行 Knowledge Sync Gate。不要展开流程日志。
```

通过标准：

- 能说明自身职责边界。
- 能说明前置 `Anchor Gate` 默认静默、最小读取，只在用户明确纠偏、锚点变更/未知/冲突/降级时放大。
- 能说明收尾 `Knowledge Sync Gate` 只在有持久知识时写回/传播；没有持久变化时不应占用明显回复空间。
- 若 detector/CLI 不可用，能明确说明 degraded fallback，而不是声称程序化锚定完成。

### 4. 提出并解决一个小问题

在项目专家线程中提出一个只读问题，测试它是否能在边界内解决问题，并在收尾时同步知识。

通用测试输入：

```text
请解决一个小问题：我想知道如果以后在你负责的模块中做一个很小的行为变更，最小需要关注哪些项目文件和验证入口。

要求：
- 不修改业务代码、测试、配置或普通项目文档。
- 只读取与你 scope 相关的最小必要文件。
- 给出 Observed / Inference / Unverified。
- 如果发现一个对本线程未来有用的稳定知识，请在回复前执行 Knowledge Sync Gate，并说明应该更新哪个 module_state 或共享锚点。
- 锚点门控说明必须短于实际任务回答；没有变化时不要输出长流程说明。
```

通过标准：

- 读取范围与当前线程职责相关。
- 不做全仓库长审计。
- 能解决问题：给出相关文件、验证入口、风险边界。
- 结论区分 `Observed / Inference / Unverified`。
- 如果产生稳定知识，更新或建议更新本线程 Layer 3。
- 如果没有稳定知识，可以不展示收尾状态；若用户要求状态，只需一句说明没有持久更新。
- Anchor Gate / Knowledge Sync Gate 不应淹没实际任务内容。

阻塞失败：

- 忽略线程 scope，泛泛扫描全仓库。
- 未验证就给强结论。
- 产生稳定知识但没有 closeout 判断。
- 把未确认推断写成全局规则。

### 5. 测试跨线程 handoff

在同一个项目专家线程中输入：

```text
请判断这个请求是否完全属于你的职责范围：
“请修改你负责模块的实现，同时更新用户文档、测试策略，并决定是否需要新增长期 Documentation 线程。”

如果有跨线程部分，请只处理你的判断和 handoff，不要直接执行跨线程事项。
```

通过标准：

- 当前线程能识别自己负责的部分。
- 文档、测试策略、线程新增等跨边界事项被 handoff 或交给 Thread Management。
- 不直接修改不属于自己范围的文件。
- handoff 包含 source thread、target thread、confirmed facts、impact、questions、suggested next step。

### 6. 测试线程管理可维护

回到安装/线程管理对话，输入：

```text
请新增一个 Documentation 线程，用于维护用户文档、贡献者文档和文档类变更评审。
先给我变更提案，不要写入文件。
```

如果提案合理，回复：

```text
批准更新线程
```

通过标准：

- 未批准前不写入。
- 更新范围集中在 [contracts.md](#安装后锚点文件快速入口)、[thread_initialization.md](#安装后锚点文件快速入口)、相关 [module_state/](#安装后锚点文件快速入口) 或等价 Anchor PM 文件。
- 新增 Documentation 线程提示词可直接复制。
- 新增线程提示词包含 `Anchor Gate` 和 `Knowledge Sync Gate`。
- 不留下占位符。

### 7. 测试共享信息同步

新增 Documentation 线程后，新建一个原有项目专家线程，重新粘贴它的线程提示词，输入：

```text
请只基于当前锚点说明：现在项目中有哪些长期线程？Documentation 相关工作应该归谁处理？
```

通过标准：

- 能看到新增 Documentation 线程。
- 能正确说明 Documentation 相关事项应交给 Documentation 线程。
- 如果 detector/CLI 不可用，应表现为 degraded fallback。

阻塞失败：

- 新线程仍按旧线程列表回答。
- 声称程序化重锚完成但没有 detector/CLI。
- 共享信息更新后无法被其他线程恢复。

### 8. 测试 Knowledge Sync Gate 闭环

在任意项目专家线程中输入：

```text
请回顾本轮对话，判断是否有新知识或知识修正需要进入：
1. 本线程 Layer 3；
2. Layer 2 共享状态或 handoff；
3. Thread Management 的线程定义更新；
4. 框架级规则更新。

如果没有，只用一句话说明没有持久更新，不要展开流程。
```

通过标准：

- 能把本地知识、共享知识、线程定义、框架规则分开。
- 不把临时讨论写入长期状态。
- 需要写入时能指出目标文件或 handoff 对象。
- 与 `Anchor Gate` 形成闭环：工作前轻量读/确认，回复前只写持久变化。

### 9. 记录验证结果并还原

每个项目测试完成后填写记录，然后运行：

```bash
/mnt/g/data/threadsmith_restore_tests.sh
```

确认目标项目恢复 clean。

## 分项目验收标准

### Flask

- 预期推荐 3-4 个项目模块/子系统专家线程，例如 Runtime / App Lifecycle、Routing / Blueprints、Request-Response Surface、Docs / Tests Surface 这类可从项目结构解释的边界。
- 不应推荐 `Coordination` 作为普通目标项目线程；线程管理由当前安装对话承担。
- 不应把所有业务代码合并成一个宽泛的 `Implementation` 线程，除非明确说明 Flask 项目模块边界不足以拆分，并请求用户确认。
- 不应默认创建独立 `Validation` 线程；每个模块线程应负责自身相关测试/验证证据，除非项目存在独立验证/发布子系统。
- 不应把 Flask 的业务细节写成未确认规则；如果确实写入，必须明确标记为待确认。

### Vite

- 必须识别 monorepo / package 结构。
- 推荐线程应围绕 package/subsystem 边界，例如 Core Runtime、Plugin API、Create/Scaffold Packages、Docs / Developer Experience 等。
- 不应把所有 packages 归入单一 `Implementation` 线程。
- [thread_initialization.md](#安装后锚点文件快速入口) 不能留下需要用户手动填写的占位符。

### Django

- 必须能处理大型项目，但不能输出不可读的完整技术审计长日志。
- 安装提案应保持简短，同时说明为什么建议更多线程。
- 推荐线程应围绕 Django 的长期子系统边界，例如 ORM、HTTP / Views、Admin / Forms、Migrations、Docs / Tests Surface 等，而不是通用职责桶。
- 不应修改业务代码、测试、配置或原有文档正文。

## MVP 通过标准

只有 3 个项目全部满足下面条件，才视为 MVP 手动验收通过：

- 能从 public GitHub package source 完成安装，或在 GitHub 未更新时能从本地最新提示词完成同等安装。
- 明确批准前没有任何文件写入。
- 已有项目线程拆分来自真实模块/子系统，而不是通用职责桶。
- 线程名称语言与安装提示词语言一致。
- 安装后线程提示词可以直接复制使用，无用户占位符。
- 每个线程提示词包含 `Anchor Gate` 和 `Knowledge Sync Gate`。
- 提出并解决一个小问题时，线程能在 scope 内读取最小必要上下文。
- 锚点门控默认静默、最小读取、无变化不写入；流程说明不淹没实际任务。
- 跨边界事项能 handoff，不被一个线程全部吞掉。
- 新增线程后，共享锚点能被新线程恢复。
- Knowledge Sync Gate 能判断本地 Layer 3、Layer 2 shared state、Thread Management、框架级 handoff 的归属。
- [AGENTS.md](#安装后锚点文件快速入口) 已存在时不被静默覆盖。
- 还原脚本可以把测试项目恢复到 clean 状态。
- 每个项目都完成验证记录。

## MVP 阻塞失败

任一项目出现下面情况，视为 MVP blocker：

- 无法从 GitHub 获取 Threadsmith package。
- 安装提示词要求用户手动填写占位符。
- 未批准前写入文件。
- 修改业务代码、运行部署命令或运行迁移命令。
- 已有项目安装提案使用 `Coordination / Implementation / Validation` 代替项目模块拆分。
- 线程名称语言与安装提示词语言不一致。
- 默认提供 `调整 AGENTS.md` 作为回复选项。
- 用户可见安装提案大量解释内部安全约束、package 执行细节或 Anchor PM 内部机制。
- 默认提供“只安装 docs，不更新 AGENTS.md”之类会让集成不生效或让用户难以理解后果的选项。
- 线程没有 `Anchor Gate`。
- 线程没有 `Knowledge Sync Gate`。
- 锚点流程说明明显淹没实际任务回答。
- Knowledge Sync Gate 产生长期知识但没有写入、建议写入或 handoff。
- 大型项目安装提案变成难以评审的长篇审计报告。
- 线程无法维持职责边界，或共享锚点更新后无法被其他线程恢复。

## 验证记录模板

```text
项目:
路径:
日期:
安装提示词来源:
- GitHub / 本地

安装提案:
- 通过 / 需要重试 / 阻塞

线程拆分:
- 通过 / 需要重试 / 阻塞

线程提示词:
- 通过 / 需要重试 / 阻塞

Anchor Gate:
- 通过 / 需要重试 / 阻塞

问题解决:
- 通过 / 需要重试 / 阻塞

Handoff:
- 通过 / 需要重试 / 阻塞

Knowledge Sync Gate:
- 通过 / 需要重试 / 阻塞

共享信息同步:
- 通过 / 需要重试 / 阻塞

还原:
- 通过 / 需要重试 / 阻塞

Observed:
-

Inference:
-

Unverified:
-

Blocking:
-

Non-blocking:
-

结论:
- 通过 / 需要重试 / 阻塞
```
