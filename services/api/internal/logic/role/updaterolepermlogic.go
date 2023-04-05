package role

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

type UpdateRolePermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRolePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRolePermLogic {
	return &UpdateRolePermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRolePermLogic) UpdateRolePerm(req *types.UpdateRolePermReq) (resp *types.UpdateRolePermResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	rpcResp, err := l.svcCtx.UamRpc.GetRoleByName(l.ctx, &uamrpc.GetRoleByNameReq{
		ClientId: clientId,
		Name:     req.RoleName,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	if rpcResp.Role == nil {
		return nil, errors.New("角色不存在: " + req.RoleName)
	}
	_, err = l.svcCtx.UamRpc.UpdateRole(l.ctx, &uamrpc.UpdateRoleReq{
		Id:          rpcResp.Role.Id,
		ClientId:    rpcResp.Role.ClientId,
		Name:        rpcResp.Role.Name,
		Desc:        rpcResp.Role.Desc,
		Permissions: req.PermKeys,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
