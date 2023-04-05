package svc

import (
	uamauth "uam/services/model/auth"
	uamclient "uam/services/model/uam_client"
	uamgroup "uam/services/model/uam_group"
	uampermission "uam/services/model/uam_permission"
	uamrel "uam/services/model/uam_rel"
	uamrole "uam/services/model/uam_role"
	uamuser "uam/services/model/user"
	"uam/services/rpc/internal/config"
	"uam/services/rpc/internal/setup"

	"github.com/zeromicro/go-queue/kq"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	DB              *gorm.DB
	AuthLocalModel  *uamauth.AuthLocalModel
	UserModel       *uamuser.UserModel
	ClientModel     *uamclient.ClientModel
	PermissionModel *uampermission.PermissionModel
	RoleModel       *uamrole.RoleModel
	GroupModel      *uamgroup.GroupModel
	RelModel        *uamrel.RelModel

	RelUpdatePusher *kq.Pusher // kafka生产者
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := setup.SetupGormMysql(c)
	setup.InitIdGenerator(c.IdGen.WorkerId)
	return &ServiceContext{
		Config: c,

		DB:              db,
		AuthLocalModel:  uamauth.NewAuthLocalModel(db),
		UserModel:       uamuser.NewUserModel(db),
		ClientModel:     uamclient.NewClientModel(db),
		PermissionModel: uampermission.NewPermissionModel(db),
		RoleModel:       uamrole.NewRoleModel(db),
		GroupModel:      uamgroup.NewGroupModel(db),
		RelModel:        uamrel.NewRelModel(db),
		RelUpdatePusher: kq.NewPusher(c.RelUpdateMq.Brokers, c.RelUpdateMq.Topic),
	}
}
