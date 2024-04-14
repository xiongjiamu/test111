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

## 6.4 列出仓库所有的tags

#### 请求：
`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/tags`

#### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | formData | string    |
|  page   | 当前的页码 | query | integer    |
|  per_page   | 每页的数量，最大为 100 | query | integer    |

### 响应
```json
[
    {
        "name": "v1.0",
        "message": "111",
        "commit": {
            "sha": "3e43581d16bc456802a1fee673b9a2a9b9618f0f",
            "date": "2024-04-14T02:59:22+00:00"
        },
        "tagger": {
            "name": "占分",
            "email": "7543745+centking@user.noreply.gitee.com",
            "date": "2024-04-14T06:18:54+00:00"
        }
    }
]
```

## 6.6 获取仓库具体路径下的内容
#### 请求：
`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/contents(/{path})`

#### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  path*  | 文件的路径 | path | string |
| ref | 分支、tag或commit。默认: 仓库的默认分支(main)  | query | string |

### 响应
```json
{
    "type": "file",
    "encoding": "base64",
    "size": 19,
    "name": "Note2.md",
    "path": "Note2.md",
    "content": "JXU2RDRCJXU4QkQ1d2ViaG9vaw==",
    "sha": "e5699fe1b360d6c799ee58b24fb5a670b1e14851",
    "url": "https://gitcode.com/api/v5/repos/daming_1/zhu_di/contents/Note2.md",
    "html_url": "https://gitcode.com/daming_1/zhu_di/blob/master/Note2.md",
    "download_url": "https://gitcode.com/daming_1/zhu_di/raw/master/Note2.md",
    "_links": {
        "self": "https://gitcode.com/api/v5/repos/daming_1/zhu_di/contents/Note2.md",
        "html": "https://gitcode.com/daming_1/zhu_di/blob/master/Note2.md"
    }
}
```

## 6.7 获取用户的某个仓库
#### 请求：
`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}`

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | formData | string    |

#### 响应
```json
{
    "id": 34165823,
    "full_name": "daming_1/zhu_di",
    "human_name": "daming/zhu_di",
    "url": "https://gitee.com/api/v5/repos/daming_1/zhu_di",
    "namespace": {
        "id": 13386953,
        "name": "daming",
        "path": "daming_1",
        "html_url": "https://gicode.com/daming_1"
    },
    "path": "zhu_di",
    "name": "zhu_di",
    "description": "朱棣",
    "private": false,
    "public": true,
    "status": "开始",
    "ssh_url_to_repo": "ssh://git@gitcode.com/group11111/repocc12.git",
    "http_url_to_repo": "https://gitcode.com/group11111/repocc12.git",
    "web_url": "https://gitcode.com/group11111/repocc12",
    "readme_url": "https://gitcode.com/group11111/repocc12/blob/main/README.md",
    "created_at": "2023-09-15T16:45:09.502+08:00",
    "updated_at": "2023-09-15T16:45:09.502+08:00",
    "creator": {
        "id": 80,
        "name": "CodeHub_beta_dev",
        "username": "CodeHub_beta_dev",
        "iam_id": "c369c68f1ff84679b5a8ed904d8bff1c",
        "nick_name": "",
        "state": "active",
        "avatar_url": "1111",
        "avatar_path": null,
        "email": "CodeHub_beta_dev@huawei.com",
        "name_cn": "CodeHub_beta_dev",
        "web_url": "https://gitcode.com/CodeHub_beta_dev",
        "tenant_name": null
    },    
}
```

## 6.9 更新文件
#### 请求：
`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/contents/{path}`

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | formData | string    |
|   path*   |   文件的路径  |   path    |  string |
|   content* |  文件内容, 要用 base64 编码 | formData | string |
|   sha* |  文件的 Blob SHA | formData | string |
| message* | 提交信息 | formData | string |

```json
{
    "content": {
        "name": "Note2.md",
        "path": "Note2.md",
        "size": 19,
        "sha": "e5699fe1b360d6c799ee58b24fb5a670b1e14851",
        "url": "https://gitee.com/api/v5/repos/daming_1/zhu_di/contents/Note2.md",
        "html_url": "https://gitee.com/daming_1/zhu_di/blob/master/Note2.md",
        "download_url": "https://gitee.com/daming_1/zhu_di/raw/master/Note2.md",
        "_links": {
            "self": "https://gitee.com/api/v5/repos/daming_1/zhu_di/contents/Note2.md",
            "html": "https://gitee.com/daming_1/zhu_di/blob/master/Note2.md"
        }
    }
}
```

## 6.9 创建组织仓库
#### 请求：
`GET https://api.gitcode.com/api/v5/orgs/{org}/repos`

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string   | 
|  name*   | 仓库名称 | formData | string   |
|  org* | 组织的路径(path/login) | path |   string |
| description	| 仓库描述 | formData |string |

