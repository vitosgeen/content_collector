package apperrors

import "net/http"

var (
	SeleniumChromeScrapperScrapNewChromeDriverServiceError = AppError{
		Message:  "Failed to create new chrome driver service",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_NEW_CHROME_DRIVER_SERVICE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SeleniumChromeScrapperScrapNewRemoteError = AppError{
		Message:  "Failed to create new remote",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_NEW_REMOTE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SeleniumChromeScrapperScrapMaximizeWindow = AppError{
		Message:  "Failed to maximize window",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_MAXIMIZE_WINDOW_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SeleniumChromeScrapperScrapExecuteScriptError = AppError{
		Message:  "Failed to execute script",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_EXECUTE_SCRIPT_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SeleniumChromeScrapperScrapDriverGet = AppError{
		Message:  "Failed to get driver",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_DRIVER_GET_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SeleniumChromeScrapperScrapPageSource = AppError{
		Message:  "Failed to get page source",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_PAGE_SOURCE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SeleniumChromeScrapperScrapProxyAuthMaximizeWindow = AppError{
		Message:  "Failed to maximize window",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_PROXY_AUTH_MAXIMIZE_WINDOW_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SeleniumChromeScrapperScrapProxyAuthDriverGet = AppError{
		Message:  "Failed to get driver",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_PROXY_AUTH_DRIVER_GET_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SeleniumChromeScrapperScrapProxyAuthPageSource = AppError{
		Message:  "Failed to get page source",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_PROXY_AUTH_PAGE_SOURCE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SeleniumChromeScrapperScrapProxyAuthNewRemoteError = AppError{
		Message:  "Failed to create new remote",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_PROXY_AUTH_NEW_REMOTE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SeleniumChromeScrapperScrapProxyAuthNewChromeDriverServiceError = AppError{
		Message:  "Failed to create new chrome driver service",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_PROXY_AUTH_NEW_CHROME_DRIVER_SERVICE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SeleniumChromeScrapperScrapBuildProxyExtensionError = AppError{
		Message:  "Failed to build proxy extension",
		Code:     "SELENIUM_CHROME_SCRAPPER_SCRAP_BUILD_PROXY_EXTENSION_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
)
