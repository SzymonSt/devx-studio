package main

import (
	"bytes"
	"devx/cmd/test/request"
	"devx/cmd/util"
	"devx/pkg/controllers"
	"devx/pkg/routes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	setupTestServer()
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
		message := "devx api is up and running"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	ClassicSurveyRouteController.RegisterRoutes(router)
	ContinuousFeedbackRouteController.RegisterRoutes(router)
	VerticalDataRouteController.RegisterRoutes(router)
	go func() { log.Fatal(server.Run(":" + config.ServerPort)) }()
}

//Test cases for ClassicSurveyController
// EMPTY

//---Test cases for ContinuousFeedbackController

//	1. test POST /continuousfeedback
//	2. test GET /continuousfeedback/:id
//  3. test DELETE /continuousfeedback/:id
func TestAddDeleteContinuousFeedback(t *testing.T) {
	client := http.Client{}

	method := "POST"
	url := "http://localhost:8080/api/continuousfeedback"
	body := request.ContinuousFeedback{
		VerticalId: "infrastruture",
		ScheduledSurveys: map[string]request.ScheduledSurvey{
			"1": {
				Name: "test",
			},
		},
	}
	payload, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)
	if result["message"] != "success" {
		t.Errorf("Expected status success, got %s", result["message"])
	}
}

//	1. test POST /continuousfeedback
//	2. test GET /continuousfeedback/:id
//	3. Test POST /continuousfeedback/answer
func TestAddAnswerContinuousFeedback(t *testing.T) {
	t.Skip("Skipping TestAddAnswerContinuousFeedback")
}

//---Test cases for VerticalDataController

//	1. Populate Database with synthetic answers data
//	2. Test GET /verticaldata/:verticalId/
func TestGetVerticalData(t *testing.T) {
	t.Skip("Skipping TestGetVerticalData")
}
