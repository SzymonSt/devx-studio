package controllers

import (
	"devx/pkg/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var scheduledSurveys []models.ScheduledSurvey
	cursor, err := ssc.dbClient.Database("devx").Collection("scheduledsurveys").Find(ctx, nil)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = cursor.All(ctx, &scheduledSurveys)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, scheduledSurveys)
}

func (ssc *ScheduledSurveyController) Get(ctx *gin.Context) {
	var scheduledSurvey models.ScheduledSurvey
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})
	}
	filter := bson.M{"_id": objectId}
	err = ssc.dbClient.Database("devx").Collection("scheduledsurveys").FindOne(ctx, filter).Decode(&scheduledSurvey)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, scheduledSurvey)
}

func (ssc *ScheduledSurveyController) Create(ctx *gin.Context) {
	var scheduledSurvey models.ScheduledSurvey
	err := ctx.BindJSON(&scheduledSurvey)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	inserResult, err := ssc.dbClient.Database("devx").Collection("scheduledsurveys").InsertOne(ctx, scheduledSurvey)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Successfully created",
		"id":      inserResult.InsertedID,
	})
}

func (ssc *ScheduledSurveyController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})
		return
	}

	var scheduledSurvey models.ScheduledSurvey
	err = ctx.BindJSON(&scheduledSurvey)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request body",
		})
	}
	updateResult, err := ssc.dbClient.Database("devx").Collection("scheduledsurveys").UpdateByID(ctx, objectId, bson.M{"$set": scheduledSurvey})
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Successfully updated",
		"id":      updateResult.UpsertedID,
	})
}

func (ssc *ScheduledSurveyController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})
		return
	}

	_, err = ssc.dbClient.Database("devx").Collection("scheduledsurveys").DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Successfully deleted",
		"id":      id,
	})
}
