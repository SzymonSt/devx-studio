package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AggregateVerticalData struct {
	Id struct {
		Time     string `json:"time" bson:"time"`
		SurveyId string `json:"surveyId" bson:"surveyId"`
	} `json:"id" bson:"_id"`
	TimeDate primitive.DateTime         `json:"timeDate" bson:"timeDate"`
	Answers  []ContinuousFeedbackAnswer `json:"answers" bson:"answers"`
}
