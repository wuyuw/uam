package logic

import (
	"context"

	"uam/services/rpc/internal/service/dao"
	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/constants"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermIdsByGroupIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermIdsByGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermIdsByGroupIdLogic {
	return &GetPermIdsByGroupIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取组关联的权限Id列表
func (l *GetPermIdsByGroupIdLogic) GetPermIdsByGroupId(in *uamrpc.GetPermIdsByGroupIdReq) (resp *uamrpc.GetPermIdsByGroupIdResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	permIds, err := dao.FindPermIdsByGroupId(l.ctx, l.svcCtx, in.GroupId)
	if err != nil {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	return &uamrpc.GetPermIdsByGroupIdResp{List: permIds}, nil
}
