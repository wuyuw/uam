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

type GetPermKeysByUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermKeysByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermKeysByUidLogic {
	return &GetPermKeysByUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取用户权限Key列表
func (l *GetPermKeysByUidLogic) GetPermKeysByUid(in *uamrpc.GetPermKeysByUidReq) (resp *uamrpc.GetPermKeysByUidResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	permIds, err := dao.FindPermIdsByUid(l.ctx, l.svcCtx, in.ClientId, in.Uid)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	permList, err := l.svcCtx.PermissionModel.FindByIds(l.ctx, permIds)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	var permKeys []string
	// 筛选类型
	if in.PermType != "" {
		permKeys = make([]string, 0)
		for _, perm := range permList {
			if perm.Type == in.PermType {
				permKeys = append(permKeys, perm.Key)
			}
		}
	} else {
		permKeys = make([]string, len(permList))
		for i, perm := range permList {
			permKeys[i] = perm.Key
		}
	}
	return &uamrpc.GetPermKeysByUidResp{List: permKeys}, nil
}
