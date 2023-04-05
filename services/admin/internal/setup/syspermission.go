package setup

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/sysadmin"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest"
)

// 客户端管理员角色
var (
	ClientAdminName  = "客户端管理员"
	ClientAdminDesc  = "可登录UAM后台进行权限管理及授权操作"
	ClientAdminPerms = []string{
		strings.Join(sysadmin.RouteGetClientOptions, " "),
		strings.Join(sysadmin.RouteGetPermList, " "),
		strings.Join(sysadmin.RouteAddPerm, " "),
		strings.Join(sysadmin.RouteEditPerm, " "),
		strings.Join(sysadmin.RouteDelPerm, " "),
		strings.Join(sysadmin.RouteGetRoleList, " "),
		strings.Join(sysadmin.RouteAddRole, " "),
		strings.Join(sysadmin.RouteEditRole, " "),
		strings.Join(sysadmin.RouteDelRole, " "),
		strings.Join(sysadmin.RouteGetGroups, " "),
		strings.Join(sysadmin.RouteAddGroup, " "),
		strings.Join(sysadmin.RouteEditGroup, " "),
		strings.Join(sysadmin.RouteDelGroup, " "),
		strings.Join(sysadmin.RouteGetUserList, " "),
		strings.Join(sysadmin.RouteEditUserPerm, " "),
	}
)

func MustSyncSysPermission(svcCtx *svc.ServiceContext) {
	if err := SyncSysPermission(svcCtx); err != nil {
		log.Fatalf("系统权限同步失败: %s", err.Error())
	}
}

// 自动创建"UAM-Admin"客户端
// 自动将所有需鉴权接口创建为权限
// 自动创建"UAM管理员"角色，并将所有UAM系统权限关联至该角色
func SyncSysPermission(svcCtx *svc.ServiceContext) error {
	var err error
	ctx := context.Background()
	if err = SyncSysClient(ctx, svcCtx); err != nil {
		return err
	}
	if err = SyncSysPerms(ctx, svcCtx); err != nil {
		return err
	}
	if err = SyncSysAdminRole(ctx, svcCtx); err != nil {
		return err
	}

	if err = SyncClientAdminRole(ctx, svcCtx); err != nil {
		return err
	}
	return nil
}

// 自动同步系统客户端
func SyncSysClient(ctx context.Context, svcCtx *svc.ServiceContext) error {
	sysClient, err := svcCtx.UamRpc.SyncSysClient(ctx, &uamrpc.SyncSysClientReq{
		AppCode: svcCtx.Config.System.ClientCode,
		Name:    svcCtx.Config.System.ClientName,
	})
	if err != nil {
		return errors.Wrap(err, "系统客户端同步失败")
	}
	svcCtx.SysClient.Id = sysClient.Client.Id
	svcCtx.SysClient.AppCode = sysClient.Client.AppCode
	svcCtx.SysClient.Name = sysClient.Client.Name
	return nil
}

// 自动将所有需鉴权接口创建为权限
func SyncSysPerms(ctx context.Context, svcCtx *svc.ServiceContext) error {
	var err error
	sysPerms := make([]*uamrpc.PermItem, 0)
	sysClientAccessPerm := &uamrpc.PermItem{
		Type:     constants.PermTypeClient,
		Key:      svcCtx.Config.System.ClientCode,
		Name:     svcCtx.Config.System.ClientName,
		Desc:     "可操作该客户端的所有权限",
		Editable: 0,
	}
	sysPerms = append(sysPerms, sysClientAccessPerm)
	for _, item := range sysadmin.ApiPermList {
		sysPerms = append(sysPerms, &uamrpc.PermItem{
			Type:     constants.PermTypeAPI,
			Key:      fmt.Sprintf("%s %s", item.Method, item.Path),
			Name:     item.Name,
			Desc:     item.Desc,
			Editable: 0,
		})
	}
	_, err = svcCtx.UamRpc.BatchAddPerm(ctx, &uamrpc.BatchAddPermReq{
		ClientId: svcCtx.SysClient.Id,
		List:     sysPerms,
	})
	if err != nil {
		return errors.Wrap(err, "系统权限同步失败")
	}
	return nil
}

