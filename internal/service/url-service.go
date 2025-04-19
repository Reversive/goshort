package service

import (
	"errors"
	"goshort/internal/domain/models"
	"goshort/internal/persistance"
)

type UrlService struct {
	UrlRepo persistance.UrlRepository
}

var ErrUrlCannotBeEmpty = errors.New("url field cannot be empty")

func (s *UrlService) HandlePostUrl(
	url string,
) (string, error) {
	if url == "" {
		return "", ErrUrlCannotBeEmpty
	}
	// TODO: Logic for shortening url
	return s.UrlRepo.Save(models.Url{Original: url, Short: "temp"})
}
