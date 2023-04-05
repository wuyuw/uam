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

type GroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListLogic) GroupList(req *types.GroupListReq) (resp *types.GroupListResp, err error) {
	// 从ctx中获取clientId
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	rpcResp, err := l.svcCtx.UamRpc.GetGroupList(l.ctx, &uamrpc.GetGroupListReq{
		ClientId: clientId,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	groupList := make([]types.Group, len(rpcResp.List))
	for i, item := range rpcResp.List {
		group := types.Group{
			Id:         item.Id,
			Name:       item.Name,
			Desc:       item.Desc,
			CreateTime: item.CreateTime,
			UpdateTime: item.UpdateTime,
		}
		groupList[i] = group
	}
	return &types.GroupListResp{
		List: groupList,
	}, nil
}
