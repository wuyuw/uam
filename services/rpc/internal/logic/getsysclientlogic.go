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

type GetSysClientLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysClientLogic {
	return &GetSysClientLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取UC后台系统客户端
func (l *GetSysClientLogic) GetSysClient(in *uamrpc.GetSysClientReq) (resp *uamrpc.GetSysClientResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	sysClient, err := l.svcCtx.ClientModel.FindSysClient(l.ctx)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("系统客户端未创建")
		} else {
			return nil, errors.Wrap(err, constants.MsgDBErr)
		}
	}
	client := &uamrpc.Client{
		Id:      sysClient.Id,
		Name:    sysClient.Name,
		AppCode: sysClient.AppCode,
		Type:    sysClient.Type,
	}

	return &uamrpc.GetSysClientResp{
		Client: client,
	}, nil
}
