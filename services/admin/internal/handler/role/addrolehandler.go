package role

import (
	"net/http"

	"uam/services/admin/internal/logic/role"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			response.FailByArgsErr(w)
			return
		}

		l := role.NewAddRoleLogic(r.Context(), svcCtx)
		_, err := l.AddRole(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
		} else {
			response.Ok(w)
		}
	}
}
