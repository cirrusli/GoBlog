package controller

import (
	"GoBlog/common"
	"GoBlog/internal/model"
	"GoBlog/internal/service"
	"net/http"
	"strconv"
)

// Like 点赞操作
func Like(w http.ResponseWriter, r *http.Request) {
	data := common.GetRequestJsonParams(r)
	typeID, _ := strconv.Atoi(data["type_id"].(string))
	likeType := data["like_type"].(bool)
	uid, _ := strconv.Atoi(data["uid"].(string))
	toUid, _ := strconv.Atoi(data["to_uid"].(string))
	status := data["status"].(bool)
	likeStructure := &model.Like{
		TypeID:   typeID,
		LikeType: likeType,
		Uid:      uid,
		ToUid:    toUid,
		Status:   status,
	}
	err := service.IsLiked(typeID, uid)
	if err != nil {
		common.Error(w, err)
		return
	}
	err = service.LikeAction(likeStructure)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, nil)
}
