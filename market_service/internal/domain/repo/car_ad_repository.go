package domain

import (
	domain "market_service/internal/domain/models"
)

type CarAdRepository interface {
	Create(ad domain.CarAd) (int, error)
	GetByID(id int) (*domain.CarAd, error)
	Update(ad domain.CarAd) error
	Delete(id int) error
	GetAll() ([]domain.CarAd, error)
}
