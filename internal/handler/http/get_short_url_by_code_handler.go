package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/bsgleison/short-url-by-sdd/internal/application/models"
	"github.com/bsgleison/short-url-by-sdd/internal/application/usecase"
	"github.com/bsgleison/short-url-by-sdd/internal/shared/validation"
)

type GetShortURLByCodeHandler struct {
	getShortURLByCodeUseCase *usecase.GetShortURLByCodeUseCase
}

func NewGetShortURLByCodeHandler(getShortURLByCodeUseCase *usecase.GetShortURLByCodeUseCase) *GetShortURLByCodeHandler {
	return &GetShortURLByCodeHandler{
		getShortURLByCodeUseCase: getShortURLByCodeUseCase,
	}
}

func (h *GetShortURLByCodeHandler) Get(w http.ResponseWriter, r *http.Request) {
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

	writeJSON(w, http.StatusOK, result.Output)
}
