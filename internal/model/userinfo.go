package model

import "gorm.io/gorm"

// User 用户信息
type User struct {
	gorm.Model
	Uid       int    `json:"uid" gorm:"primaryKey"` //EncodeID
	UserName  string `json:"username"`
	Password  string `json:"password"`
	AvatarUrl string `json:"avatarUrl" gorm:"default:https://image-1302243118.cos.ap-beijing.myqcloud.com/public/img/BLUE_GLASSES_GOPHER-1616727843503.png"` //用户头像路径
	//Age       byte  `json:"age" gorm:"default 18"`    //当该字段在数据库中无默认值时不会生效
	//Gender    byte  `json:"gender"`                   //1 is male,0 is female
	//PhoneNumber string `json:"phone_number"` //手机号
}

// LoginRes 登录时携带token
type LoginRes struct {
	Token     string `json:"token"` //含有uid
	Uid       string `json:"uid"`
	UserName  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}
