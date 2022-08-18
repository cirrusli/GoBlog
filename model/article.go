package model

import "gorm.io/gorm"

// Article 文章信息
type Article struct {
	gorm.Model
	Aid       int    `json:"aid"`    //EncodeID生成
	Author    string `json:"author"` //与MUser表中的UserName一致
	Uid       int    `json:"uid"`
	Title     string `json:"title"`
	Cover     string `json:"cover" gorm:"default:https://img-blog.csdnimg.cn/img_convert/7ad86d3945ddf7946e39d179e60c5687.png"` //封面图
	Summary   string `json:"summary" gorm:"not null"`                                                                           //首页获取文章时仅返回摘要
	Content   string `json:"content"`                                                                                           //string类型创建字段为Longtext(21亿字节)
	Hits      int    `json:"hits"`                                                                                              //点击次数
	Comments  int    `json:"comments"`                                                                                          //评论数
	Likes     int    `json:"likes"`                                                                                             //点赞数
	IsDeleted bool   `json:"is_deleted"`                                                                                        //是否为已经删除的文章
}

// ArticleRes 首页文章摘要的data部分
type ArticleRes struct {
	Aid     int    `json:"aid"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Cover   string `json:"cover"`
	Summary string `json:"summary"`
}
