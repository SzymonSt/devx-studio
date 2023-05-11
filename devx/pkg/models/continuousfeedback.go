package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContinuousFeedback struct {
	Id                 primitive.ObjectID           `json:"id" bson:"_id" omitempty:"true"`
	Name               string                       `json:"name" bson:"name"`
	VerticalId         string                       `json:"verticalId" bson:"verticalId"`
	ScheduledSurveys   map[string]ScheduledSurvey   `json:"scheduledSurveys" bson:"scheduledSurveys" omitempty:"true"`
	EventSurveys       map[string]EventSurvey       `json:"eventSurveys" bson:"eventSurveys" omitempty:"true"`             //Always empty for now
	IntegrationSurveys map[string]IntegrationSurvey `json:"integrationSurveys" bson:"integrationSurveys" omitempty:"true"` //Always empty for now
}
