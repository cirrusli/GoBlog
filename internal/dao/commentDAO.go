package dao

import (
	"SummerProject/internal/model"
	"SummerProject/utils"
)

// GetComments 根据aid获取该文章的所有评论
func GetComments(aid int) (commentList []*model.Comment, err error) {
	if err = utils.MDB.Table("comments").Where("aid=?", aid).Find(&commentList).Error; err != nil {
		return nil, err
	}
	return commentList, nil
}

// PostComment 发布评论
func PostComment(comment *model.Comment) (err error) {
	err = utils.MDB.Table("comments").Create(&comment).Error
	return err

}

// DeleteComment 删除评论
func DeleteComment(cid int) (err error) {
	err = utils.MDB.Table("comments").Where("cid=?", cid).Delete(&model.Comment{}).Error
	return err
}
