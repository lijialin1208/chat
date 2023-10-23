package model

type Session struct {
	Mid    string `json:"mid,string" bson:"mid"`
	Fid    string `json:"fid,string" bson:"fid"`
	Remark string `json:"remark" bson:"remark"`
}
