package dao

import (
	"context"
	"uam/services/rpc/internal/svc"
)

// FindPermIdsByUid 获取用户拥有的所有权限ID列表
// TODO MR
func FindPermIdsByUid(ctx context.Context, svcCtx *svc.ServiceContext, clientId, uid int64) ([]int64, error) {
	var err error
	userGroupIds, err := svcCtx.RelModel.FindGroupIdsByUid(ctx, clientId, uid)
	if err != nil {
		return nil, err
	}
	userRoleIds, err := svcCtx.RelModel.FindRoleIdsByUid(ctx, clientId, uid)
	if err != nil {
		return nil, err
	}
	groupRoleIds, err := svcCtx.RelModel.FindRoleIdsByGroupIds(ctx, userGroupIds)
	if err != nil {
		return nil, err
	}
	userRoleIds = append(userRoleIds, groupRoleIds...)
	permIds, err := svcCtx.RelModel.FindPermIdsByRoleIds(ctx, userRoleIds)
	if err != nil {
		return nil, err
	}
	return permIds, nil
}

// FindPermIdsByGroupId 获取组关联的所有权限ID列表
func FindPermIdsByGroupId(ctx context.Context, svcCtx *svc.ServiceContext, groupId int64) ([]int64, error) {
	groupRoleIds, err := svcCtx.RelModel.FindRoleIdsByGroupId(ctx, groupId)
	if err != nil {
		return nil, err
	}
	permIds, err := svcCtx.RelModel.FindPermIdsByRoleIds(ctx, groupRoleIds)
	if err != nil {
		return nil, err
	}
	return permIds, nil
}
