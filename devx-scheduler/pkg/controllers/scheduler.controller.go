package controllers

import (
	"devx-scheduler/pkg/jobs"
	"devx-scheduler/pkg/models"

	"github.com/gin-gonic/gin"
	cron "github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SchedulerController struct {
	dbClient  *mongo.Client
	cron      *cron.Cron
	slackHook string
}

func NewSchedulerController(dbClient *mongo.Client, cron *cron.Cron, slackHook string) SchedulerController {
	return SchedulerController{
		dbClient:  dbClient,
		cron:      cron,
		slackHook: slackHook,
	}
}

func (sc *SchedulerController) CreateOrUpdateJob(ctx *gin.Context) {
	var job models.Job
	err := ctx.BindJSON(&job)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request body",
		})
	}
	filter := bson.M{"surveyId": job.SurveyId, "cfId": job.CfId}
	singleResult := sc.dbClient.Database("devx-scheduler").Collection("jobs").FindOne(ctx, filter)
	if singleResult.Err() == mongo.ErrNoDocuments {
		job.Id = primitive.NewObjectID()
		jId, err := sc.cron.AddFunc(job.Cron, func() { jobs.ProcessScheduledSurvey(job.CfId, job.SurveyId, sc.dbClient, sc.slackHook) })
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		job.JobId = int64(jId)
		_, err = sc.dbClient.Database("devx-scheduler").Collection("jobs").InsertOne(ctx, job)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
	} else {
		var oldJob models.Job
		singleResult.Decode(&oldJob)
		id := cron.EntryID(oldJob.JobId)
		sc.cron.Remove(id)
		jId, err := sc.cron.AddFunc(job.Cron, func() { jobs.ProcessScheduledSurvey(job.CfId, job.SurveyId, sc.dbClient, sc.slackHook) })
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		job.JobId = int64(jId)
		_, err = sc.dbClient.Database("devx-scheduler").Collection("jobs").UpdateOne(ctx, bson.M{"_id": oldJob.Id}, bson.M{"$set": bson.M{"cron": job.Cron, "jobId": job.JobId}})
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	ctx.JSON(200, gin.H{
		"message": "Job updated successfully",
	})
}
