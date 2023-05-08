package models

type ContinuousFeedbackAnswer struct {
	VerticalId                 string                              `json:"verticalId" bson:"verticalId"`
	ContinuousFeedbackParentId string                              `json:"continuousFeedbackParentId" bson:"continuousFeedbackParentId"`
	SurveyId                   string                              `json:"surveyId" bson:"surveyId"`
	Timestamp                  string                              `json:"timestamp" bson:"timestamp"`
	Questions                  []ContinuousFeedbackAnswersQuestion `json:"questions" bson:"questions"`
}

type ContinuousFeedbackAnswersQuestion struct {
	QuestionId                 string  `json:"questionId" bson:"questionId"`
	Question                   string  `json:"question" bson:"question"`
	Score                      float64 `json:"score" bson:"score"`
	IsCalculatedInOverallScore bool    `json:"isCalculatedInOverallScore" bson:"isCalculatedInOverallScore"`
}
