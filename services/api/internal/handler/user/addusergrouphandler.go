package user

import (
	"net/http"

	"uam/services/api/internal/logic/user"
	"uam/services/api/internal/svc"
	"uam/services/api/internal/types"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户入组
func AddUserGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserGroupReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			response.FailByArgsErr(w, err.Error())
			return
		}
		l := user.NewAddUserGroupLogic(r.Context(), svcCtx)
		_, err := l.AddUserGroup(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
		} else {
			response.Ok(w)
		}
	}
}
