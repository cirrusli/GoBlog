package service

import (
	"SummerProject/dao"
	"SummerProject/model"
	"errors"
	"strconv"
)

// GetSummaryList 获取文章摘要列表
func GetSummaryList() ([]*model.ArticleRes, error) {
	summaries, err := dao.GetAllSummaries()
	if err != nil {
		return nil, errors.New("获取文章摘要列表失败！")
	}
	return summaries, nil //返回摘要列表
}

// GetArticle 获取文章详情
func GetArticle(aid int) (strData *model.StrData, err error) {
	article := new(model.Article)
	article, err = dao.GetArticle(aid)
	strData = &model.StrData{
		ID:        strconv.Itoa(int(article.ID)),
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		DeletedAt: article.DeletedAt,
		Aid:       strconv.Itoa(article.Aid),
		Author:    article.Author,
		Uid:       strconv.Itoa(article.Uid),
		Title:     article.Title,
		Cover:     article.Cover,
		Content:   article.Content,
		Hits:      strconv.Itoa(article.Hits),
		Comments:  strconv.Itoa(article.Comments),
		Likes:     strconv.Itoa(article.Likes),
	}

	if err != nil {
		return nil, errors.New("获取文章详情失败！")
	}
	return strData, nil
}

// PostAndUpdateArticle 发布或更新文章
func PostAndUpdateArticle(article *model.Article) (err error) {
	err = dao.PostAndUpdateArticle(article)
	if err != nil {
		return errors.New("文章发布(更新)失败！")
	}
	return nil
}

// DeleteArticle 删除文章
func DeleteArticle(aid int) (err error) {
	err = dao.DeleteArticle(aid)
	if err != nil {
		return errors.New("文章删除失败！")
	}
	return nil
}

// GetUserArticle 通过uid获取用户的所有文章
func GetUserArticle(uid int) (articleList []*model.Article, err error) {
	articleList, err = dao.GetUserArticle(uid)
	if err != nil {
		return nil, errors.New("该作者的文章列表获取失败！")
	}
	return articleList, nil
}
