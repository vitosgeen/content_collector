package repository

import "content_collector/internal/domain/model"

type ICollectorRepository interface {
	GetById(id string) (*model.CollectorRepository, error)
	GetByUrl(url string) (*model.CollectorRepository, error)
	GetForDelete() ([]*model.CollectorRepository, error)
	Create(collector *model.CollectorRepository) error
	Update(collector *model.CollectorRepository) error
	Delete(id string) error
}
