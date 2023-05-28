package controllers

import (
	"glucovie/internal/middleware"
	"glucovie/internal/models"
	"glucovie/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type glucoseController struct {
	service services.GlucoseServiceImpl
}

func InitGlucoseController(engine *gin.Engine, service services.GlucoseServiceImpl) {
	ac := &glucoseController{service: service}

	router := engine.Group("/glucose").Use(middleware.AuthMiddleware)
	{
		router.GET("/week", ac.GetWeekGlucoseLevel)
		router.POST("/save", ac.SaveGlucoseLevel)
	}
}

func (c *glucoseController) SaveGlucoseLevel(ctx *gin.Context) {
	var err error
	var model models.GlucoseLevel
	if err = ctx.BindJSON(&model); err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	model.UserID = ctx.GetString("user_id")

	err = c.service.SaveGlucoseLevel(&model)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully save glucose level",
	})
}

func (c *glucoseController) GetWeekGlucoseLevel(ctx *gin.Context) {
	userID := ctx.GetString("user_id")

	resp, err := c.service.GetWeekGlucoseLevel(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully get glucose level",
		"data":    resp,
	})
}
