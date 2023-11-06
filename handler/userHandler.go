package handler

import (
	"chat/dal/DB"
	"chat/dal/OOS"
	"chat/model"
	util "chat/utils"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

func LoginHandler(_ context.Context, ctx *app.RequestContext) {
	//获取请求参数
	body, err := ctx.Body()
	userBasic := &model.UserBasic{}
	err = json.Unmarshal(body, userBasic)
	if err != nil {
		ctx.JSON(400, utils.H{
			"state_msg": "请求参数不了解",
		})
		return
	}
	//根据参数查询
	user, err := DB.SelectUserByAccount(userBasic.Account)
	if err != nil {
		ctx.JSON(401, utils.H{
			"state_msg": "用户名或密码错误",
		})
		return
	}

	//验证密码
	err = util.ParsingPassword(user.Password, userBasic.Password)
	if err != nil {
		ctx.JSON(401, utils.H{
			"state_msg": "用户名或密码错误",
		})
	} else {
		//生成token
		token, err := util.GenerateToken(user.ID, user.Account)
		if err != nil {
			ctx.JSON(500, utils.H{
				"state_msg": "服务器错误",
			})
			return
		}
		userID := strconv.FormatInt(user.ID, 10)
		//redisTODO
		fmt.Println(userBasic.Mac)
		err = DB.StorageUserIDAndMac(userID, userBasic.Mac)
		if err != nil {
			ctx.JSON(500, utils.H{
				"state_msg": err.Error(),
			})
			return
		}
		userInfo, err := DB.GetUserInfoById(user.ID)
		if err != nil {
			log.Print("aaaaa")
			ctx.JSON(500, utils.H{
				"state_msg": err.Error(),
			})
			return
		}
		ctx.JSON(200, utils.H{
			"state_msg": "登陆成功",
			"token":     token,
			"id":        userID,
			"Avatar":    userInfo.Headimage,
			"NickName":  userInfo.Nickname,
		})
	}
}

func RegisterHandler(_ context.Context, ctx *app.RequestContext) {
	body, err := ctx.Body()
	userBasic := &model.UserBasic{}
	err = json.Unmarshal(body, userBasic)
	if err != nil {
		ctx.JSON(500, utils.H{
			"state_msg": "请求参数不了解",
		})
		return
	}
	//加密密码
	encryptPassword, err := util.EncryptPassword(userBasic.Password)
	if err != nil {
		ctx.String(500, "服务器出错pwd")
		return
	}
	id := util.SNOW_FLAKE.GenSnowID()
	token, err := util.GenerateToken(id, userBasic.Account)
	if err != nil {
		ctx.String(500, "服务器出错token"+err.Error())
		return
	}
	//调用dal层
	err = DB.InsertUser(id, userBasic.Account, encryptPassword)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		ctx.String(401, "用户已存在")
		return
	} else if err != nil {
		ctx.String(500, err.Error())
		return
	}

	userID := strconv.FormatInt(id, 10)
	err = DB.StorageUserIDAndMac(userID, userBasic.Mac)
	if err != nil {
		ctx.JSON(500, utils.H{
			"state_msg": err.Error(),
		})
		return
	}
	userInfo, err := DB.GetUserInfoById(id)
	if err != nil {
		ctx.JSON(500, utils.H{
			"state_msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, utils.H{
		"state_msg": "注册成功",
		"token":     token,
		"id":        id,
		"Avatar":    userInfo.Headimage,
		"NickName":  userInfo.Nickname,
	})
}

func LogoutHandler(_ context.Context, ctx *app.RequestContext) {
	//获取 userID  and  userMAC
	userID := ctx.Query("mid")
	userMAC := ctx.Query("userMAC")
	//删除 redis 中的 userID---userMAC
	err := DB.DeleteUserIDAndMac(userID, userMAC)
	if err != nil {
		ctx.JSON(500, utils.H{
			"state_msg": "退出失败",
		})
		return
	}
	ctx.JSON(200, utils.H{
		"state_msg": "退出成功",
	})
}

func GetUserInfoById(_ context.Context, ctx *app.RequestContext) {
	//获取参数
	mid := ctx.Query("mid")
	fid := ctx.Query("fid")
	result := DB.GetRemarkById(mid, fid)
	session := model.Session{}
	if err := result.Decode(&session); err != nil {
		ctx.String(500, "服务器异常解析失败")
		return
	}
	//获取备注
	remark := session.Remark
	//获取好友其他信息
	fidInt64, _ := strconv.ParseInt(fid, 10, 64)
	user, err := DB.GetUserInfoById(fidInt64)
	if err != nil {
		ctx.JSON(401, utils.H{
			"state_msg": "获取userinfo失败",
		})
		return
	}
	ctx.JSON(200, utils.H{
		"remark":    remark,
		"account":   user.Account,
		"nickname":  user.Nickname,
		"headimage": user.Headimage,
	})
}
func GetUserInfo(_ context.Context, ctx *app.RequestContext) {
	query := ctx.Query("mid")
	mid, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		ctx.JSON(401, utils.H{
			"state_msg": "获取userinfo失败",
		})
		return
	}
	user, err := DB.GetUserInfoById(mid)
	if err != nil {
		ctx.JSON(500, utils.H{
			"state_msg": "获取userinfo失败",
		})
		return
	}
	ctx.JSON(200, utils.H{
		"data": user,
	})
}

