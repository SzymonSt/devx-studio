package modules

type ContinuousFeedback struct {
	FeedbackId       string            `json:"feedbackId" bson:"feedbackId"`
	VerticalId       string            `json:"verticalId" bson:"verticalId"`
	ScheduledSurveys []ScheduledSurvey `json:"scheduledSurveys" bson:"scheduledSurveys"`
}
