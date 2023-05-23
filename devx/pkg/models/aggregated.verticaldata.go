package models

type AggregateVerticalData struct {
	Id      string                     `json:"id" bson:"_id"`
	Answers []ContinuousFeedbackAnswer `json:"answers" bson:"answers"`
}
