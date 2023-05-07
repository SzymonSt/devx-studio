package routes

import (
	"devx/pkg/controllers"

	"github.com/gin-gonic/gin"
)

type ContinuousFeedbackRouteController struct {
	continuousFeedbackController controllers.ContinuousFeedbackController
}

func NewContinuousFeedbackRouteController(continuousFeedbackController controllers.ContinuousFeedbackController) ContinuousFeedbackRouteController {
	return ContinuousFeedbackRouteController{
		continuousFeedbackController: continuousFeedbackController,
	}
}

func (ssc *ContinuousFeedbackRouteController) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("ContinuousFeedbacks")
	router.GET("/continuousfeedback", ssc.continuousFeedbackController.GetAll)
	router.GET("/continuousfeedback/:id", ssc.continuousFeedbackController.Get)
	router.POST("/continuousfeedback", ssc.continuousFeedbackController.Create)
	router.PUT("/continuousfeedback/:id", ssc.continuousFeedbackController.Update)
	router.DELETE("/continuousfeedback/:id", ssc.continuousFeedbackController.Delete)
}
