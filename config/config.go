package config

import (
	"GoBlog/internal/model"
	"fmt"
	"github.com/spf13/viper" //Viper在后台使用github.com/mitchellh/mapstructure来解析值，其默认情况下使用mapstructure tag。
	"log"
)

// Conf  better not to use global variable
var Conf *model.ConfStruct

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed:", err)
	}

	fmt.Println("read configs succeed:", viper.AllSettings())

	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatal("unmarshal config failed:", err)
	}

	log.Println(viper.GetString("Redis.FollowingList"))
	log.Println(Conf)
}
