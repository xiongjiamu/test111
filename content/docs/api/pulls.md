---
linkTitle: Pull Requests API
title: PR接口API文档
weight: 8
sidebar:
  open: false
---

## 1、获取个人PR列表

### 请求

`GET /api/v4/merge_requests`



### 参数:

|  参数名  | 传参方式 | 必填 | 类型    | 描述                                                         |
| :------: | :------: | ---- | ------- | ------------------------------------------------------------ |
|  state   |  Query   | 否   | String  | PR 状态（opened、closed、locked、merged、all）                |
|   view   |  Query   | 否   | String  | 请求返回结果视图（simple、basic）                            |
|  scope   |  Query   | 否   | String  | PR 分类范围（created_by_me、assigned_to_me、need_my_review、all） |
| order_by |  Query   | 否   | String  | 排序字段（created_at、updated_at）                           |
|   sort   |  Query   | 否   | String  | 升降序（asc、desc）                                          |
| per_page |  Query   | 否   | Integer | 分页大小                                                     |
|   page   |  Query   | 否   | Integer | 分页页码                                                     |

### 请求示例：

> /api/v4/merge_requests?view=basic&page=1&per_page=20
>



### 响应:

**view=simple**

```json
[
    {
        "id": 8316,
        "iid": 7608,
        "project_id": 1467,
        "title": "APITest_PR_Branch_87bDW7Fego",
        "description": null,
        "state": "closed",
        "created_at": "2023-09-20T17:15:00.568+08:00",
        "updated_at": "2023-09-20T17:15:25.655+08:00",
        "web_url": "https://test.gitcode.net/Opensource_Group_MR/createMergeRequest_001/merge_requests/7608",
        "merge_request_type": "MergeRequest"
    }
]
```

**view=basic**

```json
[
    {
        "id": 8316,
        "iid": 7608,
        "title": "APITest_MR_Branch_87bDW7Fego",
        "source_branch": "Branch_87bDW7Fego",
        "target_branch": "main",
        "state": "closed",
        "created_at": "2023-09-20T17:15:00.568+08:00",
        "updated_at": "2023-09-20T17:15:25.655+08:00",
        "source_project_id": 1467,
        "review_mode": "approval",
        "author": {
            "id": 186,
            "name": "gitcode_boce1",
            "username": "gitcode_boce1",
            "iam_id": "3b6a7585f48d4d5c82a1584d2733b957",
            "nick_name": "",
            "state": "active",
            "avatar_url": null,
            "avatar_path": null,
            "email": "",
            "name_cn": "gitcode_boce1",
            "web_url": "https://test.gitcode.net/gitcode_boce1",
            "tenant_name": null
        },
        "closed_at": "2023-09-20T17:15:25.454+08:00",
        "closed_by": {
            "id": 186,
            "name": "gitcode_boce1",
            "username": "gitcode_boce1",
            "iam_id": "3b6a7585f48d4d5c82a1584d2733b957",
            "nick_name": "",
            "state": "active",
            "avatar_url": null,
            "avatar_path": null,
            "email": "",
            "name_cn": "gitcode_boce1",
            "web_url": "https://test.gitcode.net/gitcode_boce1",
            "tenant_name": null
        },
        "merged_at": null,
        "merged_by": null,
        "pipeline_status": "",
        "codequality_status": null,
        "pipeline_status_with_code_quality": "",
        "notes": 0,
        "source_project": {
            "id": 1467,
            "description": null,
            "name": "createMergeRequest_001",
            "name_with_namespace": "Opensource_Group_MR / createMergeRequest_001",
            "path": "createMergeRequest_001",
            "path_with_namespace": "Opensource_Group_MR/createMergeRequest_001",
            "develop_mode": null,
            "created_at": "2023-08-24T21:18:36.711+08:00",
            "updated_at": "2023-08-24T21:18:36.711+08:00",
            "archived": false,
            "is_kia": false,
            "ssh_url_to_repo": "ssh://git@test.gitcode.net:2222/Opensource_Group_MR/createMergeRequest_001.git",
            "http_url_to_repo": "https://test.gitcode.net/Opensource_Group_MR/createMergeRequest_001.git",
            "web_url": null,
            "readme_url": null,
            "product_id": null,
            "product_name": null
        },
        "target_project": {
            "id": 1467,
            "description": null,
            "name": "createMergeRequest_001",
            "name_with_namespace": "Opensource_Group_MR / createMergeRequest_001",
            "path": "createMergeRequest_001",
            "path_with_namespace": "Opensource_Group_MR/createMergeRequest_001",
            "develop_mode": null,
            "created_at": "2023-08-24T21:18:36.711+08:00",
            "updated_at": "2023-08-24T21:18:36.711+08:00",
            "archived": false,
            "is_kia": false,
            "ssh_url_to_repo": "ssh://git@test.gitcode.net:2222/Opensource_Group_MR/createMergeRequest_001.git",
            "http_url_to_repo": "https://test.gitcode.net/Opensource_Group_MR/createMergeRequest_001.git",
            "web_url": null,
            "readme_url": null,
            "product_id": null,
            "product_name": null
        },
        "web_url": "https://test.gitcode.net/Opensource_Group_MR/createMergeRequest_001/merge_requests/7608",
        "added_lines": 1,
        "removed_lines": 0,
        "merge_request_type": "MergeRequest",
        "source_git_url": "ssh://git@test.gitcode.net:2222/Opensource_Group_MR/createMergeRequest_001.git",
        "labels": [],
        "score": 0,
        "min_merged_score": 0,
        "source_product_id": null,
        "target_product_id": null,
        "product_name": null,
        "notes_count": {
            "notes_count": 0,
            "unresolved_notes_count": 0,
            "already_resolved_count": 0,
            "need_resolved_count": 0
        },
        "milestone": null,
        "e2e_issues": []
    }
]
```



## 2、获取项目PR列表

### 请求

`GET /api/v4/projects/{project_id}/gitcode/merge_requests`



### 参数:

|        参数名        | 传参方式 | 类型    | 描述                                          |
| :------------------: | :------: | ------- | --------------------------------------------- |
|      project_id      |   path   | String  | 项目ID或项目路径                              |
|        state         |  Query   | String  | PR状态（opened、closed、locked、merged、all） |
|         view         |  Query   | String  | 请求返回结果视图（simple、basic）             |
|        labels        |  Query   | String  | 根据label筛选MR                               |
|      author_id       |  Query   | Integer | 根据创建者的用户ID筛选PR                     |
| approval_reviewer_id |  Query   | Integer | 根据评审人的用户ID筛选PR                      |
|     assignees_id     |  Query   | Integer | 根据负责人的用户ID筛选PR                      |
|       order_by       |  Query   | String  | 排序字段（created_at、updated_at）            |
|         sort         |  Query   | String  | 升降序（asc、desc）                           |
|       per_page       |  Query   | Integer | 分页大小                                      |
|         page         |  Query   | Integer | 分页页码                                      |
|      only_count      |  Query   | Boolean | 是否只返回数量                                |

### 请求示例：

> /api/v4/projects/{project_id}/gitcode/merge_requests?view=basic&page=1&per_page=20
>
> /api/v4/projects/{project_id}/gitcode/merge_requests?only_count=true



### 响应:

**only_count=true**

```Json 
{
    "all": 1,
    "opened": 1,
    "closed": 0,
    "merged": 0,
    "locked": 0
}
```

**view=simple**

```json
[
    {
        "id": 1,
        "iid": 1,
        "project_id": 6,
        "title": "test123",
        "description": "新建文件",
        "state": "opened",
        "created_at": "2023-07-28T11:46:47.088+08:00",
        "updated_at": "2023-07-28T11:46:48.637+08:00",
        "web_url": "https://test.gitcode.net/7a7657aa0aa04f438981632f623c5e52/test06/merge_requests/1",
        "merge_request_type": "MergeRequest"
    }
]
```

**view=basic**

