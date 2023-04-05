package main

import (
	"flag"
	"fmt"

	"uam/services/admin/internal/config"
	"uam/services/admin/internal/handler"
	"uam/services/admin/internal/setup"
	"uam/services/admin/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/uam-admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	setup.MustSyncSysPermission(ctx)
	setup.MustBuildRouterTree(ctx)
	setup.MustSyncApiPermission(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
