package logic

import (
	"context"
	"fmt"
	uamgroup "uam/services/model/uam_group"
	uamrel "uam/services/model/uam_rel"
	uamrole "uam/services/model/uam_role"
	"uam/services/rpc/pb/uamrpc"

	"uam/services/model"
	"uam/services/rpc/internal/svc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type AddGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupLogic {
	return &AddGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加组
func (l *AddGroupLogic) AddGroup(in *uamrpc.AddGroupReq) (resp *uamrpc.AddGroupResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	ok, err := l.hasDupGroupName(in.ClientId, in.Name)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if ok {
		return nil, errors.New(fmt.Sprintf("组已存在: %s", in.Name))
	}
	// 有效的角色
	validRoles, err := l.svcCtx.RoleModel.FindListByIds(l.ctx, in.ClientId, in.Roles)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	group := &uamgroup.UamGroup{
		ClientId: in.ClientId,
		Name:     in.Name,
		Desc:     in.Desc,
		Editable: 1,
	}

	if err = l.addGroupInTx(group, validRoles); err != nil {
		return nil, err
	}
	return &uamrpc.AddGroupResp{}, nil
}

// 是否存在重复的组名
func (l *AddGroupLogic) hasDupGroupName(clientId int64, groupName string) (bool, error) {
	dupRecord, err := l.svcCtx.GroupModel.FindOneByName(l.ctx, clientId, groupName)
	if err != nil && err != model.ErrNotFound {
		return false, err
	}
	if dupRecord != nil {
		return true, nil
	}
	return false, nil
}

func (l *AddGroupLogic) addGroupInTx(group *uamgroup.UamGroup, roles []*uamrole.UamRole) error {
	return model.ExecInTx(l.svcCtx.DB, func(tx *gorm.DB) error {
		// 添加组
		if err := l.svcCtx.GroupModel.InsertOne(l.ctx, tx, group); err != nil {
			return err
		}
		if len(roles) > 0 {
			// 关联角色
			relGroupRoles := make([]uamrel.RelGroupRole, len(roles))
			for i, role := range roles {
				relGroupRoles[i] = uamrel.RelGroupRole{
					GroupId: group.Id,
					RoleId:  role.Id,
				}
			}
			if err := tx.Create(&relGroupRoles).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
