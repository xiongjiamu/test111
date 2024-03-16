---
linkTitle: Issue API
title: Issue接口API文档
weight: 6
sidebar:
  open: false
---

### 1. 获取个人issue列表

``GET /api/v4/issues``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| state | path | 字符串 | 返回指定状态的Issue，`opened`（开启中）、`closed`（已关闭） 或 `all`（全部） | no |
| milestone | path | 字符串 | 指定关联的里程碑 | no |
| order_by | path | 字符串 | 以`created_at` 或 `updated_at`字段排序 | no |
| sort | path | 字符串 | `order by`字段排序规则 `asc` （升序）或 `desc`（降序） | no |
| search | path | 字符串 | 搜索关键字，按照标题或描述 | no |
| in | path | 字符串 | 以逗号分隔的`标题`, `描述`, 或 body | no |
| created_after | path | 字符串 | 返回给定时间以前创建的Issue | no |
| created_before | path | 字符串 | 返回给定时间以后创建的Issue | no |
| updated_after | path | 字符串 | 返回给定时间以前更新的Issue | no |
| updated_before | path | 字符串 | 返回给定时间以后更新的Issue | no |
| author_id | path | 字符串 | Issue创建者ID | no |
| assignee_id | path | 字符串 | Issue指派人ID | no |
| scope | path | 字符串 | 指定返回的Issue类型 `created_by_me`（我创建的）, `assigned_to_me`（指派给我的） or `all`（全部） | no |
| page | path | 整数 | 页码 | no |
| per_page | path | 整数 | 每页条数 | no |
| labels | path | 字符串 | 返回指定标签关联的Issue | no |
| only_count | path | 布尔值 | 取值`true`则仅返回Issue的统计数量 | no |

##### 响应

```json
{
  "page_num": 1,
  "page_size": 10,
  "total": 1,
  "page_count": 1,
  "content": [
    {
      "id": 1,
      "iid": 1,
      "project_id": 123456,
      "title": "test_issue",
      "description": "description",
      "state": "opened",
      "created_at": "2023-09-18T14:41:46.373+08:00",
      "updated_at": "2023-09-18T15:09:33.265+08:00",
      "closed_at": null,
      "closed_by": null,
      "labels": [
        {
          "id": 11026,
          "title": "bug",
          "name": "bug",
          "color": "#ED4014",
          "type": "ProjectLabel",
          "text_color": "#FFFFFF",
          "is_expired": false
        }
      ],
      "milestone": {
        "id": 1,
        "iid": 1,
        "project_id": 123456,
        "title": "test_milestone",
        "description": "description",
        "state": "active",
        "created_at": "2023-09-18T14:59:10.643+08:00",
        "updated_at": "2023-09-18T15:01:29.427+08:00",
        "due_date": "2024-01-01",
        "start_date": null,
        "project_path": "test_user/test_project",
        "web_url": "https://gitcode.com/test_user/test_project/milestones/1"
      },
      "assignees": [],
      "author": {
        "id": 173,
        "name": "test_user",
        "username": "test_user",
        "iam_id": "195f35c15e69497c8be72d15cb2199a8",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "test_user@noreply.gitcode.com",
        "name_cn": "test_user",
        "web_url": "https://gitcode.com/test_user",
        "nick_name": "",
        "tenant_name": null
      },
      "assignee": null,
      "user_notes_count": 0,
      "merge_requests_count": 0,
      "due_date": null,
      "confidential": false,
      "discussion_locked": null,
      "web_url": "https://gitcode.com/test_user/test_project/issues/1",
      "project_path_with_namespace": "test_user/test_project",
      "project": {
        "id": 590885,
        "description": "",
        "name": "dingzk_test_project",
        "name_with_namespace": "test_user / test_project",
        "path": "test_project",
        "path_with_namespace": "test_user/test_project",
        "develop_mode": "normal",
        "created_at": "2023-09-18T14:40:25.084+08:00",
        "updated_at": "2023-09-18T14:40:25.084+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "git@gitcode.com:test_user/test_project.git",
        "http_url_to_repo": "https://gitcode.com/test_user/test_project.git",
        "web_url": "https://gitcode.com/test_user/test_project",
        "readme_url": null,
        "product_id": null,
        "product_name": null,
        "license_url": null,
        "license": null
      }
    }
  ]
}
```
### 2. 创建项目下的Issue

