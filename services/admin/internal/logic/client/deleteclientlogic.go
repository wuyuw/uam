package client

import (
	"context"
	"uam/services/admin/internal/setup"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClientLogic {
	return &DeleteClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteClientLogic) DeleteClient(req *types.DeleteClientReq) (resp *types.DeleteClientResp, err error) {
	_, err = l.svcCtx.UamRpc.DeleteClient(l.ctx, &uamrpc.DeleteClientReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}

	// 刷新UC系统管理员角色权限
	if err := setup.SyncSysAdminRole(l.ctx, l.svcCtx); err != nil {
		return nil, errors.Wrap(errx.ConverRpcErr(err), "系统管理员角色权限更新失败")
	}
	return &types.DeleteClientResp{}, nil
}
