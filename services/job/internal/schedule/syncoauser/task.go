package syncoauser

import (
	"github.com/zeromicro/go-zero/core/logx"
	"uam/services/job/internal/svc"
	"uam/services/job/internal/types"
)

func SyncOaUser(svcCtx *svc.ServiceContext) types.CronTask {
	return func() {
		logx.Info("starting sync oa users...")
		defer logx.Info("sync oa users completed!")
	}
}
