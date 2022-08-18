package service

import (
	"SummerProject/dao"
	"SummerProject/middleware"
	"SummerProject/model"
	"errors"
	"log"
)

// Register 用户注册后直接登录
func Register(UserName string, Password string) (*model.LoginRes, error) {
	//用户名查重
	isDuplicate, _ := dao.GetUserByName(UserName)

	log.Println("is duplicate", isDuplicate)

	if isDuplicate != nil {
		//数据库中已有该用户名
		return nil, errors.New("duplicate username")
		//一般不能将错误详细信息返回，本项目简化了状态码
	} else {
		//无重复,在数据库中添加用户信息
		user := model.MUser{
			UserName: UserName,
			Password: Password,
		}
		err := dao.CreateUser(&user)
		if err != nil {
			return nil, errors.New("create user failed")
		}
		uid := int(user.ID)
		token, err := middleware.Award(&uid)
		if err != nil {
			return nil, errors.New("create token failed")
		}

		var registerRes = &model.LoginRes{
			Token: token,
			Uid:   uid, //todo 往前端传用户哪些信息
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
	uid := int(user.ID)
	token, err := middleware.Award(&uid)
	if err != nil {
		return nil, errors.New("create token failed")
	}

	var loginRes = &model.LoginRes{
		Token: token,
		Uid:   uid,
	}
	return loginRes, nil
}
