package group

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type GroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListLogic) GroupList(req *types.GroupListReq) (resp *types.GroupListResp, err error) {
	rpcResp, err := l.svcCtx.UamRpc.GetGroupList(l.ctx, &uamrpc.GetGroupListReq{
		ClientId: req.ClientId,
		Editable: req.Editable,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	groupList := make([]types.Group, len(rpcResp.List))
	for i, item := range rpcResp.List {
		group := types.Group{
			Id:         item.Id,
			ClientId:   item.ClientId,
			Name:       item.Name,
			Desc:       item.Desc,
			Editable:   item.Editable,
			CreateTime: item.CreateTime,
			UpdateTime: item.UpdateTime,
		}
		roleIds, err := l.getRoleIdsByGroupId(group.Id)
		if err != nil {
			return nil, errx.ConverRpcErr(err)
		}
		group.Roles = roleIds
		groupList[i] = group
	}
	return &types.GroupListResp{
		List: groupList,
	}, nil
}

func (l *GroupListLogic) getRoleIdsByGroupId(groupId int64) ([]int64, error) {
	rpcResp, err := l.svcCtx.UamRpc.GetRoleIdsByGroupId(l.ctx, &uamrpc.GetRoleIdsByGroupIdReq{
		Id: groupId,
	})
	if err != nil {
		return nil, err
	}
	res := rpcResp.Roles
	if res == nil {
		res = make([]int64, 0)
	}
	return res, nil
}
