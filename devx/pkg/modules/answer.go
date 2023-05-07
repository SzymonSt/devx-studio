package modules

type Answer struct {
	SurveyId  string `json:"surveyId"`
	Timestamp string `json:"timestamp"`
	Questions []struct {
		Question                   string  `json:"question"`
		Score                      float32 `json:"score"`
		IsCalculatedInOverallScore bool    `json:"isCalculatedInOverallScore"`
	} `json:"questions"`
}
