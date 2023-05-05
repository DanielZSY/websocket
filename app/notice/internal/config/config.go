package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"imservice/common/xorm"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql       xorm.MysqlConfig
	ImRpc       zrpc.RpcClientConf
	RelationRpc zrpc.RpcClientConf
	GroupRpc    zrpc.RpcClientConf
	UserRpc     zrpc.RpcClientConf
	MsgRpc      zrpc.RpcClientConf
}
