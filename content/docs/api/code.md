---
linkTitle: 仓库文件 API
title: 仓库文件接口API文档
weight: 6
sidebar:
  open: false
---

## 1. 文件详情

### 请求

`GET /api/v4/projects/{project_id}/repository/files`

### 参数

| 参数名        | 类型     | IN    | 必选   | 描述            |
|------------|--------|:------|:-----|---------------|
| project_id | string | path  | true | 项目id          |
| ref        | string | query | true | commit、分支、tag |
| file_path  | string | query | true | 文件路径          |

### 响应

返回码: 200

```json
{
  "name": "README.md",
  "path": "README.md",
  "size": 6,
  "encoding": "base64",
  "ref": "main",
  "blob_id": "db0fa21220c5269d2367ce5a683fc822a6edb3b0",
  "file_type": "text/plain; charset=utf-8",
  "commit": {
    "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
    "message": " ",
    "parent_ids": [
      "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
    ],
    "authored_date": "2023-09-19T10:07:44.000Z",
    "author_name": "xxx",
    "committed_date": "2023-09-19T10:07:44.000Z",
    "committer_name": "xxx",
    "short_id": "a655c6cf",
    "title": "feat: 修改README.md ",
    "author_avatar_url": "https://xxx",
    "author_nick_name": "xx",
    "committer_avatar_url": "https://xxx",
    "committer_nick_name": "xx"
  },
  "content": "YWFhYWNj",
  "is_limited": false,
  "content_sha256": "fa8b33989e4f577aa85ca4a0d305bb6f814e505c47e177d93fb5a1706ab0dc5e",
  "last_commit_id": "a48f84bfe33ede83f17c892f47d23210f6c9e53b"
}
```

## 2. 获取文件内容

### 请求

`GET /api/v4/projects/{project_id}/repository/files/{file_path}/raw`

### 参数

| 参数名        | 类型     | IN    | 必选   | 描述            |
|------------|--------|:------|:-----|---------------|
| project_id | string | path  | true | 项目id          |
| file_path  | string | path  | true | 文件路径          |
| ref        | string | query | true | commit、分支、tag |

### 响应

返回码: 200

```text
aaaacc
```

## 3. 创建分支

### 请求

`POST /api/v4/projects/{project_id}/repository/branches`

### 参数

| 参数名         | 类型     | IN   | 必选    | 描述            |
|-------------|--------|:-----|:------|---------------|
| project_id  | string | path | true  | 项目id          |
| branch      | string | body | true  | 分支名称          |
| ref         | string | body | true  | commit、分支、tag |
| description | string | body | false | 分支描述          |

### 响应

返回码: 201

| 参数名                     | 类型                           | 描述         |
|-------------------------|------------------------------|------------|
| name                    | string                       | 分支名称       |
| default                 | boolean                      | 是否默认分支     |
| commit                  | Commit object                | commit信息   |
| merged                  | boolean                      | 是否已合并      |
| protected               | boolean                      | 是否保护分支     |
| creator                 | Creator object               | 创建者信息      |
| created_at              | string                       | 创建时间       |
| description             | string                       | 分支描述       |
| create_source           | string                       | 创建来源       |
| create_source_exists    | boolean                      | 创建来源是否存在   |
| opened_mr_count         | integer                      | 开启中pr个数    |
| diverging_commit_counts | DivergingCommitCounts object | 相对默认分支推送情况 |
| developers_can_push     | boolean                      | 开发者是否能推送   |
| developers_can_merge    | boolean                      | 开发者是否能合并   |
| can_push                | boolean                      | 是否能推送      |

#### Commit

| 参数名                  | 类型              | 描述        |
|----------------------|-----------------|-----------|
| id                   | string          | commit id |
| short_id             | string          | commit短id |
| title                | string          | 提交标题      |
| message              | string          | 提交信息      |
| parent_ids           | array of string | 祖先id      |
| authored_date        | string          | 创建日期      |
| author_name          | string          | 创建者       |
| author_avatar_url    | string          | 创建者头像     |
| author_nick_name     | string          | 创建者昵称     |
| committed_date       | string          | 提交日期      |
| committer_name       | string          | 提交者       |
| committer_avatar_url | string          | 提交者头像     |
| committer_nick_name  | string          | 提交者昵称     |

#### Creator

| 参数名        | 类型     | 描述   |
|------------|--------|------|
| username   | string | 用户名  |
| nick_name  | string | 用户昵称 |
| avatar_url | string | 用户头像 |

#### DivergingCommitCounts

| 参数名    | 类型      | 描述  |
|--------|---------|-----|
| behind | integer | 滞后  |
| ahead  | integer | 超前  |

```json
{
  "name": "branch1",
  "default": false,
  "commit": {
    "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
    "message": " ",
    "parent_ids": [
      "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
    ],
    "authored_date": "2023-09-19T10:07:44.000Z",
    "author_name": "xxx",
    "committed_date": "2023-09-19T10:07:44.000Z",
    "committer_name": "xxx",
    "short_id": "a655c6cf",
    "title": "feat: 修改README.md ",
    "author_avatar_url": "https://xxx",
    "author_nick_name": "xx",
    "committer_avatar_url": "https://xxx",
    "committer_nick_name": "xx"
  },
  "merged": false,
  "protected": false,
  "created_at": "2023-09-19T20:18:13.627+08:00",
  "creator": {
    "username": "xxx",
    "nick_name": "xx",
    "avatar_url": "https://xxx"
  },
  "description": "aaa",
  "create_source": "main",
  "create_source_exists": true,
  "opened_mr_count": 0,
  "diverging_commit_counts": {
    "behind": 0,
    "ahead": 0
  },
  "developers_can_push": false,
  "developers_can_merge": false,
  "can_push": true
}
```

