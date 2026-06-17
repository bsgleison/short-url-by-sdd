package repository

import (
	"context"

	"github.com/bsgleison/short-url-by-sdd/internal/domain/entity"
)

type URLRepository interface {
	Save(ctx context.Context, url *entity.URL) error
	FindByCode(ctx context.Context, code string) (*entity.URL, error)
}
