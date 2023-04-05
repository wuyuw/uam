package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"uam/services/model"
	uamrole "uam/services/model/uam_role"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"
	"uam/tools/kqueue"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除角色
func (l *DeleteRoleLogic) DeleteRole(in *uamrpc.DeleteRoleReq) (resp *uamrpc.DeleteRoleResp, err error) {
	var (
		role *uamrole.UamRole
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	role, err = l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if role == nil {
		return &uamrpc.DeleteRoleResp{}, nil
	}
	if role.Editable == 0 {
		return nil, errors.New(constants.MsgEditableErr)
	}

	err = l.deleteRoleInTx(in.Id)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}

	return &uamrpc.DeleteRoleResp{}, nil
}

func (l *DeleteRoleLogic) deleteRoleInTx(roleId int64) error {
	return model.ExecInTx(l.svcCtx.DB, func(tx *gorm.DB) error {
		var err error
		// 删除角色
		if err = l.svcCtx.RoleModel.Delete(l.ctx, tx, roleId); err != nil {
			return err
		}
		// 删除角色-权限关联关系
		if err = l.svcCtx.RelModel.RemoveRolePermByRoleId(l.ctx, tx, roleId); err != nil {
			return err
		}
		// 删除角色-组关联关系
		if err = l.svcCtx.RelModel.RemoveGroupRoleByRoleId(l.ctx, tx, roleId); err != nil {
			return err
		}

		// 删除角色-用户关联关系
		// 关联关系删除放在MQ中进行，避免大量记录删除造成性能抖动
		msg := kqueue.RelUpdateMqMsg{
			Table: kqueue.RelUpdateTableRole,
			Id:    roleId,
		}
		body, err := json.Marshal(msg)
		if err != nil {
			return errors.Wrap(err, "MQ消息序列化失败")
		}
		if err := l.svcCtx.RelUpdatePusher.Push(string(body)); err != nil {
			l.Logger.Error("push err: %s", err)
			return errors.Wrap(err, fmt.Sprintf("MQ消息发送失败, msg: %+v", msg))
		}
		return nil
	})
}
