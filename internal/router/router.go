package router

import (
	"SummerProject/common"
	"SummerProject/config"
	"SummerProject/internal/controller"
	"SummerProject/internal/middleware"
	"SummerProject/internal/middleware/jwt"
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
	http.HandleFunc("/login", middleware.Chain(controller.LogIn, middleware.Logging()))
	http.HandleFunc("/logout", controller.LogOut)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/getuserinfo", controller.GetUserInfo)
	http.HandleFunc("/updateuserinfo", middleware.Chain(controller.UpdateUserInfo, jwt.AuthJWT()))
	//articleModule
	http.HandleFunc("/index", middleware.Chain(controller.GetArticleList, middleware.Method("GET"), middleware.Logging()))
	http.HandleFunc("/article", middleware.Chain(controller.GetArticle, middleware.Method("GET")))
	http.HandleFunc("/postarticle", middleware.Chain(controller.PostArticle, jwt.AuthJWT()))
	http.HandleFunc("/updatearticle", middleware.Chain(controller.UpdateArticle, jwt.AuthJWT()))
	http.HandleFunc("/deletearticle", middleware.Chain(controller.DeleteArticle, jwt.AuthJWT()))
	http.HandleFunc("/getuserarticle", middleware.Chain(controller.GetUserArticle, middleware.Method("GET")))
	//commentModule
	http.HandleFunc("/getcomments", middleware.Chain(controller.GetComments, middleware.Method("GET")))
	http.HandleFunc("/postcomment", middleware.Chain(controller.PostComment, jwt.AuthJWT()))
	http.HandleFunc("/deletecomment", middleware.Chain(controller.DeleteComment, jwt.AuthJWT())) //IsDeleted=true
	http.HandleFunc("/reply2comment", middleware.Chain(controller.Reply2Comment, jwt.AuthJWT()))
	//likeModule
	http.HandleFunc("/like", middleware.Chain(controller.Like, jwt.AuthJWT()))
	//followModule
	http.HandleFunc("/follow", middleware.Chain(controller.FollowUser, jwt.AuthJWT()))
	http.HandleFunc("/unfollow", middleware.Chain(controller.UnFollowUser, jwt.AuthJWT()))
	http.HandleFunc("/getfollowinglist", middleware.Chain(controller.GetFollowingList, jwt.AuthJWT())) //我关注的
	http.HandleFunc("/getfollowerlist", middleware.Chain(controller.GetFollowerList, jwt.AuthJWT()))   //关注我的
	//set up port listening
	server := http.Server{
		Addr:    config.Conf.WebService.Addr,
		Handler: http.DefaultServeMux,
	}
	err := server.ListenAndServe()
	common.PanicLog("set up router failed:", err)
}
