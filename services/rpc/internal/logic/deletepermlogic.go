package logic

import (
	"context"

	"uam/services/model"
	uampermission "uam/services/model/uam_permission"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeletePermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePermLogic {
	return &DeletePermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除权限
func (l *DeletePermLogic) DeletePerm(in *uamrpc.DeletePermReq) (resp *uamrpc.DeletePermResp, err error) {
	var (
		perm *uampermission.UamPermission
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	perm, err = l.svcCtx.PermissionModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if perm == nil {
		return &uamrpc.DeletePermResp{}, nil
	}
	if perm.Editable == 0 {
		return nil, errors.New(constants.MsgEditableErr)
	}

	if err = l.DeletePermInTx(perm.Id); err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.DeletePermResp{}, nil
}

func (l *DeletePermLogic) DeletePermInTx(permId int64) error {
	return model.ExecInTx(l.svcCtx.DB, func(tx *gorm.DB) error {
		var err error
		if err = l.svcCtx.PermissionModel.Delete(l.ctx, tx, permId); err != nil {
			return err
		}
		// 移除角色-权限关联关系
		if err = l.svcCtx.RelModel.RemoveRolePermByPermId(l.ctx, tx, permId); err != nil {
			return err
		}
		return nil
	})
}
