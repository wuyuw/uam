package permission

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type UpdatePermLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePermLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermLogic {
	return &UpdatePermLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePermLogic) UpdatePerm(req *types.UpdatePermReq) (resp *types.UpdatePermResp, err error) {
	in := &uamrpc.UpdatePermReq{
		Id:       req.Id,
		ClientId: req.ClientId,
		Type:     req.Type,
		Key:      req.Key,
		Name:     req.Name,
		Desc:     req.Desc,
	}
	_, err = l.svcCtx.UamRpc.UpdatePerm(l.ctx, in)
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return &types.UpdatePermResp{}, nil
}
