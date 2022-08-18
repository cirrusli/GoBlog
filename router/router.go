package router

import (
	"SummerProject/controller"
	"log"
	"net/http"
)

//goland:noinspection SpellCheckingInspection
func SetUpRouter() {
	//处理对应页面的请求，注意URL与前端的是否一致！！！

	//用户模块
	http.HandleFunc("/login", controller.LogIn)
	http.HandleFunc("/logout", controller.LogOut)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/updateuserinfo", controller.UpdateUserInfo)
	//文章模块
	http.HandleFunc("/articles", controller.GetArticleList)
	http.HandleFunc("/postarticle", controller.PostAndUpdateArticle)
	http.HandleFunc("/updatearticle", controller.PostAndUpdateArticle) //包括删除(IsDeleted=true)
	//评论模块
	http.HandleFunc("/getcomments", controller.GetComments)
	http.HandleFunc("/postcomment", controller.PostComment)
	http.HandleFunc("/deletecomment", controller.DeleteComment) //IsDeleted=true

	//设置端口监听
	server := http.Server{
		Addr:    "127.0.0.1:6666",
		Handler: nil,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Panicln("set up routers failed:", err.Error())
	}
}
