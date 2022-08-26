package service

import (
	"SummerProject/dao"
	"SummerProject/model"
	"errors"
)

// GetComments 根据aid获取文章的所有评论
func GetComments(aid int) (commentList []*model.Comment, err error) {
	commentList, err = dao.GetComments(aid)
	if err != nil {
		return nil, errors.New("获取评论列表失败")
	}
	return commentList, nil
}
func PostComment(comment *model.Comment) (err error) {
	err = dao.PostComment(comment)
	if err != nil {
		return errors.New("发布评论失败")
	}
	return nil
}
