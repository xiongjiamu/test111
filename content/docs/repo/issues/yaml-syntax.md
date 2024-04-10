---
linkTitle: YAML 表单语法
title: YAML 表单语法
weight: 5
---
GitCode 支持通过 YAML 语法来创建自定义的 Issue 模板，这使得项目维护者可以预定义需要提交者填写的字段。使用 YAML 构建模板可以让模板的创建和维护变得更加方便，同时也为贡献者提供了清晰的指引，确保他们能够提供所有必要的信息。希望这个简单的帮助文档能够助你一臂之力。

## 基本结构

一个基本的 GitCode Issue 模板通常包含以下部分：

* `name`: 模板的名称（必选）。
* `description`: 模板的描述，指导贡献者为什么和如何使用这个模板（必选）。
* `title`: 预设的 Issue 标题（可选）。
* `labels`: 自动添加到 Issue 上的标签（可选）。
* `assignees`: 自动分配的 GitCode 用户（可选）。
* `body`: 模板的主体，包含一系列的字段用于收集信息（可选）。

## YAML 模板示例

```yaml
name: Bug Report
description: 使用此模板来报告软件中的 bug。
title: "[BUG] Specific Issue Name"
labels: ["bug"]
assignees: 
  - username1
  - username2

body:
  - type: markdown
    attributes:
      value: |
        ## 描述 Bug
        请尽可能详细地描述你遇到的 bug。

  - type: input
    id: what-happened
    attributes:
      label: 发生了什么？
      description: 详细描述问题发生的情况。
      placeholder: 请输入详细信息
    validations:
      required: true

  - type: textarea
    id: reproduction
    attributes:
      label: 重现步骤
      description: 简单描述如何重现这个问题。
      placeholder: 1. 去 '...'
                  2. 点击 '....'
                  3. 向下滚动到 '....'
                  4. 看到错误
    validations:
      required: true

  - type: checkboxes
    id: version
    attributes:
      label: 受影响的版本
      description: 请勾选所有受影响的版本。
      options:
        - label: v1.0
          required: false
        - label: v2.0
          required: false

  - type: dropdown
    id: os
    attributes:
      label: 操作系统
      description: 你使用的操作系统。
      options:
        - label: Windows
        - label: macOS
        - label: Linux
        - label: Other
    validations:
      required: true
```

## 表单元素定义

| 键 | 说明                                                                                                                                        | 必选 | 类型   | 默认 | 有效值     |
| ---- | --------------------------------------------------------------------------------------------------------------------------------------------- | ------ | -------- | ------ | ------------ |
| `type`   | 元素类型。                                                                                                                                  | 必选 | String | -    | *`checkboxes`*<br />*`dropdown`*<br />`input`<br />`markdown`<br />`textarea`<br /> |
| `id`   | 元素的标识符，除非 `type` 设置为 `markdown`。 只能使用字母、数字、`-` 和 `_`。 在表单定义中必须是唯一的。 如果配置了id ，`id` 是 URL 查询参数预填中字段的规范标识符。 | 可选 | String | -    | -          |
| `attributes`   | 定义元素属性的一组键值对。                                                                                                                  | 必选 | 映射   | -    | -          |
| `validations`   | 设置元素约束的一组键值对。                                                                                                                  | 可选 | 映射   | -    | -          |

### 表单元素类型释义

你可以从以下类型的表单元素中选择一个类型。 每个类型都有唯一的属性和验证。：

* `markdown`: 用于添加静态文本，为用户提供额外的上下文，但并不会提交。
* `input`: 单行文本字段，适用于简短的文本输入。
* `textarea`: 多行文本字段，适合较长的描述或说明，提交人还可以在此字段中附加文件。
* `checkboxes`: 多选框，可以让用户选择多个选项。
* `dropdown`: 下拉菜单，让用户从多个选项中选择一个。

### markdown

用于添加静态文本，为用户提供额外的上下文，但并不会提交。

**attributes 定义**

| 键 | 说明                              | 必选 | 类型   | 默认 | 有效值 |
| ---- | ----------------------------------- | ------ | -------- | ------ | -------- |
| `value`   | 渲染的文本。 支持 Markdown 格式。 | 必选 | String | -    | -      |

> YAML 处理将哈希符号视为评论。 要插入 Markdown 标题，请用引号括住文本。对于多行文本，您可以使用竖线运算符。

以下是一个示例：

```yaml
body:
- type: markdown
  attributes:
    value: "## 感谢您对我们项目的反馈，这将使我们更加优秀"
- type: markdown
  attributes:
    value: |
      谢谢您抽出几分钟来给我们反馈我们的不足之处。
```

### input

单行文本字段，适用于简短的文本输入。

**attributes 定义**

| 键 | 说明                                           | 必选 | 类型   | 默认       | 有效值 |
| ---- | ------------------------------------------------ | ------ | -------- | ------------ | -------- |
| `label`   | 预期用户输入的简短描述，也以表单形式显示。     | 必选 | String | -          | -      |
| `description`   | 提供上下文或指导的字段的描述，以表单形式显示。 | 可选 | String | 空字符串<br /> | -      |
| `placeholder`   | 半透明的占位符，在字段空白时呈现。             | 可选 | String | 空字符串   | -      |
| `value`   | 字段中预填的文本。                             | 可选 | String | -          | -      |

**validations 定义**

