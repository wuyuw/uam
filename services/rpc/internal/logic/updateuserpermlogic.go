package logic

import (
	"context"
	"uam/services/model"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/collections"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateUserPermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPermLogic {
	return &UpdateUserPermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserPermLogic) UpdateUserPerm(in *uamrpc.UpdateUserPermReq) (resp *uamrpc.UpdateUserPermResp, err error) {
	var (
		delGroupIds []int64
		delRoleIds  []int64
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	delRoleIds, err = l.getDelRoleIds(in.Uid, in.ClientId, in.Roles)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	delGroupIds, err = l.getDelGroupIds(in.Uid, in.ClientId, in.Groups)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}

	if err = l.updateUserPermInTx(in.Uid, in.ClientId, in.Groups, delGroupIds, in.Roles, delRoleIds); err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.UpdateUserPermResp{}, nil
}

func (l *UpdateUserPermLogic) getDelRoleIds(uid, clientId int64, newRoleIds []int64) (delRoleIds []int64, err error) {
	oldRoleIds, err := l.svcCtx.RelModel.FindRoleIdsByUid(l.ctx, clientId, uid)
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

func (l *UpdateUserPermLogic) getDelGroupIds(uid, clientId int64, newGroupIds []int64) (delGroupIds []int64, err error) {
	oldGroupIds, err := l.svcCtx.RelModel.FindGroupIdsByUid(l.ctx, clientId, uid)
	if err != nil {
		return nil, err
	}
	if len(oldGroupIds) == 0 {
		return nil, nil
	}
	if len(newGroupIds) == 0 {
		return oldGroupIds, nil
	}
	delGroupIds = collections.SliceIntToInt64(
		collections.Difference(
			collections.SliceInt64ToInt(oldGroupIds),
			collections.SliceInt64ToInt(newGroupIds)))
	return delGroupIds, nil
}

func (l *UpdateUserPermLogic) updateUserPermInTx(uid, clientId int64, newGroupIds, delGroupIds, newRoleIds, delRoleIds []int64) error {
	return model.ExecInTx(l.svcCtx.DB, func(tx *gorm.DB) error {
		var err error
		if err = l.svcCtx.RelModel.RemoveUserGroupByGroupIds(l.ctx, tx, clientId, uid, delGroupIds); err != nil {
			return err
		}
		if err = l.svcCtx.RelModel.AddUserGroupByGroupIds(l.ctx, tx, clientId, uid, newGroupIds); err != nil {
			return err
		}
		if err = l.svcCtx.RelModel.RemoveUserRoleByRoleIds(l.ctx, tx, clientId, uid, delRoleIds); err != nil {
			return err
		}
		if err = l.svcCtx.RelModel.AddUserRoleByRoleIds(l.ctx, tx, clientId, uid, newRoleIds); err != nil {
			return err
		}
		return nil
	})
}
