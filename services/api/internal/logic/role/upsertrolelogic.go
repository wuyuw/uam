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

type UpsertRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpsertRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertRoleLogic {
	return &UpsertRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 更新或创建角色
func (l *UpsertRoleLogic) UpsertRole(req *types.UpsertRoleReq) (resp *types.UpsertRoleResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	rpcResp, err := l.svcCtx.UamRpc.UpsertRole(l.ctx, &uamrpc.UpsertRoleReq{
		ClientId: clientId,
		Name:     req.RoleName,
		Desc:     req.RoleDesc,
		Editable: 1,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	role := types.Role{
		Id:         rpcResp.Role.Id,
		Name:       rpcResp.Role.Name,
		Desc:       rpcResp.Role.Desc,
		CreateTime: rpcResp.Role.CreateTime,
		UpdateTime: rpcResp.Role.UpdateTime,
	}
	return &types.UpsertRoleResp{Role: role}, nil
}
