package models

type Job struct {
	Id       string `json:"id" bson:"_id"`
	Cron     string `json:"cron" bson:"cron"`
	SurveyId string `json:"surveyId" bson:"surveyId"`
	CfId     string `json:"cfId" bson:"cfId"`
	JobId    int64  `json:"jobId" bson:"jobId"`
}
