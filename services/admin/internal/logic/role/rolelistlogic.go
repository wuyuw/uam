package role

import (
	"context"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *RoleListLogic) RoleList(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	rpcResp, err := l.svcCtx.UamRpc.GetRoleList(l.ctx, &uamrpc.GetRoleListReq{
		ClientId: req.ClientId,
		Editable: req.Editable,
	})

	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	roleList := make([]types.Role, len(rpcResp.List))
	for i, item := range rpcResp.List {
		permList, err := l.getPermListByRoleId(item.Id)
		if err != nil {
			return nil, errx.ConverRpcErr(err)
		}
		permKeys := make([]string, len(permList))
		for i, perm := range permList {
			permKeys[i] = perm.Key
		}
		roleList[i] = types.Role{
			Id:          item.Id,
			ClientId:    item.ClientId,
			Name:        item.Name,
			Desc:        item.Desc,
			Editable:    item.Editable,
			CreateTime:  item.CreateTime,
			UpdateTime:  item.UpdateTime,
			Permissions: permKeys,
		}
	}

	return &types.RoleListResp{List: roleList}, nil
}

func (l *RoleListLogic) getPermListByRoleId(roleId int64) (permList []*uamrpc.Perm, err error) {
	rpcResp, err := l.svcCtx.UamRpc.GetPermListByRoleId(l.ctx, &uamrpc.GetPermListByRoleIdReq{
		RoleId: roleId,
	})
	if err != nil {
		return nil, err
	}
	return rpcResp.List, nil
}
