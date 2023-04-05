package user

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

type UserGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGroupsLogic {
	return &UserGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserGroupsLogic) UserGroups(req *types.UserGroupsReq) (resp *types.UserGroupsResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	rpcResp, err := l.svcCtx.UamRpc.GetGroupListByUid(l.ctx, &uamrpc.GetGroupListByUidReq{
		ClientId: clientId,
		Uid:      req.Uid,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	groupNames := make([]string, len(rpcResp.List))
	for i, item := range rpcResp.List {
		groupNames[i] = item.Name
	}
	return &types.UserGroupsResp{List: groupNames}, nil
}
