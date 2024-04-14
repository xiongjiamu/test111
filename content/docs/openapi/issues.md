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