package model

import (
	"gorm.io/gorm"
	"time"
)

// Article 文章信息
type Article struct {
	gorm.Model
	Aid      int    `json:"aid" gorm:"uniqueIndex"` //EncodeID生成
	Author   string `json:"author"`                 //与MUser表中的UserName一致
	Uid      int    `json:"uid"`
	Title    string `json:"title"`
	Cover    string `json:"cover" gorm:"default:https://img-blog.csdnimg.cn/img_convert/7ad86d3945ddf7946e39d179e60c5687.png"` //封面图
	Summary  string `json:"summary" gorm:"not null"`                                                                           //首页获取文章时仅返回摘要
	Content  string `json:"content"`                                                                                           //string类型创建字段为Longtext(21亿字节)
	Hits     int    `json:"hits"`                                                                                              //点击次数
	Comments int    `json:"comments"`                                                                                          //评论数
	Likes    int    `json:"likes"`                                                                                             //点赞数 	//是否为已经删除的文章
}

// ArticleRes 首页文章摘要的data部分
type ArticleRes struct {
	Aid     string `json:"aid"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Cover   string `json:"cover"`
	Summary string `json:"summary"`
}

// StrData 给前端传回全部是字符串
type StrData struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Aid       string         `json:"aid"`    //EncodeID生成
	Author    string         `json:"author"` //与MUser表中的UserName一致
	Uid       string         `json:"uid"`
	Title     string         `json:"title"`
	Cover     string         `json:"cover"`    //封面图
	Content   string         `json:"content"`  //string类型创建字段为Longtext(21亿字节)
	Hits      string         `json:"hits"`     //点击次数
	Comments  string         `json:"comments"` //评论数
	Likes     string         `json:"likes"`    //点赞数
}
