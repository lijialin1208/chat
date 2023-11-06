package handler

import (
	"chat/dal/DB"
	"chat/dal/OOS"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/minio/minio-go/v7"
	"log"
	"strconv"
)

func GetDynamicsHandler(_ context.Context, ctx *app.RequestContext) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	//根据page获取动态数据
	dynamics, err := DB.GetDynamics(page)
	if err != nil {
		ctx.String(500, "查询出错")
		return
	}
	fmt.Println(dynamics)
	ctx.JSON(200, utils.H{
		"data": dynamics,
	})
}
func ReleaseDynamicHandler(_ context.Context, ctx *app.RequestContext) {
	multipartForm, err := ctx.MultipartForm()
	if err != nil {
		log.Print(err)
		ctx.String(500, "发布失败")
		return
	}
	files := (multipartForm.File)["file"]

	for _, file := range files {
		filename := file.Filename
		log.Println(filename)
		filesize := file.Size
		open, err := file.Open()
		if err != nil {
			ctx.String(500, "发布失败")
			return
		}
		_, err = OOS.MINIO_CLIENT.PutObject(context.Background(), "dynamic", filename, open, filesize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
		if err != nil {
			ctx.String(500, "发布失败")
			return
		}
	}
	//err = DB.ReleaseDynamic(dynamic)
	//if err != nil {
	//	ctx.JSON(500, utils.H{
	//		"state_msg": "数据库新增失败",
	//	})
	//	return
	//}
	ctx.JSON(200, utils.H{
		"state_msg": "发布成功",
	})
}
