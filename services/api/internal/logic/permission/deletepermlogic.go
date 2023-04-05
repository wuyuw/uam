package permission

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

type DeletePermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePermLogic {
	return &DeletePermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePermLogic) DeletePerm(req *types.DeletePermReq) (resp *types.DeletePermResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	_, err = l.svcCtx.UamRpc.DeletePermByKey(l.ctx, &uamrpc.DeletePermByKeyReq{
		ClientId: clientId,
		Key:      req.Key,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