func UpdateHeadImage(_ context.Context, ctx *app.RequestContext) {
	uid, _ := ctx.Get("mid")
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(500, "更新失败")
		return
	}
	fileName := fileHeader.Filename
	fileSize := fileHeader.Size
	file, err := fileHeader.Open()
	if err != nil {
		ctx.String(500, "更新失败")
		return
	}
	_, err = OOS.MINIO_CLIENT.PutObject(context.Background(), "headimage", fileName, file, fileSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		ctx.String(500, "更新失败")
		return
	}
	// 修改用户头像
	success := DB.UpdateUserHeadImage(fileName, uid.(int64))
	if success {
		ctx.JSON(200, "http://10.224.97.223:9000/headimage/"+fileName)
		return
	} else {
		ctx.String(500, "更新失败")
		return
	}
}
func UpdateUserInfo(_ context.Context, ctx *app.RequestContext) {
	userInfo := &model.UserInfo{}
	body, _ := ctx.Body()
	err := json.Unmarshal(body, userInfo)
	if err != nil {
		ctx.String(500, "更新失败")
	}
	err = DB.UpdateUserInfo(userInfo)
	if err != nil {
		ctx.String(500, "更新失败")
	}
	ctx.JSON(200, "更新成功")
}

func GetFriends(_ context.Context, ctx *app.RequestContext) {
	mid := ctx.Query("id")
	//根据用户id查询好友
	cursor, err := DB.GetFriends(mid)
	if err != nil {
		log.Print(err)
		return
	}
	defer cursor.Close(context.TODO())
	var friendList []model.Friend
	friendList = make([]model.Friend, 0)
	for cursor.Next(context.Background()) {
		var friend model.Friend
		if err := cursor.Decode(&friend); err != nil {
			ctx.String(500, "服务器异常")
			return
		}
		friendList = append(friendList, friend)
	}
	ctx.JSON(200, utils.H{
		"friendList": friendList,
	})
}

func GetUserInfoByAccount(_ context.Context, ctx *app.RequestContext) {
	account := ctx.Query("account")
	user, err := DB.GetUserInfoByAccount(account)
	if err != nil {
		log.Print(err)
		ctx.String(500, err.Error())
		return
	}
	if user == nil {
		ctx.String(402, errors.New("为查询到该用户").Error())
		return
	}

	ctx.JSON(200, utils.H{
		"user": user,
	})
}

type Fid struct {
	Fid string `json:"fid"`
	Mid string `json:"mid"`
}

func AddFriend(_ context.Context, ctx *app.RequestContext) {
	body, _ := ctx.Body()
	fid := Fid{}
	err := json.Unmarshal(body, &fid)
	if err != nil {
		log.Print(err)
		ctx.String(500, err.Error())
		return
	}
	//存储该消息
	friend_id, err := strconv.ParseInt(fid.Fid, 10, 64)
	if err != nil {
		log.Print(err)
		ctx.String(500, err.Error())
		return
	}
	my_id, err := strconv.ParseInt(fid.Mid, 10, 64)
	if err != nil {
		log.Print(err)
		ctx.String(500, err.Error())
		return
	}
	message := &model.Message{
		Mtype:    2,
		FromID:   my_id,
		ToID:     friend_id,
		Content:  "",
		Kind:     0,
		CreateAt: strconv.FormatInt(time.Now().UnixNano(), 10),
	}
	err = DB.StorageMessage(message)
	if err != nil {
		log.Print(err)
		ctx.String(500, err.Error())
		return
	}
	//转发消息
	ctx.String(200, "添加好友成功")
}
