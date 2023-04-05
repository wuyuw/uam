package permission

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/errx"
)

type PermPageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermPageListLogic {
	return &PermPageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermPageListLogic) PermPageList(req *types.PermPageListReq) (resp *types.PermPageListResp, err error) {
	rpcResp, err := l.svcCtx.UamRpc.GetPermPageList(l.ctx, &uamrpc.GetPermPageListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		ClientId: req.ClientId,
		Type:     req.Type,
		Editable: req.Editable,
		Search:   req.Search,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	permList := make([]types.PermItem, len(rpcResp.List))
	for i, item := range rpcResp.List {
		permList[i] = types.PermItem{
			Id:         item.Id,
			ClientId:   item.ClientId,
			Type:       item.Type,
			Key:        item.Key,
			Name:       item.Name,
			Desc:       item.Desc,
			Editable:   item.Editable,
			CreateTime: item.CreateTime,
			UpdateTime: item.UpdateTime,
		}
	}
	return &types.PermPageListResp{
		Page:     rpcResp.Page,
		PageSize: rpcResp.PageSize,
		Total:    rpcResp.Total,
		List:     permList,
	}, nil
}
