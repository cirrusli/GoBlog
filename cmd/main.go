package main

import (
	"SummerProject/router"
	"SummerProject/utils"
)

func main() {
	//连接Redis
	//utils.InitRedis()

	//连接MySQL
	utils.InitMySQL()

	//雪花算法
	utils.InitSnow()

	//注册路由
	router.SetUpRouter()
}
