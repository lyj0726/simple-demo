package models

import "time"

type Comment struct {
	CommentId  int64     `json:"comment_id"`
	VideoId    int64     `json:"video_id"`
	UserId     int64     `json:"user_id"`
	Comment    string    `json:"comment"`
	CreateTime time.Time `json:"create_time"`
}
