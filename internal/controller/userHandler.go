package controller

import (
	"SummerProject/internal/service"
	"SummerProject/utils"
	"net/http"
)

// ResMsg 响应json信息的结构体
type ResMsg struct {
	Msg string `json:"msg"`
}

// Register 用户注册
func Register(w http.ResponseWriter, r *http.Request) {
	registerData := utils.GetRequestJsonParams(r)
	UserName := registerData["username"].(string)
	Password := registerData["password"].(string)

	Password = utils.Encoder(Password)

	registerRes, err := service.Register(UserName, Password)
	if err != nil {
		utils.Error(w, err)
		return
	}
	utils.Success(w, registerRes)
}

// LogIn 用户登录
func LogIn(w http.ResponseWriter, r *http.Request) {
	loginData := utils.GetRequestJsonParams(r)
	UserName := loginData["username"].(string)
	Password := loginData["password"].(string)

	Password = utils.Encoder(Password)

	loginRes, err := service.Login(UserName, Password)
	if err != nil {
		utils.Error(w, err)
		return
	}
	utils.Success(w, loginRes)
}

// LogOut 用户退出登录
func LogOut(w http.ResponseWriter, r *http.Request) {
	//token立即过期即可
}

// GetUserInfo 获取用户详细信息
func GetUserInfo(w http.ResponseWriter, r *http.Request) {

}

// UpdateUserInfo 用户信息修改
func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {

}
