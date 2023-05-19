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
	router := rg.Group("continuousfeedback")
	router.GET("/", ssc.continuousFeedbackController.GetAll)
	router.GET("/:id", ssc.continuousFeedbackController.Get)
	router.POST("/", ssc.continuousFeedbackController.Create)
	router.PUT("/:id", ssc.continuousFeedbackController.Update)
	router.DELETE("/:id", ssc.continuousFeedbackController.Delete)
	router.POST("/answer", ssc.continuousFeedbackController.PlaceAnswer)
}
