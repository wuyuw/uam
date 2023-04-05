package logic

import (
	"context"

	uamgroup "uam/services/model/uam_group"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupListLogic {
	return &GetGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupListLogic) GetGroupList(in *uamrpc.GetGroupListReq) (resp *uamrpc.GetGroupListResp, err error) {
	var (
		groups []*uamgroup.UamGroup
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	groups, err = l.svcCtx.GroupModel.FindByEditable(l.ctx, in.ClientId, in.Editable)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	list := make([]*uamrpc.Group, len(groups))
	for i, group := range groups {
		list[i] = &uamrpc.Group{
			Id:         group.Id,
			ClientId:   group.ClientId,
			Name:       group.Name,
			Desc:       group.Desc,
			Editable:   group.Editable,
			CreateTime: group.CreateTime.Format(constants.DateTimeFormatTplSlashNoSec),
			UpdateTime: group.UpdateTime.Format(constants.DateTimeFormatTplSlashNoSec),
		}
	}
	return &uamrpc.GetGroupListResp{List: list}, nil
}
