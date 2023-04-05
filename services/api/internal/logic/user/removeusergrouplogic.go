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

type RemoveUserGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserGroupLogic {
	return &RemoveUserGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 用户出组
func (l *RemoveUserGroupLogic) RemoveUserGroup(req *types.UserGroupReq) (resp *types.UserGroupResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	_, err = l.svcCtx.UamRpc.RemoveUserGroup(l.ctx, &uamrpc.RemoveUserGroupReq{
		Uid:      req.Uid,
		ClientId: clientId,
		GroupId:  req.GroupId,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
