package model

import "time"

type Dynamic struct {
	Id              int64     `json:"id,string"`
	Avatar          string    `json:"avatar"`
	NickName        string    `json:"nickName" gorm:"type:varchar(15)"`
	AuthorId        int64     `json:"authorId,string"`
	Content         string    `json:"content"`
	FileUrl         string    `json:"fileUrl"`
	CreatedAt       time.Time `json:"createdAt"`
	UpvoteCount     int       `json:"upvoteCount"`
	CommentCount    int       `json:"commentCount"`
	FatherCommentId int64     `json:"fatherCommentId,string"`
}

func (Dynamic) TableName() string {
	return "dynamic"
}
