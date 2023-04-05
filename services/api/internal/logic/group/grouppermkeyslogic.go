package group

import (
	"context"
	"uam/services/rpc/pb/uamrpc"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/api/internal/consts"
	"uam/services/api/internal/svc"
	"uam/services/api/internal/types"
	"uam/tools/contextx"
	"uam/tools/errx"
)

type GroupPermKeysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupPermKeysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPermKeysLogic {
	return &GroupPermKeysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPermKeysLogic) GroupPermKeys(req *types.GroupPermKeysReq) (resp *types.GroupPermKeysResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	rpcGroup, err := l.svcCtx.UamRpc.GetGroupByName(l.ctx, &uamrpc.GetGroupByNameReq{
		ClientId: clientId,
		Name:     req.GroupName,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	if rpcGroup.Group == nil {
		return nil, errors.Errorf("组不存在: %s", req.GroupName)
	}
	rpcResp, err := l.svcCtx.UamRpc.GetPermKeysByGroupId(l.ctx, &uamrpc.GetPermKeysByGroupIdReq{
		GroupId: rpcGroup.Group.Id,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	permKeys := make([]string, 0)
	if len(rpcResp.List) != 0 {
		permKeys = rpcResp.List
	}
	return &types.GroupPermKeysResp{List: permKeys}, nil
}
