---
linkTitle: Tags
title: Tags
weight: 10
---

Tags 是一种用于标记仓库中特定点的发布版本号或里程碑的方式，Tags 允许你标记特定的 commit。

## Tags 的主要功能

- **发布版本控制**:Tags 可以用来标记软件的发布版本,如 v1.0,v2.0 等
- **里程碑标记**:可以为项目开发中的重要阶段打上标签,如 beta、release candidate 等
- **快速回滚**:如果发现已打标签的版本存在问题,可以快速回退到标签对应的 commit
- **分支控制**:为分支打标签是常用的分支管理策略

## 添加新的 Tag

{{% steps %}}

###
访问你的 GitCode 项目

###
在项目导航栏中，点击 "代码" 选项卡

###
单击导航栏下方的「Tags」图标，进入 Tags 列表页

###
单击页面右侧的「+新建Tag」按钮

###
在新建 Tag 弹窗中选择分支信息，eg: main、develop 等

###
输入 Tag 名称，通常为版本号如 v1.0

###
输入 Tag 的描述信息

###
点击「创建」创建新的 Tag

{{% /steps %}}

## 如何使用 Tags

- 在项目的 Tags 页面浏览所有 tag
- 通过选择标签查看对应点的代码快照
- 比较不同标签之间的区别
- 使用 git checkout + tag 名称检出标签对应的版本
- 使用 git diff 查看标签版本与当前代码的不同

## 最佳实践

- 使用语义化版本号作为 tag 名
- 保持 tag 说明信息的简洁明了
- 不要删除已发布的 tag
- 定期为关键节点添加 tag
- 推荐使用轻量级 tag

Tags 功能让版本控制变得简单易用，合理利用 Tags 可以更好地管理项目的发布和里程碑。