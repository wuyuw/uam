package auth

import (
	"context"
	"github.com/pkg/errors"
	"time"

	"uam/services/rpc/pb/uamrpc"
	"uam/tools/cryptox"
	"uam/tools/errx"
	"uam/tools/jwtx"

	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrLogin = errors.New("用户名或密码错误")

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	authLocal, err := l.svcCtx.UamRpc.GetAuthLocalByUsername(l.ctx,
		&uamrpc.GetAuthLocalByUsernameReq{Username: req.Username})
	if err != nil {
		return nil, errx.ConverRpcErr(err)
	}
	// 认证记录不存在
	if authLocal.AuthLocal == nil {
		return &types.LoginResp{}, ErrLogin
	}
	// 密码校验
	if l.genPasswordHash(req.Password, authLocal.AuthLocal.Salt) != authLocal.AuthLocal.Password {
		return &types.LoginResp{}, ErrLogin
	}
	// 生成JWT Token
	accessToken, err := l.generateToken(int(authLocal.AuthLocal.Uid))
	if err != nil {
		return nil, err
	}
	// 设置redis缓存
	if err = l.svcCtx.Redis.Setex(jwtx.JWTRedisPrefix+accessToken,
		accessToken, int(l.svcCtx.Config.JwtAuth.AccessExpire*2)); err != nil {
		return nil, err
	}
	return &types.LoginResp{
		AccessToken: accessToken,
	}, nil
}

// genPasswordHash
func (l *LoginLogic) genPasswordHash(password, salt string) string {
	return cryptox.GetMd5(cryptox.GetMd5(password), salt)
}

func (l *LoginLogic) generateToken(uid int) (string, error) {
	now := time.Now().Unix()
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire

	accessToken, err := jwtx.GetJwtToken(accessSecret, now, accessExpire, map[string]interface{}{
		"uid": uid,
	})
	if err != nil {
		return "", errors.Wrap(err, "JwtToken生成失败")
	}
	return accessToken, nil
}
