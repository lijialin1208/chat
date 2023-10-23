package model

type Friend struct {
	Fid          string `json:"fid" bson:"fid"`
	FriendName   string `json:"friendName" bson:"friendName"`
	ReMark       string `json:"remark" bson:"remark"`
	FriendAvatar string `json:"friendAvatar" bson:"friendAvatar"`
}
