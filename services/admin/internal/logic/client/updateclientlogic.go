package client

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type UpdateClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClientLogic {
	return &UpdateClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateClientLogic) UpdateClient(req *types.UpdateClientReq) (resp *types.UpdateClientResp, err error) {
	in := &uamrpc.UpdateClientReq{
		Id:         req.Id,
		Name:       req.Name,
		AppCode:    req.AppCode,
		Department: req.Department,
		Maintainer: req.Maintainer,
		Status:     req.Status,
	}
	_, err = l.svcCtx.UamRpc.UpdateClient(l.ctx, in)
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return &types.UpdateClientResp{}, nil
}
