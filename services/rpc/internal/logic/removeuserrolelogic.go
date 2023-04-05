package logic

import (
	"context"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserRoleLogic {
	return &RemoveUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  用户移除角色
func (l *RemoveUserRoleLogic) RemoveUserRole(in *uamrpc.RemoveUserRoleReq) (resp *uamrpc.RemoveUserRoleResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	err = l.svcCtx.RelModel.RemoveUserRoleByRoleIds(l.ctx, nil, in.ClientId, in.Uid, []int64{in.RoleId})
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.RemoveUserRoleResp{}, nil
}
