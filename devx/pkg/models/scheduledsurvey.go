package models

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
