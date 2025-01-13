// internal/usecase/create_car_ad_usecase.go
package usecase

import (
	domain "market_service/internal/domain/models"
	"market_service/internal/infrastructure/database"
)

type CreateCarAdUseCase struct {
	CarAdRepository *database.CarAdRepository
}

func NewCreateCarAdUseCase(carAdRepo *database.CarAdRepository) *CreateCarAdUseCase {
	return &CreateCarAdUseCase{CarAdRepository: carAdRepo}
}

func (uc *CreateCarAdUseCase) Execute(carAd *domain.CarAd) error {
	return uc.CarAdRepository.CreateCarAd(carAd)
}
