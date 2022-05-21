package models

type User struct {
	UserId int64 `json:"user_id"` 
	UserName string `json:"user_name"` 
	Password string `json:"password"` 
	FollowCount int `json:"follow_count"` 
	FollowerCount int `json:"follower_count"` 
}