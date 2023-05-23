package controllers

import (
	"devx/pkg/models"
	"fmt"

	//. "github.com/ahmetalpbalkan/go-linq"
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
	//interval := ctx.Param("interval")
	verticalId := ctx.Param("verticalId")

	answers, err := vdc.collectAnswers(ctx, verticalId)
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
	fmt.Printf("answers: %+v\n", answers)

	ctx.JSON(200, answers)

}

func (vdc *VerticalDataController) collectAnswers(ctx *gin.Context, verticalId string) ([]*models.AggregateVerticalData, error) {
	var aggAnswers []*models.AggregateVerticalData
	matchVerticalIdGroupDate := bson.A{
		bson.D{{"$match", bson.D{{"verticalId", "infrastruture"}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id",
						bson.D{
							{"$dateToString",
								bson.D{
									{"format", "%Y-%m-%d"},
									{"date", "$timestamp"},
								},
							},
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
