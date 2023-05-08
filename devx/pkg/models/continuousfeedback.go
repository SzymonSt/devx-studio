package models

type ContinuousFeedback struct {
	FeedbackId         string                       `json:"feedbackId" bson:"feedbackId"`
	VerticalId         string                       `json:"verticalId" bson:"verticalId"`
	ScheduledSurveys   map[string]ScheduledSurvey   `json:"scheduledSurveys" bson:"scheduledSurveys"`
	EventSurveys       map[string]EventSurvey       `json:"eventSurveys" bson:"eventSurveys"`             //Always empty for now
	IntegrationSurveys map[string]IntegrationSurvey `json:"integrationSurveys" bson:"integrationSurveys"` //Always empty for now
}
