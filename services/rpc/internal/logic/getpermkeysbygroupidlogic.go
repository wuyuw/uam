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

type GetPermKeysByGroupIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermKeysByGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermKeysByGroupIdLogic {
	return &GetPermKeysByGroupIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取组关联的权限Key列表
func (l *GetPermKeysByGroupIdLogic) GetPermKeysByGroupId(in *uamrpc.GetPermKeysByGroupIdReq) (resp *uamrpc.GetPermKeysByGroupIdResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	permIds, err := dao.FindPermIdsByGroupId(l.ctx, l.svcCtx, in.GroupId)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	permList, err := l.svcCtx.PermissionModel.FindByIds(l.ctx, permIds)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	permKeys := make([]string, len(permList))
	for i, perm := range permList {
		permKeys[i] = perm.Key
	}
	return &uamrpc.GetPermKeysByGroupIdResp{List: permKeys}, nil
}
