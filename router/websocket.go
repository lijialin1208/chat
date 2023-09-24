package router

import (
	"chat/handler"
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/websocket"
)

var upgrader = websocket.HertzUpgrader{
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
} // use default options

func echo(_ context.Context, c *app.RequestContext) {
	err := upgrader.Upgrade(c, handler.WebsocketHandler)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
}

func WebSocketLink(h *server.Hertz) {
	h.GET("/echo", echo)
}
