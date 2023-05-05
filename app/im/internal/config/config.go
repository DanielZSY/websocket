package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"imservice/app/conn/connservice"
	"imservice/common/xorm"
)

type Config struct {
	zrpc.RpcServerConf
	ConnRpc      connservice.ConnPodsConfig
	Mysql        xorm.MysqlConfig
	Ip2RegionUrl string `json:",default=https://github.com/lionsoul2014/ip2region/raw/master/data/ip2region.xdb"`
	MsgRpc       zrpc.RpcClientConf
	NoticeRpc    zrpc.RpcClientConf
	RelationRpc  zrpc.RpcClientConf
	GroupRpc     zrpc.RpcClientConf
	UserRpc      zrpc.RpcClientConf
}
