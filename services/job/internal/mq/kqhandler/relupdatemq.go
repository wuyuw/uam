package kqhandler

import (
	"context"
	"fmt"
	"uam/services/job/internal/svc"
	"uam/services/model"
	uamgroup "uam/services/model/uam_group"
	uamrole "uam/services/model/uam_role"
	"uam/tools/constants"
	"uam/tools/kqueue"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RelUpdateMq struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *RelUpdateMq {
	return &RelUpdateMq{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (m *RelUpdateMq) Consume(key, val string) error {
	var msg kqueue.RelUpdateMqMsg
	if err := json.Unmarshal([]byte(val), &msg); err != nil {
		m.Logger.Error("NewRelUpdateMq->Consume Unmarshal failed, val: %s, err: %s", val, err)
		return err
	}
	if err := m.handle(msg); err != nil {
		m.Logger.Error("NewRelUpdateMq->handle failed, msg: %+v, err: %s", msg, err)
		return err
	}
	return nil
}

func (m *RelUpdateMq) handle(msg kqueue.RelUpdateMqMsg) error {
	switch msg.Table {
	case kqueue.RelUpdateTableGroup:
		if err := m.deleteUserGroupByGroupId(msg.Id); err != nil {
			return errors.Wrap(err, "组-用户关联关系删除失败")
		}
		m.Logger.Infof("删除'组-用户'关联关系成功, group_id: %d", msg.Id)
	case kqueue.RelUpdateTableRole:
		if err := m.deleteUserRoleByRoleId(msg.Id); err != nil {
			return errors.Wrap(err, "角色-用户关联关系删除失败")
		}
		m.Logger.Infof("删除'角色-用户'关联关系成功, role_id: %d", msg.Id)
	default:
		m.Logger.Errorf("异常消息: %+v", msg)
	}
	return nil
}

// 删除组-用户关联关系
func (m *RelUpdateMq) deleteUserGroupByGroupId(groupId int64) error {
	var (
		err   error
		group *uamgroup.UamGroup
	)
	group, err = m.svcCtx.GroupModel.FindOne(m.ctx, groupId)
	if err != nil && err != model.ErrNotFound {
		return errors.Wrap(err, constants.MsgDBErr)
	}
	if group != nil {
		return errors.New(fmt.Sprintf("组尚未删除,无法移除关联关系, group_id: %d", groupId))
	}
	if err = m.svcCtx.RelModel.RemoveUserGroupByGroupId(m.ctx, m.svcCtx.DB, groupId); err != nil {
		return errors.Wrap(err, constants.MsgDBErr)
	}
	return nil
}

// 删除角色-用户关联关系
func (m *RelUpdateMq) deleteUserRoleByRoleId(roleId int64) error {
	var (
		err  error
		role *uamrole.UamRole
	)
	role, err = m.svcCtx.RoleModel.FindOne(m.ctx, roleId)
	if err != nil && err != model.ErrNotFound {
		return errors.Wrap(err, constants.MsgDBErr)
	}
	if role != nil {
		return errors.New(fmt.Sprintf("角色尚未删除,无法移除关联关系, role_id: %d", roleId))
	}
	if err = m.svcCtx.RelModel.RemoveUserRoleByRoleId(m.ctx, m.svcCtx.DB, roleId); err != nil {
		return errors.Wrap(err, constants.MsgDBErr)
	}
	return nil
}
