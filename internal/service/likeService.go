package service

import (
	"SummerProject/internal/dao"
	"SummerProject/internal/model"
)

func LikeAction(likeStructure *model.Like) (err error) {
	err = dao.LikeAction(likeStructure)
	if err != nil {
		return err
	}
	return nil
}
