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

type GroupPermIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupPermIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPermIdsLogic {
	return &GroupPermIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPermIdsLogic) GroupPermIds(req *types.GroupPermIdsReq) (resp *types.GroupPermIdsResp, err error) {
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
	rpcResp, err := l.svcCtx.UamRpc.GetPermIdsByGroupId(l.ctx, &uamrpc.GetPermIdsByGroupIdReq{
		GroupId: rpcGroup.Group.Id,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	permIds := make([]int64, 0)
	if len(rpcResp.List) != 0 {
		permIds = rpcResp.List
	}
	return &types.GroupPermIdsResp{List: permIds}, nil
}