``POST /api/v4/projects/{project_id}/issues``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | path | 字符串 | 项目ID | yes |
| title | body | 字符串 | 标题 | yes |
| proposer_id | body | 整数 | 提出人 | no |
| discussion_locked | body | 布尔值 | 锁定Issue评论 | no |
| iid | body | 整数 | 内部ID | no |
| labels | body | 数组 | 标签 | no |
| milestone_id | body | 整数 | 里程碑ID | no |
| due_date | body | 字符串 | 截止日期 | no |
| created_at | body | 字符串 | 创建时间 | no |
| description | body | 字符串 | 描述 | no |
| discussion_to_resolve | body | 字符串 | 待解决的评论 | no |
| merge_request_to_resolve_discussions_of | body | 整数 | 转为Issue的检视意见所在的合并请求 | no |
| assignee_id | body | 整数 | 处理人 | no |
| confidential | body | 布尔值 | 机密 | no | 
##### 响应

```json
{
  "id": 1,
  "iid": 1,
  "project_id": 123456,
  "title": "test_issue",
  "description": "description",
  "state": "opened",
  "created_at": "2023-09-18T14:41:46.373+08:00",
  "updated_at": "2023-09-18T14:41:46.373+08:00",
  "closed_at": null,
  "closed_by": null,
  "labels": [
    {
      "id": 1,
      "title": "bug",
      "name": "bug",
      "color": "#ED4014",
      "type": "ProjectLabel",
      "text_color": "#FFFFFF",
      "is_expired": false
    }
  ],
  "milestone": {
    "id": 1,
    "iid": 1,
    "project_id": 123456,
    "title": "test_milestone",
    "description": "description",
    "state": "active",
    "created_at": "2023-09-18T14:59:10.643+08:00",
    "updated_at": "2023-09-18T15:01:29.427+08:00",
    "due_date": "2024-01-01",
    "start_date": null,
    "project_path": "test_user/test_project",
    "web_url": "https://gitcode.com/test_user/test_project/milestones/1"
  },
  "assignees": [],
  "author": {
    "id": 1,
    "name": "test_user",
    "username": "test_user",
    "iam_id": "abcdefghijklmnopqrstuvwxyz123456",
    "state": "active",
    "avatar_url": null,
    "avatar_path": null,
    "email": "test_user@noreply.gitcode.com",
    "name_cn": "test_user",
    "web_url": "https://gitcode.com/test_user",
    "nick_name": "",
    "tenant_name": null
  },
  "assignee": null,
  "user_notes_count": 0,
  "merge_requests_count": 0,
  "due_date": null,
  "confidential": false,
  "discussion_locked": null,
  "web_url": "https://gitcode.com/test_user/test_project/issues/1",
  "project_path_with_namespace": "test_user/test_project",
  "project": {
    "id": 123456,
    "description": "",
    "name": "dingzk_test_project",
    "name_with_namespace": "test_user / test_project",
    "path": "test_project",
    "path_with_namespace": "test_user/test_project",
    "develop_mode": "normal",
    "created_at": "2023-09-18T14:40:25.084+08:00",
    "updated_at": "2023-09-18T14:40:25.084+08:00",
    "archived": false,
    "is_kia": false,
    "ssh_url_to_repo": "git@gitcode.com:test_user/test_project.git",
    "http_url_to_repo": "https://gitcode.com/test_user/test_project.git",
    "web_url": "https://gitcode.com/test_user/test_project",
    "readme_url": null,
    "product_id": null,
    "product_name": null,
    "license_url": null,
    "license": null
  },
  "subscribed": false,
  "related_merge_request": null,
  "linked_merge_requests": []
}
```
### 3. 获取项目下的Issue列表

