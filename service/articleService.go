package service

import (
	"SummerProject/dao"
	"SummerProject/model"
)

// GetSummaryList 获取文章摘要列表
func GetSummaryList() ([]*model.ArticleRes, error) {
	summaries, err := dao.GetAllSummaries()
	if err != nil {
		return nil, err
	}
	return summaries, nil //返回摘要列表
}

// GetArticle 获取文章详情
func GetArticle(aid int) {

}

// PostAndUpdateArticle 发布或更新文章
func PostAndUpdateArticle(UserName string) {
	//todo jwt中间件鉴权

}
