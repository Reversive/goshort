package service

import (
	"encoding/base32"
	"errors"
	"goshort/internal/domain/models"
	"goshort/internal/persistance"
	"math/big"
	"sync"
)

type UrlService struct {
	UrlRepo persistance.UrlRepository
	counter *big.Int
	mtx     sync.RWMutex
}

func NewUrlService(repo persistance.UrlRepository) *UrlService {
	return &UrlService{
		UrlRepo: repo,
		counter: big.NewInt(0),
		mtx:     sync.RWMutex{},
	}
}

var ErrUrlCannotBeEmpty = errors.New("url field cannot be empty")

func (s *UrlService) HandlePostUrl(
	url string,
) (string, error) {
	if url == "" {
		return "", ErrUrlCannotBeEmpty
	}
	s.mtx.Lock()
	s.counter = s.counter.Add(s.counter, big.NewInt(1))
	b := s.counter.Bytes()
	encoding := base32.StdEncoding.WithPadding(base32.NoPadding)
	str := encoding.EncodeToString(b)
	s.mtx.Unlock()
	return s.UrlRepo.Save(models.Url{Original: url, Short: str})
}
