package logic

import (
	"context"

	"uam/services/model"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserRoleLogic {
	return &AddUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  用户添加角色
func (l *AddUserRoleLogic) AddUserRole(in *uamrpc.AddUserRoleReq) (resp *uamrpc.AddUserRoleResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.RoleId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if role == nil || role.ClientId != in.ClientId {
		return nil, errors.New("group不存在")
	}
	err = l.svcCtx.RelModel.AddUserRoleByRoleIds(l.ctx, nil, in.ClientId, in.Uid, []int64{in.RoleId})
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.AddUserRoleResp{}, nil
}
