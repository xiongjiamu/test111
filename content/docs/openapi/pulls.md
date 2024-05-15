---
linkTitle: Pull Requests
title: PR接口文档
weight: 6
sidebar:
  open: false
---

## 1. 获取Pull Request列表

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls`

### 参数

| 参数名  | 描述                                          | 类型  | 数据类型  |
| ------ |---------------------------------------------| ------  |------|
|  access_token* | 用户授权码                                       | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path)                      | path | string    |
|  repo*   | 仓库路径(path)                                  | path | string    |
|  state | 可选。Pull Request 状态: all、open、closed、locked、merged | query | string    |
|  base   | 可选。Pull Request 提交目标分支的名称                   | query | string    |
|  since | 可选。起始的更新时间，要求时间格式为 ISO 8601                 | quey | string    |
|  direction  | 可选。升序/降序                                    | query | string    |
|  sort   | 可选。排序字段: created、updated          | query | string    |
|  milestone_number | 可选。里程碑序号(id)                                | quey | integer    |
|  labels | 以逗号分隔的标签名称列表                                | quey | string    |
|  page   | 当前的页码                                       | query | integer    |
|  per_page   | 每页的数量，最大为 100                               | query | integer    |
|  author   | 可选。PR 创建者用户名                                | query | string    |
|  assignee   | 可选。评审者用户名                                   | query | string    |

### 响应

```json
[
    {
        "number": 63,
        "html_url": "https://test.gitcode.net/One/One/merge_requests/63",
        "close_related_issue": null,
        "prune_branch": false,
        "draft": false,
        "url":"https://api.gitcode.net/api/v5/repos/One/One/pulls/63",
        "labels": [
            {
                "id": 381445,
                "color": "#008672",
                "name": "help wanted",
                "title": "help wanted",
                "repository_id": 243377,
                "type": null,
                "text_color": null
            },
            {
                "id": 381446,
                "color": "#CFD240",
                "name": "invalid",
                "title": "invalid",
                "repository_id": 243377,
                "type": null,
                "text_color": null
            },
            {
                "id": 381447,
                "color": "#D876E3",
                "name": "question",
                "title": "question",
                "repository_id": 243377,
                "type": null,
                "text_color": null
            }
        ],
        "user": {
            "id": "65f94ab6f21fa3084fc04823",
            "login": "csdntest13",
            "name": "csdntest13_gitcode",
            "state": "active",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png?time=1711533165876",
            "avatar_path": null,
            "email": "",
            "name_cn": "csdntest13",
            "html_url": "https://test.gitcode.net/csdntest13",
            "tenant_name": null,
            "is_member": null
        },
        "assignees": [
            {
                "id": "64c71c3d64037b4af1c7a93f",
                "login": "green",
                "name": "百里",
                "state": "optional",
                "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/be/fb/7b9e393fbd80ca315dec249f2be6e6a7378f591609b6525798bc6d95abedc992.png?time=1712128581171",
                "avatar_path": null,
                "email": null,
                "name_cn": "green",
                "html_url": "https://test.gitcode.net/green",
                "assignee": true,
                "code_owner": false,
                "accept": false
            }
        ],
        "head": {
            "label": "test_b12",
            "ref": "test_b12",
            "sha": "fb6495834d1bf7a39dfdb44ad25e6f83c7136310",
            "user": {
                "id": "65f94ab6f21fa3084fc04823",
                "login": "csdntest13",
                "name": "csdntest13_gitcode",
                "state": "active",
                "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png?time=1711533165876",
                "avatar_path": null,
                "email": "",
                "name_cn": "csdntest13",
                "html_url": "https://test.gitcode.net/csdntest13",
                "tenant_name": null,
                "is_member": null
            },
            "repo": {
                "id": 243377,
                "full_path": "One/One",
                "human_name": "One / One",
                "name": "One",
                "path": "One",
                "description": "csdntest13的第一个项目(公开)",
                "namespace": {
                    "id": 136909,
                    "name": "One",
                    "path": "One",
                    "develop_mode": "normal",
                    "region": null,
                    "cell": "default",
                    "kind": "group",
                    "full_path": "One",
                    "full_name": "One ",
                    "parent_id": null,
                    "visibility_level": 20,
                    "enable_file_control": null,
                    "owner_id": null
                },
                "owner": {
                    "id": "65f94ab6f21fa3084fc04823",
                    "login": "csdntest13",
                    "name": "csdntest13_gitcode",
                    "state": "active",
                    "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png?time=1711533165876",
                    "avatar_path": null,
                    "email": "",
                    "name_cn": "csdntest13",
                    "html_url": "https://test.gitcode.net/csdntest13",
                    "tenant_name": null,
                    "is_member": null
                },
                "assigner": {
                    "id": "64c71c3d64037b4af1c7a93f",
                    "login": "green",
                    "name": "百里",
                    "state": null,
                    "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/be/fb/7b9e393fbd80ca315dec249f2be6e6a7378f591609b6525798bc6d95abedc992.png?time=1712128581171",
                    "avatar_path": null,
                    "email": null,
                    "name_cn": null,
                    "html_url": "https://test.gitcode.net/green",
                    "tenant_name": null,
                    "is_member": null
                },
                "private": null,
                "public": null,
                "internal": false
            }
        },
        "base": {
            "label": "dev",
            "ref": "dev",
            "sha": "0c02dd57f8945791460a141f155dd2f4bd5dea86",
            "user": {
                "id": "65f94ab6f21fa3084fc04823",
                "login": "csdntest13",
                "name": "csdntest13_gitcode",
                "state": "active",
                "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png?time=1711533165876",
                "avatar_path": null,
                "email": "",
                "name_cn": "csdntest13",
                "html_url": "https://test.gitcode.net/csdntest13",
                "tenant_name": null,
                "is_member": null
            },
            "repo": {
                "id": 243377,
                "full_path": "One/One",
                "human_name": "One / One",
                "name": "One",
                "path": "One",
                "description": "csdntest13的第一个项目(公开)",
                "namespace": {
                    "id": 136909,
                    "name": "One",
                    "path": "One",
                    "develop_mode": "normal",
                    "region": null,
                    "cell": "default",
                    "kind": "group",
                    "full_path": "One",
                    "full_name": "One ",
                    "parent_id": null,
                    "visibility_level": 20,
                    "enable_file_control": null,
                    "owner_id": null
                },
                "owner": {
                    "id": "65f94ab6f21fa3084fc04823",
                    "login": "csdntest13",
                    "name": "csdntest13_gitcode",
                    "state": "active",
                    "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png?time=1711533165876",
                    "avatar_path": null,
                    "email": "",
                    "name_cn": "csdntest13",
                    "html_url": "https://test.gitcode.net/csdntest13",
                    "tenant_name": null,
                    "is_member": null
                },
                "assigner": {
                    "id": "64c71c3d64037b4af1c7a93f",
                    "login": "green",
                    "name": "百里",
                    "state": null,
                    "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/be/fb/7b9e393fbd80ca315dec249f2be6e6a7378f591609b6525798bc6d95abedc992.png?time=1712128581171",
                    "avatar_path": null,
                    "email": null,
                    "name_cn": null,
                    "html_url": "https://test.gitcode.net/green",
                    "tenant_name": null,
                    "is_member": null
                },
                "private": null,
                "public": null,
                "internal": false
            }
        },
        "id": 70067,
        "iid": 63,
        "project_id": 243377,
        "title": "测试创建PR",
        "body": null,
        "state": "merged",
        "created_at": "2024-04-21T17:35:16.655+08:00",
        "updated_at": "2024-04-24T22:27:49.197+08:00",
        "merged_at": "2024-04-24T22:27:48.631+08:00",
        "closed_by": null,
        "closed_at": null,
        "title_html": null,
        "description_html": null,
        "target_branch": "dev",
        "source_branch": "test_b12",
        "squash_commit_message": null,
        "user_notes_count": 1,
        "upvotes": 0,
        "downvotes": 0,
        "source_project_id": 243377,
        "target_project_id": 243377,
        "work_in_progress": false,
        "milestone": null,
        "merge_when_pipeline_succeeds": false,
        "merge_status": "can_be_merged",
        "sha": "fb6495834d1bf7a39dfdb44ad25e6f83c7136310",
        "merge_commit_sha": "6c93b6e6fcf1ce1f0ce918d1a481f0500531ab72",
        "discussion_locked": null,
        "should_remove_source_branch": false,
        "force_remove_source_branch": false,
        "allow_collaboration": null,
        "allow_maintainer_to_push": null,
        "web_url": "https://test.gitcode.net/One/One/merge_requests/63",
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
        "approval_merge_request_approvers": [
            {
                "id": 233,
                "username": "wunian2011",
                "name": "wunian2011",
                "nick_name": "测试吴",
                "name_cn": "wunian2011",
                "email": null,
                "state": "approve",
                "is_codeowner": false,
                "updated_at": "2024-04-24T21:40:11.095+08:00",
                "avatar_url": null
            },
            {
                "id": 277,
                "username": "renww",
                "name": "renww",
                "nick_name": "介简介简介简介简介简介简介简介简介",
                "name_cn": "renww",
                "email": null,
                "state": "optional",
                "is_codeowner": false,
                "updated_at": "2024-04-21T17:35:18.509+08:00",
                "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ee/dc/7602704ee7dcf13f4383a72d492b1813823afba729ae6e9115877a4a0128d990.jpg?time=1711447395118"
            }
        ],
        "approval_merge_request_testers": [],
        "added_lines": 1,
        "removed_lines": 0,
        "subscribed": true,
        "changes_count": "1",
        "latest_build_started_at": null,
        "latest_build_finished_at": null,
        "first_deployed_to_production_at": null,
        "pipeline": null,
        "diff_refs": {
            "base_sha": "0c02dd57f8945791460a141f155dd2f4bd5dea86",
            "head_sha": "fb6495834d1bf7a39dfdb44ad25e6f83c7136310",
            "start_sha": "b6d44deb0ca73d7a50916d0fea02c72edd6c924e"
        },
        "merge_error": null,
        "json_merge_error": null,
        "rebase_in_progress": null,
        "diverged_commits_count": null,
        "merge_request_reviewer_list": [],
        "merge_request_review_count": 0,
        "merge_request_reviewers_count": 0,
        "notes": 1,
        "unresolved_discussions_count": 0,
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
        "source_git_url": "ssh://git@test.gitcode.net:2222/One/One.git",
        "auto_merge": false
    }
]
```

## 2. 合并Pull Request
### 请求
`PUT https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/merge`
### 参数
| 参数名      | 描述                         | 类型    | 数据类型   |
|----------|----------------------------|-------|--------|
| access_token* | 用户授权码                      | query | string |
| owner*   | 仓库所属空间地址(组织或个人的地址path)     | path  | string |
| repo*    | 仓库路径(path)                 | path  | string |
|  number*  | 第几个PR，即本仓库PR的序数 | path  | string    |
|  merge_method   | 可选。合并PR的方法，merge（合并所有提交）、squash（扁平化分支合并）和rebase（变基并合并）。默认为merge。 | body  | string    |

### 响应
```json
{
  "sha": "c20ac9624d2811a9313af29769dcf581b60c3044",
  "merged": true,
  "message": "Pull Request 已成功合并"
}
```


## 3. 获取pr关联的issue

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/issues`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  number*   | 第几个PR，即本仓库PR的序数 | path | string    |
|  page   | 当前的页码 | query | integer    |
|  per_page   | 每页的数量，最大为 100 | query | integer    |

