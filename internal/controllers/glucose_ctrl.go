package controllers

import (
	"glucovie/internal/models"
	"glucovie/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type glucoseController struct {
	service services.GlucoseServiceImpl
}

func InitGlucoseController(engine *gin.Engine, service services.GlucoseServiceImpl) {
	ac := &glucoseController{service: service}

	router := engine.Group("/glucose")
	{
		// router.GET("/:id", ac.GetUser)
		router.POST("/save", ac.SaveGlucoseLevel)
	}
}

func (c *glucoseController) SaveGlucoseLevel(ctx *gin.Context) {
	var err error
	var model models.GlucoseLevel
	if err = ctx.BindJSON(&model); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

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
