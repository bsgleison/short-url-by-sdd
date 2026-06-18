package usecase

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/bsgleison/short-url-by-sdd/internal/application/models"
	"github.com/bsgleison/short-url-by-sdd/internal/domain/repository"
	"github.com/bsgleison/short-url-by-sdd/internal/shared/validation"
)

var shortCodePattern = regexp.MustCompile(`^[a-z0-9]{7}$`)

type GetShortURLByCodeUseCase struct {
	repo repository.URLRepository
}

func NewGetShortURLByCodeUseCase(repo repository.URLRepository) *GetShortURLByCodeUseCase {
	return &GetShortURLByCodeUseCase{repo: repo}
}

func (uc *GetShortURLByCodeUseCase) Execute(ctx context.Context, input *models.GetShortURLByCodeInput) (*validation.UseCaseResult, error) {
	if input == nil || strings.TrimSpace(input.Code) == "" {
		return validation.NewFailUseCaseResult(
			validation.ValidationErrorCode,
			"code must be informed",
		), nil
	}

	code := strings.ToLower(strings.TrimSpace(input.Code))
	if !shortCodePattern.MatchString(code) {
		return validation.NewFailUseCaseResult(
			validation.ValidationErrorCode,
			"code must be a valid 7-character base32 value",
		), nil
	}

	url, err := uc.repo.FindByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if url == nil {
		return validation.NewFailUseCaseResult(
			validation.UrlCodeNotFoundCode,
			"short URL not found",
		), nil
	}

	response := &models.GetShortURLByCodeResponse{
		ID:          url.ID,
		Code:        url.Code,
		OriginalURL: url.OriginalURL,
		ShortURL:    url.ShortURL,
		Clicks:      url.Clicks,
		UsedAt:      url.UsedAt.Format(time.RFC3339),
		CreatedAt:   url.CreatedAt.Format(time.RFC3339),
	}

	result := validation.NewUseCaseResult()
	result.Output = response
	return result, nil
}
