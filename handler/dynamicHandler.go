package handler

import (
	"chat/dal/DB"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
)

func GetDynamicsHandler(_ context.Context, ctx *app.RequestContext) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	//根据page获取动态数据
	fmt.Println(page)
	dynamics, err := DB.GetDynamics(page)
	if err != nil {
		ctx.String(500, "查询出错")
		return
	}
	ctx.JSON(200, utils.H{
		"data": dynamics,
	})
}
