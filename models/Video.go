package models

import "time"

type Video struct {
	VideoId       int64     `json:"video_id"`
	CreateTime    time.Time `json:"create_time"`
	UserId        int64     `json:"user_id"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int       `json:"favorite_count"`
	CommentCount  int       `json:"comment_count"`
	Title         string    `json:"title"`
}
