// Code generated by goctl. DO NOT EDIT!
// Source: uamrpc.proto

package uam

import (
	"context"

	"uam/services/rpc/pb/uamrpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddAuthLocalReq            = uamrpc.AddAuthLocalReq
	AddAuthLocalResp           = uamrpc.AddAuthLocalResp
	AddClientReq               = uamrpc.AddClientReq
	AddClientResp              = uamrpc.AddClientResp
	AddGroupReq                = uamrpc.AddGroupReq
	AddGroupResp               = uamrpc.AddGroupResp
	AddPermReq                 = uamrpc.AddPermReq
	AddPermResp                = uamrpc.AddPermResp
	AddRoleReq                 = uamrpc.AddRoleReq
	AddRoleResp                = uamrpc.AddRoleResp
	AddUserGroupReq            = uamrpc.AddUserGroupReq
	AddUserGroupResp           = uamrpc.AddUserGroupResp
	AddUserRoleReq             = uamrpc.AddUserRoleReq
	AddUserRoleResp            = uamrpc.AddUserRoleResp
	AuthLocal                  = uamrpc.AuthLocal
	BatchAddPermReq            = uamrpc.BatchAddPermReq
	BatchAddPermResp           = uamrpc.BatchAddPermResp
	Client                     = uamrpc.Client
	DeleteClientReq            = uamrpc.DeleteClientReq
	DeleteClientResp           = uamrpc.DeleteClientResp
	DeleteGroupReq             = uamrpc.DeleteGroupReq
	DeleteGroupResp            = uamrpc.DeleteGroupResp
	DeletePermByKeyReq         = uamrpc.DeletePermByKeyReq
	DeletePermByKeyResp        = uamrpc.DeletePermByKeyResp
	DeletePermReq              = uamrpc.DeletePermReq
	DeletePermResp             = uamrpc.DeletePermResp
	DeleteRoleReq              = uamrpc.DeleteRoleReq
	DeleteRoleResp             = uamrpc.DeleteRoleResp
	GetAuthLocalByUsernameReq  = uamrpc.GetAuthLocalByUsernameReq
	GetAuthLocalByUsernameResp = uamrpc.GetAuthLocalByUsernameResp
	GetClientByCodeReq         = uamrpc.GetClientByCodeReq
	GetClientByCodeResp        = uamrpc.GetClientByCodeResp
	GetClientListByCodesReq    = uamrpc.GetClientListByCodesReq
	GetClientListByCodesResp   = uamrpc.GetClientListByCodesResp
	GetClientListReq           = uamrpc.GetClientListReq
	GetClientListResp          = uamrpc.GetClientListResp
	GetGroupByNameReq          = uamrpc.GetGroupByNameReq
	GetGroupByNameResp         = uamrpc.GetGroupByNameResp
	GetGroupListByUidReq       = uamrpc.GetGroupListByUidReq
	GetGroupListByUidResp      = uamrpc.GetGroupListByUidResp
	GetGroupListReq            = uamrpc.GetGroupListReq
	GetGroupListResp           = uamrpc.GetGroupListResp
	GetPermIdsByGroupIdReq     = uamrpc.GetPermIdsByGroupIdReq
	GetPermIdsByGroupIdResp    = uamrpc.GetPermIdsByGroupIdResp
	GetPermIdsByUidReq         = uamrpc.GetPermIdsByUidReq
	GetPermIdsByUidResp        = uamrpc.GetPermIdsByUidResp
	GetPermKeysByGroupIdReq    = uamrpc.GetPermKeysByGroupIdReq
	GetPermKeysByGroupIdResp   = uamrpc.GetPermKeysByGroupIdResp
	GetPermKeysByUidReq        = uamrpc.GetPermKeysByUidReq
	GetPermKeysByUidResp       = uamrpc.GetPermKeysByUidResp
	GetPermListByRoleIdReq     = uamrpc.GetPermListByRoleIdReq
	GetPermListByRoleIdResp    = uamrpc.GetPermListByRoleIdResp
	GetPermListReq             = uamrpc.GetPermListReq
	GetPermListResp            = uamrpc.GetPermListResp
	GetPermPageListReq         = uamrpc.GetPermPageListReq
	GetPermPageListResp        = uamrpc.GetPermPageListResp
	GetPermTypeOptionsReq      = uamrpc.GetPermTypeOptionsReq
	GetPermTypeOptionsResp     = uamrpc.GetPermTypeOptionsResp
	GetRoleByNameReq           = uamrpc.GetRoleByNameReq
	GetRoleByNameResp          = uamrpc.GetRoleByNameResp
	GetRoleIdsByGroupIdReq     = uamrpc.GetRoleIdsByGroupIdReq
	GetRoleIdsByGroupIdResp    = uamrpc.GetRoleIdsByGroupIdResp
	GetRoleListReq             = uamrpc.GetRoleListReq
	GetRoleListResp            = uamrpc.GetRoleListResp
	GetSysClientReq            = uamrpc.GetSysClientReq
	GetSysClientResp           = uamrpc.GetSysClientResp
	GetUserByGroupIdReq        = uamrpc.GetUserByGroupIdReq
	GetUserByGroupIdResp       = uamrpc.GetUserByGroupIdResp
	GetUserInfoReq             = uamrpc.GetUserInfoReq
	GetUserInfoResp            = uamrpc.GetUserInfoResp
	GetUserPageListReq         = uamrpc.GetUserPageListReq
	GetUserPageListResp        = uamrpc.GetUserPageListResp
	Group                      = uamrpc.Group
	Perm                       = uamrpc.Perm
	PermItem                   = uamrpc.PermItem
	RemoveUserGroupReq         = uamrpc.RemoveUserGroupReq
	RemoveUserGroupResp        = uamrpc.RemoveUserGroupResp
	RemoveUserRoleReq          = uamrpc.RemoveUserRoleReq
	RemoveUserRoleResp         = uamrpc.RemoveUserRoleResp
	Role                       = uamrpc.Role
	SyncSysClientReq           = uamrpc.SyncSysClientReq
	SyncSysClientResp          = uamrpc.SyncSysClientResp
	UpdateClientReq            = uamrpc.UpdateClientReq
	UpdateClientResp           = uamrpc.UpdateClientResp
	UpdateGroupReq             = uamrpc.UpdateGroupReq
	UpdateGroupResp            = uamrpc.UpdateGroupResp
	UpdatePermReq              = uamrpc.UpdatePermReq
	UpdatePermResp             = uamrpc.UpdatePermResp
	UpdateRoleReq              = uamrpc.UpdateRoleReq
	UpdateRoleResp             = uamrpc.UpdateRoleResp
	UpdateUserPermReq          = uamrpc.UpdateUserPermReq
	UpdateUserPermResp         = uamrpc.UpdateUserPermResp
	UpsertGroupReq             = uamrpc.UpsertGroupReq
	UpsertGroupResp            = uamrpc.UpsertGroupResp
	UpsertRoleReq              = uamrpc.UpsertRoleReq
	UpsertRoleResp             = uamrpc.UpsertRoleResp
	User                       = uamrpc.User
	UserInfo                   = uamrpc.UserInfo

	Uam interface {
		//  GetAuthLocalByUsername 根据用户名查询本地认证记录
		GetAuthLocalByUsername(ctx context.Context, in *GetAuthLocalByUsernameReq, opts ...grpc.CallOption) (*GetAuthLocalByUsernameResp, error)
		//  AddAuthLocal 添加本地认证记录
		AddAuthLocal(ctx context.Context, in *AddAuthLocalReq, opts ...grpc.CallOption) (*AddAuthLocalResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		GetClientList(ctx context.Context, in *GetClientListReq, opts ...grpc.CallOption) (*GetClientListResp, error)
		//  通过AppCode列表查询客户端
		GetClientListByCodes(ctx context.Context, in *GetClientListByCodesReq, opts ...grpc.CallOption) (*GetClientListByCodesResp, error)
		//  通过AppCode查询客户端
		GetClientByCode(ctx context.Context, in *GetClientByCodeReq, opts ...grpc.CallOption) (*GetClientByCodeResp, error)
		AddClient(ctx context.Context, in *AddClientReq, opts ...grpc.CallOption) (*AddClientResp, error)
		UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...grpc.CallOption) (*UpdateClientResp, error)
		DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error)
		GetPermTypeOptions(ctx context.Context, in *GetPermTypeOptionsReq, opts ...grpc.CallOption) (*GetPermTypeOptionsResp, error)
		GetPermList(ctx context.Context, in *GetPermListReq, opts ...grpc.CallOption) (*GetPermListResp, error)
		GetPermPageList(ctx context.Context, in *GetPermPageListReq, opts ...grpc.CallOption) (*GetPermPageListResp, error)
		GetPermListByRoleId(ctx context.Context, in *GetPermListByRoleIdReq, opts ...grpc.CallOption) (*GetPermListByRoleIdResp, error)
		//  获取用户权限Id列表
		GetPermIdsByUid(ctx context.Context, in *GetPermIdsByUidReq, opts ...grpc.CallOption) (*GetPermIdsByUidResp, error)
		//  获取用户权限Key列表
		GetPermKeysByUid(ctx context.Context, in *GetPermKeysByUidReq, opts ...grpc.CallOption) (*GetPermKeysByUidResp, error)
		//  获取组关联的权限Id列表
		GetPermIdsByGroupId(ctx context.Context, in *GetPermIdsByGroupIdReq, opts ...grpc.CallOption) (*GetPermIdsByGroupIdResp, error)
		//  获取组关联的权限Key列表
		GetPermKeysByGroupId(ctx context.Context, in *GetPermKeysByGroupIdReq, opts ...grpc.CallOption) (*GetPermKeysByGroupIdResp, error)
		//  添加权限
		AddPerm(ctx context.Context, in *AddPermReq, opts ...grpc.CallOption) (*AddPermResp, error)
		//  更新权限
		UpdatePerm(ctx context.Context, in *UpdatePermReq, opts ...grpc.CallOption) (*UpdatePermResp, error)
		//  根据ID删除权限
		DeletePerm(ctx context.Context, in *DeletePermReq, opts ...grpc.CallOption) (*DeletePermResp, error)
		//  根据key删除权限
		DeletePermByKey(ctx context.Context, in *DeletePermByKeyReq, opts ...grpc.CallOption) (*DeletePermByKeyResp, error)
		//  批量添加权限
		BatchAddPerm(ctx context.Context, in *BatchAddPermReq, opts ...grpc.CallOption) (*BatchAddPermResp, error)
		//  获取所有角色列表，支持按是否可编辑筛选
		GetRoleList(ctx context.Context, in *GetRoleListReq, opts ...grpc.CallOption) (*GetRoleListResp, error)
		//  通过roleName查询role
		GetRoleByName(ctx context.Context, in *GetRoleByNameReq, opts ...grpc.CallOption) (*GetRoleByNameResp, error)
		AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*AddRoleResp, error)
		//  更新或创建角色
		UpsertRole(ctx context.Context, in *UpsertRoleReq, opts ...grpc.CallOption) (*UpsertRoleResp, error)
		UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error)
		DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleResp, error)
		GetRoleIdsByGroupId(ctx context.Context, in *GetRoleIdsByGroupIdReq, opts ...grpc.CallOption) (*GetRoleIdsByGroupIdResp, error)
		//  通过groupName查询group
		GetGroupByName(ctx context.Context, in *GetGroupByNameReq, opts ...grpc.CallOption) (*GetGroupByNameResp, error)
		//  获取用户关联组列表
		GetGroupListByUid(ctx context.Context, in *GetGroupListByUidReq, opts ...grpc.CallOption) (*GetGroupListByUidResp, error)
		//  获取客户端下组列表
		GetGroupList(ctx context.Context, in *GetGroupListReq, opts ...grpc.CallOption) (*GetGroupListResp, error)
		AddGroup(ctx context.Context, in *AddGroupReq, opts ...grpc.CallOption) (*AddGroupResp, error)
		//  更新或创建角色
		UpsertGroup(ctx context.Context, in *UpsertGroupReq, opts ...grpc.CallOption) (*UpsertGroupResp, error)
		//  更新组
		UpdateGroup(ctx context.Context, in *UpdateGroupReq, opts ...grpc.CallOption) (*UpdateGroupResp, error)
		//  删除组
		DeleteGroup(ctx context.Context, in *DeleteGroupReq, opts ...grpc.CallOption) (*DeleteGroupResp, error)
		//  查询组关联用户
		GetUserByGroupId(ctx context.Context, in *GetUserByGroupIdReq, opts ...grpc.CallOption) (*GetUserByGroupIdResp, error)
		//  获取用户分页列表
		GetUserPageList(ctx context.Context, in *GetUserPageListReq, opts ...grpc.CallOption) (*GetUserPageListResp, error)
		//  更新用户权限
		UpdateUserPerm(ctx context.Context, in *UpdateUserPermReq, opts ...grpc.CallOption) (*UpdateUserPermResp, error)
		//  用户入组
		AddUserGroup(ctx context.Context, in *AddUserGroupReq, opts ...grpc.CallOption) (*AddUserGroupResp, error)
		//  用户出组
		RemoveUserGroup(ctx context.Context, in *RemoveUserGroupReq, opts ...grpc.CallOption) (*RemoveUserGroupResp, error)
		//  用户添加角色
		AddUserRole(ctx context.Context, in *AddUserRoleReq, opts ...grpc.CallOption) (*AddUserRoleResp, error)
		//  用户移除角色
		RemoveUserRole(ctx context.Context, in *RemoveUserRoleReq, opts ...grpc.CallOption) (*RemoveUserRoleResp, error)
		//  同步系统客户端
		SyncSysClient(ctx context.Context, in *SyncSysClientReq, opts ...grpc.CallOption) (*SyncSysClientResp, error)
		//  获取系统客户端
		GetSysClient(ctx context.Context, in *GetSysClientReq, opts ...grpc.CallOption) (*GetSysClientResp, error)
	}

	defaultUam struct {
		cli zrpc.Client
	}
)

