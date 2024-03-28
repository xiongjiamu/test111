---
linkTitle: Webhook API
title: WebHook 接口API文档
weight: 10
sidebar:
  open: false
---

## 1.列出仓库的WebHooks
### 请求 
`GET /api/v4/projects/{project_id}/hooks`

### 参数
| 参数名     |  位置    |   必选    | 类型    | 描述       |
| ---------- | ---------- |---------- |------- | ---------- |
| project_id |  path    |   Yes   | string  | 仓库 id    |
| page       |  query    |   NO   | integer | 分页页码 |
| per_page   |  query    |   NO   | integer | 分页数量 |


### 响应
```json
[
  {
        "id": 214246,
        "url": "http://example.com",
        "created_at": "2020-09-26T04:00:43.067+08:00",
        "push_events": true,
        "tag_push_events": true,
        "merge_requests_events": true,
        "project_id": 334507,
        "note_events": true,
        "token": "",
      	"active": true,
        "content_type": "application/json",
        "latest_log_status_code": 200
    }
]
```


## 2.创建一个仓库WebHook
### 请求 
`POST /api/v4/projects/{project_id}/hooks`

### 参数
| 参数名                | 类型    | 描述                 |
| --------------------- | ------- | -------------------- |
| project_id            | string  | 仓库 id              |
| created_at            | string  | 创建时间             |
| active                | boolean | 是否激活WebHook      |
| merge_requests_events | boolean | 是否关注合并请求事件 |
| note_events           | boolean | 是否关注评论事件     |
| url                   | string  | 远程HTTP URL         |
| token                 | string  | 认证Token            |
| push_events           | boolean | 是否关注推送事件     |
| tag_push_events       | boolean | 是否关注标签推送事件 |
| content_type          | string  | Body 格式            |


### 响应
```json
{
    "id": 166,
    "url": "https://example.com",
    "created_at": "2023-09-16T15:02:12.365+08:00",
    "push_events": true,
    "tag_push_events": true,
    "merge_requests_events": true,
    "project_id": 349013,
    "note_events": true,
    "token": "",
    "content_type": "application/json"
}
```


## 3.获取仓库单个WebHook
### 请求 
`GET /api/v4/projects/{project_id}/hooks/{hook_id}`
### 参数
| 参数名     | 类型    | 描述 |
| ---------- | ------- | ---- |
| project_id | string  |      |
| hook_id    | integer |      |


### 响应
```json
{
    "id": 166,
    "url": "https://example.com",
    "created_at": "2023-09-16T15:02:12.365+08:00",
    "push_events": true,
    "tag_push_events": true,
    "merge_requests_events": true,
    "project_id": 349013,
    "note_events": true,
    "token": "",
    "active": true,
    "content_type": "application/json"
}
```


## 4.删除一个仓库的WebHook
### 请求 
`DELETE /api/v4/projects/{project_id}/hooks/{hook_id}`
### 参数
| 参数名     | 类型    | 描述 |
| ---------- | ------- | ---- |
| project_id | string  |      |
| hook_id    | integer |      |

## 5.更新一个仓库的WebHook
### 请求 
`PUT /api/v4/projects/{project_id}/hooks/{hook_id}`
### 参数
| 参数名                | 类型    | 描述                 |
| --------------------- | ------- | -------------------- |
| project_id            | string  | 仓库 id              |
| created_at            | string  | 创建时间             |
| active                | boolean | 是否激活WebHook      |
| merge_requests_events | boolean | 是否关注合并请求事件 |
| note_events           | boolean | 是否关注评论事件     |
| url                   | string  | 远程HTTP URL         |
| token                 | string  | 认证Token            |
| push_events           | boolean | 是否关注推送事件     |
| tag_push_events       | boolean | 是否关注标签推送事件 |
| content_type          | string  | Body 格式            |
| id                    | integer | WebHook id |


### 响应
```json
{
    "id": 166,
    "url": "https://example.com",
    "created_at": "2023-09-16T15:02:12.365+08:00",
    "push_events": true,
    "tag_push_events": true,
    "merge_requests_events": true,
    "project_id": 349013,
    "note_events": true,
    "token": "",
    "content_type": "application/json"
}
```


