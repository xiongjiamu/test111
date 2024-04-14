---
linkTitle: OpenAPI Guide
title: OpenAPI 使用入门
weight: 1
sidebar:
  open: false
---

GitCode API 提供了强大的功能，允许开发者通过编程方式与 GitCode 上的资源进行交互。本文档将向您介绍如何使用 GitCode API，包括有效的 API 请求、认证方式、状态码、分页、路径参数等方面的内容。

## 有效的 API 请求

对 GitCode API 发起请求必须包括 `api` 和 API 版本，当前可用的版本为 `/api/v5`。

以下是对 GitCode API 发起请求的基本示例：

```bash
curl "https://api.gitcode.com/api/v5/users/{username}"
```

## 认证

大多数 GitCode API 请求需要认证，或者在未提供认证时只返回公共数据。每个接口的文档都会指明是否需要认证。

你可以通过[个人访问令牌](https://gitcode.com/setting/token-classic)使用 GitCode API 进行认证。

如果认证信息无效或缺失，GitCode 会返回带有 `401` 状态码的错误信息：

```json
{
  "message": "401 Unauthorized"
}
```

### 个人访问令牌

你可以通过 `access_token` 头中传递个人访问令牌来使用 API 进行认证。

请求头中的个人访问令牌的示例：

```bash
curl  "https://api.gitcode.com/api/v5/users/{username}?access_token=xxxxx"
```

## 状态码

GitCode API 旨在根据上下文和操作返回不同的状态码。这样，如果请求导致错误，你可以了解出了什么问题。

下表概述了 GitCode API 功能的一般行为：

| 请求类型 | 描述 |
|---|---|
| `GET` | 访问一个或多个资源并以 JSON 形式返回结果 |
| `POST` | 如果资源成功创建，则返回 `201 Created` 并以 JSON 形式返回新创建的资源 |
| `GET` / `PUT` | 如果资源被成功访问或修改，则返回 `200 OK`。返回的（修改后的）结果以 JSON 形式返回 |
| `DELETE` | 如果资源成功删除，则返回 `204 No Content`；如果资源计划被删除，则返回 `202 Accepted` |

下表显示了 GitCode API 请求可能的返回代码：

| 返回值 | 描述 |
|---|---|
| `200 OK` | `GET`、`PUT` 或 `DELETE` 请求成功，并且资源本身以 JSON 形式返回 |
| `201 Created` | `POST` 请求成功，并且资源以 JSON 形式返回 |
| `202 Accepted` | `GET`、`PUT` 或 `DELETE` 请求成功，并且资源计划进行处理 |
| `204 No Content` | 服务器已成功满足请求，并且在响应负载体中没有额外的内容发送 |
| `301 Moved Permanently` | 资源已被定位到由 `Location` 头给出的 URL |
| `304 Not Modified` | 资源自上次请求以来未被修改 |
| `400 Bad Request` | API 请求的必需属性缺失。例如，未给出问题的标题 |
| `401 Unauthorized` | 用户未经认证。需要有效的用户令牌 |
| `403 Forbidden` | 请求不被允许。例如，用户不允许删除项目 |
| `404 Not Found` | 无法访问资源。例如，无法找到资源的 ID，或者用户无权访问资源 |
| `405 Method Not Allowed` | 不支持请求 |
| `409 Conflict` | 冲突的资源已存在。例如，创建已存在名称的项目 |
| `412 Precondition Failed` | 请求被拒绝。这可能发生在尝试删除在此期间被修改的资源时，如果提供了 `If-Unmodified-Since` 头 |
| `422 Unprocessable` | 无法处理实体 |
| `429 Too Many Requests` | 用户超过了应用程序[速率限制](https://docs.gitcode.com/docs/api/guide/#速率限制) |
| `500 Server Error` | 处理请求时，服务器上出了问题 |
| `503 Service Unavailable` | 服务器无法处理请求，因为服务器暂时过载 |
