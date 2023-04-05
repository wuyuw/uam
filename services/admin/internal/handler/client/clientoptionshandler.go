package client

import (
	"net/http"

	"uam/services/admin/internal/logic/client"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ClientOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClientOptionsReq
		if err := httpx.Parse(r, &req); err != nil {
			response.FailByArgsErr(w)
			return
		}

		l := client.NewClientOptionsLogic(r.Context(), svcCtx)
		resp, err := l.ClientOptions(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
		} else {
			response.OkWithData(w, resp.List)
		}
	}
}
