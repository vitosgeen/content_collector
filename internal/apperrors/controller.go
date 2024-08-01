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
	ControllerCollectorGetDataByURLError = AppError{
		Message:  "Failed to get data by URL",
		Code:     "CONTROLLER_COLLECTOR_GET_DATA_BY_URL_ERR",
		HTTPCode: http.StatusBadRequest,
	}
	ControllerCollectorGetDataByURLInvalidURL = AppError{
		Message:  "Invalid URL",
		Code:     "CONTROLLER_COLLECTOR_GET_DATA_BY_URL_INVALID_URL_ERR",
		HTTPCode: http.StatusBadRequest,
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
	ValidatorCustomValidatorValidate = AppError{
		Message:  "Failed to validate",
		Code:     "VALIDATOR_CUSTOM_VALIDATOR_VALIDATE_ERR",
		HTTPCode: http.StatusBadRequest,
	}
)
