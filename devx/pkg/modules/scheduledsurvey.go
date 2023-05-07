package modules

type ScheduledSurvey struct {
	Id                         string `json:"id" bson:"_id"`
	ContinuousFeedbackParentId string `json:"continuousFeedbackParentId" bson:"continuousFeedbackParentId"`
	Name                       string `json:"name" bson:"name"`
	LastOpened                 string `json:"lastOpened" bson:"lastOpened"`
	OpenPeriod                 string `json:"openPeriod"`
	Interval                   int    `json:"interval"`
	Questions                  []struct {
		Question                   string `json:"question"`
		IsCalculatedInOverallScore bool   `json:"isCalculatedInOverallScore"`
	} `json:"questions"`
}
