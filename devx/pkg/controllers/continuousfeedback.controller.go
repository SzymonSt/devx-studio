package controllers

import (
	"devx/pkg/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContinuousFeedbackController struct {
	dbClient *mongo.Client
}

func NewContinuousFeedbackController(db *mongo.Client) ContinuousFeedbackController {
	return ContinuousFeedbackController{
		dbClient: db,
	}
}

func (cfc *ContinuousFeedbackController) GetAll(ctx *gin.Context) {
	var continuousFeedback []models.ContinuousFeedback
	cursor, err := cfc.dbClient.Database("devx").Collection("continuousfeedback").Find(ctx, nil)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = cursor.All(ctx, &continuousFeedback)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, continuousFeedback)
}

func (cfc *ContinuousFeedbackController) Get(ctx *gin.Context) {
	var continuousFeedback models.ContinuousFeedback
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})
	}
	filter := bson.M{"_id": objectId}
	err = cfc.dbClient.Database("devx").Collection("continuouseedback").FindOne(ctx, filter).Decode(&continuousFeedback)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, continuousFeedback)
}

func (cfc *ContinuousFeedbackController) Create(ctx *gin.Context) {
	var continuousFeedback models.ContinuousFeedback
	err := ctx.BindJSON(&continuousFeedback)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	inserResult, err := cfc.dbClient.Database("devx").Collection("continuousfeedback").InsertOne(ctx, continuousFeedback)
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

func (cfc *ContinuousFeedbackController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})
		return
	}

	var continuousFeedback models.ContinuousFeedback
	err = ctx.BindJSON(&continuousFeedback)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request body",
		})
	}
	updateResult, err := cfc.dbClient.Database("devx").Collection("continuousfeedback").UpdateByID(ctx, objectId, bson.M{"$set": continuousFeedback})
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

func (cfc *ContinuousFeedbackController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})
		return
	}

	_, err = cfc.dbClient.Database("devx").Collection("continuousfeedback").DeleteOne(ctx, bson.M{"_id": objectId})
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
