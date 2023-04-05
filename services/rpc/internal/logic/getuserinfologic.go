package logic

import (
	"context"
	"fmt"

	"uam/services/model"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserInfo 通过UID获取用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *uamrpc.GetUserInfoReq) (resp *uamrpc.GetUserInfoResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	user, err := l.svcCtx.UserModel.FindOneByUid(l.ctx, in.Uid)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New(fmt.Sprintf("用户不存在, UID: %d", in.Uid))
		}
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.GetUserInfoResp{
		Uid:      user.Uid,
		Nickname: user.Nickname,
		Email:    user.Email,
		Phone:    user.Phone,
	}, nil
}
