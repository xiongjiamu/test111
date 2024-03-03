---
linkTitle: CodeReview API
title: 代码评审及评论相关接口API文档
weight: 7
sidebar:
  open: false
---

## 1、获取PR 代码评审及动态列表

### 请求

`GET /api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/discussions`



### 参数:

|      参数名       | 是否必填 | 传参方式 | 类型    | 描述                                  |
| :---------------: | :------: | :------: | ------- | ------------------------------------- |
|    project_id     |    是    |   Path   | String  | 项目id或项目路径                      |
| merge_request_iid |    是    |   Path   | Integer | MR id                                 |
|      end_id       |    否    |  Query   | Integer | 上一页检视意见最大id（第一页默认为0） |
|   end_system_id   |    否    |  Query   | Integer | 上一页动态最大id（第一页默认为0）     |
|       page        |    否    |  Query   | Integer | 分页页数                              |
|     per_page      |    否    |  Query   | Integer | 分页页大小                            |

### 请求示例：

> /api/v4/projects/test%2Fproject1/merge_requests/1/discussions?page=1&per_page=20
>
> /api/v4/projects/1235/merge_requests/1/discussions?page=2&per_page=20&end_id=125&end_system_id=845



注：

> 此接口为顺序请求接口，只能从第一页开始滚动翻页，请求返回的接口数据中包含两种格式的返回体，
>
> 1、检视意见
>
> 2、系统动态
>
> 且每次请求都会返回当前页的end_id  和 end_system_id，下一页请求时需要传入上一次请求返回的数据，否则会导致分页数据错误。 



### 响应:

```json
{
  "end_id": 2337,
  "end_system_id": 1878,
  "data": [
    {
      "id": "ef2902d9bdccea82f7baf115f2557f446460372d",
      "individual_note": true,
      "notes": [
        {
          "id": 277,
          "type": null,
          "body": "创建一个不需要解决评审意见",
          "attachment": null,
          "author": {
            "id": 8,
            "name": "CSDN-liker-test3/csdndev_IDP",
            "username": "fb66cedc0d124f289a522af55549f028",
            "iam_id": null,
            "nick_name": "",
            "state": "active",
            "avatar_url": null,
            "avatar_path": null,
            "email": "",
            "name_cn": null,
            "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
            "tenant_name": null
          },
          "created_at": "2023-08-04T16:38:23.383+08:00",
          "updated_at": "2023-08-04T16:38:23.383+08:00",
          "system": false,
          "noteable_id": 7,
          "noteable_type": "MergeRequest",
          "commit_id": null,
          "resolvable": false,
          "is_reply": false,
          "resolved_by": null,
          "noteable_iid": null,
          "discussion_id": "ef2902d9bdccea82f7baf115f2557f446460372d",
          "project": "fb66cedc0d124f289a522af55549f028/yx_test",
          "diff_file": null,
          "diff": null,
          "archived": false,
          "review_categories": null,
          "review_categories_cn": "",
          "review_categories_en": "",
          "review_modules": null,
          "severity": "suggestion",
          "severity_cn": "建议",
          "severity_en": "Suggestion",
          "file_path": null,
          "line": null,
          "assignee": {
            "id": 8,
            "name": "CSDN-liker-test3/csdndev_IDP",
            "username": "fb66cedc0d124f289a522af55549f028",
            "iam_id": null,
            "nick_name": "",
            "state": "active",
            "avatar_url": null,
            "avatar_path": null,
            "email": "",
            "name_cn": null,
            "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
            "tenant_name": null
          },
          "proposer": {
            "id": 8,
            "name": "CSDN-liker-test3/csdndev_IDP",
            "username": "fb66cedc0d124f289a522af55549f028",
            "iam_id": null,
            "nick_name": "",
            "state": "active",
            "avatar_url": null,
            "avatar_path": null,
            "email": "",
            "name_cn": null,
            "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
            "tenant_name": null
          },
          "position": null,
          "resolved": null,
          "is_outdated": null,
          "issue": null
        },
        {
          "id": 279,
          "type": null,
          "body": "回复不需要解决的评审意见",
          "attachment": null,
          "author": {
            "id": 8,
            "name": "CSDN-liker-test3/csdndev_IDP",
            "username": "fb66cedc0d124f289a522af55549f028",
            "iam_id": null,
            "nick_name": "",
            "state": "active",
            "avatar_url": null,
            "avatar_path": null,
            "email": "",
            "name_cn": null,
            "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
            "tenant_name": null
          },
          "created_at": "2023-08-04T16:39:41.286+08:00",
          "updated_at": "2023-08-04T16:39:41.286+08:00",
          "system": false,
          "noteable_id": 7,
          "noteable_type": "MergeRequest",
          "commit_id": null,
          "resolvable": false,
          "is_reply": true,
          "resolved_by": null,
          "noteable_iid": null,
          "discussion_id": "ef2902d9bdccea82f7baf115f2557f446460372d",
          "project": "fb66cedc0d124f289a522af55549f028/yx_test",
          "diff_file": null,
          "diff": null,
          "archived": false,
          "review_categories": null,
          "review_categories_cn": "",
          "review_categories_en": "",
          "review_modules": null,
          "severity": "suggestion",
          "severity_cn": "建议",
          "severity_en": "Suggestion",
          "file_path": null,
          "line": null,
          "assignee": {
            "id": 8,
            "name": "CSDN-liker-test3/csdndev_IDP",
            "username": "fb66cedc0d124f289a522af55549f028",
            "iam_id": null,
            "nick_name": "",
            "state": "active",
            "avatar_url": null,
            "avatar_path": null,
            "email": "",
            "name_cn": null,
            "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
            "tenant_name": null
          },
          "proposer": {
            "id": 8,
            "name": "CSDN-liker-test3/csdndev_IDP",
            "username": "fb66cedc0d124f289a522af55549f028",
            "iam_id": null,
            "nick_name": "",
            "state": "active",
            "avatar_url": null,
            "avatar_path": null,
            "email": "",
            "name_cn": null,
            "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
            "tenant_name": null
          },
          "position": null,
          "resolved": null,
          "is_outdated": null,
          "issue": null
        }
      ],
      "project_id": 245,
      "noteable_type": "MergeRequest",
      "commit_id": null,
      "project_full_path": "fb66cedc0d124f289a522af55549f028/yx_test",
      "a_mode": null,
      "b_mode": null,
      "deleted_file": null,
      "new_file": null,
      "resolved": false,
      "archived": false,
      "review_categories": null,
      "review_categories_cn": "",
      "review_categories_en": "",
      "review_modules": null,
      "severity": "suggestion",
      "severity_cn": "建议",
      "severity_en": "Suggestion",
      "assignee": {
        "id": 8,
        "name": "CSDN-liker-test3/csdndev_IDP",
        "username": "fb66cedc0d124f289a522af55549f028",
        "iam_id": null,
        "nick_name": "",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": null,
        "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
        "tenant_name": null
      },
      "proposer": {
        "id": 8,
        "name": "CSDN-liker-test3/csdndev_IDP",
        "username": "fb66cedc0d124f289a522af55549f028",
        "iam_id": null,
        "nick_name": "",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": null,
        "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
        "tenant_name": null
      },
      "issue": null,
      "merge_request_version_params": null,
      "diff_file": null,
      "added_lines": null,
      "removed_lines": null
    },
    {
      "id": 334,
      "body": "Add reviewers: 7a7657aa0aa04f438981632f623c5e52,1f18db8aa5674f12af3cc4b5b9d60c15; ",
      "action": "mr_change",
      "author": {
        "id": 8,
        "name": "CSDN-liker-test3/csdndev_IDP",
        "username": "fb66cedc0d124f289a522af55549f028",
        "iam_id": null,
        "nick_name": "",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": null,
        "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
        "tenant_name": null
      },
      "created_at": "2023-08-07T16:08:38.004+08:00",
      "updated_at": "2023-08-07T16:08:38.004+08:00",
      "discussion_id": "ff66f5d39888f1b3d84fc98322857c33325d597a",
      "project": "fb66cedc0d124f289a522af55549f028/yx_test",
      "assignee": null,
      "proposer": null,
      "merge_request_id": 7
    }
  ]
}
```



