package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (resp *types.UpdateRoleResp, err error) {
	in := &uamrpc.UpdateRoleReq{
		Id:          req.Id,
		ClientId:    req.ClientId,
		Name:        req.Name,
		Desc:        req.Desc,
		Permissions: req.Permissions,
	}
	_, err = l.svcCtx.UamRpc.UpdateRole(l.ctx, in)
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
