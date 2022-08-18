package model

import "gorm.io/gorm"

// MUser 用户信息
type MUser struct {
	gorm.Model
	UserName  string `json:"username"`
	Password  string `json:"password"`
	AvatarUrl string `json:"avatarUrl" gorm:"default"` //用户头像路径
	//Age       byte   `json:"age" gorm:"default 18"`    //当该字段在数据库中无默认值时不会生效
	//Gender    byte   `json:"gender"`                   //1 is male,0 is female
}

// LoginRes 登录时携带token
type LoginRes struct {
	Token string `json:"token"`
	Uid   int    `json:"uid"`
}
