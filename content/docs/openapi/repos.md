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
                    "email": "13328943+gitcode_admin@user.noreply.gitcode.com"
                },
                "committer": {
                    "name": "Gitee",
                    "date": "2024-04-08T21:13:33+08:00",
                    "email": "noreply@gitcode.com"
                },
                "message": "wwqwqwq"
            }
        },
        "protected": true
    }
]

```

## 2. 获取一个组织的仓库

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

## 3 列出仓库所有的tags

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
            "email": "7543745+centking@user.noreply.gitcode.com",
            "date": "2024-04-14T06:18:54+00:00"
        }
    }
]
```

## 4 获取仓库具体路径下的内容
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

## 5. 获取用户的某个仓库
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
    "url": "https://gitcode.com/api/v5/repos/daming_1/zhu_di",
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
        "username": "CodeHub_beta_dev",
        "arts_id": "11",
        "nickname": "",
        "email": "CodeHub_beta_dev@huawei.com",
        "photo": "https://gitcode-user-img.obs.cn-north-4.myhuaweicloud.com:443/de%2Fdd%2F0de6d190dd41c0167427a5982dab156498b05573e7163ae5698d7f4fc0562ce0.png
    },    
}
```

## 6. 更新文件
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
        "url": "https://gitcode.com/api/v5/repos/daming_1/zhu_di/contents/Note2.md",
        "html_url": "https://gitcode.com/daming_1/zhu_di/blob/master/Note2.md",
        "download_url": "https://gitcode.com/daming_1/zhu_di/raw/master/Note2.md",
        "_links": {
            "self": "https://gitcode.com/api/v5/repos/daming_1/zhu_di/contents/Note2.md",
            "html": "https://gitcode.com/daming_1/zhu_di/blob/master/Note2.md"
        }
    }
}
```

## 7. 创建组织仓库
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
    "url": "https://gitcode.com/api/v5/repos/daming_1/test_create_project_2",
   
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

#### 8. 更新仓库设置
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
    "url": "https://gitcode.com/api/v5/repos/daming_1/test_create_project_2",
    "path": "test_create_project_2",
}
```

#### 9. 仓库的某个提交
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
    "url": "https://gitcode.com/api/v5/repos/daming_1/zhu_di/commits/7ffc0a0deb709143f6be12a55e218fab9233ca37",
    "sha": "7ffc0a0deb709143f6be12a55e218fab9233ca37",
    "html_url": "https://gitcode.com/daming_1/zhu_di/commit/7ffc0a0deb709143f6be12a55e218fab9233ca37",
    "comments_url": "https://gitcode.com/api/v5/repos/daming_1/zhu_di/commits/7ffc0a0deb709143f6be12a55e218fab9233ca37/comments",
    "commit": {
        "author": {
            "name": "占分",
            "date": "2024-04-14T07:25:11+00:00",
            "email": "7543745+centking@user.noreply.gitcode.com"
        },
        "committer": {
            "name": "Gitee",
            "date": "2024-04-14T07:25:11+00:00",
            "email": "noreply@gitcode.com"
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

#### 10. 获取文件Blob
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
    "sha":"e5699fe1b360d6c799ee58b24fb5a670b1e14851",
    "size":19,
    "url":"https://gitcode.com/api/v5/repos/daming_1/zhu_di/git/blobs/e5699fe1b360d6c799ee58b24fb5a670b1e14851","content":"JXU2RDRCJXU4QkQ1d2ViaG9vaw==",
    "encoding":"base64"
}
```



## 11. 创建一个仓库的Tag

### 请求

`POST https://api.gitcode.com/api/v5/repos/{owner}/{repo}/tags`

### 参数

| 参数名  | 描述  | 类型       | 数据类型  |
| ------ | ------ |----------|------|
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path     | string    |
|  repo*   | 仓库路径(path) | path     | string    |
|  refs*   | 起点名称，默认main | body     | string    |
|  tag_name*   | 新创建的标签名称 | body | string    |
|  tag_message   | Tag 描述,默认为空 | body | string    |
|  access_token* | 用户授权码 | formData | string    | 



### 响应
```json
{
  "name": "tag2",
  "message": "",
  "commit": {
    "sha": "5d165dae3b073d3e92ca91c3edcb21994361462c",
    "date": "2024-04-08T13:13:33+00:00"
  },
  "tagger": null
}
```
## 12. 创建commit评论
### 请求
`POST https://api.gitcode.com/api/v5/repos/{owner}/{repo}/commits/{sha}/comments`

### 参数

