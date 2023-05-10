package models

type ContinuousFeedback struct {
	FeedbackId         string                       `json:"feedbackId" bson:"feedbackId" omitempty:"true"`
	VerticalId         string                       `json:"verticalId" bson:"verticalId"`
	ScheduledSurveys   map[string]ScheduledSurvey   `json:"scheduledSurveys" bson:"scheduledSurveys" omitempty:"true"`
	EventSurveys       map[string]EventSurvey       `json:"eventSurveys" bson:"eventSurveys" omitempty:"true"`             //Always empty for now
	IntegrationSurveys map[string]IntegrationSurvey `json:"integrationSurveys" bson:"integrationSurveys" omitempty:"true"` //Always empty for now
}
