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
	http.HandleFunc("/getuserinfo", controller.GetUserInfo)
	http.HandleFunc("/updateuserinfo", controller.UpdateUserInfo)
	//文章模块
	http.HandleFunc("/index", controller.GetArticleList)
	http.HandleFunc("/article", controller.GetArticle)
	http.HandleFunc("/postarticle", controller.PostArticle)
	http.HandleFunc("/updatearticle", controller.UpdateArticle)
	http.HandleFunc("/deletearticle", controller.DeleteArticle)
	http.HandleFunc("/getuserarticle", controller.GetUserArticle)
	//评论模块
	http.HandleFunc("/getcomments", controller.GetComments)
	http.HandleFunc("/postcomment", controller.PostComment)
	http.HandleFunc("/deletecomment", controller.DeleteComment) //IsDeleted=true
	//点赞模块
	http.HandleFunc("/likearticle", controller.LikeArticle)
	http.HandleFunc("/likecomment", controller.LikeComment)
	//关注模块
	http.HandleFunc("/follow", controller.FollowUser)
	http.HandleFunc("/unfollow", controller.UnFollowUser)
	http.HandleFunc("/getfollowers", controller.GetFollowingList) //我关注的
	http.HandleFunc("/getfollowees", controller.GetFollowerList)  //关注我的
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
