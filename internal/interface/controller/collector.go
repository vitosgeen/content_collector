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

// GetData is a method of the CollectorController struct that handles the HTTP GET request to collect data.
// It validates the request, calls the CollectorService to collect data from the specified URL,
// and returns the collected data as a JSON response.
// If any error occurs during the process, it returns an appropriate error response.
func (controller *CollectorController) GetData(ctx echo.Context) error {
	// validate request
	request := &model.CollectorRequest{}
	err := ctx.Bind(request)
	if err != nil {
		appError := apperrors.ControllerCollectorGetDataError.AppendMessage(err)
		responseError := model.CollectResponseError{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Error:  "Invalid Object Bind",
		}
		return ctx.JSON(appError.HTTPCode, responseError)
	}
	err = ctx.Validate(request)
	if err != nil {
		appError := apperrors.ControllerCollectorGetDataError.AppendMessage(err)
		responseError := model.CollectResponseError{
			Code:      http.StatusBadRequest,
			ErrorCode: appError.Code,
			Status:    http.StatusText(http.StatusBadRequest),
			Error:     "Invalid Object Validate",
		}
		return ctx.JSON(appError.HTTPCode, responseError)
	}
	if !model.IsValidUrl(request.Url) {
		appError := apperrors.ControllerCollectorGetDataError.AppendMessage(err)
		responseError := model.CollectResponseError{
			Code:      http.StatusBadRequest,
			ErrorCode: appError.Code,
			Status:    http.StatusText(http.StatusBadRequest),
			Error:     "Invalid URL",
		}
		return ctx.JSON(appError.HTTPCode, responseError)
	}

	// call service
	scrapperData, err := controller.CollectorService.Collect(request.Url)
	if err != nil {
		appError := apperrors.ControllerCollectorCollect.AppendMessage(err)
		responseError := model.CollectResponseError{
			Code:      scrapperData.Code,
			ErrorCode: appError.Code,
			Status:    scrapperData.Status,
			Error:     appError.Message,
		}
		return ctx.JSON(appError.HTTPCode, responseError)
	}

	response := model.CollectResponse{
		Url:    request.Url,
		Data:   scrapperData.Data,
		Length: len(scrapperData.Data),
		Code:   scrapperData.Code,
		Status: scrapperData.Status,
	}

	return ctx.JSON(http.StatusOK, response)
}

// Clearing is a method of the CollectorController struct that handles the clearing of collected data.
// It takes an echo.Context as a parameter and returns an error.
// The method is responsible for clearing the collected data and returning any errors that occur during the process.
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
