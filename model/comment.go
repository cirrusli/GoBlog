package model

import (
	"gorm.io/gorm"
	"time"
)

// Comment 文章评论
type Comment struct {
	gorm.Model
	Cid     int    `json:"cid"`     //EncodeID
	Content string `json:"content"` //Content
	Uid     int    `json:"uid"`     //与MUser表中的Uid一致
	Aid     int    `json:"aid"`     //与Article表中的Aid一致
	Likes   int    `json:"likes"`   //评论点赞数
}

// CommentStrData 给前端传回全部是字符串
type CommentStrData struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Cid       string         `json:"cid"`     //EncodeID
	Content   string         `json:"content"` //Content
	Uid       string         `json:"uid"`     //与MUser表中的Uid一致
	Aid       string         `json:"aid"`     //与Article表中的Aid一致
	Likes     string         `json:"likes"`   //评论点赞数
}

// Reply 评论或回复的回复
type Reply struct {
	gorm.Model
	Rid       int    `json:"rid"`        //EncodeID
	ReplyType bool   `json:"reply_type"` //对评论回复为0，对回复的回复为1
	Content   string `json:"content"`    //Content
	Uid       int    `json:"uid"`        //与MUser表中的Uid一致
	ReplyCid  int    `json:"reply_cid"`  //所回复的评论的cid
	ToUid     int    `json:"to_uid"`     //所回复的用户的uid
}
