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


## 4. 获取仓库所有 issue

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/issues`

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  state   | Issue的状态: open（开启的）, progressing(进行中), closed（关闭的）, rejected（拒绝的）。 默认: open | query | string    |
|labels   | 用逗号分开的标签。如: bug,performance | query | string |
|sort   | 排序依据: 创建时间(created)，更新时间(updated_at)。默认: created_at | query | string |
|direction  | 排序方式: 升序(asc)，降序(desc)。默认: desc | query | string |
|since  | 起始的更新时间，要求时间格式为 ISO 8601 | query | string |
|page   | 当前的页码 | query | integer |
|per_page   | 每页的数量，最大为 100 | query | integer |
|~~schedule~~   | 计划开始日期，格式：20181006T173008+80-20181007T173008+80（区间），或者 -20181007T173008+80（小于20181007T173008+80），或者 20181006T173008+80-（大于20181006T173008+80），要求时间格式为20181006T173008+80 | query | string |
| ~~deadline~~   | 计划截止日期，格式同上 | query | string |
| created_at   | 任务创建时间，格式同上 | query | string |
| finished_at  | 任务完成时间，即任务最后一次转为已完成状态的时间点。格式同上 | query | string |
| milestone  | 根据里程碑标题。none为没里程碑的，*为所有带里程碑的 | query | string |
| assignee   | 用户的username。 none为没指派者, *为所有带有指派者的 | query | string |
| creator  | 创建Issues的用户username | query | string |
| ~~program~~  | 所属项目名称。none为没有所属项目，*为所有带所属项目的 | query | string|


### 响应

```json
[
    {
        "id": 15885154,
        "html_url": "https://gitee.com/openharmony/docs/issues/I9GH2A",
        "number": "I9GH2A",
        "state": "open",
        "title": "[Bug]: GridItem的selected属性支持双向绑定文档更新",
        "body": "### 发生了什么问题？\r\n\r\nGridItem的",
        "user": {
            "id": 6536510,
            "login": "tomkl123",
            "name": "tomkl123",
            "avatar_url": "https://gitee.com/assets/no_portrait.png",
            "html_url": "https://gitee.com/tomkl123",
            "type": "User"
        },
        "labels": [
            {
                "id": 75265085,
                "color": "DB2828",
                "name": "bug",
                "repository_id": 6511493,
                "url": "https://gitee.com/api/v5/enterprises/open_harmony/labels/bug",
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
                "avatar_url": "https://foruda.gitee.com/avatar/1677102952566165682/7928036_openharmony_admin_1622551091.png",
                "html_url": "https://gitee.com/openharmony_admin",
                "remark": "",
                "type": "User"
            },
            "assigner": {
                "id": 626123,
                "login": "landwind",
                "name": "mamingshuai",
                "avatar_url": "https://gitee.com/assets/no_portrait.png",
                "url": "https://gitee.com/api/v5/users/landwind",
                "html_url": "https://gitee.com/landwind",
                "remark": "",
                "type": "User"
            },
            "description": "OpenHarmony documentation | OpenHarmony开发者文档",
            "html_url": "https://gitee.com/openharmony/docs.git",
            "ssh_url": "git@gitee.com:openharmony/docs.git",
            "pushed_at": "2024-04-14T16:28:38+08:00",
            "created_at": "2020-08-14T14:49:40+08:00",
            "updated_at": "2024-04-14T17:46:01+08:00",
            "paas": null,
            "assignees_number": 1,
            "testers_number": 1,
            "assignee": [
                {
                    "id": 9329956,
                    "login": "Peter_1988",
                    "name": "Peter_1988",
                    "avatar_url": "https://gitee.com/assets/no_portrait.png",
                    "url": "https://gitee.com/api/v5/users/Peter_1988",
                    "html_url": "https://gitee.com/Peter_1988",
                    "remark": "",
                    "type": "User"
                }
            ],
            "testers": [
                {
                    "id": 7387629,
                    "login": "openharmony_ci",
                    "name": "openharmony_ci",
                    "avatar_url": "https://foruda.gitee.com/avatar/1677075395966509043/7387629_openharmony_ci_1656582662.png",
                    "url": "https://gitee.com/api/v5/users/openharmony_ci",
                    "html_url": "https://gitee.com/openharmony_ci",
                    "remark": "",
                    "type": "User"
                }
            ],
            "status": "开始"
        },
        "milestone": null,
        "created_at": "2024-04-14T17:01:23+08:00",
        "updated_at": "2024-04-14T17:01:42+08:00",
        "issue_state_detail": {
            "id": 877565,
            "title": "待办的",
            "color": "#8c92a4",
            "icon": "icon-task-state-21",
            "command": null,
            "serial": 0,
            "created_at": "2020-08-18T17:41:18+08:00",
            "updated_at": "2020-08-18T17:41:18+08:00"
        }
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
            "avatar_url": "https://gitee.com/assets/no_portrait.png",
            "url": "https://gitee.com/api/v5/users/fbj1121",
            "html_url": "https://gitee.com/fbj1121"
        },
        "source": null,
        "target": {
            "issue": {
                "id": 15798287,
                "title": "test2222",
                "number": "I9EM1B"
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
            "avatar_url": "https://gitee.com/assets/no_portrait.png",
            "url": "https://gitee.com/api/v5/users/fbj1121",
            "html_url": "https://gitee.com/fbj1121",
            "remark": "",
            "followers_url": "https://gitee.com/api/v5/users/fbj1121/followers",
            "following_url": "https://gitee.com/api/v5/users/fbj1121/following_url{/other_user}",
            "gists_url": "https://gitee.com/api/v5/users/fbj1121/gists{/gist_id}",
            "starred_url": "https://gitee.com/api/v5/users/fbj1121/starred{/owner}{/repo}",
            "subscriptions_url": "https://gitee.com/api/v5/users/fbj1121/subscriptions",
            "organizations_url": "https://gitee.com/api/v5/users/fbj1121/orgs",
            "repos_url": "https://gitee.com/api/v5/users/fbj1121/repos",
            "events_url": "https://gitee.com/api/v5/users/fbj1121/events{/privacy}",
            "received_events_url": "https://gitee.com/api/v5/users/fbj1121/received_events",
            "type": "User"
        },
        "source": null,
        "target": {
            "issue": {
                "id": 15798287,
                "title": "test2222",
                "number": "I9EM1B"
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
        "id": 11206285,
        "url": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/1",
        "html_url": "https://gitee.com/fbj1121/fbj-test/pulls/1",
        "diff_url": "https://gitee.com/fbj1121/fbj-test/pulls/1.diff",
        "patch_url": "https://gitee.com/fbj1121/fbj-test/pulls/1.patch",
        "issue_url": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/1/issues",
        "commits_url": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/1/commits",
        "review_comments_url": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/comments/{/number}",
        "review_comment_url": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/comments",
        "comments_url": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/1/comments",
        "number": 1,
        "close_related_issue": 1,
        "prune_branch": false,
        "state": "open",
        "assignees_number": 1,
        "testers_number": 1,
        "assignees": [
            {
                "id": 5648060,
                "login": "fbj1121",
                "name": "fbj1121",
                "avatar_url": "https://gitee.com/assets/no_portrait.png",
                "url": "https://gitee.com/api//users/fbj1121",
                "html_url": "https://gitee.com/fbj1121",
                "remark": "",
                "followers_url": "https://gitee.com/api//users/fbj1121/followers",
                "following_url": "https://gitee.com/api//users/fbj1121/following_url{/other_user}",
                "gists_url": "https://gitee.com/api//users/fbj1121/gists{/gist_id}",
                "starred_url": "https://gitee.com/api//users/fbj1121/starred{/owner}{/repo}",
                "subscriptions_url": "https://gitee.com/api//users/fbj1121/subscriptions",
                "organizations_url": "https://gitee.com/api//users/fbj1121/orgs",
                "repos_url": "https://gitee.com/api//users/fbj1121/repos",
                "events_url": "https://gitee.com/api//users/fbj1121/events{/privacy}",
                "received_events_url": "https://gitee.com/api//users/fbj1121/received_events",
                "type": "User",
                "assignee": true,
                "code_owner": false,
                "accept": false
            }
        ],
        "testers": [
            {
                "id": 5648060,
                "login": "fbj1121",
                "name": "fbj1121",
                "avatar_url": "https://gitee.com/assets/no_portrait.png",
                "url": "https://gitee.com/api//users/fbj1121",
                "html_url": "https://gitee.com/fbj1121",
                "remark": "",
                "followers_url": "https://gitee.com/api//users/fbj1121/followers",
                "following_url": "https://gitee.com/api//users/fbj1121/following_url{/other_user}",
                "gists_url": "https://gitee.com/api//users/fbj1121/gists{/gist_id}",
                "starred_url": "https://gitee.com/api//users/fbj1121/starred{/owner}{/repo}",
                "subscriptions_url": "https://gitee.com/api//users/fbj1121/subscriptions",
                "organizations_url": "https://gitee.com/api//users/fbj1121/orgs",
                "repos_url": "https://gitee.com/api//users/fbj1121/repos",
                "events_url": "https://gitee.com/api//users/fbj1121/events{/privacy}",
                "received_events_url": "https://gitee.com/api//users/fbj1121/received_events",
                "type": "User",
                "assignee": true,
                "code_owner": false,
                "accept": false
            }
        ],
        "milestone": null,
        "labels": [],
        "locked": false,
        "created_at": "2024-04-07T18:44:43+08:00",
        "updated_at": "2024-04-07T18:44:45+08:00",
        "closed_at": null,
        "draft": false,
        "merged_at": null,
        "mergeable": true,
        "can_merge_check": false,
        "_links": {
            "self": {
                "href": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/1"
            },
            "html": {
                "href": "https://gitee.com/fbj1121/fbj-test/pulls/1"
            },
            "issue": {
                "href": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/1/issues"
            },
            "comments": {
                "href": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/1/comments"
            },
            "review_comments": {
                "href": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/1/comments"
            },
            "review_comment": {
                "href": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/comments/1"
            },
            "commits": {
                "href": "https://gitee.com/api/v5/repos/fbj1121/fbj-test/pulls/1/commits"
            }
        },
        "user": {
            "id": 5648060,
            "login": "fbj1121",
            "name": "fbj1121",
            "avatar_url": "https://gitee.com/assets/no_portrait.png",
            "url": "https://gitee.com/api/v5/users/fbj1121",
            "html_url": "https://gitee.com/fbj1121",
            "remark": "",
            "followers_url": "https://gitee.com/api/v5/users/fbj1121/followers",
            "following_url": "https://gitee.com/api/v5/users/fbj1121/following_url{/other_user}",
            "gists_url": "https://gitee.com/api/v5/users/fbj1121/gists{/gist_id}",
            "starred_url": "https://gitee.com/api/v5/users/fbj1121/starred{/owner}{/repo}",
            "subscriptions_url": "https://gitee.com/api/v5/users/fbj1121/subscriptions",
            "organizations_url": "https://gitee.com/api/v5/users/fbj1121/orgs",
            "repos_url": "https://gitee.com/api/v5/users/fbj1121/repos",
            "events_url": "https://gitee.com/api/v5/users/fbj1121/events{/privacy}",
            "received_events_url": "https://gitee.com/api/v5/users/fbj1121/received_events",
            "type": "User"
        },
        "ref_pull_requests": [],
        "title": "test v1.0 to v1.1",
        "body": "test",
        "head": {
            "label": "dev",
            "ref": "dev",
            "sha": "5449c16d176d16f3fcbd2dd205ee571d1312b8b3",
            "user": {
                "id": 5648060,
                "login": "fbj1121",
                "name": "fbj1121",
                "avatar_url": "https://gitee.com/assets/no_portrait.png",
                "url": "https://gitee.com/api/v5/users/fbj1121",
                "html_url": "https://gitee.com/fbj1121",
                "remark": "",
                "followers_url": "https://gitee.com/api/v5/users/fbj1121/followers",
                "following_url": "https://gitee.com/api/v5/users/fbj1121/following_url{/other_user}",
                "gists_url": "https://gitee.com/api/v5/users/fbj1121/gists{/gist_id}",
                "starred_url": "https://gitee.com/api/v5/users/fbj1121/starred{/owner}{/repo}",
                "subscriptions_url": "https://gitee.com/api/v5/users/fbj1121/subscriptions",
                "organizations_url": "https://gitee.com/api/v5/users/fbj1121/orgs",
                "repos_url": "https://gitee.com/api/v5/users/fbj1121/repos",
                "events_url": "https://gitee.com/api/v5/users/fbj1121/events{/privacy}",
                "received_events_url": "https://gitee.com/api/v5/users/fbj1121/received_events",
                "type": "User"
            },
            "repo": {
                "id": 34059843,
                "full_name": "fbj1121/fbj-test",
                "human_name": "fbj1121/fbj-test",
                "url": "https://gitee.com/api/v5/repos/fbj1121/fbj-test",
                "namespace": {
                    "id": 5641036,
                    "type": "personal",
                    "name": "fbj1121",
                    "path": "fbj1121",
                    "html_url": "https://gitee.com/fbj1121"
                },
                "path": "fbj-test",
                "name": "fbj-test",
                "owner": {
                    "id": 5648060,
                    "login": "fbj1121",
                    "name": "fbj1121",
                    "avatar_url": "https://gitee.com/assets/no_portrait.png",
                    "url": "https://gitee.com/api/v5/users/fbj1121",
                    "html_url": "https://gitee.com/fbj1121",
                    "remark": "",
                    "followers_url": "https://gitee.com/api/v5/users/fbj1121/followers",
                    "following_url": "https://gitee.com/api/v5/users/fbj1121/following_url{/other_user}",
                    "gists_url": "https://gitee.com/api/v5/users/fbj1121/gists{/gist_id}",
                    "starred_url": "https://gitee.com/api/v5/users/fbj1121/starred{/owner}{/repo}",
                    "subscriptions_url": "https://gitee.com/api/v5/users/fbj1121/subscriptions",
                    "organizations_url": "https://gitee.com/api/v5/users/fbj1121/orgs",
                    "repos_url": "https://gitee.com/api/v5/users/fbj1121/repos",
                    "events_url": "https://gitee.com/api/v5/users/fbj1121/events{/privacy}",
                    "received_events_url": "https://gitee.com/api/v5/users/fbj1121/received_events",
                    "type": "User"
                },
                "assigner": {
                    "id": 5648060,
                    "login": "fbj1121",
                    "name": "fbj1121",
                    "avatar_url": "https://gitee.com/assets/no_portrait.png",
                    "url": "https://gitee.com/api/v5/users/fbj1121",
                    "html_url": "https://gitee.com/fbj1121",
                    "remark": "",
                    "followers_url": "https://gitee.com/api/v5/users/fbj1121/followers",
                    "following_url": "https://gitee.com/api/v5/users/fbj1121/following_url{/other_user}",
                    "gists_url": "https://gitee.com/api/v5/users/fbj1121/gists{/gist_id}",
                    "starred_url": "https://gitee.com/api/v5/users/fbj1121/starred{/owner}{/repo}",
                    "subscriptions_url": "https://gitee.com/api/v5/users/fbj1121/subscriptions",
                    "organizations_url": "https://gitee.com/api/v5/users/fbj1121/orgs",
                    "repos_url": "https://gitee.com/api/v5/users/fbj1121/repos",
                    "events_url": "https://gitee.com/api/v5/users/fbj1121/events{/privacy}",
                    "received_events_url": "https://gitee.com/api/v5/users/fbj1121/received_events",
                    "type": "User"
                },
                "description": "",
                "private": false,
                "public": true,
                "internal": false,
                "fork": false,
                "html_url": "https://gitee.com/fbj1121/fbj-test.git",
                "ssh_url": "git@gitee.com:fbj1121/fbj-test.git"
            }
        },
        "base": {
            "label": "master",
            "ref": "master",
            "sha": "ffada9268048d4e22f86fac1b4b31b4884475068",
            "user": {
                "id": 5648060,
                "login": "fbj1121",
                "name": "fbj1121",
                "avatar_url": "https://gitee.com/assets/no_portrait.png",
                "url": "https://gitee.com/api/v5/users/fbj1121",
                "html_url": "https://gitee.com/fbj1121",
                "remark": "",
                "followers_url": "https://gitee.com/api/v5/users/fbj1121/followers",
                "following_url": "https://gitee.com/api/v5/users/fbj1121/following_url{/other_user}",
                "gists_url": "https://gitee.com/api/v5/users/fbj1121/gists{/gist_id}",
                "starred_url": "https://gitee.com/api/v5/users/fbj1121/starred{/owner}{/repo}",
                "subscriptions_url": "https://gitee.com/api/v5/users/fbj1121/subscriptions",
                "organizations_url": "https://gitee.com/api/v5/users/fbj1121/orgs",
                "repos_url": "https://gitee.com/api/v5/users/fbj1121/repos",
                "events_url": "https://gitee.com/api/v5/users/fbj1121/events{/privacy}",
                "received_events_url": "https://gitee.com/api/v5/users/fbj1121/received_events",
                "type": "User"
            },
            "repo": {
                "id": 34059843,
                "full_name": "fbj1121/fbj-test",
                "human_name": "fbj1121/fbj-test",
                "url": "https://gitee.com/api/v5/repos/fbj1121/fbj-test",
                "namespace": {
                    "id": 5641036,
                    "type": "personal",
                    "name": "fbj1121",
                    "path": "fbj1121",
                    "html_url": "https://gitee.com/fbj1121"
                },
                "path": "fbj-test",
                "name": "fbj-test",
                "owner": {
                    "id": 5648060,
                    "login": "fbj1121",
                    "name": "fbj1121",
                    "avatar_url": "https://gitee.com/assets/no_portrait.png",
                    "url": "https://gitee.com/api/v5/users/fbj1121",
                    "html_url": "https://gitee.com/fbj1121",
                    "remark": "",
                    "followers_url": "https://gitee.com/api/v5/users/fbj1121/followers",
                    "following_url": "https://gitee.com/api/v5/users/fbj1121/following_url{/other_user}",
                    "gists_url": "https://gitee.com/api/v5/users/fbj1121/gists{/gist_id}",
                    "starred_url": "https://gitee.com/api/v5/users/fbj1121/starred{/owner}{/repo}",
                    "subscriptions_url": "https://gitee.com/api/v5/users/fbj1121/subscriptions",
                    "organizations_url": "https://gitee.com/api/v5/users/fbj1121/orgs",
                    "repos_url": "https://gitee.com/api/v5/users/fbj1121/repos",
                    "events_url": "https://gitee.com/api/v5/users/fbj1121/events{/privacy}",
                    "received_events_url": "https://gitee.com/api/v5/users/fbj1121/received_events",
                    "type": "User"
                },
                "assigner": {
                    "id": 5648060,
                    "login": "fbj1121",
                    "name": "fbj1121",
                    "avatar_url": "https://gitee.com/assets/no_portrait.png",
                    "url": "https://gitee.com/api/v5/users/fbj1121",
                    "html_url": "https://gitee.com/fbj1121",
                    "remark": "",
                    "followers_url": "https://gitee.com/api/v5/users/fbj1121/followers",
                    "following_url": "https://gitee.com/api/v5/users/fbj1121/following_url{/other_user}",
                    "gists_url": "https://gitee.com/api/v5/users/fbj1121/gists{/gist_id}",
                    "starred_url": "https://gitee.com/api/v5/users/fbj1121/starred{/owner}{/repo}",
                    "subscriptions_url": "https://gitee.com/api/v5/users/fbj1121/subscriptions",
                    "organizations_url": "https://gitee.com/api/v5/users/fbj1121/orgs",
                    "repos_url": "https://gitee.com/api/v5/users/fbj1121/repos",
                    "events_url": "https://gitee.com/api/v5/users/fbj1121/events{/privacy}",
                    "received_events_url": "https://gitee.com/api/v5/users/fbj1121/received_events",
                    "type": "User"
                },
                "description": "",
                "private": false,
                "public": true,
                "internal": false,
                "fork": false,
                "html_url": "https://gitee.com/fbj1121/fbj-test.git",
                "ssh_url": "git@gitee.com:fbj1121/fbj-test.git"
            }
        }
    }
]
```

