package controller

import "net/http"

// GetComments 获取当前文章的所有评论
func GetComments(w http.ResponseWriter, r *http.Request) {
	//获取ArticleID
	//todo
}

// PostComment 发布评论
func PostComment(w http.ResponseWriter, r *http.Request) {
	//获取UserID、ArticleID
}
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	//获取ID(comment表中的ID)
}