## 4. 删除分支

### 请求

`DELETE /api/v4/projects/{project_id}/repository/branches`

### 参数

| 参数名        | 类型     | IN    | 必选   | 描述   |
|------------|--------|:------|:-----|------|
| project_id | string | path  | true | 项目id |
| branch     | string | query | true | 分支名称 |

### 响应

返回码: 204

## 5. 分支列表

### 请求

`GET /api/v4/projects/{project_id}/repository/branches`

### 参数

| 参数名         | 类型      | IN    | 必选    | 默认值          | 描述                                                                                            |
|-------------|---------|:------|:------|--------------|-----------------------------------------------------------------------------------------------|
| project_id  | string  | path  | true  |              | 项目id                                                                                          |
| view        | string  | query | true  | all          | 返回数据 simple、basic、all                                                                         |
| branch_type | string  | query | false |              | 过滤保护分支 not_protect                                                                            |
| search      | string  | query | false |              | 搜索关键字                                                                                         |
| sort        | string  | query | false | updated_desc | 排序顺序 name、name_asc、updated_desc、updated_asc、created_asc、created_desc、type、mr_source、mr_target |
| query_view  | string  | query | false | all_branches | 分支类型 my_branches、active、stale、all_branches                                                    |
| page        | integer | query | false | 1            | 分页参数-当前页                                                                                      |
| per_page    | integer | query | false | 20           | 分页参数-每页条数                                                                                     |

### 响应

返回码: 200

| 参数名                     | 类型                           | 描述         | 返回数据             |
|-------------------------|------------------------------|------------|------------------|
| name                    | string                       | 分支名称       | simple、basic、all |
| default                 | boolean                      | 是否默认分支     | simple、basic、all |
| commit                  | Commit object                | commit信息   | basic、all        |
| merged                  | boolean                      | 是否已合并      | basic、all        |
| protected               | boolean                      | 是否保护分支     | basic、all        |
| creator                 | Creator object               | 创建者信息      | basic、all        |
| created_at              | string                       | 创建时间       | basic、all        |
| description             | string                       | 分支描述       | basic、all        |
| create_source           | string                       | 创建来源       | basic、all        |
| create_source_exists    | boolean                      | 创建来源是否存在   | basic、all        |
| opened_mr_count         | integer                      | 开启中pr个数    | basic、all        |
| diverging_commit_counts | DivergingCommitCounts object | 相对默认分支推送情况 | basic、all        |
| developers_can_push     | boolean                      | 开发者是否能推送   | all              |
| developers_can_merge    | boolean                      | 开发者是否能合并   | all              |
| can_push                | boolean                      | 是否能推送      | all              |

#### Commit

| 参数名                  | 类型              | 描述        |
|----------------------|-----------------|-----------|
| id                   | string          | commit id |
| short_id             | string          | commit短id |
| title                | string          | 提交标题      |
| message              | string          | 提交信息      |
| parent_ids           | array of string | 祖先id      |
| authored_date        | string          | 创建日期      |
| author_name          | string          | 创建者       |
| author_avatar_url    | string          | 创建者头像     |
| author_nick_name     | string          | 创建者昵称     |
| committed_date       | string          | 提交日期      |
| committer_name       | string          | 提交者       |
| committer_avatar_url | string          | 提交者头像     |
| committer_nick_name  | string          | 提交者昵称     |

#### Creator

| 参数名        | 类型     | 描述   |
|------------|--------|------|
| username   | string | 用户名  |
| nick_name  | string | 用户昵称 |
| avatar_url | string | 用户头像 |

#### DivergingCommitCounts

| 参数名    | 类型      | 描述  |
|--------|---------|-----|
| behind | integer | 滞后  |
| ahead  | integer | 超前  |

```json
[
  {
    "name": "branch1",
    "default": false,
    "commit": {
      "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
      "message": " ",
      "parent_ids": [
        "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
      ],
      "authored_date": "2023-09-19T10:07:44.000Z",
      "author_name": "xxx",
      "committed_date": "2023-09-19T10:07:44.000Z",
      "committer_name": "xxx",
      "short_id": "a655c6cf",
      "title": "feat: 修改README.md ",
      "author_avatar_url": "https://xxx",
      "author_nick_name": "xx",
      "committer_avatar_url": "https://xxx",
      "committer_nick_name": "xx"
    },
    "merged": false,
    "protected": false,
    "created_at": "2023-09-19T20:18:13.627+08:00",
    "creator": {
      "username": "xxx",
      "nick_name": "xx",
      "avatar_url": "https://xxx"
    },
    "description": "aaa",
    "create_source": "main",
    "create_source_exists": true,
    "latest_pipeline": null,
    "opened_mr_count": 0,
    "diverging_commit_counts": {
      "behind": 0,
      "ahead": 0
    },
    "developers_can_push": false,
    "developers_can_merge": false,
    "can_push": true
  }
]
```

## 6. 提交记录列表

### 请求

`GET /api/v4/projects/{project_id}/repository/commits`

### 参数

| 参数名        | 类型      | IN    | 必选    | 描述               |
|------------|---------|:------|:------|------------------|
| project_id | string  | path  | true  | 项目id             |
| ref_name   | string  | query | false | 分支、tag           |
| since      | string  | query | false | 时间筛选-起始时间        |
| until      | string  | query | false | 时间筛选-终止时间        |
| path       | string  | query | false | 文件路径             |
| message    | string  | query | false | 提交信息或commit id查询 |
| author     | string  | query | false | 提交者查询            |
| follow     | boolean | query | false | 文件重命名追踪参数        |
| page       | integer | query | false | 分页参数-当前页         |
| per_page   | integer | query | false | 分页参数-每页条数        |

