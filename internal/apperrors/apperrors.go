package apperrors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Message  string
	Code     string
	HTTPCode int
}

var (
	ServicesCollectorCloseError = AppError{
		Message:  "Failed to close collector service",
		Code:     "SERVICES_COLLECTOR_CLOSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCheckCollectorNewChromeDriverServiceError = AppError{
		Message:  "Failed to check collector service",
		Code:     "SERVICES_COLLECTOR_CHECK_COLLECTOR_NEW_CHROME_DRIVER_SERVICE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCheckCollectorNewRemoteError = AppError{
		Message:  "Failed to check collector service",
		Code:     "SERVICES_COLLECTOR_CHECK_COLLECTOR_NEW_REMOTE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCheckCollectorMaximizeWindowError = AppError{
		Message:  "Failed to check collector service",
		Code:     "SERVICES_COLLECTOR_CHECK_COLLECTOR_MAXIMIZE_WINDOW_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCheckCollectorDriverGetError = AppError{
		Message:  "Failed to check collector service",
		Code:     "SERVICES_COLLECTOR_CHECK_COLLECTOR_DRIVER_GET_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorDeleteOldCollectorsError = AppError{
		Message:  "Failed to delete old collectors",
		Code:     "SERVICES_COLLECTOR_DELETE_OLD_COLLECTORS_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCollectNewChromeDriverServiceError = AppError{
		Message:  "Failed to collect",
		Code:     "SERVICES_COLLECTOR_COLLECT_NEW_CHROME_DRIVER_SERVICE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCollectGetProxyRandomError = AppError{
		Message:  "Failed to collect",
		Code:     "SERVICES_COLLECTOR_COLLECT_GET_PROXY_RANDOM_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCollectScrapError = AppError{
		Message:  "Failed to collect",
		Code:     "SERVICES_COLLECTOR_COLLECT_SCRAP_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCollectGetByUrlError = AppError{
		Message:  "Failed to collect",
		Code:     "SERVICES_COLLECTOR_COLLECT_GET_BY_URL_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCollectNewRemoteError = AppError{
		Message:  "Failed to collect",
		Code:     "SERVICES_COLLECTOR_COLLECT_NEW_REMOTE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCollectMaximizeWindow = AppError{
		Message:  "Failed to collect",
		Code:     "SERVICES_COLLECTOR_COLLECT_MAXIMIZE_WINDOW_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCollectDriverGet = AppError{
		Message:  "Failed to collect",
		Code:     "SERVICES_COLLECTOR_COLLECT_DRIVER_GET_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCollectPageSource = AppError{
		Message:  "Failed to collect",
		Code:     "SERVICES_COLLECTOR_COLLECT_PAGE_SOURCE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCollectCreateError = AppError{
		Message:  "Failed to collect",
		Code:     "SERVICES_COLLECTOR_COLLECT_CREATE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServicesCollectorCollectDecodeError = AppError{
		Message:  "Failed to collect",
		Code:     "SERVICES_COLLECTOR_COLLECT_DECODE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigLoadError = AppError{
		Message:  "Failed to load env file",
		Code:     "ENV_INIT_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigParseError = AppError{
		Message:  "Failed to parse env file",
		Code:     "ENV_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigPostgresParseError = AppError{
		Message:  "Failed to parse pastgres env file",
		Code:     "ENV_POSTGRES_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigMongoParseError = AppError{
		Message:  "Failed to parse mongo env file",
		Code:     "ENV_MONGO_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigRedisParseError = AppError{
		Message:  "Failed to parse redis env file",
		Code:     "ENV_REDIS_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigSeleniumChromeParseError = AppError{
		Message:  "Failed to parse selenium chrome env file",
		Code:     "ENV_SELENIUM_CHROME_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigJwtParseError = AppError{
		Message:  "Failed to parse jwt env file",
		Code:     "ENV_CONFIG_JWT_PARSE_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}

	SqlOpenError = AppError{
		Message:  "Failed to connect database",
		Code:     "SQL_OPEN_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	PingDBError = AppError{
		Message:  "Failed ping to database",
		Code:     "PING_DB_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	PingRedisError = AppError{
		Message:  "Failed ping to redis",
		Code:     "PING_REDIS_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	ServerStartError = AppError{
		Message:  "Failed start app",
		Code:     "SERVER_START_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
)

func (appError *AppError) Error() string {
	return appError.Code + ": " + appError.Message
}

func (appError *AppError) AppendMessage(anyErrs ...interface{}) *AppError {
	return &AppError{
		Message:  fmt.Sprintf("%v : %v", appError.Message, anyErrs),
		Code:     appError.Code,
		HTTPCode: appError.HTTPCode,
	}
}

func Is(err1 error, err2 *AppError) bool {
	err, ok := err1.(*AppError)
	if !ok {
		return false
	}

	return err.Code == err2.Code
}
