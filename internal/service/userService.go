package service

import (
	"SummerProject/internal/dao"
	"SummerProject/internal/middleware"
	"SummerProject/internal/model"
	"SummerProject/utils"
	"errors"
	"log"
	"strconv"
)

// Register 用户注册后直接登录
func Register(UserName string, Password string) (*model.LoginRes, error) {
	//用户名查重
	isDuplicate, _ := dao.GetUserByName(UserName)

	log.Println("is duplicate?", isDuplicate)

	if isDuplicate != nil {
		//数据库中已有该用户名
		return nil, errors.New("duplicate username")
		//一般不能将错误详细信息返回，本项目简化了状态码
	} else {
		//无重复,在数据库中添加用户信息
		user := model.MUser{
			UserName: UserName,
			Password: Password,
			Uid:      utils.EncodeID(),
		}
		err := dao.CreateUser(&user)
		if err != nil {
			return nil, errors.New("create user failed")
		}
		token, err := middleware.Award(&user.Uid)
		if err != nil {
			return nil, errors.New("create token failed")
		}

		var registerRes = &model.LoginRes{
			Token:     token,
			Uid:       strconv.Itoa(user.Uid),
			UserName:  user.UserName,
			AvatarUrl: user.AvatarUrl,
		}
		return registerRes, nil
	}

}

// Login 用户登录，携带token
func Login(UserName string, Password string) (*model.LoginRes, error) {
	user := dao.GetUser(UserName, Password)
	if user == nil {
		return nil, errors.New("账号或密码错误！")
	}
	token, err := middleware.Award(&user.Uid)
	if err != nil {
		return nil, errors.New("create token failed")
	}

	var loginRes = &model.LoginRes{
		Token:     token,
		Uid:       strconv.Itoa(user.Uid),
		UserName:  user.UserName,
		AvatarUrl: user.AvatarUrl,
	}
	return loginRes, nil
}
