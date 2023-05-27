package controllers

import (
	"glucovie/internal/middleware"
	"glucovie/internal/models"
	"glucovie/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type eventController struct {
	service services.EventServiceImpl
}

func InitEventController(engine *gin.Engine, service services.EventServiceImpl) {
	ac := &eventController{service: service}

	router := engine.Group("/event").Use(middleware.AuthMiddleware)
	{
		router.POST("/save", ac.saveEvent)
	}
}

func (c *eventController) saveEvent(ctx *gin.Context) {
	var err error
	var model models.EventModel
	if err = ctx.BindJSON(&model); err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	model.UserID = ctx.GetString("user_id")

	err = c.service.SaveEvent(&model)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully save event",
	})
}

func (c *eventController) GetEvents(ctx *gin.Context) {
	resp, err := c.service.GetEvents(ctx.GetString("user_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully get events",
		"data":    resp,
	})
}
