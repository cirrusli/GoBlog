package model

import "gorm.io/gorm"

// Comment 文章评论
type Comment struct {
	gorm.Model
	Cid     string `json:"cid"` //EncodeID
	Comment string `json:"comment"`
	Uid     int    `json:"uid"`   //与MUser表中的Uid一致
	Aid     int    `json:"aid"`   //与Article表中的Aid一致
	Likes   int    `json:"likes"` //评论点赞数
}
type CommentRes struct {
}
