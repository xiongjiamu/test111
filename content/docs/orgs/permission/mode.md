---
linkTitle: 项目权限模式
title: 项目权限模式
weight: 3
---

目前，GitCode 中组织的项目提供了两种权限模式：独立模式和继承模式。

- **独立模式**：拥有成员管理权限的用户可以独立邀请用户并赋予权限。
- **继承模式**：组织项目的所有成员将继承自组织，继承模式下不允许单独添加非组织成员。

> 注意：所有个人项目均使用独立模式，且只支持系统角色，不支持自定义角色。

## 项目权限模式

### 独立权限模式

在独立权限模式中，组织内的项目权限管理显示出较高的灵活性，允许拥有成员管理权限的用户自主邀请其他用户加入并分配相应权限。此模式适合于那些需要细致调控项目权限的环境，比如在团队内部，根据具体任务需求动态调整各成员的权限。

**适用场景**：

- **团队项目**：团队成员需紧密协作，同时可能需要对不同成员的权限进行区分，例如开发人员、测试人员和产品经理等。
- **项目外部合作**：当项目需要引入外部合作伙伴或临时参与者时，此模式允许邀请他们参与而不必成为组织的常规成员。

### 继承权限模式

在继承权限模式下，项目内的成员权限将直接从上级组织继承，无需单独为非组织成员设置权限。这一模式适合于那些项目权限需求较为固定且较少调整的环境，有助于简化权限管理流程，提升管理效率。

**适用场景**：

- **统一权限管理**：适用于组织内部需要统一权限设置的项目，无需对每个项目单独调整权限。
- **权限稳定性**：适用于项目成员角色和权限较为固定，不常需调整权限的项目。

## 项目默认权限设置

组织管理员可以在**组织设置** - **基础设置** 中修改组织下新建项目的默认权限设置，新建的项目将默认采用此设置。

## 更改项目权限设置

组织管理员可以在**项目设置** - **项目成员** 中更改项目的权限设置，支持在独立模式和继承模式之间切换。需要注意：

1. 从独立模式切换到继承模式时，将删除不属于组织的成员，并使项目成员与组织成员的权限保持一致。
2. 从继承模式切换到独立模式时，将仅保留组织创建者，并移除所有从组织继承而来的其他项目成员。

## 组织、项目成员管理

组织和项目的成员邀请和管理设置，详见[成员管理](/docs/repo/member)