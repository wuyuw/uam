package core

import (
	"context"
	"errors"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
	"uam/tools/jwtx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// 从ctx中获取uid
	uid, err := jwtx.GetCtxValueInt64(l.ctx, "uid")
	if err != nil {
		return nil, errors.New("异常的用户信息")
	}
	rpcUserInfo, err := l.svcCtx.UamRpc.GetUserInfo(l.ctx, &uamrpc.GetUserInfoReq{Uid: int64(uid)})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	userPermKeys, err := l.svcCtx.UamRpc.GetPermKeysByUid(l.ctx, &uamrpc.GetPermKeysByUidReq{
		ClientId: l.svcCtx.SysClient.Id,
		Uid:      int64(uid),
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	userPermissions := make([]string, 0)
	if len(userPermKeys.List) != 0 {
		userPermissions = userPermKeys.List
	}
	resp = &types.UserInfoResp{
		Uid:         rpcUserInfo.Uid,
		Nickname:    rpcUserInfo.Nickname,
		Email:       rpcUserInfo.Email,
		Phone:       rpcUserInfo.Phone,
		Permissions: userPermissions,
	}
	return
}
