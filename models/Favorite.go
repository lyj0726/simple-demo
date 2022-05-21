package models

type Favorite struct {
	FavoriteId int64 `json:"favorite_id"` 
	UserId int64 `json:"user_id"` 
	VideoId int64 `json:"video_id"` 
}