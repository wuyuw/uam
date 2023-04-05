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

type GetGroupListByUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupListByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupListByUidLogic {
	return &GetGroupListByUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取用户关联组列表
func (l *GetGroupListByUidLogic) GetGroupListByUid(in *uamrpc.GetGroupListByUidReq) (resp *uamrpc.GetGroupListByUidResp, err error) {
	var (
		groups []*uamgroup.UamGroup
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	groupIds, err := l.svcCtx.RelModel.FindGroupIdsByUid(l.ctx, in.ClientId, in.Uid)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	groups, err = l.svcCtx.GroupModel.FindByIds(l.ctx, in.ClientId, groupIds)
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
	return &uamrpc.GetGroupListByUidResp{List: list}, nil
}
