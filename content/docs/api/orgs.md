---
linkTitle: 组织 API
title: 组织接口API文档
weight: 2
sidebar:
  open: false
---

## 1. 获取组织详情

### 请求

`GET /api/v4/groups/{group_id}`

### 参数

| 参数名       | 类型   | in|是否必选|default|描述                 |
| ------------ | ------ | ----|---|-------------------- | ------------ |
| group_id     | string |path|true|| 组织ID或者组织路径   |
| with_projects | bool | query |false|true|为true时，返回体携带项目MR,ISSUES数量相关信息 |
| view | enum | query |false||- all<br />- simple<br />- sum_open<br />为“all”或者null是，返回默认结构体<br />为"simple"返回最简结构体<br />为"sum_open" （with_projects为true时生效)，MR,ISSUES只返回开启中的数量|

### 响应

```json
//默认结构体
{
    "created_at": "2023-08-30T17:43:50.147+08:00",
    "description": "1233333333333",
    "full_name": "public_083111",
    "full_path": "public_0831111",
    "id": 12312,
    "name": "public_083111",
    "chinese_name": "0831test1",
    "parent_id": null,
    "path": "public_0831111",
    "web_url": "https://demoweb.com/public_0831111",
    "develop_mode": "normal",
    "region": null,
    "cell": "default",
    "email": "example@demo.com",
    "avatar": "www.111111.com",
    "home_page": "www.homepage1.com",
    "can_create_private_project": false,
    "children": null,
    "department": {
        "id": 57120,
        "target_id": 12312,
        "target_type": "Group",
        "department2": null,
        "department3": null,
        "department4": null,
        "department5": null,
        "department6": null,
        "department7": null,
        "created_at": "2023-08-30T17:43:50.166+08:00",
        "updated_at": "2023-08-30T17:43:50.166+08:00"
    },
    "department_cn": {
        "id": 57120,
        "target_id": 12312,
        "target_type": "Group",
        "department2": "",
        "department3": "",
        "department4": "",
        "department5": "",
        "department6": "",
        "department7": "",
        "created_at": "2023-08-30T17:43:50.166+08:00",
        "updated_at": "2023-08-30T17:43:50.166+08:00"
    },
    "disallowed_visibility": null,
    "group_extend": {
        "id": 12202,
        "group_id": 12312,
        "disable_fork": null,
        "is_official": true,
        "cmo_name": null,
        "email": "example@demo.com",
        "home_page": "www.homepage1.com",
        "join_way": 0,
        "can_create_private_project": false,
        "enable_file_control": false
    },
    "item_type": "Group",
    "last_owner": true,
    "lfs_enabled": false,
    "members": 1,
    "my_role": {
        "id": 15716,
        "access_level": 50,
        "source_id": 12312,
        "source_type": "Namespace",
        "user_id": 12,
        "notification_level": 3,
        "created_at": "2023-08-30T17:43:50.152+08:00",
        "updated_at": "2023-08-30T17:43:50.152+08:00",
        "created_by_id": null,
        "invite_email": null,
        "invite_token": null,
        "invite_accepted_at": null,
        "requested_at": null,
        "expires_at": null,
        "limited": false
    },
    "project_count": 0,
    "request_access_enabled": true,
    "share_with_group_lock": false,
    "star_count": 1,
    "starred": true,
    "sub_group_count": 0,
    "visibility": "public",
    // with_projects 为true时携带以下数据
    "sum": {
        // view == sum_open 时,以下两项数据为null
        "all_issues_count": 1,
        "all_merge_requests_count": 0,
        
        "open_issues_count": 1,
        "open_merge_requests_count": 0,
        "open_change_requests_count": null
    }
}
//简单结构体
{
    "created_at": "2023-08-30T17:43:50.147+08:00",
    "description": "1233333333333",
    "full_name": "public_083111",
    "full_path": "public_0831111",
    "id": 12312,
    "name": "public_083111",
    "chinese_name": "0831test1",
    "parent_id": null,
    "path": "public_0831111",
    "web_url": "https://demoweb.com/public_0831111",
    "develop_mode": "normal",
    "region": null,
    "cell": "default",
    "email": "example@demo.com",
    "avatar": "www.111111.com",
    "home_page": "www.homepage1.com"
}
```

## 2. 创建组织

### 请求

`POST /api/v4/groups`

### 参数
#### request body
```json
{
  "name": "string", //组织名称
  "chinese_name": "string", //组织中文名称
  "path": "string", //组织自身路径
  "avatar": "string",//组织头像url,目前只支持华为云obs
  "region": "string",//组织数据所属地域,待cell化上线后支持
  "cell": "string",//组织数据所属cell,待cell化上线后支持
  "parent_id": 0,//父组织id,待多层组织功能上线后支持
  "visibility": "private",//组织可见等级 private|public
  "description": "string",//组织描述
  "lfs_enabled": true,//组织是否支持lfs
}
```
### 响应