```json
[
    {
        "id": 1,
        "iid": 1,
        "title": "test123",
        "source_branch": "dev",
        "target_branch": "master",
        "state": "opened",
        "created_at": "2023-07-28T11:46:47.088+08:00",
        "updated_at": "2023-07-28T11:46:48.637+08:00",
        "source_project_id": 6,
        "review_mode": "approval",
        "author": {
            "id": 4,
            "name": "yanlp1",
            "username": "7a7657aa0aa04f438981632f623c5e52",
            "iam_id": null,
            "nick_name": "",
            "state": "active",
            "avatar_url": null,
            "avatar_path": null,
            "email": "",
            "name_cn": null,
            "web_url": "https://test.gitcode.net/7a7657aa0aa04f438981632f623c5e52",
            "tenant_name": null
        },
        "closed_at": null,
        "closed_by": null,
        "merged_at": null,
        "merged_by": null,
        "pipeline_status": "",
        "codequality_status": null,
        "pipeline_status_with_code_quality": "",
        "notes": 0,
        "source_project": {
            "id": 6,
            "description": "likun test desc",
            "name": "likun_test06",
            "name_with_namespace": "7a7657aa0aa04f438981632f623c5e52 / likun_test06",
            "path": "test06",
            "path_with_namespace": "7a7657aa0aa04f438981632f623c5e52/test06",
            "develop_mode": null,
            "created_at": "2023-07-25T11:54:30.118+08:00",
            "updated_at": "2023-07-25T11:54:30.118+08:00",
            "archived": false,
            "is_kia": false,
            "ssh_url_to_repo": "ssh://git@test.gitcode.net:2222/7a7657aa0aa04f438981632f623c5e52/test06.git",
            "http_url_to_repo": "https://test.gitcode.net/7a7657aa0aa04f438981632f623c5e52/test06.git",
            "web_url": null,
            "readme_url": null,
            "product_id": null,
            "product_name": null
        },
        "target_project": {
            "id": 6,
            "description": "likun test desc",
            "name": "likun_test06",
            "name_with_namespace": "7a7657aa0aa04f438981632f623c5e52 / likun_test06",
            "path": "test06",
            "path_with_namespace": "7a7657aa0aa04f438981632f623c5e52/test06",
            "develop_mode": null,
            "created_at": "2023-07-25T11:54:30.118+08:00",
            "updated_at": "2023-07-25T11:54:30.118+08:00",
            "archived": false,
            "is_kia": false,
            "ssh_url_to_repo": "ssh://git@test.gitcode.net:2222/7a7657aa0aa04f438981632f623c5e52/test06.git",
            "http_url_to_repo": "https://test.gitcode.net/7a7657aa0aa04f438981632f623c5e52/test06.git",
            "web_url": null,
            "readme_url": null,
            "product_id": null,
            "product_name": null
        },
        "web_url": "https://test.gitcode.net/7a7657aa0aa04f438981632f623c5e52/test06/merge_requests/1",
        "added_lines": 1,
        "removed_lines": 0,
        "merge_request_type": "MergeRequest",
        "source_git_url": "ssh://git@test.gitcode.net:2222/7a7657aa0aa04f438981632f623c5e52/test06.git",
        "labels": [],
        "score": 0,
        "min_merged_score": 0,
        "source_product_id": null,
        "target_product_id": null,
        "product_name": null,
        "notes_count": {
            "notes_count": 0,
            "unresolved_notes_count": 0,
            "already_resolved_count": 0,
            "need_resolved_count": 0
        },
        "milestone": null,
        "e2e_issues": []
    }
]
```

## 3、获取组织PR列表

### 请求

`GET /api/v4/groups/{group_id}/merge_requests`



### 参数:

|  参数名  | 传参方式 | 类型    | 描述                                          |
| :------: | :------: | ------- | --------------------------------------------- |
| group_id |   Path   | String  | 组织ID或组织路径                              |
|  state   |  Query   | String  | PR状态（opened、closed、locked、merged、all） |
|   view   |  Query   | String  | 请求返回结果视图（simple、basic）             |
| order_by |  Query   | String  | 排序字段（created_at、updated_at）            |
|   sort   |  Query   | String  | 升降序（asc、desc）                           |
| per_page |  Query   | Integer | 分页大小                                      |
|   page   |  Query   | Integer | 分页页码                                      |

### 请求示例：

> /api/v4/groups/{group_id}/merge_requests?view=basic&page=1&per_page=20
>
> /api/v4/groups/{group_id}/merge_requests?only_count=true



### 响应:

**only_count=true**

```json
{
    "all": 1,
    "opened": 1,
    "closed": 0,
    "merged": 0,
    "locked": 0
}
```

**view=simple**

```json
{
    "all": 29,
    "opened": 16,
    "merged": 9,
    "closed": 4,
    "locked": 0,
    "merge_requests": [
        {
            "id": 6551,
            "iid": 30,
            "project_id": 126,
            "title": "squash 合并",
            "description": "add\nadd\nadd:dwfewfewf\nperf:ewfewfewfewfew\nadd:cewfew口红就哦if热\n...",
            "state": "opened",
            "created_at": "2023-09-14T17:17:45.866+08:00",
            "updated_at": "2023-09-14T18:03:10.830+08:00",
            "web_url": "https://test.gitcode.net/t-qinmh/t-code-1/merge_requests/30",
            "merge_request_type": "MergeRequest"
        }
    ]
}
```

**view=basic**

```json
{
    "all": 29,
    "opened": 16,
    "merged": 9,
    "closed": 4,
    "locked": 0,
    "merge_requests": [
        {
            "id": 6551,
            "iid": 30,
            "title": "squash 合并",
            "source_branch": "empty-913-1",
            "target_branch": "empty-4",
            "state": "opened",
            "created_at": "2023-09-14T17:17:45.866+08:00",
            "updated_at": "2023-09-14T18:03:10.830+08:00",
            "source_project_id": 126,
            "review_mode": "approval",
            "author": {
                "id": 35,
                "name": "qinmh2",
                "username": "qinmh2",
                "iam_id": "2ec6da406e894bcc910a039dfac01f64",
                "nick_name": "打卡时间dswgrsegklpgvjwepioafhusfpenfpswefewfe",
                "state": "active",
                "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/fd/ab/41fdabab80765286aca56c2d33170525de83517738db2a3a95395f66912f3d1d.jpg?time=1693976634636",
                "avatar_path": null,
                "email": "",
                "name_cn": "qinmh2",
                "web_url": "https://test.gitcode.net/qinmh2",
                "tenant_name": null
            },
            "closed_at": null,
            "closed_by": null,
            "merged_at": null,
            "merged_by": null,
            "pipeline_status": "",
            "codequality_status": null,
            "pipeline_status_with_code_quality": "",
            "notes": 0,
            "source_project": {
                "id": 126,
                "description": null,
                "name": "t-code-1",
                "name_with_namespace": "t-qinmh-1 / t-code-1",
                "path": "t-code-1",
                "path_with_namespace": "t-qinmh/t-code-1",
                "develop_mode": null,
                "created_at": "2023-07-31T17:27:43.296+08:00",
                "updated_at": "2023-08-26T23:11:17.159+08:00",
                "archived": false,
                "is_kia": false,
                "ssh_url_to_repo": "ssh://git@test.gitcode.net:2222/t-qinmh/t-code-1.git",
                "http_url_to_repo": "https://test.gitcode.net/t-qinmh/t-code-1.git",
                "web_url": null,
                "readme_url": null,
                "product_id": null,
                "product_name": null
            },
            "target_project": {
                "id": 126,
                "description": null,
                "name": "t-code-1",
                "name_with_namespace": "t-qinmh-1 / t-code-1",
                "path": "t-code-1",
                "path_with_namespace": "t-qinmh/t-code-1",
                "develop_mode": null,
                "created_at": "2023-07-31T17:27:43.296+08:00",
                "updated_at": "2023-08-26T23:11:17.159+08:00",
                "archived": false,
                "is_kia": false,
                "ssh_url_to_repo": "ssh://git@test.gitcode.net:2222/t-qinmh/t-code-1.git",
                "http_url_to_repo": "https://test.gitcode.net/t-qinmh/t-code-1.git",
                "web_url": null,
                "readme_url": null,
                "product_id": null,
                "product_name": null
            },
            "web_url": "https://test.gitcode.net/t-qinmh/t-code-1/merge_requests/30",
            "added_lines": 37,
            "removed_lines": 1,
            "merge_request_type": "MergeRequest",
            "source_git_url": "ssh://git@test.gitcode.net:2222/t-qinmh/t-code-1.git",
            "labels": [],
            "score": 0,
            "min_merged_score": 0,
            "source_product_id": null,
            "target_product_id": null,
            "product_name": null,
            "notes_count": {
                "notes_count": 0,
                "unresolved_notes_count": 0,
                "already_resolved_count": 0,
                "need_resolved_count": 0
            },
            "milestone": null,
            "e2e_issues": []
        }
    ]
}
```

## 4、获取PR详情

### 请求

`GET /api/v4/projects/{project_id}/gitcode/merge_requests/{merge_request_iid}`



### 参数:

|      参数名       | 传参方式 | 类型    | 描述                          |
| :---------------: | :------: | ------- | ----------------------------- |
|    project_id     |   Path   | String  | 项目ID或项目路径              |
| merge_request_iid |   Path   | Integer | PR的IID                       |
|       view        |  Query   | String  | 返回结果视图（simple、basic） |
|    only_count     |  Query   | Boolean | 是否只返回数量                |

### 请求示例：

> /api/v4/projects/{project_id}/gitcode/merge_requests/{merge_request_iid}?view=basic
>
> /api/v4/projects/{project_id}/gitcode/merge_requests/{merge_request_iid}?only_count=true



### 响应:

**only_count=true**

```json
{
    "pipeline_count": 0,
    "binary_count": 0,
    "notes_count": 0,
    "commits_count": 1,
    "diffs_count": "1"
}
```

**view=simple**

