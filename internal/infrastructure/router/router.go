package router

import (
	"content_collector/internal/interface/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, collectorController controller.ICollectorController) *echo.Echo {
	e.POST("/collector-data", collectorController.GetData)
	e.GET("/collector-data-clearing", collectorController.GetData)

	return e
}
