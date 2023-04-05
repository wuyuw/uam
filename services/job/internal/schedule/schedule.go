package schedule

import (
	"uam/services/job/internal/schedule/syncoauser"
	"uam/services/job/internal/svc"

	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type CronService struct {
	cron   *cron.Cron
	svcCtx *svc.ServiceContext
}

func (c *CronService) Start() {
	logx.Info("Starting scheduling cron jobs...")
	c.cron.Start()
}

func (c *CronService) Stop() {
	c.cron.Stop()
}

// 注册tasks
func (c *CronService) registerTasks() {
	_, err := c.cron.AddFunc(c.svcCtx.Config.Cron.SyncOaUser, syncoauser.SyncOaUser(c.svcCtx))
	if err != nil {

	}
}

func NewCronSchedule(svcCtx *svc.ServiceContext) service.Service {
	cronSchedule := &CronService{
		cron:   cron.New(),
		svcCtx: svcCtx,
	}
	cronSchedule.registerTasks()
	return cronSchedule
}