```json
{
    "id": 1,
    "iid": 1,
    "title": "test123",
    "description": "新建文件",
    "source_branch": "dev",
    "target_branch": "master",
    "state": "opened",
    "created_at": "2023-07-28T11:46:47.088+08:00",
    "review_mode": "approval",
    "author": {
        "id": 4,
        "name": "yanlp1",
        "username": "7a7657aa0aa04f438981632f623c5e52",
        "iam_id": null,
        "nick_name": "",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": null,
        "web_url": "https://test.gitcode.net/7a7657aa0aa04f438981632f623c5e52",
        "tenant_name": null
    },
    "merge_request_type": "MergeRequest"
}
```

**view=basic**

```json
{
    "id": 61358,
    "iid": 6,
    "state": "opened",
    "title": "测试吴创建的pr",
    "milestone": null,
    "description": "测试五",
    "created_at": "2024-03-24T14:03:40.806+08:00",
    "updated_at": "2024-03-26T19:14:26.105+08:00",
    "target_branch": "main",
    "source_branch": "test_wu",
    "source_project_id": 243377,
    "target_project_id": 243377,
    "squash": false,
    "review_mode": "approval",
    "merge_when_pipeline_succeeds": false,
    "squash_commit_message": null,
    "merge_error": null,
    "json_merge_error": null,
    "sha": "0c02dd57f8945791460a141f155dd2f4bd5dea86",
    "project_id": 243377,
    "head_pipeline_id": null,
    "author": {
        "id": 233,
        "name": "wunian2011",
        "username": "wunian2011",
        "iam_id": "86f5e631aca14debbf750ab93b06395e",
        "nick_name": "测试吴",
        "state": "active",
        "avatar_url": null,
        "avatar_path": null,
        "email": "",
        "name_cn": "wunian2011",
        "web_url": "https://test.gitcode.net/wunian2011",
        "tenant_name": null,
        "is_member": null
    },
    "merged_by": null,
    "merged_at": null,
    "closed_by": null,
    "closed_at": null,
    "diff_refs": {
        "base_sha": "0c02dd57f8945791460a141f155dd2f4bd5dea86",
        "head_sha": "0c02dd57f8945791460a141f155dd2f4bd5dea86",
        "start_sha": "0c02dd57f8945791460a141f155dd2f4bd5dea86"
    },
    "merge_request_type": "MergeRequest",
    "force_remove_source_branch": false,
    "should_remove_source_branch": false,
    "merge_request_assignee_list": null,
    "merge_request_reviewer_list": null,
    "source_project": {
        "id": 243377,
        "description": "csdntest13的第一个项目(公开)",
        "name": "One",
        "name_with_namespace": "One / One",
        "path": "One",
        "path_with_namespace": "One/One",
        "develop_mode": null,
        "created_at": "2024-03-19T16:24:01.197+08:00",
        "updated_at": "2024-03-19T16:42:34.834+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "ssh://git@test.gitcode.net:2222/One/One.git",
        "http_url_to_repo": "https://test.gitcode.net/One/One.git",
        "web_url": null,
        "readme_url": null,
        "product_id": null,
        "product_name": null
    },
    "target_project": {
        "id": 243377,
        "description": "csdntest13的第一个项目(公开)",
        "name": "One",
        "name_with_namespace": "One / One",
        "path": "One",
        "path_with_namespace": "One/One",
        "develop_mode": null,
        "created_at": "2024-03-19T16:24:01.197+08:00",
        "updated_at": "2024-03-19T16:42:34.834+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "ssh://git@test.gitcode.net:2222/One/One.git",
        "http_url_to_repo": "https://test.gitcode.net/One/One.git",
        "web_url": null,
        "readme_url": null,
        "product_id": null,
        "product_name": null
    },
    "omega_mode": false,
    "is_source_branch_exist": true,
    "approval_merge_request_reviewers": null,
    "approval_merge_request_approvers": null,
    "required_reviewers": [],
    "e2e_issues": null,
    "labels": [],
    "auto_merge": null,
    "added_lines": null,
    "removed_lines": null,
    "diverged_commits_count": null,
    "force_rebuild_in_progress": false,
    "rebase_in_progress": null,
    "close_issue_when_merge": false
}
```

## 5、获取PR关联的工作项

### 请求

`GET /api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/e2e_issues`



### 参数:

|      参数名       | 传参方式 | 类型    | 描述             |
| :---------------: | :------: | ------- | ---------------- |
|    project_id     |  Query   | String  | 项目ID或项目路径 |
| merge_request_iid |  Query   | Integer | PR的IID          |

### 请求示例：

> /api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/e2e_issues



### 响应:

```json
{
    "e2e_issues": [
        {
            "id": 345313,
            "issue_type": 7,
            "linked_issue_type": null,
            "issue_num": "issue1",
            "commit_id": null,
            "merge_request_id": 496401,
            "check_fail_reason": "",
            "check_result": true,
            "issue_link": "https://test.gitcode.net/testproject/issues/1",
            "created_at": "2023-08-03T16:15:22.547+08:00",
            "mks_id": null,
            "pbi_id": null,
            "pbi_name": null,
            "source": null,
            "issue_project_id": 1111886,
            "title": "111test",
            "issue_project": null,
            "auto_c_when_mr_merged": false,
            "check_fail_reason_code": null,
            "check_fail_solution": ""
        }
    ]
}
```

## 6、获取PullRequest设置

### 请求

`get /api/v4/projects/{project_id}/merge_requests/setting`



### 参数:

|   参数名   | 传参方式 | 类型   | 描述   |
| :--------: | :------: | ------ | ------ |
| project_id |   Path   | String | 项目id |

### 请求示例：

> /api/v4/projects/test1/merge_requests/setting



### 响应:

```json
{
    "merge_request_setting": {
        "id": 7,
        "project_id": 9,
        "created_at": "2023-07-27T19:47:47.137+08:00",
        "updated_at": "2023-09-19T14:41:47.276+08:00",
        "approval_required_reviewers_enable": false,
        "approval_required_reviewers": 0,
        "disable_merge_by_self": false,
        "can_force_merge": false,
        "add_notes_after_merged": false,
        "mark_auto_merged_mr_as_closed": false,
        "can_reopen": true,
        "delete_source_branch_when_merged": false,
        "disable_squash_merge": false,
        "auto_squash_merge": false,
        "squash_merge_with_no_merge_commit": false,
        "merged_commit_author": "created_by"
    },
    "only_allow_merge_if_all_discussions_are_resolved": false,
    "only_allow_merge_if_pipeline_succeeds": false,
    "merge_method": "merge"
}
```



## 7、更改PullRequest设置

### 请求

`put /api/v4/projects/{project_id}/merge_requests/setting`



### 参数:

|                      参数名                      | 传参方式 | 类型    | 描述                                     |
| :----------------------------------------------: | :------: | ------- | ---------------------------------------- |
|                    project_id                    |   Path   | String  | 项目id                                   |
|        approval_required_reviewers_enable        |   Body   | boolean | 最小评审人数(勾选框)                     |
|           approval_required_reviewers            |   Body   | Integer | 最小评审人数(选择的数字：1~5)            |
| only_allow_merge_if_all_discussions_are_resolved |   Body   | boolean | 评审问题全部解决才能合入                 |
|      only_allow_merge_if_pipeline_succeeds       |   Body   | boolean | 流水线运行通过                           |
|              disable_merge_by_self               |   Body   | boolean | 禁止合入自己创建的PR               |
|                 can_force_merge                  |   Body   | boolean | 允许管理员强制合入                       |
|              add_notes_after_merged              |   Body   | boolean | 允许PR合并后继续做代码检视和评论   |
|          mark_auto_merged_mr_as_closed           |   Body   | boolean | 是否将自动合并的PR状态标记为关闭状态     |
|                    can_reopen                    |   Body   | boolean | 不能重新打开一个已经关闭的PR       |
|         delete_source_branch_when_merged         |   Body   | boolean | PR合入后，默认删除原分支           |
|               disable_squash_merge               |   Body   | boolean | 禁止Squash合并（合入PR时禁止Squash合并） |
|                auto_squash_merge                 |   Body   | boolean | 新建PR，默认开启Squash合并         |
|                   merge_method                   |   Body   | String  | 合并模式(三选一)                         |
|        squash_merge_with_no_merge_commit         |   Body   | boolean | Squash合并不产生Merge节点                |
|               merged_commit_author               |   Body   | String  | 使用PR(合入/创建)者生成Merge Commit      |

### 请求示例：

> /api/v4/projects/test1/merge_requests/setting

```json
{
  "approval_required_reviewers_enable": false,
  "approval_required_reviewers": 0,
  "only_allow_merge_if_all_discussions_are_resolved": false,
  "only_allow_merge_if_pipeline_succeeds": false,
  "disable_merge_by_self": false,
  "can_force_merge": false,
  "add_notes_after_merged": false,
  "mark_auto_merged_mr_as_closed": false,
  "can_reopen": false,
  "delete_source_branch_when_merged": false,
  "disable_squash_merge": false,
  "auto_squash_merge": false,
  "merge_method": "merge",
  "squash_merge_with_no_merge_commit": false,
  "merged_commit_author": "merged_by"
}
```



