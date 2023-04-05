package auth

import (
	"net/http"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"uam/services/admin/internal/logic/auth"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			response.FailByArgsErr(w)
			return
		}
		l := auth.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			if resp == nil {
				response.FailBySvcErr(w, err.Error())
			} else {
				response.FailWithMsg(w, err.Error())
			}
			return
		}
		response.OkWithData(w, resp)
	}
}
