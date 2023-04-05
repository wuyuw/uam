package auth

import (
	"net/http"
	"uam/services/admin/internal/consts"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"uam/services/admin/internal/logic/auth"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			response.FailByArgsErr(w)
			return
		}
		// 登录校验
		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			if resp == nil {
				response.FailBySvcErr(w, err.Error())
			} else {
				response.FailWithMsg(w, err.Error())
			}
			return
		}
		// 设置cookie
		accessCookie := &http.Cookie{
			Name:     consts.CookieAccessToken,
			Value:    resp.AccessToken,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, accessCookie)
		response.Ok(w)
		//indexUrl := "/"
		//// 开发环境没有nginx做反代，需要指定完整url
		//if svcCtx.Config.System.Env == constants.EnvDev {
		//	indexUrl = svcCtx.Config.System.FeIndex
		//}
		//http.Redirect(w, r, indexUrl, http.StatusFound)
	}
}
