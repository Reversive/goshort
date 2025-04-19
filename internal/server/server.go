package server

import (
	"encoding/json"
	"goshort/internal/domain/dtos"
	"log"
	"net/http"
)

func NewRouter() *http.ServeMux {
	return http.NewServeMux()
}

func HandleError(status int, err error, path string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	payload, err := json.Marshal(dtos.ApiErrorDto{
		Code:  status,
		Path:  path,
		Error: err.Error(),
	})
	if err != nil {
		log.Printf("Error while marshaling response")
		return
	}

	_, err = w.Write(payload)
	if err != nil {
		log.Printf("Error while writing response")
		return
	}
}
