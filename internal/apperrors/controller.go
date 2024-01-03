package apperrors

import "net/http"

var (
	ControllerCollectorGetDataError = AppError{
		Message:  "Failed to get data",
		Code:     "CONTROLLER_COLLECTOR_GET_DATA_ERR",
		HTTPCode: http.StatusBadRequest,
	}
	ControllerCollectorCollect = AppError{
		Message:  "Failed to collect",
		Code:     "CONTROLLER_COLLECTOR_COLLECT_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ControllerCollectorClearingDeleteOldCollectors = AppError{
		Message:  "Failed to delete old collectors",
		Code:     "CONTROLLER_COLLECTOR_CLEARING_DELETE_OLD_COLLECTORS_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorGorutineClearingLoopDeleteOldCollectors = AppError{
		Message:  "Failed to delete old collectors",
		Code:     "SERVICES_COLLECTOR_GORUTINE_CLEARING_LOOP_DELETE_OLD_COLLECTORS_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
)
