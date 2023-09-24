package model

type User struct {
	ID        int64 `gorm:"column:id"`
	Account   string
	Password  string
	Nickname  string
	Headimage string `gorm:"type:text"`
}
type UserBasic struct {
	Account  string `json:"username"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}