### 响应:

```json
{
    "merge_request_setting": {
        "id": 7,
        "project_id": 9,
        "created_at": "2023-07-27T19:47:47.137+08:00",
        "updated_at": "2023-09-20T20:12:59.620+08:00",
        "approval_required_reviewers_enable": false,
        "approval_required_reviewers": 0,
        "disable_merge_by_self": false,
        "can_force_merge": false,
        "add_notes_after_merged": false,
        "mark_auto_merged_mr_as_closed": false,
        "can_reopen": false,
        "delete_source_branch_when_merged": false,
        "disable_squash_merge": false,
        "auto_squash_merge": false,
        "squash_merge_with_no_merge_commit": false,
        "merged_commit_author": "merged_by"
    },
    "only_allow_merge_if_all_discussions_are_resolved": false,
    "only_allow_merge_if_pipeline_succeeds": false,
    "merge_method": "merge"
}
```



## 8、获取PR文件变更

### 请求

`GET /api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/changes`



### 参数:

|          参数名          | 传参方式 | 类型    | 必填 | 描述                                                         |
| :----------------------: | :------: | ------- | ---- | ------------------------------------------------------------ |
|        project_id        |   Path   | String  | 是   | 项目id或项目路径                                             |
|    merge_request_iid     |   Path   | Integer | 是   | PR iid                                                       |
|           view           |  Query   | String  | 是   | 可选项: simple                                               |
|        file_path         |  Query   | String  | 否   | 单独获取某个文件的文件变更;如果是重命名文件,使用逗号拼接,如:file_path=old_path,new_path |
| ignore_whitespace_change |  Query   | boolean | 否   | 是否忽略空格变化                                             |
|       from_diff_id       |  Query   | Integer | 否   | 版本比较的中的起点commit id                                  |
|        to_diff_id        |  Query   | Integer | 否   | 版本比较的中的终点commit id                                  |
|        commit_id         |  Query   | String  | 否   | 单独获取某个commit的文件变更                                 |

### 请求示例：

> /api/v4/projects/m0_62019520%2Ftest0828/merge_requests/4/changes?view=simple
>
> /api/v4/projects/m0_62019520%2Ftest0828/merge_requests/2/changes?view=simple&ignore_whitespace_change=false&from_diff_id=2533&to_diff_id=2532



> 注：接口返回文件变更数量最大为**1000**


### 响应:

```json
{
  "added_lines": 2,
  "removed_lines": 0,
  "diff_refs": {
    "base_sha": "8f87ba0d89f7e0107bf0d4249b4c6ab688da8bf7",
    "head_sha": "55953e9546f8d2db4131a3616740cd4b3e88678b",
    "start_sha": "8f87ba0d89f7e0107bf0d4249b4c6ab688da8bf7"
  },
  "changes_count": 2,
  "sha": "55953e9546f8d2db4131a3616740cd4b3e88678b",
  "changes": [
    {
      "old_path": "f3",
      "new_path": "f3",
      "a_mode": "0",
      "b_mode": "100644",
      "file_path": "f3",
      "new_file": true,
      "renamed_file": false,
      "deleted_file": false,
      "diff": "@@ -0,0 +1 @@\n+12312312\n\\ No newline at end of file\n",
      "binary": false,
      "too_large": false,
      "collapsed": false,
      "line_count": null,
      "added_lines": 1,
      "removed_lines": 0,
      "blob_id": "57fb78828dd16e164d3d76abac43c525f2d1ff2e"
    },
    {
      "old_path": "f4",
      "new_path": "f4",
      "a_mode": "0",
      "b_mode": "100644",
      "file_path": "f4",
      "new_file": true,
      "renamed_file": false,
      "deleted_file": false,
      "diff": "@@ -0,0 +1 @@\n+1231231\n\\ No newline at end of file\n",
      "binary": false,
      "too_large": false,
      "collapsed": false,
      "line_count": null,
      "added_lines": 1,
      "removed_lines": 0,
      "blob_id": "31ed7505e87e3e8605841f58c69fdc097c923d91"
    }
  ]
}
```

## 9、创建PR

### URI

`POST /api/v4/projects/{project_id}/merge_requests`

### 传入参数

| 参数名                | 传参方式 | 必填 | 类型    | 描述           |
| --------------------- | -------- | ---- | ------- | -------------- |
| project_id            | path     | Yes  | string  | 项目ID         |
| squash_commit_message | body     | NO   | string  | squash提交信息 |
| inner_issue_nums      | body     | NO   | string  | issue          |
| milestone_id          | body     | NO   | integer | 里程碑ID       |
| target_branch         | body     | Yes  | string  | 目标分支       |
| description           | body     | NO   | string  | 描述           |
| title                 | body     | Yes  | string  | 标题           |
| source_branch         | body     | Yes  | string  | 原分支         |
| labels                | body     | NO   | object  | 标签           |
| squash                | body     | NO   | boolean | 是否squash合入 |
| remove_source_branch  | body     | NO   | boolean | 是否删除原分支 |
| approval_reviewer_ids | body     | NO   | string  | 评审人         |
| target_project_id     | body     | NO   | integer | 目标项目       |
| assignee_ids          | body     | NO   | string  | 合并人         |


### 响应

