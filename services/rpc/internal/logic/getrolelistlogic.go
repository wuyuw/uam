package logic

import (
	"context"

	uamrole "uam/services/model/uam_role"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取角色列表
func (l *GetRoleListLogic) GetRoleList(in *uamrpc.GetRoleListReq) (resp *uamrpc.GetRoleListResp, err error) {
	var (
		roles []*uamrole.UamRole
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	roles, err = l.svcCtx.RoleModel.FindByEditable(l.ctx, in.ClientId, in.Editable)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	list := make([]*uamrpc.Role, len(roles))
	for i, role := range roles {
		list[i] = &uamrpc.Role{
			Id:         role.Id,
			ClientId:   role.ClientId,
			Name:       role.Name,
			Desc:       role.Desc,
			Editable:   role.Editable,
			CreateTime: role.CreateTime.Format(constants.DateTimeFormatTplSlashNoSec),
			UpdateTime: role.UpdateTime.Format(constants.DateTimeFormatTplSlashNoSec),
		}
	}
	return &uamrpc.GetRoleListResp{List: list}, nil
}
