---
title: v1.1 版本更新内容
tag: v1.1
summary: 2024年5月份产品版本更新内容
date: 2024-05-18
weight: 1
---

### 核心功能增强

#### 特性一：搜索功能升级

![](https://cdn-static.gitcode.com/doc/v1.1-search.png)

GitCode 搜索功能帮助用户在平台上快速找到所需信息，提高工作效率和开发体验。

- 搜索范围种类多：支持在项目、Issues、Pull Requests、用户、组织等多种内容中进行搜索，满足用户的不同需求
- 搜索结果更清晰：提供高亮显示和智能排序功能，使搜索结果更加清晰易读，帮助用户快速定位所需信息
- 支持组织内/项目内搜索：用户可以在特定组织或项目内进行搜索，方便团队协作和项目管理
- 支持搜索浏览记录：自动保存用户的搜索历史，方便用户随时查看和管理之前的搜索记录，提高工作效率


#### 特性二：权限矩阵调整

![](https://cdn-static.gitcode.com/doc/v1.1-permissions.gif)

GitCode 的权限矩阵调整功能帮助用户更灵活地管理团队成员的权限，确保项目安全和协作效率。

- 增加权限关联：支持不同资源点权限的相互依赖，自动勾选依赖权限点，避免权限配置错误
- 支持按角色筛选成员：用户可以根据角色筛选团队成员，快速查看和管理特定角色的权限分配情况，提升权限管理的便捷性

#### 特性三：Markdown 支持渲染 emoji 表情

![](https://cdn-static.gitcode.com/doc/v1.1-emoji.gif)

GitCode 的 Markdown 编辑器中支持渲染 emoji 表情功能，帮助用户在文档和评论中更生动地表达情感和强调重点。

- 多种 emoji 表情：提供丰富多样的 emoji 表情选择，满足用户在不同情境下的表达需求
- 简单易用：用户只需在 Markdown 文档中输入相应的 emoji 符号，即可轻松渲染出相应的表情，增强内容的可读性和趣味性
- 提升沟通效果：在 issue、pr 交流中使用 emoji 表情可以有效地传达情感，增进社区用户之间的沟通和理解，提升社区用户体验


#### 特性四：PR 支持关联里程碑

Pull Request（PR）支持关联里程碑功能，帮助用户更好地管理项目进度和版本发布计划。

- 清晰的进度追踪：通过将 PR 关联到里程碑，用户可以更清晰地追踪项目的进度和各个功能的实现情况，提升项目管理的透明度
- 优化版本发布：关联里程碑后，用户可以更方便地规划和管理版本发布流程，确保每个版本的功能和修复都按计划完成，提高发布效率和质量

### 其他特性升级

#### 特性升级一：CodeTree 支持空目录折叠

![](https://cdn-static.gitcode.com/doc/v1.1-empty.gif)

GitCode 在线浏览项目代码时，会自动检测并折叠空目录，使文件结构更加简洁明了，减少不必要的目录层级展示，帮助用户更高效地浏览代码。

- 提升浏览效率：用户可以更快速地定位到有内容的文件和目录，提升代码浏览和查找的效率，减少干扰
- 清晰的目录结构：折叠空目录后，代码的目录结构更加紧凑和清晰，便于用户理解和管理项目文件

#### 特性升级二：看板功能增强

![](https://cdn-static.gitcode.com/doc/v1.1-board.png)

GitCode 的看板功能增强，在原有表格视图基础上，增加了看板视图，为用户提供更灵活和高效的项目管理工具，提升团队协作效率。

- 支持看板模式：提供灵活的看板模式，用户可以根据任务状态、优先级或自定义分类来组织和展示任务，使项目管理更加直观
- 看板视图支持自定义排序：用户可以在看板视图中自定义任务的排序方式，按照需求排列任务顺序，确保重要任务得到优先处理
- issue 详情页支持查看看板信息：在 issue 详情页中，用户可以直接查看与该任务相关的看板信息，方便了解任务的当前状态和进展，提升项目管理的透明度和效率

#### 特性升级三：OpenAPI 支持获取仓库语言信息

OpenAPI 现已支持获取仓库语言信息的新功能，开发者能够轻松访问和分析项目中使用的编程语言，进一步优化资源配置和团队协作。我们的 API 持续进化，致力于为开发者提供更全面、更便捷的工具，帮助您的项目达到新的高度。

#### 特性升级四：WebIDE 支持 Vue、鸿蒙 eTS 语法高亮显示

![](https://cdn-static.gitcode.com/doc/v1.1-ets.png)

新增对 Vue 框架及鸿蒙 eTS 语法的高亮显示功能，为开发者提供更加直观和高效的编码体验。无论您是在构建前沿的 Web 应用还是开发鸿蒙 OS 相关项目，现在您的代码将更加清晰易读，错误更易发现，从而大大提升开发速度和准确性。

### 用户体验优化

#### 体验优化一：个人工作台焕新改版

![](https://cdn-static.gitcode.com/doc/v1.1-dashboard.png)

精心重塑六大板块信息：项目、动态、关注、组织、我的信息、最近参与，现以更直观的方式展示，让信息获取变得前所未有的轻松，确保您对于工作进展了如指掌。

#### 体验优化二：个人主页全新升级

![](https://cdn-static.gitcode.com/doc/v1.1-profile.png)

我们提供了一个可个性化定制专属于你的 readme，让您优雅地展示自己的成就。readme 模块不仅简化了寻找志同道合开发者的过程，还通过精选项目和生动的贡献图，使得他人对您的工作一目了然。现在让世界看到真正的您，并与全球开发者社区建立更紧密的联系吧！

#### 体验优化三：组织页焕新

![](https://cdn-static.gitcode.com/doc/v1.1-org.png)

一个专为展现团队独特风采而生的组织页，不仅能够吸引志同道合的人才了解您的团队，还能促进他们的加入。我们突出了热门项目区域，让团队的技术成就一目了然，彰显您的专业实力。同时新增的公告栏目确保团队最新动态能够迅速传达至整个社区。

#### 体验优化四：移动体验全面升级

![](https://cdn-static.gitcode.com/doc/v1.1-h5.jpg)

为了带给您更流畅的移动端体验，我们已经对50多个页面进行了精心的调整和优化。这些页面现在不仅响应更快，界面更加清晰，还能够无缝地适配各种屏幕尺寸，确保您在手机或平板上浏览时同样享受顺畅、一致的体验。我们的努力还在继续，更多页面的适配工作仍在火热进行中。

### Bug 修复

感谢社区热心用户 @Cody_G、@凌康、@哼了个哈 的反馈，我们修复了产品的相关 bug：

- 不支持查仓库目录历史的 bug 修复
- 文件最后一次提交记录显示错误的 bug 修复


[GitCode 517全新特性升级！超级搜索引擎与全功能协同！](https://mp.weixin.qq.com/s/Nex4nnsfhqVjpFr1_PonTg)