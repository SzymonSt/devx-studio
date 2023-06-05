package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AggregateVerticalDataByDateAndSurveyId struct {
	Id struct {
		Time     string `json:"time" bson:"time"`
		SurveyId string `json:"surveyId" bson:"surveyId"`
	} `json:"id" bson:"_id"`
	TimeDate primitive.DateTime         `json:"timeDate" bson:"timeDate"`
	Answers  []ContinuousFeedbackAnswer `json:"answers" bson:"answers"`
}

type AggregateVerticalDataByDate struct {
	Id struct {
		Time string `json:"time" bson:"time"`
	} `json:"id" bson:"_id"`
	TimeDate primitive.DateTime `json:"timeDate" bson:"timeDate"`
	Answers  []float64          `json:"answers" bson:"answers"`
}
