package user

import (
	"context"
	"uam/services/rpc/pb/uamrpc"

	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/tools/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPageListLogic {
	return &UserPageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPageListLogic) UserPageList(req *types.UserPageListReq) (resp *types.UserPageListResp, err error) {
	rpcResp, err := l.svcCtx.UamRpc.GetUserPageList(l.ctx, &uamrpc.GetUserPageListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		ClientId: req.ClientId,
		GroupId:  req.GroupId,
		RoleId:   req.RoleId,
		Search:   req.Search,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	userList := make([]types.UserItem, len(rpcResp.List))
	for i, item := range rpcResp.List {
		userGroups := make([]int64, 0)
		if item.Groups != nil {
			userGroups = item.Groups
		}
		userRoles := make([]int64, 0)
		if item.Roles != nil {
			userRoles = item.Roles
		}
		userList[i] = types.UserItem{
			Uid:      item.Uid,
			Nickname: item.Nickname,
			Groups:   userGroups,
			Roles:    userRoles,
		}
	}
	return &types.UserPageListResp{
		Page:     rpcResp.Page,
		PageSize: rpcResp.PageSize,
		Total:    rpcResp.Total,
		List:     userList,
	}, nil
}
