package handler

import (
	"chat/dal/DB"
	"chat/model"
	util "chat/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"log"
	"strconv"
	"time"
)

func WebsocketHandler(conn *websocket.Conn) {
	//循环处理消息逻辑
	for {
		//监听客户端消息
		_, message, err := conn.ReadMessage()
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
			data := content.Data.(map[string]interface{})
			fmt.Println("mac地址" + data["mac"].(string))
			//判断macaddress是否合法
			//存储
			DB.MacSession(data["mac"].(string), conn)
		} else if content.Kind == 1 {
			data := content.Data.(map[string]interface{})
			msg := model.Message{
				Mtype:    data["mtype"].(int),
				FromID:   data["fromID"].(int64),
				ToID:     data["toID"].(int64),
				Content:  data["content"].(string),
				Kind:     data["kind"].(int),
				CreateAt: strconv.FormatInt(time.Now().UnixNano(), 10),
			}
			//判断消息类型（单聊/群聊/添加好友）
			if msg.Mtype == 0 {
				//单聊处理逻辑
				//存储消息
				msg.CreateAt = strconv.FormatInt(time.Now().UnixNano(), 10)
				DB.StorageMessage(msg)
				//获取fromID、toID、content、kind
				//根据toID查询redis获取设备MAC地址列表
				//遍历MAC地址
				//根据MAC地址对应的websocket连接会话，转发消息
			} else if msg.Mtype == 1 {
				//群聊处理逻辑
			} else if msg.Mtype == 2 {
				//添加好友处理逻辑
			}
		} else {

		}

		//err = conn.WriteMessage(mt, message)
		//if err != nil {
		//	log.Println("write:", err)
		//	break
		//}
	}
}

func TokenCheck(_ context.Context, ctx *app.RequestContext) {
	token := ctx.Query("token")
	_, err := util.ParseToken(token)
	if err != nil {
		ctx.String(401, "token过期")
		return
	}
	ctx.String(200, "token未过期")
}