| 参数名  | 描述  | 类型       | 数据类型  |
| ------ | ------ |----------|------|
|  access_token* | 用户授权码 | formData | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path     | string    |
|  repo*   | 仓库路径(path) | path     | string    |
|  number*   | 第几个PR，即本仓库PR的序数 | path     | string    |
|  sha*   | 第几个PR，即本仓库PR的序数 | path     | string    |
|  body*   | 评论内容 | body     | string    |
|  path   | 文件的相对路径 | body     | string    |
|  position  | Diff的相对行数 | body     | integer    |

### 响应
```json
{
    "id":111,
    "body":"content"
}
```

#### 13. 新建文件
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
            "email": "13328943+gitcode_admin@user.noreply.gitcode.com",
            "date": "2024-04-11T09:15:20+00:00"
        },
        "committer": {
            "name": "Gitee",
            "email": "noreply@gitcode.com",
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



#### 14. 仓库归档

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

#### 15. 转移仓

#### 请求：

POST `https://api.gitcode.com/api/v5/orgs/{org}/projects/{repo}/transfer`

| 参数名        | 描述               | 类型 | 数据类型 |
| ------------- | ------------------ | ---- | -------- |
| access_token* | 用户授权码         | body | string   |
| org*          | 仓库所属组织       | path | string   |
| repo*         | 仓库路径           | path | string   |
| transfer_to*  | 要转移到的目标组织 | body | string   |
| password*     | 用户密码           | body | string   |

#### 返回:

```json
{
    "code": 1,
    "msg": "success"
}
```

## 16. 获取仓库目录Tree

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/git/trees/{sha}`

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  sha*   | 可以是分支名(如master)、Commit或者目录Tree的SHA | path | string    |
|recursive | 赋值为1递归获取目录 | query |integer|


### 响应

```json
{
    "tree": [
        {
            "sha": "5259c4b24f015ffdc3663e81837a730c2a108e1f",
            "name": "b",
            "type": "tree",
            "path": "a/b",
            "mode": "040000",
            "md5": "a7e86136543b019d72468ceebf71fb8e"
        }
    ]
}
```


## 17. 添加仓库成员或更新仓库成员权限

### 请求

`PUT https://api.gitcode.com/api/v5/repos/{owner}/{repo}/collaborators/{username}`

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | formData | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  username*   | 用户名(username/login)  | path | string    |
| permission* | 成员权限: 拉代码(pull)，推代码(push)，管理员(admin)。默认: push | formData |string |


### 响应

```json
{
    "id": 7543745,
    "login": "centking",
    "name": "占分",
    "avatar_url": null,
    "html_url": "https://gitcode.com/centking",
    "remark": "",
    "type": "User",
    "permissions": {
        "pull": true,
        "push": true,
        "admin": false
    }
}
```

## 18. 移除仓库成员

### 请求

`DELETE https://api.gitcode.com/api/v5/repos/{owner}/{repo}/collaborators/{username}`

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  username*   | 用户名(username/login)  | path | string    |


### 响应

```
HTTP status 204 No Content
```

## 19. 修改代码审查设置

### 请求

`PUT /api/v5/repos/{owner}/{repo}/reviewer`

### 参数

| 参数名            | 描述                                                         | 类型  | 数据类型 |
| ----------------- | ------------------------------------------------------------ | ----- | -------- |
| access_token*     | 用户授权码                                                   | query | string   |
| owner*            | 仓库所属空间地址(组织或个人的地址path)                       | path  | string   |
| repo*             | 仓库路径(path)                                               | path  | string   |
| assignees*        | 审查人员username，可多个，半角逗号分隔，如：(username1,username2) | body  | string   |
| testers*          | 测试人员username，可多个，半角逗号分隔，如：(username1,username2) | body  | string   |
| assignees_number* | 最少审查人数                                                 | body  | integer  |
| testers_number*   | 最少测试人数                                                 | body  | integer  |


### 响应

```json
{
    "id": 7543745,
    "created_at": "2024-01-24T14:33:44+08:00",
    "updated_at": "2024-04-07T21:23:08+08:00"
}
```

## 20. 新建保护分支规则

### 请求

`PUT /api/v5/repos/{owner}/{repo}/branches/setting/new`

### 参数

| 参数名        | 描述                                                         | 类型  | 数据类型 |
| ------------- | ------------------------------------------------------------ | ----- | -------- |
| access_token* | 用户授权码                                                   | query | string   |
| owner*        | 仓库所属空间地址(组织或个人的地址path)                       | path  | string   |
| repo*         | 仓库路径(path)                                               | path  | string   |
| wildcard*     | 分支/通配符                                                  | body  | string   |
| pusher*       | 可推送代码成员。developer：仓库管理员和开发者；admin：仓库管理员；none：禁止任何人推送 | body  | string   |
| merger*       | 可合并 Pull Request 成员。developer：仓库管理员和开发者；admin：仓库管理员；none：禁止任何人合并 | body  | string   |


### 响应

```text
HTTP status 200 No Content
```

