package request

type ContinuousFeedback struct {
	Id                 string              `json:"id" bson:"_id" omitempty:"true"`
	Name               string              `json:"name" bson:"name"`
	VerticalId         string              `json:"verticalId" bson:"verticalId"`
	IsCurrentlyActive  bool                `json:"isCurrentlyActive" bson:"isCurrentlyActive"`
	ResponsRate        float64             `json:"responseRate" bson:"responseRate"`
	ScheduledSurveys   []ScheduledSurvey   `json:"scheduledSurveys" bson:"scheduledSurveys" omitempty:"true"`
	EventSurveys       []EventSurvey       `json:"eventSurveys" bson:"eventSurveys" omitempty:"true"`             //Always empty for now
	IntegrationSurveys []IntegrationSurvey `json:"integrationSurveys" bson:"integrationSurveys" omitempty:"true"` //Always empty for now
}

type IntegrationSurvey struct {
}

type EventSurvey struct {
}

type ScheduledSurvey struct {
	Id                         string     `json:"id" bson:"_id" omitempty:"true"`
	ContinuousFeedbackParentId string     `json:"continuousFeedbackParentId" bson:"continuousFeedbackParentId"`
	Name                       string     `json:"name" bson:"name"`
	LastOpened                 string     `json:"lastOpened" bson:"lastOpened"`
	OpenPeriod                 string     `json:"openPeriod" bson:"openPeriod"`
	Interval                   string     `json:"interval" bson:"interval"`
	ResponseRate               float64    `json:"responseRate" bson:"responseRate"`
	Questions                  []Question `json:"questions" bson:"questions"`
}

type Question struct {
	Id                         string `json:"id" bson:"id"`
	Question                   string `json:"question" bson:"question"`
	Description                string `json:"description" bson:"description" omitempty:"true"`
	IsCalculatedInOverallScore bool   `json:"isCalculatedInOverallScore" bson:"isCalculatedInOverallScore"`
}

type Audience struct {
	Id string `json:"id" bson:"id"`
}
