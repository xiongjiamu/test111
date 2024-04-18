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

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | query | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | formData | string    |
|  state | 可选。Pull Request 状态 | query | string    |
|  base   | 可选。Pull Request 提交目标分支的名称 | query | string    |
|  since | 可选。起始的更新时间，要求时间格式为 ISO 8601 | quey | string    |
|  direction  | 可选。升序/降序 | query | string    |
|  milestone_number | 可选。里程碑序号(id) | formData | integer    |
|  page   | 当前的页码 | query | integer    |
|  per_page   | 每页的数量，最大为 100 | query | integer    |
|  author   | 可选。PR 创建者用户名 | query | string    |
|  assignee   |  可选。评审者用户名 | query | string    |
|  tester   |  可选。测试者用户名 | query | string    |

### 响应

```json
[
    {
        "id": 11217119,
        "html_url": "https://gitcode.com/OpenHarmony/ability_ability_base/pulls/341",
        "number": 341,
        "close_related_issue": 1,
        "prune_branch": false,
        "state": "open",
        "assignees_number": 1,
        "testers_number": 1,
        "assignees": [
            {
                "id": 7476845,
                "login": "ccllee",
                "name": "ccll",
                "avatar_url": null,
                "html_url": "https://gitcode.com/ccllee",
                "type": "User",
                "assignee": true,
                "code_owner": false,
                "accept": false
            }
        ],
        "milestone": null,
        "labels": [
            {
                "id": 113667295,
                "color": "20c22e",
                "name": "dco检查成功",
                "repository_id": 6511493,
                "created_at": "2021-06-18T15:08:35+08:00",
                "updated_at": "2024-04-08T21:46:16+08:00"
            }
        ],
        "created_at": "2024-04-08T19:54:38+08:00",
        "updated_at": "2024-04-08T21:48:05+08:00",
        "closed_at": null,
        "draft": false,
        "merged_at": null,
        "mergeable": true,
        "can_merge_check": false,
        "user": {
            "id": 9221401,
            "login": "hust-yf",
            "name": "yangfei",
            "avatar_url": null,
            "html_url": "https://gitcode.com/hust-yf",
            "type": "User"
        },
        "title": "新增configuration key支持sa对单个应用设置深浅色",
        "body": "**IssueNo**:Result**:",
        "head": {
			"label": "mas",
            "user": {
                "id": 9221401,
                "login": "hust-yf",
                "name": "yangfei",
                "avatar_url": null,
                "html_url": "https://gitcode.com/hust-yf",
                "type": "User"
            },
            "repo": {
                "id": 34079500,
                "full_name": "hjj-code_1/ability_ability_base",
                "human_name": "yffff/ability_ability_base",
                "path": "ability_ability_base",
                "name": "ability_ability_base",
                "assigner": {
                    "id": 9221401,
                    "login": "hust-yf",
                    "name": "yangfei",
                    "avatar_url": null,
                    "html_url": "https://gitcode.com/hust-yf",
                    "type": "User"
                },
                "description": "暂无描述",
                "internal": false,
                "html_url": "https://gitcode.com/hjj-code_1/ability_ability_base.git"
            }
        },
        "base": {
            "label": "master",
            "user": {
                "id": 626123,
                "login": "landwind",
                "name": "mamingshuai",
                "avatar_url": null,
                "html_url": "https://gitcode.com/landwind",
                "type": "User"
            },
            "repo": {
                "id": 22974255,
                "full_name": "openharmony/ability_ability_base",
                "human_name": "OpenHarmony/ability_ability_base",
                "path": "ability_ability_base",
                "name": "ability_ability_base",
                "assigner": {
                    "id": 626123,
                    "login": "landwind",
                    "name": "mamingshuai",
                    "avatar_url": null,
                    "html_url": "https://gitcode.com/landwind",
                    "type": "User"
                },
                "description": "暂无描述",
                "internal": false,
                "html_url": "https://gitcode.com/openharmony/ability_ability_base.git"
            }
        }
    }
]

```


## 2. pr提交的文件变更信息

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
                "url": "  /songyi1995/paopao/blob/bb8fd1bd1ee6cfa5a7848839d0dd0195e7e4ce4d/Photos%20Library.photoslibrary.zip",
                "commit_id": "bb8fd1bd1ee6cfa5a7848839d0dd0195e7e4ce4d",
                "commit_short_id": "bb8fd1b"
            },
            "added_lines": 0,
            "removed_lines": 0
        }
    ]
}
```

## 3. 获取pr关联的issue

### 请求

`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/issue`

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
    "title": "[bug] test"
  }
]
```

