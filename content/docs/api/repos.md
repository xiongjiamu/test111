---
linkTitle: 项目 API
title: 项目接口API文档
weight: 4
sidebar:
  open: false
---

### 1. 创建项目  
`POST /api/v4/projects`  
#### 参数  
| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| name | body | string | The name of the project | no |
| path | body | string | The path of the repository | no |
| description | body | string | The description of the project | no |
| visibility | body | string, Enum(private,public) | The visibility level of the project. | no |
| initialize_with_readme | body | boolean | Initialize a project with a README.md | no |
| initialize_with_license | body | string | Initialize a project with a license | no |
| initialize_with_gitignore | body | string | Initialize a project with a gitignore | no |
| namespace_id | body | integer | Namespace ID for the new project. Default to the user namespace. | no |
| import_url | body | string | URL from which the project is imported | no |
| mirror_repository | body | boolean | is mirror repo | no |

#### 响应  
```json
{
  "id": 98485,
  "description": null,
  "name": "repocc12",
  "name_with_namespace": "group1111 / repocc12",
  "path": "repocc12",
  "path_with_namespace": "group11111/repocc12",
  "develop_mode": "normal",
  "created_at": "2023-09-15T16:45:09.502+08:00",
  "updated_at": "2023-09-15T16:45:09.502+08:00",
  "archived": false,
  "ssh_url_to_repo": "ssh://git@gitcode-backend.cn-north-7.myhuaweicloud.com:2222/group11111/repocc12.git",
  "http_url_to_repo": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/group11111/repocc12.git",
  "web_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/group11111/repocc12",
  "readme_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/group11111/repocc12/blob/main/README.md",
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
  "open_issues_count": null,
  "open_merge_requests_count": null,
  "open_change_requests_count": null,
  "watch_count": 0,
  "last_activity_at": "2023-09-15T16:45:09.502+08:00",
  "namespace": {
    "id": 74962,
    "name": "group1111",
    "path": "group11111",
    "develop_mode": "normal",
    "region": null,
    "cell": "default",
    "kind": "group",
    "full_path": "group11111",
    "full_name": "group1111",
    "parent_id": null,
    "visibility_level": 20,
    "enable_file_control": false,
    "owner_id": null
  },
  "empty_repo": null,
  "starred": null,
  "visibility": "public",
  "owner": null,
  "creator": null,
  "creator_id": null,
  "forked_from_project": null,
  "item_type": null,
  "main_repository_language": null
}
```
### 2. 项目详情  
`GET /api/v4/projects/{project_id}`  
#### 参数  
| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| project_id | path | string | The ID of a project | yes |
| statistics | query | boolean | Include project statistics | no |
| license | query | boolean | Include project license data | no |
| view | query | string | If simple, only returns the simplest info, basic will return basic project info, all will return all info | no |

