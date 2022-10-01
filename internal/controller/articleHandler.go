package controller

import (
	"SummerProject/common"
	"SummerProject/internal/model"
	"SummerProject/internal/service"
	"github.com/yitter/idgenerator-go/idgen"
	"log"
	"net/http"
	"strconv"
)

// GetArticleList 获取文章摘要列表
//todo paginated list
func GetArticleList(w http.ResponseWriter, r *http.Request) {
	articleRes, err := service.GetSummaryList()
	if err != nil {
		common.Error(w, err)
		return
	}
	log.Println("GetArticleList: get article list succeeded")
	common.Success(w, articleRes)
}

// GetArticle 根据文章aid获取文章详细内容
func GetArticle(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("aid")
	aid, _ := strconv.Atoi(data)
	articleRes, err := service.GetArticle(aid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, articleRes)
}

// PostArticle 发布文章
func PostArticle(w http.ResponseWriter, r *http.Request) {
	//todo jwt中间件鉴权
	data := common.GetRequestJsonParams(r)
	//前端返回的文章数据如下：
	//字符串类型需要转回int
	uid, _ := strconv.Atoi(data["uid"].(string))
	//刚发布的文章哪来下面这些
	//hits, _ := strconv.Atoi(data["hits"].(string))
	//comments, _ := strconv.Atoi(data["comments"].(string))
	//likes, _ := strconv.Atoi(data["likes"].(string))
	article := &model.Article{
		Aid:     int(idgen.NextId()),
		Uid:     uid,
		Author:  data["author"].(string),
		Title:   data["title"].(string),
		Cover:   data["cover"].(string),
		Summary: data["summary"].(string),
		Content: data["content"].(string),
	}
	aid := strconv.Itoa(article.Aid)
	err := service.PostArticle(article)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, aid) //发布成功
}

// UpdateArticle 更新文章
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	//鉴权
	data := common.GetRequestJsonParams(r)
	uid, _ := strconv.Atoi(data["uid"].(string))
	aid, _ := strconv.Atoi(data["aid"].(string))
	article := &model.Article{
		Aid:     aid,
		Uid:     uid,
		Author:  data["author"].(string),
		Title:   data["title"].(string),
		Cover:   data["cover"].(string),
		Summary: data["summary"].(string),
		Content: data["content"].(string),
	}
	err := service.UpdateArticle(article)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, nil) //更新成功
}

// DeleteArticle 通过aid删除文章
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	data := common.GetRequestJsonParams(r)
	aid, _ := strconv.Atoi(data["aid"].(string))
	err := service.DeleteArticle(aid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, nil)
}

// GetUserArticle 由uid查看用户的所有文章
func GetUserArticle(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("uid")
	uid, _ := strconv.Atoi(data)
	articleList, err := service.GetUserArticle(uid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, articleList)
}
