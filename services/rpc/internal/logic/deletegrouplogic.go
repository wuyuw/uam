package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"uam/services/model"
	uamgroup "uam/services/model/uam_group"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"
	"uam/tools/kqueue"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除组
func (l *DeleteGroupLogic) DeleteGroup(in *uamrpc.DeleteGroupReq) (resp *uamrpc.DeleteGroupResp, err error) {
	var (
		group *uamgroup.UamGroup
	)
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	group, err = l.svcCtx.GroupModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if group == nil {
		return &uamrpc.DeleteGroupResp{}, nil
	}
	if group.Editable == 0 {
		return nil, errors.New(constants.MsgEditableErr)
	}

	err = l.deleteGroupInTx(in.Id)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.DeleteGroupResp{}, nil
}

func (l *DeleteGroupLogic) deleteGroupInTx(groupId int64) error {
	return model.ExecInTx(l.svcCtx.DB, func(tx *gorm.DB) error {
		// 删除组
		var err error
		if err = l.svcCtx.GroupModel.Delete(l.ctx, tx, groupId); err != nil {
			return err
		}
		// 删除组-角色关联关系
		if err = l.svcCtx.RelModel.RemoveGroupRoleByGroupId(l.ctx, tx, groupId); err != nil {
			return err
		}
		// 删除组-用户关联关系
		// 关联关系删除放在MQ中进行，避免大量记录删除造成性能抖动
		msg := kqueue.RelUpdateMqMsg{
			Table: kqueue.RelUpdateTableGroup,
			Id:    groupId,
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
