---
title: v1.0 版本更新内容
tag: v1.0
summary: 2024年4月份产品版本更新内容
date: 2024-04-24
weight: 1
---

### 核心功能增强



#### 特性一：Issue 模板

![](https://cdn-static.gitcode.com/doc/v0.1-issue-template.png)

自定义 Issue 模板，一键创建标准化问题或功能请求，让你的开发工作更加井井有条。

- 支持多种模板类型: 支持传统 markdown 和 yaml 两种模版类型，让你的Issue模板功能更加灵活、高效。
- 可自定义字段: 根据项目需求，灵活定义 Issue 模板字段，确保收集必要的信息。

👉👉👉 [Issue 模板——告别繁琐的 Issue 输入](https://docs.gitcode.com/docs/repo/issues/template/)

#### 特性二：WebIDE

![](https://cdn-static.gitcode.com/doc/v0.1-webide.png)

无需本地环境，直接在浏览器中享受流畅的开发体验，随时随地编写、运行和调试代码。

- 支持多种编程语言: 支持主流编程语言，如 JavaScript、Python、Java 等，满足不同开发需求。
- 不同设备无缝切换: 只要有网络连接，WebIDE就能让你随时随地进入开发状态，享受编程的乐趣。

#### 特性三：在线解决冲突

![](https://cdn-static.gitcode.com/doc/v0.1-confict.png)

直观显示差异，并支持在线编辑，让你轻松化解代码合并难题。

- 可视化差异展示: 清晰显示代码冲突的位置和内容，方便对比和修改。
- 支持不同解决方式: 提供多种解决冲突方案，包括手动编辑、接受一方更改等。

👉👉👉 [在线解决冲突——轻松解决代码冲突](https://docs.gitcode.com/docs/pulls/conflict/)

#### 特性四：GPG Key

![](https://cdn-static.gitcode.com/doc/v0.1-commit-gpg.png)

GPG 签名验证功能，确保代码完整性和来源可靠，为你的代码安全保驾护航。

- 支持多种 GPG 密钥格式: 支持 RSA、DSA 等多种密钥格式，满足不同用户的需求。
- 可配置签名规则: 自定义签名规则，灵活控制签名行为。

👉👉👉 [GPG Key——代码安全再升级](https://docs.gitcode.com/docs/users/gpg/)

#### 特性五：全新帮助文档

全面更新的帮助文档，支持搜索功能，助你轻松解决问题，提升工作效率。

- 覆盖全功能: 涵盖 GitCode 的所有功能和用法，提供详细的操作指南和示例。
- 支持多种搜索方式: 支持关键词搜索、全文搜索等多种搜索方式，快速找到所需信息。

### 协作功能升级

#### 协作特性一：CodeOwners

![](https://cdn-static.gitcode.com/doc/v0.1-codeowners.png)

指定代码负责人，确保每次更改都能由合适的人审查，提升代码质量和项目管理效率。

- 支持多级 CodeOwners: 可以指定多个代码负责人，并设置审查优先级，确保代码审查的及时性和有效性。
- 自动添加 CodeOwners: 当代码发生修改时，系统会自动将相应的 CodeOwners 添加为审查者，无需手动操作。

👉👉👉 [CodeOwners——代码审查更清晰](https://docs.gitcode.com/docs/repo/code/codeowners/)

#### 协作特性二：PullRequest 增强

![](https://cdn-static.gitcode.com/doc/v0.1-pr-config.png)

支持标签、审查者分配、状态管理等功能，让代码审查流程更加高效顺畅。

- 支持标签标记: 可以使用标签标记 PullRequest 的类型、优先级等信息，方便分类管理。
- 支持行内代码评论: 审查者可以针对代码提出评论和建议，促进代码的完善。

### 重大特性升级

#### 重大特性一：组织级自定义权限

![](https://cdn-static.gitcode.com/doc/v0.1-org-role.png)

精细控制每位成员的访问和操作权限，让组织管理更加灵活、安全。

- 支持自定义角色: 最多支持 50 个自定义角色，满足组织内不同用户的权限需求。
- 支持细颗粒度权限控制: 可以细化到对仓库、分支、文件等不同层级的访问和操作权限控制。

👉👉👉 [组织级自定义权限——灵活管理团队权限](https://docs.gitcode.com/docs/orgs/permission/resource/)

#### 重大特性二：组织看板功能

![](https://cdn-static.gitcode.com/doc/v0.1-org-kanban.png)

直观的看板界面，帮助团队高效跟踪项目进度，优化工作流。

- 支持自定义字段: 可以根据项目需求，自定义看板的字段和状态，灵活管理项目任务。
- 支持分组视图: 按照任意字段进行分组，直观呈现不同状态下的任务分布情况，助你更快速地洞察项目全貌。

👉👉👉 [组织看板功能——可视化项目管理](https://docs.gitcode.com/docs/board/)

#### 重大特性三：OpenAPI


支持 OpenAPI 标准，让你的 API 更易于被其他开发者和服务使用，扩展项目影响力。

- 标准通用的 OpenAPI: 遵循 OpenAPI 标准，你的 API 可以被各类开发工具和平台轻松识别和解析，使用门槛大大降低
- 丰富的文档: GitCode 提供了详尽的 API 文档，涵盖所有 API 接口定义和使用方法，方便开发者快速上手。

👉👉👉 [OpenAPI——助力项目协作与扩展](https://docs.gitcode.com/docs/openapi/)

### 用户体验优化

#### 体验优化一：GitCode首页重塑

![](https://cdn-static.gitcode.com/doc/v0.1-homepage.png)

本次首页调整突出项目信息获取，让开发者更容易发现平台优质项目：

- G-Star项目，GitCode平台优质的开源项目精选，覆盖前端组件、云原生到大模型等重要领域
- 全球开源推荐，包括GitCode全球精选项目Trending 和六大领域开源项目频道
- 开源组织推荐，推荐GitCode优质的开源组织，便于开发者通过组织发现优质项目

👉👉👉 <https://gitcode.com>

#### 体验优化二：GitCode项目首页重塑

![](https://cdn-static.gitcode.com/doc/v0.1-repo-homepage.png)

本次项目首页调整，涉及到项目介绍与代码模块整合，增加项目代码目录旨在使页面布局更加直观和高效。功能亮点如下：

- 项目名称下的标签，更便于开发者快速查看开源协议、分支、Tags及提交记录
- 增加WebIDE入口，方便开发者直接通过IDE阅读和修改代码
- 项目下载量展示，除了传统star、fork数外可直观统计项目的使用量情况

👉👉👉 <https://gitcode.com/gitcode-offical-team/hugoblox>

#### 体验优化三：GitCode消息通知页重塑

![](https://cdn-static.gitcode.com/doc/v0.1-notification.png)

本次消息通知页改版旨在完善场景功能和优化页面结构形式

- 增加聚合全部消息的消息列表和未读列表，便于用户快速查阅和处理不同类型消息
- 增加未读标记功能，便于用户回溯消息，更改消息状体
- 调整整体信息结构，突出消息重点，便于开发者更好使用消息箱

#### 体验优化四：组件优化

![](https://cdn-static.gitcode.com/doc/v0.1-toolbar-search.png)

本次优化全站侧边栏和搜索组件从信息获取角度为开发者角度考虑，更方便查找自己的开源项目和组织

- 全站侧边栏，通过关注、星标、项目和组织等多维度展示出开发者工作台的维度
- 全站搜索组件，快速展示搜索记录和最近访问的项目，快速获取用户记录


[GitCode 新版本发布！4大核心功能增强！](https://mp.weixin.qq.com/s/KOTzANfWXsxTltw2b1yp4g)