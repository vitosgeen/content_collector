package model

type CollectResponse struct {
	Url    string `json:"url"`
	Length int    `json:"length"`
	Data   string `json:"data"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}
