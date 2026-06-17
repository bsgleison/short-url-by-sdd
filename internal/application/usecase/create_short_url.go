package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/bsgleison/short-url-by-sdd/internal/application/models"
	"github.com/bsgleison/short-url-by-sdd/internal/domain/entity"
	"github.com/bsgleison/short-url-by-sdd/internal/domain/repository"
	"github.com/bsgleison/short-url-by-sdd/internal/shared/validation"
)

type CreateShortURLUseCase struct {
	repo    repository.URLRepository
	baseURL string
}

func NewCreateShortURLUseCase(repo repository.URLRepository, baseURL string) *CreateShortURLUseCase {
	return &CreateShortURLUseCase{
		repo:    repo,
		baseURL: strings.TrimRight(baseURL, "/"),
	}
}

func (uc *CreateShortURLUseCase) Execute(ctx context.Context, input *models.CreateShortURLInput) (*validation.UseCaseResult, error) {
	if input == nil || strings.TrimSpace(input.URL) == "" {
		return validation.NewFailUseCaseResult(
			validation.ValidationErrorCode,
			"URL must be informed",
		), nil
	}

	parsedURL, err := url.ParseRequestURI(input.URL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return validation.NewFailUseCaseResult(
			validation.ValidationErrorCode,
			"URL must be a valid absolute URL",
		), nil
	}

	code, err := uc.generateCode(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	shortURL := fmt.Sprintf("%s/%s", uc.baseURL, code)
	entityURL := &entity.URL{
		ID:          uuid.NewString(),
		Code:        code,
		OriginalURL: parsedURL.String(),
		ShortURL:    shortURL,
		Clicks:      0,
		UsedAt:      now,
		CreatedAt:   now,
	}

	if err := uc.repo.Save(ctx, entityURL); err != nil {
		return nil, err
	}

	response := &models.CreateShortURLResponse{
		ID:          entityURL.ID,
		Code:        entityURL.Code,
		OriginalURL: entityURL.OriginalURL,
		ShortURL:    entityURL.ShortURL,
		Clicks:      entityURL.Clicks,
		UsedAt:      entityURL.UsedAt.Format(time.RFC3339),
		CreatedAt:   entityURL.CreatedAt.Format(time.RFC3339),
	}

	result := validation.NewUseCaseResult()
	result.Output = response
	return result, nil
}

func (uc *CreateShortURLUseCase) generateCode(ctx context.Context) (string, error) {
	for i := 0; i < 10; i++ {
		bytes := make([]byte, 5)
		if _, err := rand.Read(bytes); err != nil {
			return "", err
		}

		encoded := base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString(bytes)
		code := strings.ToLower(encoded[:7])

		existing, err := uc.repo.FindByCode(ctx, code)
		if err != nil {
			return "", err
		}

		if existing == nil {
			return code, nil
		}
	}

	return "", fmt.Errorf("could not generate a unique short code")
}
