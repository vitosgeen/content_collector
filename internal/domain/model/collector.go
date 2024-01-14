package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/tebeka/selenium"
)

const (
	CollectorRepositoryStatusActive   = 1
	CollectorRepositoryStatusInactive = 0
)

type Collector struct {
	PathChromeDriver string
	PortChromeDriver int
	WebDriver        selenium.WebDriver
	ProxyIp          string
}

type CollectorRepository struct {
	ID             string     `json:"id" bson:"_id" validate:"required"`
	Url            string     `json:"url" bson:"url" validate:"required"`
	Data           string     `json:"data" bson:"data" validate:"required"`
	DataLen        int        `json:"data_len" bson:"data_len" validate:"required"`
	DataCode       int        `json:"data_code" bson:"data_code" validate:"required"`
	DataStatusText string     `json:"data_status_text" bson:"data_status_text" validate:"required"`
	Status         int        `json:"status" bson:"status" validate:"required"`
	CreatedAt      *time.Time `json:"created_at" bson:"created_at" validate:"required"`
	UpdatedAt      *time.Time `json:"updated_at" bson:"updated_at" validate:"required"`
	DeleteAt       *time.Time `json:"delete_at" bson:"delete_at" validate:"required"`
}

func NewUUID() string {
	return uuid.New().String()
}

func NewTime() time.Time {
	return time.Now()
}
