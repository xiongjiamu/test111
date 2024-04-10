---
linkTitle: 迁移项目
title: 迁移项目
weight: 3
---

GitCode 提供了通过个人访问密钥（Personal Access Token）迁移 GitHub、Gitee 项目的功能
GitCode 提供了方便的功能，允许用户通过 GitHub、Gitee 的个人访问令牌直接导入项目到 GitCode，从而简化了项目迁移和协作的过程。本文档将向你介绍什么是 GitCode 的迁移项目功能以及如何使用。

GitCode 迁移项目允许你直接从 GitHub、Gitee 导入项目，无需手动克隆和推送项目，只需提供你的 GitHub、Gitee 个人访问令牌（Personal Access Token），并选择要导入的项目，GitCode 将自动处理剩下的步骤，包括克隆和推送项目仓库到 GitCode。

> 注：目前仅支持项目 git 仓库的迁移，项目 Wiki、Issue、PR 及 Release 暂不支持迁移和导入。

## 创建个人访问令牌

要使用 GitCode 迁移项目，你需要先准备 GitHub、Gitee 的个人访问令牌，以便 GitCode 可以访问你的项目仓库。详细的个人访问令牌生成步骤，你可以参考 GitHub、Gitee 的官方网站：

- [点击查看如何生成 GitHub 个人访问令牌 ](https://docs.github.com/zh/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens)
- [点击查看如何获取 Gitee 个人访问令牌](https://gitee.com/help/articles/4336#article-header13)

接下来，你可以在 GitCode 的界面上使用简单的步骤完成项目的迁移和导入。

## 迁移项目的步骤

{{% steps %}}

### 
**登录 GitCode**：首先，请确保你已登录到 GitCode 帐户

###
**选择迁移平台**：点击导航上的“迁移项目”按钮，并选择你要迁移的平台（GitHub 或 Gitee）

### 
**选择项目所属帐户**：选择要将存储库导入到的 GitCode 帐户/组织

###
**输入个人访问令牌**：在页面的对话框中，输入对应平台的个人访问令牌。确保你的令牌具有适当的权限以访问你要导入的项目，系统会自动记住上一次有效的个人访问令牌，你也可以选择更换新的令牌

###
**点击“迁移项目”按钮**：点击页面的“迁移项目按钮”

###
**选择要导入的项目**：在页面上搜索并选择要导入的项目（建议通过 namespace 快速筛选过滤你要导入的项目）；你也可以选择批量导入多个项目

### 
**配置导入选项**：根据你的需求配置导入选项，例如项目归属、项目路径等

### 
**点击“导入项目”**：点击“导入”按钮或“批量导入”开始导入项目，系统将自动开始克隆和推送项目仓库到 GitCode

### 
**等待导入完成**：导入过程可能需要一些时间，具体取决于项目仓库的大小和网络速度。请耐心等待，直到导入完成

{{% /steps %}}

## 最佳实践

以下是使用 GitCode 迁移项目功能的最佳实践：

- 谨慎保管个人访问令牌，请确保不要与他人分享，并妥善保管
- 导入完成后，请检查项目仓库是否成功导入到 GitCode
