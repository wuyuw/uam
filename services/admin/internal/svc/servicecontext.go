package svc

import (
	"log"
	"uam/services/admin/internal/config"
	"uam/services/admin/internal/middleware"
	"uam/services/admin/internal/sysadmin"
	uamclient "uam/services/model/uam_client"
	"uam/services/rpc/uam"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	Redis  *redis.Redis
	UamRpc uam.Uam

	SysClient  *uamclient.UamClient
	RouterTree *sysadmin.RouterTree
	ApiPerms   map[string]int64

	AccessControl rest.Middleware
	JwtAuth       rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	uamRpc := uam.NewUam(zrpc.MustNewClient(c.UamRpc))
	redisClient := MustInitRedis(c)
	routerTree := sysadmin.NewRouterTree()
	sysClient := uamclient.UamClient{}
	apiPerms := make(map[string]int64)
	return &ServiceContext{
		Config:        c,
		Redis:         redisClient,
		UamRpc:        uamRpc,
		SysClient:     &sysClient,
		RouterTree:    routerTree,
		ApiPerms:      apiPerms,
		AccessControl: middleware.NewAccessControlMiddleware().HandleWrap(c, routerTree, &sysClient, apiPerms, uamRpc),
		JwtAuth:       middleware.NewJwtAuthMiddleware().HandleWrap(c, redisClient),
	}
}

func MustInitRedis(c config.Config) *redis.Redis {
	client := c.Redis.NewRedis()
	if ok := client.Ping(); !ok {
		log.Fatal("redis连接失败!")
	}
	return client
}
