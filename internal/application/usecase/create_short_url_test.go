package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/bsgleison/short-url-by-sdd/internal/application/models"
	"github.com/bsgleison/short-url-by-sdd/internal/domain/entity"
)

type fakeURLRepository struct {
	items map[string]*entity.URL
}

func newFakeURLRepository() *fakeURLRepository {
	return &fakeURLRepository{items: make(map[string]*entity.URL)}
}

func TestCreateShortURLUseCaseRejectsInvalidURL(t *testing.T) {
	repo := newFakeURLRepository()
	uc := NewCreateShortURLUseCase(repo, "http://short.com")

	result, err := uc.Execute(context.Background(), &models.CreateShortURLInput{URL: "not-a-valid-url"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if !result.HasError {
		t.Fatal("expected validation error for invalid URL")
	}
}

func TestCreateShortURLUseCaseStoresShortURL(t *testing.T) {
	repo := newFakeURLRepository()
	uc := NewCreateShortURLUseCase(repo, "http://short.com")

	result, err := uc.Execute(context.Background(), &models.CreateShortURLInput{URL: "https://example.com/path?q=1"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if result.HasError {
		t.Fatalf("expected success, got error result: %+v", result.Messages)
	}

	if result.Output == nil {
		t.Fatal("expected output to be set")
	}

	output, ok := result.Output.(*models.CreateShortURLResponse)
	if !ok {
		t.Fatalf("expected response type, got %T", result.Output)
	}

	if output.Code == "" || len(output.Code) != 7 {
		t.Fatalf("expected code length 7, got %q", output.Code)
	}

	if output.ShortURL != "http://short.com/"+output.Code {
		t.Fatalf("expected short URL %q, got %q", "http://short.com/"+output.Code, output.ShortURL)
	}

	if output.OriginalURL != "https://example.com/path?q=1" {
		t.Fatalf("expected original URL preserved, got %q", output.OriginalURL)
	}

	createdAt, err := time.Parse(time.RFC3339, output.CreatedAt)
	if err != nil {
		t.Fatalf("expected valid created timestamp, got %v", err)
	}

	if time.Since(createdAt) > time.Minute {
		t.Fatal("expected created time to be recent")
	}
}
