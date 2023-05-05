package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	msgservice "imservice/app/msg/msgService"
	"imservice/common/xorm"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql        xorm.MysqlConfig
	ImRpc        zrpc.RpcClientConf
	MsgRpc       zrpc.RpcClientConf
	RelationRpc  zrpc.RpcClientConf
	UserRpc      zrpc.RpcClientConf
	GroupRpc     zrpc.RpcClientConf
	NoticeRpc    zrpc.RpcClientConf
	MsgRpcPod    msgservice.MsgPodsConfig
	Ip2RegionUrl string
}