### 响应

返回码: 200

#### Commit

| 参数名                  | 类型              | 描述        |
|----------------------|-----------------|-----------|
| id                   | string          | commit id |
| short_id             | string          | commit短id |
| title                | string          | 提交标题      |
| message              | string          | 提交信息      |
| parent_ids           | array of string | 祖先id      |
| authored_date        | string          | 创建日期      |
| author_name          | string          | 创建者       |
| author_avatar_url    | string          | 创建者头像     |
| author_nick_name     | string          | 创建者昵称     |
| committed_date       | string          | 提交日期      |
| committer_name       | string          | 提交者       |
| committer_avatar_url | string          | 提交者头像     |
| committer_nick_name  | string          | 提交者昵称     |

```json
[
  {
    "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
    "message": " ",
    "parent_ids": [
      "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
    ],
    "authored_date": "2023-09-19T10:07:44.000Z",
    "author_name": "xxx",
    "committed_date": "2023-09-19T10:07:44.000Z",
    "committer_name": "xxx",
    "short_id": "a655c6cf",
    "title": "feat: 修改README.md ",
    "author_avatar_url": "https://xxx",
    "author_nick_name": "xx",
    "committer_avatar_url": "https://xxx",
    "committer_nick_name": "xx"
  }
]
```

## 7. 提交记录详情

### 请求

`GET /api/v4/projects/{project_id}/repository/commits/{sha}`

### 参数

| 参数名        | 类型     | IN   | 必选   | 描述            |
|------------|--------|:-----|:-----|---------------|
| project_id | string | path | true | 项目id          |
| sha        | string | path | true | commit、分支、tag |

### 响应

返回码: 200

#### Commit

| 参数名                  | 类型              | 描述        |
|----------------------|-----------------|-----------|
| id                   | string          | commit id |
| short_id             | string          | commit短id |
| title                | string          | 提交标题      |
| message              | string          | 提交信息      |
| parent_ids           | array of string | 祖先id      |
| authored_date        | string          | 创建日期      |
| author_name          | string          | 创建者       |
| author_avatar_url    | string          | 创建者头像     |
| author_nick_name     | string          | 创建者昵称     |
| committed_date       | string          | 提交日期      |
| committer_name       | string          | 提交者       |
| committer_avatar_url | string          | 提交者头像     |
| committer_nick_name  | string          | 提交者昵称     |

```json
{
  "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
  "message": " ",
  "parent_ids": [
    "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
  ],
  "authored_date": "2023-09-19T10:07:44.000Z",
  "author_name": "xxx",
  "committed_date": "2023-09-19T10:07:44.000Z",
  "committer_name": "xxx",
  "short_id": "a655c6cf",
  "title": "feat: 修改README.md ",
  "author_avatar_url": "https://xxx",
  "author_nick_name": "xx",
  "committer_avatar_url": "https://xxx",
  "committer_nick_name": "xx"
}
```

## 8. 创建tag

### 请求

`POST /api/v4/projects/{project_id}/repository/tags`

### 参数

| 参数名                 | 类型     | IN   | 必选    | 描述                        |
|---------------------|--------|:-----|:------|---------------------------|
| project_id          | string | path | true  | 项目id                      |
| tag_name            | string | body | true  | tag名称                     |
| ref                 | string | body | true  | 基于ref(commit id、分支、tag)创建 |
| message             | string | body | false | tag描述                     |
| release_description | string | body | false | release描述                 |

### 响应

返回码: 201

| 参数名     | 类型             | 描述        |
|---------|----------------|-----------|
| name    | string         | tag名称     |
| message | string         | tag描述     |
| target  | string         | tag来源     |
| commit  | Commit object  | commit信息  |
| release | Release object | release信息 |

#### Commit

| 参数名                  | 类型              | 描述        |
|----------------------|-----------------|-----------|
| id                   | string          | commit id |
| short_id             | string          | commit短id |
| title                | string          | 提交标题      |
| message              | string          | 提交信息      |
| parent_ids           | array of string | 祖先id      |
| authored_date        | string          | 创建日期      |
| author_name          | string          | 创建者       |
| author_avatar_url    | string          | 创建者头像     |
| author_nick_name     | string          | 创建者昵称     |
| committed_date       | string          | 提交日期      |
| committer_name       | string          | 提交者       |
| committer_avatar_url | string          | 提交者头像     |
| committer_nick_name  | string          | 提交者昵称     |

#### Release

| 参数名         | 类型     | 描述        |
|-------------|--------|-----------|
| tag_name    | string | tag名称     |
| description | string | release描述 |

```json
{
  "name": "t11",
  "message": "tag11",
  "target": "b96cb71b9de3866b3ac6b1f408e047bfe4f14090",
  "commit": {
    "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
    "message": " ",
    "parent_ids": [
      "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
    ],
    "authored_date": "2023-09-19T10:07:44.000Z",
    "author_name": "xxx",
    "committed_date": "2023-09-19T10:07:44.000Z",
    "committer_name": "xxx",
    "short_id": "a655c6cf",
    "title": "feat: 修改README.md ",
    "author_avatar_url": "https://xxx",
    "author_nick_name": "xx",
    "committer_avatar_url": "https://xxx",
    "committer_nick_name": "xx"
  },
  "release": {
    "tag_name": "t11",
    "description": "release11"
  }
}
```

## 9. 删除tag

### 请求

`DELETE /api/v4/projects/{project_id}/repository/tags/{tag_name}`

