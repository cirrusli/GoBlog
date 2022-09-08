package service

import (
	"SummerProject/internal/dao"
	"SummerProject/internal/model"
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
