---
linkTitle: 仓库镜像
title: 仓库镜像
weight: 10
---

GitCode 仓库镜像功能允许用户通过设置远程仓库地址来创建仓库的镜像，从而简化了在不同平台之间进行代码同步、备份的操作。

在 GitCode 中，仓库镜像功能允许用户通过设置远程仓库地址，将一个仓库的内容完全复制到另一个仓库，包括代码、分支、标签等。这使得用户可以轻松地在不同的仓库之间同步代码，备份重要的项目。

需要注意的是：

1. 目前仅开放了 Pull 镜像功能，即从远程仓库拉取更新并同步到 GitCode 的项目仓库中
1. 将 GitCode 的项目代码推送到远程仓库的 Push 镜像功能暂未开放
1. 每个项目最多设置一个 Pull、一个 Push 镜像仓库地址
1. 当同时设置了 Pull、Push 镜像时，系统将优先执行 Pull 、然后执行 Push

## 如何设置 Pull 仓库镜像

你可以通过一下方式，为你的项目添加一个 Pull 仓库镜像

{{% steps %}}

### 
**登录 GitCode**：首先，请确保你已登录 GitCode 帐户

###
**进入项目**：导航到你要创建镜像的项目页面

###
**进入仓库设置**：在项目页面上，选择“项目设置”选项

###
**选择“仓库镜像”设置**：在项目设置页面上，选择“仓库镜像”选项，并在镜像操作中选择“Pull拉取”

###
**设置远程仓库地址**：在项目地址中输入要镜像的远程仓库的地址，如果是私有项目，则需要你提供远程仓库的地址以及有该远程仓库读权限的个人访问令牌

###
**保存设置**：点击“保存”按钮保存 Pull 镜像设置

{{% /steps %}}

## Pull 仓库镜像的工作原理

当你设置了 Pull 仓库镜像后，GitCode 将会每天自动检查远程仓库的变化，并将这些变化同步拉取到本地仓库。

此外，当项目启用了 Pull 仓库镜像功能后，为了避免双写对镜像造成的影响，系统将禁用仓库的写操作（包括新建分支、推送代码、新建 Tag等操作）。

### Pull 镜像的报错

Pull 镜像时，可能会遇到以下错误并导致镜像同步失败：

- `INVALID_ARGUMENT`，Pull 镜像地址存在无效参数，遇到此类错误时：请检查或更换新的项目地址
- `INTERNAL`，系统内部错误，在镜像 Pull 镜像时遇到了未知的系统错误，遇到此类错误时无需任何操作、敬请等待错误修复
- `UNRECOGNIZED`，无法识别的 Pull 镜像，遇到此类错误时无需任何操作，系统将在下一次同步时自动尝试
- `OUT_OF_RANGE`，当前同步失败次数过多，超出限制，遇到此类错误时无需任何操作，系统将在下一次同步时自动尝试

## 最佳实践

以下是使用 GitCode 仓库镜像功能的最佳实践：

- 确保选择一个稳定、可靠的远程仓库作为镜像源，以确保代码同步的可靠性
- 建议定期检查镜像设置，以确保及时获取最新的代码变更
- 确保你有合适的权限来设置和管理仓库的镜像，避免不必要的数据泄露和安全风险