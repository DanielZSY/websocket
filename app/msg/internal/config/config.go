package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"imservice/common/pkg/mobpush"
	"imservice/common/xorm"
	"imservice/common/xtdmq"
)

type Config struct {
	zrpc.RpcServerConf
	TDMQ struct {
		Enabled bool `json:",default=true"`
		xtdmq.TDMQConfig
		xtdmq.TDMQConsumerConfig
		Producer                xtdmq.TDMQProducerConfig
		SendMsgListTaskInterval int64 `json:",default=40"`
		SendMsgListTaskNum      int   `json:",default=100"`
	}
	Mysql            xorm.MysqlConfig
	ImRpc            zrpc.RpcClientConf
	RelationRpc      zrpc.RpcClientConf
	GroupRpc         zrpc.RpcClientConf
	UserRpc          zrpc.RpcClientConf
	NoticeRpc        zrpc.RpcClientConf
	MobPush          mobpush.Config
	MobAlias         string `json:",default=deviceId,options=deviceId|userId"`
	SyncSendMsgLimit struct {
		Rate  int `json:",default=50"`  // 每秒生成
		Burst int `json:",default=100"` // 100 令牌桶最大值
	}
}
