package logic

import (
	"context"
	uampermission "uam/services/model/uam_permission"
	"uam/services/rpc/pb/uamrpc"

	"uam/services/model"
	"uam/services/rpc/internal/svc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeletePermByKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePermByKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePermByKeyLogic {
	return &DeletePermByKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  根据key删除权限
func (l *DeletePermByKeyLogic) DeletePermByKey(in *uamrpc.DeletePermByKeyReq) (resp *uamrpc.DeletePermByKeyResp, err error) {
	var (
		perm *uampermission.UamPermission
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	perm, err = l.svcCtx.PermissionModel.FindOneByKey(l.ctx, in.ClientId, in.Key)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if perm == nil {
		return &uamrpc.DeletePermByKeyResp{}, nil
	}

	if err = l.DeletePermInTx(perm.Id); err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.DeletePermByKeyResp{}, nil
}

func (l *DeletePermByKeyLogic) DeletePermInTx(permId int64) error {
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
