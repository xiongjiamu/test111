---
linkTitle: Oauth2.0
title: OAuth 文档
weight: 3
sidebar:
  open: false
---

## 1.认证接口

### 请求

`GET https://gitcode.com/oauth/authorize?client_id={client_id}&redirect_uri={redirect_uri}&response_type=code&scope={scope}&state={state}`

| 参数名           | 描述                       | 类型    | 数据类型   |
|---------------|--------------------------|-------|--------|
| client_id*    | 注册的客户端 ID                | query | string | 
| redirect_uri* | 授权后的url                  | query | string |
| scope        | 权限范围                     | query | string |
| state         | 随机字符串,<br/>用于防止跨站点请求伪造攻击 | query | string |

## 2.重定向

如果用户接受你的授权，GitCode 会重定向回您的网站，携带code参数以及你在上一步中在参数中提供的状态state。如果状态不匹配，则说明是第三方创建了请求，需要中止该过程。

`GET {redirect_uri}?code={code}&state={state}`

## 3.获取授权token

`POST https://gitcode.com/oauth/token?grant_type=authorization_code&code={code}&client_id={client_id}&client_secret={client_secret}`

| 参数名            | 描述        | 类型    | 数据类型               |
|----------------|-----------|-------|--------------------|
| grant_type     | 授权码模式     | query | authorization_code | 
| code*          | 授权码       | query | string             | 
| client_id*     | 注册的客户端 ID | query | string             | 
| client_secret* | 注册的客户端密钥  | form-data | string             |

### 响应：
```json
{
    "access_token": "eyPZPVNfsibj9tap_ibj3t3p",
    "expires_in": 1296000,
    "refresh_token": "b77ced3aee884348852160deab3697a1",
    "scope": "all_user all_key all_groups all_projects all_pr all_issue all_note all_hook all_repository",
    "created_at": "2024-04-20T09:07:59.889Z"
}
```

## 4. 使用访问令牌访问用户信息API
```text
Authorization: Bearer {access_token}
GET https://api.gitcode.com/api/v5/user
```

### 5.刷新access_token
`POST https://gitcode.com/oauth/token?grant_type=refresh_token&refresh_token={refresh_token}`
