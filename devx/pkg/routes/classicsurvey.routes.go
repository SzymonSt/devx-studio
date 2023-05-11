package routes

import (
	"devx/pkg/controllers"

	"github.com/gin-gonic/gin"
)

type ClassicSurveyRouteController struct {
	classicSurveyController controllers.ClassicSurveyController
}

func NewClassicSurveyRouteController(classicSurveyController controllers.ClassicSurveyController) ClassicSurveyRouteController {
	return ClassicSurveyRouteController{
		classicSurveyController: classicSurveyController,
	}
}

func (ssc *ClassicSurveyRouteController) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("classicsurvey")
	router.GET("/", ssc.classicSurveyController.GetAll)
	router.GET("/:id", ssc.classicSurveyController.Get)
	router.POST("/", ssc.classicSurveyController.Create)
	router.PUT("/:id", ssc.classicSurveyController.Update)
	router.DELETE("/:id", ssc.classicSurveyController.Delete)
}
