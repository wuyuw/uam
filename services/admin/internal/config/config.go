package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Redis  redis.RedisConf
	System struct {
		Env        string // 运行环境: dev/test/prod
		FeIndex    string // 前端首页地址: dev环境存在跨域问题
		ClientName string // 系统客户端名称
		ClientCode string // 系统客户端appCode
		AdminRole  string // 系统管理员角色名
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	UamRpc zrpc.RpcClientConf
}
