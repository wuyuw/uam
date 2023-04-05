package user

import (
	"context"
	"errors"
	"uam/services/rpc/pb/uamrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/api/internal/consts"
	"uam/services/api/internal/svc"
	"uam/services/api/internal/types"
	"uam/tools/contextx"
	"uam/tools/errx"
)

type AddUserGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserGroupLogic {
	return &AddUserGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 用户入组
func (l *AddUserGroupLogic) AddUserGroup(req *types.UserGroupReq) (resp *types.UserGroupResp, err error) {
	clientId, ok := l.ctx.Value(contextx.CtxKey(consts.CtxKeyClientId)).(int64)
	if !ok {
		return nil, errors.New("异常的客户端")
	}
	_, err = l.svcCtx.UamRpc.AddUserGroup(l.ctx, &uamrpc.AddUserGroupReq{
		Uid:      req.Uid,
		ClientId: clientId,
		GroupId:  req.GroupId,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return
}
