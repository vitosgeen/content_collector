package model

import "net/url"

type CollectResponse struct {
	Url    string `json:"url"`
	Length int    `json:"length"`
	Data   string `json:"data"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type CollectResponseError struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

func IsValidUrl(validateUrl string) bool {
	validatedUrl, err := url.Parse(validateUrl)
	if err != nil {
		return false
	}
	if validatedUrl.Scheme == "" || validatedUrl.Host == "" {
		return false
	}
	return true
}
