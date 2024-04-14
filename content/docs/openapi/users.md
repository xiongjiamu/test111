---
linkTitle: 用户账号 Users
title: 用户账号API文档
weight: 2
sidebar:
  open: false
---

## 1. 获取一个用户

### 请求

`GET https://api.gitcode.com/api/v5/users/{username}`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
| username* | 用户名(username) | path | string    |

### 响应

```
{
    "avatar_url": null,
    "events_url": null,
    "followers_url": null,
    "following_url": null,
    "gists_url": null,
    "html_url": "https://gitcode.com/theo6789",
    "id": "64dc3b13b8b9504cec223ab1",
    "login": "theo6789",
    "member_role": null,
    "name": "TheoCui",
    "organizations_url": null,
    "received_events_url": null,
    "remark": null,
    "repos_url": null,
    "starred_url": null,
    "subscriptions_url": null,
    "type": "User",
    "url": null,
    "bio": "",
    "blog": "",
    "company": "",
    "created_at": null,
    "email": "cuizk@csdn.net",
    "followers": 0,
    "following": 1,
    "profession": null,
    "public_gists": null,
    "public_repos": null,
    "qq": null,
    "stared": null,
    "updated_at": null,
    "watched": null,
    "wechat": null,
    "weibo": null
}
```


## 2. 获取授权用户的资料

### 请求

`GET https://api.gitcode.com/api/v5/user`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 

### 响应

```
{
    "avatar_url": null,
    "events_url": null,
    "followers_url": null,
    "following_url": null,
    "gists_url": null,
    "html_url": "https://gitcode.com/yinlin",
    "id": "6601148c5079ba0d1c00fc0e",
    "login": "yinlin",
    "member_role": null,
    "name": "yinlin-昵称",
    "organizations_url": null,
    "received_events_url": null,
    "remark": null,
    "repos_url": null,
    "starred_url": null,
    "subscriptions_url": null,
    "type": "User",
    "url": null,
    "bio": "",
    "blog": "",
    "company": "",
    "created_at": null,
    "email": "",
    "followers": 0,
    "following": 0,
    "profession": null,
    "public_gists": null,
    "public_repos": null,
    "qq": null,
    "stared": null,
    "updated_at": null,
    "watched": null,
    "wechat": null,
    "weibo": null
}
```