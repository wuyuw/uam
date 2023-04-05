package logic

import (
	"context"
	"fmt"

	"uam/services/model"
	uamrole "uam/services/model/uam_role"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *uamrpc.UpdateRoleReq) (resp *uamrpc.UpdateRoleResp, err error) {
	var (
		role             *uamrole.UamRole
		delPermissionIds []int64 // 待删除权限列表
		addPermissionIds []int64 // 待添加权限列表
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	role, err = l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	// 查重
	if role.Name != in.Name {
		ok, err := l.hasDupRoleName(in.ClientId, in.Name)
		if err != nil {
			return nil, errors.Wrap(err, constants.MsgDBErr)
		}
		if ok {
			return nil, errors.New(fmt.Sprintf("角色已存在: %s", in.Name))
		}
	}
	role.Name = in.Name
	role.Desc = in.Desc

	delPermissionIds, addPermissionIds, err = l.getDelAndAddPermIds(role.Id, role.ClientId, in.Permissions)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}

	if err = l.updateRoleInTx(role, delPermissionIds, addPermissionIds); err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}

	return &uamrpc.UpdateRoleResp{}, nil
}

func (l *UpdateRoleLogic) hasDupRoleName(clientId int64, roleName string) (bool, error) {
	dupRole, err := l.svcCtx.RoleModel.FindOneByName(l.ctx, clientId, roleName)
	if err != nil && err != model.ErrNotFound {
		return false, err
	}
	if dupRole != nil {
		return true, nil
	}
	return false, nil
}

func (l *UpdateRoleLogic) getDelAndAddPermIds(roleId, clientId int64, newPermKeys []string) (delIds []int64, addIds []int64, err error) {
	var (
		oldPermissionIds []int64 // 旧权限列表
		delPermissionIds []int64 // 待删除权限列表
		newPermissionIds []int64 // 待添加权限列表
	)
	// 旧的权限列表
	oldPermissionIds, err = l.svcCtx.RelModel.FindPermIdsByRoleIds(l.ctx, []int64{roleId})
	if err != nil {
		return nil, nil, err
	}

	// 新的权限列表
	newPermissions, err := l.svcCtx.PermissionModel.FindByKeys(l.ctx, clientId, newPermKeys)
	if err != nil {
		return nil, nil, err
	}

	newPermMap := make(map[int64]bool)
	for _, perm := range newPermissions {
		newPermMap[perm.Id] = true
		newPermissionIds = append(newPermissionIds, perm.Id)
	}
	for _, permId := range oldPermissionIds {
		if !newPermMap[permId] {
			delPermissionIds = append(delPermissionIds, permId)
		}
	}
	return delPermissionIds, newPermissionIds, nil
}

func (l *UpdateRoleLogic) updateRoleInTx(role *uamrole.UamRole, delPermissionIds, addPermissionIds []int64) (err error) {
	return model.ExecInTx(l.svcCtx.DB, func(tx *gorm.DB) error {
		var err error
		if err = tx.Updates(role).Error; err != nil {
			return err
		}
		if err = l.svcCtx.RelModel.RemoveRolePermByPermIds(l.ctx, tx, role.Id, delPermissionIds); err != nil {
			return err
		}
		if err = l.svcCtx.RelModel.AddRolePermByPermIds(l.ctx, tx, role.Id, addPermissionIds); err != nil {
			return err
		}
		return nil
	})
}
