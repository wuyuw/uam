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

type GetRoleByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByNameLogic {
	return &GetRoleByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  通过roleName查询role
func (l *GetRoleByNameLogic) GetRoleByName(in *uamrpc.GetRoleByNameReq) (resp *uamrpc.GetRoleByNameResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	role, err := l.svcCtx.RoleModel.FindOneByName(l.ctx, in.ClientId, in.Name)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if role == nil {
		return &uamrpc.GetRoleByNameResp{Role: nil}, nil
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
	return &uamrpc.GetRoleByNameResp{Role: rpcRole}, nil
}
