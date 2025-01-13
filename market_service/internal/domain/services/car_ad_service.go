package domain

import domain "market_service/internal/domain/repo"

type CarAdService struct {
	repo domain.CarAdRepository
}

func NewCarAdService(repo domain.CarAdRepository) *CarAdService {
	return &CarAdService{repo: repo}
}
