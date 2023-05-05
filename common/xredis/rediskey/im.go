package rediskey

import "imservice/common/xredis"

func ConvSetting(convId string, userId string) string {
	return "s:model:conv_setting:" + convId + ":" + userId
}

func ConvSettingExpire() int {
	return xredis.ExpireMinutes(5)
}
