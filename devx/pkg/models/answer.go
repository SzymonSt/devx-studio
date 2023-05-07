package models

type Answer struct {
	ContinuousFeedbackParentId string `json:"continuousFeedbackParentId" bson:"continuousFeedbackParentId"`
	SurveyId                   string `json:"surveyId"`
	Timestamp                  string `json:"timestamp"`
	Questions                  []struct {
		QuestionId                 string  `json:"questionId"`
		Question                   string  `json:"question"`
		Score                      float32 `json:"score"`
		IsCalculatedInOverallScore bool    `json:"isCalculatedInOverallScore"`
	} `json:"questions"`
}
