package logic

import (
	"context"
	"fmt"
	"time"

	"uam/services/model"
	uamclient "uam/services/model/uam_client"
	uampermission "uam/services/model/uam_permission"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"
	"uam/tools/cryptox"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type AddClientLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClientLogic {
	return &AddClientLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加客户端
func (l *AddClientLogic) AddClient(in *uamrpc.AddClientReq) (resp *uamrpc.AddClientResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	client, err := l.svcCtx.ClientModel.FindOneByAppCode(l.ctx, in.AppCode)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if client != nil {
		return nil, errors.New(fmt.Sprintf("client already exists, app_code: %s", in.AppCode))
	}
	// 获取UAM系统客户端
	sysClient, err := l.svcCtx.ClientModel.FindSysClient(l.ctx)
	if err != nil {
		if err != model.ErrNotFound {
			err = errors.Wrap(err, constants.MsgDBErr)
		} else {
			err = errors.New("uam系统客户端不存在")
		}
		return nil, err
	}
	privateKey := cryptox.GetMd5(
		l.svcCtx.Config.Client.PrivateKeySalt,
		in.AppCode,
		fmt.Sprintf("%d", time.Now().Unix()))
	client = &uamclient.UamClient{
		Name:       in.Name,
		AppCode:    in.AppCode,
		PrivateKey: privateKey,
		Department: in.Department,
		Maintainer: in.Maintainer,
	}
	if err = l.AddClientInTx(sysClient.Id, client); err != nil {
		return nil, err
	}
	return &uamrpc.AddClientResp{}, nil
}

func (l *AddClientLogic) AddClientInTx(sysClientId int64, client *uamclient.UamClient) error {
	return model.ExecInTx(l.svcCtx.DB, func(tx *gorm.DB) error {
		if err := l.svcCtx.ClientModel.InsertOne(l.ctx, tx, client); err != nil {
			return err
		}
		// 添加客户端自动添加对应权限
		clientPerm := &uampermission.UamPermission{
			ClientId: sysClientId,
			Type:     constants.PermTypeClient,
			Key:      client.AppCode,
			Name:     client.Name,
			Desc:     "可操作该客户端的权限",
			Editable: 0,
		}
		if err := l.svcCtx.PermissionModel.InsertOne(l.ctx, tx, clientPerm); err != nil {
			return err
		}
		return nil
	})
}