## 2、创建MR评审意见

### 请求

`POST /api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/discussions`

### 参数

| 参数名            | 位置 | 必选 | 类型        | 描述                                 |
| ----------------- | ---- | ---- | ----------- | ------------------------------------ |
| project_id        | path | Yes  | string      | 项目id或项目路径                     |
| merge_request_iid | path | Yes  | integer     | MRiid                                |
| body              | body | NO   | string      | 评论内容                             |
| line_types        | body | NO   | string      | 评论代码行类型                       |
| need_to_resolve   | body | NO   | boolean     | 是否需要解决，若为true则会影响MR合入 |
| position          | body | NO   | PositionDto | 位置信息标记检视意见在代码中的位置   |
| severity          | body | YES  | string      | 严重等级，默认suggestion即可         |

PositionDto

| 参数名                   | 类型    | 描述                                          |
| ------------------------ | ------- | --------------------------------------------- |
| base_sha                 | string  | 文件基础sha值                                 |
| start_sha                | string  | 文件起始sha值，基于当前MR                     |
| head_sha                 | string  | 文件最新sha值，基于当前MR                     |
| position_type            | string  | 位置类型                                      |
| new_path                 | string  | 文件新路径                                    |
| old_path                 | string  | 文件旧路径（仅在文件移动时新旧路径不同）      |
| new_line                 | integer | 在最新提交中行数                              |
| old_line                 | integer | 在上一次提交中文件中行数，若旧文件不存在取 -1 |
| ignore_whitespace_change | boolean | 是否忽略空格变更                              |

### 请求示例：

创建一个待解决检视意见:

```json
{
    "body": "新建一个需要解决的评审意见",
    "need_to_resolve": true
}
```

创建一个文件页检视意见：

