package config

import (
	"SummerProject/internal/model"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Conf todo better not to use global variable
var Conf *model.ConfStruct

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("read config.toml failed:", err.Error())
	}
	fmt.Println("I can print this ", viper.AllSettings())

	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalln("unmarshal failed:", err.Error())
	}
	log.Println(viper.GetString("Redis.FollowingList"))
	log.Println(Conf)
}
