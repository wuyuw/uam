package logic

import (
	"context"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermListByRoleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermListByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermListByRoleIdLogic {
	return &GetPermListByRoleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过roleId获取权限列表
func (l *GetPermListByRoleIdLogic) GetPermListByRoleId(in *uamrpc.GetPermListByRoleIdReq) (resp *uamrpc.GetPermListByRoleIdResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	permIds, err := l.svcCtx.RelModel.FindPermIdsByRoleIds(l.ctx, []int64{in.RoleId})
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	permList, err := l.svcCtx.PermissionModel.FindByIds(l.ctx, permIds)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	list := make([]*uamrpc.Perm, len(permList))
	for i, item := range permList {
		list[i] = &uamrpc.Perm{
			Id:         item.Id,
			ClientId:   item.ClientId,
			Type:       item.Type,
			Key:        item.Key,
			Name:       item.Name,
			Desc:       item.Desc,
			Editable:   item.Editable,
			CreateTime: item.CreateTime.Format(constants.DateTimeFormatTplSlashNoSec),
			UpdateTime: item.UpdateTime.Format(constants.DateTimeFormatTplSlashNoSec),
		}
	}

	return &uamrpc.GetPermListByRoleIdResp{
		List: list,
	}, nil
}
