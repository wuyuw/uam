package logic

import (
	"context"
	"github.com/pkg/errors"
	"uam/services/model"
	"uam/tools/constants"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthLocalByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuthLocalByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthLocalByUsernameLogic {
	return &GetAuthLocalByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetAuthLocalByUsername 根据用户名查询本地认证记录
func (l *GetAuthLocalByUsernameLogic) GetAuthLocalByUsername(in *uamrpc.GetAuthLocalByUsernameReq) (resp *uamrpc.GetAuthLocalByUsernameResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	authLocal, err := l.svcCtx.AuthLocalModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			return &uamrpc.GetAuthLocalByUsernameResp{AuthLocal: nil}, nil
		}
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}

	return &uamrpc.GetAuthLocalByUsernameResp{
		AuthLocal: &uamrpc.AuthLocal{
			Uid:      authLocal.Uid,
			Username: authLocal.Username,
			Password: authLocal.Password,
			Salt:     authLocal.Salt,
		},
	}, nil
}
