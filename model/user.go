package model

type User struct {
	ID        int64 `gorm:"column:id" json:",Id,string"`
	Account   string
	Password  string
	Nickname  string
	Headimage string `gorm:"type:text"`
}
type UserBasic struct {
	Account  string `json:"username"`
	Password string `json:"password"`
	Mac      string `json:"mac"`
}
type UserInfo struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

func (User) TableName() string {
	return "user"
}
