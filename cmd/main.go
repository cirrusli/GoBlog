package main

import (
	"SummerProject/common"
	"SummerProject/internal/router"
)

func main() {
	//连接Redis
	//common.InitRedis()

	//连接MySQL
	common.InitMySQL()

	//雪花算法优化版
	common.InitID()

	//注册路由
	router.InitRouter()
}