```xml
{
    "id": 572877,
    "iid": 5,
    "project_id": 1227960,
    "title": "test",
    "description": "新建文件 test1116,\n新建文件 test1118",
    "state": "opened",
    "created_at": "2023-09-21T10:06:56.400+08:00",
    "updated_at": "2023-09-21T10:06:57.425+08:00",
    "merged_by": null,
    "merged_at": null,
    "closed_by": null,
    "closed_at": null,
    "title_html": null,
    "description_html": null,
    "target_branch": "master",
    "source_branch": "dev",
    "squash_commit_message": null,
    "user_notes_count": 0,
    "upvotes": 0,
    "downvotes": 0,
    "author": {
        "id": 88839,
        "name": "gaoyu 30043904",
        "username": "g30043904",
        "state": "active",
        "avatar_url": "https://w3.huawei.com/w3lab/rest/yellowpage/face/30043904/120",
        "avatar_path": null,
        "email": "gaoyu104@huawei.com",
        "name_cn": "高宇 30043904",
        "web_url": "https://gamma.codehub.huawei.com/g30043904",
        "nick_name": null,
        "tenant_name": null
    },
    "assignee": null,
    "source_project_id": 1227960,
    "target_project_id": 1227960,
    "labels": [],
    "work_in_progress": false,
    "milestone": null,
    "merge_when_pipeline_succeeds": false,
    "merge_status": "unchecked",
    "sha": "9d58e3fa00c9407efe47f0a79e351b1ea459e31f",
    "merge_commit_sha": null,
    "discussion_locked": null,
    "should_remove_source_branch": false,
    "force_remove_source_branch": false,
    "allow_collaboration": null,
    "allow_maintainer_to_push": null,
    "web_url": "https://gamma.codehub.huawei.com/g30043904/test35/merge_requests/5",
    "time_stats": {
        "time_estimate": null,
        "total_time_spent": 0,
        "human_time_estimate": null,
        "human_total_time_spent": null
    },
    "squash": true,
    "merge_request_type": "MergeRequest",
    "has_pre_merge_ref": false,
    "review_mode": "approval",
    "is_source_branch_exist": true,
    "approval_merge_request_reviewers": [],
    "approval_merge_request_approvers": [],
    "source_project": {
        "id": 1227960,
        "description": "",
        "name": "test35",
        "name_with_namespace": "g30043904 / test35",
        "path": "test35",
        "path_with_namespace": "g30043904/test35",
        "develop_mode": "normal",
        "created_at": "2023-09-19T22:47:14.938+08:00",
        "updated_at": "2023-09-19T22:47:14.938+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "ssh://git@dgg-g.gamma.codehub.huawei.com:2222/g30043904/test35.git",
        "http_url_to_repo": "https://dgg-g.gamma.codehub.huawei.com/g30043904/test35.git",
        "web_url": "https://gamma.codehub.huawei.com/g30043904/test35",
        "readme_url": "https://gamma.codehub.huawei.com/g30043904/test35/blob/master/README.md",
        "product_id": null,
        "product_name": null,
        "default_branch": "master",
        "tag_list": [],
        "license_url": null,
        "license": null,
        "avatar_url": null,
        "star_count": 0,
        "forks_count": 0,
        "open_issues_count": 0,
        "open_merge_requests_count": 1,
        "open_change_requests_count": null,
        "last_activity_at": "2023-09-21T10:06:40.907+08:00",
        "namespace": {
            "id": 280009,
            "name": "g30043904",
            "path": "g30043904",
            "develop_mode": "normal",
            "region": null,
            "cell": null,
            "kind": "user",
            "full_path": "g30043904",
            "full_name": "g30043904 ",
            "parent_id": null,
            "visibility_level": 20,
            "enable_file_control": null,
            "owner_id": 88839
        },
        "empty_repo": false,
        "starred": false,
        "visibility": "internal",
        "security": "open_source",
        "has_updated_kia": false,
        "network_type": "green",
        "owner": {
            "id": 88839,
            "name": "gaoyu 30043904",
            "username": "g30043904",
            "state": "active",
            "avatar_url": "https://w3.huawei.com/w3lab/rest/yellowpage/face/30043904/120",
            "avatar_path": null,
            "email": "gaoyu104@huawei.com",
            "name_cn": "高宇 30043904",
            "web_url": "https://gamma.codehub.huawei.com/g30043904",
            "nick_name": null,
            "tenant_name": null
        },
        "creator": {
            "id": 88839,
            "name": "gaoyu 30043904",
            "username": "g30043904",
            "state": "active",
            "avatar_url": "https://w3.huawei.com/w3lab/rest/yellowpage/face/30043904/120",
            "avatar_path": null,
            "email": "gaoyu104@huawei.com",
            "name_cn": "高宇 30043904",
            "web_url": "https://gamma.codehub.huawei.com/g30043904",
            "nick_name": null,
            "tenant_name": null
        },
        "creator_id": 88839,
        "forked_from_project": null,
        "item_type": "Project",
        "main_repository_language": [
            null,
            null
        ],
        "statistics": null,
        "branch_count": null,
        "tag_count": null,
        "member_count": null,
        "repo_replica_urls": null,
        "open_external_wiki": true
    },
    "target_project": {
        "id": 1227960,
        "description": "",
        "name": "test35",
        "name_with_namespace": "g30043904 / test35",
        "path": "test35",
        "path_with_namespace": "g30043904/test35",
        "develop_mode": "normal",
        "created_at": "2023-09-19T22:47:14.938+08:00",
        "updated_at": "2023-09-19T22:47:14.938+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "ssh://git@dgg-g.gamma.codehub.huawei.com:2222/g30043904/test35.git",
        "http_url_to_repo": "https://dgg-g.gamma.codehub.huawei.com/g30043904/test35.git",
        "web_url": "https://gamma.codehub.huawei.com/g30043904/test35",
        "readme_url": "https://gamma.codehub.huawei.com/g30043904/test35/blob/master/README.md",
        "product_id": null,
        "product_name": null,
        "default_branch": "master",
        "tag_list": [],
        "license_url": null,
        "license": null,
        "avatar_url": null,
        "star_count": 0,
        "forks_count": 0,
        "open_issues_count": 0,
        "open_merge_requests_count": 1,
        "open_change_requests_count": null,
        "last_activity_at": "2023-09-21T10:06:40.907+08:00",
        "namespace": {
            "id": 280009,
            "name": "g30043904",
            "path": "g30043904",
            "develop_mode": "normal",
            "region": null,
            "cell": null,
            "kind": "user",
            "full_path": "g30043904",
            "full_name": "g30043904 ",
            "parent_id": null,
            "visibility_level": 20,
            "enable_file_control": null,
            "owner_id": 88839
        },
        "empty_repo": false,
        "starred": false,
        "visibility": "internal",
        "security": "open_source",
        "has_updated_kia": false,
        "network_type": "green",
        "owner": {
            "id": 88839,
            "name": "gaoyu 30043904",
            "username": "g30043904",
            "state": "active",
            "avatar_url": "https://w3.huawei.com/w3lab/rest/yellowpage/face/30043904/120",
            "avatar_path": null,
            "email": "gaoyu104@huawei.com",
            "name_cn": "高宇 30043904",
            "web_url": "https://gamma.codehub.huawei.com/g30043904",
            "nick_name": null,
            "tenant_name": null
        },
        "creator": {
            "id": 88839,
            "name": "gaoyu 30043904",
            "username": "g30043904",
            "state": "active",
            "avatar_url": "https://w3.huawei.com/w3lab/rest/yellowpage/face/30043904/120",
            "avatar_path": null,
            "email": "gaoyu104@huawei.com",
            "name_cn": "高宇 30043904",
            "web_url": "https://gamma.codehub.huawei.com/g30043904",
            "nick_name": null,
            "tenant_name": null
        },
        "creator_id": 88839,
        "forked_from_project": null,
        "item_type": "Project",
        "main_repository_language": [
            null,
            null
        ],
        "statistics": null,
        "branch_count": null,
        "tag_count": null,
        "member_count": null,
        "repo_replica_urls": null,
        "open_external_wiki": true
    },
    "added_lines": 1,
    "removed_lines": 0,
    "subscribed": true,
    "changes_count": "1",
    "latest_build_started_at": null,
    "latest_build_finished_at": null,
    "first_deployed_to_production_at": null,
    "pipeline": null,
    "diff_refs": {
        "base_sha": "03a8b68aed65fe6966510a30c4836a28ecdcf715",
        "head_sha": "9d58e3fa00c9407efe47f0a79e351b1ea459e31f",
        "start_sha": "9e648596ac7c150e4674dc885b3d1bb31fa32ba5"
    },
    "merge_error": null,
    "json_merge_error": null,
    "rebase_in_progress": null,
    "diverged_commits_count": null,
    "merge_request_assignee_list": [],
    "merge_request_reviewer_list": [],
    "user": {
        "can_merge": true
    },
    "merge_request_review_count": 0,
    "merge_request_reviewers_count": 0,
    "notes": 0,
    "unresolved_discussions_count": 0,
    "e2e_issues": [],
    "gate_check": true,
    "head_pipeline_id": null,
    "pipeline_status": "",
    "codequality_status": "success",
    "pipeline_status_with_code_quality": "",
    "from_forked_project": false,
    "forked_project_name": null,
    "can_delete_source_branch": true,
    "required_reviewers": [],
    "omega_mode": false,
    "root_mr_locked_detail": null,
    "source_git_url": "ssh://git@dgg-g.gamma.codehub.huawei.com:2222/g30043904/test35.git",
    "auto_merge": null
}
```

## 10、更新PR

### URI

`PUT /api/v4/projects/{project_id}/merge_requests/{merge_request_iid}`

### 传入参数

| 参数名                | 传参方式 | 必填 | 类型    | 描述           |
| --------------------- | -------- | ---- | ------- | -------------- |
| project_id            | path     | Yes  | string  | 项目ID         |
| merge_request_iid     | path     | Yes  | integer | iid            |
| squash_commit_message | body     | NO   | string  | squash提交信息 |
| inner_issue_nums      | body     | NO   | string  | issue          |
| milestone_id          | body     | NO   | integer | 里程碑ID       |
| target_branch         | body     | NO   | string  | 目标分支       |
| description           | body     | NO   | string  | 描述           |
| title                 | body     | NO   | string  | 标题           |
| labels                | body     | NO   | object  | 标签           |
| squash                | body     | NO   | boolean | squash合入     |
| remove_source_branch  | body     | NO   | boolean | 是否删除原分支 |
| state_event           | body     | NO   | string  | 状态           |
| assignee_ids          | body     | NO   | string  | 合并人         |

### 响应