func NewUam(cli zrpc.Client) Uam {
	return &defaultUam{
		cli: cli,
	}
}

//  GetAuthLocalByUsername 根据用户名查询本地认证记录
func (m *defaultUam) GetAuthLocalByUsername(ctx context.Context, in *GetAuthLocalByUsernameReq, opts ...grpc.CallOption) (*GetAuthLocalByUsernameResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetAuthLocalByUsername(ctx, in, opts...)
}

//  AddAuthLocal 添加本地认证记录
func (m *defaultUam) AddAuthLocal(ctx context.Context, in *AddAuthLocalReq, opts ...grpc.CallOption) (*AddAuthLocalResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.AddAuthLocal(ctx, in, opts...)
}

func (m *defaultUam) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUam) GetClientList(ctx context.Context, in *GetClientListReq, opts ...grpc.CallOption) (*GetClientListResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetClientList(ctx, in, opts...)
}

//  通过AppCode列表查询客户端
func (m *defaultUam) GetClientListByCodes(ctx context.Context, in *GetClientListByCodesReq, opts ...grpc.CallOption) (*GetClientListByCodesResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetClientListByCodes(ctx, in, opts...)
}

//  通过AppCode查询客户端
func (m *defaultUam) GetClientByCode(ctx context.Context, in *GetClientByCodeReq, opts ...grpc.CallOption) (*GetClientByCodeResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetClientByCode(ctx, in, opts...)
}

