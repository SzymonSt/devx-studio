package routes

import (
	"devx/pkg/controllers"

	"github.com/gin-gonic/gin"
)

type VerticalDataRoutes struct {
	VerticalDataController controllers.VerticalDataController
}

func NewVerticalDataRoutes(verticalDataController controllers.VerticalDataController) VerticalDataRoutes {
	return VerticalDataRoutes{
		VerticalDataController: verticalDataController,
	}
}

func (vdr *VerticalDataRoutes) RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("verticaldata")
	router.GET("overall/:interval/:verticalId", vdr.VerticalDataController.GetOverallVerticalData)
	router.GET("detailed/:interval/:verticalId", vdr.VerticalDataController.GetDetailedVerticalData)
}
