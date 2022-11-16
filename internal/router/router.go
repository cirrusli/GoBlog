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
	//test connection
	http.HandleFunc("/testConn", middleware.Chain(testConn, middleware.SetHeaders()))
	//test middleware
	http.Handle("/test", middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "test middleware")
		log.Println("执行中间件后的函数")
	})))
	http.HandleFunc("/hello", middleware.Chain(middleware.Hello, middleware.Method("GET"), middleware.Logging()))
	//userModule
	http.HandleFunc("/login", middleware.Chain(controller.LogIn, middleware.Logging(), middleware.SetHeaders()))
	http.HandleFunc("/logout", middleware.Chain(controller.LogOut, middleware.SetHeaders()))
	http.HandleFunc("/register", middleware.Chain(controller.Register, middleware.SetHeaders()))
	http.HandleFunc("/getuserinfo", middleware.Chain(controller.GetUserInfo, middleware.SetHeaders()))
	http.HandleFunc("/updateuserinfo", middleware.Chain(controller.UpdateUserInfo, jwt.AuthJWT(), middleware.SetHeaders()))
	//articleModule
	http.HandleFunc("/index", middleware.Chain(controller.GetArticleList, middleware.Method("GET"), middleware.Logging(), middleware.SetHeaders()))
	http.HandleFunc("/article", middleware.Chain(controller.GetArticle, middleware.Method("GET"), middleware.SetHeaders()))
	http.HandleFunc("/postarticle", middleware.Chain(controller.PostArticle, jwt.AuthJWT(), middleware.SetHeaders()))
	http.HandleFunc("/updatearticle", middleware.Chain(controller.UpdateArticle, jwt.AuthJWT(), middleware.SetHeaders()))
	http.HandleFunc("/deletearticle", middleware.Chain(controller.DeleteArticle, jwt.AuthJWT(), middleware.SetHeaders()))
	http.HandleFunc("/getuserarticle", middleware.Chain(controller.GetUserArticle, middleware.Method("GET"), middleware.SetHeaders()))
	//commentModule
	http.HandleFunc("/getcomments", middleware.Chain(controller.GetComments, middleware.Method("GET"), middleware.SetHeaders()))
	http.HandleFunc("/postcomment", middleware.Chain(controller.PostComment, jwt.AuthJWT(), middleware.SetHeaders()))
	http.HandleFunc("/deletecomment", middleware.Chain(controller.DeleteComment, jwt.AuthJWT(), middleware.SetHeaders())) //IsDeleted=true
	http.HandleFunc("/reply2comment", middleware.Chain(controller.Reply2Comment, jwt.AuthJWT(), middleware.SetHeaders()))
	//likeModule
	http.HandleFunc("/like", middleware.Chain(controller.Like, jwt.AuthJWT(), middleware.SetHeaders()))
	//followModule
	http.HandleFunc("/follow", middleware.Chain(controller.FollowUser, jwt.AuthJWT(), middleware.SetHeaders()))
	http.HandleFunc("/unfollow", middleware.Chain(controller.UnFollowUser, jwt.AuthJWT(), middleware.SetHeaders()))
	http.HandleFunc("/getfollowinglist", middleware.Chain(controller.GetFollowingList, jwt.AuthJWT(), middleware.SetHeaders())) //我关注的
	http.HandleFunc("/getfollowerlist", middleware.Chain(controller.GetFollowerList, jwt.AuthJWT(), middleware.SetHeaders()))   //关注我的
	//set up port listening
	server := http.Server{
		Addr:    config.Conf.WebService.Addr,
		Handler: http.DefaultServeMux,
	}
	err := server.ListenAndServe()
	common.PanicLog("set up router failed:", err)
}
func testConn(w http.ResponseWriter, r *http.Request) {
	log.Println("testConn succeeded,the request IP:", common.GetRemoteIP(r))
	_, _ = w.Write([]byte("success"))
}
