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

type AddUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserRoleLogic {
	return &AddUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 用户添加角色
func (l *AddUserRoleLogic) AddUserRole(req *types.UserRoleReq) (resp *types.UserRoleResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	_, err = l.svcCtx.UamRpc.AddUserRole(l.ctx, &uamrpc.AddUserRoleReq{
		Uid:      req.Uid,
		ClientId: clientId,
		RoleId:   req.RoleId,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
