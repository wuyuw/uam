package jwtx

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"net/http/httputil"

	"github.com/zeromicro/go-zero/rest/token"
)

type CtxKey string

const (
	jwtAudience    = "aud"
	jwtExpire      = "exp"
	jwtId          = "jti"
	jwtIssueAt     = "iat"
	jwtIssuer      = "iss"
	jwtNotBefore   = "nbf"
	jwtSubject     = "sub"
	noDetailReason = "no detail reason"
)

var (
	errInvalidToken = errors.New("invalid auth token")
	errNoClaims     = errors.New("no auth params")
	ErrSvcError     = errors.New("service error")
)

type (
	// A AuthorizeOptions is authorize options.
	AuthorizeOptions struct {
		PrevSecret     string
		Callback       UnauthorizedCallback
		IgnoreCallback IgnoreAuthCallback
		MockClaims     jwt.MapClaims
	}

	// UnauthorizedCallback defines the method of unauthorized callback.
	UnauthorizedCallback func(w http.ResponseWriter, r *http.Request, err error)

	// IgnoreAuthCallback defines the method whether to ignore jwt auth.
	IgnoreAuthCallback func() bool

	// AuthorizeOption defines the method to customize an AuthorizeOptions.
	AuthorizeOption func(opts *AuthorizeOptions)
)

// Authorize returns an authorization middleware.
func Authorize(secret string, renewalSeconds int64, redis *redis.Redis, opts ...AuthorizeOption) func(http.Handler) http.Handler {
	var authOpts AuthorizeOptions
	for _, opt := range opts {
		opt(&authOpts)
	}

	parser := token.NewTokenParser()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// ignore auth
			if ignoreAuth(authOpts.IgnoreCallback) {
				// mock claims
				ctx := r.Context()
				if authOpts.MockClaims != nil {
					for k, v := range authOpts.MockClaims {
						ctx = context.WithValue(ctx, CtxKey(k), v)
					}
				}
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			tok, err := parser.ParseToken(r, secret, authOpts.PrevSecret)
			if err != nil {
				unauthorized(w, r, err, authOpts.Callback)
				return
			}

			// 判断redis缓存是否存在
			keyExist, err := redis.Exists(JWTRedisPrefix + tok.Raw)
			if err != nil {
				logx.Errorf("redis error: %v", err)
				unauthorized(w, r, ErrSvcError, authOpts.Callback)
				return
			}
			if !tok.Valid {
				// 判断redis缓存是否过期，没过期则自动续期
				if !keyExist {
					unauthorized(w, r, errInvalidToken, authOpts.Callback)
					return
				}
				if err := redis.Expire(JWTRedisPrefix+tok.Raw, int(renewalSeconds)); err != nil {
					logx.Errorf("redis error: %v", err)
					unauthorized(w, r, ErrSvcError, authOpts.Callback)
					return
				}
			}
			// 如果有效，但是redis中不存在，则设置redis缓存
			if !keyExist {
				if err := redis.Setex(JWTRedisPrefix+tok.Raw, tok.Raw, int(renewalSeconds)); err != nil {
					// 即使设置失败，后续只要tok在有效期仍可正常访问，直到过期后，由于redis不存在，则直接认证失败
					logx.Errorf("redis error: %v", err)
				}
			}

			claims, ok := tok.Claims.(jwt.MapClaims)
			if !ok {
				unauthorized(w, r, errNoClaims, authOpts.Callback)
				return
			}

			ctx := r.Context()
			for k, v := range claims {
				switch k {
				case jwtAudience, jwtExpire, jwtId, jwtIssueAt, jwtIssuer, jwtNotBefore, jwtSubject:
					// ignore the standard claims
				default:
					ctx = context.WithValue(ctx, CtxKey(k), v)
				}
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// WithPrevSecret returns an AuthorizeOption with setting previous secret.
func WithPrevSecret(secret string) AuthorizeOption {
	return func(opts *AuthorizeOptions) {
		opts.PrevSecret = secret
	}
}

// WithUnauthorizedCallback returns an AuthorizeOption with setting unauthorized callback.
func WithUnauthorizedCallback(callback UnauthorizedCallback) AuthorizeOption {
	return func(opts *AuthorizeOptions) {
		opts.Callback = callback
	}
}

// WithIgnoreCallback returns an AuthorizeOption with setting ignore jwt auth callback.
func WithIgnoreCallback(callback IgnoreAuthCallback) AuthorizeOption {
	return func(opts *AuthorizeOptions) {
		opts.IgnoreCallback = callback
	}
}

func WithMockClaims(claims jwt.MapClaims) AuthorizeOption {
	return func(opts *AuthorizeOptions) {
		opts.MockClaims = claims
	}
}

func detailAuthLog(r *http.Request, reason string) {
	// discard dump error, only for debug purpose
	details, _ := httputil.DumpRequest(r, true)
	logx.Errorf("authorize failed: %s\n=> %+v", reason, string(details))
}

func unauthorized(w http.ResponseWriter, r *http.Request, err error, callback UnauthorizedCallback) {
	// writer := response.NewHeaderOnceResponseWriter(w)

	if err != nil {
		detailAuthLog(r, err.Error())
	} else {
		detailAuthLog(r, noDetailReason)
	}

	// let callback go first, to make sure we respond with user-defined HTTP header
	if callback != nil {
		callback(w, r, err)
	}

	// if user not setting HTTP header, we set header with 401
	w.WriteHeader(http.StatusUnauthorized)
}

func ignoreAuth(callback IgnoreAuthCallback) bool {
	if callback != nil {
		return callback()
	}
	return false
}