func (m *defaultUam) AddClient(ctx context.Context, in *AddClientReq, opts ...grpc.CallOption) (*AddClientResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.AddClient(ctx, in, opts...)
}

func (m *defaultUam) UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...grpc.CallOption) (*UpdateClientResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.UpdateClient(ctx, in, opts...)
}

func (m *defaultUam) DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.DeleteClient(ctx, in, opts...)
}

func (m *defaultUam) GetPermTypeOptions(ctx context.Context, in *GetPermTypeOptionsReq, opts ...grpc.CallOption) (*GetPermTypeOptionsResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetPermTypeOptions(ctx, in, opts...)
}

func (m *defaultUam) GetPermList(ctx context.Context, in *GetPermListReq, opts ...grpc.CallOption) (*GetPermListResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetPermList(ctx, in, opts...)
}

func (m *defaultUam) GetPermPageList(ctx context.Context, in *GetPermPageListReq, opts ...grpc.CallOption) (*GetPermPageListResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetPermPageList(ctx, in, opts...)
}

func (m *defaultUam) GetPermListByRoleId(ctx context.Context, in *GetPermListByRoleIdReq, opts ...grpc.CallOption) (*GetPermListByRoleIdResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetPermListByRoleId(ctx, in, opts...)
}

