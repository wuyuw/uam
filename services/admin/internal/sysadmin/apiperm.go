package sysadmin

import "net/http"

type ApiPerm struct {
	Method string
	Path   string
	Name   string
	Desc   string
}

type Route []string

var (
	RouteGetClientList    = Route{http.MethodGet, "/uam/admin/v1/clients"}
	RouteAddClient        = Route{http.MethodPost, "/uam/admin/v1/client"}
	RouteEditClient       = Route{http.MethodPut, "/uam/admin/v1/client/:id"}
	RouteDelClient        = Route{http.MethodDelete, "/uam/admin/v1/client/:id"}
	RouteGetClientOptions = Route{http.MethodGet, "/uam/admin/v1/client-options"}
	RouteGetPermList      = Route{http.MethodGet, "/uam/admin/v1/permissions"}
	RouteAddPerm          = Route{http.MethodPost, "/uam/admin/v1/permission"}
	RouteEditPerm         = Route{http.MethodPut, "/uam/admin/v1/permission/:id"}
	RouteDelPerm          = Route{http.MethodDelete, "/uam/admin/v1/permission/:id"}
	RouteGetRoleList      = Route{http.MethodGet, "/uam/admin/v1/roles"}
	RouteAddRole          = Route{http.MethodPost, "/uam/admin/v1/role"}
	RouteEditRole         = Route{http.MethodPut, "/uam/admin/v1/role/:id"}
	RouteDelRole          = Route{http.MethodDelete, "/uam/admin/v1/role/:id"}
	RouteGetGroups        = Route{http.MethodGet, "/uam/admin/v1/groups"}
	RouteAddGroup         = Route{http.MethodPost, "/uam/admin/v1/group"}
	RouteEditGroup        = Route{http.MethodPut, "/uam/admin/v1/group/:id"}
	RouteDelGroup         = Route{http.MethodDelete, "/uam/admin/v1/group/:id"}
	RouteGetUserList      = Route{http.MethodGet, "/uam/admin/v1/users"}
	RouteEditUserPerm     = Route{http.MethodPut, "/uam/admin/v1/user/:uid"}
)

// ApiPermList API权限列表
var ApiPermList = []ApiPerm{
	{Method: RouteGetClientList[0], Path: RouteGetClientList[1], Name: "获取客户端列表", Desc: "可获取客户端列表"},
	{Method: RouteAddClient[0], Path: RouteAddClient[1], Name: "添加客户端", Desc: "可添加客户端"},
	{Method: RouteEditClient[0], Path: RouteEditClient[1], Name: "编辑客户端", Desc: "可编辑客户端"},
	{Method: RouteDelClient[0], Path: RouteDelClient[1], Name: "删除客户端", Desc: "可删除客户端"},
	{Method: RouteGetClientOptions[0], Path: RouteGetClientOptions[1], Name: "获取客户端可选项", Desc: "可获取客户端可选项"},

	{Method: RouteGetPermList[0], Path: RouteGetPermList[1], Name: "获取权限列表", Desc: "可获取权限列表"},
	{Method: RouteAddPerm[0], Path: RouteAddPerm[1], Name: "添加权限", Desc: "可添加权限"},
	{Method: RouteEditPerm[0], Path: RouteEditPerm[1], Name: "编辑权限", Desc: "可编辑权限"},
	{Method: RouteDelPerm[0], Path: RouteDelPerm[1], Name: "删除权限", Desc: "可删除权限"},

	{Method: RouteGetRoleList[0], Path: RouteGetRoleList[1], Name: "获取角色列表", Desc: "可获取角色列表"},
	{Method: RouteAddRole[0], Path: RouteAddRole[1], Name: "添加角色", Desc: "可添加角色"},
	{Method: RouteEditRole[0], Path: RouteEditRole[1], Name: "编辑角色", Desc: "可编辑角色"},
	{Method: RouteDelRole[0], Path: RouteDelRole[1], Name: "删除角色", Desc: "可删除角色"},

	{Method: RouteGetGroups[0], Path: RouteGetGroups[1], Name: "获取组列表", Desc: "可获取组列表"},
	{Method: RouteAddGroup[0], Path: RouteAddGroup[1], Name: "添加组", Desc: "可添加组"},
	{Method: RouteEditGroup[0], Path: RouteEditGroup[1], Name: "编辑组", Desc: "可编辑组"},
	{Method: RouteDelGroup[0], Path: RouteDelGroup[1], Name: "删除组", Desc: "可删除组"},

	{Method: RouteGetUserList[0], Path: RouteGetUserList[1], Name: "获取用户列表", Desc: "可获取用户列表"},
	{Method: RouteEditUserPerm[0], Path: RouteEditUserPerm[1], Name: "编辑用户权限", Desc: "可编辑用户权限"},
}
