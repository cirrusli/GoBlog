package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	TypeID   int  `json:"type_id"`   //对应文章或评论（回复）的ID
	LikeType bool `json:"like_type"` //0是作品（文章）点赞，1是评论（回复）点赞
	Uid      int  `json:"uid"`       //与MUser表中的Uid一致
	ToUid    int  `json:"to_uid"`    //被点赞的用户的uid
	Status   bool `json:"status"`    //0为取消赞，1为点赞

}
