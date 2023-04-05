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

type AddPermBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPermBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPermBatchLogic {
	return &AddPermBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPermBatchLogic) AddPermBatch(req *types.AddPermBatchReq) (resp *types.AddPermBatchResp, err error) {
	// 从ctx中获取clientId
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	permList := make([]*uamrpc.PermItem, len(req.Permissions))
	for i, item := range req.Permissions {
		permList[i] = &uamrpc.PermItem{
			Type:     item.Type,
			Key:      item.Key,
			Name:     item.Name,
			Desc:     item.Desc,
			Editable: 1,
		}
	}
	_, err = l.svcCtx.UamRpc.BatchAddPerm(l.ctx, &uamrpc.BatchAddPermReq{
		ClientId: clientId,
		List:     permList,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
