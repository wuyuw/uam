package logic

import (
	"context"

	"uam/services/model"
	uamclient "uam/services/model/uam_client"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SyncSysClientLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncSysClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncSysClientLogic {
	return &SyncSysClientLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 同步UC系统客户端
func (l *SyncSysClientLogic) SyncSysClient(in *uamrpc.SyncSysClientReq) (resp *uamrpc.SyncSysClientResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	sysClient, err := l.svcCtx.ClientModel.FindSysClient(l.ctx)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if sysClient != nil {
		if sysClient.Name != in.Name || sysClient.AppCode != in.AppCode {
			sysClient.Name = in.Name
			sysClient.AppCode = in.AppCode
			if err := l.svcCtx.ClientModel.Update(l.ctx, sysClient); err != nil {
				return nil, errors.Wrap(err, constants.MsgDBErr)
			}
		}
	} else {
		sysClient = &uamclient.UamClient{
			AppCode: in.AppCode,
			Name:    in.Name,
			Type:    2,
		}
		if err := l.svcCtx.ClientModel.InsertOne(l.ctx, l.svcCtx.DB, sysClient); err != nil {
			return nil, errors.Wrap(err, constants.MsgDBErr)
		}
	}
	client := &uamrpc.Client{
		Id:      sysClient.Id,
		Name:    sysClient.Name,
		AppCode: sysClient.AppCode,
		Type:    sysClient.Type,
	}
	return &uamrpc.SyncSysClientResp{Client: client}, nil
}
