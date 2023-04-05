package user

import (
	"context"
	"uam/services/rpc/pb/uamrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/tools/errx"
)

type UpdateUserPermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPermLogic {
	return &UpdateUserPermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserPermLogic) UpdateUserPerm(req *types.UpdateUserPermReq) (resp *types.UpdateUserPermResp, err error) {
	_, err = l.svcCtx.UamRpc.UpdateUserPerm(l.ctx, &uamrpc.UpdateUserPermReq{
		Uid:      req.Uid,
		ClientId: req.ClientId,
		Groups:   req.Groups,
		Roles:    req.Roles,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
