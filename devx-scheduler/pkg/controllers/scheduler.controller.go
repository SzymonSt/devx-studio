package controllers

import (
	"devx-scheduler/pkg/jobs"
	"devx-scheduler/pkg/models"
	"fmt"

	"github.com/gin-gonic/gin"
	cron "github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SchedulerController struct {
	cron *cron.Cron
}

func NewSchedulerController(cron *cron.Cron) SchedulerController {
	return SchedulerController{
		cron: cron,
	}
}

func (sc *SchedulerController) CreateOrUpdateJob(ctx *gin.Context) {
	var job models.Job
	err := ctx.BindJSON(&job)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	job.Id = primitive.NewObjectID()

	_, err = sc.cron.AddFunc(job.Cron, func() { jobs.ProcessScheduledSurvey(job.CfId, job.SurveyId) })
	if err != nil {
		fmt.Printf("[ERROR] %s", err.Error())
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Job updated successfully",
	})
}
