package setup

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"uam/services/admin/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"
)

func MustSyncApiPermission(svcCtx *svc.ServiceContext) {
	if err := SyncApiPermission(svcCtx); err != nil {
		log.Fatalf("API 权限同步失败: %s", err.Error())
	}
}

func SyncApiPermission(svcCtx *svc.ServiceContext) error {
	var err error
	ctx := context.Background()
	permList, err := svcCtx.UamRpc.GetPermList(ctx, &uamrpc.GetPermListReq{ClientId: svcCtx.SysClient.Id})
	if err != nil {
		return errors.Wrap(err, "系统权限获取失败")
	}
	if svcCtx.ApiPerms == nil {
		return errors.Wrap(err, "svcCtx.ApiPerms未初始化")
	}
	for _, perm := range permList.List {
		if perm.Type == constants.PermTypeAPI {
			svcCtx.ApiPerms[perm.Key] = perm.Id
		}
	}
	return nil
}
