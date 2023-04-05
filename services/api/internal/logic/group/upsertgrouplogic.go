package group

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

type UpsertGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpsertGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertGroupLogic {
	return &UpsertGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpsertGroupLogic) UpsertGroup(req *types.UpsertGroupReq) (resp *types.UpsertGroupResp, err error) {
	// 从ctx中获取clientId
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	rpcResp, err := l.svcCtx.UamRpc.UpsertGroup(l.ctx, &uamrpc.UpsertGroupReq{
		ClientId: clientId,
		Name:     req.GroupName,
		Desc:     req.GroupDesc,
		Editable: 1,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	group := types.Group{
		Id:         rpcResp.Group.Id,
		Name:       rpcResp.Group.Name,
		Desc:       rpcResp.Group.Desc,
		CreateTime: rpcResp.Group.CreateTime,
		UpdateTime: rpcResp.Group.UpdateTime,
	}
	return &types.UpsertGroupResp{
		Group: group,
	}, nil
}
