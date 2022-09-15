package controller

import (
	"SummerProject/common"
	"SummerProject/internal/service"
	"log"
	"net/http"
)

// ResMsg 响应json信息的结构体
type ResMsg struct {
	Msg string `json:"msg"`
}

// Register 用户注册
func Register(w http.ResponseWriter, r *http.Request) {
	registerData := common.GetRequestJsonParams(r)
	UserName := registerData["username"].(string)
	Password := registerData["password"].(string)

	Password = common.Encoder(Password)

	registerRes, err := service.Register(UserName, Password)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, registerRes)
}

// LogIn 用户登录
func LogIn(w http.ResponseWriter, r *http.Request) {
	loginData := common.GetRequestJsonParams(r)
	UserName := loginData["username"].(string)
	Password := loginData["password"].(string)

	Password = common.Encoder(Password)

	loginRes, err := service.Login(UserName, Password)
	if err != nil {
		log.Println("error")
		common.Error(w, err)
		return
	}
	log.Println("success")
	common.Success(w, loginRes)
}

// LogOut 用户退出登录
func LogOut(w http.ResponseWriter, r *http.Request) {
	//token立即过期即可
}

// GetUserInfo 获取用户详细信息
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	//data := common.GetRequestJsonParams(r)
	//uid, _ := strconv.Atoi(data["uid"].(string))

}

// UpdateUserInfo 用户信息修改
func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	//todo 修改用户名，密码，头像
}
