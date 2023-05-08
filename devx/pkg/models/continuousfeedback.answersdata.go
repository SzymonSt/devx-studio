package models

type ContinuousFeedbackAnswersData struct {
	VerticalId   string        `json:"verticalId"`
	OverallScore Score         `json:"overallScore"`
	SurveyScores []SurveyScore `json:"surveyScores"`
}
