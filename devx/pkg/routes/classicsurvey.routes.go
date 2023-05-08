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
	router := rg.Group("classicsurveys")
	router.GET("/classicsurvey", ssc.classicSurveyController.GetAll)
	router.GET("/classicsurvey/:id", ssc.classicSurveyController.Get)
	router.POST("/classicsurvey", ssc.classicSurveyController.Create)
	router.PUT("/classicsurvey/:id", ssc.classicSurveyController.Update)
	router.DELETE("/classicsurvey/:id", ssc.classicSurveyController.Delete)
}
