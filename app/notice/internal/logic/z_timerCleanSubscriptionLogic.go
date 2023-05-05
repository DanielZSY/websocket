package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"imservice/app/notice/internal/svc"
)

type TimerCleanSubscriptionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTimerCleanSubscriptionLogic(svcCtx *svc.ServiceContext) *TimerCleanSubscriptionLogic {
	l := &TimerCleanSubscriptionLogic{svcCtx: svcCtx, ctx: context.Background()}
	l.Logger = logx.WithContext(l.ctx)
	return l
}

func (l *TimerCleanSubscriptionLogic) Start() {

}
