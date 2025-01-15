package interfaces

import (
	"encoding/json"
	domain "market_service/internal/domain/models"
	"market_service/internal/interfaces/dto"
	"market_service/internal/usecase"
	"market_service/pkg/logger"
	"market_service/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type CarAdHandler struct {
	UseCase *usecase.CreateCarAdUseCase
	Logger  *logger.Logger
}

func NewCarAdHandler(r *mux.Router, useCase *usecase.CreateCarAdUseCase, logger *logger.Logger) {
	handler := &CarAdHandler{UseCase: useCase, Logger: logger}
	r.HandleFunc("/car_ads", handler.CreateCarAd).Methods("POST")
}

func (h *CarAdHandler) CreateCarAd(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCarAdRequest

	h.Logger.Info("Handle CreateCarAd request")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.Error("Invalid request body", zap.Error(err))
		utils.JSONResponse(w, http.StatusBadRequest, "error", nil, "Invalid request body")
		return
	}

	// Валидируем запрос
	validator := utils.NewValidator()
	if err := validator.Validate(req); err != nil {
		h.Logger.Error("Err validation JSON", zap.Error(err))
		utils.JSONResponse(w, http.StatusBadRequest, "error", nil, err.Error())
		return
	}

	ad := domain.NewCarAd(req.Title, req.Description, req.Price)

	if err := h.UseCase.Execute(ad); err != nil {
		h.Logger.Error("Failed to create car ad", zap.Error(err))
		utils.JSONResponse(w, http.StatusInternalServerError, "error", nil, "Failed to create car ad")
		return
	}

	utils.JSONResponse(w, http.StatusCreated, "success", ad, "")
}
