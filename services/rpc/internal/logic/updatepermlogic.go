package logic

import (
	"context"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermLogic {
	return &UpdatePermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新权限
func (l *UpdatePermLogic) UpdatePerm(in *uamrpc.UpdatePermReq) (resp *uamrpc.UpdatePermResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	perm, err := l.svcCtx.PermissionModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if perm.Key != in.Key {
		return nil, errors.New("update permission error: can not update field `key`")
	}
	perm.Type = in.Type
	perm.Name = in.Name
	perm.Desc = in.Desc
	if err := l.svcCtx.PermissionModel.Update(l.ctx, perm); err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.UpdatePermResp{}, nil
}
