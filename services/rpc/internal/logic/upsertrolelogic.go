package logic

import (
	"context"

	"uam/services/model"
	uamrole "uam/services/model/uam_role"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpsertRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpsertRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertRoleLogic {
	return &UpsertRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  更新或创建角色-不关联权限
func (l *UpsertRoleLogic) UpsertRole(in *uamrpc.UpsertRoleReq) (resp *uamrpc.UpsertRoleResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	role, err := l.svcCtx.RoleModel.FindOneByName(l.ctx, in.ClientId, in.Name)
	if err != nil {
		if err != model.ErrNotFound {
			return nil, errors.Wrap(err, constants.MsgDBErr)
		}
		role = &uamrole.UamRole{
			ClientId: in.ClientId,
			Name:     in.Name,
			Desc:     in.Desc,
			Editable: in.Editable,
		}
		err = l.svcCtx.RoleModel.InsertOne(l.ctx, nil, role)
		if err != nil {
			return nil, errors.Wrap(err, constants.MsgDBErr)
		}
	}
	role.Desc = in.Desc
	role.Editable = in.Editable

	if err = l.svcCtx.RoleModel.Update(l.ctx, role); err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	rpcRole := &uamrpc.Role{
		Id:         role.Id,
		ClientId:   role.ClientId,
		Name:       role.Name,
		Desc:       role.Desc,
		Editable:   role.Editable,
		CreateTime: role.CreateTime.Format(constants.DateTimeFormatTplSlashNoSec),
		UpdateTime: role.UpdateTime.Format(constants.DateTimeFormatTplSlashNoSec),
	}
	return &uamrpc.UpsertRoleResp{Role: rpcRole}, nil
}
