package routers

import (
	controller2 "SummerProject/internal/controller"
	"log"
	"net/http"
)

//goland:noinspection SpellCheckingInspection
func SetUpRouter() {
	//处理对应页面的请求，注意URL与前端的是否一致！！！

	//用户模块
	http.HandleFunc("/login", controller2.LogIn)
	http.HandleFunc("/logout", controller2.LogOut)
	http.HandleFunc("/register", controller2.Register)
	http.HandleFunc("/getuserinfo", controller2.GetUserInfo)
	http.HandleFunc("/updateuserinfo", controller2.UpdateUserInfo)
	//文章模块
	http.HandleFunc("/index", controller2.GetArticleList)
	http.HandleFunc("/article", controller2.GetArticle)
	http.HandleFunc("/postarticle", controller2.PostArticle)
	http.HandleFunc("/updatearticle", controller2.UpdateArticle)
	http.HandleFunc("/deletearticle", controller2.DeleteArticle)
	http.HandleFunc("/getuserarticle", controller2.GetUserArticle)
	//评论模块
	http.HandleFunc("/getcomments", controller2.GetComments)
	http.HandleFunc("/postcomment", controller2.PostComment)
	http.HandleFunc("/deletecomment", controller2.DeleteComment) //IsDeleted=true
	//点赞模块
	http.HandleFunc("/likearticle", controller2.LikeArticle)
	http.HandleFunc("/likecomment", controller2.LikeComment)
	//关注模块
	http.HandleFunc("/follow", controller2.FollowUser)
	http.HandleFunc("/unfollow", controller2.UnFollowUser)
	http.HandleFunc("/getfollowinglist", controller2.GetFollowingList) //我关注的
	http.HandleFunc("/getfollowerlist", controller2.GetFollowerList)   //关注我的
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
