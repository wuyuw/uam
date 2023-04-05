package logic

import (
	"context"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermTypeOptionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermTypeOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermTypeOptionsLogic {
	return &GetPermTypeOptionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取权限类型可选项
func (l *GetPermTypeOptionsLogic) GetPermTypeOptions(in *uamrpc.GetPermTypeOptionsReq) (resp *uamrpc.GetPermTypeOptionsResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	typeOptions, err := l.svcCtx.PermissionModel.FindTypeOptions(l.ctx, in.ClientId)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.GetPermTypeOptionsResp{
		List: typeOptions,
	}, nil
}
