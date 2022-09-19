package main

import (
	"SummerProject/common"
	"SummerProject/config"
	"SummerProject/internal/middleware/rabbitmq"
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

	//启动rabbitmq服务
	rabbitmq.SendInit()
	rabbitmq.ReceiveInit()

	//注册路由
	router.InitRouter()
}
