---
linkTitle: Cherry Pick
title: Cherry Pick
weight: 10
---

Cherry-Pick 用于选择性地将一个或多个提交从一个分支应用到另一个分支，而不必合并整个分支的更改。这使你可以非常精确地选择性地引入代码更改，而不必合并整个分支。

### 如何使用 Cherry-Pick

使用 Cherry-Pick 非常简单：

{{% steps %}}

###
**打开 Cherry-Pick 操作**：在 GitCode 项目中，导航到你要 Cherry-Pick 的提交，单击提交标题打开提交详情页

###
**Cherry-Pick**：在提交详细信息页面的右上角，单击「Cherry-Pick」按钮

###
**创建 Cherry-Pick pull requests**：在弹出的确认框中，确认要执行 Cherry-Pick 操作

###
**审核和合并**：pull requests会进入审查流程，一旦获得批准，选择的提交将被应用到目标分支中

{{% /steps %}}

**注意：Cherry-Pick 也存在一些限制：**

- **可能会引入冲突**：如果选择的提交与目标分支中的其他更改发生冲突，你需要手动解决这些冲突
- **提交历史可能变得混乱**：频繁使用 Cherry-Pick 可能会导致提交历史变得复杂，需要谨慎使用

### 最佳实践

以下是使用pull requests中 Cherry-Pick 功能的最佳实践：

- **仔细选择提交**：在执行 Cherry-Pick 之前，请仔细选择要引入的提交，确保只引入必要的更改。
- **处理冲突**：如果 Cherry-Pick 操作引入冲突，及时处理这些冲突以确保代码的稳定性。
- **谨慎使用**：虽然 Cherry-Pick 是一个强大的工具，但请谨慎使用，以避免提交历史混乱和代码维护的问题。

Cherry-Pick 功能使你能够更灵活地控制分支之间的代码流，特别是在需要将特定更改引入到不同分支时。