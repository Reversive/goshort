package persistance

import (
	"errors"
	"goshort/internal/domain/models"
	"sync"
)

var ErrNotFound = errors.New("short code not found")

type UrlRepositoryStackImpl struct {
	Store map[string]string
	Mtx   sync.RWMutex
}

func NewStackRepository() *UrlRepositoryStackImpl {
	return &UrlRepositoryStackImpl{Store: make(map[string]string)}
}

func (u *UrlRepositoryStackImpl) Save(url models.Url) (string, error) {
	u.Mtx.Lock()
	defer u.Mtx.Unlock()
	u.Store[url.Short] = url.Original
	return url.Short, nil
}

func (u *UrlRepositoryStackImpl) Get(short string) (string, error) {
	u.Mtx.RLock()
	defer u.Mtx.RUnlock()
	original, ok := u.Store[short]
	if !ok {
		return "", ErrNotFound
	}
	return original, nil
}