### 参数

| 参数名        | 类型     | IN   | 必选   | 描述    |
|------------|--------|:-----|:-----|-------|
| project_id | string | path | true | 项目id  |
| tag_name   | string | path | true | tag名称 |

### 响应

返回码: 204

## 10. tag列表

### 请求

`GET /api/v4/projects/{project_id}/repository/tags`

### 参数

| 参数名        | 类型      | IN    | 必选    | 默认值     | 描述                        |
|------------|---------|:------|:------|---------|---------------------------|
| project_id | string  | path  | true  |         | 项目id                      |
| order_by   | string  | query | false | updated | 排序方式 name、updated、created |
| sort       | string  | query | false | desc    | 排序顺序 asc、desc             |
| view       | string  | query | false | detail  | 返回数据 simple、basic、detail  |
| search     | string  | query | false |         | 搜索关键字                     |
| creator    | string  | query | false |         | 创建者                       |
| page       | integer | query | false | 1       | 分页参数-当前页                  |
| per_page   | integer | query | false | 20      | 分页参数-每页条数                 |

### 响应

返回码: 200

| 参数名                  | 类型             | 描述       | 返回数据                |
|----------------------|----------------|----------|---------------------|
| name                 | string         | tag名称    | simple、basic、detail |
| message              | string         | tag描述    | simple、basic、detail |
| target               | string         | tag来源    | simple、basic、detail |
| commit               | Commit object  | commit信息 | basic、detail        |
| protected            | boolean        | 是否保护tag  | detail              |
| created_at           | string         | 创建时间     | detail              |
| creator              | Creator object | 创建者信息    | detail              |
| create_source        | string         | 创建来源     | detail              |
| create_source_exists | boolean        | 创建来源是否存在 | detail              |

#### Commit

| 参数名                  | 类型              | 描述        |
|----------------------|-----------------|-----------|
| id                   | string          | commit id |
| short_id             | string          | commit短id |
| title                | string          | 提交标题      |
| message              | string          | 提交信息      |
| parent_ids           | array of string | 祖先id      |
| authored_date        | string          | 创建日期      |
| author_name          | string          | 创建者       |
| author_avatar_url    | string          | 创建者头像     |
| author_nick_name     | string          | 创建者昵称     |
| committed_date       | string          | 提交日期      |
| committer_name       | string          | 提交者       |
| committer_avatar_url | string          | 提交者头像     |
| committer_nick_name  | string          | 提交者昵称     |

#### Creator

| 参数名        | 类型     | 描述   |
|------------|--------|------|
| username   | string | 用户名  |
| nick_name  | string | 用户昵称 |
| avatar_url | string | 用户头像 |

```json
[
  {
    "name": "t9",
    "message": "tag5",
    "target": "b96cb71b9de3866b3ac6b1f408e047bfe4f14090",
    "commit": {
      "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
      "message": " ",
      "parent_ids": [
        "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
      ],
      "authored_date": "2023-09-19T10:07:44.000Z",
      "author_name": "xxx",
      "committed_date": "2023-09-19T10:07:44.000Z",
      "committer_name": "xxx",
      "short_id": "a655c6cf",
      "title": "feat: 修改README.md ",
      "author_avatar_url": "https://xxx",
      "author_nick_name": "xx",
      "committer_avatar_url": "https://xxx",
      "committer_nick_name": "xx"
    },
    "protected": false,
    "created_at": "2023-09-20 11:03:03 +0800",
    "creator": {
      "username": "xxx",
      "nick_name": "xx",
      "avatar_url": "https://xxx"
    },
    "create_source": "b96cb71b9de3866b3ac6b1f408e047bfe4f14090",
    "create_source_exists": true
  }
]
```

## 11. 创建release

### 请求

`POST /api/v4/projects/{project_id}/releases`

### 参数

| 参数名         | 类型     | IN   | 必选   | 描述                      |
|-------------|--------|:-----|:-----|-------------------------|
| project_id  | string | path | true | 项目id                    |
| tag_name    | string | body | true | tag名称                   |
| ref         | string | body | true | 依赖ref(commit id、分支、tag) |
| name        | string | body | true | release名称               |
| description | string | body | true | release描述               |

### 响应

返回码: 201

| 参数名         | 类型                  | 描述        |
|-------------|---------------------|-----------|
| tag_name    | string              | tag名称     |
| description | string              | release描述 |
| name        | string              | release名称 |
| created_at  | string              | 创建时间      |
| author      | UserBasicDto object | 创建人       |
| commit      | Commit object       | commit信息  |
| assets      | AssetsDto object    | 附件        |

#### UserBasicDto

| 参数名           | 类型     | 描述   |
|---------------|--------|------|
| id            | string | id   |
| name          | string | 名称   |
| username      | string | 用户名  |
| nick_name     | string | 昵称   |
| authored_date | string | 创建日期 |
| avatar_url    | string | 头像   |
| email         | string | 邮箱   |

#### Commit

| 参数名                  | 类型              | 描述        |
|----------------------|-----------------|-----------|
| id                   | string          | commit id |
| short_id             | string          | commit短id |
| title                | string          | 提交标题      |
| message              | string          | 提交信息      |
| parent_ids           | array of string | 祖先id      |
| authored_date        | string          | 创建日期      |
| author_name          | string          | 创建者       |
| author_avatar_url    | string          | 创建者头像     |
| author_nick_name     | string          | 创建者昵称     |
| committed_date       | string          | 提交日期      |
| committer_name       | string          | 提交者       |
| committer_avatar_url | string          | 提交者头像     |
| committer_nick_name  | string          | 提交者昵称     |

#### AssetsDto

