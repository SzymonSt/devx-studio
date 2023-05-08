package controllers

import (
	"devx/pkg/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClassicSurveyController struct {
	dbClient *mongo.Client
}

func NewClassicSurveyController(dbClient *mongo.Client) ClassicSurveyController {
	return ClassicSurveyController{
		dbClient: dbClient,
	}
}

func (ssc *ClassicSurveyController) GetAll(ctx *gin.Context) {
	var classicSurveys []models.ClassicSurvey
	cursor, err := ssc.dbClient.Database("devx").Collection("classicsurveys").Find(ctx, nil)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = cursor.All(ctx, &classicSurveys)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, classicSurveys)
}

func (ssc *ClassicSurveyController) Get(ctx *gin.Context) {
	var classicSurvey models.ClassicSurvey
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})
	}
	filter := bson.M{"_id": objectId}
	err = ssc.dbClient.Database("devx").Collection("classicsurveys").FindOne(ctx, filter).Decode(&classicSurvey)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, classicSurvey)
}

func (ssc *ClassicSurveyController) Create(ctx *gin.Context) {
	var classicSurvey models.ClassicSurvey
	err := ctx.BindJSON(&classicSurvey)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	inserResult, err := ssc.dbClient.Database("devx").Collection("classicsurveys").InsertOne(ctx, classicSurvey)
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

func (ssc *ClassicSurveyController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})
		return
	}

	var classicSurvey models.ClassicSurvey
	err = ctx.BindJSON(&classicSurvey)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request body",
		})
	}
	updateResult, err := ssc.dbClient.Database("devx").Collection("classicsurveys").UpdateByID(ctx, objectId, bson.M{"$set": classicSurvey})
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

func (ssc *ClassicSurveyController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})
		return
	}

	_, err = ssc.dbClient.Database("devx").Collection("classicsurveys").DeleteOne(ctx, bson.M{"_id": objectId})
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
