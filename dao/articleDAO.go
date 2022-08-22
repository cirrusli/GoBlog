package dao

import (
	"SummerProject/model"
	"SummerProject/utils"
)

// GetAllSummaries 用户登录成功后，在主页展示文章摘要
func GetAllSummaries() (articleRes []*model.ArticleRes, err error) {
	//在article表中查询选定字段
	err = utils.MDB.Table("articles").
		Select("aid,author,title,cover,summary").Where("deleted_at", nil).Scan(&articleRes).Error
	if err != nil {
		return nil, err
	}
	return
}

// GetArticle 根据aid获取文章详情
func GetArticle(aid int) (*model.Article, error) {
	article := new(model.Article)
	if err := utils.MDB.Where("aid=?", aid).First(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

// PostAndUpdateArticle 用户发布/更新文章
func PostAndUpdateArticle(article *model.Article) (err error) {
	if err = utils.MDB.Table("articles").Where("aid=?", article.Aid).Updates(article).Error; err != nil {
		return err
	}
	return err
}

// DeleteArticle 用户删除文章（软删除）
func DeleteArticle(aid int) (err error) {
	//var article *model.Article仅仅声明，没有分配值
	article := new(model.Article)
	if err = utils.MDB.Table("articles").Where("aid=?", aid).Delete(article).Error; err != nil {
		return err
	}
	return err
}
func GetUserArticle(uid int) (articleList []*model.Article, err error) {
	if err = utils.MDB.Table("articles").Where("uid=?", uid).Find(&articleList).Error; err != nil {
		return nil, err
	}
	return articleList, nil
}
