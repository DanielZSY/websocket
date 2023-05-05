package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"imservice/common/xorm"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql        xorm.MysqlConfig
	ImRpc        zrpc.RpcClientConf
	UserRpc      zrpc.RpcClientConf
	MsgRpc       zrpc.RpcClientConf
	NoticeRpc    zrpc.RpcClientConf
	Ip2RegionUrl string `json:",default=https://github.com/lionsoul2014/ip2region/raw/master/data/ip2region.xdb"`
}