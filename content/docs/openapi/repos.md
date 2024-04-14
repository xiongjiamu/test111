---
linkTitle: 仓库Repositories
title: 仓库接口文档
weight: 7
sidebar:
  open: false
---

## 1. 获取仓库所有分支

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/branches`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | formData | string    |
|  sort | 排序字段 name/updated | query | string    |
|  base   | 可选。Pull Request 提交目标分支的名称 | query | string    |
|  direction | 排序方向 asc/desc | quey | string    |
|  page   | 当前的页码 | query | integer    |
|  per_page   | 每页的数量，最大为 100 | query | integer    |

### 响应

```json
[
    {
        "name": "master",
        "commit": {
            "sha": "5d165dae3b073d3e92ca91c3edcb21994361462c",
            "commit": {
                "author": {
                    "name": "GitCode2023",
                    "date": "2024-04-08T21:13:33+08:00",
                    "email": "13328943+gitcode_admin@user.noreply.gitee.com"
                },
                "committer": {
                    "name": "Gitee",
                    "date": "2024-04-08T21:13:33+08:00",
                    "email": "noreply@gitee.com"
                },
                "message": "wwqwqwq"
            }
        },
        "protected": true
    }
]

```

## 5. 获取一个组织的仓库

### 请求

`GET https://api.gitcode.com/api/v5/orgs/{org}/repos`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  org* | 组织的路径(path/login) | path | string    |
|  type	| 筛选仓库的类型，可以是 all, public, private。默认: all | query | string    |
|  page   | 当前的页码 | query | integer    |
|  per_page   | 每页的数量，最大为 100 | query | integer    |

### 响应

```json
[
    {
        "id": 29724198,
        "full_name": "openharmony/.gitee",
		"namespace": {
            "id": 6486504,
            "type": "group",
            "name": "OpenHarmony",
            "path": "openharmony",
            "html_url": "https://gitcode.com/openharmony"
        },
        "path": ".gitee",
        "name": ".gitee",
        "description": "",
        "private": false,
        "public": true,
        "internal": false,
        "fork": false,
        "html_url": "https://gitcode.com/openharmony/.gitee.git",
        "forks_count": 4,
        "stargazers_count": 0,
        "watchers_count": 1,
        "default_branch": "master",
        "open_issues_count": 0,
        "license": null,
        "project_creator": "landwind",
        "pushed_at": "2024-02-06T18:25:26+08:00",
        "created_at": "2023-06-16T10:55:42+08:00",
        "updated_at": "2024-03-29T14:59:46+08:00",
        "status": "开始"
    }
]

```