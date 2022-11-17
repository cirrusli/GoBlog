package redis

import (
	"SummerProject/common"
	"github.com/go-redis/redis"
)

var RDB *redis.Client

func InitRedis() {
	//todo 关注，粉丝，互关分别是三个数据库
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := RDB.Ping().Result()
	if err != nil {
		common.PanicLog("connect Redis failed:", err)
	}

}
