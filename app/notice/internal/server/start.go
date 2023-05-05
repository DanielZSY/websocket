package server

import "imservice/app/notice/internal/logic"

func (s *NoticeServiceServer) Start() {
	{
		l := logic.NewTimerCleanSubscriptionLogic(s.svcCtx)
		go l.Start()
	}
}
