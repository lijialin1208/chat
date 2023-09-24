package main

import (
	"chat/dal/initDB"
	"chat/router"
	"chat/utils"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithHostPorts("10.224.94.43:80"))
	h.NoHijackConnPool = true
	router.InitRouter(h)
	router.WebSocketLink(h)
	initDB.InitMysql()
	initDB.InitRedis()
	initDB.InitMongoDB()
	utils.InitSnowFlake(0)
	h.Spin()
}
