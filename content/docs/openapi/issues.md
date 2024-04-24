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
|  repo*   | 仓库路径 | formData | string    |
|  title* | Issue标题 | formData | string    |
|  body   | Issue描述 | formData | string    |
|  assignee | Issue负责人的个人空间地址 | formData | string    |
|  milestone  | 里程碑序号 | formData | integer    |
|  labels | 用逗号分开的标签，名称要求长度在 2-20 之间且非特殊字符。如: bug,performance | formData | string    |
|  security_hole   | 是否是私有issue(默认为false) | formData | string    |

### 响应

```json
{
    "id": 152642,
    "html_url": "https://test.gitcode.net/dengmengmian/test01/issues/15",
    "number": 15,
    "state": "opened",
    "title": "半月据",
    "body": "节油料被引系活力级少本化段维家住实。常气前步证时第样日所阶效温界到量。个导土机技亲布接增论始高世收圆流级集。此般区才听党机达两收文斗公加白。代军前分写第图美市与道及间。",
    "user": {
        "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
        "events_url": null,
        "followers_url": null,
        "following_url": null,
        "gists_url": null,
        "html_url": "https://test.gitcode.net/dengmengmian",
        "id": "661ce4eab470b1430d456154",
        "login": "dengmengmian",
        "member_role": null,
        "name": "麻凡_",
        "organizations_url": null,
        "received_events_url": null,
        "remark": null,
        "repos_url": null,
        "starred_url": null,
        "subscriptions_url": null,
        "type": null,
        "url": null
    },
    "assignee": null,
    "repository": {
        "id": 152642,
        "full_name": "dengmengmian/test01",
        "path": "test01",
        "name": "test01",
        "description": "",
        "created_at": "2024-04-18T14:35:15.479+08:00",
        "updated_at": "2024-04-18T14:35:15.479+08:00"
    },
    "created_at": "2024-04-18T14:35:15.479+08:00",
    "updated_at": "2024-04-18T14:35:15.479+08:00",
    "finished_at": null,
    "labels": [
        {
            "id": 382379,
            "name": "enim",
            "color": "#428BCA"
        },
        {
            "id": 382378,
            "name": "proident",
            "color": "#428BCA"
        },
        {
            "id": 382377,
            "name": "qui",
            "color": "#428BCA"
        }
    ]
}

```


## 2. 更新Issue

### 请求

