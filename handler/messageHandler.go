package handler

import (
	"chat/dal/DB"
	"chat/dal/OOS"
	"chat/model"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/minio/minio-go/v7"
	"strconv"
)

func GetMessagesHandler(_ context.Context, ctx *app.RequestContext) {
	fid := ctx.Query("fid")
	frid, err := strconv.ParseInt(fid, 10, 64)
	if err != nil {
		ctx.String(401, "参数有误")
		return
	}
	mid, exists := ctx.Get("mid")
	if !exists {
		ctx.String(500, "服务器异常getToken")
		return
	}
	myid := mid.(int64)
	curser, err := DB.GetMessages(myid, frid)
	if err != nil {
		ctx.String(500, err.Error())
		return
	}
	defer curser.Close(context.Background())
	var messages []model.MessagePlus
	messages = make([]model.MessagePlus, 0)
	for curser.Next(context.Background()) {
		var message model.MessagePlus
		if err := curser.Decode(&message); err != nil {
			ctx.String(500, "服务器异常")
			return
		}
		messages = append(messages, message)
	}
	ctx.JSON(200, utils.H{
		"data": messages,
	})
}

func UploadVoice(_ context.Context, ctx *app.RequestContext) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(500, "发送失败")
		return
	}
	fileName := fileHeader.Filename
	fileSize := fileHeader.Size
	file, err := fileHeader.Open()
	if err != nil {
		ctx.String(500, "发送失败")
		return
	}
	_, err = OOS.MINIO_CLIENT.PutObject(context.Background(), "voice", fileName, file, fileSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		ctx.String(500, "发送失败")
		return
	}
	ctx.JSON(200, utils.H{
		"data": "http://10.224.97.223:9000/voice/" + fileName,
	})
}
