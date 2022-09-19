package model

// ConfStruct 配置文件结构体
type ConfStruct struct {
	MySQL MySQL
	Redis Redis
	RabbitMQ
	WebService WebService
}

//MySQL 配置模块
type MySQL struct {
	UserName     string `mapstructure:"UserName"`
	Password     string `mapstructure:"Password"`
	Addr         string `mapstructure:"Addr"`
	DatabaseName string `mapstructure:"DatabaseName"`
}

// Redis todo 无法解析递归结构的数据，加不加mapstructure都没用
// Redis 配置模块
type Redis struct {
	FollowingList    FollowingList    `mapstructure:"FollowingList"`
	FollowersList    FollowersList    `mapstructure:"FollowersList"`
	MutualFollowList MutualFollowList `mapstructure:"MutualFollowList"`
}
type FollowingList struct {
	Addr     string `mapstructure:"Addr"`
	Password string `mapstructure:"Password"`
	DB       int    `mapstructure:"DB"`
}
type FollowersList struct {
	Addr     string `mapstructure:"Addr"`
	Password string `mapstructure:"Password"`
	DB       int    `mapstructure:"DB"`
}
type MutualFollowList struct {
	Addr     string `mapstructure:"Addr"`
	Password string `mapstructure:"Password"`
	DB       int    `mapstructure:"DB"`
}

//RabbitMQ 配置模块
type RabbitMQ struct {
}

// WebService 服务器配置模块
type WebService struct {
	Addr string `mapstructure:"Addr"`
}
