package logic

import (
	"context"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveUserGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserGroupLogic {
	return &RemoveUserGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户出组
func (l *RemoveUserGroupLogic) RemoveUserGroup(in *uamrpc.RemoveUserGroupReq) (resp *uamrpc.RemoveUserGroupResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	err = l.svcCtx.RelModel.RemoveUserGroupByGroupIds(l.ctx, nil, in.ClientId, in.Uid, []int64{in.GroupId})
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.RemoveUserGroupResp{}, nil
}
