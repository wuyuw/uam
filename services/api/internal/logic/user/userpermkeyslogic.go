package user

import (
	"context"
	"errors"
	"uam/services/rpc/pb/uamrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/api/internal/consts"
	"uam/services/api/internal/svc"
	"uam/services/api/internal/types"
	"uam/tools/contextx"
	"uam/tools/errx"
)

type UserPermKeysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPermKeysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPermKeysLogic {
	return &UserPermKeysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取用户所有权限Key列表
func (l *UserPermKeysLogic) UserPermKeys(req *types.UserPermKeysReq) (resp *types.UserPermKeysResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	rpcResp, err := l.svcCtx.UamRpc.GetPermKeysByUid(l.ctx, &uamrpc.GetPermKeysByUidReq{
		ClientId: clientId,
		Uid:      req.Uid,
		PermType: req.PermType,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	permKeys := make([]string, 0)
	if len(rpcResp.List) != 0 {
		permKeys = rpcResp.List
	}
	return &types.UserPermKeysResp{List: permKeys}, nil
}
