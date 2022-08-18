package controller

import (
	"SummerProject/service"
	"SummerProject/utils"
	"net/http"
)

// GetArticleList 获取文章摘要列表
func GetArticleList(w http.ResponseWriter, r *http.Request) {
	articleRes, err := service.GetSummaryList()
	if err != nil {
		utils.Error(w, err)
		return
	}
	utils.Success(w, articleRes)
}

// GetArticle 根据文章aid获取文章详细内容
func GetArticle(w http.ResponseWriter, r *http.Request) {

}

// PostAndUpdateArticle 发布或更新(删除)文章
func PostAndUpdateArticle(w http.ResponseWriter, r *http.Request) {

	//
	//uid := claim.Uid
	//method := r.Method
	//switch method {
	//case http.MethodPost:
	//	//Post Article
	//	params := utils.GetRequestJsonParams(r)
	//
	//}

}
