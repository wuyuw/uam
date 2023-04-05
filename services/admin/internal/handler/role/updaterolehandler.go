package role

import (
	"net/http"

	"uam/services/admin/internal/logic/role"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			response.FailByArgsErr(w)
			return
		}

		l := role.NewUpdateRoleLogic(r.Context(), svcCtx)
		_, err := l.UpdateRole(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
		} else {
			response.Ok(w)
		}
	}
}
