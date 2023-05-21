package jobs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProcessScheduledSurvey(cfId string, surveyId string, mongoClient *mongo.Client) {
	fmt.Println("===Processing scheduled survey===")
	var continuousFeedback map[string]interface{}
	cfIDParsed, err := primitive.ObjectIDFromHex(cfId)
	if err != nil {
		entry := bson.M{
			"log":       "ERROR",
			"timestamp": primitive.DateTime(primitive.NewDateTimeFromTime(time.Now())),
			"error":     err.Error(),
		}
		_, err = mongoClient.Database("devx-scheduler").Collection("cronlog").InsertOne(context.Background(), entry)
		if err != nil {
			return
		}
	}
	filter := bson.M{"_id": cfIDParsed}
	err = mongoClient.Database("devx").Collection("continuousfeedback").FindOne(context.Background(), filter).Decode(&continuousFeedback)
	if err != nil {
		entry := bson.M{
			"log":       "ERROR",
			"timestamp": primitive.DateTime(primitive.NewDateTimeFromTime(time.Now())),
			"error":     err.Error(),
		}
		_, err = mongoClient.Database("devx-scheduler").Collection("cronlog").InsertOne(context.Background(), entry)
		if err != nil {
			return
		}
	}
	var survey map[string]interface{}
	for _, s := range continuousFeedback["scheduledSurveys"].(primitive.A) {
		surveyTmp := s.(map[string]interface{})
		if surveyTmp["_id"] == surveyId {
			survey = surveyTmp
			break
		}
	}
	if survey != nil {
		entry := bson.M{
			"log":       fmt.Sprintf("OK! Survey %s processed", survey["name"]),
			"timestamp": primitive.DateTime(primitive.NewDateTimeFromTime(time.Now())),
		}
		_, err = mongoClient.Database("devx-scheduler").Collection("cronlog").InsertOne(context.Background(), entry)
		if err != nil {
			return
		}
		//===================================================
		//TBD: sending message via integrated chat(slack/ms teams)
		//===================================================
	}
}
