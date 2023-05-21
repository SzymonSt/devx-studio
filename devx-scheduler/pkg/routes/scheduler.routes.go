package routes

import (
	"devx-scheduler/pkg/controllers"

	"github.com/gin-gonic/gin"
)

type SchedulerRoutes struct {
	SchedulerController controllers.SchedulerController
}

func NewSchedulerRoutes(schedulerController controllers.SchedulerController) SchedulerRoutes {
	return SchedulerRoutes{
		SchedulerController: schedulerController,
	}
}

func (sr *SchedulerRoutes) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("scheduler")
	router.POST("/job", sr.SchedulerController.CreateOrUpdateJob)
}
