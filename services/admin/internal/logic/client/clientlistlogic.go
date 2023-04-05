package client

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type ClientListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientListLogic {
	return &ClientListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientListLogic) ClientList(req *types.ClientListReq) (resp *types.ClientListResp, err error) {
	rpcResp, err := l.svcCtx.UamRpc.GetClientList(l.ctx, &uamrpc.GetClientListReq{})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	clients := make([]types.ClientItem, len(rpcResp.List))
	for i, item := range rpcResp.List {
		clients[i] = types.ClientItem{
			Id:         item.Id,
			Name:       item.Name,
			AppCode:    item.AppCode,
			PrivateKey: item.PrivateKey,
			Department: item.Department,
			Maintainer: item.Maintainer,
			Status:     item.Status,
			Type:       item.Type,
			CreateTime: item.CreateTime,
			UpdateTime: item.UpdateTime,
		}
	}
	return &types.ClientListResp{
		List: clients,
	}, nil
}
