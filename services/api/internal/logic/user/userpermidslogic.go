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

type UserPermIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPermIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPermIdsLogic {
	return &UserPermIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取用户权限ID列表
func (l *UserPermIdsLogic) UserPermIds(req *types.UserPermIdsReq) (resp *types.UserPermIdsResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	rpcResp, err := l.svcCtx.UamRpc.GetPermIdsByUid(l.ctx, &uamrpc.GetPermIdsByUidReq{
		ClientId: clientId,
		Uid:      req.Uid,
		PermType: req.PermType,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	permIds := make([]int64, 0)
	if len(rpcResp.List) != 0 {
		permIds = rpcResp.List
	}
	return &types.UserPermIdsResp{List: permIds}, nil
}
