package handler

import (
	"chat/dal/DB"
	"chat/model"
	util "chat/utils"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"gorm.io/gorm"
	"strconv"
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
		ctx.JSON(200, utils.H{
			"state_msg": "登陆成功",
			"token":     token,
			"id":        userID,
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
	ctx.JSON(200, utils.H{
		"state_msg": "注册成功",
		"token":     token,
		"id":        id,
	})
}

func GetUserInfoById(_ context.Context, ctx *app.RequestContext) {
	//获取参数
	mid, exists := ctx.Get("mid")
	if !exists {
		ctx.String(500, "服务器异常getToken")
		return
	}
	myid := mid.(int64)
	query := ctx.Query("fid")
	fid, err2 := strconv.ParseInt(query, 10, 64)
	if err2 != nil {
		ctx.String(401, "参数有误")
		return
	}
	result := DB.GetRemarkById(myid, fid)
	session := model.Session{}
	if err := result.Decode(&session); err != nil {
		ctx.String(500, "服务器异常解析失败")
		return
	}
	//获取备注
	remark := session.Remark
	//获取好友其他信息
	user, err := DB.GetUserInfoById(fid)
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
