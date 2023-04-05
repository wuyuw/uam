package group

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type DeleteGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGroupLogic) DeleteGroup(req *types.DeleteGroupReq) (resp *types.DeleteGroupResp, err error) {
	_, err = l.svcCtx.UamRpc.DeleteGroup(l.ctx, &uamrpc.DeleteGroupReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
