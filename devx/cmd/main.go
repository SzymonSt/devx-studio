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
	ClassicSurveyController := controllers.NewClassicSurveyController(dbClient)
	ClassicSurveyRouteController := routes.NewClassicSurveyRouteController(ClassicSurveyController)

	ContinuousFeedbackController := controllers.NewContinuousFeedbackController(dbClient)
	ContinuousFeedbackRouteController := routes.NewContinuousFeedbackRouteController(ContinuousFeedbackController)

	VerticalDataController := controllers.NewVerticalDataController(dbClient)
	VerticalDataRouteController := routes.NewVerticalDataRoutes(VerticalDataController)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthz", func(ctx *gin.Context) {
		message := "devx api is up and running"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	ClassicSurveyRouteController.RegisterRoutes(router)
	ContinuousFeedbackRouteController.RegisterRoutes(router)
	VerticalDataRouteController.RegisterRoutes(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
