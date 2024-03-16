---
linkTitle: 用户 API
title: 用户接口API文档
weight: 2
sidebar:
  open: false
---

## 1. 获取 SSH KEY 列表

### 请求

`GET /api/v4/user/keys`

### 参数

| 参数名       | 类型   | in|是否必选|default|描述                 |
| ------------ | ------ | ----|---|-------------------- | ------------ |
| page     | integer | query | false    | 1       | 分页参数，查询第page页                 |
| per_page | integer | query | false    | 20      | 指定每页查询的个数                     |
| search | string | query |false||模糊查询指定标题的ssh_key,大小写不敏感|

### 响应

```
[
    {
        "id": 4,
        "title": "demo1",,
        "key": "****",
        "created_at": "2023-07-26T14:55:11.655+08:00",
        "last_used_at": null
    },
    {
        "id": 10,
        "title": "demo2",
        "key": "****",
        "created_at": "2023-07-29T11:24:25.929+08:00",
        "last_used_at": null
    }
]
```
## 2. 添加 SSH KEY

### 请求

`POST /api/v4/user/keys`

### 参数

#### request body

``` json
{
  "key": "string", //必须为合法的SSH公钥格式
  "title": "string"
}
```
### 响应

```
{
    "id": 8992,
    "title": "demo3",
    "key": "****",
    "created_at": "2023-09-22T15:52:11.567+08:00",
    "last_used_at": null
}
```
## 3. 删除指定 SSH KEY

### 请求

`DELETE /api/v4/user/keys/{id}`

### 参数

| 参数名       | 类型   | in|是否必选|default|描述                 |
| ------------ | ------ | ----|---|-------------------- | ------------ |
| id     | integer | path | true    |       | 指定删除的SSH_KEY的id      |

### 响应

状态码204即为删除成功