```xml
{
    "id": 22580,
    "iid": 64,
    "project_id": 9,
    "title": "test12345",
    "description": "新建文件 test123234",
    "state": "opened",
    "created_at": "2023-09-08T22:17:16.427+08:00",
    "updated_at": "2023-09-08T22:22:57.020+08:00",
    "merged_by": null,
    "merged_at": null,
    "closed_by": null,
    "closed_at": null,
    "title_html": null,
    "description_html": null,
    "target_branch": "master",
    "source_branch": "test1",
    "squash_commit_message": null,
    "user_notes_count": 0,
    "upvotes": 0,
    "downvotes": 0,
    "author": {
        "id": 12,
        "name": "f4ManagerTest",
        "username": "f4ManagerTest-update",
        "iam_id": "a618e34bd5704be3ae3395dfede06041_k",
        "nick_name": "测试我有没有一头小毛驴",
        "state": "active",
        "avatar_url": "https://www.google.com.hk/url?sa=i&url=http%3A%2F%2F699pic.com%2F&psig=AOvVaw1eJLkLXz8b5cXfICVndJ-9&ust=1693911080514000&source=images&cd=vfe&opi=89978449&ved=0CA4QjRxqFwoTCLCO49PkkIEDFQAAAAAdAAAAABAE",
        "avatar_path": null,
        "email": "ManagerTest4@example.com",
        "name_cn": null,
        "web_url": "https://gamma.codehub.huawei.com/f4ManagerTest-update",
        "tenant_name": null
    },
    "assignee": null,
    "source_project_id": 9,
    "target_project_id": 9,
    "labels": [],
    "work_in_progress": false,
    "milestone": null,
    "merge_when_pipeline_succeeds": false,
    "merge_status": "unchecked",
    "sha": "e3b9d5590984eac6b88bfa75ef5d9d95df6f53a1",
    "merge_commit_sha": null,
    "discussion_locked": null,
    "should_remove_source_branch": false,
    "force_remove_source_branch": false,
    "allow_collaboration": null,
    "allow_maintainer_to_push": null,
    "web_url": "https://gamma.codehub.huawei.com/f4ManagerTest-update/009/merge_requests/64",
    "time_stats": {
        "time_estimate": null,
        "total_time_spent": 0,
        "human_time_estimate": null,
        "human_total_time_spent": null
    },
    "squash": false,
    "merge_request_type": "MergeRequest",
    "has_pre_merge_ref": false,
    "review_mode": "approval",
    "is_source_branch_exist": true,
    "approval_merge_request_reviewers": [],
    "approval_merge_request_approvers": [],
    "source_project": {
        "id": 9,
        "description": null,
        "name": "009",
        "name_with_namespace": "f4ManagerTest-update / 009",
        "path": "009",
        "path_with_namespace": "f4ManagerTest-update/009",
        "develop_mode": "normal",
        "created_at": "2023-07-24T10:06:38.413+08:00",
        "updated_at": "2023-07-24T10:06:38.413+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "ssh://git@dgg-g.dev.huawei.com:1111/f4ManagerTest-update/009.git",
        "http_url_to_repo": "https://dgg-g.dev.huawei.com/f4ManagerTest-update/009.git",
        "web_url": "https://gamma.codehub.huawei.com/f4ManagerTest-update/009",
        "readme_url": "https://gamma.codehub.huawei.com/f4ManagerTest-update/009/blob/main/README.md",
        "product_id": null,
        "product_name": null,
        "default_branch": "master",
        "tag_list": [],
        "license_url": null,
        "license": {
            "key": "Creative Commons Attribution Share Alike 4.0",
            "name": null,
            "nickname": null,
            "html_url": null,
            "source_url": null
        },
        "avatar_url": null,
        "star_count": 0,
        "forks_count": 1,
        "open_issues_count": 13,
        "open_merge_requests_count": 9,
        "open_change_requests_count": null,
        "watch_count": 0,
        "last_activity_at": "2023-09-08T22:16:50.655+08:00",
        "namespace": {
            "id": 2,
            "name": "f4ManagerTest-update",
            "path": "f4ManagerTest-update",
            "develop_mode": "normal",
            "region": null,
            "cell": null,
            "kind": "user",
            "full_path": "f4ManagerTest-update",
            "full_name": "f4ManagerTest-update ",
            "parent_id": null,
            "visibility_level": 20,
            "enable_file_control": null,
            "owner_id": 12
        },
        "empty_repo": false,
        "starred": false,
        "visibility": "public",
        "security": "internal",
        "has_updated_kia": false,
        "network_type": "green",
        "owner": {
            "id": 12,
            "name": "f4ManagerTest",
            "username": "f4ManagerTest-update",
            "iam_id": "a618e34bd5704be3ae3395dfede06041_k",
            "nick_name": "测试我有没有一头小毛驴",
            "state": "active",
            "avatar_url": "https://www.google.com.hk/url?sa=i&url=http%3A%2F%2F699pic.com%2F&psig=AOvVaw1eJLkLXz8b5cXfICVndJ-9&ust=1693911080514000&source=images&cd=vfe&opi=89978449&ved=0CA4QjRxqFwoTCLCO49PkkIEDFQAAAAAdAAAAABAE",
            "avatar_path": null,
            "email": "ManagerTest4@example.com",
            "name_cn": null,
            "web_url": "https://gamma.codehub.huawei.com/f4ManagerTest-update",
            "tenant_name": null
        },
        "creator": {
            "id": 12,
            "name": "f4ManagerTest",
            "username": "f4ManagerTest-update",
            "iam_id": "a618e34bd5704be3ae3395dfede06041_k",
            "nick_name": "测试我有没有一头小毛驴",
            "state": "active",
            "avatar_url": "https://www.google.com.hk/url?sa=i&url=http%3A%2F%2F699pic.com%2F&psig=AOvVaw1eJLkLXz8b5cXfICVndJ-9&ust=1693911080514000&source=images&cd=vfe&opi=89978449&ved=0CA4QjRxqFwoTCLCO49PkkIEDFQAAAAAdAAAAABAE",
            "avatar_path": null,
            "email": "ManagerTest4@example.com",
            "name_cn": null,
            "web_url": "https://gamma.codehub.huawei.com/f4ManagerTest-update",
            "tenant_name": null
        },
        "creator_id": 12,
        "forked_from_project": null,
        "item_type": "Project",
        "main_repository_language": [
            null,
            null
        ],
        "statistics": null,
        "branch_count": null,
        "tag_count": null,
        "member_count": null,
        "repo_replica_urls": null,
        "open_external_wiki": true,
        "release_count": null
    },
    "target_project": {
        "id": 9,
        "description": null,
        "name": "009",
        "name_with_namespace": "f4ManagerTest-update / 009",
        "path": "009",
        "path_with_namespace": "f4ManagerTest-update/009",
        "develop_mode": "normal",
        "created_at": "2023-07-24T10:06:38.413+08:00",
        "updated_at": "2023-07-24T10:06:38.413+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "ssh://git@dgg-g.dev.huawei.com:1111/f4ManagerTest-update/009.git",
        "http_url_to_repo": "https://dgg-g.dev.huawei.com/f4ManagerTest-update/009.git",
        "web_url": "https://gamma.codehub.huawei.com/f4ManagerTest-update/009",
        "readme_url": "https://gamma.codehub.huawei.com/f4ManagerTest-update/009/blob/main/README.md",
        "product_id": null,
        "product_name": null,
        "default_branch": "master",
        "tag_list": [],
        "license_url": null,
        "license": {
            "key": "Creative Commons Attribution Share Alike 4.0",
            "name": null,
            "nickname": null,
            "html_url": null,
            "source_url": null
        },
        "avatar_url": null,
        "star_count": 0,
        "forks_count": 1,
        "open_issues_count": 13,
        "open_merge_requests_count": 9,
        "open_change_requests_count": null,
        "watch_count": 0,
        "last_activity_at": "2023-09-08T22:16:50.655+08:00",
        "namespace": {
            "id": 2,
            "name": "f4ManagerTest-update",
            "path": "f4ManagerTest-update",
            "develop_mode": "normal",
            "region": null,
            "cell": null,
            "kind": "user",
            "full_path": "f4ManagerTest-update",
            "full_name": "f4ManagerTest-update ",
            "parent_id": null,
            "visibility_level": 20,
            "enable_file_control": null,
            "owner_id": 12
        },
        "empty_repo": false,
        "starred": false,
        "visibility": "public",
        "security": "internal",
        "has_updated_kia": false,
        "network_type": "green",
        "owner": {
            "id": 12,
            "name": "f4ManagerTest",
            "username": "f4ManagerTest-update",
            "iam_id": "a618e34bd5704be3ae3395dfede06041_k",
            "nick_name": "测试我有没有一头小毛驴",
            "state": "active",
            "avatar_url": "https://www.google.com.hk/url?sa=i&url=http%3A%2F%2F699pic.com%2F&psig=AOvVaw1eJLkLXz8b5cXfICVndJ-9&ust=1693911080514000&source=images&cd=vfe&opi=89978449&ved=0CA4QjRxqFwoTCLCO49PkkIEDFQAAAAAdAAAAABAE",
            "avatar_path": null,
            "email": "ManagerTest4@example.com",
            "name_cn": null,
            "web_url": "https://gamma.codehub.huawei.com/f4ManagerTest-update",
            "tenant_name": null
        },
        "creator": {
            "id": 12,
            "name": "f4ManagerTest",
            "username": "f4ManagerTest-update",
            "iam_id": "a618e34bd5704be3ae3395dfede06041_k",
            "nick_name": "测试我有没有一头小毛驴",
            "state": "active",
            "avatar_url": "https://www.google.com.hk/url?sa=i&url=http%3A%2F%2F699pic.com%2F&psig=AOvVaw1eJLkLXz8b5cXfICVndJ-9&ust=1693911080514000&source=images&cd=vfe&opi=89978449&ved=0CA4QjRxqFwoTCLCO49PkkIEDFQAAAAAdAAAAABAE",
            "avatar_path": null,
            "email": "ManagerTest4@example.com",
            "name_cn": null,
            "web_url": "https://gamma.codehub.huawei.com/f4ManagerTest-update",
            "tenant_name": null
        },
        "creator_id": 12,
        "forked_from_project": null,
        "item_type": "Project",
        "main_repository_language": [
            null,
            null
        ],
        "statistics": null,
        "branch_count": null,
        "tag_count": null,
        "member_count": null,
        "repo_replica_urls": null,
        "open_external_wiki": true,
        "release_count": null
    },
    "added_lines": 4,
    "removed_lines": 0,
    "subscribed": true,
    "changes_count": "5",
    "latest_build_started_at": null,
    "latest_build_finished_at": null,
    "first_deployed_to_production_at": null,
    "pipeline": null,
    "diff_refs": {
        "base_sha": "73f01c735a13bd7f7552116ae1380afa18aa1f4c",
        "head_sha": "e3b9d5590984eac6b88bfa75ef5d9d95df6f53a1",
        "start_sha": "c4ebe0aec07a5ab90755e3b1c84024118dad467c"
    },
    "merge_error": null,
    "json_merge_error": null,
    "rebase_in_progress": null,
    "diverged_commits_count": null,
    "merge_request_assignee_list": [],
    "merge_request_reviewer_list": [],
    "user": {
        "can_merge": true
    },
    "merge_request_review_count": 0,
    "merge_request_reviewers_count": 0,
    "notes": 0,
    "unresolved_discussions_count": 0,
    "e2e_issues": [],
    "gate_check": true,
    "head_pipeline_id": null,
    "pipeline_status": "",
    "codequality_status": "success",
    "pipeline_status_with_code_quality": "",
    "from_forked_project": false,
    "forked_project_name": null,
    "can_delete_source_branch": true,
    "required_reviewers": [],
    "omega_mode": false,
    "root_mr_locked_detail": null,
    "source_git_url": "ssh://git@dgg-g.dev.huawei.com:1111/f4ManagerTest-update/009.git",
    "auto_merge": null
}
```

