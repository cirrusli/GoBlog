package service

import (
	"GoBlog/internal/dao"
	"GoBlog/internal/model"
	"errors"
)

func LikeAction(likeStructure *model.Like) (err error) {
	//todo add redis cache
	err = dao.LikeAction(likeStructure)
	if err != nil {
		return errors.New("点赞失败")
	}
	return nil
}
func IsLiked(typeID, uid int) (err error) {
	if dao.IsLiked(typeID, uid) {
		return errors.New("already liked")
	}
	return nil
}
