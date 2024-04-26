---
linkTitle: Organizations
title: 组织接口文档
weight: 4
sidebar:
  open: false
---

## 1. 列出用户所属的组织

### 请求

`GET https://api.gitcode.com/api/v5/users/{username}/orgs`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  username* | 用户名(username/login) | path | string    |
|  page | 当前的页码 | query | string    |
|  per_page | 每页的数量，默认为20，最大为100 | query | string    |

### 响应

```json
[
    {
        "id": 133039,
        "login": "openharmony",
        "name": "OpenHarmony",
        "avatar_url": null,
        "repos_url": null,
        "events_url": null,
        "members_url": null,
        "description": "OpenHarmony是由开放原子开源基金会（OpenAtom Foundation）孵化及运营的开源项目，目标是面向全场景、全连接、全智能时代，搭建一个智能终端设备操作系统的框架和平台，促进万物互联产业的繁荣发展。",
        "follow_count": 3
    }
]
```

## 2. 列出授权用户所属的企业

### 请求

`GET https://api.gitcode.com/api/v5/users/enterprises`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  page | 当前的页码 | query | string    |
|  per_page | 每页的数量，默认为20，最大为100 | query | string    |

### 响应

```json
[
    {
        "id": 133039,
        "login": "openharmony",
        "path": "openharmony",
        "name": "OpenHarmony",
        "avatar_url": null,
        "repos_url": null,
        "events_url": null,
        "members_url": null,
        "description": "OpenHarmony是由开放原子开源基金会（OpenAtom Foundation）孵化及运营的开源项目，目标是面向全场景、全连接、全智能时代，搭建一个智能终端设备操作系统的框架和平台，促进万物互联产业的繁荣发展。",
        "follow_count": 3
    }
]
```

## 3. 获取组织成员详情

### 请求

`GET https://api.gitcode.com/api/v5/orgs/{org}/members/{username}`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  org* | 组织的路径(path/login) | path | string    |
|  username* | 用户名(username/login) | path | string    |

### 响应

```json
{
    "id": 133039,
    "path":"openharmony",
    "name":"",
    "url":"",
    "avatar_url": null,
    "user":{
        "id":"64dc3b13b8b9504cec223ab1",
        "login":"theo6789",
        "name":"TheoCui",
        "avatar_url": null,
        "html_url": "https://gitcode.com/theo6789",
    }
}
```

## 4. 获取一个企业

### 请求

`GET https://api.gitcode.com/api/v5/enterprises/{enterprise}`

### 参数

| 参数名  | 描述                  | 类型  | 数据类型  |
| ------ |---------------------| ------  |------|
|  access_token* | 用户授权码               | query | string    | 
|  enterprise* | 企业的路径(path/login)   | path | string    |

### 响应

```json
{
  "path": "wz_test_4",
  "avatar_url": "",
  "created_at": "2024-04-12T15:45:58.249+08:00",
  "description": "",
  "email": null,
  "enterprise": "wz_test_4",
  "events_url": null,
  "follow_count": 0,
  "html_url": null,
  "id": 144239,
  "location": null,
  "login": "wz_test_4",
  "members": 1,
  "members_url": null,
  "name": "wz_test_4",
  "owner": {
    "id": 0,
    "login": null,
    "name": null,
    "avatar_url": null,
    "url": null,
    "html_url": null
  },
  "remark": null,
  "private_repos": 0,
  "public": true,
  "public_repos": 0,
  "repos_url": null,
  "type": "Group",
  "url": null
}
```
## 3 获取一个组织信息
### 请求
`GET https://api.gitcode.com/api/v5/orgs/{org}`

### 参数
| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    |
|  org* | 组织的路径(path/login) | path | string    |

### 响应
```json
{
   "id": 6486504,
   "login": "openharmony",
   "name": "OpenHarmony",
   "avatar_url": "",
   "repos_url": "https://api.gitcode.com/openharmony/repos",
   "events_url": "https://api.gitcode.com/openharmony/events",
   "members_url": "https://api.gitcode.com/openharmony/members{/member}",
   "description": "OpenHarmony是由开放原子开源基金会（OpenAtom Foundation）孵化及运营的开源项目，目标是面向全场景、全连接、全智能时代，搭建一个智能终端设备操作系统的框架和平台，促进万物互联产业的繁荣发展。\r\n",
   "enterprise": "openharmony",
   "follow_count": 40819
}
```
## 4 企业 Pull Request 列表
### 请求
`GET https://api.gitcode.com/api/v5/enterprise/{enterprise}/pull_requests`

### 参数
| 参数名           | 描述  | 类型  | 数据类型  |
|---------------| ------ | ------  |------|
| access_token* | 用户授权码 | query | string    |
| enterprise*   | 企业的路径(path/login) | path | string    |
| repo          |可选。仓库路径(path) | query | string    |
| state         |可选。Pull Request 状态 | query | string    |
| sort        |可选。排序字段，默认按创建时间 | query | string    |
| direction        |可选。升序/降序 | query | string    |
| page        |当前的页码 | query | string    |
| per_page        |每页的数量，最大为 100 | query | string    |

### 响应
```json
[
  {
    "id": 71020,
    "url": "https://test.gitcode.net/api/v5/repos/test/test/1",
    "html_url": "https://test.gitcode.net/test/test/1",
    "number": 1,
    "state": "merged",
    "assignees_number": 0,
    "testers_number": 0,
    "assignees": [],
    "testers": [],
    "mergeable": null,
    "can_merge_check": true,
    "head": {
      "ref": "main",
      "sha": "d874402d259744a00121c2cff0febc8554339aef",
      "repo": {
        "path": "test",
        "name_space": {
          "path": "repo-dev"
        },
        "assigner": {
          "id": 708,
          "login": "Lzm_0916",
          "name": "Lzm_0916"
        }
      }
    },
    "base": {
      "ref": null,
      "sha": null,
      "repo": {
        "path": "test",
        "name_space": {
          "path": "repo-dev"
        },
        "assigner": {
          "id": 708,
          "login": "Lzm_0916",
          "name": "Lzm_0916"
        }
      }
    }
  }
]
```

