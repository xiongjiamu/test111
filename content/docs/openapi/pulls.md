---
linkTitle: Pull Requests
title: PR接口文档
weight: 6
sidebar:
  open: false
---

## 1. 获取Pull Request列表

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | formData | string    |
|  state | 可选。Pull Request 状态 | query | string    |
|  base   | 可选。Pull Request 提交目标分支的名称 | query | string    |
|  since | 可选。起始的更新时间，要求时间格式为 ISO 8601 | quey | string    |
|  direction  | 可选。升序/降序 | query | string    |
|  milestone_number | 可选。里程碑序号(id) | formData | integer    |
|  page   | 当前的页码 | query | integer    |
|  per_page   | 每页的数量，最大为 100 | query | integer    |
|  author   | 可选。PR 创建者用户名 | query | string    |
|  assignee   |  可选。评审者用户名 | query | string    |
|  tester   |  可选。测试者用户名 | query | string    |

### 响应

```json
[
    {
        "id": 11217119,
        "html_url": "https://gitcode.com/OpenHarmony/ability_ability_base/pulls/341",
        "number": 341,
        "close_related_issue": 1,
        "prune_branch": false,
        "state": "open",
        "assignees_number": 1,
        "testers_number": 1,
        "assignees": [
            {
                "id": 7476845,
                "login": "ccllee",
                "name": "ccll",
                "avatar_url": null,
                "html_url": "https://gitcode.com/ccllee",
                "type": "User",
                "assignee": true,
                "code_owner": false,
                "accept": false
            }
        ],
        "milestone": null,
        "labels": [
            {
                "id": 113667295,
                "color": "20c22e",
                "name": "dco检查成功",
                "repository_id": 6511493,
                "created_at": "2021-06-18T15:08:35+08:00",
                "updated_at": "2024-04-08T21:46:16+08:00"
            }
        ],
        "created_at": "2024-04-08T19:54:38+08:00",
        "updated_at": "2024-04-08T21:48:05+08:00",
        "closed_at": null,
        "draft": false,
        "merged_at": null,
        "mergeable": true,
        "can_merge_check": false,
        "user": {
            "id": 9221401,
            "login": "hust-yf",
            "name": "yangfei",
            "avatar_url": null,
            "html_url": "https://gitcode.com/hust-yf",
            "type": "User"
        },
        "title": "新增configuration key支持sa对单个应用设置深浅色",
        "body": "**IssueNo**:Result**:",
        "head": {
			"label": "mas",
            "user": {
                "id": 9221401,
                "login": "hust-yf",
                "name": "yangfei",
                "avatar_url": null,
                "html_url": "https://gitcode.com/hust-yf",
                "type": "User"
            },
            "repo": {
                "id": 34079500,
                "full_name": "hjj-code_1/ability_ability_base",
                "human_name": "yffff/ability_ability_base",
                "path": "ability_ability_base",
                "name": "ability_ability_base",
                "assigner": {
                    "id": 9221401,
                    "login": "hust-yf",
                    "name": "yangfei",
                    "avatar_url": null,
                    "html_url": "https://gitee.com/hust-yf",
                    "type": "User"
                },
                "description": "暂无描述",
                "internal": false,
                "html_url": "https://gitcode.com/hjj-code_1/ability_ability_base.git"
            }
        },
        "base": {
            "label": "master",
            "user": {
                "id": 626123,
                "login": "landwind",
                "name": "mamingshuai",
                "avatar_url": null,
                "html_url": "https://gitcode.com/landwind",
                "type": "User"
            },
            "repo": {
                "id": 22974255,
                "full_name": "openharmony/ability_ability_base",
                "human_name": "OpenHarmony/ability_ability_base",
                "path": "ability_ability_base",
                "name": "ability_ability_base",
                "assigner": {
                    "id": 626123,
                    "login": "landwind",
                    "name": "mamingshuai",
                    "avatar_url": null,
                    "html_url": "https://gitcode.com/landwind",
                    "type": "User"
                },
                "description": "暂无描述",
                "internal": false,
                "html_url": "https://gitcode.com/openharmony/ability_ability_base.git",
            }
        }
    }
]

```
