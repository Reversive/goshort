package persistance

import "goshort/internal/domain/models"

type UrlRepository interface {
	Save(url models.Url) (string, error)
	Get(short string) (string, error)
}
