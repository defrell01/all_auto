// internal/infrastructure/database/car_ad_repo_sqlx.go
package database

import (
	domain "market_service/internal/domain/models"

	"github.com/jmoiron/sqlx"
)

type CarAdRepository struct {
	DB *sqlx.DB
}

// Конструктор для CarAdRepository
func NewCarAdRepository(db *sqlx.DB) *CarAdRepository {
	return &CarAdRepository{DB: db}
}

// Метод для добавления нового объявления
func (r *CarAdRepository) CreateCarAd(carAd *domain.CarAd) error {
	_, err := r.DB.Exec(`
		INSERT INTO car_ads (title, description, price)
		VALUES ($1, $2, $3)`,
		carAd.Title, carAd.Description, carAd.Price,
	)
	return err
}
