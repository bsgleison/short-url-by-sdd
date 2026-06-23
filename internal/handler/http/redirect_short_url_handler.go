package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/bsgleison/short-url-by-sdd/internal/application/models"
	"github.com/bsgleison/short-url-by-sdd/internal/application/usecase"
	"github.com/bsgleison/short-url-by-sdd/internal/shared/validation"
)

type URLClickedPublisher interface {
	Publish(ctx context.Context, code string, id string) error
}

type RedirectShortURLHandler struct {
	getShortURLByCodeUseCase *usecase.GetShortURLByCodeUseCase
	publisher                URLClickedPublisher
}

func NewRedirectShortURLHandler(
	getShortURLByCodeUseCase *usecase.GetShortURLByCodeUseCase,
	publisher URLClickedPublisher,
) *RedirectShortURLHandler {
	return &RedirectShortURLHandler{
		getShortURLByCodeUseCase: getShortURLByCodeUseCase,
		publisher:                publisher,
	}
}

func (h *RedirectShortURLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	input := models.GetShortURLByCodeInput{Code: code}

	result, err := h.getShortURLByCodeUseCase.Execute(r.Context(), &input)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if result.HasError {
		status := http.StatusBadRequest
		for _, message := range result.Messages {
			if message != nil && message.Code == validation.UrlCodeNotFoundCode {
				status = http.StatusNotFound
				break
			}
		}
		writeJSON(w, status, result)
		return
	}

	response, ok := result.Output.(*models.GetShortURLByCodeResponse)
	if !ok {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "unexpected response format"})
		return
	}

	messageId := uuid.NewString()

	if err := h.publisher.Publish(r.Context(), response.Code, messageId); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	http.Redirect(w, r, response.OriginalURL, http.StatusMovedPermanently)
}
