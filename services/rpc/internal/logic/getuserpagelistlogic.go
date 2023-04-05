package logic

import (
	"context"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/collections"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPageListLogic {
	return &GetUserPageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户分页列表
func (l *GetUserPageListLogic) GetUserPageList(in *uamrpc.GetUserPageListReq) (resp *uamrpc.GetUserPageListResp, err error) {
	var intersectUids []int64
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	// 满足查询条件的uid列表
	searchUids, err := l.svcCtx.UserModel.FindUidsBySearch(l.ctx, in.Search)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if len(searchUids) != 0 {
		intersectUids = searchUids
	}

	// 组关联的uid列表
	groupUids, err := l.svcCtx.RelModel.FindUidsByGroupId(l.ctx, in.ClientId, in.GroupId)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if len(groupUids) != 0 {
		if len(intersectUids) == 0 {
			intersectUids = groupUids
		} else {
			intersectUids = collections.SliceIntToInt64(
				collections.Intersect(collections.SliceInt64ToInt(groupUids),
					collections.SliceInt64ToInt(intersectUids)))
		}
	}

	// 角色直接关联的uid列表
	roleUids, err := l.svcCtx.RelModel.FindUidsByRoleId(l.ctx, in.ClientId, in.RoleId)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if len(roleUids) != 0 {
		if len(intersectUids) == 0 {
			intersectUids = roleUids
		} else {
			intersectUids = collections.SliceIntToInt64(
				collections.Intersect(collections.SliceInt64ToInt(roleUids),
					collections.SliceInt64ToInt(intersectUids)))
		}
	}
	total := int64(len(intersectUids))
	page := in.Page
	pageSize := in.PageSize

	offset := pageSize * (page - 1)
	if total <= offset {
		page = 1
		offset = 0
	}
	nextOffset := offset + pageSize
	if total < nextOffset {
		nextOffset = total
	}
	pageUids := intersectUids[offset:nextOffset]
	pageUsers, err := l.svcCtx.UserModel.FindByUids(l.ctx, pageUids)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	resList := make([]*uamrpc.User, len(pageUsers))
	for i, item := range pageUsers {
		user := &uamrpc.User{
			Uid:      item.Uid,
			Nickname: item.Nickname,
		}
		groupIds, err := l.svcCtx.RelModel.FindGroupIdsByUid(l.ctx, in.ClientId, item.Uid)
		if err != nil {
			return nil, errors.Wrap(err, constants.MsgDBErr)
		}
		roleIds, err := l.svcCtx.RelModel.FindRoleIdsByUid(l.ctx, in.ClientId, item.Uid)
		if err != nil {
			return nil, errors.Wrap(err, constants.MsgDBErr)
		}
		user.Groups = groupIds
		user.Roles = roleIds
		resList[i] = user
	}
	return &uamrpc.GetUserPageListResp{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		List:     resList,
	}, nil
}
