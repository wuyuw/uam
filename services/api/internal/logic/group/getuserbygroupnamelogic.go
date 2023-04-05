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

type GetUserByGroupNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByGroupNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByGroupNameLogic {
	return &GetUserByGroupNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByGroupNameLogic) GetUserByGroupName(req *types.GetUserByGroupNameReq) (resp *types.GetUserByGroupNameResp, err error) {
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
		return &types.GetUserByGroupNameResp{List: []types.User{}}, nil
	}
	rpcResp, err := l.svcCtx.UamRpc.GetUserByGroupId(l.ctx, &uamrpc.GetUserByGroupIdReq{
		ClientId: clientId,
		GroupId:  rpcGroup.Group.Id,
	})
	userList := make([]types.User, len(rpcResp.List))
	for i, item := range rpcResp.List {
		userList[i] = types.User{
			Uid:      item.Uid,
			Nickname: item.Nickname,
			Email:    item.Email,
			Phone:    item.Phone,
		}
	}
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return &types.GetUserByGroupNameResp{List: userList}, nil
}
