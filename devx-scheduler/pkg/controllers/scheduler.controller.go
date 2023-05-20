package controllers

import (
	"devx-scheduler/pkg/jobs"
	"devx-scheduler/pkg/models"

	"github.com/gin-gonic/gin"
	cron "github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SchedulerController struct {
	dbClient *mongo.Client
	cron     *cron.Cron
}

func NewSchedulerController(dbClient *mongo.Client, cron *cron.Cron) SchedulerController {
	return SchedulerController{
		dbClient: dbClient,
		cron:     cron,
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
	filter := bson.M{"surveyId": job.Id, "cfId": job.CfId}
	singleResult := sc.dbClient.Database("devx-scheduler").Collection("jobs").FindOne(ctx, filter)
	if singleResult.Err() == mongo.ErrNoDocuments {
		jId, err := sc.cron.AddFunc(job.Cron, func() { jobs.ProcessScheduledSurvey(job.CfId, job.Id) })
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Internal server error",
			})
		}
		job.JobId = int64(jId)
		_, err = sc.dbClient.Database("devx-schedule").Collection("jobs").InsertOne(ctx, job)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Internal server error",
			})
		}
	} else {
		id := cron.EntryID(job.JobId)
		sc.cron.Remove(id)
		jId, err := sc.cron.AddFunc(job.Cron, func() { jobs.ProcessScheduledSurvey(job.CfId, job.Id) })
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Internal server error",
			})
		}
		job.JobId = int64(jId)
		_, err = sc.dbClient.Database("devx-schedule").Collection("jobs").UpdateOne(ctx, filter, bson.M{"$set": job})
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Internal server error",
			})
		}
	}
	ctx.JSON(200, gin.H{
		"message": "Job updated successfully",
	})
}
