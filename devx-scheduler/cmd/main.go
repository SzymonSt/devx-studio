package main

import (
	"context"
	"devx-scheduler/cmd/util"
	"devx-scheduler/pkg/controllers"
	"devx-scheduler/pkg/jobs"
	"devx-scheduler/pkg/models"
	"devx-scheduler/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cron "github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var (
		server = gin.Default()
	)
	config, err := util.LaodConfig()
	if err != nil {
		log.Fatal(err)
	}

	//Set up cron
	c := cron.New()

	//Set up mongo client
	// mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.DBUri))
	// if err != nil {
	// 	fmt.Println("[ERROR]" + err.Error())
	// }
	// err = loadExistingJobs(c, mongoClient, config.SlackHook)
	// if err != nil {
	// 	fmt.Println("[ERROR]" + err.Error())
	// }

	//Set up http server
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	server.Use(cors.New(corsConfig))

	schedulerController := controllers.NewSchedulerController(c)
	schedulerRouteController := routes.NewSchedulerRoutes(schedulerController)

	router := server.Group("/api")
	router.GET("/healthz", func(ctx *gin.Context) {
		message := "devx-scheduler api is up and running"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	schedulerRouteController.RegisterRoutes(router)

	c.Start()
	log.Fatal(server.Run(":" + config.ServerPort))

}

func loadExistingJobs(c *cron.Cron, m *mongo.Client, slackHook string) (err error) {
	var fetchedJobs []models.Job
	getResult, err := m.Database("devx-scheduler").Collection("jobs").Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	err = getResult.All(context.Background(), &fetchedJobs)
	if err != nil {
		return err
	}
	for _, job := range fetchedJobs {
		jId, err := c.AddFunc(job.Cron, func() { jobs.ProcessScheduledSurvey(job.CfId, job.SurveyId) })
		if err != nil {
			return err
		}
		job.JobId = int64(jId)
		_, err = m.Database("devx-scheduler").Collection("jobs").UpdateOne(context.Background(), bson.M{"_id": job.Id}, bson.M{"$set": bson.M{"jobId": job.JobId}})
		if err != nil {
			return err
		}
		fmt.Printf("cron %s scheduled for survey: %s", job.Cron, job.SurveyId)
	}
	return nil
}
