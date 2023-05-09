package controllers

import (
	"devx/pkg/models"

	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/gin-gonic/gin"
	"github.com/montanaflynn/stats"
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

	scores := make([]float64, 0)
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

	surveyScores := make([]*models.SurveyScore, 0)
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

func calculateOverallScore(scores []float64) (overallScore models.Score) {
	var mean float64
	var percentile95 float64
	var percentile99 float64

	mean, _ = stats.Mean(scores)
	percentile95, _ = stats.Percentile(scores, 95)
	percentile99, _ = stats.Percentile(scores, 99)

	return models.Score{
		Mean:         mean,
		Percentile95: percentile95,
		Percentile99: percentile99,
	}
}

func calculatePerQuestionPerSurveyScore(scores []*models.ContinuousFeedbackAnswer) *models.SurveyScore {
	questionScoresTotals := make([]*models.QuestionScore, 0)
	for _, answer := range scores {
		groupedQuestionScores := map[string][]*models.ContinuousFeedbackAnswersQuestion{}
		From(answer.Questions).GroupBy(
			func(question interface{}) interface{} {
				return question.(*models.ContinuousFeedbackAnswersQuestion).QuestionId
			},
			func(question interface{}) interface{} {
				return question
			},
		).ToMap(&groupedQuestionScores)

		for qid, question := range groupedQuestionScores {
			questionScores := make([]float64, 0)
			for _, questionAnswer := range question {
				questionScores = append(questionScores, questionAnswer.Score)
			}
			questionScoresTotals = append(questionScoresTotals, &models.QuestionScore{
				QuestionId:      qid,
				QuestionContent: question[0].Question,
				Score:           calculateOverallScore(questionScores),
			})
		}

	}

	return &models.SurveyScore{
		SurveyName:     scores[0].SurveyName,
		SurveyId:       scores[0].SurveyId,
		CFId:           scores[0].ContinuousFeedbackParentId,
		CFName:         scores[0].ContinuousFeedbackName,
		QuestionScores: questionScoresTotals,
	}
}
