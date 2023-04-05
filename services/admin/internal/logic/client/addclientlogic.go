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

type AddClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClientLogic {
	return &AddClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddClientLogic) AddClient(req *types.AddClientReq) (resp *types.AddClientResp, err error) {
	in := &uamrpc.AddClientReq{
		Name:       req.Name,
		AppCode:    req.AppCode,
		Department: req.Department,
		Maintainer: req.Maintainer,
	}
	_, err = l.svcCtx.UamRpc.AddClient(l.ctx, in)
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	// 刷新UC系统管理员角色权限
	if err := setup.SyncSysAdminRole(l.ctx, l.svcCtx); err != nil {
		return nil, errors.Wrap(errx.ConverRpcErr(err), "系统管理员角色权限更新失败")
	}
	return &types.AddClientResp{}, nil
}
