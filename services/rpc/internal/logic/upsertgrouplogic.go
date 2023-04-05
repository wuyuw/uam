package logic

import (
	"context"

	"uam/services/model"
	uamgroup "uam/services/model/uam_group"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpsertGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpsertGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertGroupLogic {
	return &UpsertGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  更新或创建角色
func (l *UpsertGroupLogic) UpsertGroup(in *uamrpc.UpsertGroupReq) (resp *uamrpc.UpsertGroupResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	group, err := l.svcCtx.GroupModel.FindOneByName(l.ctx, in.ClientId, in.Name)
	if err != nil {
		if err != model.ErrNotFound {
			return nil, errors.Wrap(err, constants.MsgDBErr)
		}
		group = &uamgroup.UamGroup{
			ClientId: in.ClientId,
			Name:     in.Name,
			Desc:     in.Desc,
			Editable: in.Editable,
		}
		err = l.svcCtx.GroupModel.InsertOne(l.ctx, nil, group)
		if err != nil {
			return nil, errors.Wrap(err, constants.MsgDBErr)
		}
	}
	group.Desc = in.Desc
	group.Editable = in.Editable

	if err = l.svcCtx.GroupModel.Update(l.ctx, group); err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	rpcGroup := &uamrpc.Group{
		Id:         group.Id,
		ClientId:   group.ClientId,
		Name:       group.Name,
		Desc:       group.Desc,
		Editable:   group.Editable,
		CreateTime: group.CreateTime.Format(constants.DateTimeFormatTplSlashNoSec),
		UpdateTime: group.UpdateTime.Format(constants.DateTimeFormatTplSlashNoSec),
	}
	return &uamrpc.UpsertGroupResp{
		Group: rpcGroup,
	}, nil
}
