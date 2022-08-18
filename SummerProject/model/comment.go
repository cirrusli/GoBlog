package model

import "gorm.io/gorm"

// Comment 文章评论
type Comment struct {
	gorm.Model
	Comment   string `json:"comment"`
	UserID    int    `json:"user_id"`    //与MUser表中的ID一致（也是uid）
	ArticleID int    `json:"article_id"` //与Article表中的ID一致
	Likes     int    `json:"likes"`      //评论点赞数
	IsDeleted bool   `json:"is_deleted"` //采用软删除
}