| 参数名     | 类型                         | 描述   |
|---------|----------------------------|------|
| count   | integer                    | 个数   |
| sources | array of  ReleaseSourceDto | 默认附件 |
| assets  | array of  AssetDto         | 上传附件 |

#### ReleaseSourceDto

| 参数名    | 类型     | 描述   |
|--------|--------|------|
| format | string | 格式   |
| url    | string | 下载地址 |

#### AssetDto

| 参数名        | 类型         | 描述         |
|------------|------------|------------|
| id         | integer    | id         |
| project_id | integer    | 项目id       |
| release_id | integer    | release_id |
| attachment | string     | 附件名称       |
| size       | BigDecimal | 附件大小       |
| created_at | string     | 创建时间       |

```json
{
  "tag_name": "tag2",
  "description": "release2",
  "name": "release2",
  "created_at": "2023-09-20T17:31:55.128+08:00",
  "commit": {
    "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
    "message": " ",
    "parent_ids": [
      "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
    ],
    "authored_date": "2023-09-19T10:07:44.000Z",
    "author_name": "xxx",
    "committed_date": "2023-09-19T10:07:44.000Z",
    "committer_name": "xxx",
    "short_id": "a655c6cf",
    "title": "feat: 修改README.md ",
    "author_avatar_url": "https://xxx",
    "author_nick_name": "xx",
    "committer_avatar_url": "https://xxx",
    "committer_nick_name": "xx"
  },
  "assets": {
    "count": 5,
    "sources": [
      {
        "format": "zip",
        "url": "https:/xxxxxx.zip"
      },
      {
        "format": "tar.gz",
        "url": "https://xxxx.tar.gz"
      },
      {
        "format": "tar.bz2",
        "url": "https://xxx.tar.bz2"
      },
      {
        "format": "tar",
        "url": "https://xxx.tar"
      }
    ],
    "assets": [
      {
        "id": xx,
        "project_id": xx,
        "release_id": xx,
        "attachment": "截图aaaaa.PNG",
        "size": 0,
        "created_at": "2024-03-27T00:58:03.661+08:00"
      }
    ]
  }
}
```

## 12. 删除release

### 请求

`DELETE /api/v4/projects/{project_id}/releases/{tag_name}`

### 参数

| 参数名        | 类型     | IN   | 必选   | 描述    |
|------------|--------|:-----|:-----|-------|
| project_id | string | path | true | 项目id  |
| tag_name   | string | path | true | tag名称 |

### 响应

返回码: 200

| 参数名         | 类型                  | 描述        |
|-------------|---------------------|-----------|
| tag_name    | string              | tag名称     |
| description | string              | release描述 |
| name        | string              | release名称 |
| created_at  | string              | 创建时间      |
| author      | UserBasicDto object | 创建人       |
| commit      | Commit object       | commit信息  |
| assets      | AssetsDto object    | 附件        |

#### UserBasicDto

| 参数名           | 类型     | 描述   |
|---------------|--------|------|
| id            | string | id   |
| name          | string | 名称   |
| username      | string | 用户名  |
| nick_name     | string | 昵称   |
| authored_date | string | 创建日期 |
| avatar_url    | string | 头像   |
| email         | string | 邮箱   |

#### Commit

| 参数名                  | 类型              | 描述        |
|----------------------|-----------------|-----------|
| id                   | string          | commit id |
| short_id             | string          | commit短id |
| title                | string          | 提交标题      |
| message              | string          | 提交信息      |
| parent_ids           | array of string | 祖先id      |
| authored_date        | string          | 创建日期      |
| author_name          | string          | 创建者       |
| author_avatar_url    | string          | 创建者头像     |
| author_nick_name     | string          | 创建者昵称     |
| committed_date       | string          | 提交日期      |
| committer_name       | string          | 提交者       |
| committer_avatar_url | string          | 提交者头像     |
| committer_nick_name  | string          | 提交者昵称     |

#### AssetsDto

| 参数名     | 类型                         | 描述   |
|---------|----------------------------|------|
| count   | integer                    | 个数   |
| sources | array of  ReleaseSourceDto | 默认附件 |
| assets  | array of  AssetDto         | 上传附件 |

#### ReleaseSourceDto

| 参数名    | 类型     | 描述   |
|--------|--------|------|
| format | string | 格式   |
| url    | string | 下载地址 |

#### AssetDto

| 参数名        | 类型         | 描述         |
|------------|------------|------------|
| id         | integer    | id         |
| project_id | integer    | 项目id       |
| release_id | integer    | release_id |
| attachment | string     | 附件名称       |
| size       | BigDecimal | 附件大小       |
| created_at | string     | 创建时间       |

```json
{
  "tag_name": "tag2",
  "description": "release2",
  "name": "release2",
  "created_at": "2023-09-20T17:31:55.128+08:00",
  "commit": {
    "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
    "message": " ",
    "parent_ids": [
      "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
    ],
    "authored_date": "2023-09-19T10:07:44.000Z",
    "author_name": "xxx",
    "committed_date": "2023-09-19T10:07:44.000Z",
    "committer_name": "xxx",
    "short_id": "a655c6cf",
    "title": "feat: 修改README.md ",
    "author_avatar_url": "https://xxx",
    "author_nick_name": "xx",
    "committer_avatar_url": "https://xxx",
    "committer_nick_name": "xx"
  },
  "assets": {
    "count": 5,
    "sources": [
      {
        "format": "zip",
        "url": "https:/xxxxxx.zip"
      },
      {
        "format": "tar.gz",
        "url": "https://xxxx.tar.gz"
      },
      {
        "format": "tar.bz2",
        "url": "https://xxx.tar.bz2"
      },
      {
        "format": "tar",
        "url": "https://xxx.tar"
      }
    ],
    "assets": [
      {
        "id": xx,
        "project_id": xx,
        "release_id": xx,
        "attachment": "截图aaaaa.PNG",
        "size": 0,
        "created_at": "2024-03-27T00:58:03.661+08:00"
      }
    ]
  }
}
```

