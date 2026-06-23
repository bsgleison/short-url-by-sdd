package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/bsgleison/short-url-by-sdd/internal/application/models"
	"github.com/bsgleison/short-url-by-sdd/internal/domain/entity"
)

func (f *fakeURLRepository) Save(_ context.Context, url *entity.URL) error {
	if url == nil {
		return nil
	}
	f.items[url.Code] = url
	return nil
}

func (f *fakeURLRepository) FindByCode(_ context.Context, code string) (*entity.URL, error) {
	for _, item := range f.items {
		if item.Code == code {
			return item, nil
		}
	}
	return nil, nil
}

func TestURLClickedUseCaseRejectsInvalidInput(t *testing.T) {
	repo := newFakeURLRepository()
	uc := NewURLClickedUseCase(repo)

	err := uc.Execute(context.Background(), nil)
	if err != nil {
		t.Fatalf("expected nil error for nil input, got %v", err)
	}

	msg := &models.URLClickedMessage{Code: ""}
	err = uc.Execute(context.Background(), msg)
	if err != nil {
		t.Fatalf("expected nil error for empty code, got %v", err)
	}
}

func TestURLClickedUseCaseIncrementsClicks(t *testing.T) {
	repo := newFakeURLRepository()
	uc := NewURLClickedUseCase(repo)

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

	err := uc.Execute(context.Background(), &models.URLClickedMessage{Code: "abc1234"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	updated := repo.items["abc1234"]
	if updated.Clicks != 3 {
		t.Fatalf("expected clicks to be 3, got %d", updated.Clicks)
	}

	if updated.UsedAt.Before(now) {
		t.Fatalf("expected UsedAt to be updated, got %v", updated.UsedAt)
	}
}

func TestURLClickedUseCaseDiscardsNotFound(t *testing.T) {
	repo := newFakeURLRepository()
	uc := NewURLClickedUseCase(repo)

	err := uc.Execute(context.Background(), &models.URLClickedMessage{Code: "nonexistent"})
	if err != nil {
		t.Fatalf("expected nil error for not found, got %v", err)
	}
}
