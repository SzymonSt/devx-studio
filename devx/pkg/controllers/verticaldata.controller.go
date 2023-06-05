package controllers

import (
	"devx/pkg/models"
	"fmt"

	//. "github.com/ahmetalpbalkan/go-linq"
	"github.com/gin-gonic/gin"
	"github.com/montanaflynn/stats"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gonum.org/v1/gonum/stat"
)

type VerticalDataController struct {
	dbClient *mongo.Client
}

func NewVerticalDataController(dbClient *mongo.Client) VerticalDataController {
	return VerticalDataController{
		dbClient: dbClient,
	}
}

func (vdc *VerticalDataController) GetOverallVerticalData(ctx *gin.Context) {
	//interval := ctx.Param("interval")
	verticalId := ctx.Param("verticalId")

	answers, err := vdc.collectVerticalAnswersGroupedByDate(ctx, verticalId)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	if len(answers) == 0 {
		ctx.JSON(200, nil)
		return
	}
	score := calculateOverallScore(answers)

	ctx.JSON(200, score)

}
func (vdc *VerticalDataController) GetDetailedVerticalData(ctx *gin.Context) {
	//interval := ctx.Param("interval")
	verticalId := ctx.Param("verticalId")

	answers, err := vdc.collectVerticalAnswersGroupedByDateAndSurveyId(ctx, verticalId)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	if len(answers) == 0 {
		ctx.JSON(200, nil)
		return
	}
	//score := calculateSurveysScore(answers)

	ctx.JSON(200, answers)
}

func (vdc *VerticalDataController) collectVerticalAnswersGroupedByDateAndSurveyId(ctx *gin.Context, verticalId string) ([]*models.AggregateVerticalDataByDateAndSurveyId, error) {
	var aggAnswers []*models.AggregateVerticalDataByDateAndSurveyId
	matchVerticalIdGroupDateAndSurveyId := bson.A{
		bson.D{{"$match", bson.D{{"verticalId", "infrastruture"}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id",
						bson.D{
							{"time",
								bson.D{
									{"$dateToString",
										bson.D{
											{"format", "%Y-%m-%d"},
											{"date", "$timestamp"},
										},
									},
								},
							},
							{"surveyId", "$surveyId"},
						},
					},
					{"answers",
						bson.D{
							{"$push",
								bson.D{
									{"surveyId", "$surveyId"},
									{"continuousFeedbackName", "$continuousFeedbackName"},
									{"surveyName", "$surveyName"},
									{"questions", "$questions"},
									{"timestamp", "$timestamp"},
									{"continuousFeedbackParentId", "$continuousFeedbackParentId"},
									{"verticalId", "$verticalId"},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$set",
				bson.D{
					{"timeDate",
						bson.D{
							{"$dateFromString",
								bson.D{
									{"dateString", "$_id.time"},
									{"format", "%Y-%m-%d"},
								},
							},
						},
					},
				},
			},
		},
		bson.D{{"$sort", bson.D{{"timeDate", 1}}}},
	}
	answersResult, err := vdc.dbClient.Database("devx").Collection("continuousfeedbackanswers").Aggregate(ctx, matchVerticalIdGroupDateAndSurveyId)
	if err != nil {
		return nil, err
	}
	err = answersResult.All(ctx, &aggAnswers)
	if err != nil {
		return nil, err
	}
	return aggAnswers, nil

}

func (vdc *VerticalDataController) collectVerticalAnswersGroupedByDate(ctx *gin.Context, verticalId string) ([]*models.AggregateVerticalDataByDate, error) {
	var aggAnswers []*models.AggregateVerticalDataByDate
	matchVerticalIdGroupDate := bson.A{
		bson.D{{"$match", bson.D{{"verticalId", "infrastruture"}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id",
						bson.D{
							{"time",
								bson.D{
									{"$dateToString",
										bson.D{
											{"format", "%Y-%m-%d"},
											{"date", "$timestamp"},
										},
									},
								},
							},
						},
					},
					{"answers", bson.D{{"$push", "$questions"}}},
				},
			},
		},
		bson.D{{"$unwind", "$answers"}},
		bson.D{{"$unwind", "$answers"}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$_id"},
					{"questionAnswers", bson.D{{"$push", "$answers"}}},
				},
			},
		},
		bson.D{{"$project", bson.D{{"answers", "$questionAnswers.score"}}}},
		bson.D{{"$sort", bson.D{{"_id", 1}}}},
	}
	answersResult, err := vdc.dbClient.Database("devx").Collection("continuousfeedbackanswers").Aggregate(ctx, matchVerticalIdGroupDate)
	if err != nil {
		return nil, err
	}
	err = answersResult.All(ctx, &aggAnswers)
	if err != nil {
		return nil, err
	}
	return aggAnswers, nil
}

func calculateOverallScore(answers []*models.AggregateVerticalDataByDate) []*models.ScoreData {
	scores := make([]*models.ScoreData, 0)
	for _, answer := range answers {
		p95, err := stats.Percentile(answer.Answers, 95)
		if err != nil {
			fmt.Println(err)
		}
		p99, err := stats.Percentile(answer.Answers, 99)
		if err != nil {
			fmt.Println(err)
		}
		score := models.Score{
			Mean:         stat.Mean(answer.Answers, nil),
			Percentile95: p95,
			Percentile99: p99,
		}
		scoreData := &models.ScoreData{
			Timestamp: answer.Id.Time,
			Score:     score,
		}
		scores = append(scores, scoreData)
	}
	return scores
}

func calculateSurveysScore(answers []*models.AggregateVerticalDataByDateAndSurveyId) []*models.SurveyScore {
	return nil
}