## 13. 更新release

### 请求

`PUT /api/v4/projects/{project_id}/releases/{tag_name}`

### 参数

| 参数名         | 类型     | IN   | 必选   | 描述        |
|-------------|--------|:-----|:-----|-----------|
| project_id  | string | path | true | 项目id      |
| tag_name    | string | path | true | tag名称     |
| name        | string | body | true | release名称 |
| description | string | body | true | release描述 |

### 响应

返回码: 200

| 参数名         | 类型                  | 描述        |
|-------------|---------------------|-----------|
| tag_name    | string              | tag名称     |
| description | string              | release描述 |
| name        | string              | release名称 |
| created_at  | string              | 创建时间      |
| author      | UserBasicDto object | 创建人       |
| commit      | Commit object       | commit信息  |
| assets      | AssetsDto object    | 附件        |

#### UserBasicDto

| 参数名           | 类型     | 描述   |
|---------------|--------|------|
| id            | string | id   |
| name          | string | 名称   |
| username      | string | 用户名  |
| nick_name     | string | 昵称   |
| authored_date | string | 创建日期 |
| avatar_url    | string | 头像   |
| email         | string | 邮箱   |

#### Commit

| 参数名                  | 类型              | 描述        |
|----------------------|-----------------|-----------|
| id                   | string          | commit id |
| short_id             | string          | commit短id |
| title                | string          | 提交标题      |
| message              | string          | 提交信息      |
| parent_ids           | array of string | 祖先id      |
| authored_date        | string          | 创建日期      |
| author_name          | string          | 创建者       |
| author_avatar_url    | string          | 创建者头像     |
| author_nick_name     | string          | 创建者昵称     |
| committed_date       | string          | 提交日期      |
| committer_name       | string          | 提交者       |
| committer_avatar_url | string          | 提交者头像     |
| committer_nick_name  | string          | 提交者昵称     |

#### AssetsDto

| 参数名     | 类型                         | 描述   |
|---------|----------------------------|------|
| count   | integer                    | 个数   |
| sources | array of  ReleaseSourceDto | 默认附件 |
| assets  | array of  AssetDto         | 上传附件 |

#### ReleaseSourceDto

| 参数名    | 类型     | 描述   |
|--------|--------|------|
| format | string | 格式   |
| url    | string | 下载地址 |

#### AssetDto

| 参数名        | 类型         | 描述         |
|------------|------------|------------|
| id         | integer    | id         |
| project_id | integer    | 项目id       |
| release_id | integer    | release_id |
| attachment | string     | 附件名称       |
| size       | BigDecimal | 附件大小       |
| created_at | string     | 创建时间       |

```json
{
  "tag_name": "tag2",
  "description": "release2",
  "name": "release2",
  "created_at": "2023-09-20T17:31:55.128+08:00",
  "commit": {
    "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
    "message": " ",
    "parent_ids": [
      "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
    ],
    "authored_date": "2023-09-19T10:07:44.000Z",
    "author_name": "xxx",
    "committed_date": "2023-09-19T10:07:44.000Z",
    "committer_name": "xxx",
    "short_id": "a655c6cf",
    "title": "feat: 修改README.md ",
    "author_avatar_url": "https://xxx",
    "author_nick_name": "xx",
    "committer_avatar_url": "https://xxx",
    "committer_nick_name": "xx"
  },
  "assets": {
    "count": 5,
    "sources": [
      {
        "format": "zip",
        "url": "https:/xxxxxx.zip"
      },
      {
        "format": "tar.gz",
        "url": "https://xxxx.tar.gz"
      },
      {
        "format": "tar.bz2",
        "url": "https://xxx.tar.bz2"
      },
      {
        "format": "tar",
        "url": "https://xxx.tar"
      }
    ],
    "assets": [
      {
        "id": xx,
        "project_id": xx,
        "release_id": xx,
        "attachment": "截图aaaaa.PNG",
        "size": 0,
        "created_at": "2024-03-27T00:58:03.661+08:00"
      }
    ]
  }
}
```

## 14. release详情

### 请求

`GET /api/v4/projects/{project_id}/releases/{tag_name}`

### 参数

| 参数名        | 类型     | IN   | 必选   | 描述    |
|------------|--------|:-----|:-----|-------|
| project_id | string | path | true | 项目id  |
| tag_name   | string | path | true | tag名称 |

### 响应

返回码: 200

| 参数名         | 类型                  | 描述        |
|-------------|---------------------|-----------|
| tag_name    | string              | tag名称     |
| description | string              | release描述 |
| name        | string              | release名称 |
| created_at  | string              | 创建时间      |
| author      | UserBasicDto object | 创建人       |
| commit      | Commit object       | commit信息  |
| assets      | AssetsDto object    | 附件        |

#### UserBasicDto

| 参数名           | 类型     | 描述   |
|---------------|--------|------|
| id            | string | id   |
| name          | string | 名称   |
| username      | string | 用户名  |
| nick_name     | string | 昵称   |
| authored_date | string | 创建日期 |
| avatar_url    | string | 头像   |
| email         | string | 邮箱   |

#### Commit

