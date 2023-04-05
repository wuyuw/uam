package logic

import (
	"context"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByGroupIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByGroupIdLogic {
	return &GetUserByGroupIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  查询组关联用户
func (l *GetUserByGroupIdLogic) GetUserByGroupId(in *uamrpc.GetUserByGroupIdReq) (resp *uamrpc.GetUserByGroupIdResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	uids, err := l.svcCtx.RelModel.FindUidsByGroupId(l.ctx, in.ClientId, in.GroupId)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	userList, err := l.svcCtx.UserModel.FindByUids(l.ctx, uids)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	list := make([]*uamrpc.UserInfo, len(userList))
	for i, item := range userList {
		list[i] = &uamrpc.UserInfo{
			Uid:      item.Uid,
			Nickname: item.Nickname,
			Email:    item.Email,
			Phone:    item.Phone,
		}
	}
	return &uamrpc.GetUserByGroupIdResp{List: list}, nil
}
