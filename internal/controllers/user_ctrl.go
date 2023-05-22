package controllers

import (
	"glucovie/internal/models"
	"glucovie/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	service services.UserServiceImpl
}

func InitAuthController(engine *gin.Engine, service services.UserServiceImpl) {
	ac := &authController{service: service}

	router := engine.Group("/auth/user")
	{
		router.GET("/:id", ac.GetUser)
		router.POST("/register", ac.Register)
		router.POST("/login", ac.Login)
		router.DELETE("/delete/:id", ac.DeleteUser)
		router.PUT("/update/:id", ac.UpdateUser)
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var err error
	var user models.UserCredentials
	if err = ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := c.service.Login(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *authController) Register(ctx *gin.Context) {
	var user *models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := user.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.service.Register(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user was created successfully",
	})
}

func (c *authController) GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	user, err := c.service.GetUser(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *authController) DeleteUser(ctx *gin.Context) {}
func (c *authController) UpdateUser(ctx *gin.Context) {}
