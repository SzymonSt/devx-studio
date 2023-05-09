package main

import (
	"devx/cmd/util"
	"devx/pkg/controllers"
	"devx/pkg/routes"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setupTestServer() {
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
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	ClassicSurveyRouteController.RegisterRoutes(router)
	ContinuousFeedbackRouteController.RegisterRoutes(router)
	VerticalDataRouteController.RegisterRoutes(router)
	go log.Fatal(server.Run(":" + config.ServerPort))
}

//Test cases for ClassicSurveyController

//Test cases for ContinuousFeedbackController
//	1. test POST /continuousfeedback
//	2. test GET /continuousfeedback/:id
//	3. Test POST /continuousfeedback/answer

//Test cases for VerticalDataController
//	1. Populate Database with synthetic answers data
//	2. Test /verticaldata/:verticalId/
