package handler

import (
	"chat/dal/DB"
	"chat/model"
	util "chat/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/websocket"
	"log"
	"strconv"
	"time"
	"unsafe"
)

func WebsocketHandler(conn *websocket.Conn) {
	//循环处理消息逻辑
	for {
		//监听客户端消息
		messageType, message, err := conn.ReadMessage()
		fmt.Println(messageType)
		if err != nil {
			log.Println("read:", err)
			break
		}
		//对message进行解析
		log.Printf("recv: %s", message)
		content := model.Content{}
		err = json.Unmarshal(message, &content)
		if err != nil {
			//返回消息发送失败
			fmt.Println("json解析失败", err)
		}

		if content.Kind == 0 {
			mac := content.Data.(string)
			//判断macaddress是否合法
			//存储
			err := DB.StorageMacAndSession(mac, conn)
			if err != nil {
				fmt.Println(err)
			}
		} else if content.Kind == 1 {
			data := content.Data.(map[string]interface{})
			fromID, err := strconv.ParseInt(data["fromID"].(string), 10, 64)
			if err != nil {
				log.Println("消息解析失败")
			}
			toID, err := strconv.ParseInt(data["toID"].(string), 10, 64)
			if err != nil {
				log.Println("消息解析失败")
			}
			msg := &model.Message{
				Mtype:    int(data["mtype"].(float64)),
				FromID:   fromID,
				ToID:     toID,
				Content:  data["content"].(string),
				Kind:     int(data["kind"].(float64)),
				CreateAt: strconv.FormatInt(time.Now().UnixNano(), 10),
			}
			//判断消息类型（单聊/群聊/添加好友）
			if msg.Mtype == 0 {
				//单聊处理逻辑
				//存储消息
				//msg.CreateAt = strconv.FormatInt(time.Now().UnixNano(), 10)
				DB.StorageMessage(msg)
				//获取fromID、toID、content、kind
				//根据toID查询redis获取设备MAC地址列表
				//DB.GetUsersMac(strconv.FormatInt(msg.ToID, 10))
				//遍历MAC地址
				macSlice := DB.GetUsersMac(data["toID"].(string))
				macList := macSlice.Val()
				for _, mac := range macList {
					//获取session
					cmd := DB.GetMacAndSession(mac)
					session := cmd.Val()
					log.Println(session)
					ptr, err := strconv.ParseInt(session, 0, 64)
					if err != nil {
						log.Println(err)
					}
					//Student
					toSession := *(**websocket.Conn)(unsafe.Pointer(uintptr(ptr)))
					sendMessage, err := json.Marshal(*msg)
					if err != nil {
						log.Println(err)
					}
					//转发消息
					if toSession == nil {
						continue
					}
					err = toSession.WriteMessage(websocket.TextMessage, sendMessage)
					if err != nil {
						log.Println(err)
					}
				}
				//根据MAC地址对应的websocket连接会话，转发消息
			} else if msg.Mtype == 1 {
				//群聊处理逻辑
			} else if msg.Mtype == 2 {
				//添加好友处理逻辑
			}
		} else {

		}
	}
}

func TokenCheck(_ context.Context, ctx *app.RequestContext) {
	token := ctx.Query("token")
	user, err := util.ParseToken(token)
	if err != nil {
		ctx.String(401, "token过期")
		return
	}
	userInfo, err := DB.GetUserInfoById(user.ID)
	if err != nil {
		ctx.JSON(500, utils.H{
			"state_msg": err.Error(),
		})
		return
	}
	userID := strconv.FormatInt(user.ID, 10)
	ctx.JSON(200, utils.H{
		"id":       userID,
		"Avatar":   userInfo.Headimage,
		"NickName": userInfo.Nickname,
	})
}
