package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	TypeID int  `json:"typeId"` //对应文章或评论（回复）的ID
	Type   bool `json:"type"`   //0是作品（文章）点赞，1是评论（回复）点赞
	Uid    int  `json:"uid"`    //与MUser表中的Uid一致
	Status bool `json:"status"` //0为取消赞，1为点赞
}
