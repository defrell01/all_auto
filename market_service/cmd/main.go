package main

import (
	"fmt"
	"log"
	"market_service/internal/infrastructure/config"
	"market_service/internal/infrastructure/database"
	interfaces "market_service/internal/interfaces/http"
	"market_service/internal/usecase"
	"market_service/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func createTables(db *sqlx.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS car_ads (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		price NUMERIC NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
	log.Println("Tables initialized successfully!")
}

func main() {
	// Загрузка конфигурации
	configPath := "config/config.yaml"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger, err := logger.NewLogger(cfg.Logger.Logstash.Address)

	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("Logger initialized successfully")

	db, err := database.NewDB(*cfg)
	if err != nil {
		logger.Error("Failed to initialize database", zap.Error(err))
	}
	defer db.Close()
	createTables(db)

	carAdRepo := database.NewCarAdRepository(db)

	// Инициализация UseCase для создания объявления
	carAdUsecase := usecase.NewCreateCarAdUseCase(carAdRepo)

	// Инициализация HTTP-хендлеров
	r := mux.NewRouter()
	interfaces.NewCarAdHandler(r, carAdUsecase)

	// Запуск HTTP-сервера
	address := fmt.Sprintf(":%d", cfg.App.Port)
	logger.Info("Starting server", zap.String("address", address))
	err = http.ListenAndServe(address, r)

	if err != nil {
		logger.Error("Failed to start server", zap.Error(err))
	}

}
