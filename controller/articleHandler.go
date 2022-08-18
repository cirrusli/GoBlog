package controller

import (
	"net/http"
)

// GetArticleList 获取文章摘要列表
func GetArticleList(w http.ResponseWriter, r *http.Request) {

}

// GetArticle 根据文章aid获取文章详细内容
func GetArticle(w http.ResponseWriter, r *http.Request) {

}

// PostAndUpdateArticle 发布或更新(删除)文章
func PostAndUpdateArticle(w http.ResponseWriter, r *http.Request) {
	//判断用户是否登录
	//token := r.Header.Get("Authorization")
	//_, claim, err := middleware.ParseToken(token)
	//if err != nil {
	//	utils.Error(w, errors.New("登录已过期，请重新登录！"))
	//}//todo 这部分使用中间件处理
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
