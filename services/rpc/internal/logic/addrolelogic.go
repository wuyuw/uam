package logic

import (
	"context"
	"fmt"

	"uam/services/model"
	uampermission "uam/services/model/uam_permission"
	uamrel "uam/services/model/uam_rel"
	uamrole "uam/services/model/uam_role"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type AddRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRoleLogic) AddRole(in *uamrpc.AddRoleReq) (resp *uamrpc.AddRoleResp, err error) {
	var (
		role *uamrole.UamRole
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	ok, err := l.hasDupRoleName(in.ClientId, in.Name)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if ok {
		return nil, errors.New(fmt.Sprintf("角色已存在: clientId: %d role: %s", in.ClientId, in.Name))
	}

	// 有效的权限
	validPerms, err := l.svcCtx.PermissionModel.FindByKeys(l.ctx, in.ClientId, in.Permissions)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}

	role = &uamrole.UamRole{
		ClientId: in.ClientId,
		Name:     in.Name,
		Desc:     in.Desc,
		Editable: 1,
	}

	if err = l.addRoleInTx(role, validPerms); err != nil {
		return nil, err
	}

	return &uamrpc.AddRoleResp{}, nil
}

// 是否存在重复的角色名
func (l *AddRoleLogic) hasDupRoleName(clientId int64, roleName string) (bool, error) {
	dupRole, err := l.svcCtx.RoleModel.FindOneByName(l.ctx, clientId, roleName)
	if err != nil && err != model.ErrNotFound {
		return false, err
	}
	if dupRole != nil {
		return true, nil
	}
	return false, nil
}

func (l *AddRoleLogic) addRoleInTx(role *uamrole.UamRole, permList []*uampermission.UamPermission) error {
	return model.ExecInTx(l.svcCtx.DB, func(tx *gorm.DB) error {
		// 添加角色
		if err := l.svcCtx.RoleModel.InsertOne(l.ctx, tx, role); err != nil {
			return err
		}
		if len(permList) > 0 {
			// 关联权限
			relRolePermRecords := make([]uamrel.RelRolePermission, len(permList))
			for i, perm := range permList {
				relRolePermRecords[i] = uamrel.RelRolePermission{
					RoleId:       role.Id,
					PermissionId: perm.Id,
				}
			}
			if err := tx.Create(&relRolePermRecords).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
