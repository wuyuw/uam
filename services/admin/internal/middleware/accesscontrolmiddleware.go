package middleware

import (
	"fmt"
	"net/http"
	"uam/services/admin/internal/config"
	"uam/services/admin/internal/sysadmin"
	uamclient "uam/services/model/uam_client"
	"uam/services/rpc/uam"
	"uam/tools/jwtx"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type AccessControlMiddleware struct {
}

func NewAccessControlMiddleware() *AccessControlMiddleware {
	return &AccessControlMiddleware{}
}

func (m *AccessControlMiddleware) HandleWrap(
	c config.Config,
	routerTree *sysadmin.RouterTree,
	sysClient *uamclient.UamClient,
	apiPerms map[string]int64,
	uamRpc uam.Uam) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 访问权限控制
			// 从ctx中获取uid
			uid, err := jwtx.GetCtxValueInt64(r.Context(), "uid")
			if err != nil {
				logx.Error(err)
				response.FailByForbidden(w, "异常的用户信息")
				return
			}
			if len(apiPerms) == 0 {
				response.FailBySvcErr(w, "系统权限为初始化")
				return
			}
			apiPerm, err := routerTree.Search(r)
			if err != nil {
				response.FailByForbidden(w, "异常的请求")
				return
			}
			apiPermId, ok := apiPerms[fmt.Sprintf("%s %s", apiPerm.Method, apiPerm.Path)]
			if !ok {
				// 不在系统API权限范围内的接口不做鉴权
				next(w, r)
				return
			}
			// 获取用户权限列表
			userPerms, err := uamRpc.GetPermIdsByUid(r.Context(), &uam.GetPermIdsByUidReq{
				ClientId: sysClient.Id,
				Uid:      int64(uid),
			})
			if err != nil {
				response.FailBySvcErr(w, "用户权限获取失败")
				return
			}
			for _, permId := range userPerms.List {
				// 有访问权限
				if apiPermId == permId {
					next(w, r)
					return
				}
			}
			response.FailByForbidden(w, "无操作权限")
		}
	}
}
