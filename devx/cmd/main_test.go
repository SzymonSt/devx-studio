package main

import (
	"bytes"
	"devx/cmd/test/request"
	"devx/cmd/util"
	"devx/pkg/controllers"
	"devx/pkg/routes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	setupTestServer()
	client := http.Client{}
	method := "GET"
	url := "http://localhost:8080/api/healthz"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	for {
		res, err := client.Do(req)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		if res.StatusCode == 200 {
			break
		}
	}

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
	corsConfig.AllowOrigins = []string{"*", config.ClientOrigin}
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
func TestCreateDeleteContinuousFeedback(t *testing.T) {
	client := http.Client{}

	//1. test POST /continuousfeedback
	method := "POST"
	url := "http://localhost:8080/api/continuousfeedback"
	body := request.ContinuousFeedback{
		Name:       "Infra continuous feedback",
		VerticalId: "infrastruture",
		ScheduledSurveys: map[string]request.ScheduledSurvey{
			"1": {
				Name:       "IaC Survey",
				LastOpened: "",
				OpenPeriod: "48h",
				Interval:   "0 0 1 * *", //random cron
				Questions: []request.Question{
					{
						QuestionId:                 "1",
						Question:                   "How would you rate coverage of IaC modules on most cloud resources provided by Platform Engineering Team?",
						IsCalculatedInOverallScore: true,
					},
					{
						QuestionId:                 "2",
						Question:                   "How would you rate ease of setting up and providing parameters for IaC modules?",
						IsCalculatedInOverallScore: true,
					},
					{
						QuestionId:                 "3",
						Question:                   "How would you rate ease of making changes to development and production environments using IaC modules?",
						IsCalculatedInOverallScore: true,
					},
				},
			},
			"2": {
				Name:       "Infrastructure observability survey",
				LastOpened: "",
				OpenPeriod: "72h",
				Interval:   "0 0 1 * *", //random cron
				Questions: []request.Question{
					{
						QuestionId:                 "1",
						Question:                   "How would you rate quality of infrastructre alerts?",
						IsCalculatedInOverallScore: true,
					},
					{
						QuestionId:                 "2",
						Question:                   "How would you rate ease of debugging infrastructure issues?",
						IsCalculatedInOverallScore: true,
					},
					{
						QuestionId:                 "3",
						Question:                   "How would you rate guides for infrastructure troubleshooting provided by Platform Engineering Team?",
						IsCalculatedInOverallScore: false,
					},
				},
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

	var createResult map[string]interface{}
	json.NewDecoder(res.Body).Decode(&createResult)
	if createResult["message"] != "success" {
		t.Errorf("Expected status success, got %s", createResult["message"])
	}

	//2. test GET /continuousfeedback/:id
	method = "GET"
	url = fmt.Sprintf("http://localhost:8080/api/continuousfeedback/%s", createResult["id"])
	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	res, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}
	var getResult map[string]interface{}
	json.NewDecoder(res.Body).Decode(&getResult)
	if getResult["verticalId"] != "infrastruture" {
		t.Errorf("Expected vertical infrastruture, got %s", getResult["Vertical"])
	}

	//3. test DELETE /continuousfeedback/:id
	method = "DELETE"
	url = fmt.Sprintf("http://localhost:8080/api/continuousfeedback/%s", createResult["id"])
	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	res, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}

}

//	1. test POST /continuousfeedback
//	2. test GET /continuousfeedback/:id
//	3. Test POST /continuousfeedback/answer
func TestAddAnswerContinuousFeedback(t *testing.T) {
	client := http.Client{}

	//1. test POST /continuousfeedback
	method := "POST"
	url := "http://localhost:8080/api/continuousfeedback"
	body := request.ContinuousFeedback{
		Name:       "Infra continuous feedback",
		VerticalId: "infrastruture",
		ScheduledSurveys: map[string]request.ScheduledSurvey{
			"1": {
				Name:       "IaC Survey",
				LastOpened: "",
				OpenPeriod: "48h",
				Interval:   "0 0 1 * *", //random cron
				Questions: []request.Question{
					{
						QuestionId:                 "1",
						Question:                   "How would you rate coverage of IaC modules on most cloud resources provided by Platform Engineering Team?",
						IsCalculatedInOverallScore: true,
					},
					{
						QuestionId:                 "2",
						Question:                   "How would you rate ease of setting up and providing parameters for IaC modules?",
						IsCalculatedInOverallScore: true,
					},
					{
						QuestionId:                 "3",
						Question:                   "How would you rate ease of making changes to development and production environments using IaC modules?",
						IsCalculatedInOverallScore: true,
					},
				},
			},
			"2": {
				Name:       "Infrastructure observability survey",
				LastOpened: "",
				OpenPeriod: "72h",
				Interval:   "0 0 1 * *", //random cron
				Questions: []request.Question{
					{
						QuestionId:                 "1",
						Question:                   "How would you rate quality of infrastructre alerts?",
						IsCalculatedInOverallScore: true,
					},
					{
						QuestionId:                 "2",
						Question:                   "How would you rate ease of debugging infrastructure issues?",
						IsCalculatedInOverallScore: true,
					},
					{
						QuestionId:                 "3",
						Question:                   "How would you rate guides for infrastructure troubleshooting provided by Platform Engineering Team?",
						IsCalculatedInOverallScore: false,
					},
				},
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

	var createResult map[string]interface{}
	json.NewDecoder(res.Body).Decode(&createResult)
	if createResult["message"] != "success" {
		t.Errorf("Expected status success, got %s", createResult["message"])
	}

	//2. test GET /continuousfeedback/:id
	method = "GET"
	url = fmt.Sprintf("http://localhost:8080/api/continuousfeedback/%s", createResult["id"])
	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	res, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}
	var getResult map[string]interface{}
	json.NewDecoder(res.Body).Decode(&getResult)
	if getResult["verticalId"] != "infrastruture" {
		t.Errorf("Expected vertical infrastruture, got %s", getResult["verticalId"])
	}

	//3. Test POST /continuousfeedback/answer
	method = "POST"
	url = "http://localhost:8080/api/continuousfeedback/answer"
	bodyAnswer := request.ContinuousFeedbackAnswer{
		VerticalId:                 getResult["verticalId"].(string),
		ContinuousFeedbackParentId: getResult["id"].(string),
		ContinuousFeedbackName:     getResult["name"].(string),
		SurveyId:                   "1",
		SurveyName:                 getResult["scheduledSurveys"].(map[string]interface{})["1"].(map[string]interface{})["name"].(string),
		Timestamp:                  time.Now().UTC().Format(time.RFC3339),
		Questions: []*request.ContinuousFeedbackAnswersQuestion{
			{
				QuestionId:                 "1",
				Question:                   "How would you rate coverage of IaC modules on most cloud resources provided by Platform Engineering Team?",
				Score:                      4.5,
				IsCalculatedInOverallScore: true,
			},
			{
				QuestionId:                 "2",
				Question:                   "How would you rate ease of setting up and providing parameters for IaC modules?",
				Score:                      3.5,
				IsCalculatedInOverallScore: true,
			},
			{
				QuestionId:                 "3",
				Question:                   "How would you rate ease of making changes to development and production environments using IaC modules?",
				Score:                      4.0,
				IsCalculatedInOverallScore: true,
			},
		},
	}
	payload, err = json.Marshal(bodyAnswer)
	if err != nil {
		panic(err)
	}
	req, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}
	var answerResult map[string]interface{}
	json.NewDecoder(res.Body).Decode(&answerResult)
	if answerResult["message"] != "success" {
		t.Errorf("Expected status success, got %s", answerResult["message"])
	}
}

//---Test cases for VerticalDataController

//	1. Populate Database with synthetic answers data
//	2. Test GET /verticaldata/:verticalId/
func TestGetVerticalData(t *testing.T) {
	client := http.Client{}
	jsonFileCf, err := os.Open("test/data/cfs.json")
	if err != nil {
		panic(err)
	}
	defer jsonFileCf.Close()

	byteValueCf, _ := ioutil.ReadAll(jsonFileCf)
	var cfs []request.ContinuousFeedback
	var cfId string
	json.Unmarshal(byteValueCf, &cfs)
	for _, cf := range cfs {
		payload, err := json.Marshal(cf)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/api/continuousfeedback", bytes.NewBuffer(payload))
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
		var createResult map[string]interface{}
		json.NewDecoder(res.Body).Decode(&createResult)

		cfId = createResult["id"].(string)
	}

	jsonFileAnswer, err := os.Open("test/data/answers.json")
	if err != nil {
		panic(err)
	}
	defer jsonFileAnswer.Close()

	byteValueAnswer, _ := ioutil.ReadAll(jsonFileAnswer)
	var answers []request.ContinuousFeedbackAnswer
	json.Unmarshal(byteValueAnswer, &answers)
	for _, answer := range answers {
		payload, err := json.Marshal(answer)
		if err != nil {
			panic(err)
		}
		answer.ContinuousFeedbackParentId = cfId
		answer.Timestamp = time.Now().UTC().Format(time.RFC3339)
		req, err := http.NewRequest("POST", "http://localhost:8080/api/continuousfeedback/answer", bytes.NewBuffer(payload))
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
	}

	//2. Test GET /verticaldata/:verticalId/
	method := "GET"
	url := "http://localhost:8080/api/verticaldata/infrastruture"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}
	var verticalDataResult map[string]interface{}
	json.NewDecoder(res.Body).Decode(&verticalDataResult)
	fmt.Printf("verticalDataResult: %v\n", verticalDataResult)
}