```json
{
    "body": "新建一个评审意见",
    "line_types": "new",
    "position": {
        "base_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
        "head_sha": "3726ebde5bbcc996fe20af164586678ec09d3d15",
        "start_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
        "position_type": "text",
        "new_path": "dev",
        "old_path": "dev",
        "new_line": 1,
        "old_line": -1,
        "ignore_whitespace_change": false
    },
    "severity": "suggestion"
}
```

### 响应：

```json
{
    "id": "765e58c6d85db93175c413c9e8b61d3908bebf17",
    "individual_note": false,
    "notes": [
        {
            "id": 17571,
            "type": "DiffNote",
            "body": "新建一个评审意见",
            "attachment": null,
            "author": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "created_at": "2023-09-20T16:28:30.267+08:00",
            "updated_at": "2023-09-20T16:28:30.267+08:00",
            "system": false,
            "noteable_id": 4000,
            "noteable_type": "MergeRequest",
            "commit_id": null,
            "resolvable": true,
            "is_reply": false,
            "resolved_by": null,
            "noteable_iid": 1,
            "discussion_id": "765e58c6d85db93175c413c9e8b61d3908bebf17",
            "project": "admini/minority",
            "diff_file": "dev",
            "diff": "@@ -0,0 +1,0 @@\n+dev\n",
            "archived": false,
            "review_categories": null,
            "review_categories_cn": "",
            "review_categories_en": "",
            "review_modules": null,
            "severity": "suggestion",
            "severity_cn": "建议",
            "severity_en": "Suggestion",
            "file_path": null,
            "line": null,
            "assignee": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "proposer": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "position": {
                "base_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
                "start_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
                "head_sha": "3726ebde5bbcc996fe20af164586678ec09d3d15",
                "old_path": "dev",
                "new_path": "dev",
                "position_type": "text",
                "old_line": null,
                "new_line": 1
            },
            "resolved": false,
            "is_outdated": false,
            "issue": null
        }
    ],
    "project_id": 179098,
    "noteable_type": "MergeRequest",
    "commit_id": null,
    "project_full_path": "admini/minority",
    "deleted_file": false,
    "new_file": true,
    "resolved": false,
    "archived": false,
    "review_categories": null,
    "review_categories_cn": "",
    "review_categories_en": "",
    "review_modules": null,
    "severity": "suggestion",
    "severity_cn": "建议",
    "severity_en": "Suggestion",
    "assignee": {
        "id": 107,
        "name": "admini",
        "username": "admini",
        "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "admini",
        "web_url": "https://gitcode.com/admini",
        "nick_name": "admini",
        "tenant_name": null
    },
    "proposer": {
        "id": 107,
        "name": "admini",
        "username": "admini",
        "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "admini",
        "web_url": "https://gitcode.com/admini",
        "nick_name": "admini",
        "tenant_name": null
    },
    "issue": null,
    "merge_request_version_params": {
        "diff_id": null,
        "start_sha": null,
        "commit_id": null
    },
    "diff_file": "@@ -0,0 +1,0 @@\n+dev\n",
    "added_lines": 1,
    "removed_lines": 0,
    "amode": null,
    "bmode": null
}
```

## 3、回复PR 代码评审或评论

### 请求：

`POST /api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/discussions/{discussion_id}/notes`

### 参数：

| 参数名            | 位置 | 必选 | 类型    | 描述                         |
| ----------------- | ---- | ---- | ------- | ---------------------------- |
| project_id        | path | Yes  | string  | 项目id或项目路径             |
| merge_request_iid | path | Yes  | integer | MRiid                        |
| discussion_id     | path | Yes  | string  | 评论id，创建评论时会返回此id |
| body              | body | NO   | string  | 回复内容                     |

### 请求示例：

```json
{
    "body": "回复一个评论"
}
```

### 响应：

```json
{
    "id": 17593,
    "type": "DiffNote",
    "body": "回复一个评论",
    "attachment": null,
    "author": {
        "id": 107,
        "name": "admini",
        "username": "admini",
        "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "admini",
        "web_url": "https://gitcode.com/admini",
        "nick_name": "admini",
        "tenant_name": null
    },
    "created_at": "2023-09-20T17:09:54.719+08:00",
    "updated_at": "2023-09-20T17:09:54.719+08:00",
    "system": false,
    "noteable_id": 4000,
    "noteable_type": "MergeRequest",
    "commit_id": null,
    "resolvable": true,
    "is_reply": true,
    "resolved_by": null,
    "noteable_iid": 1,
    "discussion_id": "765e58c6d85db93175c413c9e8b61d3908bebf17",
    "project": "admini/minority",
    "diff_file": "dev",
    "diff": "@@ -0,0 +1,0 @@\n+dev\n",
    "archived": false,
    "review_categories": null,
    "review_categories_cn": "",
    "review_categories_en": "",
    "review_modules": null,
    "severity": "suggestion",
    "severity_cn": "建议",
    "severity_en": "Suggestion",
    "file_path": null,
    "line": null,
    "assignee": {
        "id": 107,
        "name": "admini",
        "username": "admini",
        "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "admini",
        "web_url": "https://gitcode.com/admini",
        "nick_name": "admini",
        "tenant_name": null
    },
    "proposer": {
        "id": 107,
        "name": "admini",
        "username": "admini",
        "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "admini",
        "web_url": "https://gitcode.com/admini",
        "nick_name": "admini",
        "tenant_name": null
    },
    "position": {
        "base_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
        "start_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
        "head_sha": "3726ebde5bbcc996fe20af164586678ec09d3d15",
        "old_path": "dev",
        "new_path": "dev",
        "position_type": "text",
        "old_line": null,
        "new_line": 1
    },
    "resolved": false,
    "is_outdated": false,
    "issue": null
}
```



