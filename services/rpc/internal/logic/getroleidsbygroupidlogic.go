package logic

import (
	"context"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleIdsByGroupIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleIdsByGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleIdsByGroupIdLogic {
	return &GetRoleIdsByGroupIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过groupId获取角色ID列表
func (l *GetRoleIdsByGroupIdLogic) GetRoleIdsByGroupId(in *uamrpc.GetRoleIdsByGroupIdReq) (resp *uamrpc.GetRoleIdsByGroupIdResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	roleIds, err := l.svcCtx.RelModel.FindRoleIdsByGroupId(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if roleIds == nil {
		roleIds = make([]int64, 0)
	}
	return &uamrpc.GetRoleIdsByGroupIdResp{Roles: roleIds}, nil
}
