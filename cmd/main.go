package main

import (
	"SummerProject/common"
	"SummerProject/config"
	"SummerProject/internal/router"
)

func main() {
	//初始化配置项
	config.InitConfig()

	//连接Redis
	//redis.InitRedis()

	//连接MySQL
	common.InitMySQL()

	//雪花算法优化版
	common.InitID()

	//启动RabbitMQ
	//rabbitmq.SendInit()
	//rabbitmq.ReceiveInit()

	//注册路由
	router.InitRouter()
}
