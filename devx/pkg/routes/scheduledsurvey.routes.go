package routes

import (
	"devx/pkg/controllers"

	"github.com/gin-gonic/gin"
)

type ScheduledSurveyRouteController struct {
	scheduledSurveyController controllers.ScheduledSurveyController
}

func NewScheduledSurveyRouteController(scheduledSurveyController controllers.ScheduledSurveyController) ScheduledSurveyRouteController {
	return ScheduledSurveyRouteController{
		scheduledSurveyController: scheduledSurveyController,
	}
}

func (ssc *ScheduledSurveyRouteController) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("scheduledsurveys")
	router.GET("/scheduledsurvey", ssc.scheduledSurveyController.GetAll)
	router.GET("/scheduledsurvey/:id", ssc.scheduledSurveyController.Get)
	router.POST("/scheduledsurvey", ssc.scheduledSurveyController.Create)
	router.PUT("/scheduledsurvey/:id", ssc.scheduledSurveyController.Update)
	router.DELETE("/scheduledsurvey/:id", ssc.scheduledSurveyController.Delete)
}
