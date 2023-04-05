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

type AddUserGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserGroupLogic {
	return &AddUserGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  用户入组
func (l *AddUserGroupLogic) AddUserGroup(in *uamrpc.AddUserGroupReq) (resp *uamrpc.AddUserGroupResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	group, err := l.svcCtx.GroupModel.FindOne(l.ctx, in.GroupId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if group == nil || group.ClientId != in.ClientId {
		return nil, errors.New("group不存在")
	}
	err = l.svcCtx.RelModel.AddUserGroupByGroupIds(l.ctx, nil, in.ClientId, in.Uid, []int64{in.GroupId})
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.AddUserGroupResp{}, nil
}
