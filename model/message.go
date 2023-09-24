package model

import "time"

type Content struct {
	Kind int         `json:"kind"`
	Data interface{} `json:"data"`
}

type MacAddress struct {
	Mac string `json:"mac"`
}

type Message struct {
	Mtype    int       `json:"mtype" bson:"mtype"` //0表示单聊/1表示群聊/2表示添加好友
	FromID   int64     `json:"fromID,string" bson:"fromID"`
	ToID     int64     `json:"toID,string" bson:"toID"`
	Content  string    `json:"content" bson:"content"`
	Kind     int       `json:"kind" bson:"kind"`
	CreateAt time.Time `json:"createAt" bson:"createAt"`
}
type MessagePlus struct {
	ID       string    `json:"ID" bson:"_id"`
	Mtype    int       `json:"mtype" bson:"mtype"` //0表示单聊/1表示群聊/2表示添加好友
	FromID   int64     `json:"fromID" bson:"fromID"`
	ToID     int64     `json:"toID" bson:"toID"`
	Content  string    `json:"content" bson:"content"`
	Kind     int       `json:"kind" bson:"kind"`
	CreateAt time.Time `json:"createAt" bson:"createAt"`
}
