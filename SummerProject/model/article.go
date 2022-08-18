package model

import "gorm.io/gorm"

// Article 文章信息
type Article struct {
	gorm.Model
	Author    string `json:"author"` //与MUser表中的UserName一致
	Title     string `json:"title"`
	Summary   string `json:"summary" gorm:"not null"` //首页获取文章时仅返回摘要
	Content   string `json:"content"`                 //string类型创建字段为Longtext(21亿字节)
	Hits      int    `json:"hits"`                    //点击次数
	Comments  int    `json:"comments"`                //评论数
	Likes     int    `json:"likes"`                   //点赞数
	IsDeleted bool   `json:"is_deleted"`              //是否为已经删除的文章
}
