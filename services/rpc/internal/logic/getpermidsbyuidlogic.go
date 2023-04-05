package logic

import (
	"context"

	"uam/services/rpc/internal/service/dao"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermIdsByUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermIdsByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermIdsByUidLogic {
	return &GetPermIdsByUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取用户权限Id列表
func (l *GetPermIdsByUidLogic) GetPermIdsByUid(in *uamrpc.GetPermIdsByUidReq) (resp *uamrpc.GetPermIdsByUidResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	permIds, err := dao.FindPermIdsByUid(l.ctx, l.svcCtx, in.ClientId, in.Uid)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	// 过滤类型
	if in.PermType != "" {
		permList, err := l.svcCtx.PermissionModel.FindByIds(l.ctx, permIds)
		if err != nil {
			return nil, errors.Wrap(err, constants.MsgDBErr)
		}
		permIds = make([]int64, 0)
		for _, perm := range permList {
			if perm.Type == in.PermType {
				permIds = append(permIds, perm.Id)
			}
		}
	}
	return &uamrpc.GetPermIdsByUidResp{
		List: permIds,
	}, nil
}