//  获取用户权限Id列表
func (m *defaultUam) GetPermIdsByUid(ctx context.Context, in *GetPermIdsByUidReq, opts ...grpc.CallOption) (*GetPermIdsByUidResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetPermIdsByUid(ctx, in, opts...)
}

//  获取用户权限Key列表
func (m *defaultUam) GetPermKeysByUid(ctx context.Context, in *GetPermKeysByUidReq, opts ...grpc.CallOption) (*GetPermKeysByUidResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetPermKeysByUid(ctx, in, opts...)
}

//  获取组关联的权限Id列表
func (m *defaultUam) GetPermIdsByGroupId(ctx context.Context, in *GetPermIdsByGroupIdReq, opts ...grpc.CallOption) (*GetPermIdsByGroupIdResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetPermIdsByGroupId(ctx, in, opts...)
}

//  获取组关联的权限Key列表
func (m *defaultUam) GetPermKeysByGroupId(ctx context.Context, in *GetPermKeysByGroupIdReq, opts ...grpc.CallOption) (*GetPermKeysByGroupIdResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetPermKeysByGroupId(ctx, in, opts...)
}

//  添加权限
func (m *defaultUam) AddPerm(ctx context.Context, in *AddPermReq, opts ...grpc.CallOption) (*AddPermResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.AddPerm(ctx, in, opts...)
}

