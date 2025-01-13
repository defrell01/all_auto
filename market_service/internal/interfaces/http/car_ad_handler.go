package interfaces

import (
	"encoding/json"
	domain "market_service/internal/domain/models"
	"market_service/internal/interfaces/dto"
	"market_service/internal/usecase"
	"market_service/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type CarAdHandler struct {
	UseCase *usecase.CreateCarAdUseCase
}

func NewCarAdHandler(r *mux.Router, useCase *usecase.CreateCarAdUseCase) {
	handler := &CarAdHandler{UseCase: useCase}
	r.HandleFunc("/car_ads", handler.CreateCarAd).Methods("POST")
}

func (h *CarAdHandler) CreateCarAd(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCarAdRequest

	// Декодируем тело запроса
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "error", nil, "Invalid request body")
		return
	}

	// Валидируем запрос
	validator := utils.NewValidator()
	if err := validator.Validate(req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "error", nil, err.Error())
		return
	}

	// Создаем объект объявления
	ad := domain.NewCarAd(req.Title, req.Description, req.Price)

	// Выполняем бизнес-логику
	if err := h.UseCase.Execute(ad); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "error", nil, "Failed to create car ad")
		return
	}

	utils.JSONResponse(w, http.StatusCreated, "success", ad, "")
}
