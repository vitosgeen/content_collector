package controller

import (
	"fmt"
	"net/http"
	"time"

	"content_collector/internal/apperrors"
	"content_collector/internal/domain/model"
	"content_collector/internal/services"

	"github.com/labstack/echo/v4"
)

var (
	ClearingInterval = 60
	ClearingStart    = 0
)

type ICollectorController interface {
	GetData(c echo.Context) error
	Clearing(c echo.Context) error
	// Create(c echo.Context) error
	// Update(c echo.Context) error
	// Delete(c echo.Context) error
}

type CollectorController struct {
	CollectorService services.ICollectorService
}

func NewCollectorController(collectorService services.ICollectorService) ICollectorController {
	return &CollectorController{
		CollectorService: collectorService,
	}
}

func (controller *CollectorController) GetData(ctx echo.Context) error {
	// validate request
	request := &model.CollectorRequest{}
	err := ctx.Bind(request)
	if err != nil {
		appError := apperrors.ControllerCollectorGetDataError.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Message)
	}
	err = ctx.Validate(request)
	if err != nil {
		appError := apperrors.ControllerCollectorGetDataError.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Message)
	}
	// call service
	html, err := controller.CollectorService.Collect(request.Url)
	if err != nil {
		appError := apperrors.ControllerCollectorCollect.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Message)
	}

	response := model.CollectResponse{
		Url:    request.Url,
		Data:   html,
		Length: len(html),
	}

	return ctx.JSON(http.StatusOK, response)
}

func (controller *CollectorController) Clearing(ctx echo.Context) error {
	if ClearingStart != 0 {
		responseStr := fmt.Sprintf("Clearing already started. Start: %d", ClearingStart)
		return ctx.JSON(http.StatusOK, responseStr)
	}

	err := controller.CollectorService.DeleteOldCollectors()
	if err != nil {
		appError := apperrors.ControllerCollectorClearingDeleteOldCollectors.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Message)
	}

	go gorutineClearingLoop(controller.CollectorService)

	return ctx.JSON(http.StatusOK, "OK")
}

func gorutineClearingLoop(collectorService services.ICollectorService) {
	defer func() {
		ClearingStart = 0
	}()
	for {
		ClearingStart++
		err := collectorService.DeleteOldCollectors()
		if err != nil {
			apperrors.ServicesCollectorGorutineClearingLoopDeleteOldCollectors.AppendMessage(err) //nolint:errcheck
			break
		}

		time.Sleep(time.Duration(ClearingInterval) * time.Second)
	}
}
