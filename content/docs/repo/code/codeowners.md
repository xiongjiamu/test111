---
linkTitle: Code Owners
title: Code Owners
weight: 12
---

## Code Owners简介

在GitCode网站上使用Code Owners功能可以帮助项目团队确保代码的关键部分由正确和合适的人员进行审查。Code Owners功能允许您为仓库中的特定文件或文件夹指定“所有者”。这些所有者会在相关文件或文件夹被修改并创建Pull Request时自动被请求审查代码。本文将引导您如何在GitCode中设置和使用Code Owners功能，以提升代码审查的效率和质量。

## 第一步：创建CODEOWNERS文件

要启用Code Owners功能，首先需要在您的项目仓库中创建一个名为`CODEOWNERS`​的文件。您可以将此文件放置在以下任一位置：

* 仓库的根目录
* ​`.gitcode/`​目录（如果存在）
* ​`docs/`​目录

‍

选择一个位置创建文件。GitHub会按照以下顺序检查这些位置：

* ​`CODEOWNERS`​ 在仓库的根目录。
* ​`.gitcode/CODEOWNERS`​
* ​`docs/CODEOWNERS`​

### 创建CODEOWNERS文件

1. 登录到您的GitCode账户并导航到相应的仓库。
2. 进入仓库后，点击“新建文件”按钮。
3. 将新文件命名为`CODEOWNERS`​，并选择保存位置。
4. 在文件编辑器中定义Code Owners规则（详见下方规则定义部分）。
5. 提交新文件。

## 第二步：定义Code Owners规则

在`CODEOWNERS`​文件中，您可以指定个人用户、组织或电子邮件地址作为特定文件或路径的Code Owners。每个规则一行，支持使用`#`​作为注释。下面是几个示例规则：

```
# 这是一个注释。
# 指定根目录下README文件的Code Owner
README.md @username

# 指定scripts文件夹的Code Owner
/scripts/ @username

# 指定所有JS文件的Code Owner
*.js @username

# 指定docs文件夹的多个Code Owners
/docs/ @username1 @username2

# 指定一个团队为某个目录的Code Owners
/docs/ @org/team-name

# 使用邮箱地址指定Code Owner
*.go username@example.com
```

### CODEOWNERS文件语法转义

警告：gitignore 文件有一些语法规则在 CODEOWNERS 文件中不起作用：

* 使用 `\`​ 对以 `#`​ 开头的模式进行转义，就可将其视为模式而不是注释
* 使用 `!`​ 否定模式
* 使用 `[ ]`​ 定义字符范围

CODEOWNERS 文件使用遵循 [gitignore](https://git-scm.com/docs/gitignore#_pattern_format) 文件中所用的大多数相同规则的模式。 模式后接一个或多个使用标准 `@username`​ 或 `@org/team-name`​ 格式的 GitCode用户名或组织名称。

如果要匹配两个或多个具有相同模式的代码所有者，所有的代码所有者必须位于同一行。 如果代码所有者不在同一行，模式会仅匹配最后提到的代码所有者。假设您的`CODEOWNERS`​文件中包含以下规则：

```
# 指定所有JS文件的Code Owners
*.js @frontendTeam

# 指定特定路径下所有文件的Code Owners
/src/scripts/ @javascriptExpert

# 项目根目录下的特定文件
/important.js @projectLead
```

如果有一个Pull Request修改了`/src/scripts/important.js`​文件，那么仅最后一行的规则才会被匹配，本次修改也将指派给@projectLead进行审查。

## 第三步：代码审查流程

当项目中的代码发生变化，并且这些变化涉及到Code Owners负责的文件时，相应的Code Owners会自动被请求审查相关的Pull Request。

### 自动请求审查

* GitCode会根据`CODEOWNERS`​文件中的规则自动请求Code Owners审查Pull Request。
* 如果有多个Code Owners，他们都会被添加到审查请求列表中。

## 管理CODEOWNERS文件

* 仅项目的维护者或具有管理权限的用户可以添加、修改或删除`CODEOWNERS`​文件。
* 强烈建议通过Pull Request来管理`CODEOWNERS`​文件的更改，以便进行代码审查。

## 常见问题解答

### Q: 如果Code Owners中的一个人审查通过，Pull Request就可以合并吗？

A: 这取决于您的项目设置。通常，至少需要一个Code Owner的批准才能合并。但具体规则可根据项目需求自定义设置。

### Q: 如果我们需要更改Code Owner，应该如何操作？

A: 通过编辑`CODEOWNERS`​文件，添加、删除或修改相应的规则即可。建议这些更改通过Pull Request完成，确保团队的其他成员都有机会进行审查。

‍