``GET /api/v4/projects/{project_id}/issues``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | path | 字符串 | 项目ID | yes |
| state | path | 字符串 | 返回指定状态的Issue，`opened`（开启中）、`closed`（已关闭） 或 `all`（全部） | no |
| milestone | path | 字符串 | 指定关联的里程碑 | no |
| order_by | path | 字符串 | `order by`字段排序规则 `asc` （升序）或 `desc`（降序） | no |
| sort | path | 字符串 | `order by`字段排序规则 `asc` （升序）或 `desc`（降序） | no |
| search | path | 字符串 | 搜索关键字，按照标题或描述 | no |
| in | path | 字符串 | 以逗号分隔的标题, 描述, 或body | no |
| created_after | path | 字符串 | 返回给定时间以前创建的Issue | no |
| created_before | path | 字符串 | 返回给定时间以后创建的Issue | no |
| updated_after | path | 字符串 | 返回给定时间以前更新的Issue | no |
| updated_before | path | 字符串 | 返回给定时间以后更新的Issue | no |
| author_id | path | 字符串 | Issue创建者ID | no |
| assignee_id | path | 字符串 | Issue指派人ID | no |
| scope | path | 字符串 | 指定返回的Issue类型 `created_by_me`（我创建的）, `assigned_to_me`（指派给我的） 或者 `all`（全部） | no |
| page | path | 整数 | 页码 | no |
| per_page | path | 整数 | 每页条数 | no |
| labels | path | 字符串 | 返回指定标签关联的Issue | no |

##### 响应

```json
{
  "opened": 1,
  "closed": 0,
  "all": 1,
  "issues": [
    {
      "id": 1,
      "iid": 1,
      "project_id": 123456,
      "title": "test_issue",
      "description": null,
      "state": "opened",
      "created_at": "2023-09-18T14:41:46.373+08:00",
      "updated_at": "2023-09-18T14:41:46.373+08:00",
      "closed_at": null,
      "closed_by": null,
      "labels": [
        {
          "id": 1,
          "title": "bug",
          "name": "bug",
          "color": "#ED4014",
          "type": "ProjectLabel",
          "text_color": "#FFFFFF",
          "is_expired": false
        }
      ],
      "milestone": {
        "id": 1,
        "iid": 1,
        "project_id": 123456,
        "title": "test_milestone",
        "description": "description",
        "state": "active",
        "created_at": "2023-09-18T14:59:10.643+08:00",
        "updated_at": "2023-09-18T15:01:29.427+08:00",
        "due_date": "2024-01-01",
        "start_date": null,
        "project_path": "test_user/test_project",
        "web_url": "https://gitcode.com/test_user/test_project/milestones/1"
      },
      "assignees": [],
      "author": {
        "id": 1,
        "name": "test_user",
        "username": "test_user",
        "iam_id": "abcdefghijklmnopqrstuvwxyz123456",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "test_user@noreply.gitcode.com",
        "name_cn": "test_user",
        "web_url": "https://gitcode.com/test_user",
        "nick_name": "",
        "tenant_name": null
      },
      "assignee": null,
      "user_notes_count": 0,
      "merge_requests_count": 0,
      "due_date": null,
      "confidential": false,
      "discussion_locked": null,
      "web_url": "https://gitcode.com/test_user/test_project/issues/1",
      "project_path_with_namespace": "test_user/test_project",
      "project": {
        "id": 123456,
        "description": "",
        "name": "test_project",
        "name_with_namespace": "test_user / test_project",
        "path": "test_project",
        "path_with_namespace": "test_user/test_project",
        "develop_mode": "normal",
        "created_at": "2023-09-18T14:40:25.084+08:00",
        "updated_at": "2023-09-18T14:40:25.084+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "git@gitcode.com:test_user/test_project.git",
        "http_url_to_repo": "https://gitcode.com/test_user/test_project.git",
        "web_url": "https://gitcode.com/test_user/test_project",
        "readme_url": null,
        "product_id": null,
        "product_name": null,
        "license_url": null,
        "license": null
      }
    }
  ]
}
```
### 4. 获取项目Issue详情

``GET /api/v4/projects/{project_id}/issues/{issue_iid}``

##### 参数

| 参数名 | 位置 | 描述 | 类型 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id |项目ID | path | 字符串 | yes |
| issue_iid | Issue内部ID |  path | 整数 | yes |

##### 响应

