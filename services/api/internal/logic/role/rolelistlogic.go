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

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取全部角色列表
func (l *RoleListLogic) RoleList(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	rpcResp, err := l.svcCtx.UamRpc.GetRoleList(l.ctx, &uamrpc.GetRoleListReq{
		ClientId: clientId,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	roleList := make([]types.Role, len(rpcResp.List))
	for i, item := range rpcResp.List {
		roleList[i] = types.Role{
			Id:         item.Id,
			Name:       item.Name,
			Desc:       item.Desc,
			CreateTime: item.CreateTime,
			UpdateTime: item.UpdateTime,
		}
	}
	return &types.RoleListResp{List: roleList}, nil
}
