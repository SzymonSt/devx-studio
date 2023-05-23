package models

type ContinuousFeedbackAnswersData struct {
	VerticalId   string         `json:"verticalId"`
	OverallScore []*ScoreData   `json:"overallScore"`
	SurveyScores []*SurveyScore `json:"surveyScores"`
}
