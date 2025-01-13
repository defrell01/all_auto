package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func JSONResponse(w http.ResponseWriter, statusCode int, status string, data interface{}, errMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Status: status,
		Data:   data,
		Error:  errMessage,
	}

	_ = json.NewEncoder(w).Encode(response)
}
