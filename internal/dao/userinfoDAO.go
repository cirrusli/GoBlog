package dao

import (
	"SummerProject/common"
	"SummerProject/internal/model"
)

// CreateUser 创建用户信息
func CreateUser(user *model.User) (err error) {
	err = common.MDB.Create(user).Error //use pointer to create
	return err
}

// GetUser 验证用户信息
func GetUser(UserName string, Password string) *model.User {
	user := new(model.User)
	if err := common.MDB.Where("user_name=? and password=?", UserName, Password).First(user).Error; err != nil {
		return nil
	}
	return user
}

// UpdateUser 更新用户信息
func UpdateUser(user *model.User) (err error) {
	err = common.MDB.Save(user).Error
	return err
}

// DeleteUser 通过Uid删除用户信息
func DeleteUser(uid int) (err error) {
	err = common.MDB.Where("id=?", uid).Delete(&model.User{}).Error
	return err
}

// GetAllUsers 后台查询所有用户信息
func GetAllUsers() (users []*model.User, err error) {
	if err = common.MDB.Find(&users).Error; err != nil {
		return nil, err
	}
	return
}

// GetUserByName 通过name查询用户(可对用户信息查重)
func GetUserByName(name string) (user *model.User, err error) {
	//在返回值中定义的指针没有默认值，应当先初始化！！！
	user = new(model.User)
	if err = common.MDB.Where("user_name=?", name).First(user).Error; err != nil {
		return nil, err
	}
	return
}