## 4、修改PR 代码评审或评论

### 请求：

`PUT /api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/discussions/{discussion_id}/notes/{note_id}`

### 参数：

| 参数名            | 位置 | 必选 | 类型    | 描述                                 |
| ----------------- | ---- | ---- | ------- | ------------------------------------ |
| project_id        | path | Yes  | string  | 项目id或项目路径                     |
| merge_request_iid | path | Yes  | integer | MRiid                                |
| discussion_id     | path | Yes  | string  | 评论id，主评论及回复此id一致         |
| note_id           | path | Yes  | integer | 评论记录id，每个评论或回复，此id唯一 |
| body              | body | NO   | string  | 评论或描述内容                       |
| serverity         | body | YES  | string  | 严重等级，默认取suggestion即可       |

### 请求示例：

```json
{
    "body": "修一个评审意见",
    "severity": "suggestion"
}
```

### 响应：

```json
{
    "id": 17571,
    "type": "DiffNote",
    "body": "修一个评审意见",
    "attachment": null,
    "author": {
        "id": 107,
        "name": "admini",
        "username": "admini",
        "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "admini",
        "web_url": "https://gitcode.com/admini",
        "nick_name": "admini",
        "tenant_name": null
    },
    "created_at": "2023-09-20T16:28:30.267+08:00",
    "updated_at": "2023-09-20T17:12:37.090+08:00",
    "system": false,
    "noteable_id": 4000,
    "noteable_type": "MergeRequest",
    "commit_id": null,
    "resolvable": true,
    "is_reply": false,
    "resolved_by": null,
    "noteable_iid": 1,
    "discussion_id": "765e58c6d85db93175c413c9e8b61d3908bebf17",
    "project": "admini/minority",
    "diff_file": "dev",
    "diff": "@@ -0,0 +1,0 @@\n+dev\n",
    "archived": false,
    "review_categories": null,
    "review_categories_cn": "",
    "review_categories_en": "",
    "review_modules": null,
    "severity": "suggestion",
    "severity_cn": "建议",
    "severity_en": "Suggestion",
    "file_path": null,
    "line": null,
    "assignee": {
        "id": 107,
        "name": "admini",
        "username": "admini",
        "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "admini",
        "web_url": "https://gitcode.com/admini",
        "nick_name": "admini",
        "tenant_name": null
    },
    "proposer": {
        "id": 107,
        "name": "admini",
        "username": "admini",
        "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "admini",
        "web_url": "https://gitcode.com/admini",
        "nick_name": "admini",
        "tenant_name": null
    },
    "position": {
        "base_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
        "start_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
        "head_sha": "3726ebde5bbcc996fe20af164586678ec09d3d15",
        "old_path": "dev",
        "new_path": "dev",
        "position_type": "text",
        "old_line": null,
        "new_line": 1
    },
    "resolved": false,
    "is_outdated": false,
    "issue": null
}
```

## 5、删除检视意见或评论

### 请求：

`DELETE /api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/discussions/{discussion_id}/notes/{note_id}`

### 参数：

| 参数名            | 位置 | 必选 | 类型    | 描述                                 |
| ----------------- | ---- | ---- | ------- | ------------------------------------ |
| project_id        | path | Yes  | string  | 项目id或项目路径                     |
| merge_request_iid | path | Yes  | integer | MRiid                                |
| discussion_id     | path | Yes  | string  | 评论id，主评论及回复此id一致         |
| note_id           | path | Yes  | integer | 评论记录id，每个评论或回复，此id唯一 |

### 请求示例：

