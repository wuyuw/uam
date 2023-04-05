package permission

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type DeletePermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePermLogic {
	return &DeletePermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePermLogic) DeletePerm(req *types.DeletePermReq) (resp *types.DeletePermResp, err error) {
	_, err = l.svcCtx.UamRpc.DeletePerm(l.ctx, &uamrpc.DeletePermReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return &types.DeletePermResp{}, nil
}
