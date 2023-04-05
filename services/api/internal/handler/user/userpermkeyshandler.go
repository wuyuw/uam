package user

import (
	"net/http"

	"uam/services/api/internal/logic/user"
	"uam/services/api/internal/svc"
	"uam/services/api/internal/types"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserPermKeysHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPermKeysReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.WithContext(r.Context()).Error(err)
			response.FailByArgsErr(w)
			return
		}

		l := user.NewUserPermKeysLogic(r.Context(), svcCtx)
		resp, err := l.UserPermKeys(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
		} else {
			response.OkWithData(w, resp.List)
		}
	}
}