#### 响应  
```json
{
  "id": 98485,
  "description": null,
  "name": "repocc12",
  "name_with_namespace": "group1111 / repocc12",
  "path": "repocc12",
  "path_with_namespace": "group11111/repocc12",
  "develop_mode": "normal",
  "created_at": "2023-09-15T16:45:09.502+08:00",
  "updated_at": "2023-09-15T16:45:09.502+08:00",
  "archived": false,
  "ssh_url_to_repo": "ssh://git@gitcode.com/group11111/repocc12.git",
  "http_url_to_repo": "https://gitcode.com/group11111/repocc12.git",
  "web_url": "https://gitcode.com/group11111/repocc12",
  "readme_url": "https://gitcode.com/group11111/repocc12/blob/main/README.md",
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
  "last_activity_at": "2023-09-15T18:44:19.163+08:00",
  "namespace": {
    "id": 74962,
    "name": "group1111",
    "path": "group11111",
    "develop_mode": "normal",
    "region": null,
    "cell": "default",
    "kind": "group",
    "full_path": "group11111",
    "full_name": "group1111 ",
    "parent_id": null,
    "visibility_level": 20,
    "enable_file_control": null,
    "owner_id": null
  },
  "empty_repo": false,
  "starred": false,
  "visibility": "public",
  "owner": null,
  "creator": {
    "id": 80,
    "name": "CodeHub_beta_dev",
    "username": "CodeHub_beta_dev",
    "iam_id": "c369c68f1ff84679b5a8ed904d8bff1c",
    "nick_name": "",
    "state": "active",
    "avatar_url": "1111",
    "avatar_path": null,
    "email": "CodeHub_beta_dev@huawei.com",
    "name_cn": "CodeHub_beta_dev",
    "web_url": "https://gitcode.com/CodeHub_beta_dev",
    "tenant_name": null
  },
  "creator_id": 80,
  "forked_from_project": null,
  "item_type": "Project",
  "main_repository_language": null,
  "statistics": null,
  "branch_count": 1,
  "tag_count": 0,
  "member_count": 1,
  "repo_replica_urls": null,
  "open_external_wiki": true,
  "release_count": 0,
  "_links": {
    "self": "https://gitcode.com/api/v4/projects/98485.json",
    "issues": "https://gitcode.com/api/v4/projects/98485/issues.json",
    "merge_requests": "https://gitcode.com/api/v4/projects/98485/merge_requests",
    "repo_branches": "https://gitcode.com/api/v4/projects/98485/repository/branches.json",
    "labels": "https://gitcode.com/api/v4/projects/98485/labels.json",
    "events": "https://gitcode.com/api/v4/projects/98485/events.json",
    "members": "https://gitcode.com/api/v4/projects/98485/members.json"
  },
  "resolve_outdated_diff_discussions": false,
  "container_registry_enabled": true,
  "issues_enabled": true,
  "merge_requests_enabled": true,
  "wiki_enabled": false,
  "jobs_enabled": true,
  "snippets_enabled": true,
  "shared_runners_enabled": true,
  "lfs_enabled": true,
  "import_status": "none",
  "fork_status": "none",
  "import_url": null,
  "import_error": null,
  "runners_token": null,
  "public_jobs": true,
  "ci_config_path": null,
  "shared_with_groups": null,
  "only_allow_merge_if_pipeline_succeeds": false,
  "request_access_enabled": false,
  "only_allow_merge_if_all_discussions_are_resolved": false,
  "printing_merge_request_link_enabled": true,
  "merge_method": "merge",
  "manifest_view_enable": false,
  "main_zone": "cn-north-7-op",
  "partition": null,
  "fork_network_projects": [],
  "permissions": {
    "project_access": null,
    "group_access": null
  }
}
```
### 3. 获取项目保护分支列表  
`GET /api/v4/projects/{project_id}/protected_branches`  
#### 参数  
| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| project_id | path | string | The ID of a project | yes |
| search | query | string | Return list of branches matching the search criteria | no |
| view | query | string | If simple, only returns the name | no |
| page | query | integer | Current page number | no |
| per_page | query | integer | Number of items per page | no |

#### 响应  
```json
{
    "total": 2,
    "protected_branches": [
        {
            "name": "b1",
            "updated_at": "2023-09-19T20:11:51.948+08:00",
            "push_users": [],
            "push_users_access_level_count": null,
            "merge_users": [],
            "merge_users_access_level_count": null,
            "merged": false,
            "protected": true,
            "developers_can_push": false,
            "developers_can_merge": true,
            "committer_can_push": false,
            "committer_can_merge": false,
            "master_can_push": false,
            "master_can_merge": false,
            "maintainer_can_push": false,
            "maintainer_can_merge": false,
            "owner_can_push": true,
            "owner_can_merge": true,
            "no_one_can_push": false,
            "no_one_can_merge": false
        },
        {
            "name": "b2",
            "updated_at": "2023-09-19T20:11:51.948+08:00",
            "push_users": [],
            "push_users_access_level_count": null,
            "merge_users": [],
            "merge_users_access_level_count": null,
            "merged": false,
            "protected": true,
            "developers_can_push": false,
            "developers_can_merge": true,
            "committer_can_push": false,
            "committer_can_merge": false,
            "master_can_push": false,
            "master_can_merge": false,
            "maintainer_can_push": false,
            "maintainer_can_merge": false,
            "owner_can_push": true,
            "owner_can_merge": true,
            "no_one_can_push": false,
            "no_one_can_merge": false
        }
    ]
}
```
### 4. 批量创建保护分支  
`POST /api/v4/projects/{project_id}/batch/protected_branches`  
#### 参数  
| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| project_id | path | string | The ID of a project | yes |
| names | body | array | The names of the branches | yes |
| access_level_string | body | AccessLevelString |  | no |

