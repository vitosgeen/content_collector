package router

import (
	"content_collector/internal/interface/controller"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, collectorController controller.ICollectorController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &controller.CustomValidator{Validator: validator.New()}

	e.POST("/collector-data", collectorController.GetData)
	e.GET("/collector-data-clearing", collectorController.Clearing)
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	return e
}
