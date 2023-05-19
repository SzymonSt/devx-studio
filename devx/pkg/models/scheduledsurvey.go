package models

type ScheduledSurvey struct {
	Id                         string  `json:"id" bson:"_id"`
	ContinuousFeedbackParentId string  `json:"continuousFeedbackParentId" bson:"continuousFeedbackParentId"`
	Name                       string  `json:"name" bson:"name"`
	LastOpened                 string  `json:"lastOpened" bson:"lastOpened"`
	OpenPeriod                 string  `json:"openPeriod" bson:"openPeriod"`
	Interval                   string  `json:"interval" bson:"interval"`
	ResponseRate               float64 `json:"responseRate" bson:"responseRate"`
	Audience                   []struct {
		Id string `json:"id" bson:"id"`
	} `json:"audience" bson:"audience"`
	Questions []struct {
		Id                         string `json:"id" bson:"id"`
		Question                   string `json:"question" bson:"question"`
		Description                string `json:"description" bson:"description" omitempty:"true"`
		IsCalculatedInOverallScore bool   `json:"isCalculatedInOverallScore" bson:"isCalculatedInOverallScore"`
	} `json:"questions" bson:"questions"`
}
