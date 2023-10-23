package middleware

import (
	"chat/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"strings"
)

func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		tokenBytes := c.GetHeader("Authorization")
		token := string(tokenBytes)
		split := strings.Split(token, " ")[1]
		parseToken, err := utils.ParseToken(split)
		if split == "" || err != nil {
			c.AbortWithMsg("未登录", 301)
		} else {
			c.Set("mid", parseToken.ID)
			c.Next(ctx)
		}
		defer func() {
			err := recover()
			if err != nil {
				log.Print(err)
				c.AbortWithMsg("token有误", 301)
			}
		}()
	}
}
