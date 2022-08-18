package dao

import (
	"SummerProject/model"
	"SummerProject/utils"
)

// PostComment 发布评论
func PostComment(comment *model.Comment) (err error) {
	err = utils.MDB.Create(&comment).Error
	return err

}

// DeleteComment 删除评论
func DeleteComment(commentID int) (err error) {
	err = utils.MDB.Delete(&model.Comment{}, commentID).Error
	return err
}
