package handler

import (
	"encoding/json"
	"errors"
	"goshort/internal/domain/dtos"
	"goshort/internal/persistance"
	"goshort/internal/server"
	"goshort/internal/service"
	"log"
	"net/http"
)

type UrlHandler struct {
	Service service.UrlService
}

func (handler *UrlHandler) ShortenUrlHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if req.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer func() {
		err := req.Body.Close()
		if err != nil {
			log.Printf("error closing request body: %v", err)
		}
	}()

	var bodyDto dtos.PostUrlDTO
	err := json.NewDecoder(req.Body).Decode(&bodyDto)
	if err != nil {
		server.HandleError(http.StatusBadRequest, err, req.RequestURI, w)
		return
	}

	url, err := handler.Service.HandlePostUrl(bodyDto.URL)

	if err != nil {
		if errors.Is(err, service.ErrUrlCannotBeEmpty) {
			server.HandleError(http.StatusBadRequest, err, req.RequestURI, w)
		} else if errors.Is(err, persistance.ErrNotFound) {
			server.HandleError(http.StatusNotFound, err, req.RequestURI, w)
		}
		return
	}

	w.Header().Set(
		"Location",
		url,
	)
	w.WriteHeader(http.StatusCreated)
}