```json
{
  "id": 1,
  "iid": 1,
  "project_id": 123456,
  "title": "test_issue",
  "description": "description",
  "state": "opened",
  "created_at": "2023-09-18T14:41:46.373+08:00",
  "updated_at": "2023-09-18T14:41:46.373+08:00",
  "closed_at": null,
  "closed_by": null,
  "labels": [
    {
      "id": 1,
      "title": "bug",
      "name": "bug",
      "color": "#ED4014",
      "type": "ProjectLabel",
      "text_color": "#FFFFFF",
      "is_expired": false
    }
  ],
  "milestone": {
    "id": 1,
    "iid": 1,
    "project_id": 123456,
    "title": "test_milestone",
    "description": "description",
    "state": "active",
    "created_at": "2023-09-18T14:59:10.643+08:00",
    "updated_at": "2023-09-18T15:01:29.427+08:00",
    "due_date": "2024-01-01",
    "start_date": null,
    "project_path": "test_user/test_project",
    "web_url": "https://gitcode.com/test_user/test_project/milestones/1"
  },
  "assignees": [],
  "author": {
    "id": 1,
    "name": "test_user",
    "username": "test_user",
    "iam_id": "abcdefghijklmnopqrstuvwxyz123456",
    "state": "active",
    "avatar_url": null,
    "avatar_path": null,
    "email": "test_user@noreply.gitcode.com",
    "name_cn": "test_user",
    "web_url": "https://gitcode.com/test_user",
    "nick_name": "",
    "tenant_name": null
  },
  "assignee": null,
  "user_notes_count": 0,
  "merge_requests_count": 0,
  "due_date": null,
  "confidential": false,
  "discussion_locked": null,
  "web_url": "https://gitcode.com/test_user/test_project/issues/1",
  "project_path_with_namespace": "test_user/test_project",
  "project": {
    "id": 590885,
    "description": "",
    "name": "dingzk_test_project",
    "name_with_namespace": "test_user / test_project",
    "path": "test_project",
    "path_with_namespace": "test_user/test_project",
    "develop_mode": "normal",
    "created_at": "2023-09-18T14:40:25.084+08:00",
    "updated_at": "2023-09-18T14:40:25.084+08:00",
    "archived": false,
    "is_kia": false,
    "ssh_url_to_repo": "git@gitcode.com:test_user/test_project.git",
    "http_url_to_repo": "https://gitcode.com/test_user/test_project.git",
    "web_url": "https://gitcode.com/test_user/test_project",
    "readme_url": null,
    "product_id": null,
    "product_name": null,
    "license_url": null,
    "license": null
  },
  "links": null,
  "subscribed": false,
  "related_merge_request": null,
  "linked_merge_requests": [],
  "related_note_params": null
}
```
### 5. 更新项目下的Issue

``PUT /api/v4/projects/{project_id}/issues/{issue_iid}``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | path | 字符串 | 项目ID | yes |
| issue_iid |  body | 整数 | Issue内部ID | yes |
| title | body | 字符串 | 标题 | yes |  
| proposer_id | body | 整数 | 提出人 | no |
| discussion_locked | body | 布尔值 | 锁定评论 | no |
| milestone_id | body | 整数 | 里程碑ID | no |
| due_date | body | 字符串 | 截止日期 | no |
| description | body | 字符串 | 描述 | no |
| created_at | body | 字符串 | 创建时间 | no |
| labels | body | 数组 | 标签 | no |
| updated_at | body | 字符串 | 更新日期 | no |
| state_event | body | 字符串 | 状态（开启/关闭） | no |
| assignee_id | body | 整数 | 处理人ID | no |
| confidential | body | 布尔值 | 机密 | no |

##### 响应

```json
{
  "id": 1,
  "iid": 1,
  "project_id": 123456,
  "title": "test_issue",
  "description": "description",
  "state": "opened",
  "created_at": "2023-09-18T14:41:46.373+08:00",
  "updated_at": "2023-09-18T15:09:33.265+08:00",
  "closed_at": null,
  "closed_by": null,
  "labels": [
    {
      "id": 11026,
      "title": "bug",
      "name": "bug",
      "color": "#ED4014",
      "type": "ProjectLabel",
      "text_color": "#FFFFFF",
      "is_expired": false
    }
  ],
  "milestone": {
    "id": 1,
    "iid": 1,
    "project_id": 123456,
    "title": "test_milestone",
    "description": "description",
    "state": "active",
    "created_at": "2023-09-18T14:59:10.643+08:00",
    "updated_at": "2023-09-18T15:01:29.427+08:00",
    "due_date": "2024-01-01",
    "start_date": null,
    "project_path": "test_user/test_project",
    "web_url": "https://gitcode.com/test_user/test_project/milestones/1"
  },
  "assignees": [],
  "author": {
    "id": 1,
    "name": "test_user",
    "username": "test_user",
    "iam_id": "abcdefghijklmnopqrstuvwxyz123456",
    "state": "active",
    "avatar_url": null,
    "avatar_path": null,
    "email": "test_user@noreply.gitcode.com",
    "name_cn": "test_user",
    "web_url": "https://gitcode.com/test_user",
    "nick_name": "",
    "tenant_name": null
  },
  "assignee": null,
  "user_notes_count": 0,
  "merge_requests_count": 0,
  "due_date": null,
  "confidential": false,
  "discussion_locked": null,
  "web_url": "https://gitcode.com/test_user/test_project/issues/1",
  "project_path_with_namespace": "test_user/test_project",
  "project": {
    "id": 590885,
    "description": "",
    "name": "dingzk_test_project",
    "name_with_namespace": "test_user / test_project",
    "path": "test_project",
    "path_with_namespace": "test_user/test_project",
    "develop_mode": "normal",
    "created_at": "2023-09-18T14:40:25.084+08:00",
    "updated_at": "2023-09-18T14:40:25.084+08:00",
    "archived": false,
    "is_kia": false,
    "ssh_url_to_repo": "git@gitcode.com:test_user/test_project.git",
    "http_url_to_repo": "https://gitcode.com/test_user/test_project.git",
    "web_url": "https://gitcode.com/test_user/test_project",
    "readme_url": null,
    "product_id": null,
    "product_name": null,
    "license_url": null,
    "license": null
  },
  "links": null,
  "subscribed": false,
  "related_merge_request": null,
  "linked_merge_requests": []
}
```

