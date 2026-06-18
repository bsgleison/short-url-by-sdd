package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/bsgleison/short-url-by-sdd/internal/application/models"
	"github.com/bsgleison/short-url-by-sdd/internal/domain/entity"
)

func TestGetShortURLByCodeUseCaseRejectsInvalidCode(t *testing.T) {
	repo := newFakeURLRepository()
	uc := NewGetShortURLByCodeUseCase(repo)

	result, err := uc.Execute(context.Background(), &models.GetShortURLByCodeInput{Code: "invalid"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if !result.HasError {
		t.Fatal("expected validation error for invalid code")
	}
}

func TestGetShortURLByCodeUseCaseReturnsStoredURL(t *testing.T) {
	repo := newFakeURLRepository()
	now := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	stored := &entity.URL{
		ID:          "id-1",
		Code:        "abc1234",
		OriginalURL: "https://example.com",
		ShortURL:    "http://short.com/abc1234",
		Clicks:      2,
		UsedAt:      now,
		CreatedAt:   now,
	}
	if err := repo.Save(context.Background(), stored); err != nil {
		t.Fatalf("expected save to succeed, got %v", err)
	}

	uc := NewGetShortURLByCodeUseCase(repo)
	result, err := uc.Execute(context.Background(), &models.GetShortURLByCodeInput{Code: "abc1234"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if result.HasError {
		t.Fatalf("expected success, got error result: %+v", result.Messages)
	}

	output, ok := result.Output.(*models.GetShortURLByCodeResponse)
	if !ok {
		t.Fatalf("expected response type, got %T", result.Output)
	}

	if output.ID != stored.ID || output.Code != stored.Code || output.OriginalURL != stored.OriginalURL ||
		output.ShortURL != stored.ShortURL || output.Clicks != stored.Clicks ||
		output.UsedAt != stored.UsedAt.Format(time.RFC3339) || output.CreatedAt != stored.CreatedAt.Format(time.RFC3339) {
		t.Fatalf("unexpected output: %+v", output)
	}
}