## 6.测试WebHook是否发送成功
### 请求 
`POST /api/v4/projects/{project_id}/hooks/{hook_id}/test`
### 参数
| 参数名     | 类型    | 描述       |
| ---------- | ------- | ---------- |
| project_id | string  | 仓库 id    |
| hook_id    | integer | WebHook id |
| trigger    | string  | 触发的事件 |


### 响应
```json
{
    "message": "Hook executed successfully: HTTP 200"
}
```


## 7.回传PR流水线构建状态
### 请求 
`POST /api/v4/projects/{project_id}/statuses/{sha}`

### 参数
| 参数名            | 类型    | 描述         |
| ----------------- | ------- | ------------ |
| project_id        | string  | 仓库 id      |
| sha               | string  | commit id    |
| merge_request_id  | string  | 合并请求 id  |
| target_url        | string  | 跳转链接     |
| description       | string  | 描述信息     |
| merge_request_iid | integer | 合并请求 iid |
| ref               | string  | 分支名称     |
| stage             | string  | 阶段名称     |
| trigger_user      | string  | 触发用户     |
| build_id          | string  | 构建 id      |
| name              | string  | 任务名称     |
| state             | string  | 状态         |


### 响应
```json
{
    "id": 1025,
    "sha": "dd39828ef7ab863d7108d527c26d27f7c67708d5",
    "ref": "test",
    "status": "success",
    "name": "cloudpipeline",
    "target_url": "https://example.com",
    "created_at": "2023-09-01T15:47:55.810+08:00",
    "started_at": null,
    "finished_at": "2023-09-01T15:47:55.817+08:00",
    "allow_failure": false,
    "trigger_user": null,
    "updated_at": "2023-09-16T15:21:20.043+08:00",
    "rebuild_failed_url": null,
    "third_build_id": "2b5c17e12f8640a38ab609263724d595",
    "description": "",
    "author": {
        "id": 12,
        "name": "tom",
        "username": "tom",
        "nick_name": "tom",
        "state": "active",
        "avatar_url": "https://example.com",
        "avatar_path": null,
        "email": "tom@example.com",
        "name_cn": null,
        "web_url": "https://example.com",
        "tenant_name": null
    }
}
```


## 8.回传PR流水线质量报告
### 请求 
`POST /api/v4/projects/{project_id}/commits/{sha}/quality`
### 参数
| 参数名            | 类型    | 描述         |
| ----------------- | ------- | ------------ |
| project_id        | string  | 仓库 id      |
| sha               | string  | commit id    |
| merge_request_id  | integer | 合并请求 id  |
| ref               | string  | 分支名称     |
| stage             | string  | 阶段名称     |
| message           | string  | 提示信息     |
| sha               | string  | commit id    |
| url               | string  | 跳转链接     |
| quality_data      | object  | 质量报告     |
| status            | string  | 状态         |
| merge_request_iid | integer | 合并请求 iid |


### 响应
```json
{
    "pipeline_id": 5,
    "stage": "stage",
    "status": "success",
    "url": "",
    "message": null,
    "result": [
        {
            "time": "308",
            "tool": "codecheck",
            "result": [
                {
                    "real": "0",
                    "field": "critical",
                    "is_gray": false,
                    "expected": "<=0",
                    "field_url": null
                },
                {
                    "real": "0",
                    "field": "major",
                    "is_gray": false,
                    "expected": "<=0",
                    "field_url": null
                },
                {
                    "real": "0",
                    "field": "minor",
                    "is_gray": false,
                    "expected": "<=0",
                    "field_url": null
                },
                {
                    "real": "0",
                    "field": "info",
                    "is_gray": true,
                    "expected": "",
                    "field_url": null
                }
            ],
            "status": "failed",
            "log_url": "https://example.com",
            "indicator": "",
            "line_break_flag": false
        }
    ],
    "version": 987,
    "created_at": "2023-08-14T15:33:41.742+08:00",
    "updated_at": "2023-09-16T15:27:13.160+08:00"
}
```
