package core

import (
	"github.com/golang-jwt/jwt/v4/request"
	"net/http"
	"uam/services/admin/internal/types"

	"uam/services/admin/internal/consts"
	"uam/services/admin/internal/logic/core"
	"uam/services/admin/internal/svc"
	"uam/tools/response"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := request.AuthorizationHeaderExtractor.ExtractToken(r)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
			return
		}
		req := types.LogoutReq{
			Token: tokenString,
		}
		l := core.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
			return
		}
		// 删除cookie中的token(jwt token在过期前一直有效，这里只是君子协定)
		accessCookie := &http.Cookie{
			Name:     consts.CookieAccessToken,
			Value:    "",
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, accessCookie)
		// http.Redirect(w, r, resp.RedirectUrl, http.StatusSeeOther)
		response.OkWithData(w, resp)
	}
}