### 响应

```json
[
  {
    "number": "1",
    "title": "[bug] test",
    "state": "open",
    "title": "进行稳定性测试",
    "url":"https://api.gitcode.com/api/v5/repos/sytest/paopao/issues/1",
    "body": "发生什么问题了？",
    "user": {
      "id": "681",
      "login": "test",
      "name": "test"
    },
    "labels": [
      {
        "color": "#008672",
        "name": "help wanted",
        "id": 381445,
        "title": "help wanted",
        "type": null,
        "textColor": "#FFFFFF"
      }
    ]
  }
]
```

## 4. 提交pull request 评论
### 请求

`POST https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/comments`

### 参数

| 参数名  | 描述  | 类型    | 数据类型  |
| ------ | ------ |-------|------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path  | string    |
|  repo* | 仓库路径(path) | path  | string    |
|  number* | 第几个PR，即本仓库PR的序数 | path  | string    |
|  body* | 评论内容 | body  | string    |
|  path  | 文件的相对路径 | body  | string    |
|  position  | Diff的相对行数 | body  | integer    |


### 响应
```json
{
    "id": "97219c08d421e55cfa841deca16a30f5d7269e10",
    "body":"22222"
}
```



## 5. Pull Request Commit文件列表

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/files`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
| number* | 第几个PR，即本仓库PR的序数 | path | integer |


### 响应

```json
[
    {
        "sha": "45e1211262a0ed24eeb85ac37f7776259ef0e7e1",
        "filename": "README.md",
        "status": null,
        "additions": "3",
        "deletions": "1",
        "blob_url": "https://ra w.gitcode.com/zzero/demo/blob/45e1211262a0ed24eeb85ac37f7776259ef0e7e1/README.md",
        "raw_url": "https://ra w.gitcode.com/zzero/demo/raw/45e1211262a0ed24eeb85ac37f7776259ef0e7e1/README.md",
        "patch": {
            "diff": "@@ -13,4 +13,6 @@ demo\n \r\n > covid_19 一个模拟感染人群爆发的小动画\r\n \r\n-> leetcode 算法解答\n\\ No newline at end of file\n+> leetcode 算法解答\r\n+\r\n+> juc包测试\n\\ No newline at end of file\n",
            "new_path": "README.md",
            "old_path": "README.md",
            "a_mode": "100644",
            "b_mode": "100644",
            "new_file": false,
            "renamed_file": false,
            "deleted_file": false,
            "too_large": false
        }
    },
    {
        "sha": "45e1211262a0ed24eeb85ac37f7776259ef0e7e1",
        "filename": "src/main/java/com/zhzh/sc/demo/juc/lock/VolatileDemo.java",
        "status": null,
        "additions": "3",
        "deletions": "0",
        "blob_url": "https://ra w.gitcode.com/zzero/demo/blob/45e1211262a0ed24eeb85ac37f7776259ef0e7e1/src/main/java/com/zhzh/sc/demo/juc/lock/VolatileDemo.java",
        "raw_url": "https://ra w.gitcode.com/zzero/demo/raw/45e1211262a0ed24eeb85ac37f7776259ef0e7e1/src/main/java/com/zhzh/sc/demo/juc/lock/VolatileDemo.java",
        "patch": {
            "diff": "@@ -15,6 +15,9 @@ public class VolatileDemo {\n         System.out.println(\"service end\");\n     }\n \n+    /**\n+     * 测试方法入口\n+     */\n     public static void main(String[] args) throws InterruptedException {\n         VolatileDemo v = new VolatileDemo();\n         new Thread(v::service, \"thread-1\").start();\n",
            "new_path": "src/main/java/com/zhzh/sc/demo/juc/lock/VolatileDemo.java",
            "old_path": "src/main/java/com/zhzh/sc/demo/juc/lock/VolatileDemo.java",
            "a_mode": "100644",
            "b_mode": "100644",
            "new_file": false,
            "renamed_file": false,
            "deleted_file": false,
            "too_large": false
        }
    }
]
```




## 6. 获取某个Pull Request的所有评论

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/comments`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
| number* | 第几个PR，即本仓库PR的序数 | path | integer |
| page | 当前的页码 | query | integer |
| per_page | 每页的数量，最大为 100 | query | integer |
| direction | 可选。升序/降序(asc/desc) | query | integer |
| comment_type | 可选。筛选评论类型。代码行评论/pr普通评论:diff_comment/pr_comment | query | string |
| end_id | 【*】上一页检视意见最大id（第一页默认为0） | query | integer |