| 参数名                  | 类型              | 描述        |
|----------------------|-----------------|-----------|
| id                   | string          | commit id |
| short_id             | string          | commit短id |
| title                | string          | 提交标题      |
| message              | string          | 提交信息      |
| parent_ids           | array of string | 祖先id      |
| authored_date        | string          | 创建日期      |
| author_name          | string          | 创建者       |
| author_avatar_url    | string          | 创建者头像     |
| author_nick_name     | string          | 创建者昵称     |
| committed_date       | string          | 提交日期      |
| committer_name       | string          | 提交者       |
| committer_avatar_url | string          | 提交者头像     |
| committer_nick_name  | string          | 提交者昵称     |

#### AssetsDto

| 参数名     | 类型                         | 描述   |
|---------|----------------------------|------|
| count   | integer                    | 个数   |
| sources | array of  ReleaseSourceDto | 默认附件 |
| assets  | array of  AssetDto         | 上传附件 |

#### ReleaseSourceDto

| 参数名    | 类型     | 描述   |
|--------|--------|------|
| format | string | 格式   |
| url    | string | 下载地址 |

#### AssetDto

| 参数名        | 类型         | 描述         |
|------------|------------|------------|
| id         | integer    | id         |
| project_id | integer    | 项目id       |
| release_id | integer    | release_id |
| attachment | string     | 附件名称       |
| size       | BigDecimal | 附件大小       |
| created_at | string     | 创建时间       |

```json
{
  "tag_name": "tag2",
  "description": "release2",
  "name": "release2",
  "created_at": "2023-09-20T17:31:55.128+08:00",
  "commit": {
    "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
    "message": " ",
    "parent_ids": [
      "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
    ],
    "authored_date": "2023-09-19T10:07:44.000Z",
    "author_name": "xxx",
    "committed_date": "2023-09-19T10:07:44.000Z",
    "committer_name": "xxx",
    "short_id": "a655c6cf",
    "title": "feat: 修改README.md ",
    "author_avatar_url": "https://xxx",
    "author_nick_name": "xx",
    "committer_avatar_url": "https://xxx",
    "committer_nick_name": "xx"
  },
  "assets": {
    "count": 5,
    "sources": [
      {
        "format": "zip",
        "url": "https:/xxxxxx.zip"
      },
      {
        "format": "tar.gz",
        "url": "https://xxxx.tar.gz"
      },
      {
        "format": "tar.bz2",
        "url": "https://xxx.tar.bz2"
      },
      {
        "format": "tar",
        "url": "https://xxx.tar"
      }
    ],
    "assets": [
      {
        "id": xx,
        "project_id": xx,
        "release_id": xx,
        "attachment": "截图aaaaa.PNG",
        "size": 0,
        "created_at": "2024-03-27T00:58:03.661+08:00"
      }
    ]
  }
}
```

## 15. release列表

### 请求

`GET /api/v4/projects/{project_id}/releases`

### 参数

| 参数名        | 类型      | IN    | 必选    | 默认值 | 描述        |
|------------|---------|:------|:------|-----|-----------|
| project_id | string  | path  | true  |     | 项目id      |
| tag_name   | string  | path  | true  |     | tag名称     |
| page       | integer | query | false | 1   | 分页参数-当前页  |
| per_page   | integer | query | false | 20  | 分页参数-每页条数 |

### 响应

返回码: 200

| 参数名         | 类型                  | 描述        |
|-------------|---------------------|-----------|
| tag_name    | string              | tag名称     |
| description | string              | release描述 |
| name        | string              | release名称 |
| created_at  | string              | 创建时间      |
| author      | UserBasicDto object | 创建人       |
| commit      | Commit object       | commit信息  |
| assets      | AssetsDto object    | 附件        |

#### UserBasicDto

| 参数名           | 类型     | 描述   |
|---------------|--------|------|
| id            | string | id   |
| name          | string | 名称   |
| username      | string | 用户名  |
| nick_name     | string | 昵称   |
| authored_date | string | 创建日期 |
| avatar_url    | string | 头像   |
| email         | string | 邮箱   |

#### Commit

| 参数名                  | 类型              | 描述        |
|----------------------|-----------------|-----------|
| id                   | string          | commit id |
| short_id             | string          | commit短id |
| title                | string          | 提交标题      |
| message              | string          | 提交信息      |
| parent_ids           | array of string | 祖先id      |
| authored_date        | string          | 创建日期      |
| author_name          | string          | 创建者       |
| author_avatar_url    | string          | 创建者头像     |
| author_nick_name     | string          | 创建者昵称     |
| committed_date       | string          | 提交日期      |
| committer_name       | string          | 提交者       |
| committer_avatar_url | string          | 提交者头像     |
| committer_nick_name  | string          | 提交者昵称     |

#### AssetsDto

| 参数名     | 类型                         | 描述   |
|---------|----------------------------|------|
| count   | integer                    | 个数   |
| sources | array of  ReleaseSourceDto | 默认附件 |
| assets  | array of  AssetDto         | 上传附件 |

#### ReleaseSourceDto

| 参数名    | 类型     | 描述   |
|--------|--------|------|
| format | string | 格式   |
| url    | string | 下载地址 |

#### AssetDto

| 参数名        | 类型         | 描述         |
|------------|------------|------------|
| id         | integer    | id         |
| project_id | integer    | 项目id       |
| release_id | integer    | release_id |
| attachment | string     | 附件名称       |
| size       | BigDecimal | 附件大小       |
| created_at | string     | 创建时间       |