//  更新权限
func (m *defaultUam) UpdatePerm(ctx context.Context, in *UpdatePermReq, opts ...grpc.CallOption) (*UpdatePermResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.UpdatePerm(ctx, in, opts...)
}

//  根据ID删除权限
func (m *defaultUam) DeletePerm(ctx context.Context, in *DeletePermReq, opts ...grpc.CallOption) (*DeletePermResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.DeletePerm(ctx, in, opts...)
}

//  根据key删除权限
func (m *defaultUam) DeletePermByKey(ctx context.Context, in *DeletePermByKeyReq, opts ...grpc.CallOption) (*DeletePermByKeyResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.DeletePermByKey(ctx, in, opts...)
}

//  批量添加权限
func (m *defaultUam) BatchAddPerm(ctx context.Context, in *BatchAddPermReq, opts ...grpc.CallOption) (*BatchAddPermResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.BatchAddPerm(ctx, in, opts...)
}

//  获取所有角色列表，支持按是否可编辑筛选
func (m *defaultUam) GetRoleList(ctx context.Context, in *GetRoleListReq, opts ...grpc.CallOption) (*GetRoleListResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetRoleList(ctx, in, opts...)
}

//  通过roleName查询role
func (m *defaultUam) GetRoleByName(ctx context.Context, in *GetRoleByNameReq, opts ...grpc.CallOption) (*GetRoleByNameResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetRoleByName(ctx, in, opts...)
}

func (m *defaultUam) AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*AddRoleResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.AddRole(ctx, in, opts...)
}

//  更新或创建角色
func (m *defaultUam) UpsertRole(ctx context.Context, in *UpsertRoleReq, opts ...grpc.CallOption) (*UpsertRoleResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.UpsertRole(ctx, in, opts...)
}

func (m *defaultUam) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

func (m *defaultUam) DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