```json
{
    "id": 81878,
    "web_url": "https://demoweb.com/0922_wjh",
    "name": "0922_wjh",
    "chinese_name": "0831test1",
    "path": "0922_wjh",
    "develop_mode": "normal",
    "region": null,
    "cell": "default",
    "description": "1233333333333",
    "visibility": "public",
    "lfs_enabled": null,
    "avatar": "www.111111.com",
    "email": null,
    "home_page": null,
    "request_access_enabled": "false",
    "full_name": "0922_wjh",
    "full_path": "0922_wjh",
    "parent_id": null,
    "projects": [],
    "shared_projects": [],
    "can_create_private_project": "false"
}
```
## 3. 获取组织列表

### 请求

`get /api/v4/groups/list`

### 参数
| 参数名       | 类型   | in|是否必选|default|描述                 |
| ------------ | ------ | ----|---|-------------------- | ------------ |
| parent_id | integer |query|false|| 限制得到的组织父组织为parent_id(待支持多层组织后生效) |
| skip_groups | List<Integer> | query |false||需要排除的组织 |
| search | string | query |false|| 模糊查询name或者path中带有search指定字符的组织(大小写不敏感) |
| owned | bool | query |false|| 是否只查询用户有管理员角色的组织                             |
| min_access_level | enum | query |false|| - 10<br /> - 30<br /> - 50<br />查询用户拥有权限>=min_access_level的组织 |
| starred | boolean | query |false|| 查询用户关注的组织                                           |
| maintained | boolean | query |false|| 查询用户管理的的组织(与owned的区别在,会查询用户有管理员角色的组织以及其子组织) |
| all_available | boolean | query |false|true| 为true时，会查询所有公开组织                                 |
| order_by | enum | query |false|created_at| - name<br /> - path<br /> - id<br /> - created_at<br /> - updated_at |
| sort | enum | query |false|desc| - asc <br />- desc                                           |
| page | integer | query |false|1| 分页参数，查询第page页                                       |
| per_page | integer | query |false|20| 指定每页查询的个数                                           |
| simple | boolean | query |false|false| 是否只返回最简结构体                                         |
特别说明:

1. owned,starred,maintained,min_access_level 指定的逻辑，只能生效其中一个，优先级按照此处给的顺序从大到小。
2. 不传search且不是查询关注的组织时，只会查询根组织
3. 只有在不传owned,starred,maintained,min_access_level其中任意一项时，all_available参数生效，当其为false时，会查询用户有权限的组织，当不传或者为true时，会额外带上所有合法的公开组织。
4. 当不传入token时，owned,starred,maintained,min_access_level全部失效，查询结果为合法的公开组织，其他参数仍然可以指定

### 响应

响应头中会带有分页相关信息
#### response body

```json
[
    {
        "id": 81878,
        "name": "0922_wjh",
        "chinese_name": "0831test1",
        "web_url": "https://demoweb.com/0922_wjh",
        "lfs_enabled": true,
        "full_name": "0922_wjh",
        "full_path": "0922_wjh",
        "path": "0922_wjh",
        "visibility": "public",
        "description": "1233333333333",
        "request_access_enabled": false,
        "share_with_group_lock": false,
        "avatar": "www.111111.com",
        "item_type": "Group",
        "parent_id": null,
        "my_role": null,
        "member_count": 1,
        "department": {
            "id": 216217,
            "target_id": 81878,
            "target_type": "Group",
            "department2": null,
            "department3": null,
            "department4": null,
            "department5": null,
            "department6": null,
            "department7": null,
            "created_at": "2023-09-22T11:45:28.943+08:00",
            "updated_at": "2023-09-22T11:45:28.943+08:00"
        },
        "created_at": "2023-09-22T11:45:28.555+08:00",
        "project_count": 0,
        "group_extend": null,
        "sub_group_count": 0,
        "last_owner": null,
        "can_create_private_project": false,
        "starred": null,
        "develop_mode": "normal",
        "region": null,
        "cell": "default",
        "moderation_result": true
    },
    {
        "id": 81844,
        "name": "API_test__name_eEnrt8KfT7_level_02",
        "chinese_name": "",
        "web_url": "https://demoweb.com/API_test__path_ifWVvBR121/API_test__path_ifWVvBR121_level_02",
        "lfs_enabled": true,
        "full_name": "API_test__name_eEnrt8KfT7 / API_test__name_eEnrt8KfT7_level_02",
        "full_path": "API_test__path_ifWVvBR121/API_test__path_ifWVvBR121_level_02",
        "path": "API_test__path_ifWVvBR121_level_02",
        "visibility": "public",
        "description": "API_test__name_eEnrt8KfT7_level_02",
        "request_access_enabled": false,
        "share_with_group_lock": false,
        "avatar": "",
        "item_type": "Group",
        "parent_id": 81843,
        "my_role": null,
        "member_count": 1,
        "department": {
            "id": 215946,
            "target_id": 81844,
            "target_type": "Group",
            "department2": null,
            "department3": null,
            "department4": null,
            "department5": null,
            "department6": null,
            "department7": null,
            "created_at": "2023-09-21T15:46:16.035+08:00",
            "updated_at": "2023-09-21T15:46:16.035+08:00"
        },
        "created_at": "2023-09-21T15:46:16.023+08:00",
        "project_count": 0,
        "group_extend": null,
        "sub_group_count": 0,
        "last_owner": null,
        "can_create_private_project": false,
        "starred": null,
        "develop_mode": "normal",
        "region": null,
        "cell": "default",
        "moderation_result": true
    }]
```