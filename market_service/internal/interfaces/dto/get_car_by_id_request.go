package dto

type GetCarById struct {
	Id string `json:"id" validate:"requred"`
}
