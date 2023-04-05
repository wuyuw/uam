package middleware

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"
	"uam/services/api/internal/consts"
	"uam/services/rpc/pb/uamrpc"
	"uam/services/rpc/uam"
	"uam/tools/contextx"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const CacheKeyClient = "client:appcode:"

type (
	Client struct {
		Id         int64
		Name       string
		AppCode    string
		PrivateKey string
		Status     int64
	}
	ApiCommonQueryParams struct {
		AppCode     string `form:"appCode"`
		RequestTime int64  `form:"requestTime"`
		Sign        string `form:"sign"`
	}
	ApiCommonBodyParams struct {
		AppCode     string `json:"appCode"`
		RequestTime int64  `json:"requestTime"`
		Sign        string `json:"sign"`
	}
	ApiCommonParams struct {
		AppCode     string
		RequestTime int64
		Sign        string
	}
)

type ApiAuthMiddleware struct {
}

func NewApiAuthMiddleware() *ApiAuthMiddleware {
	return &ApiAuthMiddleware{}
}

func (m *ApiAuthMiddleware) HandleWrap(cache *collection.Cache, uamRpc uam.Uam) rest.Middleware {

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// copy request.Body
			reqBody := []byte{}
			if r.Body != nil {
				reqBody, _ = ioutil.ReadAll(r.Body)
			}
			cpReqBody := make([]byte, len(reqBody))
			copy(cpReqBody, reqBody)
			// 读取后重置Body
			r.Body = ioutil.NopCloser(bytes.NewBuffer(cpReqBody))
			var params ApiCommonParams
			if r.Method == http.MethodGet {
				var queryParams ApiCommonQueryParams
				if err := httpx.Parse(r, &queryParams); err != nil {
					response.FailByArgsErr(w, err.Error())
					return
				}
				params = ApiCommonParams(queryParams)
			} else {
				var bodyParams ApiCommonBodyParams
				if err := httpx.Parse(r, &bodyParams); err != nil {
					response.FailByArgsErr(w, err.Error())
					return
				}
				params = ApiCommonParams(bodyParams)
			}
			// 再次重置Body
			r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
			// 获取缓存，如果 client 不存在的，则会调用 func 去生成缓存
			clientIface, err := cache.Take(CacheKeyClient+params.AppCode, func() (interface{}, error) {
				return m.getClientByCode(r.Context(), uamRpc, params.AppCode)
			})
			if err != nil {
				response.FailBySvcErr(w, err.Error())
				return
			}
			client := clientIface.(*Client)
			if client == nil {
				response.FailByInvalidClient(w)
				return
			}
			if ok := m.checkTime(params.RequestTime); !ok {
				response.FailByAuthFail(w)
				return
			}
			if ok := m.checkSign(client, params); !ok {
				response.FailByAuthFail(w)
				return
			}
			// 客户端被禁用
			if client.Status != 0 {
				response.FailByForbidden(w)
				return
			}
			// 将ClientId添加至ctx
			ctx := context.WithValue(r.Context(), contextx.CtxKey(consts.CtxKeyClientId), client.Id)
			next(w, r.WithContext(ctx))
		}
	}
}

func (m *ApiAuthMiddleware) getClientByCode(ctx context.Context, uamRpc uam.Uam, appCode string) (*Client, error) {
	rpcResp, err := uamRpc.GetClientByCode(ctx, &uamrpc.GetClientByCodeReq{
		AppCode: appCode,
	})
	if err != nil {
		return nil, err
	}
	if rpcResp.Client == nil {
		return nil, nil
	}
	return &Client{
		Id:         rpcResp.Client.Id,
		Name:       rpcResp.Client.Name,
		AppCode:    rpcResp.Client.AppCode,
		PrivateKey: rpcResp.Client.PrivateKey,
		Status:     rpcResp.Client.Status,
	}, nil
}

func (m *ApiAuthMiddleware) checkTime(requestTime int64) bool {
	delta := math.Abs(float64(time.Now().Unix() - requestTime))
	return delta < 3000
}

func (m *ApiAuthMiddleware) checkSign(client *Client, params ApiCommonParams) bool {
	raw := fmt.Sprintf("appCode=%s&privateKey=%s&requestTime=%d", client.AppCode, client.PrivateKey, params.RequestTime)
	return fmt.Sprintf("%x", md5.Sum([]byte(raw))) == params.Sign
}
