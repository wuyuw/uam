package client

import (
	"context"
	"errors"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"
	"uam/tools/errx"
	"uam/tools/jwtx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientOptionsLogic {
	return &ClientOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientOptionsLogic) ClientOptions(req *types.ClientOptionsReq) (resp *types.ClientOptionsResp, err error) {
	// 获取用户有权访问的客户端列表
	uid, err := jwtx.GetCtxValueInt64(l.ctx, "uid")
	if err != nil {
		l.Logger.Error(err)
		return nil, errors.New("异常的用户信息")
	}
	// 获取用户权限列表
	userPerms, err := l.svcCtx.UamRpc.GetPermIdsByUid(l.ctx, &uamrpc.GetPermIdsByUidReq{
		ClientId: l.svcCtx.SysClient.Id,
		Uid:      int64(uid),
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	if len(userPerms.List) == 0 {
		return &types.ClientOptionsResp{List: []types.ClientOption{}}, nil
	}
	// 获取所有类型为客户端的权限
	clientPerms, err := l.svcCtx.UamRpc.GetPermList(l.ctx, &uamrpc.GetPermListReq{
		ClientId: l.svcCtx.SysClient.Id,
		Type:     constants.PermTypeClient,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	userClientCodes := l.getUserClientCodes(userPerms.List, clientPerms.List)
	if len(userClientCodes) == 0 {
		return &types.ClientOptionsResp{List: []types.ClientOption{}}, nil
	}
	clientListResp, err := l.svcCtx.UamRpc.GetClientListByCodes(l.ctx, &uamrpc.GetClientListByCodesReq{
		AppCodes: userClientCodes,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	clientOptions := make([]types.ClientOption, len(clientListResp.List))
	for i, client := range clientListResp.List {
		option := types.ClientOption{
			Id:      client.Id,
			Name:    client.Name,
			AppCode: client.AppCode,
		}
		groupOptions, err := l.getGroupOptionsByClientId(client.Id)
		if err != nil {
			return nil, errx.ConverRpcErr(err)
		}
		option.Groups = groupOptions
		roleOptions, err := l.getRoleOptionsByClientId(client.Id)
		if err != nil {
			return nil, errx.ConverRpcErr(err)
		}
		option.Roles = roleOptions

		permOptions, err := l.getPermOptionsByClientId(client.Id)
		if err != nil {
			return nil, errx.ConverRpcErr(err)
		}
		option.Perms = permOptions

		permTypeSet := make(map[string]struct{})
		for _, perm := range permOptions {
			_, ok := permTypeSet[perm.Type]
			if !ok {
				permTypeSet[perm.Type] = struct{}{}
			}
		}
		permTypeOptions := make([]string, 0, len(permTypeSet))
		for k := range permTypeSet {
			permTypeOptions = append(permTypeOptions, k)
		}
		if err != nil {
			return nil, errx.ConverRpcErr(err)
		}
		option.PermTypes = permTypeOptions

		clientOptions[i] = option
	}

	return &types.ClientOptionsResp{List: clientOptions}, nil
}

// 获取用户有权访问的客户端AppCode列表
func (l *ClientOptionsLogic) getUserClientCodes(userPermIds []int64, clientPerms []*uamrpc.Perm) []string {
	var userClientCodes []string
	userPermIdMap := make(map[int64]struct{})
	for _, permId := range userPermIds {
		userPermIdMap[permId] = struct{}{}
	}
	for _, perm := range clientPerms {
		if _, ok := userPermIdMap[perm.Id]; ok {
			userClientCodes = append(userClientCodes, perm.Key)
		}
	}
	return userClientCodes
}

func (l *ClientOptionsLogic) getGroupOptionsByClientId(clientId int64) ([]types.GroupOption, error) {
	rpcResp, err := l.svcCtx.UamRpc.GetGroupList(l.ctx, &uamrpc.GetGroupListReq{ClientId: clientId})
	if err != nil {
		return nil, err
	}
	groupOptions := make([]types.GroupOption, len(rpcResp.List))
	for i, item := range rpcResp.List {
		groupOptions[i] = types.GroupOption{
			Id:       item.Id,
			ClientId: item.ClientId,
			Name:     item.Name,
			Desc:     item.Desc,
		}
	}
	return groupOptions, nil
}

func (l *ClientOptionsLogic) getRoleOptionsByClientId(clientId int64) ([]types.RoleOption, error) {
	rpcResp, err := l.svcCtx.UamRpc.GetRoleList(l.ctx, &uamrpc.GetRoleListReq{ClientId: clientId})
	if err != nil {
		return nil, err
	}
	roleOptions := make([]types.RoleOption, len(rpcResp.List))
	for i, item := range rpcResp.List {
		roleOptions[i] = types.RoleOption{
			Id:       item.Id,
			ClientId: item.ClientId,
			Name:     item.Name,
			Desc:     item.Desc,
		}
	}
	return roleOptions, nil
}

func (l *ClientOptionsLogic) getPermOptionsByClientId(clientId int64) ([]types.PermOption, error) {
	rpcResp, err := l.svcCtx.UamRpc.GetPermList(l.ctx, &uamrpc.GetPermListReq{ClientId: clientId})
	if err != nil {
		return nil, err
	}
	permOptions := make([]types.PermOption, len(rpcResp.List))
	for i, item := range rpcResp.List {
		permOptions[i] = types.PermOption{
			Id:       item.Id,
			ClientId: item.ClientId,
			Type:     item.Type,
			Key:      item.Key,
			Name:     item.Name,
			Desc:     item.Desc,
		}
	}
	return permOptions, nil
}
