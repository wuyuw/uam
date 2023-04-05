package logic

import (
	"context"
	"fmt"

	"uam/services/model"
	uampermission "uam/services/model/uam_permission"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddPermLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPermLogic {
	return &AddPermLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加权限
func (l *AddPermLogic) AddPerm(in *uamrpc.AddPermReq) (resp *uamrpc.AddPermResp, err error) {
	var (
		perm *uampermission.UamPermission
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	perm, err = l.svcCtx.PermissionModel.FindOneByKey(l.ctx, in.ClientId, in.Key)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if perm != nil {
		return nil, errors.New(fmt.Sprintf("permission already exists, clientId: %d key: %s", in.ClientId, in.Key))
	}
	perm = &uampermission.UamPermission{
		ClientId: in.ClientId,
		Type:     in.Type,
		Key:      in.Key,
		Name:     in.Name,
		Desc:     in.Desc,
		Editable: 1,
	}
	err = l.svcCtx.PermissionModel.InsertOne(l.ctx, l.svcCtx.DB, perm)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.AddPermResp{}, nil
}
