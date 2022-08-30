package main

import (
	"SummerProject/routers"
	"SummerProject/utils"
)

func main() {
	//连接Redis
	//utils.InitRedis()

	//连接MySQL
	utils.InitMySQL()

	//雪花算法优化版
	utils.InitID()

	//注册路由
	routers.InitRouter()
}