```json
{
    "id": 34171993,
    "full_name": "daming_1/test_create_project_2",
    "human_name": "daming/test_create_project_2",
    "url": "https://gitee.com/api/v5/repos/daming_1/test_create_project_2",
   
    "path": "test_create_project_2",
    "name": "test_create_project_2",
    
    "description": "描述",
    "private": false,
    "public": true,
    "namespace": {
        "id": 74962,
        "name": "group1111",
        "path": "group11111",
        "develop_mode": "normal",
        "region": null,
        "cell": "default",
        "kind": "group",
        "full_path": "group11111",
        "full_name": "group1111",
        "parent_id": null,
        "visibility_level": 20,
        "enable_file_control": false,
        "owner_id": null
    },
    "empty_repo": null,
    "starred": null,
    "visibility": "public",
    "owner": null,
    "creator": null,
    "creator_id": null,
    "forked_from_project": null,
    "item_type": null,
    "main_repository_language": null
}
```

#### 6.11 更新仓库设置
#### 请求：
`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}`

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string   | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | formData | string    |
| name* | 仓库名称 | formData | string |
| description | 仓库描述 | formData | string |
| homepage	|   主页(eg: https://gitcode.com) | formData |string |
| path	 |  更新仓库路径 |  formData    |   string |

#### 返回:
```json
{
    "id": 34171993,
    "full_name": "daming_1/test_create_project_2",
    "human_name": "daming/test_create_project_2",
    "url": "https://gitee.com/api/v5/repos/daming_1/test_create_project_2",
    "path": "test_create_project_2",
}
```

#### 6.20. 仓库的某个提交
#### 请求：
`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/commits/{sha}`

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string   | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | formData | string    |
| sha* |  提交的SHA值或者分支名 | path | string |

#### 返回:
```json
{
    "url": "https://gitee.com/api/v5/repos/daming_1/zhu_di/commits/7ffc0a0deb709143f6be12a55e218fab9233ca37",
    "sha": "7ffc0a0deb709143f6be12a55e218fab9233ca37",
    "html_url": "https://gitee.com/daming_1/zhu_di/commit/7ffc0a0deb709143f6be12a55e218fab9233ca37",
    "comments_url": "https://gitee.com/api/v5/repos/daming_1/zhu_di/commits/7ffc0a0deb709143f6be12a55e218fab9233ca37/comments",
    "commit": {
        "author": {
            "name": "占分",
            "date": "2024-04-14T07:25:11+00:00",
            "email": "7543745+centking@user.noreply.gitee.com"
        },
        "committer": {
            "name": "Gitee",
            "date": "2024-04-14T07:25:11+00:00",
            "email": "noreply@gitee.com"
        },
        "message": "提交信息测试",
    },
    "stats": {  
        "additions": 1,
        "deletions": 1,
        "total": 2
    }
}
```

#### 6.20. 新建文件
#### 请求：
POST `https://api.gitcode.com/api/v5/repos/{owner}/{repo}/contents/{path}`

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | body | string   |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
| path* | 文件路径 | path | string |
| content* | 文件内容, 要用 base64 编码 | body | string |
| message* | 提交的 commit 信息 | body | string |
| branch | 提交的分支 | body | string |

#### 返回:
```json
{
    "commit": {
        "sha": "668cb104692b30d537b07f3721df9956d073d343",
        "author": {
            "name": "GitCode2023",
            "email": "13328943+gitcode_admin@user.noreply.gitee.com",
            "date": "2024-04-11T09:15:20+00:00"
        },
        "committer": {
            "name": "Gitee",
            "email": "noreply@gitee.com",
            "date": "2024-04-11T09:15:20+00:00"
        },
        "message": "22222"
        "parents": [
            {
                    "sha":
      "0117aa5c6bc8e33d18ad8911afa3cbb54a1faabe"
            }
        ]
    }
}
```



#### 6.21. 仓库归档

#### 请求：

PUT `https://api.gitcode.com/api/v5/org/{org}/repo/{repo}/status`

| 参数名        | 描述                       | 类型 | 数据类型 |
| ------------- | -------------------------- | ---- | -------- |
| access_token* | 用户授权码                 | body | string   |
| org*          | 仓库所属组织               | path | string   |
| repo*         | 仓库路径(path)             | path | string   |
| status*       | 仓库状态，0：开始，2：关闭 | body | integer  |
| password*     | 用户密码                   | body | string   |

#### 返回:

```json
{
    "code": 1,
    "msg": "success"
}
```

#### 6.21. 转移仓

#### 请求：

POST `https://api.gitcode.com/api/v5/org/{org}/repo/{owner}/transfer`

| 参数名        | 描述               | 类型 | 数据类型 |
| ------------- | ------------------ | ---- | -------- |
| access_token* | 用户授权码         | body | string   |
| org*          | 仓库所属组织       | path | string   |
| owner*        | 仓库owner          | path | string   |
| transfer_to*  | 要转移到的目标组织 | body | integer  |
| password*     | 用户密码           | body | string   |

#### 返回:

```json
{
    "code": 1,
    "msg": "success"
}
```

