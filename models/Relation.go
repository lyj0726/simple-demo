package models

import "time"

type Relation struct {
	RelationId int64     `json:"relation_id"`
	FolloweeId int64     `json:"followee_id"`
	FollowerId int64     `json:"follower_id"`
	CreateTime time.Time `json:"create_time"`
}
