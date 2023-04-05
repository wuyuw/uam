package svc

import (
	"gorm.io/gorm"
	"uam/services/job/internal/config"
	"uam/services/job/internal/setup"
	uamgroup "uam/services/model/uam_group"
	uamrel "uam/services/model/uam_rel"
	uamrole "uam/services/model/uam_role"
	"uam/services/model/user"
)

type ServiceContext struct {
	Config config.Config

	DB         *gorm.DB
	UserModel  *user.UserModel
	RoleModel  *uamrole.RoleModel
	GroupModel *uamgroup.GroupModel
	RelModel   *uamrel.RelModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := setup.SetupGormMysql(c)
	return &ServiceContext{
		Config:     c,
		DB:         db,
		UserModel:  user.NewUserModel(db),
		RoleModel:  uamrole.NewRoleModel(db),
		GroupModel: uamgroup.NewGroupModel(db),
		RelModel:   uamrel.NewRelModel(db),
	}
}
