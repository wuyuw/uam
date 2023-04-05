package core

import (
	"net/http"

	"uam/services/admin/internal/logic/core"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			response.FailByArgsErr(w)
			return
		}

		l := core.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
		} else {
			response.OkWithData(w, resp)
		}
	}
}
