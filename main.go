package main

import (
	"SummerProject/router"
	"SummerProject/utils"
)

func main() {
	//连接Redis
	//dao.InitRedis()

	//连接MySQL
	utils.InitMySQL()

	//注册路由
	router.SetUpRouter()

}
