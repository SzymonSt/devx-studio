package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContinuousFeedbackAnswer struct {
	Id                         primitive.ObjectID                   `json:"id" bson:"_id" omitempty:"true"`
	VerticalId                 string                               `json:"verticalId" bson:"verticalId"`
	ContinuousFeedbackParentId string                               `json:"continuousFeedbackParentId" bson:"continuousFeedbackParentId"`
	ContinuousFeedbackName     string                               `json:"continuousFeedbackName" bson:"continuousFeedbackName"`
	SurveyId                   string                               `json:"surveyId" bson:"surveyId"`
	SurveyName                 string                               `json:"surveyName" bson:"surveyName"`
	Timestamp                  primitive.DateTime                   `json:"timestamp" bson:"timestamp"`
	Questions                  []*ContinuousFeedbackAnswersQuestion `json:"questions" bson:"questions"`
}

type ContinuousFeedbackAnswersQuestion struct {
	QuestionId                 string  `json:"questionId" bson:"questionId"`
	Question                   string  `json:"question" bson:"question"`
	Score                      float64 `json:"score" bson:"score"`
	IsCalculatedInOverallScore bool    `json:"isCalculatedInOverallScore" bson:"isCalculatedInOverallScore"`
}
