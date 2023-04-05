package core

import (
	"context"
	"uam/tools/jwtx"

	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (resp *types.LogoutResp, err error) {
	_, err = l.svcCtx.Redis.Del(jwtx.JWTRedisPrefix + req.Token)
	if err != nil {
		l.Logger.Errorf("redis error: %v", err)
		return nil, err
	}
	return &types.LogoutResp{}, nil
}
