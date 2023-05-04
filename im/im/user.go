package im

import (
	"fmt"
	"imservice/im/benchmark"
)
import log "github.com/sirupsen/logrus"
import "github.com/gomodule/redigo/redis"

func GetSyncKey(appid int64, uid int64) int64 {
	conn := benchmark.redis_pool.Get()
	defer conn.Close()

	key := fmt.Sprintf("users_%d_%d", appid, uid)

	origin, err := redis.Int64(conn.Do("HGET", key, "sync_key"))
	if err != nil && err != redis.ErrNil {
		log.Info("hget error:", err)
		return 0
	}
	return origin
}

func GetGroupSyncKey(appid int64, uid int64, group_id int64) int64 {
	conn := benchmark.redis_pool.Get()
	defer conn.Close()

	key := fmt.Sprintf("users_%d_%d", appid, uid)
	field := fmt.Sprintf("group_sync_key_%d", group_id)

	origin, err := redis.Int64(conn.Do("HGET", key, field))
	if err != nil && err != redis.ErrNil {
		log.Info("hget error:", err)
		return 0
	}
	return origin
}

func SaveSyncKey(appid int64, uid int64, sync_key int64) {
	conn := benchmark.redis_pool.Get()
	defer conn.Close()

	key := fmt.Sprintf("users_%d_%d", appid, uid)

	_, err := conn.Do("HSET", key, "sync_key", sync_key)
	if err != nil {
		log.Warning("hset error:", err)
	}
}

func SaveGroupSyncKey(appid int64, uid int64, group_id int64, sync_key int64) {
	conn := benchmark.redis_pool.Get()
	defer conn.Close()

	key := fmt.Sprintf("users_%d_%d", appid, uid)
	field := fmt.Sprintf("group_sync_key_%d", group_id)

	_, err := conn.Do("HSET", key, field, sync_key)
	if err != nil {
		log.Warning("hset error:", err)
	}
}

func GetUserPreferences(appid int64, uid int64) (int, bool, error) {
	conn := benchmark.redis_pool.Get()
	defer conn.Close()

	key := fmt.Sprintf("users_%d_%d", appid, uid)

	reply, err := redis.Values(conn.Do("HMGET", key, "forbidden", "notification_on"))
	if err != nil {
		log.Info("hget error:", err)
		return 0, false, err
	}

	//电脑在线，手机新消息通知
	var notification_on int
	//用户禁言
	var forbidden int
	_, err = redis.Scan(reply, &forbidden, &notification_on)
	if err != nil {
		log.Warning("scan error:", err)
		return 0, false, err
	}

	return forbidden, notification_on != 0, nil
}

func SetUserUnreadCount(appid int64, uid int64, count int32) {
	conn := benchmark.redis_pool.Get()
	defer conn.Close()

	key := fmt.Sprintf("users_%d_%d", appid, uid)
	_, err := conn.Do("HSET", key, "unread", count)
	if err != nil {
		log.Info("hset err:", err)
	}
}