### 6. 获取组织下的Issue列表

#### GET /api/v4/groups/{group_id}/issues

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| group_id | path | 字符串 | 组织ID | yes |
| state | path | 字符串 | 返回指定状态的Issue，`opened`（开启中）、`closed`（已关闭） 或 `all`（全部） | no |
| milestone | path | 字符串 | 指定关联的里程碑 | no |
| order_by | path | 字符串 | `order by`字段排序规则 `asc` （升序）或 `desc`（降序） | no |
| sort | path | 字符串 | `order by`字段排序规则 `asc` （升序）或 `desc`（降序） | no |
| search | path | 字符串 | 搜索关键字，按照标题或描述 | no |
| in | path | 字符串 | 以逗号分隔的标题, 描述, 或body | no |
| created_after | path | 字符串 | 返回给定时间以前创建的Issue | no |
| created_before | path | 字符串 | 返回给定时间以后创建的Issue | no |
| updated_after | path | 字符串 | 返回给定时间以前更新的Issue | no |
| updated_before | path | 字符串 | 返回给定时间以后更新的Issue | no |
| author_id | path | 字符串 | Issue创建者ID | no |
| assignee_id | path | 字符串 | Issue指派人ID | no |
| scope | path | 字符串 | 指定返回的Issue类型 `created_by_me`（我创建的）, `assigned_to_me`（指派给我的） or `all`（全部） | no |
| page | path | 整数 | 页码 | no |
| per_page | path | 整数 | 每页条数 | no |
| labels | path| 字符串 | 返回指定标签关联的Issue | no |

##### 响应

```json
{
  "page_num": 1,
  "page_size": 10,
  "total": 1,
  "page_count": 1,
  "content": {
    "opened": 1,
    "closed": 0,
    "all": 1,
    "issues": [
      {
        "id": 1,
        "iid": 1,
        "project_id": 123456,
        "title": "test_issue",
        "description": "description",
        "state": "opened",
        "created_at": "2023-09-18T14:41:46.373+08:00",
        "updated_at": "2023-09-18T15:09:33.265+08:00",
        "closed_at": null,
        "closed_by": null,
        "labels": [
          {
            "id": 1,
            "title": "bug",
            "name": "bug",
            "color": "#ED4014",
            "type": "ProjectLabel",
            "text_color": "#FFFFFF",
            "is_expired": false
          }
        ],
        "milestone": {
          "id": 1,
          "iid": 1,
          "project_id": 123456,
          "title": "test_milestone",
          "description": "description",
          "state": "active",
          "created_at": "2023-09-18T14:59:10.643+08:00",
          "updated_at": "2023-09-18T15:01:29.427+08:00",
          "due_date": "2024-01-01",
          "start_date": null,
          "project_path": "test_group/test_project",
          "web_url": "https://gitcode.com/test_group/test_project/milestones/1"
        },
        "assignees": [],
        "author": {
          "id": 1,
          "name": "test_user",
          "username": "test_user",
          "iam_id": "abcdefghijklmnopqrstuvwxyz123456",
          "state": "active",
          "avatar_url": null,
          "avatar_path": null,
          "email": "test_user@noreply.gitcode.com",
          "name_cn": "test_user",
          "web_url": "https://gitcode.com/test_user",
          "nick_name": "",
          "tenant_name": null
        },
        "assignee": null,
        "user_notes_count": 0,
        "merge_requests_count": 0,
        "due_date": null,
        "confidential": false,
        "discussion_locked": null,
        "web_url": "https://gitcode.com/test_group/test_project/issues/1",
        "project_path_with_namespace": "test_group/test_project",
        "project": {
          "id": 590885,
          "description": "",
          "name": "dingzk_test_project",
          "name_with_namespace": "test_group / test_project",
          "path": "test_project",
          "path_with_namespace": "test_group/test_project",
          "develop_mode": "normal",
          "created_at": "2023-09-18T14:40:25.084+08:00",
          "updated_at": "2023-09-18T14:40:25.084+08:00",
          "archived": false,
          "is_kia": false,
          "ssh_url_to_repo": "git@gitcode.com:test_group/test_project.git",
          "http_url_to_repo": "https://gitcode.com/test_group/test_project.git",
          "web_url": "https://gitcode.com/test_group/test_project",
          "readme_url": null,
          "product_id": null,
          "product_name": null,
          "license_url": null,
          "license": null
        }
      }
    ]
  }
}
```

