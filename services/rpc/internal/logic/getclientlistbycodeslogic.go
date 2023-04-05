package logic

import (
	"context"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetClientListByCodesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClientListByCodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClientListByCodesLogic {
	return &GetClientListByCodesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  通过AppCode列表查询客户端
func (l *GetClientListByCodesLogic) GetClientListByCodes(in *uamrpc.GetClientListByCodesReq) (resp *uamrpc.GetClientListByCodesResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	clients, err := l.svcCtx.ClientModel.FindByCodes(l.ctx, in.AppCodes)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	list := make([]*uamrpc.Client, len(clients))
	for i, client := range clients {
		list[i] = &uamrpc.Client{
			Id:         client.Id,
			Name:       client.Name,
			AppCode:    client.AppCode,
			PrivateKey: client.PrivateKey,
			Department: client.Department,
			Maintainer: client.Maintainer,
			Status:     client.Status,
			Type:       client.Type,
			CreateTime: client.CreateTime.Format(constants.DateTimeFormatTplSlashNoSec),
			UpdateTime: client.UpdateTime.Format(constants.DateTimeFormatTplSlashNoSec),
		}
	}
	return &uamrpc.GetClientListByCodesResp{List: list}, nil
}
