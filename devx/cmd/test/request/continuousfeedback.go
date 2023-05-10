package request

type ContinuousFeedback struct {
	FeedbackId         string                       `json:"feedbackId" bson:"feedbackId" omitempty:"true"`
	VerticalId         string                       `json:"verticalId" bson:"verticalId"`
	ScheduledSurveys   map[string]ScheduledSurvey   `json:"scheduledSurveys" bson:"scheduledSurveys" omitempty:"true"`
	EventSurveys       map[string]EventSurvey       `json:"eventSurveys" bson:"eventSurveys" omitempty:"true"`             //Always empty for now
	IntegrationSurveys map[string]IntegrationSurvey `json:"integrationSurveys" bson:"integrationSurveys" omitempty:"true"` //Always empty for now
}

type IntegrationSurvey struct {
}

type EventSurvey struct {
}

type ScheduledSurvey struct {
	Id                         string `json:"id" bson:"_id"`
	ContinuousFeedbackParentId string `json:"continuousFeedbackParentId" bson:"continuousFeedbackParentId"`
	Name                       string `json:"name" bson:"name"`
	LastOpened                 string `json:"lastOpened" bson:"lastOpened"`
	OpenPeriod                 string `json:"openPeriod" bson:"openPeriod"`
	Interval                   int    `json:"interval" bson:"interval"`
	Questions                  []struct {
		QuestionId                 string `json:"questionId" bson:"questionId"`
		Question                   string `json:"question" bson:"question"`
		IsCalculatedInOverallScore bool   `json:"isCalculatedInOverallScore" bson:"isCalculatedInOverallScore"`
	} `json:"questions" bson:"questions"`
}