### 7. 获取项目下的标签列表

``GET /api/v4/projects/{project_id}/labels``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | path | 字符串 | 项目ID | yes |
| sort | path | 字符串 | `order by`字段排序规则 `asc` （升序）或 `desc`（降序） | no |
| search | path | 字符串 | 搜索关键字，按照标题或描述 | no |
| include_expired | path | 布尔值 | 取值`true` 则返回包括以超期的Label | no |
| most_active_ones | path | 布尔值 | 取值`true`, 则返回最常使用的Label | no |
| view | path | 字符串 | 取值`simple`、返回简要信息, 取值`basic` 、返回基础信息, 取值`detail`、返回详细信息 | no |
| page | path | 整数 | 页码 | no |
| per_page | path | 整数 | 每页条数 | no |

##### 响应

```json
{
  "page_num": 1,
  "page_size": 1,
  "total": 8,
  "page_count": 8,
  "content": [
    {
      "id": 1,
      "name": "bug",
      "color": "#ED4014",
      "description": "Something isn't working",
      "text_color": "#FFFFFF",
      "expires_at": null,
      "is_expired": false
    }
  ]
}
```

### 8. 创建项目下的标签

``POST /api/v4/projects/{project_id}/labels``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | path | 字符串 | 项目ID | yes |
| name | body | 字符串 | 名字 | yes |
| color | body | 字符串 | 颜色 | yes |
| expires_at | body | 字符串 | 过期日期 | no |
| description | body | 字符串 | 描述 | no |

##### 响应

```json
{
  "id": 1,
  "name": "test_label",
  "color": "#2865E0",
  "description": "decription",
  "text_color": "#FFFFFF",
  "expires_at": null,
  "is_expired": false,
  "open_issues_count": 0,
  "closed_issues_count": 0,
  "open_merge_requests_count": 0,
  "subscribed": false,
  "priority": 0,
  "is_project_label": true
}
```

### 9. 更新项目下的标签

``PUT /api/v4/projects/{project_id}/labels``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | path | 字符串 | 项目ID | yes |
| name | body | 字符串 | 名字 | yes |
| new_name | body | 字符串 | 新名字 | yes |
| color | body | 字符串 | 颜色 | yes |
| expires_at | body | 字符串 | 过期日期 | no |
| description | body | 字符串 | 描述 | no |

##### 响应

```json
{
  "id": 1,
  "name": "test_label1",
  "color": "#2865E0",
  "description": "decription",
  "text_color": "#FFFFFF",
  "expires_at": null,
  "is_expired": false,
  "open_issues_count": 0,
  "closed_issues_count": 0,
  "open_merge_requests_count": 0,
  "subscribed": false,
  "priority": 0,
  "is_project_label": true
}
```

### 10. 删除项目下的标签

``DELETE /api/v4/projects/{project_id}/labels``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | path | 字符串 | 项目ID | yes |
| name | body | 字符串 | 待删除Label标题 | yes |

##### 响应

```json

```

### 11. 创建项目下的里程碑

