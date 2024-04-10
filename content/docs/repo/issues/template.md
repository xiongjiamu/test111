---
linkTitle: Issue 模版
title: Issue 模板
weight: 4
---

创建一个有效的GitCode Issue模板可以帮助项目维护者更快地理解和解决问题。

### 什么是GitCode Issue模板？

GitCode Issue模板是预先填充的表单，用于引导用户提供有关问题的所有必要细节。这样做可以提高问题处理的效率，确保维护者获得解决问题所需的所有信息。当前GitCode Issue模板支持如下两种填写类型：

* [Markdown](#markdown配置模板)：传统的 Issue 模板，由若干 `.md` 文件组成。一般用户 Issue 的标题和正文的规范提示，对用户限制较弱。
* [表单YAML](#表单yaml配置模板)：你可以创建具有可自定义 Web 表单字段的Issue模板。 您可以通过在仓库中使用议题表单鼓励贡献者包含特定的结构化信息。 Issue模版使用  YAML 编写。 有关详细信息，请参阅[“YAML 表单语法”](../yaml-syntax)。 如果你不熟悉 YAML 并且想要了解详细信息，请参阅“[在五分钟内了解 YAML](https://learnxinyminutes.com/docs/yaml/)”。

### 如何创建GitCode Issue模板？

在你的项目的**默认分支**的根目录下，创建一个 `.gitcode/ISSUE_TEMPLATE/` 目录（兼容`.github`目录）。在该目录中，你可以创建一个或多个`.md`/`.yml`/`.yaml`文件。每个文件将作为一个独立的模板。

### Markdown配置模板

你可以在新建的 `.md`文件中配置 front-matter 信息，包括 Issue 模板的介绍、Issue 标题、指定负责人、指定 Label 等。同时在新建的 `.md` 文件中添加正文内容，该内容会作为用户新建 Issue 时预设的内容填充到描述中。

以下是一个简单的 Issue 模板示例，你可以根据你自己代码库的需要进行调整：

```markdown
---
name: "Bug 报告"
about: "报告一个问题帮助我们改进"
title: "【BUG】:"
labels: ["BUG"]
assignees: 'username'
---

### BUG 类型

<!-- 
请在这里描述你所遇到的 BUG 类型，以便我们更快定位问题，比如 UI、功能、体验等 
-->

### 复现步骤

<!-- 
请在这里描述你遇到该 BUG 时的页面及步骤
-->
```

上面这个模板表示的是一个用于【报告一个问题帮助我们改进】的【Bug 报告】Issue 模板，在使用该模板创建时，会创建一个标题默认以 （【BUG】:）开头、默认指派给【username】、Labels 是【BUG】的 Issue。

#### Front-matter 介绍

目前我们支持以下几种 markdown 的 front-matter 配置：

| 字段      | 说明                      | 备注                                                                        |
| ----------- | --------------------------- | ----------------------------------------------------------------------------- |
| name      | 模板名称                  | 含中文使用双引号                                                            |
| about     | 模板解释说明              | 含中文使用双引号                                                            |
| titile    | Issue 预设标题            | 含中文使用双引号                                                            |
| labels    | Issue 的 labels，支持多个 | 多个需要使用中括号，当含有不存在的 labels 时，在创建 Issue 时不显示该 label |
| assignees | Issue 默认指派人          | 指派人的 username                                                           |

### 表单YAML配置模板

Issue模板同时使用表单配置可以支持预设默认指派的用户和标签 (label)，同时支持自定义表单类型（输入/下拉/单选/多选/代码块等），并设置表单项是否必填。以下是 Issue 模板配置的示例：

```yaml
name: Bug 报告
description: 报告一个问题帮助我们改进
title: "[BUG] "
labels: ["bug"]
assignees:
  - username
body:
  - type: markdown
    attributes:
      value: |
        ## 谢谢您的报告！
        请在提交之前，检查该问题是否已经被报告过。
      
  - type: input
    id: what-happened
    attributes:
      label: 发生了什么？
      description: 请尽可能详细地描述问题。
      placeholder: 请在这里输入详细信息
    validations:
      required: true
  
  - type: checkboxes
    id: terms
    attributes:
      label: 确认项
      description: 请确认以下信息。
      options:
        - label: 我已经搜索过了现有的issue，确信这个问题是新的。
          required: true
        - label: 我已经阅读了说明文档。
          required: false
  
  - type: dropdown
    id: version
    attributes:
      label: 受影响的版本
      description: 请选择受影响的软件版本。
      options:
        - 1.0
        - 2.0
        - 3.0
        - 我不确定
    validations:
      required: true

# 更多字段和配置可以根据需要添加

```

具体字段释义如下：

| 字段        | 说明                            | 备注                                                                                                                                                    |
| ------------- | --------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name        | 模板名称，必填                  | 用于定义模版的名称                                                                                                                                      |
| description | 模板解释说明，必填              | 用于解释模版的作用                                                                                                                                      |
| titile      | Issue 预设标题，可选            | 含中文使用双引号                                                                                                                                        |
| labels      | Issue 的 labels，支持多个，可选 | 多个需要使用中括号，当含有不存在的 labels 时，在创建 Issue 时不显示该 label                                                                             |
| assignees   | Issue 默认指派人，可选          | 预设的被指派用户的 username，使用列表或使用逗号分隔。在 Issue 创建同时将 Issue 指派给具体用户。被指派用户必须属于项目成员，当被指派的用户不时将被忽略。 |
| body        | 必填                            | Issue 模板表单配置内容，具体可参考[“YAML 表单语法”](../yaml-syntax)                                                                                                 |

### 配置模版选择器（`config.yml`）

可以通过向 `.gitcode/ISSUE_TEMPLATE` 文件夹添加 `config.yml` 文件来自定义用户在项目中创建新问题时看到的模板选择器。

下面是`config.yml`文件的一个示例。

```yaml
# 该YAML文件用于配置issue模板选择器的选项
blank_issues_enabled: false  # 禁用创建空白issues的选项
contact_links:
  - name: 官方支持
    url: https://support.example.com
    about: 请通过我们的支持页面获取官方帮助。
  - name: 特性请求
    url: https://features.example.com
    about: 告诉我们你希望增加的新功能！
```

以下是对上述`config.yml`示例内容的解释：

* `blank_issues_enabled: false`：这会禁止用户创建不使用模板的空白问题。如果你希望用户也能创建不遵循模板格式的issues，可以将此值改为`true`。
* `contact_links`：这是一个列表，其中可以包含指向外部资源的链接，用户在创建issue之前可以访问这些资源以获得帮助或信息。

  * `name`：链接的名称，它将显示在issue创建页面上。
  * `url`：指向帮助页面或其他资源的链接地址。
  * `about`：简短描述，说明链接的内容或用途。

### 最佳实践

* **简洁明了：**  保持模板简单明了，避免过多不相关的要求。
* **鼓励详细报告：**  鼓励用户提供尽可能多的细节，但也说明哪些信息是必需的。
* **持续更新：**  根据用户反馈和项目需求调整和优化模板。

通过遵循这些指导原则，你可以创建清晰、高效的GitCode Issue模板，帮助加快问题诊断和解决过程。