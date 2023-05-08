package models

type Score struct {
	Mean         float32 `json:"mean"`
	Percentile95 float32 `json:"percentile95"`
	Percentile99 float32 `json:"percentile99"`
}

type SurveyScore struct {
	SurveyName     string          `json:"surveyName"`
	SurveyId       string          `json:"surveyId"`
	CFId           string          `json:"cfId"`
	CFName         string          `json:"cfName"`
	QuestionScores []QuestionScore `json:"questionScores"`
}

type QuestionScore struct {
	QuestionId      string `json:"questionId"`
	QuestionContent string `json:"questionContent"`
	Score           `json:"score"`
}
