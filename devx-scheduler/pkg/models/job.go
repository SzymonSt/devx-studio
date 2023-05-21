package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	Id       primitive.ObjectID `json:"id" bson:"_id" omitempty:"true"`
	Cron     string             `json:"cron" bson:"cron"`
	SurveyId string             `json:"surveyId" bson:"surveyId"`
	CfId     string             `json:"cfId" bson:"cfId"`
	JobId    int64              `json:"jobId" bson:"jobId"`
}
