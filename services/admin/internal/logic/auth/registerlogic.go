package auth

import (
	"context"
	"github.com/pkg/errors"
	"uam/services/rpc/pb/uamrpc"
	"uam/tools/cryptox"
	"uam/tools/errx"
	"uam/tools/randstr"

	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	if req.Password != req.DupPassword {
		return &types.RegisterResp{}, errors.New("两次输入密码不一致")
	}
	authLocal, err := l.svcCtx.UamRpc.GetAuthLocalByUsername(l.ctx, &uamrpc.GetAuthLocalByUsernameReq{
		Username: req.Username,
	})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	if authLocal.AuthLocal != nil {
		return &types.RegisterResp{}, errors.New("用户名已存在")
	}
	// 创建用户
	salt := randstr.RandStr(4)
	password := l.genPasswordHash(req.Password, salt)
	if _, err = l.svcCtx.UamRpc.AddAuthLocal(l.ctx, &uamrpc.AddAuthLocalReq{
		Username: req.Username,
		Password: password,
		Salt:     salt,
	}); err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	return &types.RegisterResp{}, nil
}

// genPasswordHash
func (l *RegisterLogic) genPasswordHash(password, salt string) string {
	return cryptox.GetMd5(cryptox.GetMd5(password), salt)
}