func (m *defaultUam) GetRoleIdsByGroupId(ctx context.Context, in *GetRoleIdsByGroupIdReq, opts ...grpc.CallOption) (*GetRoleIdsByGroupIdResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetRoleIdsByGroupId(ctx, in, opts...)
}

//  通过groupName查询group
func (m *defaultUam) GetGroupByName(ctx context.Context, in *GetGroupByNameReq, opts ...grpc.CallOption) (*GetGroupByNameResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetGroupByName(ctx, in, opts...)
}

//  获取用户关联组列表
func (m *defaultUam) GetGroupListByUid(ctx context.Context, in *GetGroupListByUidReq, opts ...grpc.CallOption) (*GetGroupListByUidResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetGroupListByUid(ctx, in, opts...)
}

//  获取客户端下组列表
func (m *defaultUam) GetGroupList(ctx context.Context, in *GetGroupListReq, opts ...grpc.CallOption) (*GetGroupListResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetGroupList(ctx, in, opts...)
}

func (m *defaultUam) AddGroup(ctx context.Context, in *AddGroupReq, opts ...grpc.CallOption) (*AddGroupResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.AddGroup(ctx, in, opts...)
}

//  更新或创建角色
func (m *defaultUam) UpsertGroup(ctx context.Context, in *UpsertGroupReq, opts ...grpc.CallOption) (*UpsertGroupResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.UpsertGroup(ctx, in, opts...)
}

//  更新组
func (m *defaultUam) UpdateGroup(ctx context.Context, in *UpdateGroupReq, opts ...grpc.CallOption) (*UpdateGroupResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.UpdateGroup(ctx, in, opts...)
}

//  删除组
func (m *defaultUam) DeleteGroup(ctx context.Context, in *DeleteGroupReq, opts ...grpc.CallOption) (*DeleteGroupResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.DeleteGroup(ctx, in, opts...)
}

//  查询组关联用户
func (m *defaultUam) GetUserByGroupId(ctx context.Context, in *GetUserByGroupIdReq, opts ...grpc.CallOption) (*GetUserByGroupIdResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetUserByGroupId(ctx, in, opts...)
}

//  获取用户分页列表
func (m *defaultUam) GetUserPageList(ctx context.Context, in *GetUserPageListReq, opts ...grpc.CallOption) (*GetUserPageListResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetUserPageList(ctx, in, opts...)
}

//  更新用户权限
func (m *defaultUam) UpdateUserPerm(ctx context.Context, in *UpdateUserPermReq, opts ...grpc.CallOption) (*UpdateUserPermResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.UpdateUserPerm(ctx, in, opts...)
}

//  用户入组
func (m *defaultUam) AddUserGroup(ctx context.Context, in *AddUserGroupReq, opts ...grpc.CallOption) (*AddUserGroupResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.AddUserGroup(ctx, in, opts...)
}

//  用户出组
func (m *defaultUam) RemoveUserGroup(ctx context.Context, in *RemoveUserGroupReq, opts ...grpc.CallOption) (*RemoveUserGroupResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.RemoveUserGroup(ctx, in, opts...)
}

//  用户添加角色
func (m *defaultUam) AddUserRole(ctx context.Context, in *AddUserRoleReq, opts ...grpc.CallOption) (*AddUserRoleResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.AddUserRole(ctx, in, opts...)
}

//  用户移除角色
func (m *defaultUam) RemoveUserRole(ctx context.Context, in *RemoveUserRoleReq, opts ...grpc.CallOption) (*RemoveUserRoleResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.RemoveUserRole(ctx, in, opts...)
}

//  同步系统客户端
func (m *defaultUam) SyncSysClient(ctx context.Context, in *SyncSysClientReq, opts ...grpc.CallOption) (*SyncSysClientResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.SyncSysClient(ctx, in, opts...)
}

//  获取系统客户端
func (m *defaultUam) GetSysClient(ctx context.Context, in *GetSysClientReq, opts ...grpc.CallOption) (*GetSysClientResp, error) {
	client := uamrpc.NewUamClient(m.cli.Conn())
	return client.GetSysClient(ctx, in, opts...)
}