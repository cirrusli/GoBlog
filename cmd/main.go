package main

import (
	"GoBlog/common"
	"GoBlog/config"
	"GoBlog/internal/router"
	"fmt"
)

func init() {

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

}
func main() {
	fmt.Println("初始化成功！")

	router.InitRouter()
}