#### AccessLevelString  
| 参数名 | 类型 | 描述 | 必选 |
| :----: | :---------: | :------: | :------: |
| push_access_levels | string | Access levels allowed to push, split by "," | no |
| merge_access_levels | string | Access levels allowed to merge, split by "," | no |

#### 响应  
```json
[
    "b1",
    "b2"
]
```
### 5. 修改保护分支  
`PUT /api/v4/projects/{project_id}/repository/branches/protect`  
#### 参数  
| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| project_id | path | string |  | yes |
| name | body | string |  | yes |
| access_level_string | body | AccessLevelString |  | no |

#### AccessLevelString  
| 参数名 | 类型 | 描述 | 必选 |
| :----: | :---------: | :------: | :------: |
| push_access_levels | string | Access levels allowed to push, split by "," | no |
| merge_access_levels | string | Access levels allowed to merge, split by "," | no |

#### 响应  
```json
{
    "name": "b1",
    "updated_at": "2023-09-19T20:11:51.948+08:00",
    "push_users": [],
    "push_users_access_level_count": null,
    "merge_users": [],
    "merge_users_access_level_count": null,
    "merged": false,
    "protected": true,
    "developers_can_push": true,
    "developers_can_merge": true,
    "committer_can_push": false,
    "committer_can_merge": false,
    "master_can_push": false,
    "master_can_merge": false,
    "maintainer_can_push": false,
    "maintainer_can_merge": false,
    "owner_can_push": false,
    "owner_can_merge": false,
    "no_one_can_push": false,
    "no_one_can_merge": false
}
```
### 6. 获取保护分支详情  
`GET /api/v4/projects/{project_id}/protected_branches/{branch_name}`  
#### 参数  
| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| project_id | path | string | The ID of a project | yes |
| branch_name | path | string | The name of the branch or wildcard | yes |

#### 响应  
```json
{
    "name": "b2",
    "updated_at": "2023-09-19T20:11:51.948+08:00",
    "push_users": [],
    "push_users_access_level_count": null,
    "merge_users": [],
    "merge_users_access_level_count": null,
    "merged": false,
    "protected": true,
    "developers_can_push": false,
    "developers_can_merge": true,
    "committer_can_push": false,
    "committer_can_merge": false,
    "master_can_push": false,
    "master_can_merge": false,
    "maintainer_can_push": false,
    "maintainer_can_merge": false,
    "owner_can_push": true,
    "owner_can_merge": true,
    "no_one_can_push": false,
    "no_one_can_merge": false,
    "push_access_levels": [
        {
            "access_level": 50,
            "access_level_description": null
        }
    ],
    "merge_access_levels": [
        {
            "access_level": 30,
            "access_level_description": "Developers + Committer + Maintainers"
        }
    ],
    "read_access_levels": null,
    "create_change_access_levels": null,
    "create_delete_access_levels": null
}
```
### 7. 删除保护分支  
`DELETE /api/v4/projects/{project_id}/protected_branches/{branch_name}`  
#### 参数  
| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| project_id | path | string | The ID of a project | yes |
| branch_name | path | string | The name of the branch or wildcard | yes |

#### 响应  
```json
{}
```
### 8. 创建保护tag  
`POST /api/v4/projects/{project_id}/protected_tags`  
#### 参数  
| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| project_id | path | string | The ID of a project | yes |
| name | body | string |  | no |
| create_access_level | body | integer |  | no |

#### 响应  
```json
{
    "name": "t1",
    "create_access_levels": [
        {
            "access_level": 30,
            "access_level_description": "Developers + Committer + Maintainers"
        }
    ]
}
```
### 9. 删除保护tag  
`DELETE /api/v4/projects/{project_id}/protected_tags/{name}`  
#### 参数  
| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| project_id | path | string | The ID of a project | yes |
| name | path | string | The name of a tag | yes |

