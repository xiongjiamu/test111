---
linkTitle: 任务 Issues
title: Issues接口文档
weight: 5
sidebar:
  open: false
---

## 1. 创建Issue

### 请求

`POST https://api.gitcode.com/api/v5/repos/{owner}/issues`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | formData | string    |
|  title* | Issue标题 | formData | string    |
|  body   | Issue描述 | formData | string    |
|  assignee | Issue负责人的个人空间地址 | formData | string    |
|  milestone  | 里程碑序号 | formData | integer    |
|  labels | 用逗号分开的标签，名称要求长度在 2-20 之间且非特殊字符。如: bug,performance | formData | string    |
|  security_hole   | 是否是私有issue(默认为false) | formData | string    |

### 响应

```json
{
	"id": 150663,
	"repository_url": "https://gitcode.com/qinmh2/zhishiku2222",
	"html_url": "https://gitcode.com/qinmh2/zhishiku2222/issues/3",
	"number": 3,
	"state": "open",
	"title": "123",
	"body": "test123",
	"user": {
		"id": 655,
		"login": "vinnie",
		"name": "Vinnie",
		"avatar_url": null,
		"html_url": "https://gitcode.com/vinnie",
		"type": "User"
	},
	"labels": [],
	"assignee": null,
	"repository": {
		"id": 155366,
		"full_name": "qinmh2/zhishiku2222",
		"path": "fbj-zhishiku2222",
		"name": "知识库",
		"description": "regret额度无发",
		"created_at": "2024-01-04T18:32:17.789+08:00",
		"updated_at": "2024-01-04T18:32:17.789+08:00"
	},
	"milestone": null,
	"created_at": "2024-04-14T12:05:45.712+08:00",
	"updated_at": "2024-04-14T12:05:45.712+08:00",
    "finished_at":null
}

```


## 2. 更新Issue

### 请求

