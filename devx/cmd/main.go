package main

import (
	"devx/cmd/util"
	"devx/pkg/controllers"
	"devx/pkg/routes"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		server = gin.Default()
	)
	config, err := util.LaodConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbClient := util.ConnectDB(config.DBUri)
	ScheduledSurveyController := controllers.NewScheduledSurveyController(dbClient)
	ScheduledSurveyRouteController := routes.NewScheduledSurveyRouteController(ScheduledSurveyController)

	ContinuousFeedbackController := controllers.NewContinuousFeedbackController(dbClient)
	ContinuousFeedbackRouteController := routes.NewContinuousFeedbackRouteController(ContinuousFeedbackController)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthz", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	ScheduledSurveyRouteController.RegisterRoutes(router)
	ContinuousFeedbackRouteController.RegisterRoutes(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
