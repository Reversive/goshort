package main

import (
	"fmt"
	"goshort/internal/handler"
	"goshort/internal/persistance"
	"goshort/internal/server"
	"goshort/internal/service"
	"log"
	"net/http"
)

func main() {
	var router = server.NewRouter()
	urlRepo, err := persistance.NewPsqlRepository()
	if err != nil {
		fmt.Printf("error while connecting to the db: %v", err.Error())
		return
	}
	defer urlRepo.DbPool.Close()

	var urlService = service.NewUrlService(urlRepo)
	var urlHandler = handler.NewUrlHandler(urlService)

	router.HandleFunc("/urls", urlHandler.ShortenUrlHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
