# UAM

统一访问管理中心--UAM (Unify Access Management Center) 

注意：本项目包含多个服务，包括"UAM-后台管理-后端"、"UAM-API服务"、"UAM-RPC服务"、"UAM-Job任务调度服务"，
"UAM-后台管理-前端"项目在这里

UAM是基于RBCA模型的统一用户权限管理中心，支持任意需要进行访问控制的系统来接入，支持Rest API及RPC接入。

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

