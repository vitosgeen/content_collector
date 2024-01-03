package model

type CollectorRequest struct {
	Url string `json:"url" validate:"required"`
}