### 响应

```json
[
  {
    "id": "de772738e6dab92174c0e86c052ccf9bed24f747",
    "body": "111",
    "created_at": "2024-04-19T07:48:59.755+00:00",
    "updated_at": "2024-04-19T07:48:59.755+00:00",
    "user": {
      "id": 708,
      "login": "Lzm_0916",
      "name": "Lzm_0916",
      "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/cb/da/6cb18d9ae9f1a94b4f640d3b848351c352c7869f33d0cb68e7acad4f224c4e23.png",
      "html_url": "https://test.gitcode.net/Lzm_0916"
    }
  }
]
```



## 7. 创建Pull Request

### 请求

`POST https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
| title* | 必填。Pull Request 标题 | body | string |
| head* | 必填。Pull Request 提交的源分支。格式：branch 或者：username:branch | body | string |
| base* | 必填。Pull Request 提交目标分支的名称 | body | string |
| body | 可选。Pull Request 内容 | body | string |
| milestone_number | 可选。里程碑序号(id) | body | integer |
| labels | 用逗号分开的标签，名称要求长度在 2-20 之间且非特殊字符。如: bug,performance | body | string |
| issue | 可选。Pull Request的标题和内容可以根据指定的Issue Id自动填充 | body | string |
| assignees | 可选。审查人员username，可多个，半角逗号分隔，如：(username1,username2), 注意: 当仓库代码审查设置中已设置【指派审查人员】则此选项无效 | body | string |
| testers | 可选。测试人员username，可多个，半角逗号分隔，如：(username1,username2), 注意: 当仓库代码审查设置中已设置【指派测试人员】则此选项无效 | body | string |
| prune_source_branch | 可选。合并PR后是否删除源分支，默认false（不删除） | body | boolean |
| draft | 是否设置为草稿 | body | boolean |
| squash | 接受 Pull Request 时使用扁平化（Squash）合并 | body | boolean |
| squash_commit_message | squash提交信息 | body | string |


### 响应

```json
{
    "id": 11264998,
    "url": "https://gitcode.com/api/v5/repos/zzero/demo/pulls/6",
    "html_url": "https://gitcode.com/zzero/demo/pulls/6",
    "diff_url": "https://gitcode.com/zzero/demo/pulls/6.diff",
    "patch_url": "https://gitcode.com/zzero/demo/pulls/6.patch",
    "issue_url": "https://gitcode.com/api/v5/repos/zzero/demo/pulls/6/issues",
    "commits_url": "https://gitcode.com/api/v5/repos/zzero/demo/pulls/6/commits",
    "review_comments_url": "https://gitcode.com/api/v5/repos/zzero/demo/pulls/comments/{/number}",
    "review_comment_url": "https://gitcode.com/api/v5/repos/zzero/demo/pulls/comments",
    "comments_url": "https://gitcode.com/api/v5/repos/zzero/demo/pulls/6/comments",
    "number": 6,
    "title": "测试创建PR",
    "description": "update: 更新文件 dev_001.txt \nupdate: 更新文件 dev_001.txt ",
    "state": "opened",
    "created_at": "2024-04-14T20:53:13.185+08:00",
    "updated_at": "2024-04-14T20:53:22.634+08:00",
    "merged_by": null,
    "merged_at": null,
    "closed_by": null,
    "closed_at": null,
    "title_html": null,
    "description_html": null,
    "target_branch": "test_b5",
    "source_branch": "dev",
    "squash_commit_message": null,
    "user_notes_count": 0,
    "upvotes": 0,
    "downvotes": 0,
    "author": {
        "id": 494,
        "name": "csdntest13",
        "username": "csdntest13",
        "iam_id": "d8b3e018b2364546b946886a669d50fc",
        "nick_name": "csdntest13_gitcode",
        "state": "active",
        "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png?time=1711533165876",
        "avatar_path": null,
        "email": "csdntest13@noreply.gitcode.com",
        "name_cn": "csdntest13",
        "web_url": "https://gitcode.com/csdntest13",
        "tenant_name": null,
        "is_member": null
    },
    "assignee": null,
    "source_project_id": 243377,
    "target_project_id": 243377,
    "labels": [
        {
            "color": "#008672",
            "name": "help wanted",
            "id": 381445,
            "title": "help wanted",
            "type": null,
            "textColor": "#FFFFFF"
        },
        {
            "color": "#CFD240",
            "name": "invalid",
            "id": 381446,
            "title": "invalid",
            "type": null,
            "textColor": "#FFFFFF"
        },
        {
            "color": "#D876E3",
            "name": "question",
            "id": 381447,
            "title": "question",
            "type": null,
            "textColor": "#333333"
        }
    ],
    "work_in_progress": false,
    "milestone": null,
    "merge_when_pipeline_succeeds": false,
    "merge_status": "unchecked",
    "sha": "8da7a5c35e71deeb0bf1d9ecae70449c574749f2",
    "merge_commit_sha": null,
    "discussion_locked": null,
    "should_remove_source_branch": false,
    "force_remove_source_branch": false,
    "allow_collaboration": null,
    "allow_maintainer_to_push": null,
    "web_url": "https://gitcode.com/One/One/merge_requests/53",
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
    "approval_merge_request_reviewers": [
        {
            "id": 43,
            "username": "green",
            "name": "green",
            "nick_name": null,
            "name_cn": "green",
            "email": null,
            "state": "optional",
            "is_codeowner": false,
            "updated_at": "2024-04-14T20:53:23.021+08:00",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/be/fb/7b9e393fbd80ca315dec249f2be6e6a7378f591609b6525798bc6d95abedc992.png?time=1712128581171"
        }
    ],
    "approval_merge_request_approvers": [
        {
            "id": 277,
            "username": "renww",
            "name": "renww",
            "nick_name": null,
            "name_cn": "renww",
            "email": null,
            "state": "optional",
            "is_codeowner": false,
            "updated_at": "2024-04-14T20:53:23.751+08:00",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ee/dc/7602704ee7dcf13f4383a72d492b1813823afba729ae6e9115877a4a0128d990.jpg?time=1711447395118"
        }
    ],
    "approval_merge_request_testers": [
        {
            "id": 43,
            "username": "green",
            "name": "green",
            "nick_name": null,
            "name_cn": "green",
            "email": null,
            "state": "optional",
            "is_codeowner": false,
            "updated_at": "2024-04-14T20:53:23.755+08:00",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/be/fb/7b9e393fbd80ca315dec249f2be6e6a7378f591609b6525798bc6d95abedc992.png?time=1712128581171"
        },
        {
            "id": 277,
            "username": "renww",
            "name": "renww",
            "nick_name": null,
            "name_cn": "renww",
            "email": null,
            "state": "optional",
            "is_codeowner": false,
            "updated_at": "2024-04-14T20:53:23.755+08:00",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ee/dc/7602704ee7dcf13f4383a72d492b1813823afba729ae6e9115877a4a0128d990.jpg?time=1711447395118"
        }
    ],
    "source_project": {
        "id": 243377,
        "description": "csdntest13的第一个项目(公开)",
        "name": "One",
        "name_with_namespace": "One / One",
        "path": "One",
        "path_with_namespace": "One/One",
        "develop_mode": "normal",
        "created_at": "2024-03-19T16:24:01.197+08:00",
        "updated_at": "2024-03-19T16:42:34.834+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "ssh://git@gitcode.com:2222/One/One.git",
        "http_url_to_repo": "https://gitcode.com/One/One.git",
        "web_url": "https://gitcode.com/One/One",
        "readme_url": "https://gitcode.com/One/One/blob/main/README.md",
        "product_id": "28f96caf52004e81ab0bc38d60d11940",
        "product_name": null,
        "member_mgnt_mode": 3,
        "default_branch": "main",
        "tag_list": [],
        "license_url": null,
        "license": {
            "key": "Apache_License_v2.0",
            "name": null,
            "nickname": null,
            "html_url": null,
            "source_url": null
        },
        "avatar_url": null,
        "star_count": 1,
        "forks_count": 0,
        "open_issues_count": 108,
        "open_merge_requests_count": 32,
        "open_change_requests_count": null,
        "watch_count": 1,
        "last_activity_at": "2024-04-14T20:43:58.602+08:00",
        "namespace": {
            "id": 136909,
            "name": "One",
            "path": "One",
            "develop_mode": "normal",
            "region": null,
            "cell": "default",
            "kind": "group",
            "full_path": "One",
            "full_name": "One ",
            "parent_id": null,
            "visibility_level": 20,
            "enable_file_control": null,
            "owner_id": null
        },
        "empty_repo": false,
        "starred": false,
        "visibility": "public",
        "security": "internal",
        "has_updated_kia": false,
        "network_type": "green",
        "owner": null,
        "creator": {
            "id": 494,
            "name": "csdntest13",
            "username": "csdntest13",
            "iam_id": "d8b3e018b2364546b946886a669d50fc",
            "nick_name": "csdntest13_gitcode",
            "state": "active",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png?time=1711533165876",
            "avatar_path": null,
            "email": "csdntest13@noreply.gitcode.com",
            "name_cn": "csdntest13",
            "web_url": "https://gitcode.com/csdntest13",
            "tenant_name": null,
            "is_member": null
        },
        "creator_id": 494,
        "forked_from_project": null,
        "item_type": "Project",
        "main_repository_language": [
            "Text",
            "#cccccc"
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
        "id": 243377,
        "description": "csdntest13的第一个项目(公开)",
        "name": "One",
        "name_with_namespace": "One / One",
        "path": "One",
        "path_with_namespace": "One/One",
        "develop_mode": "normal",
        "created_at": "2024-03-19T16:24:01.197+08:00",
        "updated_at": "2024-03-19T16:42:34.834+08:00",
        "archived": false,
        "is_kia": false,
        "ssh_url_to_repo": "ssh://git@gitcode.com:2222/One/One.git",
        "http_url_to_repo": "https://gitcode.com/One/One.git",
        "web_url": "https://gitcode.com/One/One",
        "readme_url": "https://gitcode.com/One/One/blob/main/README.md",
        "product_id": "28f96caf52004e81ab0bc38d60d11940",
        "product_name": null,
        "member_mgnt_mode": 3,
        "default_branch": "main",
        "tag_list": [],
        "license_url": null,
        "license": {
            "key": "Apache_License_v2.0",
            "name": null,
            "nickname": null,
            "html_url": null,
            "source_url": null
        },
        "avatar_url": null,
        "star_count": 1,
        "forks_count": 0,
        "open_issues_count": 108,
        "open_merge_requests_count": 32,
        "open_change_requests_count": null,
        "watch_count": 1,
        "last_activity_at": "2024-04-14T20:43:58.602+08:00",
        "namespace": {
            "id": 136909,
            "name": "One",
            "path": "One",
            "develop_mode": "normal",
            "region": null,
            "cell": "default",
            "kind": "group",
            "full_path": "One",
            "full_name": "One ",
            "parent_id": null,
            "visibility_level": 20,
            "enable_file_control": null,
            "owner_id": null
        },
        "empty_repo": false,
        "starred": false,
        "visibility": "public",
        "security": "internal",
        "has_updated_kia": false,
        "network_type": "green",
        "owner": null,
        "creator": {
            "id": 494,
            "name": "csdntest13",
            "username": "csdntest13",
            "iam_id": "d8b3e018b2364546b946886a669d50fc",
            "nick_name": "csdntest13_gitcode",
            "state": "active",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png?time=1711533165876",
            "avatar_path": null,
            "email": "csdntest13@noreply.gitcode.com",
            "name_cn": "csdntest13",
            "web_url": "https://gitcode.com/csdntest13",
            "tenant_name": null,
            "is_member": null
        },
        "creator_id": 494,
        "forked_from_project": null,
        "item_type": "Project",
        "main_repository_language": [
            "Text",
            "#cccccc"
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
    "added_lines": 19860,
    "removed_lines": 1,
    "subscribed": true,
    "changes_count": "6",
    "latest_build_started_at": null,
    "latest_build_finished_at": null,
    "first_deployed_to_production_at": null,
    "pipeline": null,
    "diff_refs": {
        "base_sha": "0c02dd57f8945791460a141f155dd2f4bd5dea86",
        "head_sha": "8da7a5c35e71deeb0bf1d9ecae70449c574749f2",
        "start_sha": "fb6495834d1bf7a39dfdb44ad25e6f83c7136310"
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
    "e2e_issues": [
        {
            "id": 13588,
            "issue_type": 7,
            "linked_issue_type": null,
            "issue_num": "issue100",
            "commit_id": null,
            "merge_request_id": 68253,
            "check_fail_reason": "",
            "check_result": true,
            "issue_link": "https://gitcode.com/One/One/issues/100",
            "created_at": "2024-04-14T20:53:23.772+08:00",
            "mks_id": null,
            "pbi_id": null,
            "pbi_name": null,
            "source": null,
            "issue_project_id": 243377,
            "title": "第boudoirripinings-45个issue",
            "issue_project": null,
            "auto_c_when_mr_merged": false
        }
    ],
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
    "source_git_url": "ssh://git@gitcode.com:2222/One/One.git",
    "auto_merge": null
}
```


## 8. 更新Pull Request信息

### 请求

`PATCH https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}`

### 参数

| 参数名                | 描述                                                         | 类型  | 数据类型 |
| ------------------- | ------------------------------------------------------------ | ----- | -------- |
| access_token*       | 用户授权码                                                   | query | string   |
| owner*              | 仓库所属空间地址(组织或个人的地址path)                       | path  | string   |
| repo*               | 仓库路径(path)                                               | path  | string   |
| number*             | 第几个PR，即本仓库PR的序数                                   | path  | integer  |
| title               | 可选。Pull Request 标题                                      | body  | string   |
| body                | 可选。Pull Request 内容                                      | body  | string   |
| state               | 可选。Pull Request 状态                                      | body  | string   |
| milestone_number    | 可选。里程碑序号(id)                                         | body  | integer  |
| labels              | 用逗号分开的标签，名称要求长度在 2-20 之间且非特殊字符。如: bug,performance | body  | string   |
| draft               | 是否设置为草稿                                               | body  | boolean  |


### 响应

```json
{
    "title": "test_b3->dev",
    "body": "new: 新增文件 test_b3 \n2223333444444",
    "state": "opened",
    "created_at": "2024-03-28T22:23:29.999+08:00",
    "updated_at": "2024-04-14T21:06:52.499+08:00"
}
```



## 9. 获取项目下所有pr

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/pulls/related_review_data`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    |
| pr_ids* | pr id | query | string    |


### 响应

```json
[
  {
    "committerName": "杨妮",
    "CommitterEmail": "neen.yang@huawei.com",
    "sUbSystem": "openharmony/docs",
    "isHWEmployee": true,
    "prId": 11135558,
    "pr": 40135,
    "prState": "merged",
    "prLink": "https://gitcode.com/openharmony/docs/pulls/40135",
    "changedAdditions": 2,
    "changedDeletions": 2,
    "issueNo": "",
    "reviewTime": "2024-03-28T11:19:33+08:00",
    "prBugCount": 0,
    "sumPrBugCount": 2,
    "prBugList": {
    },
    "prReviewTime": 0.5,
    "noteCount": 1
  }
]
```


## 10. 获取单个Pull Request
### 请求
`GET  https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  number*   | 第几个PR，即本仓库PR的序数 | path | string    |

### 响应
```json
{
  "id": 111,
  "html_url": "http://gitcode.com/sytest/paopao/pull/1",
  "number": 1,
  "url":"https://api.gitcode.com/api/v5/repos/sytest/paopao/pulls/1",
  "issue_url":"https://api.gitcode.com/api/v5/repos/sytest/paopao/pulls/1/issues",
  "state": "open",
  "assignees_number": 1,
  "assignees": [
    {
      "id": 2,
      "login": "test",
      "name": "test_web",
      "avatar_url": "http://gitcode.com/sytest/paopao/pull/1.png",
      "html_url": "http://gitcode.com/sytest/paopao/pull/1",
      "assigness": true,
      "code_owner": false,
      "accept": true
    }
  ],
  "testers": [
    {
      "id": 2,
      "login": "test",
      "name": "test_web",
      "avatar_url": "http://gitcode.com/sytest/paopao/pull/1.png",
      "html_url": "http://gitcode.com/sytest/paopao/pull/1"
    }
  ],
  "labels": [
    {
      "id": 222,
      "name": "label1",
      "repository_id": 1,
      "created_at": "",
      "updated_at": ""
    }
  ],
  "created_at": "",
  "updated_at": "",
  "closed__at": "",
  "draft": false,
  "merged_at": "",
  "can_merge_check": false,
  "mergeable": true,
  "body": "Description",
  "user": {
    "id": "userId",
    "login": "test"
  },
  "head": {
    "label": "test",
    "ref": "test",
    "sha": "91861a9668041fc1c0ff51d1db66b6297179f5e6",
    "repo":{
      "path":"paopao",
      "name":"paopao"
    }
  },
  "base": {
    "label": "main",
    "ref": "main",
    "sha": "91861a9668041fc1c0ff51d1db66b6297179f5e6",
    "repo":{
      "path":"paopao",
      "name":"paopao"
    }
  }
}
```

## 11. 获取某Pull Request的所有Commit信息
### 请求
`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/commits`

### 参数
| 参数名           | 描述                               | 类型             | 数据类型   |
|---------------|----------------------------------|----------------|--------|
|  access_token* | 用户授权码 | query | string    |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  number*   | 第几个PR，即本仓库PR的序数 | path | string    |

###响应
```json
[
  {
    "sha": "91861a9668041fc1c0ff51d1db66b6297179f5e6",
    "html_url": "https://gitcode.com/sytest/paopao/blob/91861a9668041fc1c0ff51d1db66b6297179f5e6",
    "commit": {
      "author": {
        "name": "test",
        "email": "test@test.com",
        "date": "2024-03-28T11:19:33+08:00"
      },
      "committer": {
        "name": "test",
        "email": "test@test.com"
      },
      "message":"!5 333 * add 1/2/3/4. * add 1/2/3. "
    },
    "author": {
      "id": "id123",
      "login": "test",
      "name": "test",
      "avatar_url": "https://gitcode/pic.png",
      "html_url": "https://gitcode.com/test"
    },
    "committer": {
      "id": "id123",
      "login": "test",
      "name": "test",
      "avatar_url": "https://gitcode/pic.png",
      "html_url": "https://gitcode.com/test"
    },
    "parents":{
      "sha":"2e208a1e38f6a5a7b0cc3787688067ba082a8bb7",
      "shas":["2e208a1e38f6a5a7b0cc3787688067ba082a8bb7"]
    }
  }
]
```

## 12. 创建 Pull Request 标签
### 请求
`POST https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/labels`

### 参数
| 参数名           | 描述                       | 类型             | 数据类型   |
|---------------|--------------------------|----------------|--------|
|  access_token* | 用户授权码                    | query | string |
|  owner* | 仓库所属空间地址(组织或个人的地址path)   | path | string |
|  repo*   | 仓库路径(path)               | path | string |
|  number*   | 第几个PR，即本仓库PR的序数          | path | string |
|  labels*   | 添加的标签 如: ["feat", "bug"] | body | array  |

### 响应
```json
[
  {
    "color": "#008672",
    "name": "help wanted",
    "id": 381445,
    "title": "help wanted",
    "type": null,
    "textColor": "#FFFFFF"
  }
]
```
### Response Code
```text
HTTP status 201 No Content
```

## 13. 删除 Pull Request 标签

### 请求
`DELETE https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/labels/{name}`

### 参数
| 参数名           | 描述                       | 类型             | 数据类型   |
|---------------|--------------------------|----------------|--------|
|  access_token* | 用户授权码                    | query | string |
|  owner* | 仓库所属空间地址(组织或个人的地址path)   | path | string |
|  repo*   | 仓库路径(path)               | path | string |
|  number*   | 第几个PR，即本仓库PR的序数          | path | string |
|  name*   |  标签名称(批量删除用英文逗号分隔，如: bug,feature)    | path | string |

### 响应
```text
HTTP status 204 No Content
```


## 14. 处理 Pull Request 测试
### 请求
`POST https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/test`
### 参数
| 参数名      | 描述                       | 类型    | 数据类型   |
|----------|--------------------------|-------|--------|
| access_token* | 用户授权码                    | query | string |
| owner*   | 仓库所属空间地址(组织或个人的地址path)   | path  | string |
| repo*    | 仓库路径(path)               | path  | string |
| force      | 是否强制测试通过（默认否），只对管理员生效          | body  | 	boolean |

### 响应
```text
HTTP status 204 No Content
```

## 15. 处理 Pull Request 审查
### 请求
`POST https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/review`
### 参数
| 参数名      | 描述                     | 类型    | 数据类型   |
|----------|------------------------|-------|--------|
| access_token* | 用户授权码                  | query | string |
| owner*   | 仓库所属空间地址(组织或个人的地址path) | path  | string |
| repo*    | 仓库路径(path)             | path  | string |
| force      | 是否强制审查通过（默认否），只对管理员生效  | body  | 	boolean |

### 响应
```text
HTTP status 204 No Content
```

## 16. 获取某个Pull Request的操作日志
### 请求
`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/operate_logs`
### 参数
| 参数名      | 描述                     | 类型    | 数据类型   |
|----------|------------------------|-------|--------|
| access_token* | 用户授权码                  | query | string |
| owner*   | 仓库所属空间地址(组织或个人的地址path) | path  | string |
| repo*    | 仓库路径(path)             | path  | string |
| number*      | 第几个PR，即本仓库PR的序数  | body  | 	boolean |
| sort     | 按递减(desc)排序，默认：递减  | query  | 	boolean |

### 响应
```json
[
    {
        "content": "Create mr issue link: **第boudoirripinings-24个issue** #79",
        "id": 274531,
        "action": "add_mr_issue_link",
        "merge_request_id": 70067,
        "created_at": "2024-04-23T11:32:08.522+08:00",
        "updated_at": "2024-04-23T11:32:08.522+08:00",
        "discussion_id": "18a5ab21f57cda175b8eabc2ec829a9e04d4d458",
        "project": "One/One",
        "assignee": null,
        "proposer": null,
        "user": {
            "id": "65f94ab6f21fa3084fc04823",
            "name": "csdntest13",
            "login": "csdntest13",
            "iam_id": "d8b3e018b2364546b946886a669d50fc",
            "nick_name": "csdntest13_gitcode",
            "state": "active",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png?time=1711533165876",
            "avatar_path": null,
            "email": "csdntest13@noreply.gitcode.com",
            "name_cn": "csdntest13",
            "web_url": "https://test.gitcode.net/csdntest13",
            "tenant_name": null,
            "is_member": null
        }
    },
    {
        "content": "Create mr issue link: **第boudoirripinings-25个issue** #80",
        "id": 274529,
        "action": "add_mr_issue_link",
        "merge_request_id": 70067,
        "created_at": "2024-04-23T11:32:07.588+08:00",
        "updated_at": "2024-04-23T11:32:07.588+08:00",
        "discussion_id": "9b4b01dbe059dbdc120afd8bdf9fd865d4ea42b1",
        "project": "One/One",
        "assignee": null,
        "proposer": null,
        "user": {
            "id": "65f94ab6f21fa3084fc04823",
            "name": "csdntest13",
            "login": "csdntest13",
            "iam_id": "d8b3e018b2364546b946886a669d50fc",
            "nick_name": "csdntest13_gitcode",
            "state": "active",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png?time=1711533165876",
            "avatar_path": null,
            "email": "csdntest13@noreply.gitcode.com",
            "name_cn": "csdntest13",
            "web_url": "https://test.gitcode.net/csdntest13",
            "tenant_name": null,
            "is_member": null
        }
    }
]
```

## 17. 获取某个 Pull Request 的所有标签
### 请求
`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/labels`

### 参数
| 参数名           | 描述                               | 类型             | 数据类型   |
|---------------|----------------------------------|----------------|--------|
|  access_token* | 用户授权码 | query | string    |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  number*   | 第几个PR，即本仓库PR的序数 | path | string    |
|  page   | 当前的页码 | query | integer    |
|  per_page   | 每页的数量，最大为 100 | query | integer    |

###响应
```json
[
  {
    "id": 18517,
    "color": "#ED4014",
    "name": "bug",
    "repository_id": 198606,
    "url": "",
    "created_at": "2024-02-23",
    "updated_at": "2024-02-23",
    "text_color": "#FFFFFF"
  },
  {
    "id": 383740,
    "color": "#428BCA",
    "name": "performance",
    "repository_id": 198606,
    "url": "",
    "created_at": "2024-04-20",
    "updated_at": "2024-04-20",
    "text_color": "#FFFFFF"
  }
]
```

## 18. 重置 Pull Request 测试 的状态
### 请求
`PATCH https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/testers`
### 参数
| 参数名      | 描述                       | 类型    | 数据类型   |
|----------|--------------------------|-------|--------|
| access_token* | 用户授权码                    | query | string |
| owner*   | 仓库所属空间地址(组织或个人的地址path)   | path  | string |
| repo*    | 仓库路径(path)               | path  | string |
| reset_all      | 是否重置所有测试人，默认：false，只对管理员生效          | body  | 	boolean |

### 响应
```text
HTTP status 204 No Content
```

## 19. 重置 Pull Request 审查 的状态
### 请求
`PATCH https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/assignees`
### 参数
| 参数名      | 描述                         | 类型    | 数据类型   |
|----------|----------------------------|-------|--------|
| access_token* | 用户授权码                      | query | string |
| owner*   | 仓库所属空间地址(组织或个人的地址path)     | path  | string |
| repo*    | 仓库路径(path)                 | path  | string |
| reset_all      | 是否重置所有审查人，默认：false，只对管理员生效 | body  | 	boolean |

### 响应
```text
HTTP status 204 No Content
```


## 20. pr提交的文件变更信息

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/files.json`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  number*   | 第几个PR，即本仓库PR的序数 | path | string    |

### 响应

```json
{
    "code": 0,
    "added_lines": 2,
    "removed_lines": 0,
    "count": "3",
    "diff_refs": {
        "base_sha": "79548c6fd379d6f2e0574341255c95d1e0c7760c",
        "start_sha": "79548c6fd379d6f2e0574341255c95d1e0c7760c",
        "head_sha": "bb8fd1bd1ee6cfa5a7848839d0dd0195e7e4ce4d"
    },
    "diffs": [
        {
            "new_blob_id": "45e071c0bc5e62f63730b54a8375f45b8356379b",
            "statistic": {
                "type": "new_file",
                "path": "Photos Library.photoslibrary.zip",
                "old_path": "Photos Library.photoslibrary.zip",
                "new_path": "Photos Library.photoslibrary.zip",
                "viewed": false
            },
            "head": {
                "type": "none_deleted_file",
                "url": "https://raw.gitcode.com/songyi1995/paopao/raw/bb8fd1bd1ee6cfa5a7848839d0dd0195e7e4ce4d/Photos%20Library.photoslibrary.zip",
                "commit_id": "bb8fd1bd1ee6cfa5a7848839d0dd0195e7e4ce4d",
                "commit_short_id": "bb8fd1b"
            },
            "added_lines": 0,
            "removed_lines": 0
        }
    ]
}
```

## 21. 获取文件内容

### 请求
`GET https://raw.gitcode.com/{owner}/{repo}/raw/{head_sha}/{name}`

### 注意
```text
请求20 pr提交的文件变更信息的接口,拿到diffs.head.url返回的地址信息 直接浏览器访问即可。
```

## 22. 企业 Pull Request 列表
### 请求
`GET https://api.gitcode.com/api/v5/enterprise/{enterprise}/pull_requests`

### 参数
| 参数名           | 描述  | 类型  | 数据类型  |
|---------------| ------ | ------  |------|
| access_token* | 用户授权码 | query | string    |
| enterprise*   | 企业的路径(path/login) | path | string    |
| repo          |可选。仓库路径(path) | query | string    |
| state         |可选。Pull Request 状态 | query | string    |
| sort        |可选。排序字段，默认按创建时间 | query | string    |
| direction        |可选。升序/降序 | query | string    |
| page        |当前的页码 | query | string    |
| per_page        |每页的数量，最大为 100 | query | string    |

### 响应

```json
[
  {
    "id": 71020,
    "url": "https://test.gitcode.net/api/v5/repos/test/test/1",
    "html_url": "https://test.gitcode.net/test/test/1",
    "number": 1,
    "state": "merged",
    "assignees_number": 0,
    "testers_number": 0,
    "assignees": [],
    "testers": [],
    "mergeable": null,
    "can_merge_check": true,
    "head": {
      "ref": "main",
      "sha": "d874402d259744a00121c2cff0febc8554339aef",
      "repo": {
        "path": "test",
        "name_space": {
          "path": "repo-dev"
        },
        "assigner": {
          "id": "uuid",
          "login": "Lzm_0916",
          "name": "Lzm_0916"
        }
      }
    },
    "base": {
      "ref": null,
      "sha": null,
      "repo": {
        "path": "test",
        "name_space": {
          "path": "repo-dev"
        }
      }
    }
  }
]
```
