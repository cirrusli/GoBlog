package common

import (
	"github.com/go-redis/redis"
	"log"
)

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := RDB.Ping().Result()
	if err != nil {
		log.Panicln("connect Redis failed:", err.Error())
	}

}
