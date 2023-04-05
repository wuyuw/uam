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

type RemoveUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserRoleLogic {
	return &RemoveUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveUserRoleLogic) RemoveUserRole(req *types.UserRoleReq) (resp *types.UserRoleResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	_, err = l.svcCtx.UamRpc.RemoveUserRole(l.ctx, &uamrpc.RemoveUserRoleReq{
		Uid:      req.Uid,
		ClientId: clientId,
		RoleId:   req.RoleId,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
