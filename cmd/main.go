package main

import (
	"fmt"
	glucovieapi "glucovie"
	"glucovie/internal/controllers"
	"glucovie/internal/repositories"
	"glucovie/internal/services"
	"glucovie/pkg/mongodb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var database *mongo.Database

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	fmt.Println(int(time.Now().Add(-(time.Hour * 24)).Weekday()))

	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository)
	controllers.InitAuthController(router, userService)

	glucoseRepository := repositories.NewGlucoseRepository(database)
	glucoseService := services.NewGlucoseService(glucoseRepository)
	controllers.InitGlucoseController(router, glucoseService)

	eventRepository := repositories.NewEventRepository(database)
	eventService := services.NewEventService(eventRepository)
	controllers.InitEventController(router, eventService)

	srv := glucovieapi.NewHTTPServer(":8000", router)
	srv.Start()
}

func init() {
	database = mongodb.GetMongoDBConnection()
}
