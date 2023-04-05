package logic

import (
	"context"

	"uam/services/model"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetClientByCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClientByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClientByCodeLogic {
	return &GetClientByCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  通过AppCode查询客户端
func (l *GetClientByCodeLogic) GetClientByCode(in *uamrpc.GetClientByCodeReq) (resp *uamrpc.GetClientByCodeResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	client, err := l.svcCtx.ClientModel.FindOneByAppCode(l.ctx, in.AppCode)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if client == nil {
		return &uamrpc.GetClientByCodeResp{
			Client: nil,
		}, nil
	}
	clientPtr := &uamrpc.Client{
		Id:         client.Id,
		Name:       client.Name,
		AppCode:    client.AppCode,
		PrivateKey: client.PrivateKey,
		Status:     client.Status,
	}
	return &uamrpc.GetClientByCodeResp{
		Client: clientPtr,
	}, nil
}
