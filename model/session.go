package model

type Session struct {
	Mid    int64  `json:"mid,string" bson:"mid"`
	Fid    int64  `json:"fid,string" bson:"fid"`
	Remark string `json:"remark" bson:"remark"`
}
