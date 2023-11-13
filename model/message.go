package model

type Content struct {
	Kind int         `json:"kind"`
	Data interface{} `json:"data"`
}

type MacAddress struct {
	Mac string `json:"mac"`
}
type MessageBasic struct {
	Mtype   int    `json:"mtype" bson:"mtype"` //0表示单聊/1表示群聊/2表示添加好友
	FromID  int64  `json:"fromID,string" bson:"fromID"`
	ToID    int64  `json:"toID,string" bson:"toID"`
	Content string `json:"content" bson:"content"`
	Kind    int    `json:"kind" bson:"kind"`
}
type Message struct {
	Mtype    int    `json:"mtype" bson:"mtype"` //0表示单聊/1表示群聊/2表示添加好友
	FromID   int64  `json:"fromID,string" bson:"fromID"`
	ToID     int64  `json:"toID,string" bson:"toID"`
	Content  string `json:"content" bson:"content"`
	Kind     int    `json:"kind" bson:"kind"` //0表示文本/1表示图片/2表示语音
	CreateAt string `json:"createAt" bson:"createAt"`
	IsRead   bool   `json:"isRead" bson:"isRead"`
	Length   int    `json:"length" bson:"length"`
}
type MessagePlus struct {
	ID       string `json:"ID" bson:"_id"`
	Mtype    int    `json:"mtype" bson:"mtype"` //0表示单聊/1表示群聊/2表示添加好友
	FromID   int64  `json:"fromID" bson:"fromID"`
	ToID     int64  `json:"toID" bson:"toID"`
	Content  string `json:"content" bson:"content"`
	Kind     int    `json:"kind" bson:"kind"`
	CreateAt string `json:"createAt" bson:"createAt"`
}
