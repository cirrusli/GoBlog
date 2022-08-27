package controller

import (
	"SummerProject/model"
	"SummerProject/service"
	"SummerProject/utils"
	"net/http"
	"strconv"
)

// GetComments 获取当前文章的所有评论
func GetComments(w http.ResponseWriter, r *http.Request) {
	//获取ArticleID
	data := r.URL.Query().Get("aid")
	aid, _ := strconv.Atoi(data)
	commentList, err := service.GetComments(aid)
	if err != nil {
		utils.Error(w, err)
	}
	utils.Success(w, commentList)
}

// PostComment 发布评论
func PostComment(w http.ResponseWriter, r *http.Request) {
	data := utils.GetRequestJsonParams(r)
	uid, _ := strconv.Atoi(data["uid"].(string))
	aid, _ := strconv.Atoi(data["aid"].(string))

	comment := &model.Comment{
		Cid:     utils.EncodeID(),
		Content: data["content"].(string),
		Uid:     uid,
		Aid:     aid,
	}
	cid := strconv.Itoa(comment.Cid)
	err := service.PostComment(comment)
	if err != nil {
		utils.Error(w, err)
	}
	utils.Success(w, cid)
}
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	//获取ID(comment表中的ID)
}