## 11、合并PR

### URI

`PUT /api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/merge`

### 传入参数

| 参数名                      | 传参方式 | 必填 | 类型    | 描述           |
| --------------------------- | -------- | ---- | ------- | -------------- |
| project_id                  | path     | Yes  | string  | 项目ID         |
| merge_request_iid           | path     | Yes  | integer | iid            |
| merge_request_id            | body     | NO   | integer | PR Id           |
| squash_commit_message       | body     | NO   | string  | squash提交信息 |
| force_merge                 | body     | NO   | boolean | 是否强制合入   |
| squash                      | body     | NO   | boolean | 是否squash合入 |
| merge_commit_message        | body     | NO   | string  | 合并提交信息   |
| should_remove_source_branch | body     | NO   | boolean | 是否删除原分支 |

### 响应

```xml
{
    "id": 46564,
    "iid": 2,
    "project_id": 134181,
    "title": "test234",
    "description": "",
    "state": "merged",
    "created_at": "2023-09-20T16:34:54.132+08:00",
    "updated_at": "2023-09-20T16:35:18.098+08:00",
    "merged_by": {
        "id": 12,
        "name": "f4ManagerTest",
        "username": "f4ManagerTest-update",
        "iam_id": "a618e34bd5704be3ae3395dfede06041_k",
        "nick_name": "测试我有没有一头小毛驴",
        "state": "active",
        "avatar_url": "https://www.google.com.hk/url?sa=i&url=http%3A%2F%2F699pic.com%2F&psig=AOvVaw1eJLkLXz8b5cXfICVndJ-9&ust=1693911080514000&source=images&cd=vfe&opi=89978449&ved=0CA4QjRxqFwoTCLCO49PkkIEDFQAAAAAdAAAAABAE",
        "avatar_path": null,
        "email": "ManagerTest4@example.com",
        "name_cn": null,
        "web_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update",
        "tenant_name": null
    },
    "merged_at": "2023-09-20T16:35:17.858+08:00",
    "closed_by": null,
    "closed_at": null,
    "title_html": null,
    "description_html": null,
    "target_branch": "main",
    "source_branch": "dev",
    "squash_commit_message": null,
    "user_notes_count": 0,
    "upvotes": 0,
    "downvotes": 0,
    "author": {
        "id": 12,
        "name": "f4ManagerTest",
        "username": "f4ManagerTest-update",
        "iam_id": "a618e34bd5704be3ae3395dfede06041_k",
        "nick_name": "测试我有没有一头小毛驴",
        "state": "active",
        "avatar_url": "https://www.google.com.hk/url?sa=i&url=http%3A%2F%2F699pic.com%2F&psig=AOvVaw1eJLkLXz8b5cXfICVndJ-9&ust=1693911080514000&source=images&cd=vfe&opi=89978449&ved=0CA4QjRxqFwoTCLCO49PkkIEDFQAAAAAdAAAAABAE",
        "avatar_path": null,
        "email": "ManagerTest4@example.com",
        "name_cn": null,
        "web_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update",
        "tenant_name": null
    },
    "assignee": null,
    "source_project_id": 134181,
    "target_project_id": 134181,
    "labels": [],
    "work_in_progress": false,
    "milestone": null,
    "merge_when_pipeline_succeeds": false,
    "merge_status": "can_be_merged",
    "sha": "6a41b32ab5ca4aa38f849fd280a75d41b78cf3d7",
    "merge_commit_sha": "cb04c6d83f74bf90745857c61900ce51cbcbc104",
    "discussion_locked": null,
    "should_remove_source_branch": false,
    "force_remove_source_branch": false,
    "allow_collaboration": null,
    "allow_maintainer_to_push": null,
    "web_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update/gyTest/merge_requests/2",
    "time_stats": {
        "time_estimate": null,
        "total_time_spent": 0,
        "human_time_estimate": null,
        "human_total_time_spent": null
    },
    "squash": false,
    "merge_request_type": "MergeRequest",
    "has_pre_merge_ref": false,
    "review_mode": "approval",
    "is_source_branch_exist": true,
    "approval_merge_request_reviewers": [],
    "approval_merge_request_approvers": [],
    "source_project": {
        "id": 134181,
        "description": null,
        "name": "gyTest",
        "name_with_namespace": "f4ManagerTest-update / gyTest",
        "path": "gyTest",
        "path_with_namespace": "f4ManagerTest-update/gyTest",
        "develop_mode": "normal",
        "created_at": "2023-09-20T15:26:15.744+08:00",
        "updated_at": "2023-09-20T15:26:15.744+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "ssh://git@gitcode-backend.cn-north-7.myhuaweicloud.com:2222/f4ManagerTest-update/gyTest.git",
        "http_url_to_repo": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update/gyTest.git",
        "web_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update/gyTest",
        "readme_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update/gyTest/blob/main/README.md",
        "product_id": null,
        "product_name": null,
        "default_branch": "main",
        "tag_list": [],
        "license_url": null,
        "license": {
            "key": null,
            "name": null,
            "nickname": null,
            "html_url": null,
            "source_url": null
        },
        "avatar_url": null,
        "star_count": 0,
        "forks_count": 0,
        "open_issues_count": 0,
        "open_merge_requests_count": 0,
        "open_change_requests_count": null,
        "watch_count": 0,
        "last_activity_at": "2023-09-20T16:35:17.870+08:00",
        "namespace": {
            "id": 2,
            "name": "f4ManagerTest-update",
            "path": "f4ManagerTest-update",
            "develop_mode": "normal",
            "region": null,
            "cell": null,
            "kind": "user",
            "full_path": "f4ManagerTest-update",
            "full_name": "f4ManagerTest-update ",
            "parent_id": null,
            "visibility_level": 20,
            "enable_file_control": null,
            "owner_id": 12
        },
        "empty_repo": false,
        "starred": false,
        "visibility": "public",
        "security": "internal",
        "has_updated_kia": false,
        "network_type": "green",
        "owner": {
            "id": 12,
            "name": "f4ManagerTest",
            "username": "f4ManagerTest-update",
            "iam_id": "a618e34bd5704be3ae3395dfede06041_k",
            "nick_name": "测试我有没有一头小毛驴",
            "state": "active",
            "avatar_url": "https://www.google.com.hk/url?sa=i&url=http%3A%2F%2F699pic.com%2F&psig=AOvVaw1eJLkLXz8b5cXfICVndJ-9&ust=1693911080514000&source=images&cd=vfe&opi=89978449&ved=0CA4QjRxqFwoTCLCO49PkkIEDFQAAAAAdAAAAABAE",
            "avatar_path": null,
            "email": "ManagerTest4@example.com",
            "name_cn": null,
            "web_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update",
            "tenant_name": null
        },
        "creator": {
            "id": 12,
            "name": "f4ManagerTest",
            "username": "f4ManagerTest-update",
            "iam_id": "a618e34bd5704be3ae3395dfede06041_k",
            "nick_name": "测试我有没有一头小毛驴",
            "state": "active",
            "avatar_url": "https://www.google.com.hk/url?sa=i&url=http%3A%2F%2F699pic.com%2F&psig=AOvVaw1eJLkLXz8b5cXfICVndJ-9&ust=1693911080514000&source=images&cd=vfe&opi=89978449&ved=0CA4QjRxqFwoTCLCO49PkkIEDFQAAAAAdAAAAABAE",
            "avatar_path": null,
            "email": "ManagerTest4@example.com",
            "name_cn": null,
            "web_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update",
            "tenant_name": null
        },
        "creator_id": 12,
        "forked_from_project": null,
        "item_type": "Project",
        "main_repository_language": [
            null,
            null
        ],
        "mirror_project_data": null,
        "statistics": null,
        "branch_count": null,
        "tag_count": null,
        "member_count": null,
        "repo_replica_urls": null,
        "open_external_wiki": true,
        "release_count": null
    },
    "target_project": {
        "id": 134181,
        "description": null,
        "name": "gyTest",
        "name_with_namespace": "f4ManagerTest-update / gyTest",
        "path": "gyTest",
        "path_with_namespace": "f4ManagerTest-update/gyTest",
        "develop_mode": "normal",
        "created_at": "2023-09-20T15:26:15.744+08:00",
        "updated_at": "2023-09-20T15:26:15.744+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "ssh://git@gitcode-backend.cn-north-7.myhuaweicloud.com:2222/f4ManagerTest-update/gyTest.git",
        "http_url_to_repo": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update/gyTest.git",
        "web_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update/gyTest",
        "readme_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update/gyTest/blob/main/README.md",
        "product_id": null,
        "product_name": null,
        "default_branch": "main",
        "tag_list": [],
        "license_url": null,
        "license": {
            "key": null,
            "name": null,
            "nickname": null,
            "html_url": null,
            "source_url": null
        },
        "avatar_url": null,
        "star_count": 0,
        "forks_count": 0,
        "open_issues_count": 0,
        "open_merge_requests_count": 0,
        "open_change_requests_count": null,
        "watch_count": 0,
        "last_activity_at": "2023-09-20T16:35:17.870+08:00",
        "namespace": {
            "id": 2,
            "name": "f4ManagerTest-update",
            "path": "f4ManagerTest-update",
            "develop_mode": "normal",
            "region": null,
            "cell": null,
            "kind": "user",
            "full_path": "f4ManagerTest-update",
            "full_name": "f4ManagerTest-update ",
            "parent_id": null,
            "visibility_level": 20,
            "enable_file_control": null,
            "owner_id": 12
        },
        "empty_repo": false,
        "starred": false,
        "visibility": "public",
        "security": "internal",
        "has_updated_kia": false,
        "network_type": "green",
        "owner": {
            "id": 12,
            "name": "f4ManagerTest",
            "username": "f4ManagerTest-update",
            "iam_id": "a618e34bd5704be3ae3395dfede06041_k",
            "nick_name": "测试我有没有一头小毛驴",
            "state": "active",
            "avatar_url": "https://www.google.com.hk/url?sa=i&url=http%3A%2F%2F699pic.com%2F&psig=AOvVaw1eJLkLXz8b5cXfICVndJ-9&ust=1693911080514000&source=images&cd=vfe&opi=89978449&ved=0CA4QjRxqFwoTCLCO49PkkIEDFQAAAAAdAAAAABAE",
            "avatar_path": null,
            "email": "ManagerTest4@example.com",
            "name_cn": null,
            "web_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update",
            "tenant_name": null
        },
        "creator": {
            "id": 12,
            "name": "f4ManagerTest",
            "username": "f4ManagerTest-update",
            "iam_id": "a618e34bd5704be3ae3395dfede06041_k",
            "nick_name": "测试我有没有一头小毛驴",
            "state": "active",
            "avatar_url": "https://www.google.com.hk/url?sa=i&url=http%3A%2F%2F699pic.com%2F&psig=AOvVaw1eJLkLXz8b5cXfICVndJ-9&ust=1693911080514000&source=images&cd=vfe&opi=89978449&ved=0CA4QjRxqFwoTCLCO49PkkIEDFQAAAAAdAAAAABAE",
            "avatar_path": null,
            "email": "ManagerTest4@example.com",
            "name_cn": null,
            "web_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update",
            "tenant_name": null
        },
        "creator_id": 12,
        "forked_from_project": null,
        "item_type": "Project",
        "main_repository_language": [
            null,
            null
        ],
        "mirror_project_data": null,
        "statistics": null,
        "branch_count": null,
        "tag_count": null,
        "member_count": null,
        "repo_replica_urls": null,
        "open_external_wiki": true,
        "release_count": null
    },
    "added_lines": 1,
    "removed_lines": 0,
    "subscribed": true,
    "changes_count": "1",
    "latest_build_started_at": null,
    "latest_build_finished_at": null,
    "first_deployed_to_production_at": null,
    "pipeline": null,
    "diff_refs": {
        "base_sha": "7e96181973808f036f08cf1a83d1ea35c9deea00",
        "head_sha": "6a41b32ab5ca4aa38f849fd280a75d41b78cf3d7",
        "start_sha": "3d8a36b6bd762a606d979885d286358c1be5cceb"
    },
    "merge_error": null,
    "json_merge_error": null,
    "rebase_in_progress": null,
    "diverged_commits_count": null,
    "merge_request_assignee_list": [],
    "merge_request_reviewer_list": [],
    "user": {
        "can_merge": true
    },
    "merge_request_review_count": 0,
    "merge_request_reviewers_count": 0,
    "notes": 0,
    "unresolved_discussions_count": 0,
    "e2e_issues": [],
    "gate_check": true,
    "head_pipeline_id": null,
    "pipeline_status": "",
    "codequality_status": "success",
    "pipeline_status_with_code_quality": "",
    "from_forked_project": false,
    "forked_project_name": null,
    "can_delete_source_branch": true,
    "required_reviewers": [],
    "omega_mode": false,
    "root_mr_locked_detail": null,
    "source_git_url": "ssh://git@gitcode-backend.cn-north-7.myhuaweicloud.com:2222/f4ManagerTest-update/gyTest.git",
    "auto_merge": null
}
```

