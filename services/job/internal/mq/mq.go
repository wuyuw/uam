package mq

import (
	"context"
	"uam/services/job/internal/mq/kqhandler"
	"uam/services/job/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func NewKqServices(svcCtx *svc.ServiceContext) []service.Service {
	ctx := context.Background()
	return []service.Service{
		kq.MustNewQueue(svcCtx.Config.RelUpdateMq, kqhandler.NewRelUpdateMq(ctx, svcCtx)),
	}
}
