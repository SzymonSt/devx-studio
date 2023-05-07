package models

type AnswersData struct {
	VerticalId   string `json:"verticalId"`
	OverallScore []struct {
		Score []float32 `json:"score"`
	}
	Surveys []struct {
		CFId     string   `json:"cfId"`
		SurveyId string   `json:"surveyId"`
		Answers  []Answer `json:"answers"`
	}
}
