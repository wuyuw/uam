# UAM

统一访问管理中心--UAM (Unify Access Management Center)

UAM是基于RBCA模型的统一用户权限管理中心，支持任意需要进行访问控制的系统来接入，支持Rest API及RPC接入。

本项目包含多个服务，包括"UAM-后台管理-后端"、"UAM-API服务"、"UAM-RPC服务"、"UAM-Job任务调度服务"，
[UAM-后台管理-前端](https://github.com/wuyuw/uam-admin-fe)项目在这里

技术栈：`go-zero`、`gorm`、`Mysql`、`Redis`、`Kafka`、`etcd`

### 系统架构

<img title="" src="https://github.com/wuyuw/uam/blob/master/images/uam-framework.png?raw=true" alt="UAM架构" data-align="inline">

### 架构说明

系统提供了一个后台管理平台，UAM管理员和各应用管理员可登录后台对各种资源实体进行CRUD操作。

应用接入后会获得`app_id`和`app_secret`，按照接入规范即可访问`Rest API服务`获取权限相关资源，从而实现访问控制。

`UAM-Admin`及`Rest API`均通过调用`RPC Service`实现对数据层的操作，通过`etcd`实现服务注册、发现。

对于一些异步操作，通过基于`Kafka`实现的异步任务队列在`Task Worker`中异步执行，另外定时任务类需求可通过`Cron Job`来配置执行。

目前`UAM-Admin`提供了一套简单的登录注册模块来实现用户添加，后续可同步企业内部用户表替换现有用户表，可接入OAuth认证方式替换账号密码登录。

当前登录用户是通过JWT实现认证，通过redis缓存实现JWT Token续期，避免Token过期造成用户体验问题。


### 应用接入流程
假设应用App01需要接入UAM系统，`App01管理员`需要向`UAM系统管理员`提交接入申请，

申请通过后`UAM系统管理员`在UAM后台添加客户端接入记录，并将系统生成的`app_id`和`app_secret`发给`App01管理员`，
并授予`App01管理员`在UAM后台操作`App01`资源的权限。

后续`App01管理员`也可登录UAM后台对`App01`系统下的资源实体进行CURD操作和用户访问权限管理，

`App01开发人员`通过`app_id`和`app_secret`即可访问UAM开放的`Rest API服务`获取`App01`下的资源数据和用户权限数据，
以实现访问控制管理。


### 资源实体关系

<img title="" src="https://github.com/wuyuw/uam/blob/master/images/uam-resoures.png?raw=true" alt="UAM资源实体" data-align="inline">


## 1 开发环境搭建

### 1.1 环境依赖
自行搭建以下组件

- Mysql

- Redis

- etcd

- Kafka

### 1.2 数据库准备

1. 创建`uam`数据库

2. 创建所需表：deploy/sql/*.sql


### 1.3 启动job服务

```bash
# 更新services/job/etc/uam-job.yaml配置文件
cd services/job
go run job.go
```

### 1.4 启动rpc服务

```bash
# 更新services/rpc/etc/uamrpc.yaml配置文件
cd services/rpc
go run uamrpc.go
```

### 1.4 启动Rest API服务

```bash
# 更新services/api/etc/uam-api.yaml配置文件
cd services/api
go run uam.go
```

### 1.5 启动Admin API服务

```bash
# 更新services/admin/etc/uam-admin-api.yaml配置文件
cd services/admin
go run uam-admin.go
```

## 2 定制开发

### admin后端

#### 新增接口

1. 进入admin目录
   
   ```bash
   cd services/admin
   ```

2. 创建新接口依赖的类型api文件
   
   ```bash
   mkdir desc/role
   vim desc/role/role.api
   ```

3. 定义接口依赖的类型
   
   ```bash
   syntax = "v1"
   
   info(
       title: "角色管理"
       desc: "角色管理"
       author: "will515"
       email: "wuyuw515@gmail.com"
       version: 1.0
   )
   
   // 获取所有角色
   type (
       Role {
           Id          int64    `json:"id"`
           ClientId    int64    `json:"client_id"`
           Name        string   `json:"name"`
           Desc        string   `json:"desc"`
           Editable    int64    `json:"editable"`
           CreateTime  string   `json:"create_time"`
           UpdateTime  string   `json:"update_time"`
           Permissions []string `json:"permissions"`
       }
   
       RoleListReq {
           ClientId int64  `form:"client_id"`
           Editable string `form:"editable,optional"`
       }
   
       RoleListResp {
           List []Role `json:"list"`
       }
   )
   ```

4. 在desc/admin.api文件中导入刚创建的类型文件，并定义接口
   
   ```bash
   import (
       # "core/core.api"
       # "user/user.api"
       # "client/client.api"
       # "permission/permission.api"
       # 导入类型文件
       "role/role.api"
       # "group/group.api"
   )
   
   # 定义接口
   @server(
       prefix: "/uam/admin/v1"
       group: role
       middleware: JwtAuth, AccessControl
   )
   service uam-admin-api {
       @doc "获取角色列表"
       @handler RoleList
       get /roles (RoleListReq) returns (RoleListResp)
   }
   ```

5. 使用命令行工具goctl生成代码模板
   
   ```bash
   # 在uam/services/admin目录下执行
   goctl api go -api desc/admin.api --dir .
   ```

6. 修改生成的handler和logic目录下对应的包





### 定义数据表模型

1. 创建表对应的包
   
   ```bash
   mkdir model/user
   ```
   
   

2. 分别创建model/user/gorm.go和model/user/model.go两个文件，定义orm结构体和封装数据库操作的model
   
   gorm.go
   
   ```go
   package user
   
   import "time"
   
   const TableUser = "user"
   
   type User struct {
   	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
   	Uid        int64     `gorm:"column:uid"`      // uid
   	Nickname   string    `gorm:"column:nickname"` // 花名
   	Email      string    `gorm:"column:email"`    // 邮箱
   	Phone      string    `gorm:"column:phone"`    // 手机
   	CreateTime time.Time `gorm:"column:create_time"`
   	UpdateTime time.Time `gorm:"column:update_time"`
   }
   
   func (User) TableName() string {
   	return TableUser
   }
   
   ```
   
   model.go
   
   ```go
   package user
   
   import (
   	"context"
   	"fmt"
   	"uam/services/model"
   
   	"gorm.io/gorm"
   )
   
   type UserModel struct {
   	table string
   	db    *gorm.DB
   }
   
   func NewUserModel(db *gorm.DB) *UserModel {
   	return &UserModel{
   		db:    db,
   		table: TableUser,
   	}
   }
   
   // FindOneByUid 根据uid查询用户信息
   func (m *UserModel) FindOneByUid(ctx context.Context, uid int64) (*User, error) {
   	var (
   		err  error
   		user User
   	)
   	db := m.db.Table(m.table)
   	err = db.Where("`uid` = ?", uid).First(&user).Error
   	switch err {
   	case nil:
   		return &user, nil
   	case gorm.ErrRecordNotFound:
   		return nil, model.ErrNotFound
   	default:
   		return nil, err
   	}
   }
   ```





### RPC接口开发

1. 更新proto文件，定义类型和接口
   
   services/rpc/uamprc.proto
   
   ```protobuf
   // 获取用户信息
   message GetUserInfoReq {
     int64  uid = 1;
   }
   message GetUserInfoResp {
     int64 uid = 1;
     string nickname = 2;
     string email = 3;
     string phone = 4;
   }
   
   
   
   // services
   service Uam {
     // 定义接口
     rpc getUserInfo(GetUserInfoReq) returns(GetUserInfoResp);
   }
   ```
   
   

2. 生成rpc服务代码
   
   ```bash
   goctl rpc protoc uamrpc.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
   ```

3. 添加RPC接口业务逻辑

## 3 基于 Github Actions 的 CI 执行自动构建镜像

1. Fork 代码仓库

2. 在仓库导航栏的`Settings` -> `Secrets and variables` -> `Actions` -> `New repository secret` 添加 Docker Hub 个人账户的账户名和密码

   账户变量名: `DOCKERHUB_TOKEN` 
   
   密码变量名: `DOCKERHUB_USERNAME`

3. 拉取代码

   ```bash
   $ git clone git@github.com:${usename}/uam.git
   ```

4. 打 tag，CI会根据tag的前缀构建对应服务的docker镜像并上传到Docker Hub

   ```bash
   $ git tag admin-0.0.1
   $ git tag rpc-0.0.1
   $ git tag job-0.0.1
   $ git tag api-0.0.1
   $ git push origin admin-0.0.1
   ```

5. 点击仓库导航栏的`Actions`选项卡查看 Workflow 执行情况

6. 在 Docker Hub 仓库中确认镜像是否上传成功