| 键 | 说明                                            | 必选 | 类型 | 默认  | 有效值 |
| ---- | ------------------------------------------------- | ------ | ------ | ------- | -------- |
| `required`   | 防止在元素完成之前提交表单。 仅适用于公开项目。 | 可选 | 布尔 | false | -      |

以下是一个示例：

```yaml
body:
- type: input
  id: input
  attributes:
    label: Bug 发生频率
    description: "您大概多久遇到一次此类 bug?"
    placeholder: "例如：当我访问登录页面时大概两天一次"
  validations:
    required: true
```

### textarea

多行文本字段，适合较长的描述或说明，提交人还可以在此字段中附加文件。

**attributes 定义**

| 键 | 说明                                                                                                  | 必选 | 类型   | 默认       | 有效值                              |
| ---- | ------------------------------------------------------------------------------------------------------- | ------ | -------- | ------------ | ------------------------------------- |
| `label`   | 预期用户输入的简短描述，也以表单形式显示。                                                            | 必选 | String | -          | -                                   |
| `description`   | 提供上下文或指导的字段的描述，以表单形式显示。                                                        | 可选 | String | 空字符串<br /> | -                                   |
| `placeholder`   | 半透明的占位符，在字段空白时呈现。                                                                    | 可选 | String | 空字符串   | -                                   |
| `value`   | 字段中预填的文本。                                                                                    | 可选 | String | -          | -                                   |
| `render`   | 如果提供了值，提交的文本将格式化为代码块。 提供此键时，文本区域将不会扩展到文件附件或 Markdown 编辑。 | 可选 | String | -          | 已知的语言。 有关详细信息，请参阅[语言类型定义声明](https://github.com/github-linguist/linguist/blob/master/lib/linguist/languages.yml)。 |

**validations 定义**

| 键 | 说明                                            | 必选 | 类型 | 默认  | 有效值 |
| ---- | ------------------------------------------------- | ------ | ------ | ------- | -------- |
| `required`   | 防止在元素完成之前提交表单。 仅适用于公开项目。 | 可选 | 布尔 | false | -      |

以下是一个示例：

```yaml
body:
- type: textarea
  id: repro
  attributes:
    label: 复现步骤
    description: "你如何触发这个bug？请一步一步地的带着我们复现它。"
    value: |
      1.
      2.
      3.
      ...
    render: bash
  validations:
    required: true
```

### checkboxes

多选框，可以让用户选择多个选项。

**attributes 定义**

| 键 | 说明                                               | 必选 | 类型   | 默认       | 有效值 |
| ---- | ---------------------------------------------------- | ------ | -------- | ------------ | -------- |
| `label`   | 预期用户输入的简短描述，也以表单形式显示。         | 必选 | String | -          | -      |
| `description`   | 复选框的描述，以表单形式显示。 支持Markdown 格式。 | 可选 | String | 空字符串<br /> | -      |
| `options`   | 用户可以选择的复选框列表，语法参见下文。           | 必选 | Array  | -          | -      |

对于 `options` 数组中的每个值，可以设置以下键。

| 键 | 说明                                                                              | 必选 | 类型   | 默认  | 有效值 |
| ---- | ----------------------------------------------------------------------------------- | ------ | -------- | ------- | -------- |
| `label`   | 选项的标识符，显示在表单中。 支持 Markdown 用于粗体或斜体文本格式化和超文本链接。 | 必选 | String | -     | -      |
| `required`   | 防止在元素完成之前提交表单。 仅适用于公开项目。                                   | 可选 | 布尔   | false | -      |

**validations 定义**

| 键 | 说明                                            | 必选 | 类型 | 默认  | 有效值 |
| ---- | ------------------------------------------------- | ------ | ------ | ------- | -------- |
| `required`   | 防止在元素完成之前提交表单。 仅适用于公开项目。 | 可选 | 布尔 | false | -      |

以下是一个示例：

```yaml
body:
- type: checkboxes
  id: operating-systems
  attributes:
    label: 你所使用的操作系统是?
    description: 最少选择一个.
    options:
      - label: macOS
      - label: Windows
      - label: Linux
```

### dropdown

下拉菜单，让用户从多个选项中选择一个。

**attributes 定义**

| 键 | 说明                                                                          | 必选 | 类型       | 默认       | 有效值 |
| ---- | ------------------------------------------------------------------------------- | ------ | ------------ | ------------ | -------- |
| `label`   | 预期用户输入的简短描述，也以表单形式显示。                                    | 必选 | String     | -          | -      |
| `description`   | 提供上下文或指导的字段的描述，以表单形式显示。                                | 可选 | String     | 空字符串<br /> | -      |
| `multiple`<br /> | 确定用户是否可以选择多个选项。                                                | 可选 | 布尔       | false      | -      |
| `options`   | 用户可以选择的选项。 不能为空，所有选项必须是不同的。                         | 可选 | 字符串数组 | -          | -      |
| `default`   | `options` 数组中预选选项的索引。 指定了默认选项时，不能包含“None”或“n/a”作为选项。 | 可选 | Integer    | -          | -      |

**validations 定义**

| 键 | 说明                                            | 必选 | 类型 | 默认  | 有效值 |
| ---- | ------------------------------------------------- | ------ | ------ | ------- | -------- |
| `required`   | 防止在元素完成之前提交表单。 仅适用于公开项目。 | 可选 | 布尔 | false | -      |

以下是一个示例：

```yaml
body:
- type: dropdown
  id: download
  attributes:
    label: 你从哪里获知这个软件的?
    options:
      - GitCode
      - CSDN
      - Github
    default: 0
  validations:
    required: true
```