package logic

import (
	"context"

	uamclient "uam/services/model/uam_client"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClientLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClientLogic {
	return &UpdateClientLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新客户端
func (l *UpdateClientLogic) UpdateClient(in *uamrpc.UpdateClientReq) (resp *uamrpc.UpdateClientResp, err error) {
	var (
		client *uamclient.UamClient
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	client, err = l.svcCtx.ClientModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if client.AppCode != in.AppCode {
		return nil, errors.New("app_code禁止更新操作")
	}
	client.Name = in.Name
	client.Department = in.Department
	client.Maintainer = in.Maintainer
	client.Status = in.Status
	err = l.svcCtx.ClientModel.Update(l.ctx, client)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.UpdateClientResp{}, nil
}
