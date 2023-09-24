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
	userGroup.Use(middleware.AuthMiddleware()).GET("/getUserInfoById", handler.GetUserInfoById)
	userGroup.Use(middleware.AuthMiddleware()).GET("/getUserInfo", handler.GetUserInfo)

	messageGroup := hertz.Group("/api/message").Use(middleware.AuthMiddleware())
	messageGroup.Use(middleware.AuthMiddleware()).GET("/getMessages", handler.GetMessagesHandler)

	dynamicGroup := hertz.Group("/api/dynamic")
	dynamicGroup.Use(middleware.AuthMiddleware()).GET("/getDynamics", handler.GetDynamicsHandler)
}