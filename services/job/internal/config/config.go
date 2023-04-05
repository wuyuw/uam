package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf
	Mysql       MysqlConfig
	Cron        Cron
	RelUpdateMq kq.KqConf
}
