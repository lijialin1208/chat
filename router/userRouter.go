package router

import (
	"chat/handler"
	"chat/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter(hertz *server.Hertz) {
	hertz.GET("/api/tokenCheck", handler.TokenCheck)
	userGroup := hertz.Group("/api/user")
	userGroup.POST("/login", handler.LoginHandler)
	userGroup.POST("/register", handler.RegisterHandler)
	userGroup.Use(middleware.AuthMiddleware()).GET("/logout", handler.LogoutHandler)
	userGroup.Use(middleware.AuthMiddleware()).GET("/getUserInfoById", handler.GetUserInfoById)
	userGroup.Use(middleware.AuthMiddleware()).GET("/getFriends", handler.GetFriends)
	userGroup.Use(middleware.AuthMiddleware()).GET("/getUserInfo", handler.GetUserInfo)
	userGroup.Use(middleware.AuthMiddleware()).POST("/updateHeadImage", handler.UpdateHeadImage)
	userGroup.Use(middleware.AuthMiddleware()).POST("/updateUserInfo", handler.UpdateUserInfo)

	messageGroup := hertz.Group("/api/message").Use(middleware.AuthMiddleware())
	messageGroup.Use(middleware.AuthMiddleware()).GET("/getMessages", handler.GetMessagesHandler)

	dynamicGroup := hertz.Group("/api/dynamic")
	dynamicGroup.Use(middleware.AuthMiddleware()).GET("/getDynamics", handler.GetDynamicsHandler)
}
