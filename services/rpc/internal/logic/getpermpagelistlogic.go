package logic

import (
	"context"

	uampermission "uam/services/model/uam_permission"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermPageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermPageListLogic {
	return &GetPermPageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取权限分页列表
func (l *GetPermPageListLogic) GetPermPageList(in *uamrpc.GetPermPageListReq) (resp *uamrpc.GetPermPageListResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	pageListResp, err := l.svcCtx.PermissionModel.FindPageList(l.ctx,
		in.Page, in.PageSize, in.ClientId, in.Type, in.Editable, in.Search)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	permList := pageListResp.List.([]*uampermission.UamPermission)
	list := make([]*uamrpc.Perm, len(permList))
	for i, perm := range permList {
		list[i] = &uamrpc.Perm{
			Id:         perm.Id,
			ClientId:   perm.ClientId,
			Type:       perm.Type,
			Key:        perm.Key,
			Name:       perm.Name,
			Desc:       perm.Desc,
			Editable:   perm.Editable,
			CreateTime: perm.CreateTime.Format(constants.DateTimeFormatTplSlashNoSec),
			UpdateTime: perm.UpdateTime.Format(constants.DateTimeFormatTplSlashNoSec),
		}
	}
	return &uamrpc.GetPermPageListResp{
		Page:     pageListResp.Page,
		PageSize: pageListResp.PageSize,
		Total:    pageListResp.Total,
		List:     list,
	}, nil
}
