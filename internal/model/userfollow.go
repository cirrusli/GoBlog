package model

// UserFollow 用户的关注与粉丝列表
type UserFollow struct {
	Uid            int `json:"uid"`         //与MUser的ID相同
	FanID          int `json:"fan_id"`      //粉丝列表
	FollowerID     int `json:"follower_id"` //关注列表
	FansCount      int `json:"fans_count"`
	FollowersCount int `json:"followers_count"`
}
