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
	"gorm.io/gorm"
)

type DeleteClientLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClientLogic {
	return &DeleteClientLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TODO MR
func (l *DeleteClientLogic) DeleteClient(in *uamrpc.DeleteClientReq) (resp *uamrpc.DeleteClientResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	client, err := l.svcCtx.ClientModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if client == nil {
		return &uamrpc.DeleteClientResp{}, nil
	}
	// 获取uam系统客户端
	sysClient, err := l.svcCtx.ClientModel.FindSysClient(l.ctx)
	if err != nil {
		if err != model.ErrNotFound {
			err = errors.Wrap(err, constants.MsgDBErr)
		} else {
			err = errors.New("uam系统客户端不存在")
		}
		return nil, err
	}
	if err = l.DeleteClientInTx(sysClient.Id, client); err != nil {
		return nil, err
	}
	return &uamrpc.DeleteClientResp{}, nil
}

func (l *DeleteClientLogic) DeleteClientInTx(sysClientId int64, client *uamclient.UamClient) error {
	return model.ExecInTx(l.svcCtx.DB, func(tx *gorm.DB) error {
		if err := l.svcCtx.ClientModel.Delete(l.ctx, tx, client.Id); err != nil {
			return err
		}
		clientPerm, err := l.svcCtx.PermissionModel.FindOneByKey(l.ctx, sysClientId, client.AppCode)
		if err != nil {
			if err != model.ErrNotFound {
				err = errors.Wrap(err, constants.MsgDBErr)
			} else {
				err = errors.New("客户端对应权限不存在")
			}
			return err
		}
		if err := l.svcCtx.PermissionModel.Delete(l.ctx, tx, clientPerm.Id); err != nil {
			return err
		}
		return nil
	})
}
