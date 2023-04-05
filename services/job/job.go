package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"uam/services/job/internal/config"
	"uam/services/job/internal/mq"
	"uam/services/job/internal/schedule"
	"uam/services/job/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/uam-job.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		log.Fatal(err)
	}
	ctx := svc.NewServiceContext(c)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	// 添加Cron调度服务
	serviceGroup.Add(schedule.NewCronSchedule(ctx))
	logx.Info("cron调度任务初始化完成.")
	// 添加MQ服务
	for _, kqService := range mq.NewKqServices(ctx) {
		serviceGroup.Add(kqService)
	}
	logx.Info("KQ消息队列初始化完成.")

	serviceGroup.Start()
	//捕捉信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-ch
		logx.Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			logx.Infof("stop service")
			serviceGroup.Stop()
			logx.Info("service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
