package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type notificationController struct {
}

func InitNotificationController(engine *gin.Engine) {
	h := &notificationController{}

	notificationRouter := engine.Group("/sock")
	{
		notificationRouter.GET("/ws", h.notifyUser)
	}
}

func (c *notificationController) notifyUser(ctx *gin.Context) {
	log.Printf("connected to socket\n")
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("Error upgrading connection to websocket: %s", err.Error())
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing websocket connection: %s", err.Error())
		}
	}()
	for {
		err := conn.WriteJSON("Hi")
		if err != nil {
			log.Printf("Error writing message to websocket: %s", err.Error())
			return
		}
	}
}
