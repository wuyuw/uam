package logic

import (
	"context"
	"fmt"

	"uam/services/model"
	uamgroup "uam/services/model/uam_group"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/collections"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新组
func (l *UpdateGroupLogic) UpdateGroup(in *uamrpc.UpdateGroupReq) (resp *uamrpc.UpdateGroupResp, err error) {
	var (
		group      *uamgroup.UamGroup
		delRoleIds []int64 // 待删除角色列表
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	group, err = l.svcCtx.GroupModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	// 查重
	if group.Name != in.Name {
		ok, err := l.hasDupGroupName(in.ClientId, in.Name)
		if err != nil {
			return nil, errors.Wrap(err, constants.MsgDBErr)
		}
		if ok {
			return nil, errors.New(fmt.Sprintf("组已存在: %s", in.Name))
		}
	}
	group.Name = in.Name
	group.Desc = in.Desc

	delRoleIds, err = l.getDelRoleIds(group.Id, group.ClientId, in.Roles)
	if err != nil {
		return nil, err
	}

	if err = l.updateGroupInTx(group, delRoleIds, in.Roles); err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.UpdateGroupResp{}, nil
}

func (l *UpdateGroupLogic) hasDupGroupName(clientId int64, groupName string) (bool, error) {
	dupRecord, err := l.svcCtx.GroupModel.FindOneByName(l.ctx, clientId, groupName)
	if err != nil && err != model.ErrNotFound {
		return false, err
	}
	if dupRecord != nil {
		return true, nil
	}
	return false, nil
}

func (l *UpdateGroupLogic) getDelRoleIds(groupId, clientId int64, newRoleIds []int64) (delIds []int64, err error) {
	var (
		oldRoleIds []int64 // 旧角色列表
		delRoleIds []int64 // 待移除角色列表
	)
	oldRoleIds, err = l.svcCtx.RelModel.FindRoleIdsByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, err
	}
	if len(oldRoleIds) == 0 {
		return nil, nil
	}
	if len(newRoleIds) == 0 {
		return oldRoleIds, nil
	}

	delRoleIds = collections.SliceIntToInt64(
		collections.Difference(
			collections.SliceInt64ToInt(oldRoleIds),
			collections.SliceInt64ToInt(newRoleIds)))

	return delRoleIds, nil
}

func (l *UpdateGroupLogic) updateGroupInTx(group *uamgroup.UamGroup, delRoleIds, addRoleIds []int64) (err error) {
	return model.ExecInTx(l.svcCtx.DB, func(tx *gorm.DB) error {
		var err error
		if err = tx.Updates(group).Error; err != nil {
			return err
		}
		if err = l.svcCtx.RelModel.RemoveGroupRoleByRoleIds(l.ctx, tx, group.Id, delRoleIds); err != nil {
			return err
		}
		if err = l.svcCtx.RelModel.AddGroupRoleByRoleIds(l.ctx, tx, group.Id, addRoleIds); err != nil {
			return err
		}
		return nil
	})
}
