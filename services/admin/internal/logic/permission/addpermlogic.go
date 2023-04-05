package permission

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type AddPermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPermLogic {
	return &AddPermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPermLogic) AddPerm(req *types.AddPermReq) (resp *types.AddPermResp, err error) {
	in := &uamrpc.AddPermReq{
		ClientId: req.ClientId,
		Type:     req.Type,
		Key:      req.Key,
		Name:     req.Name,
		Desc:     req.Desc,
	}
	_, err = l.svcCtx.UamRpc.AddPerm(l.ctx, in)
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return &types.AddPermResp{}, nil
}
