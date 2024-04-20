---
linkTitle: 组织 Organizations
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

## 2. 获取组织成员详情

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
