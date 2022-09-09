package dao

import (
	"SummerProject/common"
	"SummerProject/internal/model"
	"gorm.io/gorm"
)

// LikeAction 点赞/取消赞操作
func LikeAction(likeStructure *model.Like) (err error) {
	if err = common.MDB.Table("likes").Create(likeStructure).Error; err != nil {
		return err
	}
	//判断是点赞还是取消点赞操作
	if likeStructure.Status {
		//True(1)为 点赞操作

		//处理被点赞内容的类型
		if likeStructure.LikeType {
			//True(1)为 评论（回复）点赞
			//todo how to merge like operations
			if err = common.MDB.Table("comments").Where("cid=?", likeStructure.TypeID).UpdateColumn("likes", gorm.Expr("likes+?", 1)).Error; err != nil {
				return err
			}
		} else {
			//False(0)为 文章点赞
			if err = common.MDB.Table("articles").Where("aid=?", likeStructure.TypeID).UpdateColumn("likes", gorm.Expr("likes+?", 1)).Error; err != nil {
				return err
			}
		}
	} else {
		//False(0)为 取消点赞操作

		//处理被点赞内容的类型
		if likeStructure.LikeType {
			//True(1)为 取消评论（回复）点赞
			if err = common.MDB.Table("comments").Where("cid=?", likeStructure.TypeID).UpdateColumn("likes", gorm.Expr("likes-?", 1)).Error; err != nil {
				return err
			}
		} else {
			//False(0)为 取消文章点赞
			if err = common.MDB.Table("articles").Where("aid=?", likeStructure.TypeID).UpdateColumn("likes", gorm.Expr("likes-?", 1)).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// IsLiked 是否已点赞
func IsLiked(typeID, uid int) bool {
	like := &model.Like{}
	like = new(model.Like)
	if err := common.MDB.Table("likes").Where("type_id=? and uid=?", typeID, uid).First(like).Error; err != nil {
		return false
	}
	return true
}