## 4. 提交pull request 评论
### 请求

`POST https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/comments`

### 参数

| 参数名  | 描述  | 类型  | 数据类型  |
| ------ | ------ | ------  |------|
|  access_token* | 用户授权码 | formData | string    | 
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  number*   | 第几个PR，即本仓库PR的序数 | path | string    |
|  body*   | 评论内容 | formData | string    |


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
| direction | 可选。升序/降序 | query | integer |
| comment_type | 可选。筛选评论类型。代码行评论/pr普通评论 | query | string |
| end_id | 【*】上一页检视意见最大id（第一页默认为0） | query | integer |


### 响应

```json
[
    {
        "url": "https://api.gitcode.com/api/v5/repos/zzero/demo/pulls/comments/26594610",
        "id": 26594610,
        "path": null,
        "position": null,
        "original_position": null,
        "new_line": null,
        "commit_id": null,
        "original_commit_id": null,
        "user": {
            "id": 76523,
            "login": "zzero",
            "name": "insight",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png",
            "url": "https://gitcode.com/api/v5/users/zzero",
            "html_url": "https://gitcode.com/zzero",
            "state": "active",
            "email": "",
            "is_member": ""
        },
        "created_at": "2024-04-14T15:46:37+08:00",
        "updated_at": "2024-04-14T15:46:37+08:00",
        "body": "第一次评论，有点紧张",
        "html_url": "https://gitcode.com/zzero/demo/pulls/1#note_26594610",
        "comment_type": "pr_comment"
    },
    {
        "url": "https://gitcode.com/api/v5/repos/zzero/demo/pulls/comments/26595312",
        "id": 26595312,
        "path": "README.md",
        "position": null,
        "original_position": null,
        "new_line": 18,
        "commit_id": null,
        "original_commit_id": null,
        "user": {
            "id": 76523,
            "login": "zzero",
            "name": "insight",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ec/ba/4e7c4661b6154a7dd088d9fe64b4893383a2e318bf362350ce18d44df6ac7e37.png",
            "url": "https://gitcode.com/api/v5/users/zzero",
            "html_url": "https://gitcode.com/zzero",
            "state": "active",
            "email": "",
            "is_member": ""
        },
        "created_at": "2024-04-14T16:33:51+08:00",
        "updated_at": "2024-04-14T16:33:51+08:00",
        "body": "代码行评论",
        "html_url": "https://gitcode.com/zzero/demo/pulls/1#note_26595312",
        "pull_request_url": "https://gitcode.com/api/v5/repos/zzero/demo/pulls/1",
        "comment_type": "diff_comment"
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




## 17. 更新Pull Request信息

### 请求

`PATCH https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}`

### 参数

| 参数名                | 描述                                                         | 类型  | 数据类型 |
| --------------------- | ------------------------------------------------------------ | ----- | -------- |
| access_token*         | 用户授权码                                                   | query | string   |
| owner*                | 仓库所属空间地址(组织或个人的地址path)                       | path  | string   |
| repo*                 | 仓库路径(path)                                               | path  | string   |
| number*               | 第几个PR，即本仓库PR的序数                                   | path  | integer  |
| title*                | 必填。Pull Request 标题                                      | body  | string   |
| body                  | 可选。Pull Request 内容                                      | body  | string   |
| state                 | 可选。Pull Request 状态                                      | body  | string   |
| milestone_number      | 可选。里程碑序号(id)                                         | body  | integer  |
| labels                | 用逗号分开的标签，名称要求长度在 2-20 之间且非特殊字符。如: bug,performance | body  | string   |
| issue                 | 可选。Pull Request的标题和内容可以根据指定的Issue Id自动填充 | body  | string   |
| assignees             | 可选。审查人员username，可多个，半角逗号分隔，如：(username1,username2), 注意: 当仓库代码审查设置中已设置【指派审查人员】则此选项无效 | body  | string   |
| testers               | 可选。测试人员username，可多个，半角逗号分隔，如：(username1,username2), 注意: 当仓库代码审查设置中已设置【指派测试人员】则此选项无效 | body  | string   |
| prune_source_branch   | 可选。合并PR后是否删除源分支，默认false（不删除）            | body  | boolean  |
| draft                 | 是否设置为草稿                                               | body  | boolean  |
| squash                | 接受 Pull Request 时使用扁平化（Squash）合并                 | body  | boolean  |
| squash_commit_message | squash提交信息                                               | body  | string   |


### 响应

```json
{
    "id": 62874,
    "iid": 27,
    "project_id": 243377,
    "title": "test_b3->dev",
    "description": "new: 新增文件 test_b3 \n2223333444444",
    "state": "opened",
    "created_at": "2024-03-28T22:23:29.999+08:00",
    "updated_at": "2024-04-14T21:06:52.499+08:00",
    "merged_by": null,
    "merged_at": null,
    "closed_by": null,
    "closed_at": null,
    "title_html": null,
    "description_html": null,
    "target_branch": "main",
    "source_branch": "test_b3",
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
    "labels": [],
    "work_in_progress": false,
    "milestone": null,
    "merge_when_pipeline_succeeds": false,
    "merge_status": "unchecked",
    "sha": "618e3dfaa9e017f0389efd01b6aa25795ba2742f",
    "merge_commit_sha": null,
    "discussion_locked": null,
    "should_remove_source_branch": false,
    "force_remove_source_branch": false,
    "allow_collaboration": null,
    "allow_maintainer_to_push": null,
    "web_url": "https://gitcode.com/One/One/merge_requests/27",
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
    "approval_merge_request_approvers": [
        {
            "id": 58,
            "username": "yx_test",
            "name": "yx_test",
            "nick_name": null,
            "name_cn": "yx_test",
            "email": null,
            "state": "optional",
            "is_codeowner": false,
            "updated_at": "2024-03-28T22:23:40.724+08:00",
            "avatar_url": null
        },
        {
            "id": 233,
            "username": "wunian2011",
            "name": "wunian2011",
            "nick_name": null,
            "name_cn": "wunian2011",
            "email": null,
            "state": "optional",
            "is_codeowner": false,
            "updated_at": "2024-03-28T22:23:40.724+08:00",
            "avatar_url": null
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
            "updated_at": "2024-03-28T22:23:40.735+08:00",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/be/fb/7b9e393fbd80ca315dec249f2be6e6a7378f591609b6525798bc6d95abedc992.png?time=1712128581171"
        },
        {
            "id": 58,
            "username": "yx_test",
            "name": "yx_test",
            "nick_name": null,
            "name_cn": "yx_test",
            "email": null,
            "state": "optional",
            "is_codeowner": false,
            "updated_at": "2024-03-28T22:23:40.735+08:00",
            "avatar_url": null
        },
        {
            "id": 233,
            "username": "wunian2011",
            "name": "wunian2011",
            "nick_name": null,
            "name_cn": "wunian2011",
            "email": null,
            "state": "optional",
            "is_codeowner": false,
            "updated_at": "2024-03-28T22:23:40.735+08:00",
            "avatar_url": null
        },
        {
            "id": 437,
            "username": "wanzun",
            "name": "wanzun",
            "nick_name": null,
            "name_cn": "wanzun",
            "email": null,
            "state": "optional",
            "is_codeowner": false,
            "updated_at": "2024-03-28T22:23:40.735+08:00",
            "avatar_url": "https://gitcode-img.obs.cn-south-1.myhuaweicloud.com:443/ae/dc/135a26e881d62c81e3034e9d7cd8aec053934209ffdd644938497f130ccfbe2d.jpeg?time=1711339131224"
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
    "added_lines": 1,
    "removed_lines": 0,
    "subscribed": true,
    "changes_count": "1",
    "latest_build_started_at": null,
    "latest_build_finished_at": null,
    "first_deployed_to_production_at": null,
    "pipeline": null,
    "diff_refs": {
        "base_sha": "fb6495834d1bf7a39dfdb44ad25e6f83c7136310",
        "head_sha": "618e3dfaa9e017f0389efd01b6aa25795ba2742f",
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
            "id": 13591,
            "issue_type": 7,
            "linked_issue_type": null,
            "issue_num": "issue2",
            "commit_id": null,
            "merge_request_id": 62874,
            "check_fail_reason": "",
            "check_result": true,
            "issue_link": "https://gitcode.com/One/One/issues/2",
            "created_at": "2024-04-14T21:06:52.492+08:00",
            "mks_id": null,
            "pbi_id": null,
            "pbi_name": null,
            "source": null,
            "issue_project_id": 243377,
            "title": "第二个issue你是个明白人，我明白你明白的意思。我也是明白人，明白人就应该明白我明白你明白的意思。只要大家都明白明白人应该明白我明白你明白的意思，这样网络环境就是充满明白人所明白。你是个明白人，我明",
            "issue_project": null,
            "auto_c_when_mr_merged": false
        },
        {
            "id": 13590,
            "issue_type": 7,
            "linked_issue_type": null,
            "issue_num": "issue1",
            "commit_id": null,
            "merge_request_id": 62874,
            "check_fail_reason": "",
            "check_result": true,
            "issue_link": "https://gitcode.com/One/One/issues/1",
            "created_at": "2024-04-14T21:06:52.482+08:00",
            "mks_id": null,
            "pbi_id": null,
            "pbi_name": null,
            "source": null,
            "issue_project_id": 243377,
            "title": "这是一个issue",
            "issue_project": null,
            "auto_c_when_mr_merged": false
        },
        {
            "id": 13589,
            "issue_type": 7,
            "linked_issue_type": null,
            "issue_num": "issue1",
            "commit_id": null,
            "merge_request_id": 62874,
            "check_fail_reason": "",
            "check_result": true,
            "issue_link": "https://gitcode.com/WhatCanIDo/W_One/issues/1",
            "created_at": "2024-04-14T21:06:52.478+08:00",
            "mks_id": null,
            "pbi_id": null,
            "pbi_name": null,
            "source": null,
            "issue_project_id": 256563,
            "title": "问题很大",
            "issue_project": {
                "id": 256563,
                "description": "",
                "name": "W_One",
                "name_with_namespace": "我能怎么办呢 / W_One",
                "path": "W_One",
                "path_with_namespace": "WhatCanIDo/W_One",
                "develop_mode": "normal",
                "created_at": "2024-03-28T17:14:28.241+08:00",
                "updated_at": "2024-03-28T17:14:28.241+08:00",
                "archived": false,
                "is_kia": false,
                "ssh_url_to_repo": "ssh://git@gitcode.com:2222/WhatCanIDo/W_One.git",
                "http_url_to_repo": "https://gitcode.com/WhatCanIDo/W_One.git",
                "web_url": "https://gitcode.com/WhatCanIDo/W_One",
                "readme_url": null,
                "product_id": "bb82c64b4f1c4e759b2e81d7f4621030",
                "product_name": null,
                "member_mgnt_mode": 3
            },
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



## 18. 获取项目下所有pr

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


## 19. 获取单个Pull Request
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
  "state": "open",
  "assignees_number": 1,
  "assignees": [
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
  "user": {
    "id": "userId",
    "login": "test"
  },
  "head": {
    "repo": {
      "assigner": {
        "id": "userId",
        "login": "test",
        "name": "test_web",
        "avatar_url": "http://gitcode.com/sytest/paopao/pull/1.png",
        "html_url": "http://gitcode.com/sytest/paopao/pull/1"
      }
    }
  },
  "base": {
    "ref": "main",
    "repo": {
      "assigner": {
        "id": "userId",
        "login": "test",
        "name": "test_web",
        "avatar_url": "http://gitcode.com/sytest/paopao/pull/1.png",
        "html_url": "http://gitcode.com/sytest/paopao/pull/1"
      }
    }
  }
}
```

## 20. 创建commit评论

### 请求

`POST https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/commits`

### 参数

| 参数名           | 描述                               | 类型             | 数据类型   |
|---------------|----------------------------------|----------------|--------|
| access_token* | 用户授权码                            | formData          | string | 
| owner*        | 仓库所属空间地址(组织或个人的地址path)           | path           | string |
| repo*         | 仓库路径(path)                       | path           | string |
| sha*          | 评论的sha值                          | path           | string |
| body*         | 评论的内容                            | formData       | string |
| path          | 文件的相对路径                          | formData       | string |
| position      | Diff的相对行数                        | formData       | string |

### 响应
```json
{
  "id": "12312sadsa",
  "created_at": "2024-03-28T11:19:33+08:00",
  "updated_at": "2024-03-28T11:19:33+08:00"
}
```

## 21. 获取某Pull Request的所有Commit信息
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
      }
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
    }
  }
]
```
## 22. 获取Pull Request关联的issue

### 请求
`GET https://api.gitcode.com/api/v5/repos/{owner}/{repo}/pulls/{number}/issues`

### 参数
| 参数名           | 描述                               | 类型             | 数据类型   |
|---------------|----------------------------------|----------------|--------|
|  access_token* | 用户授权码 | query | string    |
|  owner* | 仓库所属空间地址(组织或个人的地址path) | path | string    |
|  repo*   | 仓库路径(path) | path | string    |
|  number*   | 第几个PR，即本仓库PR的序数 | path | string    |

### 响应
```json
[
  {
    "number": "id"
  }
]
```
