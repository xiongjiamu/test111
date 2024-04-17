---
linkTitle: Oauth2.0
title: OAuth文档
weight: 3
sidebar:
  open: false
---

# OAuth 获取认证步骤

## A.授权码模式

### 1.认证接口

#### 请求

`GET https://gitcode.com/oauth/authorize?client_id={client_id}&client_secret={client_secret}&redirect_uri={redirect_uri}&response_type=code&scope={scope}&state={state}`

| 参数名           | 描述                       | 类型    | 数据类型   |
|---------------|--------------------------|-------|--------|
| client_id*    | 注册的客户端 ID                | query | string | 
| client_secret* | 注册的客户端密钥  | query | string             |
| redirect_uri* | 授权后的url                  | query | string |
| scope*        | 权限范围                     | query | string |
| state         | 随机字符串,<br/>用于防止跨站点请求伪造攻击 | query | string |

### 2.重定向

如果用户接受你的授权，GitCode 会重定向回您的网站，携带code参数以及你在上一步中在参数中提供的状态state。如果状态不匹配，则说明是第三方创建了请求，需要中止该过程。

`GET {redirect_uri}?code={code}&state={state}`

### 3.获取授权token

`POST https://web-api.gitcode.com/uc/api/v1/oauth/token?grant_type=authorization_code&code={code}&client_id={client_id}&client_secret={client_secret}`

| 参数名            | 描述        | 类型    | 数据类型               |
|----------------|-----------|-------|--------------------|
| grant_type     | 授权码模式     | query | authorization_code | 
| code*          | 授权码       | query | string             | 
| client_id*     | 注册的客户端 ID | query | string             | 
| client_secret* | 注册的客户端密钥  | form-data | string             |

#### 响应：
```json
{
  "access_token": "xxxx",
  "token_type": "",
  "expires_in": "",
  "refresh_token": "xxx",
  "scope": "userinfo",
  "created_at": ""
}
```

### 4. 使用访问令牌访问用户信息API
```text
Authorization: Bearer {access_token}
GET https://api.gitcode.com/api/v5/user
```
