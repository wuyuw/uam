package logic

import (
	"context"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermListLogic {
	return &GetPermListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPermListLogic) GetPermList(in *uamrpc.GetPermListReq) (resp *uamrpc.GetPermListResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	permList, err := l.svcCtx.PermissionModel.FindByType(l.ctx, in.ClientId, in.Type)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
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

	return &uamrpc.GetPermListResp{List: list}, nil
}
