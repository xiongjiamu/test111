---
linkTitle: OpenAPI Guide
title: OpenAPI 使用入门
weight: 1
sidebar:
  open: false
---

GitCode API 提供了强大的功能，允许开发者通过编程方式与 GitCode 上的资源进行交互。本文档将向您介绍如何使用 GitCode API，包括有效的 API 请求、认证方式、状态码、分页、路径参数等方面的内容。

## 有效的 API 请求

对 GitCode API 发起请求必须包括 `api` 和 API 版本，当前可用的版本为 `/api/v4`。

以下是对 GitCode API 发起请求的基本示例：

```bash
curl "https://api.gitcode.com/api/v4/projects"
```

## 认证

大多数 GitCode API 请求需要认证，或者在未提供认证时只返回公共数据。每个接口的文档都会指明是否需要认证。

你可以通过以下几种方式之一使用 GitCode API 进行认证：

- OAuth 2.0 令牌
- [个人访问令牌](https://docs.gitcode.com/docs/users/pat/)
- 会话 cookie

如果认证信息无效或缺失，GitCode 会返回带有 `401` 状态码的错误信息：

```json
{
  "message": "401 Unauthorized"
}
```

### OAuth 2.0 令牌

你可以通过在 `access_token` 参数或 `Authorization` 头中传递 OAuth 2.0 令牌来使用 API 进行认证。

使用参数中的 OAuth 2.0 令牌的示例：

```bash
curl "https://api.gitcode.com/api/v4/projects?access_token=OAUTH-TOKEN"
```

请求头中的 OAuth 2.0 令牌的示例：

```bash
curl --header "Authorization: Bearer OAUTH-TOKEN" "https://api.gitcode.com/api/v4/projects"
```

所有 OAuth 访问令牌在创建后两小时内有效。你可以使用 `refresh_token` 参数来刷新令牌。

### 个人访问令牌

你可以通过 `PRIVATE-TOKEN` 头中传递个人访问令牌来使用 API 进行认证。

请求头中的个人访问令牌的示例：

```bash
curl --header "PRIVATE-TOKEN: <your_access_token>" "https://api.gitcode.com/api/v4/projects"
```

你还可以使用与 OAuth 兼容的头来使用个人访问令牌：

```bash
curl --header "Authorization: Bearer <your_access_token>" "https://api.gitcode.com/api/v4/projects"
```

### 会话 cookie

用户登录 GitCode 后系统会设置一个 `GitCode_ACCESS_TOKEN` cookie。如果此 cookie 存在，API 将使用它进行认证。不支持使用 API 生成新的会话 cookie。

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

## 重定向

在项目路径更改之后，REST API 可以响应重定向，用户应能够处理此类响应。用户应跟随重定向并重复请求到 `Location` 头指定的 URI。

项目移动到不同路径的示例：

```bash
curl --verbose "https://api.gitcode.com/api/v4/projects/gitcode-org%2Fold-path-project"
```

响应为：

```bash
...
< Location: http://gitcode.example.com/api/v4/projects/81
...
此资源已被永久移动到 https://api.gitcode.com/api/v4/projects/81
```

## 分页

GitCode 支持以下分页方法：

- 基于偏移的分页。默认方法，除 `users` 接口外，所有接口均可用。
- 基于键集的分页。添加到选定的接口。

对于大型集合，出于性能原因考虑你应该使用键集分页（如果可用）而不是偏移分页。

### 基于偏移的分页

有时，返回的结果跨越许多页面。当列出资源时，你可以传递以下参数：

|参数|描述|
|---|---|
|`page`|页码（默认：`1`） |
|`per_page`|每页列出的项目数量（默认：`20`，最大：`100`） |

在以下示例中，我们每页列出 50 个 Namespaces：

```bash
curl --request GET --header "PRIVATE-TOKEN: <your_access_token>" "https://api.gitcode.com/api/v4/namespaces?per_page=50"
```

#### 分页 `Link` 头

每个响应返回 `Link` 头，它们的 `rel` 设置为 `prev`、`next`、`first` 或 `last` 并包含相关 URL。确保使用这些链接而不是生成自己的 URL。

在以下 cURL 示例中，我们将输出限制为每页三个项目 (`per_page=3`)，我们请求 ID 为 `8` 的问题的第二页 (`page=2`) 的评论，该问题属于 ID 为 `9` 的项目：

```bash
curl --head --header "PRIVATE-TOKEN: <your_access_token>" "https://api.gitcode.com/api/v4/projects/9/issues/8/notes?per_page=3&page=2"
```

响应为：

```bash
HTTP/2 200 OK
cache-control: no-cache
content-length: 1103
content-type: application/json
date: Mon, 18 Jan 2016 09:43:18 GMT
link: <https://api.gitcode.com/api/v4/projects/8/issues/8/notes?page=1&per_page=3>; rel="prev", <https://api.gitcode.com/api/v4/projects/8/issues/8/notes?page=3&per_page=3>; rel="next", <https://api.gitcode.com/api/v4/projects/8/issues/8/notes?page=1&per_page=3>; rel="first", <https://api.gitcode.com/api/v4/projects/8/issues/8/notes?page=3&per_page=3>; rel="last"
status: 200 OK
vary: Origin
x-next-page: 3
x-page: 2
x-per-page: 3
x-prev-page: 1
x-request-id: 732ad4ee-9870-4866-a199-a9db0cde3c86
x-runtime: 0.108688
x-total: 8
x-total-pages: 3
```

#### 其他分页头

GitCode 还返回以下附加分页头：

|头|描述|
|---|---|
|`x-next-page`|下一页的索引 |
|`x-page`|当前页面的索引（从 1 开始） |
|`x-per-page`|每页的项目数 |
|`x-prev-page`|上一页的索引 |
|`x-total`|项目总数 |
|`x-total-pages`|总页数 |

### 基于键集的分页

键集分页允许更高效地检索页面，与基于偏移的分页相比，运行时间与集合的大小无关。

此方法由以下参数控制。`order_by` 和 `sort` 都是必需的。

|参数|必需|描述|
|---|---|---|
|`pagination`|是|`keyset`（启用键集分页） |
|`per_page`|否|每页列出的项目数量（默认：`20`，最大：`100`） |
|`order_by`|是|排序依据的列 |
|`sort`|是|排序顺序（`asc` 或 `desc`） |

在以下示例中，我们每页列出 50 个 [项目](https://docs.gitcode.com/docs/api/repos/#11-%e9%a1%b9%e7%9b%ae%e5%88%97%e8%a1%a8)，按 `id` 升序排序。

```bash
curl --request GET --header "PRIVATE-TOKEN: <your_access_token>" "https://api.gitcode.com/api/v4/projects?pagination=keyset&per_page=50&order_by=id&sort=asc"
```

响应头包括指向下一页的链接。例如：

```bash
HTTP/1.1 200 OK
...
Link: <https://api.gitcode.com/api/v4/projects?pagination=keyset&per_page=50&order_by=id&sort=asc&id_after=42>; rel="next"
Status: 200 OK
...
```

下一页的链接包含一个额外的过滤器 `id_after=42`，用于排除已检索的记录。

作为另一个示例，以下请求使用键集分页按 `name` 升序每页列出 50 个 [组](https://docs.gitcode.com/docs/api/orgs/#3-%e8%8e%b7%e5%8f%96%e7%bb%84%e7%bb%87%e5%88%97%e8%a1%a8)：

```bash
curl --request GET --header "PRIVATE-TOKEN: <your_access_token>" "https://api.gitcode.com/api/v4/groups?pagination=keyset&per_page=50&order_by=name&sort=asc"
```

响应头包括指向下一页的链接：

```bash
HTTP/1.1 200 OK
...
Link: <https://api.gitcode.com/api/v4/groups?pagination=keyset&per_page=50&order_by=name&sort=asc&cursor=eyJuYW1lIjoiRmxpZ2h0anMiLCJpZCI6IjI2IiwiX2tkIjoibiJ9>; rel="next"
Status: 200 OK
...
```

下一页的链接包含一个额外的过滤器 `cursor=eyJuYW1lIjoiRmxpZ2h0anMiLCJpZCI6IjI2IiwiX2tkIjoibiJ9`，用于排除已检索的记录。

过滤器的类型取决于使用的 `order_by` 选项，我们可以有不止一个额外的过滤器。当到达集合的末尾且没有其他记录可检索时，`Link` 头不存在，结果数组为空。

你应该仅使用给定的链接检索下一页，而不是构建自己的 URL。除了显示的头之外，我们不公开其他分页头。

### 分页响应头

出于性能原因，如果查询返回超过 10,000 条记录，GitCode 不返回以下头信息：

- `x-total`
- `x-total-pages`
- `rel="last"` `link`

## 路径参数

如果接口有路径参数，文档会用前缀冒号显示它们。

例如：

```bash
DELETE /projects/:id/share/:group_id
```

需要将 `:id` 路径参数替换为项目 ID，将 `:group_id` 替换为组的 ID。不应包含冒号 `:`。

对于 ID 为 `5` 项目的 cURL 请求为：

```bash
curl --request DELETE --header "PRIVATE-TOKEN: <your_access_token>" "https://api.gitcode.com/api/v4/projects/5"
```

路径参数需要被 URL 编码。如果没有遵循这一规则，则无法匹配到 API 接口，并且会收到 404 错误。如果 API 前面有其他组件（例如 Apache），确保它不会对 URL 编码过的路径参数进行解码。

## 命名空间路径编码

使用命名空间 API 请求时，确保 `NAMESPACE/PROJECT_PATH` 被 URL 编码。

例如，`/` 被表示为 `%2F`：

```bash
GET /api/v4/projects/diaspora%2Fdiaspora
```

项目的 _路径_ 不一定与其 _名称_ 相同。项目路径可以在项目的 URL 或项目设置中找到，位于 **项目设置 > 高级设置 > 更改路径** 部分。

## 文件路径、分支和标签名称编码

如果文件路径、分支或标签包含 `/`，确保它被 URL 编码。

例如，`/` 被表示为 `%2F`：

```bash
GET /api/v4/projects/1/repository/files/src%2FREADME.md?ref=main
GET /api/v4/projects/1/branches/my%2Fbranch/commits
GET /api/v4/projects/1/repository/tags/my%2Ftag
```

## 请求负载

API 请求可以使用作为[查询字符串](https://en.wikipedia.org/wiki/Query_string)或[负载体](https://datatracker.ietf.org/doc/html/draft-ietf-httpbis-p3-payload-14#section-3.2)发送的参数。GET 请求通常发送查询字符串，而 PUT 或 POST 请求通常发送负载体：

- 查询字符串：

    ```bash
    curl --request POST "https://api.gitcode.com/api/v4/projects?name=<example-name>&description=<example-description>"
    ```

- 请求负载（JSON）：

    ```bash
    curl --request POST --header "Content-Type: application/json" \
         --data '{"name":"<example-name>", "description":"<example-description>"}' "https://api.gitcode.com/api/v4/projects"
    ```

URL 编码的查询字符串有长度限制。太大的请求会导致 `414 Request-URI Too Large` 错误消息。这可以通过使用负载体来解决。

## 编码 `array` 和 `hash` 类型的 API 参数

你可以使用 `array` 和 `hash` 类型的参数请求 API：

### `array`

`import_sources` 是一个 `array` 类型的参数：

```
curl --request POST --header "PRIVATE-TOKEN: <your_access_token>" \
-d "import_sources[]=GitCode" \
-d "import_sources[]=bitbucket" \
"https://api.gitcode.com/api/v4/some_endpoint"
```

### `hash`

`override_params` 是一个 `hash` 类型的参数：

```
curl --request POST --header "PRIVATE-TOKEN: <your_access_token>" \
--form "namespace=email" \
--form "path=impapi" \
--form "file=@/path/to/somefile.txt" \
--form "override_params[visibility]=private" \
--form "override_params[some_other_param]=some_value" \
"https://api.gitcode.com/api/v4/projects/import"
```

### 哈希数组

`variables` 是一个包含哈希键值对 `[{ 'key': 'UPLOAD_TO_S3', 'value': 'true' }]` 的 `array` 类型参数：

```
curl --globoff --request POST --header "PRIVATE-TOKEN: <your_access_token>" \
"https://api.gitcode.com/api/v4/projects/169/pipeline?ref=master&variables[0][key]=VAR1&variables[0][value]=hello&variables[1][key]=VAR2&variables[1][value]=world"

curl --request POST --header "PRIVATE-TOKEN: <your_access_token>" \
--header "Content-Type: application/json" \
--data '{ "ref": "master", "variables": [ {"key": "VAR1", "value": "hello"}, {"key": "VAR2", "value": "world"} ] }' \
"https://api.gitcode.com/api/v4/projects/169/pipeline"
```

## `id` 与 `iid`

某些资源有两个名称相似的字段。例如，[ISSUE](https://docs.gitcode.com/docs/api/issue/)、[PR](https://docs.gitcode.com/docs/api/pulls/)和[项目里程碑](https://docs.gitcode.com/docs/api/issue/#12-%e8%8e%b7%e5%8f%96%e9%a1%b9%e7%9b%ae%e4%b8%8b%e7%9a%84%e9%87%8c%e7%a8%8b%e7%a2%91%e5%88%97%e8%a1%a8)。这些字段是：

- `id`：在所有项目中唯一的 ID。
- `iid`：附加的内部 ID（在网页 UI 中显示），在单个项目的范围内唯一。

如果资源同时有 `iid` 字段和 `id` 字段，通常使用 `iid` 字段而不是 `id` 来获取资源。

例如，假设一个项目的 `id: 42` 有一个问题的 `id: 46` 和 `iid: 5`。在这种情况下：

- 检索问题的有效 API 请求是 `GET /projects/42/issues/5`。
- 检索问题的无效 API 请求是 `GET /projects/42/issues/46`。


## `null` 与 `false`

在 API 响应中，某些布尔字段可以有 `null` 值。一个 `null` 布尔值没有默认值，既不是 `true` 也不是 `false`。GitCode 将布尔字段中的 `null` 值与 `false` 相同对待。

在布尔参数中，你应该只设置 `true` 或 `false` 值（不是 `null`）。

## 数据验证与错误报告

在使用 API 时，你可能会遇到验证错误，此时 API 会返回 HTTP `400` 错误。

这类错误出现在以下情况：

- API 请求缺少一个必需的属性（例如，没有提供问题的标题）。
- 属性未通过验证（例如，用户简介太长）。

当属性缺失时，你会收到类似以下内容：

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
{
    "message":"400 (Bad request) \"title\" not given"
}
```

当发生验证错误时，错误消息不同。它们包含了验证错误的所有细节：

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
{
    "message": {
        "bio": [
            "is too long (maximum is 255 characters)"
        ]
    }
}
```

这使得错误消息更易于机器阅读。格式可以描述如下：

```
{
    "message": {
        "<property-name>": [
            "<error-message>",
            "<error-message>",
            ...
        ],
        "<embed-entity>": {
            "<property-name>": [
                "<error-message>",
                "<error-message>",
                ...
            ],
        }
    }
}
```

## 未知路由

当你尝试访问不存在的 API URL 时，你会收到 404 Not Found 消息。

```
HTTP/1.1 404 Not Found
Content-Type: application/json
{
    "error": "404 Not Found"
}
```

## 编码 ISO 8601 日期中的 `+`

如果你需要在查询参数中包含一个 `+`，你可能需要使用 `%2B` 替代，这是由于 [W3 推荐](https://www.w3.org/Addressing/URL/4_URI_Recommentations.html) 导致 `+` 被解释为空格。例如，在 ISO 8601 日期中，你可能想包含一个特定的时间，如：

```
2017-10-17T23:11:13.000+05:30
```

正确的查询参数编码应该是：

```
2017-10-17T23:11:13.000%2B05:30
```

## 速率限制

GitCode 限制了你在特定时间内可以发出的 REST API 请求的数量。此限制有助于防止滥用和拒绝服务攻击，并确保所有用户都可以使用 API。

一般来说，你可以根据下文描述的认证方式计算 REST API 的主要速率限制。

### 未经认证用户的主要速率限制

如果你只是获取公共数据，可以进行未经认证的请求。未经认证的请求与发起请求的原始 IP 地址相关联，而不是与发出请求的用户或应用程序相关联。

未经认证请求的主要速率限制是每小时 60 次请求。

### 经认证用户的主要速率限制

你可以使用个人访问令牌进行 API 请求。此外，你可以授权 GitCode 应用或 OAuth 应用代表你发出 API 请求。

所有这些请求都计入你每小时 5,000 次请求的个人速率限制。

<!-->
### GitCode 应用安装的主要速率限制

使用安装访问令牌认证的 GitCode 应用使用安装的最低速率限制，即每小时 5,000 次请求。
<!-->

### OAuth 应用的主要速率限制

由 OAuth 应用生成的 OAuth 访问令牌的主要速率限制由经认证用户的主要速率限制决定。这个速率限制与其他 GitCode 应用或 OAuth 应用代表该用户发出的任何请求以及用户使用个人访问令牌发出的任何请求相结合。

OAuth 应用还可以使用其客户端 ID 和客户端密钥获取公共数据。例如：

```shell
curl -u YOUR_CLIENT_ID:YOUR_CLIENT_SECRET -I https://api.gitcode.com/meta
```

对于这些请求，每个 OAuth 应用每小时的速率限制为 5,000 次请求。

**注意：**永远不要在客户端代码或在用户设备上运行的代码中包含你应用的客户端密钥。客户端密钥可用于为授权你的应用的用户生成 OAuth 访问令牌，因此你应始终保持客户端密钥的安全。

## 关于次要速率限制

除了主要速率限制之外，GitCode 还强制执行次要速率限制，以防止滥用并保持 API 对所有用户的可用性。

如果你：

- _发出太多并发请求。_ 允许的最大并发请求数为 100。此限制在 REST API 和 GraphQL API 之间共享。
- _每分钟对单个接口发出太多请求。_ REST API 接口每分钟最多允许 900 点，GraphQL API 接口每分钟最多允许 2,000 点。有关点数的更多信息，请参考计算次要速率限制的点数。
- _每分钟发出太多请求。_ 允许的最大 CPU 时间为每 60 秒的真实时间中的 90 秒。GraphQL API 的 CPU 时间不得超过 60 秒。你可以通过测量 API 请求的总响应时间来大致估算 CPU 时间。
- _在短时间内在 GitCode 上创建过多内容。_ 通常，每分钟最多允许 80 个生成内容的请求，每小时最多允许 500 个生成内容的请求。某些接口的内容创建限制更低。内容创建限制包括在 GitCode 网页界面上以及通过 REST API 执行的操作。

这些次要速率限制可能会在不事先通知的情况下更改。你也可能因为未公开的原因遇到次要速率限制。

### 计算次要速率限制的点数

某些次要速率限制是通过请求的点值来确定的。

|请求|点数|
|---|---|
|不包含变更的 GraphQL 请求|1|
|包含变更的 GraphQL 请求|5|
|大多数 REST API `GET`、`HEAD` 和 `OPTIONS` 请求|1|
|大多数 REST API `POST`、`PATCH`、`PUT` 或 `DELETE` 请求|5|

某些 REST API 接口有不公开的不同点成本。

## 检查你的速率限制状态

你可以使用每个响应发送的头信息来确定你的主要速率限制的当前状态。

|头信息名称|描述|
|---|---|
|`x-ratelimit-limit`|你每小时可以发出的最大请求数|
|`x-ratelimit-remaining`|当前速率限制窗口中剩余的请求数|
|`x-ratelimit-used`|你在当前速率限制窗口中已发出的请求数|
|`x-ratelimit-reset`|当前速率限制窗口重置的时间，以 UTC 纪元秒为单位|
|`x-ratelimit-resource`|请求计入的速率限制资源。|

你还可以调用 `GET /rate_limit` 接口来检查你的速率限制。调用此接口不会占用你的速率限制。

没有方法可以检查你的次要速率限制状态。

## 超出速率限制

如果你超出了主要速率限制，你将收到 `403` 或 `429` 响应，且 `x-ratelimit-remaining` 头将为 `0`。你不应在 `x-ratelimit-reset` 头指定的时间之前重试你的请求。

继续在你被速率限制时发出请求可能会导致你的集成被禁用。