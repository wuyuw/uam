package group

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type AddGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupLogic {
	return &AddGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGroupLogic) AddGroup(req *types.AddGroupReq) (resp *types.AddGroupResp, err error) {
	in := &uamrpc.AddGroupReq{
		ClientId: req.ClientId,
		Name:     req.Name,
		Desc:     req.Desc,
		Roles:    req.Roles,
	}
	_, err = l.svcCtx.UamRpc.AddGroup(l.ctx, in)
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
