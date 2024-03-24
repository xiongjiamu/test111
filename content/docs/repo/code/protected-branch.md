---
linkTitle: 保护分支
title: 保护分支
weight: 7
---

保护分支功用于确保代码的安全性、稳定性和一致性。它允许项目管理员和团队限制对特定分支的访问和更改，以防止不必要的错误合并和破坏代码库的稳定性。

保护分支通过对特定分支应用一系列保护措施，以确保代码库的安全性和一致性。这些措施包括：

- **强制代码评审**：要合并更改到受保护分支，必须经过代码pull requests评审并获得批准
- **限制合并权限**：只有特定的人或团队成员才能合并更改
- **禁止直接推送**：阻止直接向受保护分支提交更改，通常要求通过pull requests来提交更改
- **要求流水线检查通过**：确保提交的更改通过持续集成/持续交付（CI/CD）流程的检查

## 设置保护分支

设置保护分支的步骤如下：


{{% steps %}}

###
**进入项目设置**：进入你的 GitCode 项目，然后点击导航栏中的“项目设置”选项卡

###
**选择保护分支选项**：在项目设置页面的左侧导航栏中，选择“保护分支”选项

###
**选择分支**：在“保护分支规则”中，选择要保护的分支

###
**设置保护规则**：启用所需的保护规则，例如允许推送、允许合并以及是否允许强制推送等

###
**启用规则**：单击「确定」按钮以应用保护规则

{{% /steps %}}