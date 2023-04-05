package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoleLogic) AddRole(req *types.AddRoleReq) (resp *types.AddRoleResp, err error) {
	in := &uamrpc.AddRoleReq{
		ClientId:    req.ClientId,
		Name:        req.Name,
		Desc:        req.Desc,
		Permissions: req.Permissions,
	}
	_, err = l.svcCtx.UamRpc.AddRole(l.ctx, in)
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