`POST https://api.gitcode.com/api/v5/repos/{owner}/issues/{number}`

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | formData | string    |
|  number*   | Issue 编号(区分大小写，无需添加 # 号) | path | string    |
|  title | Issue标题 | formData | string    |
|  body   | Issue描述 | formData | string    |
|  state   | Issue 状态，open（开启的）、closed（关闭的） | formData | string    |
|  assignee | Issue负责人的个人空间地址 | formData | string    |
|  milestone  | 里程碑序号 | formData | integer    |
|  labels | 用逗号分开的标签，名称要求长度在 2-20 之间且非特殊字符。如: bug,performance | formData | string    |
|  security_hole   | 是否是私有issue(默认为false) | formData | string    |

### 响应

```json
{
	"id": 150663,
	"repository_url": "https://gitcode.com/qinmh2/zhishiku2222",
	"html_url": "https://gitcode.com/qinmh2/zhishiku2222/issues/3",
	"number": 3,
	"state": "open",
	"title": "123",
	"body": "test123",
	"user": {
		"id": 655,
		"login": "vinnie",
		"name": "Vinnie",
		"avatar_url": null,
		"html_url": "https://gitcode.com/vinnie",
		"type": "User"
	},
	"labels": [],
	"assignee": null,
	"repository": {
		"id": 155366,
		"full_name": "qinmh2/zhishiku2222",
		"path": "fbj-zhishiku2222",
		"name": "知识库",
		"description": "regret额度无发",
		"created_at": "2024-01-04T18:32:17.789+08:00",
		"updated_at": "2024-01-04T18:32:17.789+08:00"
	},
	"milestone": null,
	"created_at": "2024-04-14T12:05:45.712+08:00",
	"updated_at": "2024-04-14T12:05:45.712+08:00",
    "finished_at":null
}
```


## 3. 仓库的某个Issue

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/issues/{number}`

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  number*   | Issue 编号(区分大小写，无需添加 # 号) | path | string    |

### 响应

```json
{
	"id": 150663,
	"repository_url": "https://gitcode.com/qinmh2/zhishiku2222",
	"html_url": "https://gitcode.com/qinmh2/zhishiku2222/issues/3",
	"number": 3,
	"state": "open",
	"title": "123",
	"body": "test123",
	"user": {
		"id": 655,
		"login": "vinnie",
		"name": "Vinnie",
		"avatar_url": null,
		"html_url": "https://gitcode.com/vinnie",
		"type": "User"
	},
	"labels": [],
	"assignee": null,
	"repository": {
		"id": 155366,
		"full_name": "qinmh2/zhishiku2222",
		"path": "fbj-zhishiku2222",
		"name": "知识库",
		"description": "regret额度无发",
		"created_at": "2024-01-04T18:32:17.789+08:00",
		"updated_at": "2024-01-04T18:32:17.789+08:00",
        "assigner": {
            "id": 5648060,
            "login": "fbj1121",
            "name": "fbj1121",
            "avatar_url": null,
            "html_url": "https://gitcode.com/fbj1121",
            "type": "User"
        },
        "pushed_at": "2024-01-04T18:32:17.789+08:00",
        "created_at": "2024-01-04T18:32:17.789+08:00",
		"updated_at": "2024-01-04T18:32:17.789+08:00",
        "assignees_number": 1,
        "testers_number": 1,
        "assignee": [
            {
                "id": 5648060,
                "login": "fbj1121",
                "name": "fbj1121",
                "avatar_url": null,
                "html_url": "https://gitcode.com/fbj1121",
                "type": "User"
            }
        ],
        "testers": [
            {
                "id": 5648060,
                "login": "fbj1121",
                "name": "fbj1121",
                "avatar_url": null,
                "html_url": "https://gitcode.com/fbj1121",
                "type": "User"
            }
        ],
	},
	"milestone": null,
	"created_at": "2024-04-14T12:05:45.712+08:00",
	"updated_at": "2024-04-14T12:05:45.712+08:00",
    "finished_at": null
}
```


## 4. 获取仓库所有 issues

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/issues`

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  state   | Issue的状态: open（开启的）,  closed（关闭的）。 默认: open | query | string    |
|labels   | 用逗号分开的标签。如: bug,performance | query | string |
|sort   | 排序依据: 创建时间(created)，更新时间(updated_at)。默认: created_at | query | string |
|direction  | 排序方式: 升序(asc)，降序(desc)。默认: desc | query | string |
|since  | 起始的更新时间，要求时间格式为 ISO 8601 | query | string |
|page   | 当前的页码 | query | integer |
|per_page   | 每页的数量，最大为 100 | query | integer |
| created_at   | 任务创建时间，格式同上 | query | string |
| finished_at  | 任务完成时间，即任务最后一次转为已完成状态的时间点。格式同上 | query | string |
| milestone  | 根据里程碑标题。none为没里程碑的，*为所有带里程碑的 | query | string |
| assignee   | Issue指派人ID | query | string |
| creator  | 创建Issues的用户username | query | string |


### 响应

```json
[
    {
        "id": 15885154,
        "html_url": "https://gitcode.com/test_user/test_project/issues/1",
        "number": "1",
        "state": "opened",
        "title": "[Bug]: GridItem的selected属性支持双向绑定文档更新",
        "body": "### 发生了什么问题？\r\n\r\nGridItem的",
        "user": {
            "id": 6536510,
            "login": "tomkl123",
            "name": "tomkl123",
            "avatar_url": null,
            "html_url": "https://gitcode.com/tomkl123",
            "type": "User"
        },
        "labels": [
            {
                "id": 75265085,
                "color": "DB2828",
                "name": "bug",
                "created_at": "2020-08-18T17:41:20+08:00",
                "updated_at": "2024-04-14T19:24:01+08:00"
            }
        ],
        "assignee": null,
        "repository": {
            "id": 10919030,
            "full_name": "openharmony/docs",
            "human_name": "OpenHarmony/docs",
            "path": "docs",
            "name": "docs",
            "owner": {
                "id": 7928036,
                "login": "openharmony_admin",
                "name": "openharmony",
                "avatar_url": null,
                "html_url": "https://gitcode.com/openharmony_admin",
                "remark": "",
                "type": "User"
            },
            "assigner": {
                "id": 626123,
                "login": "landwind",
                "name": "mamingshuai",
                "avatar_url": null,
                "html_url": "https://gitcode.com/landwind",
                "remark": "",
                "type": "User"
            },
            "description": "OpenHarmony documentation | OpenHarmony开发者文档",
            "html_url": "https://gitcode.com/openharmony/docs.git",
            "ssh_url": "git@gitcode.com:openharmony/docs.git",
            "pushed_at": "2024-04-14T16:28:38+08:00",
            "created_at": "2020-08-14T14:49:40+08:00",
            "updated_at": "2024-04-14T17:46:01+08:00",
            "assignees_number": 1,
            "testers_number": 1,
            "assignee": [
                {
                    "id": 9329956,
                    "login": "Peter_1988",
                    "name": "Peter_1988",
                    "avatar_url": null,
                    "html_url": "https://gitcode.com/Peter_1988",
                    "remark": "",
                    "type": "User"
                }
            ],
            "testers": [
                {
                    "id": 7387629,
                    "login": "openharmony_ci",
                    "name": "openharmony_ci",
                    "avatar_url": null,
                    "html_url": "https://gitcode.com/openharmony_ci",
                    "remark": "",
                    "type": "User"
                }
            ],
        },
        "created_at": "2024-04-14T17:01:23+08:00",
        "updated_at": "2024-04-14T17:01:42+08:00",
    }
]
```


## 5. 获取仓库某个Issue所有的评论

### 请求

`GET /api/v4/projects/{project_id}/issues/{issue_iid}/notes`

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
| access_token | 用户授权码 | query | string |
| owner*|  仓库所属空间地址(企业、组织或个人的地址path) | path  | string |
| repo*| 仓库路径(path)  | path  | string |
| number*| Issue 编号(区分大小写，无需添加 # 号)  | path  | string |
| since |Only comments updated at or after this time are returned. This is a timestamp in ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ  | query | string |
| page | 当前的页码 | query | integer |
| per_page | 每页的数量，最大为 100 | query | integer |
| order |排序顺序: asc(default),desc | query | string |


### 响应

```json
[
    {
        "id": 26383103,
        "body": "测试评论",
        "user": {
            "id": 5648060,
            "login": "fbj1121",
            "name": "fbj1121",
            "avatar_url": null,
            "html_url": "https://gitcode.com/fbj1121"
        },
        "source": null,
        "target": {
            "issue": {
                "id": 15798287,
                "title": "test2222",
                "number": 1
            },
            "pull_request": null
        },
        "created_at": "2024-04-07T19:00:56+08:00",
        "updated_at": "2024-04-07T19:00:56+08:00"
    }
]
```


## 6. 获取仓库所有 iusse 评论

### 请求 暂无

`GET `

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
| access_token  | 用户授权码 | query | string |
| owner* |  仓库所属空间地址(企业、组织或个人的地址path) | path  | string |
| repo* | 仓库路径(path)  | path  | string |
| sort  | Either created or updated. Default: created | query | string |
| direction  |Either asc or desc. Ignored without the sort parameter. | query | string |
| since  |Only comments updated at or after this time are returned. This is a timestamp in ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ  | query | string |
| page  | 当前的页码 | query | integer |
| per_page  | 每页的数量，最大为 100 | query | integer |


### 响应

```json
[
    {
        "id": 26383103,
        "body": "测试评论",
        "user": {
            "id": 5648060,
            "login": "fbj1121",
            "name": "fbj1121",
            "avatar_url": null,
            "html_url": "https://gitcode.com/fbj1121",
            "type": "User"
        },
        "source": null,
        "target": {
            "issue": {
                "id": 15798287,
                "title": "test2222",
                "number": 1
            },
            "pull_request": null
        },
        "created_at": "2024-04-07T19:00:56+08:00",
        "updated_at": "2024-04-07T19:00:56+08:00"
    }
]
```


## 7. 获取 iusse 关联的 pull requests

### 请求 暂无

`GET `

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
| access_token  |   用户授权码  | query  | string |
| owner* |  仓库所属空间地址(企业、组织或个人的地址path)  | path   | string |
| repo  |   仓库路径(path) | query  | string |
| number* | Issue 编号(区分大小写，无需添加 # 号)   | path   | string |


### 响应

```json
[
    {
        "id": 1,
        "html_url": "https://gitcode.com/foo/bar/merge_requests/1",
        "number": 1,
        "state": "opened",
        "assignees_number": 1,
        "testers_number": 1,
        "milestone": null,
        "created_at": "2024-04-07T18:44:43+08:00",
        "updated_at": "2024-04-07T18:44:45+08:00",
        "merged_at": "2024-04-07T18:44:43+08:00",
        "closed_at":"2024-04-07T18:44:45+08:00",
        "user": {
            "id": 5648060,
            "login": "fbj1121",
            "name": "fbj1121",
            "avatar_url": null,
            "html_url": "https://gitcode.com/fbj1121",
            "type": "User"
        },
        "title": "test v1.0 to v1.1",
        "body": "test",
    }
]
```
