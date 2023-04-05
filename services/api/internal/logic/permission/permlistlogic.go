package permission

import (
	"context"
	"errors"
	"uam/services/rpc/pb/uamrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/api/internal/consts"
	"uam/services/api/internal/svc"
	"uam/services/api/internal/types"
	"uam/tools/contextx"
	"uam/tools/errx"
)

type PermListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermListLogic {
	return &PermListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermListLogic) PermList(req *types.PermListReq) (resp *types.PermListResp, err error) {
	// 从ctx中获取clientId
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	rpcResp, err := l.svcCtx.UamRpc.GetPermList(l.ctx, &uamrpc.GetPermListReq{
		ClientId: clientId,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	permList := make([]types.PermItem, len(rpcResp.List))
	for i, item := range rpcResp.List {
		permList[i] = types.PermItem{
			Id:         item.Id,
			Type:       item.Type,
			Key:        item.Key,
			Name:       item.Name,
			Desc:       item.Desc,
			CreateTime: item.CreateTime,
			UpdateTime: item.UpdateTime,
		}
	}
	return &types.PermListResp{
		List: permList,
	}, nil
}
