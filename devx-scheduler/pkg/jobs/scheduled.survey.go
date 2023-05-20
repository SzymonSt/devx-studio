package jobs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ProcessScheduledSurvey(cfId string, surveyId string) {
	client := http.Client{}

	method := "GET"
	url := fmt.Sprintf("http://localhost:8080/api/continuousfeedback/%s", cfId)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if res.StatusCode != 200 {
		fmt.Printf("Expected status 200, got %d", res.StatusCode)
	}
	var getResult map[string]interface{}
	json.NewDecoder(res.Body).Decode(&getResult)
	var survey map[string]interface{}
	for _, s := range getResult["scheduledSurveys"].([]interface{}) {
		if s.(map[string]interface{})["surveyId"] == surveyId {
			survey = s.(map[string]interface{})
			break
		}
	}
	_ = survey["audience"]
}
