package models

import "github.com/google/uuid"

type CarAd struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewCarAd(title, description string, price float64) *CarAd {
	return &CarAd{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Price:       price,
	}
}
