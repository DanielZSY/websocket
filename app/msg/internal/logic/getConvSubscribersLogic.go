package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"imservice/common/utils"
	"imservice/common/xredis/rediskey"
	"time"

	"imservice/app/msg/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConvSubscribersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConvSubscribersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConvSubscribersLogic {
	return &GetConvSubscribersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetConvSubscribers 获取一个会话里所有的消息订阅者
func (l *GetConvSubscribersLogic) GetConvSubscribers(in *pb.GetConvSubscribersReq) (*pb.GetConvSubscribersResp, error) {
	// ZRANGEBYSCORE conv:subscribers:group:1 min +inf
	min := time.Now().AddDate(0, 0, -1).UnixMilli()
	if in.LastActiveTime != nil {
		min = *in.LastActiveTime
	}
	val, err := l.svcCtx.Redis().ZrangebyscoreWithScoresCtx(l.ctx, rediskey.ConvMembersSubscribed(in.ConvId), min, time.Now().UnixMilli()+1000*60*60)
	if err != nil {
		if err == redis.Nil {
			return &pb.GetConvSubscribersResp{}, nil
		}
		l.Errorf("get conv subscribers error: %v", err)
		return &pb.GetConvSubscribersResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	userIds := make([]string, 0)
	for _, pair := range val {
		userId := rediskey.ConvMembersSubscribedSplit(pair.Key)
		userIds = append(userIds, userId)
	}
	return &pb.GetConvSubscribersResp{
		UserIdList: utils.Set(userIds),
	}, nil
}