#### 响应  
```json
{}
```
### 10. 获取组织下项目列表  
`GET /api/v4/groups/{id}/projects`  
#### 参数  
| 参数名 | 位置 | 类型 | 描述 | 必选 |
| :----: | :----: | :---------: | :------: | :------: |
| id | path | string | The ID of a group | yes |
| archived | query | string | Limit by archived status | no |
| visibility | query | string | Limit by visibility | no |
| search | query | string | Return list of authorized projects matching the search criteria | no |
| order_by | query | string | Return projects ordered by field | no |
| sort | query | string | Return projects sorted in ascending and descending order | no |
| simple | query | boolean | Return only the ID, URL, name, and path of each project | no |
| owned | query | boolean | Limit by owned by authenticated user | no |
| starred | query | boolean | Limit by starred status | no |
| include_subgroups | query | boolean | Includes projects in subgroups of this group | no |
| page | query | integer | Current page number | no |
| per_page | query | integer | Number of items per page | no |
| with_programming_language | query | string | Limit to repositories which use the given programming language | no |

#### 响应  
```json
[
  {
    "id": 98485,
    "description": null,
    "name": "repocc12",
    "name_with_namespace": "group1111 / repocc12",
    "path": "repocc12",
    "path_with_namespace": "group11111/repocc12",
    "develop_mode": null,
    "created_at": "2023-09-15T16:45:09.502+08:00",
    "updated_at": "2023-09-15T16:45:09.502+08:00",
    "archived": false,
    "ssh_url_to_repo": "ssh://git@gitcode.com/group11111/repocc12.git",
    "http_url_to_repo": "https://gitcode.com/group11111/repocc12.git",
    "web_url": null,
    "readme_url": null,
    "product_id": null,
    "product_name": null
  }
]
```
### 11. 项目列表
`GET /api/v4/projects`
#### 参数
| 参数名 |  位置   | 类型 |                               描述                               |  必选   |
| :----: |:-----:| :---------: |:--------------------------------------------------------------:|:-----:|
| order_by | query | string |                Return projects ordered by field                | false |
| sort | query | string |    Return projects sorted in ascending and descending order    | false |
| archived | query | boolean |                    Limit by archived status                    | false |
| visibility | query | string |                      Limit by visibility                       | false |
| search | query | string |      Return list of projects matching the search criteria      | false |
| owned | query | boolean |              Limit by owned by authenticated user              | false |
| starred | query | boolean |                    Limit by starred status                     | false |
| membership | query | boolean |     Limit by projects that the current user is a member of     | false |
| with_programming_language | query | string | Limit to repositories which use the given programming language | false |
| page | query | integer |                      Current page number                       | false |
| per_page | query | integer |                    Number of items per page                    | false |
| is_forks | query | boolean |        If true,return the fork projects of current_user        | false |
| user_created | query | boolean |    If true,only return the projects created at current_user    | false |
| user_id | query | integer |                 query projects of target user                  | false |
| namespace_id | query | integer |                      The id of namespace                       | false |
| fork_from_project_id | query | integer |                    fork from which project                     | false |
| ids | query | string |                    Limit by project id list                    | false |

#### 响应
```json
[
  {
    "id": 138162,
    "name": "test142",
    "namespace": "f4ManagerTest-update/test142",
    "path": "test142",
    "develop_mode": "normal",
    "tag_list": [],
    "visibility": "public",
    "star_count": 0,
    "forks_count": 0,
    "open_issues_count": 0,
    "open_merge_requests_count": 0,
    "open_change_requests_count": 0,
    "starred": false,
    "name_with_namespace": "f4ManagerTest-update / test142",
    "web_url": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update/test142",
    "last_activity_at": "2023-09-27T15:51:49.033+08:00",
    "main_repository_language": [
      null,
      null
    ],
    "forked_from_project": null,
    "permissions": 40,
    "archived": false,
    "member_count": 1,
    "description": null,
    "repository_size": 0.0,
    "ssh_url_to_repo": "ssh://git@gitcode-backend.cn-north-7.myhuaweicloud.com:2222/f4ManagerTest-update/test142.git",
    "http_url_to_repo": "https://gitcode-backend.cn-north-7.myhuaweicloud.com/f4ManagerTest-update/test142.git",
    "license": null,
    "mirror_project_data": null,
    "namespace_info": {
      "id": 2,
      "name": "f4ManagerTest-update",
      "path": "f4ManagerTest-update",
      "develop_mode": "normal",
      "region": null,
      "cell": null,
      "kind": "user",
      "full_path": "f4ManagerTest-update",
      "full_name": "f4ManagerTest-update",
      "parent_id": null,
      "visibility_level": 20,
      "enable_file_control": false,
      "owner_id": 12
    }
  }
]
```