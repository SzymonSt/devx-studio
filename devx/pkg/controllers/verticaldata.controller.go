package controllers

import (
	"devx/pkg/models"

	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VerticalDataController struct {
	dbClient *mongo.Client
}

func NewVerticalDataController(dbClient *mongo.Client) VerticalDataController {
	return VerticalDataController{
		dbClient: dbClient,
	}
}

func (vdc *VerticalDataController) GetVerticalData(ctx *gin.Context) {
	verticalId := ctx.Param("verticalId")

	answers, err := vdc.collectAnswers(ctx, verticalId)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if len(answers) == 0 {
		ctx.JSON(200, nil)
		return
	}

	scores := make([]float32, 0)
	for _, answer := range answers {
		for _, question := range answer.Questions {
			if question.IsCalculatedInOverallScore {
				scores = append(scores, question.Score)
			}
		}
	}
	overallScore := calculateOverallScore(scores)

	groupedAnswers := map[string][]*models.ContinuousFeedbackAnswer{}
	From(answers).GroupBy(
		func(answer interface{}) interface{} {
			return answer.(*models.ContinuousFeedbackAnswer).SurveyId
		},
		func(answer interface{}) interface{} {
			return answer
		},
	).ToMap(&groupedAnswers)

	surveyScores := make([]models.SurveyScore, 0)
	for _, answers := range groupedAnswers {
		surveyScore := calculatePerQuestionPerSurveyScore(answers)
		surveyScores = append(surveyScores, surveyScore)
	}

	verticalData := models.ContinuousFeedbackAnswersData{
		VerticalId:   verticalId,
		OverallScore: overallScore,
		SurveyScores: surveyScores,
	}

	ctx.JSON(200, verticalData)

}

func (vdc *VerticalDataController) collectAnswers(ctx *gin.Context, verticalId string) ([]*models.ContinuousFeedbackAnswer, error) {

	var answers []*models.ContinuousFeedbackAnswer
	filter := bson.M{"verticalId": verticalId}
	answersResult, err := vdc.dbClient.Database("devx").Collection("continuousfeedbackanswers").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = answersResult.All(ctx, &answers)
	if err != nil {
		return nil, err
	}
	return answers, nil

}

func calculateOverallScore(scores []float32) (overallScore models.Score) {
	return models.Score{}
}

func calculatePerQuestionPerSurveyScore(scores []*models.ContinuousFeedbackAnswer) models.SurveyScore {
	return models.SurveyScore{}
}
