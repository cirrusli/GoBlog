package common

import (
	"SummerProject/config"
	"SummerProject/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var MDB *gorm.DB

func InitMySQL() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.MySQL.UserName, config.Conf.MySQL.Password, config.Conf.MySQL.Addr, config.Conf.MySQL.DatabaseName)
	// DSN data source name
	//想要正确的处理time.Time,需要带上 parseTime 参数，
	//要支持完整的UTF-8编码，需要将 charset=utf8 更改为 charset=utf8mb4

	var err error
	MDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("connect MySQL failed:", err.Error())
	}
	// 检测User结构体对应的表是否存在
	//if dao.MDB.Migrator().HasTable(&model.MUser{}) {
	//	log.Println("exist")
	//} else {
	//	log.Println("not exist")
	//}

	//绑定模型,自动创建对应的表
	err = MDB.AutoMigrate(&model.MUser{}, &model.Article{}, &model.Comment{}, &model.Reply{}, &model.Like{})
	if err != nil {
		log.Panicln("migrate model MUser failed:", err.Error())
	}
}
