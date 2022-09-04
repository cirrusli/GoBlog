package router

import (
	"SummerProject/config"
	"SummerProject/internal/controller"
	"SummerProject/internal/middleware"
	"fmt"
	"log"
	"net/http"
)

//goland:noinspection SpellCheckingInspection
func InitRouter() {
	//attention:frontend request URL!!!
	//test middleware
	http.Handle("/test", middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "test middleware")
		log.Println("执行中间件后的函数")
	})))
	http.HandleFunc("/hello", middleware.Chain(middleware.Hello, middleware.Method("GET"), middleware.Logging()))
	//userModule
	http.HandleFunc("/login", controller.LogIn)
	http.HandleFunc("/logout", controller.LogOut)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/getuserinfo", controller.GetUserInfo)
	http.HandleFunc("/updateuserinfo", controller.UpdateUserInfo)
	//articleModule
	http.HandleFunc("/index", controller.GetArticleList)
	http.HandleFunc("/article", controller.GetArticle)
	http.HandleFunc("/postarticle", controller.PostArticle)
	http.HandleFunc("/updatearticle", controller.UpdateArticle)
	http.HandleFunc("/deletearticle", controller.DeleteArticle)
	http.HandleFunc("/getuserarticle", controller.GetUserArticle)
	//commentModule
	http.HandleFunc("/getcomments", controller.GetComments)
	http.HandleFunc("/postcomment", controller.PostComment)
	http.HandleFunc("/deletecomment", controller.DeleteComment) //IsDeleted=true
	http.HandleFunc("/reply2comment", controller.Reply2Comment)
	//likeModule
	http.HandleFunc("/likearticle", controller.LikeArticle)
	http.HandleFunc("/likecomment", controller.LikeComment)
	//followModule
	http.HandleFunc("/follow", controller.FollowUser)
	http.HandleFunc("/unfollow", controller.UnFollowUser)
	http.HandleFunc("/getfollowinglist", controller.GetFollowingList) //我关注的
	http.HandleFunc("/getfollowerlist", controller.GetFollowerList)   //关注我的
	//set up port listening
	server := http.Server{
		Addr:    config.Conf.WebService.Addr,
		Handler: http.DefaultServeMux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Panicln("set up router failed:", err.Error())
	}
}
