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

type GetGroupByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupByNameLogic {
	return &GetGroupByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  通过groupName查询group
func (l *GetGroupByNameLogic) GetGroupByName(in *uamrpc.GetGroupByNameReq) (resp *uamrpc.GetGroupByNameResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	group, err := l.svcCtx.GroupModel.FindOneByName(l.ctx, in.ClientId, in.Name)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if group == nil {
		return &uamrpc.GetGroupByNameResp{Group: nil}, nil
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
	return &uamrpc.GetGroupByNameResp{
		Group: rpcGroup,
	}, nil
}