func SyncSysAdminRole(ctx context.Context, svcCtx *svc.ServiceContext) error {

	var err error
	errMsg := "系统管理员角色同步失败"
	// 更新或创建管理员角色
	sysAdmin, err := svcCtx.UamRpc.UpsertRole(ctx, &uamrpc.UpsertRoleReq{
		ClientId: svcCtx.SysClient.Id,
		Name:     svcCtx.Config.System.AdminRole,
		Desc:     "UAM系统管理员",
		Editable: 0,
	})
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	// 将系统所有权限关联至系统管理员
	allPerms, err := svcCtx.UamRpc.GetPermList(ctx, &uamrpc.GetPermListReq{ClientId: svcCtx.SysClient.Id})
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	permKeys := make([]string, len(allPerms.List))
	for i, item := range allPerms.List {
		permKeys[i] = item.Key
	}
	_, err = svcCtx.UamRpc.UpdateRole(ctx, &uamrpc.UpdateRoleReq{
		Id:          sysAdmin.Role.Id,
		ClientId:    svcCtx.SysClient.Id,
		Name:        sysAdmin.Role.Name,
		Desc:        sysAdmin.Role.Desc,
		Permissions: permKeys,
	})
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	return nil
}

func SyncClientAdminRole(ctx context.Context, svcCtx *svc.ServiceContext) error {
	var err error
	errMsg := "客户端管理员角色同步失败"
	// 更新或创建管理员角色
	clientAdmin, err := svcCtx.UamRpc.UpsertRole(ctx, &uamrpc.UpsertRoleReq{
		ClientId: svcCtx.SysClient.Id,
		Name:     ClientAdminName,
		Desc:     ClientAdminDesc,
		Editable: 0,
	})
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	// 将权限关联至客户端管理员
	_, err = svcCtx.UamRpc.UpdateRole(ctx, &uamrpc.UpdateRoleReq{
		Id:          clientAdmin.Role.Id,
		ClientId:    svcCtx.SysClient.Id,
		Name:        clientAdmin.Role.Name,
		Desc:        clientAdmin.Role.Desc,
		Permissions: ClientAdminPerms,
	})
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	return nil
}

// 通过反射获取server上已注册的路由表
func GetRoutesByReflect(server *rest.Server) []rest.Route {
	// 获取server
	rfvServerPtr := reflect.ValueOf(server)
	rfvServer := rfvServerPtr.Elem()
	// 获取ngin字段
	rfvNginPtr := rfvServer.FieldByName("ngin")
	rfvNgin := rfvNginPtr.Elem()
	// 获取featureRoutes字段
	rfvFeaturedRoutes := rfvNgin.FieldByName("routes")
	// 对于未导出字段，通过unsafe.Pointer访问
	rfvFeaturedRoutes = reflect.NewAt(rfvFeaturedRoutes.Type(), unsafe.Pointer(rfvFeaturedRoutes.UnsafeAddr())).Elem()
	// 遍历featureRoutes字段
	allRoutes := make([]rest.Route, 0)
	for i := 0; i < rfvFeaturedRoutes.Len(); i++ {
		// 获取routes字段
		rfvRoutes := rfvFeaturedRoutes.Index(i).FieldByName("routes")
		// 对于未导出字段，通过unsafe.Pointer访问
		rfvRoutes = reflect.NewAt(rfvRoutes.Type(), unsafe.Pointer(rfvRoutes.UnsafeAddr())).Elem()
		// 转换为导出的rest.Route类型
		routes := rfvRoutes.Interface().([]rest.Route)
		allRoutes = append(allRoutes, routes...)
	}
	return allRoutes
}
