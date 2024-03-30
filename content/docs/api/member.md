---
linkTitle: 成员 API
title: 组织、项目成员 API
weight: 5
sidebar:
  open: false
---

### 1. 获取成员列表

``GET /api/v4/projects/{project_id}/members/all``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| project_id | path | 字符串 | 项目ID | yes |
| query | query | 字符串 | 搜索关键字 | no |
| sort | query | 字符串 | 排序方式 | no |
| access_level | query | 字符串 | 返回指定access_level的成员 | no |
| join_way | query | 字符串 | 添加成员的方式 | no |
| with_department | query | 布尔值 | 是否返回部门信息 | no |
| current_user_at_top | query | 布尔值 | 当前用户处于列表第一个 | no |
| simple | query | 布尔值 | 仅返回`user.id`,`user.name`,`user.username`,`access_level`字段 | no |
| page | path | 整数 | 页码 | no |
| per_page | path | 整数 | 每页条数 | no |

##### 响应

```json
[
    {
        "id": 12863,
        "name": "public_083111",
        "name_cn": "0831test1",
        "username": "0831test1",
        "avatar_url": "https://www.111111.com",
        "email": "example@demo.com",
        "access_level": 50,
        "expires_at": null,
        "type": "GroupMember",
        "state": "active",
        "avatar_path": null,
        "custom_attributes": null,
        "web_url": "https://demoweb.com/public_0831111",
        "invite_email": null,
        "limited": false,
        "last_owner": null,
        "is_current_source_member": false,
        "last_source_owner": true,
        "join_way": "inherit",
        "domain_group": null,
        "inherit_from": "AuthTest",
        "inherit_from_path": "AuthTest",
        "source_name": "AuthTest",
        "member_roles": [
            {
                "access_level": 50,
                "expires_at": null,
                "type": "GroupMember",
                "source_name": "AuthTest",
                "group": {
                    "id": 18906,
                    "path": "AuthTest",
                    "name": "AuthTest",
                    "visibility": "internal",
                    "full_path": "AuthTest",
                    "full_name": "AuthTest"
                },
                "branches": null
            }
        ],
        "department": null,
        "committer_system_from": false
    }
]
```