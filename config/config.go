package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type MySQL struct {
	UserName     string `toml:"UserName"`
	Password     string `toml:"Password"`
	Addr         string `toml:"Addr"`
	DatabaseName string `toml:"DatabaseName"`
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("read config.toml failed:", err.Error())
	}
	fmt.Println("I can print this ", viper.AllSettings())

	var mysql MySQL
	err = viper.Unmarshal(&mysql)
	if err != nil {
		log.Fatalln("unmarshal failed:", err.Error())
	}
	log.Println(viper.GetString("MySQL.PassWord"))

}
func MySQLConfig() MySQL {
	var mysql MySQL
	return mysql
}