``/api/v4/projects/admini%2Fminority/merge_requests/1/discussions/765e58c6d85db93175c413c9e8b61d3908bebf17`

### 响应：

无



## 6、解决或取消解决检视意见

### 请求：

`PUT /api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/discussions/{discussion_id}`

### 参数：

| 参数名            | 位置 | 必选 | 类型    | 描述                         |
| ----------------- | ---- | ---- | ------- | ---------------------------- |
| project_id        | path | Yes  | string  | 项目id或项目路径             |
| merge_request_iid | path | Yes  | integer | MRiid                        |
| discussion_id     | path | Yes  | string  | 评论id，主评论及回复此id一致 |
| resolved          | body | Yes  | boolean | 解决传true，取消解决传false  |

### 请求示例：

`/api/v4/projects/admini%2Fminority/merge_requests/1/discussions/765e58c6d85db93175c413c9e8b61d3908bebf17`

```json
{
    "resolved": true
}
```

### 响应：

```json
{
    "id": "765e58c6d85db93175c413c9e8b61d3908bebf17",
    "individual_note": false,
    "notes": [
        {
            "id": 17571,
            "type": "DiffNote",
            "body": "修一个评审意见",
            "attachment": null,
            "author": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "created_at": "2023-09-20T16:28:30.267+08:00",
            "updated_at": "2023-09-20T19:37:07.267+08:00",
            "system": false,
            "noteable_id": 4000,
            "noteable_type": "MergeRequest",
            "commit_id": null,
            "resolvable": true,
            "is_reply": false,
            "resolved_by": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "noteable_iid": 1,
            "discussion_id": "765e58c6d85db93175c413c9e8b61d3908bebf17",
            "project": "admini/minority",
            "diff_file": "dev",
            "diff": "@@ -0,0 +1,0 @@\n+dev\n",
            "archived": false,
            "review_categories": null,
            "review_categories_cn": "",
            "review_categories_en": "",
            "review_modules": null,
            "severity": "suggestion",
            "severity_cn": "建议",
            "severity_en": "Suggestion",
            "file_path": null,
            "line": null,
            "assignee": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "proposer": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "position": {
                "base_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
                "start_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
                "head_sha": "3726ebde5bbcc996fe20af164586678ec09d3d15",
                "old_path": "dev",
                "new_path": "dev",
                "position_type": "text",
                "old_line": null,
                "new_line": 1
            },
            "resolved": true,
            "is_outdated": false,
            "issue": null
        },
        {
            "id": 17593,
            "type": "DiffNote",
            "body": "回复一个评论",
            "attachment": null,
            "author": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "created_at": "2023-09-20T17:09:54.719+08:00",
            "updated_at": "2023-09-20T19:37:07.267+08:00",
            "system": false,
            "noteable_id": 4000,
            "noteable_type": "MergeRequest",
            "commit_id": null,
            "resolvable": true,
            "is_reply": true,
            "resolved_by": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "noteable_iid": 1,
            "discussion_id": "765e58c6d85db93175c413c9e8b61d3908bebf17",
            "project": "admini/minority",
            "diff_file": "dev",
            "diff": "@@ -0,0 +1,0 @@\n+dev\n",
            "archived": false,
            "review_categories": null,
            "review_categories_cn": "",
            "review_categories_en": "",
            "review_modules": null,
            "severity": "suggestion",
            "severity_cn": "建议",
            "severity_en": "Suggestion",
            "file_path": null,
            "line": null,
            "assignee": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "proposer": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "position": {
                "base_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
                "start_sha": "896b22be71b808f8471ef4d21890b3bc70f39b62",
                "head_sha": "3726ebde5bbcc996fe20af164586678ec09d3d15",
                "old_path": "dev",
                "new_path": "dev",
                "position_type": "text",
                "old_line": null,
                "new_line": 1
            },
            "resolved": true,
            "is_outdated": false,
            "issue": null
        }
    ],
    "project_id": 179098,
    "noteable_type": "MergeRequest",
    "commit_id": null,
    "project_full_path": "admini/minority",
    "deleted_file": false,
    "new_file": true,
    "resolved": true,
    "archived": false,
    "review_categories": null,
    "review_categories_cn": "",
    "review_categories_en": "",
    "review_modules": null,
    "severity": "suggestion",
    "severity_cn": "建议",
    "severity_en": "Suggestion",
    "assignee": {
        "id": 107,
        "name": "admini",
        "username": "admini",
        "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "admini",
        "web_url": "https://gitcode.com/admini",
        "nick_name": "admini",
        "tenant_name": null
    },
    "proposer": {
        "id": 107,
        "name": "admini",
        "username": "admini",
        "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "admini",
        "web_url": "https://gitcode.com/admini",
        "nick_name": "admini",
        "tenant_name": null
    },
    "issue": null,
    "merge_request_version_params": {
        "diff_id": null,
        "start_sha": null,
        "commit_id": null
    },
    "diff_file": "@@ -0,0 +1,0 @@\n+dev\n",
    "added_lines": 1,
    "removed_lines": 0,
    "amode": null,
    "bmode": null
}
```

## 7、获取commit评论列表

### 请求：

`GET /api/v4/projects/{project_id}/repository/commits/{commit_id}/discussions`

### 参数：

| 参数名     | 位置  | 必选 | 类型    | 描述             |
| ---------- | ----- | ---- | ------- | ---------------- |
| project_id | path  | Yes  | string  | 项目id或项目路径 |
| commit_id  | path  | Yes  | string  | commit_id        |
| page       | query | NO   | integer | 分页页码         |
| per_page   | query | NO   | integer | 分页大小         |

### 请求示例：

`/api/v4/projects/admini%2Fminority/repository/commits/a10515ee66bc2a31940963d2cceb01ba981c9bbc/discussions`

### 响应：

```json
[
    {
        "id": "23e972408db76d6ed0e63b45d06f07238989b878",
        "individual_note": false,
        "notes": [
            {
                "id": 367,
                "noteable_type": "Commit",
                "commit_id": "747155b8fe6fe68006f6c064ffe5b0e2f935f07a",
                "discussion_id": "23e972408db76d6ed0e63b45d06f07238989b878",
                "project": "fb66cedc0d124f289a522af55549f028/yx_test",
                "type": "DiffNote",
                "body": "添加一个commit页的评论22",
                "diff_file": "dev",
                "diff": "@@ -0,0 +1,0 @@\n+抱�u�]\n",
                "attachment": null,
                "author": {
                    "id": 8,
                    "name": "CSDN-liker-test3/csdndev_IDP",
                    "username": "fb66cedc0d124f289a522af55549f028",
                    "iam_id": null,
                    "nick_name": "",
                    "state": "active",
                    "avatar_url": null,
                    "avatar_path": null,
                    "email": "",
                    "name_cn": null,
                    "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
                    "tenant_name": null
                },
                "created_at": "2023-08-07T16:57:29.282+08:00",
                "updated_at": "2023-08-07T17:13:44.012+08:00",
                "system": false,
                "noteable_id": null,
                "position": {
                    "base_sha": "02715cfe49ddc91ed45e57f3a829a66882e2dc7e",
                    "start_sha": "02715cfe49ddc91ed45e57f3a829a66882e2dc7e",
                    "head_sha": "747155b8fe6fe68006f6c064ffe5b0e2f935f07a",
                    "old_path": "dev",
                    "new_path": "dev",
                    "position_type": "text",
                    "old_line": null,
                    "new_line": 1
                },
                "resolvable": true,
                "resolved": true,
                "resolved_by": {
                    "id": 8,
                    "name": "CSDN-liker-test3/csdndev_IDP",
                    "username": "fb66cedc0d124f289a522af55549f028",
                    "iam_id": null,
                    "nick_name": "",
                    "state": "active",
                    "avatar_url": null,
                    "avatar_path": null,
                    "email": "",
                    "name_cn": null,
                    "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
                    "tenant_name": null
                },
                "archived": false,
                "noteable_iid": null,
                "review_categories": null,
                "review_categories_cn": "",
                "review_categories_en": "",
                "review_modules": null,
                "severity": "suggestion",
                "severity_cn": "建议",
                "severity_en": "Suggestion",
                "file_path": null,
                "line": null,
                "assignee": null,
                "proposer": {
                    "id": 8,
                    "name": "CSDN-liker-test3/csdndev_IDP",
                    "username": "fb66cedc0d124f289a522af55549f028",
                    "iam_id": null,
                    "nick_name": "",
                    "state": "active",
                    "avatar_url": null,
                    "avatar_path": null,
                    "email": "",
                    "name_cn": null,
                    "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
                    "tenant_name": null
                },
                "issue": null,
                "is_reply": false
            }
        ],
        "project_id": 245,
        "noteable_type": "Commit",
        "commit_id": "747155b8fe6fe68006f6c064ffe5b0e2f935f07a",
        "project_full_path": "fb66cedc0d124f289a522af55549f028/yx_test",
        "a_mode": "0",
        "b_mode": "100644",
        "deleted_file": false,
        "new_file": true,
        "diff_file": "@@ -0,0 +1,0 @@\n+抱�u�]\n",
        "added_lines": 1,
        "removed_lines": 0,
        "resolved": true,
        "archived": false,
        "review_categories": null,
        "review_categories_cn": "",
        "review_categories_en": "",
        "review_modules": null,
        "severity": "suggestion",
        "severity_cn": "建议",
        "severity_en": "Suggestion",
        "assignee": null,
        "proposer": {
            "id": 8,
            "name": "CSDN-liker-test3/csdndev_IDP",
            "username": "fb66cedc0d124f289a522af55549f028",
            "iam_id": null,
            "nick_name": "",
            "state": "active",
            "avatar_url": null,
            "avatar_path": null,
            "email": "",
            "name_cn": null,
            "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
            "tenant_name": null
        },
        "issue": null,
        "merge_request_version_params": null
    },
    {
        "id": "431b31a2f17be82f53b4a6245908643c4cdbbc9a",
        "individual_note": true,
        "notes": [
            {
                "id": 469,
                "noteable_type": "Commit",
                "commit_id": "747155b8fe6fe68006f6c064ffe5b0e2f935f07a",
                "discussion_id": "431b31a2f17be82f53b4a6245908643c4cdbbc9a",
                "project": "fb66cedc0d124f289a522af55549f028/yx_test",
                "type": null,
                "body": "给这个commit添加一个评论1",
                "diff_file": null,
                "diff": "",
                "attachment": null,
                "author": {
                    "id": 8,
                    "name": "CSDN-liker-test3/csdndev_IDP",
                    "username": "fb66cedc0d124f289a522af55549f028",
                    "iam_id": null,
                    "nick_name": "",
                    "state": "active",
                    "avatar_url": null,
                    "avatar_path": null,
                    "email": "",
                    "name_cn": null,
                    "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
                    "tenant_name": null
                },
                "created_at": "2023-08-07T17:15:21.914+08:00",
                "updated_at": "2023-08-07T17:15:21.914+08:00",
                "system": false,
                "noteable_id": null,
                "position": null,
                "resolvable": false,
                "resolved": null,
                "resolved_by": null,
                "archived": false,
                "noteable_iid": null,
                "review_categories": null,
                "review_categories_cn": "",
                "review_categories_en": "",
                "review_modules": null,
                "severity": "suggestion",
                "severity_cn": "建议",
                "severity_en": "Suggestion",
                "file_path": null,
                "line": null,
                "assignee": null,
                "proposer": {
                    "id": 8,
                    "name": "CSDN-liker-test3/csdndev_IDP",
                    "username": "fb66cedc0d124f289a522af55549f028",
                    "iam_id": null,
                    "nick_name": "",
                    "state": "active",
                    "avatar_url": null,
                    "avatar_path": null,
                    "email": "",
                    "name_cn": null,
                    "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
                    "tenant_name": null
                },
                "issue": null,
                "is_reply": false
            }
        ],
        "project_id": 245,
        "noteable_type": "Commit",
        "commit_id": "747155b8fe6fe68006f6c064ffe5b0e2f935f07a",
        "project_full_path": "fb66cedc0d124f289a522af55549f028/yx_test",
        "a_mode": null,
        "b_mode": null,
        "deleted_file": null,
        "new_file": null,
        "diff_file": "",
        "added_lines": null,
        "removed_lines": null,
        "resolved": null,
        "archived": false,
        "review_categories": null,
        "review_categories_cn": "",
        "review_categories_en": "",
        "review_modules": null,
        "severity": "suggestion",
        "severity_cn": "建议",
        "severity_en": "Suggestion",
        "assignee": null,
        "proposer": {
            "id": 8,
            "name": "CSDN-liker-test3/csdndev_IDP",
            "username": "fb66cedc0d124f289a522af55549f028",
            "iam_id": null,
            "nick_name": "",
            "state": "active",
            "avatar_url": null,
            "avatar_path": null,
            "email": "",
            "name_cn": null,
            "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
            "tenant_name": null
        },
        "issue": null,
        "merge_request_version_params": null
    }
]
```

## 8、创建Issue评论

### 请求：

`POST /api/v4/projects/{project_id}/issues/{issue_iid}/discussions`

### 参数：

| 参数名     | 位置 | 必选 | 类型    | 描述             |
| ---------- | ---- | ---- | ------- | ---------------- |
| project_id | path | Yes  | string  | 项目id或项目路径 |
| issue_iid  | path | Yes  | integer | issue_iid        |
| body       | body | Yes  | string  | 评论本体         |

### 请求示例：

`/api/v4/projects/admini%2Fminority/issues/1/discussions`

```json
{
    "body": "这是一个Issue评论"
}
```



### 响应：

```json
{
    "id": "0b6e276df8046b87e3421cf5b64b681b8cf50751",
    "individual_note": false,
    "notes": [
        {
            "id": 17682,
            "type": "DiscussionNote",
            "body": "这是一个Issue评论",
            "attachment": null,
            "author": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            },
            "created_at": "2023-09-20T19:47:00.529+08:00",
            "updated_at": "2023-09-20T19:47:00.529+08:00",
            "system": false,
            "noteable_id": 6106,
            "noteable_type": "Issue",
            "commit_id": null,
            "resolvable": false,
            "is_reply": false,
            "resolved_by": null,
            "noteable_iid": 1,
            "discussion_id": "0b6e276df8046b87e3421cf5b64b681b8cf50751",
            "project": "admini/minority",
            "diff_file": null,
            "diff": "",
            "archived": false,
            "review_categories": null,
            "review_categories_cn": "",
            "review_categories_en": "",
            "review_modules": null,
            "severity": "suggestion",
            "severity_cn": "建议",
            "severity_en": "Suggestion",
            "file_path": null,
            "line": null,
            "assignee": null,
            "proposer": {
                "id": 107,
                "name": "admini",
                "username": "admini",
                "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
                "state": "active",
                "avatar_url": null,
                "avatar_path": null,
                "email": "",
                "name_cn": "admini",
                "web_url": "https://gitcode.com/admini",
                "nick_name": "admini",
                "tenant_name": null
            }
        }
    ],
    "project_id": 179098,
    "noteable_type": "Issue",
    "commit_id": null,
    "project_full_path": "admini/minority",
    "deleted_file": null,
    "new_file": null,
    "diff_file": null,
    "added_lines": null,
    "removed_lines": null,
    "resolved": false,
    "archived": false,
    "review_categories": null,
    "review_categories_cn": "",
    "review_categories_en": "",
    "review_modules": null,
    "severity": "suggestion",
    "severity_cn": "建议",
    "severity_en": "Suggestion",
    "assignee": null,
    "proposer": {
        "id": 107,
        "name": "admini",
        "username": "admini",
        "iam_id": "c2b0a8285bec42dc93263adb2a870f9e",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "admini",
        "web_url": "https://gitcode.com/admini",
        "nick_name": "admini",
        "tenant_name": null
    },
    "issue": null,
    "merge_request_version_params": null,
    "amode": null,
    "bmode": null
}
```

## 9、回复Issue评论

### 请求：

`POST /api/v4/projects/{project_id}/issues/{issue_iid}/discussions/{discussion_id}/notes`

### 参数：

| 参数名        | 位置 | 必选 | 类型    | 描述             |
| ------------- | ---- | ---- | ------- | ---------------- |
| project_id    | path | Yes  | string  | 项目id或项目路径 |
| issue_iid     | path | Yes  | integer | issue_iid        |
| discussion_id | path | Yes  | string  | 评论id           |
| body          | body | Yes  | string  | 评论内容         |

### 请求示例：

`/api/v4/projects/admini%2Fminority/issues/1/discussions/3dfe8997bd3dbbc04f6f49319b6ae617c3c9e2a8/notes`

```json
{
    "body": "回复一个Issue评论"
}
```

### 响应：

```json
{
    "id": 22389,
    "type": "DiscussionNote",
    "body": "回复一个Issue评论",
    "attachment": null,
    "author": {
        "id": 8,
        "name": "CSDN-liker-test3/csdndev_IDP",
        "username": "fb66cedc0d124f289a522af55549f028",
        "iam_id": null,
        "nick_name": "",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": null,
        "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
        "tenant_name": null
    },
    "created_at": "2023-09-20T19:54:35.688+08:00",
    "updated_at": "2023-09-20T19:54:35.688+08:00",
    "system": false,
    "noteable_id": 333,
    "noteable_type": "Issue",
    "commit_id": null,
    "resolvable": false,
    "is_reply": true,
    "resolved_by": null,
    "noteable_iid": 1,
    "discussion_id": "3dfe8997bd3dbbc04f6f49319b6ae617c3c9e2a8",
    "project": "fb66cedc0d124f289a522af55549f028/yx_test",
    "diff_file": null,
    "diff": "",
    "archived": false,
    "review_categories": null,
    "review_categories_cn": "",
    "review_categories_en": "",
    "review_modules": null,
    "severity": "suggestion",
    "severity_cn": "建议",
    "severity_en": "Suggestion",
    "file_path": null,
    "line": null,
    "assignee": null,
    "proposer": {
        "id": 8,
        "name": "CSDN-liker-test3/csdndev_IDP",
        "username": "fb66cedc0d124f289a522af55549f028",
        "iam_id": null,
        "nick_name": "",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": null,
        "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
        "tenant_name": null
    }
}
```

## 10、修改Issue评论

### 请求：

`PUT /api/v4/projects/{project_id}/issues/{issue_iid}/notes/{note_id}`

### 参数：

| 参数名     | 位置 | 必选 | 类型    | 描述         |
| ---------- | ---- | ---- | ------- | ------------ |
| project_id | path | Yes  | string  | 项目ID       |
| issue_iid  | path | Yes  | integer | Issue_iid    |
| note_id    | path | Yes  | string  | 记录id       |
| body       | body | Yes  | string  | 修改评论内容 |

### 请求示例：

`/api/v4/projects/245/issues/1/notes/22388`

```json
{
    "body": "这是一个修改后的评论"
}
```



### 响应：

```json
{
    "id": 22388,
    "type": "DiscussionNote",
    "body": "这是一个修改后的评论",
    "attachment": null,
    "author": {
        "id": 8,
        "name": "CSDN-liker-test3/csdndev_IDP",
        "username": "fb66cedc0d124f289a522af55549f028",
        "iam_id": null,
        "nick_name": "",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": null,
        "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
        "tenant_name": null
    },
    "created_at": "2023-09-20T19:53:38.863+08:00",
    "updated_at": "2023-09-20T19:59:15.711+08:00",
    "system": false,
    "noteable_id": 333,
    "noteable_type": "Issue",
    "commit_id": null,
    "resolvable": false,
    "is_reply": false,
    "resolved_by": null,
    "noteable_iid": 1,
    "discussion_id": "3dfe8997bd3dbbc04f6f49319b6ae617c3c9e2a8",
    "project": "fb66cedc0d124f289a522af55549f028/yx_test",
    "diff_file": null,
    "diff": "",
    "archived": false,
    "review_categories": null,
    "review_categories_cn": "",
    "review_categories_en": "",
    "review_modules": null,
    "severity": "suggestion",
    "severity_cn": "建议",
    "severity_en": "Suggestion",
    "file_path": null,
    "line": null,
    "assignee": null,
    "proposer": {
        "id": 8,
        "name": "CSDN-liker-test3/csdndev_IDP",
        "username": "fb66cedc0d124f289a522af55549f028",
        "iam_id": null,
        "nick_name": "",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": null,
        "web_url": "https://test.gitcode.net/fb66cedc0d124f289a522af55549f028",
        "tenant_name": null
    }
}
```



## 11、删除Issue评论

### 请求：

`GET /api/v4/projects/{project_id}/issues/{issue_iid}/notes/{note_id}`

### 参数：

| 参数名     | 位置 | 必选 | 类型    | 描述      |
| ---------- | ---- | ---- | ------- | --------- |
| project_id | path | Yes  | string  | 项目ID    |
| issue_iid  | path | Yes  | integer | Issue_iid |
| note_id    | path | Yes  | string  | 记录id    |

### 请求示例：

`/api/v4/projects/245/issues/1/notes/22388`

### 响应：

无