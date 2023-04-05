package middleware

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"uam/services/admin/internal/config"
	"uam/services/admin/internal/consts"
	"uam/tools/constants"
	"uam/tools/jwtx"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest"
)

type JwtAuthMiddleware struct {
}

func NewJwtAuthMiddleware() *JwtAuthMiddleware {
	return &JwtAuthMiddleware{}
}

func (m *JwtAuthMiddleware) HandleWrap(c config.Config, redis *redis.Redis) rest.Middleware {

	//var (
	//	ignoreAuth bool
	//	mockClaims jwt.MapClaims
	//)
	// 开发环境忽略登录认证，mock用户信息
	//if c.System.Env == constants.EnvDev {
	//	ignoreAuth = true
	//	mockClaims = make(jwt.MapClaims)
	//	mockClaims["uid"] = json.Number("7841")
	//}
	handler := jwtx.Authorize(
		c.JwtAuth.AccessSecret,
		c.JwtAuth.AccessExpire,
		redis,
		// 认证失败回调
		jwtx.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
			if err == jwtx.ErrSvcError {
				response.FailBySvcErr(w, err.Error())
			}
			response.FailByNoLogin(w, map[string]string{
				"error": err.Error(),
			})
		}),
		//jwtx.WithIgnoreCallback(func() bool {
		//	// 开发环境不做登录认证
		//	return ignoreAuth
		//}),
		//jwtx.WithMockClaims(mockClaims),
	)
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 开发环境从cookie中获取token并设置Authorization
			if c.System.Env == constants.EnvDev {
				accessToken, err := r.Cookie(consts.CookieAccessToken)
				if err != nil {
					fmt.Println(err)
				}
				if accessToken != nil {
					r.Header.Set("Authorization", "Bearer "+accessToken.Value)
				}
			}
			handler(next).ServeHTTP(w, r)
		}
	}
}
