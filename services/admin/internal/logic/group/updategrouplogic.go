package group

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type UpdateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGroupLogic) UpdateGroup(req *types.UpdateGroupReq) (resp *types.UpdateGroupResp, err error) {
	in := &uamrpc.UpdateGroupReq{
		Id:       req.Id,
		ClientId: req.ClientId,
		Name:     req.Name,
		Desc:     req.Desc,
		Roles:    req.Roles,
	}
	_, err = l.svcCtx.UamRpc.UpdateGroup(l.ctx, in)
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
