package DB

import (
	"chat/dal/initDB"
	"chat/model"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func SelectUserByAccount(username string) (user *model.User, err error) {
	m := model.User{}
	result := initDB.MYSQL_DB.Where("account = ?", username).Limit(1).First(&m)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.User{}, result.Error
	} else if result.Error != nil {
		return &model.User{}, result.Error
	}
	return &m, nil
}

func InsertUser(id int64, username string, password string) error {
	user := model.User{
		ID:        id,
		Account:   username,
		Password:  password,
		Nickname:  "chat" + strconv.FormatInt(time.Now().Unix(), 10),
		Headimage: "https://cn.bing.com/images/search?q=%e5%a4%b4%e5%83%8f&id=D7D19EEE8EF16C8AE2714AAD1F96855E4028D5C0&FORM=IQFRBA",
	}
	result := initDB.MYSQL_DB.Create(&user)
	if result.Error != nil && result.RowsAffected == 1 {
		return nil
	} else {
		return result.Error
	}
}

func GetUserInfoById(fid int64) (user *model.User, err error) {
	m := model.User{}
	result := initDB.MYSQL_DB.Select("account,nickname,headimage").Where("id = ?", fid).Limit(1).First(&m)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &model.User{}, result.Error
	} else if result.Error != nil {
		return &model.User{}, result.Error
	}
	return &m, nil
}

func GetDynamics(page int) ([]model.Dynamic, error) {
	//var count int64
	dynamics := make([]model.Dynamic, 0)
	result := initDB.MYSQL_DB.Model(&model.Dynamic{}).Offset(page * 15).Limit(15).Find(&dynamics)
	//result.Count(&count)
	if result.Error != nil {
		return dynamics, result.Error
	}
	return dynamics, nil
}
func UpdateUserHeadImage(fileName string, uid int64) bool {
	headImage := "http://10.224.97.223:9000/headimage/" + fileName
	result := initDB.MYSQL_DB.Model(&model.User{}).Where("id = ?", uid).Update("headimage", headImage)
	if result.Error != nil {
		return false
	}
	return true
}

func UpdateUserInfo(userInfo *model.UserInfo) error {
	result := initDB.MYSQL_DB.Model(&model.User{}).Where("id = ?", userInfo.ID).Update("nickanme", userInfo.Nickname)
	return result.Error
}

func ReleaseDynamic(dynamic *model.Dynamic) error {
	result := initDB.MYSQL_DB.Model(&model.Dynamic{}).Create(dynamic)
	return result.Error
}

func GetUserInfoByAccount(account string) (user *model.User, err error) {
	result := initDB.MYSQL_DB.Model(&model.User{}).Where("account = ?", account).First(&user)
	if result.Error != nil {
		return nil, err
	}
	return user, nil
}
