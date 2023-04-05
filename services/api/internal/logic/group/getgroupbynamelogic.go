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

type GetGroupByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupByNameLogic {
	return &GetGroupByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupByNameLogic) GetGroupByName(req *types.GetGroupByNameReq) (resp *types.GetGroupByNameResp, err error) {
	// 从ctx中获取clientId
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	rpcResp, err := l.svcCtx.UamRpc.GetGroupByName(l.ctx, &uamrpc.GetGroupByNameReq{
		ClientId: clientId,
		Name:     req.GroupName,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	if rpcResp.Group == nil {
		return &types.GetGroupByNameResp{List: []types.Group{}}, nil
	}
	group := types.Group{
		Id:         rpcResp.Group.Id,
		Name:       rpcResp.Group.Name,
		Desc:       rpcResp.Group.Desc,
		CreateTime: rpcResp.Group.CreateTime,
		UpdateTime: rpcResp.Group.UpdateTime,
	}
	return &types.GetGroupByNameResp{List: []types.Group{group}}, nil
}
