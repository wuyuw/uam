package svc

import (
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"time"
	"uam/services/api/internal/config"
	"uam/services/api/internal/middleware"
	"uam/services/rpc/uam"
)

type ServiceContext struct {
	Config  config.Config
	Redis   *redis.Redis
	Cache   *collection.Cache
	UamRpc  uam.Uam
	ApiAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	uamRpc := uam.NewUam(zrpc.MustNewClient(c.UamRpc))
	cache := MustNewCache()
	return &ServiceContext{
		Config:  c,
		Cache:   cache,
		Redis:   MustInitRedis(c),
		UamRpc:  uamRpc,
		ApiAuth: middleware.NewApiAuthMiddleware().HandleWrap(cache, uamRpc),
	}
}

func MustInitRedis(c config.Config) *redis.Redis {
	client := c.Redis.NewRedis()
	if ok := client.Ping(); !ok {
		log.Fatal("redis连接失败!")
	}
	return client
}

func MustNewCache() *collection.Cache {
	cache, err := collection.NewCache(time.Minute, collection.WithLimit(1000))
	if err != nil {
		log.Fatalf("缓存初始化失败: %s", err)
	}
	return cache
}