`PATCH https://api.gitcode.com/api/v5/repos/{owner}/issues/{number}`

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径 | formData | string    |
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
    "id": 152467,
    "html_url": "https://test.gitcode.net/dengmengmian/test01/issues/14",
    "number": 14,
    "state": "closed",
    "title": "取属且阶",
    "body": "速军间问备题意自系建技至速。那照与受证们老则使六么信。联不格决白转数特先到接单备心样本及。比论受感此中成要则片会受争里领周局。",
    "user": {
        "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
        "events_url": null,
        "followers_url": null,
        "following_url": null,
        "gists_url": null,
        "html_url": "https://test.gitcode.net/dengmengmian",
        "id": "661ce4eab470b1430d456154",
        "login": "dengmengmian",
        "member_role": null,
        "name": "麻凡_",
        "organizations_url": null,
        "received_events_url": null,
        "remark": null,
        "repos_url": null,
        "starred_url": null,
        "subscriptions_url": null,
        "type": null,
        "url": null
    },
    "assignee": {
        "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
        "events_url": null,
        "followers_url": null,
        "following_url": null,
        "gists_url": null,
        "html_url": "https://test.gitcode.net/dengmengmian",
        "id": "661ce4eab470b1430d456154",
        "login": "dengmengmian",
        "member_role": null,
        "name": "麻凡_",
        "organizations_url": null,
        "received_events_url": null,
        "remark": null,
        "repos_url": null,
        "starred_url": null,
        "subscriptions_url": null,
        "type": null,
        "url": null
    },
    "repository": {
        "id": 152467,
        "full_name": "dengmengmian/test01",
        "path": "test01",
        "name": "test01",
        "description": "",
        "created_at": "2024-04-16T14:38:43.464+08:00",
        "updated_at": "2024-04-18T18:27:21.955+08:00"
    },
    "created_at": "2024-04-16T14:38:43.464+08:00",
    "updated_at": "2024-04-18T18:27:21.955+08:00",
    "finished_at": "2024-04-16T14:49:45.166+08:00",
    "labels": [
        {
            "id": 382389,
            "name": "ad",
            "color": "#428BCA"
        },
        {
            "id": 382388,
            "name": "id",
            "color": "#428BCA"
        }
    ]
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
    "id": 152212,
    "html_url": "https://test.gitcode.net/dengmengmian/test01/issues/3",
    "number": 3,
    "state": "opened",
    "title": "查员种金交片",
    "body": "而很资七图数指反系并物众示易今高。运边月发红条亲才调二心点上米面世其分。由众计比维选作小指件每酸一见基历。向九又中国层合感内两米或自很转的。",
    "user": {
        "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
        "events_url": null,
        "followers_url": null,
        "following_url": null,
        "gists_url": null,
        "html_url": "https://test.gitcode.net/dengmengmian",
        "id": "661ce4eab470b1430d456154",
        "login": "dengmengmian",
        "member_role": null,
        "name": "麻凡_",
        "organizations_url": null,
        "received_events_url": null,
        "remark": null,
        "repos_url": null,
        "starred_url": null,
        "subscriptions_url": null,
        "type": null,
        "url": null
    },
    "assignee": null,
    "repository": {
        "id": 280713,
        "full_name": "dengmengmian / test01",
        "path": "test01",
        "name": "test01",
        "description": "",
        "created_at": "2024-04-15T16:27:45.090+08:00",
        "updated_at": "2024-04-15T16:27:45.090+08:00",
        "assigner": null,
        "pushed_at": null,
        "paas": null,
        "assignees_number": null,
        "testers_number": null,
        "assignee": null,
        "testers": null
    },
    "created_at": "2024-04-15T21:58:21.188+08:00",
    "updated_at": "2024-04-15T21:58:21.188+08:00",
    "finished_at": null,
    "labels": [],
    "priority": null,
    "issue_type": null,
    "issue_state": "opened",
    "issue_state_detail": null
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
        "id": 152642,
        "html_url": "https://test.gitcode.net/dengmengmian/test01/issues/15",
        "number": 15,
        "state": "opened",
        "title": "半月据",
        "body": "节油料被引系活力级少本化段维家住实。常气前步证时第样日所阶效温界到量。个导土机技亲布接增论始高世收圆流级集。此般区才听党机达两收文斗公加白。代军前分写第图美市与道及间。",
        "user": {
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
            "events_url": null,
            "followers_url": null,
            "following_url": null,
            "gists_url": null,
            "html_url": "https://test.gitcode.net/dengmengmian",
            "id": "661ce4eab470b1430d456154",
            "login": "dengmengmian",
            "member_role": null,
            "name": "麻凡_",
            "organizations_url": null,
            "received_events_url": null,
            "remark": null,
            "repos_url": null,
            "starred_url": null,
            "subscriptions_url": null,
            "type": null,
            "url": null
        },
        "assignee": {
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
            "events_url": null,
            "followers_url": null,
            "following_url": null,
            "gists_url": null,
            "html_url": "https://test.gitcode.net/dengmengmian",
            "id": "661ce4eab470b1430d456154",
            "login": "dengmengmian",
            "member_role": null,
            "name": "麻凡_",
            "organizations_url": null,
            "received_events_url": null,
            "remark": null,
            "repos_url": null,
            "starred_url": null,
            "subscriptions_url": null,
            "type": null,
            "url": null
        },
        "repository": {
            "id": 280713,
            "full_name": "dengmengmian / test01",
            "path": "test01",
            "name": "test01",
            "description": "",
            "created_at": "2024-04-15T16:27:45.090+08:00",
            "updated_at": "2024-04-15T16:27:45.090+08:00",
            "assigner": null,
            "pushed_at": null,
            "paas": null,
            "assignees_number": null,
            "testers_number": null,
            "assignee": null,
            "testers": null
        },
        "created_at": "2024-04-18T14:35:15.479+08:00",
        "updated_at": "2024-04-20T15:20:30.111+08:00",
        "finished_at": null,
        "labels": [
            {
                "id": 382379,
                "name": "enim",
                "color": "#428BCA"
            },
            {
                "id": 382378,
                "name": "proident",
                "color": "#428BCA"
            },
            {
                "id": 382377,
                "name": "qui",
                "color": "#428BCA"
            }
        ],
        "priority": null,
        "issue_type": null,
        "issue_state": "opened",
        "issue_state_detail": null
    },
    {
        "id": 152467,
        "html_url": "https://test.gitcode.net/dengmengmian/test01/issues/14",
        "number": 14,
        "state": "closed",
        "title": "取属且阶",
        "body": "速军间问备题意自系建技至速。那照与受证们老则使六么信。联不格决白转数特先到接单备心样本及。比论受感此中成要则片会受争里领周局。",
        "user": {
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
            "events_url": null,
            "followers_url": null,
            "following_url": null,
            "gists_url": null,
            "html_url": "https://test.gitcode.net/dengmengmian",
            "id": "661ce4eab470b1430d456154",
            "login": "dengmengmian",
            "member_role": null,
            "name": "麻凡_",
            "organizations_url": null,
            "received_events_url": null,
            "remark": null,
            "repos_url": null,
            "starred_url": null,
            "subscriptions_url": null,
            "type": null,
            "url": null
        },
        "assignee": {
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
            "events_url": null,
            "followers_url": null,
            "following_url": null,
            "gists_url": null,
            "html_url": "https://test.gitcode.net/dengmengmian",
            "id": "661ce4eab470b1430d456154",
            "login": "dengmengmian",
            "member_role": null,
            "name": "麻凡_",
            "organizations_url": null,
            "received_events_url": null,
            "remark": null,
            "repos_url": null,
            "starred_url": null,
            "subscriptions_url": null,
            "type": null,
            "url": null
        },
        "repository": {
            "id": 280713,
            "full_name": "dengmengmian / test01",
            "path": "test01",
            "name": "test01",
            "description": "",
            "created_at": "2024-04-15T16:27:45.090+08:00",
            "updated_at": "2024-04-15T16:27:45.090+08:00",
            "assigner": null,
            "pushed_at": null,
            "paas": null,
            "assignees_number": null,
            "testers_number": null,
            "assignee": null,
            "testers": null
        },
        "created_at": "2024-04-16T14:38:43.464+08:00",
        "updated_at": "2024-04-18T18:27:21.955+08:00",
        "finished_at": "2024-04-16T14:49:45.166+08:00",
        "labels": [
            {
                "id": 382389,
                "name": "ad",
                "color": "#428BCA"
            },
            {
                "id": 382388,
                "name": "id",
                "color": "#428BCA"
            }
        ],
        "priority": null,
        "issue_type": null,
        "issue_state": "closed",
        "issue_state_detail": null
    }
]
```


## 5. 获取仓库某个Issue所有的评论

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/issues/{number}/comments`

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
        "id": 271624,
        "body": "评论内容。",
        "user": {
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
            "events_url": null,
            "followers_url": null,
            "following_url": null,
            "gists_url": null,
            "html_url": "https://test.gitcode.net/dengmengmian",
            "id": "661ce4eab470b1430d456154",
            "login": "dengmengmian",
            "member_role": null,
            "name": "麻凡_",
            "organizations_url": null,
            "received_events_url": null,
            "remark": null,
            "repos_url": null,
            "starred_url": null,
            "subscriptions_url": null,
            "type": null,
            "url": null
        },
        "target": {
            "issue": {
                "id": 152134,
                "title": "",
                "nubmer": 1
            }
        },
        "created_at": "2024-04-19T17:50:18.199+08:00",
        "updated_at": "2024-04-19T17:50:18.199+08:00"
    }
]
```


## 6. 获取仓库所有 iusse 评论

### 请求 暂无

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/issues/comments `

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
        "id": 272201,
        "body": "daetete",
        "user": {
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
            "events_url": null,
            "followers_url": null,
            "following_url": null,
            "gists_url": null,
            "html_url": "https://test.gitcode.net/dengmengmian",
            "id": "661ce4eab470b1430d456154",
            "login": "dengmengmian",
            "member_role": null,
            "name": "麻凡_",
            "organizations_url": null,
            "received_events_url": null,
            "remark": null,
            "repos_url": null,
            "starred_url": null,
            "subscriptions_url": null,
            "type": null,
            "url": null
        },
        "target": {
            "issue": {
                "id": 152642,
                "title": "半月据",
                "nubmer": 15
            }
        },
        "created_at": "2024-04-20T15:20:30.104+08:00",
        "updated_at": null
    }
]
```


## 7. 获取 issue 关联的 pull requests

### 请求 暂无

`GET  https://api.gitcode.com/api/v5/repos/{owner}/issues/{number}/pull_requests `

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
| access_token  |   用户授权码  | query  | string |
| owner* |  仓库所属空间地址(企业、组织或个人的地址path)  | path   | string |
| repo*  |   仓库路径(path) | query  | string |
| number* | Issue 编号(区分大小写，无需添加 # 号)   | path   | string |


### 响应

```json
[
    {
        "id": 68525,
        "html_url": "https://test.gitcode.net/dengmengmian/test01/merge_requests/1",
        "number": 1,
        "state": "opened",
        "title": "new: 新增文件 1.text",
        "body": "new: 新增文件 1.text ",
        "created_at": "2024-04-15T16:29:11.035+08:00",
        "updated_at": "2024-04-15T16:29:17.769+08:00",
        "merged_at": null,
        "closed_at": null
    }
]
```

## 8 获取企业某个Issue所有标签

### 请求
`GET https://api.gitcode.com/api/v5/enterprises/{enterprise}/issues/{number}/labels`

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
| access_token  |   用户授权码  | query  | string |
| enterprise* |   企业名(path)  | path   | string |
| number*  |   Issue 编号(区分大小写，无需添加 # 号) | path  | string |
| page  | 当前的页码 | query | integer |
| per_page  | 每页的数量，最大为 100 | query | integer |

### 响应
```json
[
  {
    "color": "#008672",
    "name": "help wanted",
    "id": 381445,
    "title": "help wanted",
    "type": null,
    "textColor": "#FFFFFF"
  }
]
```
## 9 创建Issue标签
### 请求
`POST https://api.gitcode.com/api/v5/repos/{owner}/{repo}/issues/{number}/labels`

### 参数
| 参数名  | 描述                       | 类型  | 数据类型  |
|---------------|--------------------------|----------------|--------|
|  access_token* | 用户授权码                    | query | string |
|  owner* | 仓库所属空间地址(组织或个人的地址path)   | path | string |
|  repo*   | 仓库路径(path)               | path | string |
|  number*   | issue编号                  | path | string |
|  labels*   | 添加的标签 如: ["feat", "bug"] | body | array  |

### 响应
```json
[
  {
    "color": "#008672",
    "name": "help wanted",
    "id": 381445,
    "title": "help wanted",
    "type": null,
    "textColor": "#FFFFFF"
  }
]
```

## 10 删除Issue标签

### 请求
`DELETE https://api.gitcode.com/api/v5/repos/{owner}/{repo}/issues/{number}/labels/{name}`


### 参数
| 参数名           | 描述                       | 类型             | 数据类型   |
|---------------|--------------------------|----------------|--------|
|  access_token* | 用户授权码                    | query | string |
|  owner* | 仓库所属空间地址(组织或个人的地址path)   | path | string |
|  repo*   | 仓库路径(path)               | path | string |
|  number*   | 第几个PR，即本仓库PR的序数          | path | string |
|  name*   |  标签名称(批量删除用英文逗号分隔，如: bug,feature)    | path | string |

### 响应
`204`

## 11 创建Issue评论
### 请求
`POST https://api.gitcode.com/api/v5/repos/{owner}/{repo}/issues/{number}/comments`

### 参数
| 参数名  | 描述                       | 类型  | 数据类型  |
|---------------|--------------------------|----------------|--------|
|  access_token* | 用户授权码                    | query | string |
|  owner* | 仓库所属空间地址(组织或个人的地址path)   | path | string |
|  repo*   | 仓库路径(path)               | path | string |
|  number*   | issue编号                  | path | string |


### 响应
```json
{
    "id": 271624,
    "body": "评论内容。",
    "user": {
        "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
        "events_url": null,
        "followers_url": null,
        "following_url": null,
        "gists_url": null,
        "html_url": "https://test.gitcode.net/dengmengmian",
        "id": "661ce4eab470b1430d456154",
        "login": "dengmengmian",
        "member_role": null,
        "name": "麻凡_",
        "organizations_url": null,
        "received_events_url": null,
        "remark": null,
        "repos_url": null,
        "starred_url": null,
        "subscriptions_url": null,
        "type": null,
        "url": null
    },
    "target": {
        "issue": {
            "id": 152134,
            "title": "",
            "nubmer": 1
        }
    },
    "created_at": null,
    "updated_at": null
}
```
## 11 获取某个iusse下的操作日志
### 请求
`POST https://api.gitcode.com/api/v5/repos/{owner}/issues/{number}/operate_logs`

### 参数
| 参数名  | 描述                       | 类型  | 数据类型  |
|---------------|--------------------------|----------------|--------|
|  access_token* | 用户授权码                    | query | string |
|  owner* | 仓库所属空间地址(组织或个人的地址path)   | path | string |
|  repo*   | 仓库路径(path)               | path | string |
|  number*   | issue编号                  | path | string |

### 响应
```json
[
    {
        "id": 272199,
        "user": {
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
            "events_url": null,
            "followers_url": null,
            "following_url": null,
            "gists_url": null,
            "html_url": "https://test.gitcode.net/dengmengmian",
            "id": "661ce4eab470b1430d456154",
            "login": "dengmengmian",
            "member_role": null,
            "name": "麻凡_",
            "organizations_url": null,
            "received_events_url": null,
            "remark": null,
            "repos_url": null,
            "starred_url": null,
            "subscriptions_url": null,
            "type": null,
            "url": null
        },
        "content": "Create issue mr links: **new: 新增文件 1.text** #1",
        "created_at": "2024-04-20T15:20:24.009+08:00",
        "action_type": "add_issue_mr_link",
        "update_at": "2024-04-20T15:20:24.009+08:00",
        "title": "new: 新增文件 1.text",
        "body": "new: 新增文件 1.text ",
        "head": {
            "ref": "develop",
            "sha": "dd954d3a779edc86dae5b4b60c7f24dd0f195bf4",
            "repo": {
                "path": "test01",
                "name": "test01"
            },
            "assigner": {
                "login": "dengmengmian",
                "name": "麻凡_"
            }
        },
        "base": {
            "ref": "main",
            "sha": "32cff0d8faaa0c044d0f94957e656051986e8403",
            "repo": {
                "path": "test01",
                "name": "test01"
            },
            "assigner": null
        },
        "issue_id": "152642"
    },
    {
        "id": 272198,
        "user": {
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fa/fe/f32a9fecc53e890afbd48fd098b0f6c5f20f062581400c76c85e5baab3f0d5b2.png",
            "events_url": null,
            "followers_url": null,
            "following_url": null,
            "gists_url": null,
            "html_url": "https://test.gitcode.net/dengmengmian",
            "id": "661ce4eab470b1430d456154",
            "login": "dengmengmian",
            "member_role": null,
            "name": "麻凡_",
            "organizations_url": null,
            "received_events_url": null,
            "remark": null,
            "repos_url": null,
            "starred_url": null,
            "subscriptions_url": null,
            "type": null,
            "url": null
        },
        "content": "changed milestone to testew",
        "created_at": "2024-04-20T15:20:09.305+08:00",
        "action_type": "milestone",
        "update_at": "2024-04-20T15:20:09.305+08:00",
        "title": null,
        "body": null,
        "head": null,
        "base": null,
        "issue_id": "152642"
    }
]
```