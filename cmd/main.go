package main

import (
	"context"
	"fmt"
	glucovieapi "glucovie"
	"glucovie/internal/controllers"
	"glucovie/internal/models"
	"glucovie/internal/repositories"
	"glucovie/internal/services"
	"glucovie/pkg/mongodb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/mongo"
)

var database *mongo.Database

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository)
	controllers.InitAuthController(router, userService)

	glucoseRepository := repositories.NewGlucoseRepository(database)
	glucoseService := services.NewGlucoseService(glucoseRepository)
	controllers.InitGlucoseController(router, glucoseService)

	eventRepository := repositories.NewEventRepository(database)
	eventService := services.NewEventService(eventRepository)
	controllers.InitEventController(router, eventService)

	controllers.InitNotificationController(router)

	go func() {
		users, err := userRepository.GetAllUsersIDs(context.Background())
		if err != nil {
			panic(err)
		}

		runScheduler(users, glucoseService)
	}()

	srv := glucovieapi.NewHTTPServer(":8000", router)
	srv.Start()
}

func init() {
	database = mongodb.GetMongoDBConnection()
}

func runScheduler(users []models.User, svc services.GlucoseServiceImpl) {
	fmt.Println("Cron Started !")
	s := gocron.NewScheduler(time.UTC)
	job, err := s.Every(1).Week().Do(func() {
		for _, v := range users {
			if v.EmailNotification {
				res, err := svc.GetWeekGlucoseLevel(v.ID.Hex())
				if err != nil {
					panic(err)
				}

				err = services.SendEmail(v.DocMail, res)
				if err != nil {
					panic(err)
				}
			}
		}
	})
	if err != nil {
		panic(err)
	}
	err = job.Error()
	if err != nil {
		panic(err)
	}
	s.StartAsync()
}
