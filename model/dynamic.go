package model

import "time"

type Dynamic struct {
	Id              int64
	AuthorId        int64
	Content         string
	FileUrl         string
	CreatedAt       time.Time
	UpvoteCount     int
	CommentCount    int
	FatherCommentId int64
}

func (Dynamic) TableName() string {
	return "dynamic"
}
