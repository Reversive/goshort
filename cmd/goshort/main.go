package main

import (
	"goshort/internal/handler"
	"goshort/internal/persistance"
	"goshort/internal/server"
	"goshort/internal/service"
	"log"
	"net/http"
)

func main() {
	var router = server.NewRouter()
	var urlRepo = persistance.New()
	var urlService = service.NewUrlService(urlRepo)
	var urlHandler = handler.NewUrlHandler(urlService)

	router.HandleFunc("/urls", urlHandler.ShortenUrlHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
