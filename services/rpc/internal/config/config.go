package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql       MysqlConfig
	Client      ClientConfig
	IdGen       IdGenConfig
	RelUpdateMq struct {
		Brokers []string // kafka brokers
		Topic   string   // topic
	}
}
