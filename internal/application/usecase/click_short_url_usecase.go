package usecase

import (
	"context"
	"time"

	"github.com/bsgleison/short-url-by-sdd/internal/application/models"
	"github.com/bsgleison/short-url-by-sdd/internal/domain/repository"
)

type URLClickedUseCase struct {
	repo repository.URLRepository
}

func NewURLClickedUseCase(repo repository.URLRepository) *URLClickedUseCase {
	return &URLClickedUseCase{repo: repo}
}

func (uc *URLClickedUseCase) Execute(ctx context.Context, input *models.URLClickedMessage) error {
	if input == nil || input.Code == "" {
		return nil
	}

	url, err := uc.repo.FindByCode(ctx, input.Code)
	if err != nil {
		return err
	}

	if url == nil {
		return nil
	}

	url.Clicks++
	url.UsedAt = time.Now()

	return uc.repo.Save(ctx, url)
}
