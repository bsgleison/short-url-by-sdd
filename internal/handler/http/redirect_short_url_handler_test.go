package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/bsgleison/short-url-by-sdd/internal/application/usecase"
	"github.com/bsgleison/short-url-by-sdd/internal/domain/entity"
	"github.com/bsgleison/short-url-by-sdd/internal/domain/repository"
)

type fakeURLRepositoryForRedirect struct {
	items map[string]*entity.URL
}

func newFakeURLRepositoryForRedirect() *fakeURLRepositoryForRedirect {
	return &fakeURLRepositoryForRedirect{items: make(map[string]*entity.URL)}
}

func (f *fakeURLRepositoryForRedirect) Save(_ context.Context, url *entity.URL) error {
	f.items[url.ID] = url
	return nil
}

func (f *fakeURLRepositoryForRedirect) FindByCode(_ context.Context, code string) (*entity.URL, error) {
	for _, item := range f.items {
		if item.Code == code {
			return item, nil
		}
	}
	return nil, nil
}

var _ repository.URLRepository = (*fakeURLRepositoryForRedirect)(nil)

type fakeURLClickedPublisher struct {
	publishedCodes []string
}

func (f *fakeURLClickedPublisher) Publish(_ context.Context, code string, id string) error {
	f.publishedCodes = append(f.publishedCodes, code)
	return nil
}

func TestRedirectShortURLHandlerRedirectsAndPublishesEvent(t *testing.T) {
	repo := newFakeURLRepositoryForRedirect()
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

	publisher := &fakeURLClickedPublisher{}
	getUseCase := usecase.NewGetShortURLByCodeUseCase(repo)
	handler := NewRedirectShortURLHandler(getUseCase, publisher)

	req := httptest.NewRequest(http.MethodGet, "/abc1234", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("code", stored.Code)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	handler.Redirect(rr, req)

	if rr.Code != http.StatusMovedPermanently {
		t.Fatalf("expected redirect status %d, got %d", http.StatusMovedPermanently, rr.Code)
	}

	if rr.Header().Get("Location") != stored.OriginalURL {
		t.Fatalf("expected redirect location %q, got %q", stored.OriginalURL, rr.Header().Get("Location"))
	}

	if len(publisher.publishedCodes) != 1 || publisher.publishedCodes[0] != stored.Code {
		t.Fatalf("expected event to publish code %q, got %#v", stored.Code, publisher.publishedCodes)
	}
}
