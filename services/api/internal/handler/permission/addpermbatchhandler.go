package permission

import (
	"net/http"

	"uam/services/api/internal/logic/permission"
	"uam/services/api/internal/svc"
	"uam/services/api/internal/types"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddPermBatchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddPermBatchReq
		if err := httpx.Parse(r, &req); err != nil {
			response.FailByArgsErr(w)
			return
		}

		l := permission.NewAddPermBatchLogic(r.Context(), svcCtx)
		_, err := l.AddPermBatch(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
		} else {
			response.Ok(w)
		}
	}
}
