package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScheduledSurveyController struct {
	dbClient *mongo.Client
}

func NewScheduledSurveyController(dbClient *mongo.Client) ScheduledSurveyController {
	return ScheduledSurveyController{
		dbClient: dbClient,
	}
}

func (ssc *ScheduledSurveyController) GetAll(ctx *gin.Context) {
}

func (ssc *ScheduledSurveyController) Get(ctx *gin.Context) {

}

func (ssc *ScheduledSurveyController) Create(ctx *gin.Context) {

}

func (ssc *ScheduledSurveyController) Update(ctx *gin.Context) {

}

func (ssc *ScheduledSurveyController) Delete(ctx *gin.Context) {

}