## 12.更新PR标签接口

```xml
'/api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/labels':
put:
    description: update merge request labels
    produces:
    - application/json
    consumes:
    - application/json
    parameters:
    - in: path
        name: project_id
        description: The ID of a project
        type: string
        required: true
    - in: path
        name: merge_request_iid
        description: The merge request iid
        type: integer
        format: int32
        required: true
    - in: query
        name: labels
        description: Comma-separated list of label names
        type: string
        required: true
    responses:
    '200':
        description: Update merge request labels
        schema:
        type: object
    tags:
    - mergeRequest
    operationId: updateMergeRequestLabels

```
### 传入参数示例
```xml
{
    "labels": "BUG1,BUG2,BUG3"  // String，必填，使用逗号分隔的label titles
}

```

### 响应参数示例
```xml
{
    "id": 87655,
    "iid": 44,
    "project_id": 182759,
    "title": "test234",
    "description": "",
    "state": "opened",
    "created_at": "2024-03-23T14:07:30.619+08:00",
    "updated_at": "2024-03-23T14:10:34.626+08:00",
    "merged_by": null,
    "merged_at": null,
    "closed_by": null,
    "closed_at": null,
    "title_html": null,
    "description_html": null,
    "target_branch": "main",
    "source_branch": "dev",
    "squash_commit_message": null,
    "user_notes_count": 0,
    "upvotes": 0,
    "downvotes": 0,
    "author": {},
    "assignee": null,
    "source_project_id": 182759,
    "target_project_id": 182759,
    "labels": [
        {
            "color": "#428BCA",
            "name": "BUG1",
            "id": 1412560,
            "title": "BUG1",
            "type": null,
            "textColor": "#FFFFFF"
        },
        {
            "color": "#428BCA",
            "name": "BUG2",
            "id": 1412561,
            "title": "BUG2",
            "type": null,
            "textColor": "#FFFFFF"
        },
        {
            "color": "#428BCA",
            "name": "BUG3",
            "id": 1412562,
            "title": "BUG3",
            "type": null,
            "textColor": "#FFFFFF"
        }
    ]
}

```

## 13.删除PR标签接口

```xml
'/api/v4/projects/{project_id}/merge_requests/{merge_request_iid}/labels':
delete:
    description: update merge request labels
    produces:
    - application/json
    consumes:
    - application/json
    parameters:
    - in: path
        name: project_id
        description: The ID of a project
        type: string
        required: true
    - in: path
        name: merge_request_iid
        description: The merge request iid
        type: integer
        format: int32
        required: true
    - in: query
        name: labels
        description: Comma-separated list of label names
        type: string
        required: true
    responses:
    '200':
        description: Delete merge request labels
        schema:
        type: object
    tags:
    - mergeRequest
    operationId: deleteMergeRequestLabels

```
### 传入参数示例
```xml
{
    "labels": "BUG1,BUG2,BUG3"  // String，必填，使用逗号分隔的label titles
}

```

### 响应参数示例
```xml
{
    "id": 87655,
    "iid": 44,
    "project_id": 182759,
    "title": "test234",
    "description": "",
    "state": "opened",
    "created_at": "2024-03-23T14:07:30.619+08:00",
    "updated_at": "2024-03-23T14:10:34.626+08:00",
    "merged_by": null,
    "merged_at": null,
    "closed_by": null,
    "closed_at": null,
    "title_html": null,
    "description_html": null,
    "target_branch": "main",
    "source_branch": "dev",
    "squash_commit_message": null,
    "user_notes_count": 0,
    "upvotes": 0,
    "downvotes": 0,
    "author": {},
    "assignee": null,
    "source_project_id": 182759,
    "target_project_id": 182759,
    "labels": [
        {
            "color": "#428BCA",
            "name": "BUG2",
            "id": 1412561,
            "title": "BUG2",
            "type": null,
            "textColor": "#FFFFFF"
        },
        {
            "color": "#428BCA",
            "name": "BUG3",
            "id": 1412562,
            "title": "BUG3",
            "type": null,
            "textColor": "#FFFFFF"
        }
    ]
}

```