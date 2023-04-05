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

type GetRoleByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByNameLogic {
	return &GetRoleByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleByNameLogic) GetRoleByName(req *types.GetRoleByNameReq) (resp *types.GetRoleByNameResp, err error) {
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
		return &types.GetRoleByNameResp{List: []types.Role{}}, nil
	}
	role := types.Role{
		Id:         rpcResp.Role.Id,
		Name:       rpcResp.Role.Name,
		Desc:       rpcResp.Role.Desc,
		CreateTime: rpcResp.Role.CreateTime,
		UpdateTime: rpcResp.Role.UpdateTime,
	}
	return &types.GetRoleByNameResp{List: []types.Role{role}}, nil
}