#### POST /api/v4/projects/{project_id}/milestones

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | path | 字符串 | 项目ID | yes |
| title | body | 字符串 | 标题 | yes |
| due_date | body | 字符串 | 截止日期 | yes |
| description | body | 字符串 | 描述 | no |
| state_event | body | 字符串 | 状态 | no |
| start_date | body | 字符串 | 开始日期 | no |

##### 响应

```json
{
  "id": 1,
  "iid": 1,
  "project_id": 123456,
  "title": "test_milestone",
  "description": "description",
  "state": "active",
  "created_at": "2023-09-18T14:59:10.643+08:00",
  "updated_at": "2023-09-18T14:59:10.643+08:00",
  "due_date": "2024-01-01",
  "start_date": null,
  "project_path": "test_user/test_project",
  "web_url": "https://gitcode.com/test_user/test_project/milestones/1"
}
```

### 12. 获取项目下的里程碑列表

``GET /api/v4/projects/{project_id}/milestones``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | path | 字符串 | 项目ID | yes |
| sort | body | 字符串 | Return milestones sorted in `asc` or `desc` order. | no |
| state | body | 字符串 | 返回指定状态的里程碑，`active`（开启中）、`closed`（已关闭） 或 `all`（全部） | no |
| iids |  body | 整数 | 里程碑内部ID | no |
| search | body | 字符串 | 搜索关键字，按照标题或描述 | no |
| page | body | 整数 | 页码 | no |
| per_page | body | 整数 | 每页条数 | no |

##### 响应

```json
{
  "page_num": 1,
  "page_size": 1,
  "total": 1,
  "page_count": 1,
  "content": {
    "all": 1,
    "closed": 0,
    "opened": 1,
    "milestones": [
      {
        "id": 1,
        "iid": 1,
        "project_id": 123456,
        "title": "test_milestone",
        "description": "description",
        "state": "active",
        "created_at": "2023-09-18T14:59:10.643+08:00",
        "updated_at": "2023-09-18T14:59:10.643+08:00",
        "due_date": "2024-01-01",
        "start_date": null,
        "project_path": "test_user/test_project",
        "web_url": "https://gitcode.com/test_user/test_project/milestones/1",
        "issues_count": 0,
        "merge_requests_count": 0,
        "progress": 0
      }
    ]
  }
}
```

### 13. 删除项目下的里程碑

``DELETE /api/v4/projects/{project_id}/milestones/{milestone_id}``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | path | 字符串 | 项目ID | yes |
| milestone_id |  path | 整数 | 里程碑ID | yes |

##### 响应

```json

```

### 14. 更新项目下的里程碑

``PUT /api/v4/projects/{project_id}/milestones/{milestone_id}``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | path | 字符串 | 项目ID | yes |
| milestone_id |  path | 整数 | 里程碑ID | yes |
| title | body | 字符串 | 标题 | yes |
| due_date | body | 字符串 | 截止日期 | yes |
| description | body | 字符串 | 描述 | no |
| state_event | body | 字符串 | 状态 | no |
| start_date | body | 字符串 | 开始日期 | no |

##### 响应

```json
{
  "id": 1,
  "iid": 1,
  "project_id": 123456,
  "title": "test_milestone1",
  "description": "description",
  "state": "active",
  "created_at": "2023-09-18T14:59:10.643+08:00",
  "updated_at": "2023-09-18T15:01:29.427+08:00",
  "due_date": "2024-01-01",
  "start_date": null,
  "project_path": "test_user/test_project",
  "web_url": "https://gitcode.com/test_user/test_project/milestones/1"
}
```

### 15. 获取项目下的里程碑详情

``GET /api/v4/projects/{project_id}/milestones/{milestone_id}``

##### 参数

| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :-------: |
| project_id | body | 字符串 | 项目ID | yes |
| milestone_id | body | 整数 | 里程碑ID | yes |

##### 响应

```json
{
  "id": 1,
  "iid": 1,
  "project_id": 123456,
  "title": "test_milestone",
  "description": "description",
  "state": "active",
  "created_at": "2023-09-18T14:59:10.643+08:00",
  "updated_at": "2023-09-18T15:01:29.427+08:00",
  "due_date": "2024-01-01",
  "start_date": null,
  "project_path": "test_user/test_project",
  "web_url": "https://gitcode.com/test_user/test_project/milestones/1",
  "issues_count": 0,
  "merge_requests_count": 0,
  "progress": 0
}
```
