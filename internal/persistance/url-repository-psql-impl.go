package persistance

import (
	"context"
	"goshort/internal/domain/models"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UrlRepositoryPsqlImpl struct {
	DbPool *pgxpool.Pool
}

func NewPsqlRepository() (*UrlRepositoryPsqlImpl, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("PSQL_URL"))
	if err != nil {
		return nil, err
	}
	return &UrlRepositoryPsqlImpl{DbPool: pool}, nil
}

func (u *UrlRepositoryPsqlImpl) Save(url models.Url) (string, error) {
	return "", nil
}

func (u *UrlRepositoryPsqlImpl) Get(short string) (string, error) {
	return "", nil
}