```json
[
  {
    "tag_name": "tag2",
    "description": "release2",
    "name": "release2",
    "created_at": "2023-09-20T17:31:55.128+08:00",
    "commit": {
      "id": "a655c6cf680db548e33d05e0bb499b1de944515a",
      "message": " ",
      "parent_ids": [
        "7e71b59e15797afe6ee8f101f2cf85e1b4a8aa54"
      ],
      "authored_date": "2023-09-19T10:07:44.000Z",
      "author_name": "xxx",
      "committed_date": "2023-09-19T10:07:44.000Z",
      "committer_name": "xxx",
      "short_id": "a655c6cf",
      "title": "feat: 修改README.md ",
      "author_avatar_url": "https://xxx",
      "author_nick_name": "xx",
      "committer_avatar_url": "https://xxx",
      "committer_nick_name": "xx"
    },
    "assets": {
      "count": 5,
      "sources": [
        {
          "format": "zip",
          "url": "https:/xxxxxx.zip"
        },
        {
          "format": "tar.gz",
          "url": "https://xxxx.tar.gz"
        },
        {
          "format": "tar.bz2",
          "url": "https://xxx.tar.bz2"
        },
        {
          "format": "tar",
          "url": "https://xxx.tar"
        }
      ],
      "assets": [
        {
          "id": xx,
          "project_id": xx,
          "release_id": xx,
          "attachment": "截图aaaaa.PNG",
          "size": 0,
          "created_at": "2024-03-27T00:58:03.661+08:00"
        }
      ]
    }
  }
]
```

## 16. 语言排行

### 请求

`GET /api/v4/projects/{project_id}/repository/languages`

### 参数

| 参数名        | 类型     | IN   | 必选   | 默认值 | 描述   |
|------------|--------|:-----|:-----|-----|------|
| project_id | string | path | true |     | 项目id |

### 响应

返回码: 200

| 参数名       | 类型                    | 描述   |
|-----------|-----------------------|------|
| languages | array of LanguagesDto | 语言列表 |
| status    | string                | -    |

#### LanguagesDto

| 参数名   | 类型     | 描述   |
|-------|--------|------|
| color | string | 颜色   |
| label | string | 语言名称 |
| value | double | 占比   |

```json
{
  "languages": [
    {
      "color": "#375eab",
      "label": "Go",
      "value": 100
    }
  ],
  "status": null
}
```

## 17. 更新文件

### 请求

`PUT /api/v4/projects/{project_id}/repository/files`

### 参数

| 参数名            | 类型     | IN   | 必选    | 描述           |
|----------------|--------|:-----|:------|--------------|
| project_id     | string | path | true  | 项目id         |
| name           | string | body | false | 文件名称         |
| file_path      | string | body | true  | 文件路径         |
| branch         | string | body | true  | 分支名称         |
| commit_message | string | body | true  | 提交信息         |
| start_branch   | string | body | false | 基于分支         |
| author_email   | string | body | false | author邮箱     |
| author_name    | string | body | false | author名称     |
| content        | string | body | true  | 文件内容         |
| encoding       | string | body | false | 文件编码         |
| last_commit_id | string | body | false | 文件最后修改commit |

### 响应

返回码: 200

| 参数名       | 类型     | 描述   |
|-----------|--------|------|
| file_path | string | 文件路径 |
| branch    | string | 分支名称 |

```json
{
  "file_path": "README.md",
  "branch": "master"
}
```

## 18. 获取目录Tree

### 请求

`GET /api/v4/projects/{project_id}/repository/tree`

### 参数

| 参数名        | 类型      | IN    | 必选    | 默认值   | 描述            |
|------------|---------|:------|:------|-------|---------------|
| project_id | string  | path  | true  |       | 项目id          |
| ref        | string  | path  | false | 默认分支  | commit、分支、tag |
| path       | string  | path  | false |       | 文件路径          |
| recursive  | boolean | path  | false | false | 是否递归获取        |
| page       | integer | query | false | 1     | 分页参数-当前页      |
| per_page   | integer | query | false | 20    | 分页参数-每页条数     |

### 响应

返回码: 200

| 参数名              | 类型     | 描述     |
|------------------|--------|--------|
| id               | string | blobId |
| name             | string | 文件名称   |
| type             | string | 文件类型   |
| path             | string | 文件路径   |
| mode             | string | 文件权限   |
| submodule_link   | string | 子模块链接  |
| submodule_branch | string | 子模块分支  |
| md5              | string | MD5    |

```json
[
  {
    "id": "500c78f00047fb43177839468f83476f8fb0f497",
    "name": ".gitmodules",
    "type": "blob",
    "path": ".gitmodules",
    "mode": "100644",
    "submodule_link": null,
    "submodule_branch": null,
    "md5": "8903239df476d7401cf9e76af0252622"
  },
  {
    "id": "396d52643794602484520360bd5682da7c02dcd3",
    "name": "README.md",
    "type": "blob",
    "path": "README.md",
    "mode": "100644",
    "submodule_link": null,
    "submodule_branch": null,
    "md5": "04c6e90faac2675aa89e2176d2eec7d8"
  },
  {
    "id": "efd261bf79519c997d1c2ac4154798d551f022dd",
    "name": "demo",
    "type": "blob",
    "path": "demo",
    "mode": "100644",
    "submodule_link": null,
    "submodule_branch": null,
    "md5": "fe01ce2a7fbac8fafaed7c982a04e229"
  },
  {
    "id": "83cd5cf8ae91ee1f40eb13f9e1352b8fba42bb08",
    "name": "240223demo1",
    "type": "commit",
    "path": "240223demo1",
    "mode": "160000",
    "submodule_link": "/xxxxx/240223demo1/tree/83cd5cf8ae91ee1f40eb13f9e1352b8fba42bb08",
    "submodule_branch": "dev",
    "md5": "e1d936a2bd619753065a3c15255f4a2a"
  }
]
```
