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

type BatchAddPermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchAddPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchAddPermLogic {
	return &BatchAddPermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  批量添加权限
func (l *BatchAddPermLogic) BatchAddPerm(in *uamrpc.BatchAddPermReq) (resp *uamrpc.BatchAddPermResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	permList := make([]*uampermission.UamPermission, len(in.List))
	for i, item := range in.List {
		permList[i] = &uampermission.UamPermission{
			ClientId: in.ClientId,
			Type:     item.Type,
			Key:      item.Key,
			Name:     item.Name,
			Desc:     item.Desc,
			Editable: item.Editable,
		}
	}
	if err := l.svcCtx.PermissionModel.UpsertPermList(l.ctx, permList); err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.BatchAddPermResp{}, nil
}
