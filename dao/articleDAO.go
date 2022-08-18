package dao

import (
	"SummerProject/model"
	"SummerProject/utils"
)

// GetAllSummaries 用户登录成功后，在主页展示文章摘要
func GetAllSummaries() (ArticleRes []*model.ArticleRes, err error) {
	//在article表中查询选定字段
	err = utils.MDB.Table("articles").
		Select("aid,author,title,cover,summary").Scan(&ArticleRes).Error
	//todo 这里如果传int类型的ID字段为什么传进JSON就是0了。。
	if err != nil {
		return nil, err
	}
	return
}

// PostArticle 用户发布文章
func PostArticle(article *model.Article) (err error) {
	if err = utils.MDB.Table("articles").Save(article).Error; err != nil {
		return err
	}
	return
}

// UpdateArticle 修改、编辑、更新文章
func UpdateArticle() {

}

// DeleteArticle 删除文章
func DeleteArticle() {

}
