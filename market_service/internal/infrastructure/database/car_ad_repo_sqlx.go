// internal/infrastructure/database/car_ad_repo_sqlx.go
package database

import (
	domain "market_service/internal/domain/models"
	"market_service/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type CarAdRepository struct {
	DB     *sqlx.DB
	Logger *logger.Logger
}

// Конструктор для CarAdRepository
func NewCarAdRepository(db *sqlx.DB, logger *logger.Logger) *CarAdRepository {
	return &CarAdRepository{DB: db, Logger: logger}
}

// Метод для добавления нового объявления
func (r *CarAdRepository) CreateCarAd(carAd *domain.CarAd) error {

	r.Logger.Info("Creating new car ad in database")
	_, err := r.DB.Exec(`
		INSERT INTO car_ads (title, description, price)
		VALUES ($1, $2, $3)`,
		carAd.Title, carAd.Description, carAd.Price,
	)
	return err
}
